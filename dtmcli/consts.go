package dtmcli

// Status transaction status
type Status = string

const (
	StatusPrepared  Status = "prepared"  // transaction is in prepared
	StatusSubmitted Status = "submitted" // transaction is submmited
	StatusSucceed   Status = "succeed"   // transaction has succeeded
	StatusFailed    Status = "failed"    // transaction has failed
)

// Branch transaction branch
type Branch = string

const (
	BranchTry     Branch = "try"     // TCC try branch
	BranchConfirm Branch = "confirm" // TCC confirm branch
	BranchCancel  Branch = "cancel"  // TCC cancel branch

	BranchAction     Branch = "action"     // Message, Saga, XA transaction action branch
	BranchCompensate Branch = "compensate" // Saga transaction compensate branch
	BranchCommit     Branch = "commit"     // XA transaction commit branch
	BranchRollback   Branch = "rollback"   // XA transaction rollback branch
)

// Result transaction result
type Result = string

const (
	ResultSuccess Result = "SUCCESS" // transaction branch success
	ResultFailure Result = "FAILURE" // transaction branch failed
)
