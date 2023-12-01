package client

func Client(apiId int, apiHash string) (*TGoCall, error) {

	c := &TGoCall{
		ApiID:   apiId,
		ApiHash: apiHash,
	}

	return c, nil
}
