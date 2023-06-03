var (
    rdb = &RedisClient{} // this variable will make the RedisClient global visibile in the 'main' scope
)

func main() {
    ctx := context.Background() // contexts in go

    err := rdb.InitClient(ctx, "redis:6379", "")
    if err != nil {
       errMsg := fmt.Sprintf("failed to init Redis client, err: %v", err)
       log.Fatal(errMsg)
    }

    r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
    if err != nil {
       log.Fatal(err)
    }

    svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
       ServiceName: "demo.rpc.server",
    }))

    err = svr.Run()
    if err != nil {
       log.Println(err.Error())
    }
}
