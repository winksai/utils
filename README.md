# utils
在git里拉取代码 go get github.com/winksai/utils/utils

             
                func InitRdb() {
                global.Rdb = redis.NewClient(&redis.Options{
                Addr:     "",
                Password: "", // no password set
                DB:       0,   // use default DB
                })
                    pong, err := global.Rdb.Ping().Result()
                    if err != nil {
                        panic(err)
                    }
                }
            
                func InitDB() {
                dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local")
                db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
                Logger: logger.Default.LogMode(logger.Info),
                })
                
                    sqlDB, err := db.DB()
                    // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
                    sqlDB.SetMaxIdleConns(10)
                    // SetMaxOpenConns sets the maximum number of open connections to the database.
                    sqlDB.SetMaxOpenConns(100)
                    // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
                    sqlDB.SetConnMaxLifetime(time.Hour)

                    if err != nil {
                        panic(err)
                    }
                }
              
                    func InitConfig() {
                    viper.SetConfigFile("")
                    err := viper.ReadInConfig()
                    if err != nil {
                    panic(err)
                    }
                    viper.Unmarshal()
                    }
               
                func InitZap() {
                // 创建日志目录，确保日志目录存在
                // 使用 os.MkdirAll 来创建目录，0777 表示最大权限
                os.MkdirAll("指定目录", 0777)

                // 配置 Zap 日志库
                config := zap.NewDevelopmentConfig() // 创建开发模式的 Zap 配置（默认日志级别为 debug）
                
                // 修改输出路径，将日志输出到指定文件

                config.OutputPaths = []string{
                "", // 只输出到文件
                "stdout",          // 输出到控制台
                }
                config.ErrorOutputPaths = []string{
                "", // 错误日志也输出到指定文件
                }
            
                // 构建并返回一个 Zap logger 实例
                // 该 logger 实例会根据上面的配置输出日志
                build, _ := config.Build()
            
                // 使用 ReplaceGlobals 将创建的 logger 替换为全局日志记录器
                zap.ReplaceGlobals(build)
            
                // 控制台输出会自动生效
                // 在使用 `zap` 打印日志时，会输出到指定的文件以及控制台
                }








