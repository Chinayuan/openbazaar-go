package pstoremem

import (
	"sync"

	peer "gx/ipfs/QmTRhk7cgjUf2gfQ3p2M9KPECNZEW9XUrmHcFCgog4cPgB/go-libp2p-peer"
	pstore "gx/ipfs/QmTTJcDL3gsnGDALjh2fDGg1onGRUdVgNL2hU2WEZcVrMX/go-libp2p-peerstore"
)

type memoryPeerMetadata struct {
	// store other data, like versions
	//ds ds.ThreadSafeDatastore
	ds     map[string]interface{}
	dslock sync.Mutex
}

var _ pstore.PeerMetadata = (*memoryPeerMetadata)(nil)

func NewPeerMetadata() pstore.PeerMetadata {
	return &memoryPeerMetadata{
		ds: make(map[string]interface{}),
	}
}

func (ps *memoryPeerMetadata) Put(p peer.ID, key string, val interface{}) error {
	//dsk := ds.NewKey(string(p) + "/" + key)
	//return ps.ds.Put(dsk, val)
	ps.dslock.Lock()
	defer ps.dslock.Unlock()
	ps.ds[string(p)+"/"+key] = val
	return nil
}

func (ps *memoryPeerMetadata) Get(p peer.ID, key string) (interface{}, error) {
	//dsk := ds.NewKey(string(p) + "/" + key)
	//return ps.ds.Get(dsk)

	ps.dslock.Lock()
	defer ps.dslock.Unlock()
	i, ok := ps.ds[string(p)+"/"+key]
	if !ok {
		return nil, pstore.ErrNotFound
	}
	return i, nil
}
