package wompatti

type WompattiServer struct {
	NewContext func() *WompattiContext
}

func CreateWompattiServer(
	wompattiServerPort string,
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
	dev bool,
) {
	contextGenerator := NewContextGenerator(
		dbUser,
		dbPass,
		dbIp,
		dbPort,
		dbName,
	)

	CreateWompattiService(contextGenerator.NewContext, wompattiServicePort)
}
