package gcp

func GCS_OS_C3_TR01() {
	return
}

func GCS_OS_C3_TR02() {
	return
}

func GCS_OS_C3_TR02() {
	err := CCC_OS_C3_TR02_T01()
	if err != nil {
		return err
	}
	CCC_OS_C3_TR02_T02()
	if err != nil {
		return err
	}
	CCC_OS_C3_TR02_T03()
	if err != nil {
		return err
	}
	return
}

func CCC_OS_C3_TR02_T01() error {
	// Non-admin user should not be able to create a bucket
	err := createBucket()
	if err != nil {
		// ensure error is due to permissions (expected)
		// if not, fail()
	}
	return pass()
}
