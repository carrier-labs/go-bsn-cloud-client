// Package models contains shared data structures for the BSN.Cloud API client.
package models

// PlayerFamily is an enumeration of all supported player families in BSN.Cloud.
type PlayerFamily string

const (
	// PlayerFamilyTiger represents the Tiger family.
	PlayerFamilyTiger PlayerFamily = "Tiger"
	// PlayerFamilyPantera represents the Pantera family.
	PlayerFamilyPantera PlayerFamily = "Pantera"
	// PlayerFamilyImpala represents the Impala family.
	PlayerFamilyImpala PlayerFamily = "Impala"
	// PlayerFamilyMalibu represents the Malibu family.
	PlayerFamilyMalibu PlayerFamily = "Malibu"
	// PlayerFamilyPagani represents the Pagani family.
	PlayerFamilyPagani PlayerFamily = "Pagani"
	// PlayerFamilySebring represents the Sebring family.
	PlayerFamilySebring PlayerFamily = "Sebring"
	// PlayerFamilyRaptor represents the Raptor family.
	PlayerFamilyRaptor PlayerFamily = "Raptor"
	// PlayerFamilyCobra represents the Cobra family.
	PlayerFamilyCobra PlayerFamily = "Cobra"
	// PlayerFamilyUnknown represents an unknown family.
	PlayerFamilyUnknown PlayerFamily = "Unknown"
)
