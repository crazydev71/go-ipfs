package coreunix

import (
	core "github.com/ipfs/go-ipfs/core"
	dag "gx/ipfs/QmSRe5UvVPLJ6LAtVH9gQZwEL4nbck5b5zNe4MChh3LJHk/go-merkledag"
	ft "gx/ipfs/QmcLmU3rW9jmxPf2SZuvS15Dhtji4TivmuEca58HTdPRBm/go-unixfs"
	cid "gx/ipfs/QmcRoKTXnq18qQRZFa4jWwWvMQkxzWRpxCwcpCCFgnLUGi/go-cid"
)

func AddMetadataTo(n *core.IpfsNode, skey string, m *ft.Metadata) (string, error) {
	c, err := cid.Decode(skey)
	if err != nil {
		return "", err
	}

	nd, err := n.DAG.Get(n.Context(), c)
	if err != nil {
		return "", err
	}

	mdnode := new(dag.ProtoNode)
	mdata, err := ft.BytesForMetadata(m)
	if err != nil {
		return "", err
	}

	mdnode.SetData(mdata)
	if err := mdnode.AddNodeLink("file", nd); err != nil {
		return "", err
	}

	err = n.DAG.Add(n.Context(), mdnode)
	if err != nil {
		return "", err
	}

	return mdnode.Cid().String(), nil
}

func Metadata(n *core.IpfsNode, skey string) (*ft.Metadata, error) {
	c, err := cid.Decode(skey)
	if err != nil {
		return nil, err
	}

	nd, err := n.DAG.Get(n.Context(), c)
	if err != nil {
		return nil, err
	}

	pbnd, ok := nd.(*dag.ProtoNode)
	if !ok {
		return nil, dag.ErrNotProtobuf
	}

	return ft.MetadataFromBytes(pbnd.Data())
}
