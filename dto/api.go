package dto

type StartSessionResponse struct {
	// Nonce authorizes a Websocket connection to be opened.
	// Nonce expires in 10 seconds.
	Nonce string `json:"nonce"`
}

type ConnectRequest struct {
	Nonce string `json:"nonce"`

	// Username is a keyring. Client can declare a Username, otherwise server will
	// allocate one based on config.
	// Username is actually a namespace on the client. It allows the same inboxes to be
	// reused on another server without a collision in cookie storage.
	Username string `json:"username"`
	// Password is used to salt a server-side hash for client-side encryption.
	// See EncryptionKey below.
	Password string `json:"secret"`

	// Inboxes are vanity addresses the user may want to claim.
	// The server host needs to check if those addresses are currently in use.
	// Authentic claims will evict existing sockets for those inboxes.
	Inboxes []string `json:"inboxes"`
}

type ConnectResponse struct {
	Success bool `json:"success"`

	// SessionID is used for keepalives
	SessionID string

	// Key is a crypto key for clients to use for encryption-at-rest.
	// The key is hash(username, password, secret).
	Key string `json:"key"`

	// AllowedInboxes is a subset of Inboxes.
	AllowedInboxes []string `json:"inboxes"`
}
