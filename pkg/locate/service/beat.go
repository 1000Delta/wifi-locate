package service

type BeatReq struct{}

type BeatResp struct{}

// Beat used to check rpc service
func Beat(req BeatReq, resp *BeatResp) error {
	return nil
}