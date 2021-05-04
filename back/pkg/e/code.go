package e

const (
    SUCCESS                                   = 200
    InvalidParams                             = 400
    ERROR                                     = 500
    VerificationFailedPleaseTryAgain          = 4000
    VerificationHasFailedPleaseReAcquire      = 4001
    VerificationFrequentFailedPleaseReAcquire = 4002

    ErrorAuthCheckTokenFail    = 20001
    ErrorAuthCheckTokenTimeout = 20002
    ErrorAuthToken             = 20003

    ErrorAuth                  = 20004
)
