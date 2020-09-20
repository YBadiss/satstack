package bus

// Status indicates the state of LSS with regards to the readiness of the
// connected Bitcoin Core node.
type Status string

const (
	// Initializing is a Status to indicate the initial state of LSS, while it
	// is warming up. This is typically the case when LSS is launched, until
	// any new Status is assigned.
	Initializing Status = "initializing"

	// Ready is a Status to indicate that LSS is ready to accept explorer API
	// requests from Ledger Live.
	Ready Status = "ready"

	// Syncing is a Status to indicate that the Bitcoin Core node is currently
	// downloading and validating blocks.
	Syncing Status = "syncing"

	// Scanning is a Status to indicate that the Bitcoin Core node is currently
	// importing account descriptors into its wallet.
	Scanning Status = "scanning"
)
