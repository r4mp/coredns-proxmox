package proxmox

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

var (
	backend            = os.Getenv("CDNS_PX_BACKEND")
	tokenId            = os.Getenv("CDNS_PX_TOKEN_ID")
	tokenSecret        = os.Getenv("CDNS_PX_TOKEN_SECRET")
	insecure           = os.Getenv("CDNS_PX_INSECURE")
	nodeName           = os.Getenv("CDNS_PX_NODE_NAME")
	vmName             = os.Getenv("CDNS_PX_VM_NAME")
	vmIPv4             = os.Getenv("CDNS_PX_VM_IP_V4")
	awaitedAnswersIPv4 = os.Getenv("CDNS_PX_AWAITED_ANSWERS_IP_V4")
	vmIPv6             = os.Getenv("CDNS_PX_VM_IP_V6")
	awaitedAnswersIPv6 = os.Getenv("CDNS_PX_AWAITED_ANSWERS_IP_V6")
)

//	func TestExample(t *testing.T) {
//		// Create a new Example Plugin. Use the test.ErrorHandler as the next plugin.
//		x := Example{Next: test.ErrorHandler()}
//
//		// Setup a new output buffer that is *not* standard output, so we can check if
//		// example is really being printed.
//		b := &bytes.Buffer{}
//		golog.SetOutput(b)
//
//		ctx := context.TODO()
//		r := new(dns.Msg)
//		r.SetQuestion("example.org.", dns.TypeA)
//		// Create a new Recorder that captures the result, this isn't actually used in this test
//		// as it just serves as something that implements the dns.ResponseWriter interface.
//		rec := dnstest.NewRecorder(&test.ResponseWriter{})
//
//		// Call our plugin directly, and check the result.
//		x.ServeDNS(ctx, rec, r)
//		if a := b.String(); !strings.Contains(a, "[INFO] plugin/example: example") {
//			t.Errorf("Failed to print '%s', got %s", "[INFO] plugin/example: example", a)
//		}
//	}
func TestGetNodeNames(t *testing.T) {
	//t.Log(nodes)

	pve := Proxmox{Backend: backend, TokenId: tokenId, TokenSecret: tokenSecret, Insecure: insecure}
	info, err := pve.GetNodes()

	if err != nil {
		t.Error(err)
	}
	for _, node := range info {
		t.Log(node.Node)
	}
}

func TestGetVMNames(t *testing.T) {
	//t.Log(nodes)

	pve := Proxmox{Backend: backend, TokenId: tokenId, TokenSecret: tokenSecret, Insecure: insecure}
	info, err := pve.GetVMs(nodeName)

	if err != nil {
		t.Error(err)
	}
	for _, node := range info {
		t.Log(node.Name)
	}
}

func TestProxmox_GetIPs(t *testing.T) {
	pve := Proxmox{Backend: backend, TokenId: tokenId, TokenSecret: tokenSecret, Insecure: insecure}

	ips, err := pve.GetIPs(nodeName)
	if err != nil {
		t.Error(err)
	}
	for _, ip := range ips {
		t.Log(ip)
	}
}

func TestProxmox_GetIPsByNameIPv4(t *testing.T) {
	pve := Proxmox{Backend: backend, TokenId: tokenId, TokenSecret: tokenSecret, Insecure: insecure}

	ctx := context.TODO()
	r := new(dns.Msg)
	r.SetQuestion(vmName, dns.TypeA)
	// Create a new Recorder that captures the result, this isn't actually used in this test
	// as it just serves as something that implements the dns.ResponseWriter interface.
	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	// Call our plugin directly, and check the result.
	_, err := pve.ServeDNS(ctx, rec, r)
	if err != nil {
		t.Error(err)
	}

	t.Log(rec.Msg)

	iAwaitedAnswersIPv4, err := strconv.Atoi(awaitedAnswersIPv4)
	if err != nil {
		t.Error(err)
	}

	if a := rec.Msg.Answer; len(a) != iAwaitedAnswersIPv4 {
		t.Errorf("Expected %d answer, got %d", iAwaitedAnswersIPv4, len(a))
	}
	if a := rec.Msg.Answer[0].(*dns.A).A.String(); a != vmIPv4 {
		t.Errorf("Expected %s, got %s", vmIPv4, a)
	}

}

func TestProxmox_GetIPsByNameIPv6(t *testing.T) {
	pve := Proxmox{Backend: backend, TokenId: tokenId, TokenSecret: tokenSecret, Insecure: insecure}

	ctx := context.TODO()
	r := new(dns.Msg)
	r.SetQuestion(vmName, dns.TypeAAAA)
	// Create a new Recorder that captures the result, this isn't actually used in this test
	// as it just serves as something that implements the dns.ResponseWriter interface.
	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	// Call our plugin directly, and check the result.
	_, err := pve.ServeDNS(ctx, rec, r)
	if err != nil {
		t.Error(err)
	}

	t.Log(rec.Msg)

	iAwaitedAnswersIPv6, err := strconv.Atoi(awaitedAnswersIPv6)
	if err != nil {
		t.Error(err)
	}

	if aaaa := rec.Msg.Answer; len(aaaa) != iAwaitedAnswersIPv6 {
		t.Errorf("Expected %d answer, got %d", iAwaitedAnswersIPv6, len(aaaa))
	}
	if aaaa := rec.Msg.Answer[0].(*dns.AAAA).AAAA.String(); aaaa != vmIPv6 {
		t.Errorf("Expected %s, got %s", vmIPv6, aaaa)
	}

}
