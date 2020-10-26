// Copyright © 2019 Leonardo Javier Gago <ljgago@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server // import "github.com/ljgago/adbus/cmd/server"

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ljgago/adbus/cmd"
	pb "github.com/ljgago/adbus/pkg/api/v1"
	"github.com/ljgago/adbus/pkg/config"
	"github.com/ljgago/adbus/pkg/log"
)

const logName string = ".adbus/adbus-server.log"

// commandDefine represents the subcommand
var commandDefine = &cobra.Command{
	Use:   "server [flags]",
	Short: "Start the Proxy Server",
	Long: `
Start the Proxy Server mode.

The proxy server expose an REST API service for interact 
with a pub-sub message broker and it send commands
for communicate with the remote devices.`,
	DisableAutoGenTag: true,
	Run: func(command *cobra.Command, args []string) {
		//	options, err := configAPI()
		//	if err != nil {
		//		log.Fatal().Str("type", "server").Err(err).Msg("")
		//	}
		runServer(args)
	},
}

func init() {
	cmd.Root.AddCommand(commandDefine)

	flags := commandDefine.Flags()
	flags.String("port-http", "8088", "use port for web clients (format: [port])")
	flags.String("port-grpc", "8089", "use port for grpc clients (format: [port])")
	flags.StringP("broker-url", "b", "127.0.0.1:4222", "broker server address (format: <ip>[:port])")
	flags.String("broker-access-key", "admin", "access key for connect")
	flags.String("broker-secret-key", "123456789", "secret key for connect")
	flags.StringP("storage-url", "s", "127.0.0.1:9000", "url endpoint for s3 node ")
	flags.String("storage-access-key", "admin", "access key for connect")
	flags.String("storage-secret-key", "123456789", "secret key for connect")
	flags.String("database-url", "postgresql://maxroach@localhost:26257/bank?sslmode=disable", "database address")
	// flags.StringP("token", "t", "123456789", "token for connect")
	// flags.Bool("broker-external", false, "use external broker server (default false)")

	viper.BindPFlag("adbus.server.port-http", flags.Lookup("port-http"))
	viper.BindPFlag("adbus.server.port-grpc", flags.Lookup("port-grpc"))
	viper.BindPFlag("adbus.server.broker.url", flags.Lookup("broker-url"))
	viper.BindPFlag("adbus.server.broker.access-key", flags.Lookup("broker-access-key"))
	viper.BindPFlag("adbus.server.broker.secret-key", flags.Lookup("broker-secret-key"))
	viper.BindPFlag("adbus.server.storage.url", flags.Lookup("storage-url"))
	viper.BindPFlag("adbus.server.storage.access-key", flags.Lookup("storage-access-key"))
	viper.BindPFlag("adbus.server.storage.secret-key", flags.Lookup("storage-secret-key"))
	viper.BindPFlag("adbus.server.database", flags.Lookup("database-url"))
	// viper.BindPFlag("adbus.server.broker-external", flags.Lookup("broker-external"))

	viper.BindEnv("adbus.server.port-http", "ADBUS_SERVER_PORT_HTTP")
	viper.BindEnv("adbus.server.port-grpc", "ADBUS_SERVER_PORT_GRPC")
	viper.BindEnv("adbus.server.broker.url", "ADBUS_SERVER_BROKER_URL")
	viper.BindEnv("adbus.server.broker.access-key", "ADBUS_SERVER_BROKER_ACCESS_KEY")
	viper.BindEnv("adbus.server.broker.secret-key", "ADBUS_SERVER_BROKER_SECRET_KEY")
	viper.BindEnv("adbus.server.storage.url", "ADBUS_SERVER_STORAGE_URL")
	viper.BindEnv("adbus.server.storage.access-key", "ADBUS_SERVER_STORAGE_ACCESS_KEY")
	viper.BindEnv("adbus.server.storage.secret-key", "ADBUS_SERVER_STORAGE_SECRET_KEY")
	viper.BindEnv("adbus.server.database", "ADBUS_SERVER_DATABASE_URL")
	//viper.BindEnv("broker-external", "ADBUS_BROKER_EXTERNAL")
	//viper.AllSettings()

}

/*
func configAPI() (*api.Config, error) {
	options := &server.Options{}
	if options.Port = viper.GetString("adbus.server.port"); options.Port == "" {
		return nil, errors.New("'web-port' the web server need a port")
	}
	if options.BrokerURL = viper.GetString("adbus.server.broker-url"); options.BrokerURL == "" {
		return nil, errors.New("'broker-url' the broker server address")
	}
	if options.Endpoint = viper.GetString("adbus.server.endpoint"); options.Endpoint == "" {
		return nil, errors.New("'endpoint' need an endpoint url")
	}
	if options.AccessKey = viper.GetString("adbus.server.access-key"); options.AccessKey == "" {
		return nil, errors.New("'access-key' need access key")
	}
	if options.SecretKey = viper.GetString("adbus.server.secret-key"); options.SecretKey == "" {
		return nil, errors.New("'secret-key' need a secret key")
	}
	options.BrokerExternal = viper.GetBool("adbus.server.broker-external")
	options.Config = viper.GetString("adbus.config")
	log.Init(options.Config, logName)
	return options, nil
}
*/

//func stanInit(natsURL string) {
//	// Set options nats
//	nOpts := &natsd.Options{}
//	url := strings.Split(natsURL, ":")
//	if len(url) == 2 {
//		//opts.Host = url[0]
//		port, err := strconv.Atoi(url[1])
//		if err != nil {
//			nOpts.Port = port
//		}
//	}
//
//	sOpts := stand.GetDefaultOptions()
//	// Force the streaming server to setup its own signal handler
//	sOpts.HandleSignals = true
//	// override the NoSigs for NATS since Streaming has its own signal handler
//	nOpts.NoSigs = true
//	// Without this option set to true, the logger is not configured.
//	//sOpts.EnableLogging = true
//	// This will invoke RunServerWithOpts but on Windows, may run it as a service.
//	log.Info().Str("type", "nats-server").Msg("Listening for NATS devices connections on :" + url[1])
//	if _, err := stand.Run(sOpts, nOpts); err != nil {
//		log.Fatal().Str("type", "stan").Err(err).Msg("")
//		//fmt.Println(err)
//		os.Exit(1)
//	}
//	// Waiting
//	<-make(chan struct{})
//}

func runServer(args []string) {
	// Start the log
	path := viper.GetString("adbus.config")
	mode := viper.GetString("adbus.mode")
	log.Init(path, mode, logName)

	// Load the config
	cfg, err := config.GetServer()
	fmt.Println("Puerto: " + cfg.PortHTTP)

	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}

	// Start the nats client
	reconnectDelay := 10 * time.Second
	// Connect Options.
	brokerOpts := []nats.Option{nats.Name("SERVER")}
	// User / Password
	brokerOpts = append(brokerOpts, nats.UserInfo(cfg.Broker.AccessKey, cfg.Broker.SecretKey))
	// Token security
	//brokerOpts = append(brokerOpts, broker.Token(opts.Token))

	// Delay for reconnect
	brokerOpts = append(brokerOpts, nats.ReconnectWait(reconnectDelay))
	// Max reconnections forever
	brokerOpts = append(brokerOpts, nats.MaxReconnects(-1))
	brokerOpts = append(brokerOpts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Info().Str("type", "server").Msg("Disconnected: trying to reconnect ...")
	}))
	brokerOpts = append(brokerOpts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Info().Str("type", "server").Msgf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	brokerOpts = append(brokerOpts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal().Str("type", "server").Msg("Exiting, no servers available")
	}))

	var conn *nats.Conn
	showlog := true
	for {
		bc, err := nats.Connect(cfg.Broker.URL, brokerOpts...)
		if err != nil {
			if showlog {
				log.Error().Str("type", "server").Err(err).Msg("Desconected: trying to connect ...")
				showlog = false
			}
		} else {
			//log.Info().Str("type", "server").Msgf("Server conected to %s", cfg.Broker.URL)
			fmt.Println("Broker connected to " + cfg.Broker.URL)
			conn = bc
			break
		}
	}

	// Init gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.PortGRPC)
	if err != nil {
		log.Fatal().Str("type", "server").Err(err).Msg("")
	}
	// Creates a new gRPC server with UnaryInterceptor
	grpcServer := grpc.NewServer(
		withServerUnaryInterceptor(),
	)
	reflection.Register(grpcServer)
	pb.RegisterDeviceServiceServer(grpcServer, &deviceService{
		cfg:  cfg,
		conn: conn,
	})
	fmt.Println("gRPC Server API listen on :" + cfg.PortGRPC)
	go grpcServer.Serve(lis)

	// Init HTTP server gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterDeviceServiceHandlerFromEndpoint(ctx, mux, ":"+cfg.PortGRPC, opts)
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}

	fmt.Println("HTTP Server API listen on :" + cfg.PortHTTP)
	err = http.ListenAndServe(":"+cfg.PortHTTP, mux)
	if err != nil {
		log.Error().Str("type", "server").Err(err).Msg("")
	}
	// Waiting
	// <-make(chan struct{})
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}
