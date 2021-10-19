package services

import (
	"net"
	"strconv"
	"time"

	"github.com/chrj/smtpd"
	"github.com/fiffu/vprmail/config"
	"github.com/fiffu/vprmail/dto"
)

type Mail struct {
	timestamp time.Time
	origin    smtpd.Peer
	body      smtpd.Envelope
	err       error
}

type SMTPService struct {
	Bus    chan dto.IMessage
	server *smtpd.Server
	Port   string
}

func NewSMTPService(cfg config.Config, bus chan dto.IMessage) SMTPService {
	svc := SMTPService{
		Port: strconv.Itoa(cfg.SMTPPort),
		Bus:  bus,
	}
	svc.server = &smtpd.Server{
		Hostname:          cfg.SMTPHostName,
		WelcomeMessage:    cfg.SMTPWelcomeMsg,
		ConnectionChecker: svc.onConnect(),
		HeloChecker:       svc.onHelo(),
		SenderChecker:     svc.onMailFrom(),
		Handler:           svc.onMail(),
	}
	return svc
}

// go Start()
func (s *SMTPService) Start() error {
	listener, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}
	s.server.Serve(listener)
	return nil
}

func (s *SMTPService) onConnect() func(smtpd.Peer) error {
	return func(peerConnect smtpd.Peer) error {
		return nil
	}
}

func (s *SMTPService) onMailFrom() func(smtpd.Peer, string) error {
	return func(peerSender smtpd.Peer, peerAddress string) error {
		return nil
	}
}

func (s *SMTPService) onHelo() func(smtpd.Peer, string) error {
	return func(peerHelo smtpd.Peer, name string) error {
		return nil
	}
}

func (s *SMTPService) onMail() func(smtpd.Peer, smtpd.Envelope) error {
	return func(peer smtpd.Peer, env smtpd.Envelope) error {
		s.Bus <- &Mail{
			timestamp: time.Now(),
			origin:    peer,
			body:      env,
			err:       nil,
		}
		return nil
	}
}

func (m *Mail) Error() error {
	return m.err
}

func (m *Mail) Timestamp() time.Time {
	return m.timestamp
}

func (m *Mail) Origin() string {
	return m.origin.ServerName
}

func (m *Mail) Sender() string {
	return m.body.Sender
}

func (m *Mail) Recipients() []string {
	return m.body.Recipients
}

func (m *Mail) Contents() []byte {
	return m.body.Data
}
