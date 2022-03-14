package tls

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/http2"
)

func TestClientTLS(t *testing.T) {
	ts := httptest.NewTLSServer( // 1
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.TLS == nil { // 2
					u := "https://" + r.Host + r.RequestURI
					http.Redirect(w, r, u, http.StatusMovedPermanently)
					return
				}
				w.WriteHeader(http.StatusOK)
			},
		),
	)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL) // 3
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; actual status %d",
			http.StatusOK, resp.StatusCode)
	}

	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256}, // 4
			MinVersion:       tls.VersionTLS12,
		},
	}

	err = http2.ConfigureTransport(tp) // 5
	if err != nil {
		t.Fatal(err)
	}

	client2 := &http.Client{Transport: tp}

	_, err = client2.Get(ts.URL)
	if err == nil || !strings.Contains(err.Error(), "certificate signed by unknown authority") {
		t.Fatalf("expected unknown authority error; actual: %q", err)
	}

	tp.TLSClientConfig.InsecureSkipVerify = true // 6

	resp, err = client2.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; actual status %d", http.StatusOK, resp.StatusCode)
	}
}

/*
1. httptest.NewTLSServer
	- HTTPS 서버를 반환
	- 새로운 인증서 생성을 포함하여 HTTPS 서버 초기화를 위한 TLS 세부 환경구성까지 처리해 줌
	- 아직 신뢰할 수 없는 상태

2. if r.TLS == nil {...}
	- 서버에서 HTTP로 클라이언트의 요청을 받으면 요청 객체의 TLS 필드는 nil이 됨
	- 클라이언트의 요청을 HTTPS로 리다이렉트 시킬 수 있음

3. ts.Client().Get(ts.URL)
	- 서버 객체의 Client 메소드는 서버의 인증서를 신뢰하는 *http.Client 객체를 반환함
	- 이 클라이언트를 이용하여 핸들러 내의 TLS와 관련된 코드를 테스트할 수 있음

4. []tls.CurveID{tls.CurveP256}
	- 새로운 트랜스포트를 생성하여 TLS 구성을 정의하는 부분
	- 이 트랜스포트를 사용하도록 http2를 구성한 뒤 클라이언트 트랜스포트의 기본 TLS 구성을 오버라이딩 함
	- 클라이언트의 TLS 구성 설정의 CurvePreferences 필드 값은 P-256으로 설정하는 것이 좋음
		- P-256은 P-384나 P-521과는 달리 시간차 공격에 저항이 있음
		- P-256을 사용하면 클라이언트는 TLS 협상(negotiation)에 최소한 1.2 이상의 버전을 사용함
			- P-256, P-384, P-521 : 미국 국립표준기술연구소의 디지털 서명 표준

5. http2.ConfigureTransport(tp)
	- 트랜스포트는 더 이상 기본 TLS 구성을 사용하지 않기에 클라이언트는 HTTP/2를 기본적으로 지원하지 않음
	- HTTP/2를 사용하려면 명시적으로 HTTP/2를 사용하기 위한 함수에 트랜스포트를 전달해주어야 함
	- HTTP/2를 사용하진 않지만 트랜스포트의 TLS 구성을 오버라이딩할 경우 HTTP/2 지원이 제거됨

6. tp.TLSClientConfig.InsecureSkipVerify = true
	- 명시적으로 신뢰할 인증서를 선택하지 않으면 클라이언트는 운영체제가 신뢰하는 인증 저장소의 인증서를 신뢰함
	- 테스트 서버로 보내는 첫 번째 요청은 클라이언트가 서버가 보내는 인증서의 서명자를 신뢰하지 않기 때문에 실패하여 에러가 발생함
	- 이를 우회하기 위해 InsecureSkipVerify 필드의 값을 true로 설정하여 클라이언트의 트랜스포트가 서버의 인증서를 검증하지 않도록 할 수 있음
	- 보안상의 이유로 좋지 않은 방법이며 인증서 고정(certificate pinning)이 더 나은 방법임

*/
