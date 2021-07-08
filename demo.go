```
func initMysql(conf string) (*sql.DB, error) {
	db, err := sql.Open("mysql", conf)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func monitorit(db *sql.DB) {
	defer db.Close()
	var name, logStatus, logPath string
	getStatus := "SHOW VARIABLES LIKE 'general_log'"
	getPath := "SHOW VARIABLES LIKE 'general_log_file'"
	err := db.QueryRow(getStatus).Scan(&name, &logStatus)
	if err != nil {
		color.Red("获取日志状态失败！")
		return
	}
	if logStatus != "ON" {
		color.White("日志未开启！")
		setStatus := "set global general_log=ON"
		err := db.QueryRow(setStatus).Scan(&name, &logStatus)
		if err != nil && logStatus != "ON" {
			color.Red("开启日志失败！")
			return
		}
	}
	err = db.QueryRow(getPath).Scan(&name, &logPath)
	if err != nil {
		color.Red("获取日志路径失败！")
		return
	}
	printit(logPath)
}
```
