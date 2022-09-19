package errdef


import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeOK-0]
	_ = x[CodeCanceled-1]
	_ = x[CodeUnknown-2]
	_ = x[CodeInvalidArgument-3]
	_ = x[CodeDeadlineExceeded-4]
	_ = x[CodeNotFound-5]
	_ = x[CodeAlreadyExists-6]
	_ = x[CodePermissionDenied-7]
	_ = x[CodeResourceExhausted-8]
	_ = x[CodeFailedPrecondition-9]
	_ = x[CodeAborted-10]
	_ = x[CodeOutOfRange-11]
	_ = x[CodeUnimplemented-12]
	_ = x[CodeInternal-13]
	_ = x[CodeUnavailable-14]
	_ = x[CodeDataLoss-15]
	_ = x[CodeUnauthenticated-16]
}

const _ErrorCode_name = "OKCanceledUnknownInvalidArgumentDeadlineExceededNotFoundAlreadyExistsPermissionDeniedResourceExhaustedFailedPreconditionAbortedOutOfRangeUnimplementedInternalUnavailableDataLossUnauthenticated"

var _ErrorCode_index = [...]uint8{0, 2, 10, 17, 32, 48, 56, 69, 85, 102, 120, 127, 137, 150, 158, 169, 177, 192}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}