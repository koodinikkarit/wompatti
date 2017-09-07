package WompattiContext

func NewContextGenerator(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
	dev bool,
) *ContextGenerator {
	return &ContextGenerator{
		dbUser: dbUser,
		dbPass: dbPass,
		dbIP:   dbIP,
		dbPort: dbPort,
		dbName: dbName,
		dev:    dev,
	}
}
