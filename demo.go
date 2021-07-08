var (
	flagHost   string
	flagPort   int
	flagUser   string
	flagPwd    string
	flagFilter string
)

func init() {
	flag.StringVar(&flagHost, "h", "127.0.0.1", "host")
	flag.IntVar(&flagPort, "P", 3306, "port")
	flag.StringVar(&flagUser, "u", "root", "user")
	flag.StringVar(&flagPwd, "p", "123456", "password")
	flag.StringVar(&flagFilter, "f", "Query", "filter string")

	flag.Usage = usage

}
func usage() {
	fmt.Fprintf(os.Stderr, `MysqlMonitor by xwuyi.
Usage: MysqlMonitor [-h 127.0.0.1] [-P 3306] [-u root] [-p 123456] [-f "Query"]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql", flagUser, flagPwd, flagHost, flagPort)
	cdb, err := initMysql(dbConfig)
	if err != nil {
		color.Red("MySQL连接失败！")
		return
	}
	color.Green("MySQL连接成功！")
	monitorit(cdb)
}
