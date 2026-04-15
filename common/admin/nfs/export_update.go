//go:build ceph_preview && !(nautilus || octopus)
// +build ceph_preview,!nautilus,!octopus

package nfs

import (
	"github.com/ceph/go-ceph/internal/commands"
)

// UpdateCephFSExport will update an existing NFS export for a CephFS file system.
//
// Similar To:
//
//	ceph nfs export update
func (nfsa *Admin) UpdateCephFSExport(spec CephFSExportSpec) (
	*ExportResult, error) {
	// ---
	f := &cephFSExportFields{
		Prefix:           "nfs export update",
		Format:           "json",
		CephFSExportSpec: spec,
	}
	return parseExportResult(commands.MarshalMgrCommand(nfsa.conn, f))
}
