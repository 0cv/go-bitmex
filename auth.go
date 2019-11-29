package bitmex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/adampointer/go-bitmex/types"
)

type sigParams struct {
	secret, method, path, body string
	expires                    time.Time
}

func (s *sigParams) expiryString() string {
	return fmt.Sprintf("%d", s.expires.Unix())
}

func calculateSignature(params *sigParams) (string, error) {
	raw := fmt.Sprintf("%s%s%d%s", params.method, params.path, params.expires.Unix(), params.body)
	sig := hmac.New(sha256.New, []byte(params.secret))

	if _, err := sig.Write([]byte(raw)); err != nil {
		return "", err
	}
	return hex.EncodeToString(sig.Sum(nil)), nil
}

func websocketAuthCommand(key, secret string) (*types.Command, error) {
	req := &sigParams{
		method:  "GET",
		path:    "/realtime",
		secret:  secret,
		body:    "",
		expires: expiryTime(),
	}
	sig, err := calculateSignature(req)
	if err != nil {
		return nil, err
	}
	cmd := &types.Command{
		Op:   types.CommandOpAuth,
		Args: types.CommandArgs{key, req.expires.Unix(), sig},
	}
	return cmd, nil
}

func expiryTime() time.Time {
	return time.Now().Add(5 * time.Minute)
}
