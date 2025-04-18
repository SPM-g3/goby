// Code generated by Kitex v0.12.2. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	user "github.com/bitdance-panic/gobuy/app/rpc/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUser": kitex.NewMethodInfo(
		getUserHandler,
		newUserServiceGetUserArgs,
		newUserServiceGetUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updateUser": kitex.NewMethodInfo(
		updateUserHandler,
		newUserServiceUpdateUserArgs,
		newUserServiceUpdateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"removeUser": kitex.NewMethodInfo(
		removeUserHandler,
		newUserServiceRemoveUserArgs,
		newUserServiceRemoveUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"blockUser": kitex.NewMethodInfo(
		blockUserHandler,
		newUserServiceBlockUserArgs,
		newUserServiceBlockUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"unblockUser": kitex.NewMethodInfo(
		unblockUserHandler,
		newUserServiceUnblockUserArgs,
		newUserServiceUnblockUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"adminListUser": kitex.NewMethodInfo(
		adminListUserHandler,
		newUserServiceAdminListUserArgs,
		newUserServiceAdminListUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Seller": kitex.NewMethodInfo(
		sellerHandler,
		newUserServiceSellerArgs,
		newUserServiceSellerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserArgs)
	realResult := result.(*user.UserServiceGetUserResult)
	success, err := handler.(user.UserService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return user.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return user.NewUserServiceGetUserResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateUserArgs)
	realResult := result.(*user.UserServiceUpdateUserResult)
	success, err := handler.(user.UserService).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateUserArgs() interface{} {
	return user.NewUserServiceUpdateUserArgs()
}

func newUserServiceUpdateUserResult() interface{} {
	return user.NewUserServiceUpdateUserResult()
}

func removeUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRemoveUserArgs)
	realResult := result.(*user.UserServiceRemoveUserResult)
	success, err := handler.(user.UserService).RemoveUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRemoveUserArgs() interface{} {
	return user.NewUserServiceRemoveUserArgs()
}

func newUserServiceRemoveUserResult() interface{} {
	return user.NewUserServiceRemoveUserResult()
}

func blockUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceBlockUserArgs)
	realResult := result.(*user.UserServiceBlockUserResult)
	success, err := handler.(user.UserService).BlockUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceBlockUserArgs() interface{} {
	return user.NewUserServiceBlockUserArgs()
}

func newUserServiceBlockUserResult() interface{} {
	return user.NewUserServiceBlockUserResult()
}

func unblockUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUnblockUserArgs)
	realResult := result.(*user.UserServiceUnblockUserResult)
	success, err := handler.(user.UserService).UnblockUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUnblockUserArgs() interface{} {
	return user.NewUserServiceUnblockUserArgs()
}

func newUserServiceUnblockUserResult() interface{} {
	return user.NewUserServiceUnblockUserResult()
}

func adminListUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAdminListUserArgs)
	realResult := result.(*user.UserServiceAdminListUserResult)
	success, err := handler.(user.UserService).AdminListUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAdminListUserArgs() interface{} {
	return user.NewUserServiceAdminListUserArgs()
}

func newUserServiceAdminListUserResult() interface{} {
	return user.NewUserServiceAdminListUserResult()
}

func sellerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSellerArgs)
	realResult := result.(*user.UserServiceSellerResult)
	success, err := handler.(user.UserService).Seller(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSellerArgs() interface{} {
	return user.NewUserServiceSellerArgs()
}

func newUserServiceSellerResult() interface{} {
	return user.NewUserServiceSellerResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *user.GetUserReq) (r *user.GetUserResp, err error) {
	var _args user.UserServiceGetUserArgs
	_args.Req = req
	var _result user.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (r *user.UpdateUserResp, err error) {
	var _args user.UserServiceUpdateUserArgs
	_args.Req = req
	var _result user.UserServiceUpdateUserResult
	if err = p.c.Call(ctx, "updateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveUser(ctx context.Context, req *user.RemoveUserReq) (r *user.RemoveUserResp, err error) {
	var _args user.UserServiceRemoveUserArgs
	_args.Req = req
	var _result user.UserServiceRemoveUserResult
	if err = p.c.Call(ctx, "removeUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BlockUser(ctx context.Context, req *user.BlockUserReq) (r *user.BlockUserResp, err error) {
	var _args user.UserServiceBlockUserArgs
	_args.Req = req
	var _result user.UserServiceBlockUserResult
	if err = p.c.Call(ctx, "blockUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnblockUser(ctx context.Context, req *user.UnblockUserReq) (r *user.UnblockUserResp, err error) {
	var _args user.UserServiceUnblockUserArgs
	_args.Req = req
	var _result user.UserServiceUnblockUserResult
	if err = p.c.Call(ctx, "unblockUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AdminListUser(ctx context.Context, req *user.AdminListUserReq) (r *user.AdminListUserResp, err error) {
	var _args user.UserServiceAdminListUserArgs
	_args.Req = req
	var _result user.UserServiceAdminListUserResult
	if err = p.c.Call(ctx, "adminListUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Seller(ctx context.Context, req *user.SellerReq) (r *user.SellerResp, err error) {
	var _args user.UserServiceSellerArgs
	_args.Req = req
	var _result user.UserServiceSellerResult
	if err = p.c.Call(ctx, "Seller", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
