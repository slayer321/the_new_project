proto:
	protoc api_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc user_profile_svc/user_service/pb/*.proto --go_out=. --go-grpc_out=.
