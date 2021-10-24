package hub

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/fiffu/vapormail/dto"
	"github.com/fiffu/vapormail/utils"
)

func (h *HubService) Hash(username, password string) string {
	salt := h.config.ServerSecret
	s := username + salt + password
	md5 := md5.Sum([]byte(s))
	return hex.EncodeToString(md5[:])
}

func (h *HubService) NewInbox() string {
	return utils.RandomName(2, ".")
}

func (h *HubService) NewNonce() string {
	n := utils.RandomUUID()
	h.nonces[n] = time.Now()
	return n
}

func (h *HubService) NewSessionID() string {
	return utils.RandomUUID()
}

func (h *HubService) NewUsername() string {
	return utils.TimestampString()
}

func (h *HubService) ValidateNonce(n string) bool {
	issued, ok := h.nonces[n]
	if !ok {
		return false
	}
	nonceTTL := h.config.ClientHeartbeat
	return time.Since(issued) <= time.Duration(nonceTTL)
}

func (h *HubService) NewSession(
	nonce, username, password string,
	inboxes []string,
) (string, []string, error) {

	if ok := h.ValidateNonce(nonce); !ok {
		return "", nil, errors.New("invalid nonce")
	}
	var allowedInboxes []string
	key := h.Hash(username, password)
	// Examine inbox claims. If any claim has an existing client connection,
	// and this key matches that client's key, evict the old connection.
	for _, i := range inboxes {
		ibx := dto.Inbox(i)
		client, found := h.clients[ibx]
		if !found {
			allowedInboxes = append(allowedInboxes, i)
		} else if key == client.key {
			client.Evict()
			allowedInboxes = append(allowedInboxes, i)
		} else {
			// key didn't match existing client's, so not pushing to acceptedInboxes
			continue
		}
	}
	if len(allowedInboxes) == 0 {
		// If no inboxes, randomly allocate one
		allowedInboxes = append(allowedInboxes, h.NewInbox())
	}
	return h.NewSessionID(), allowedInboxes, nil
}
