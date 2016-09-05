package yourtool

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/.well-known/acme-challenge/", letsEncryptHadler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func letsEncryptHadler(w http.ResponseWriter, r *http.Request) {
	key := r.RequestURI[len("/.well-known/acme-challenge/"):]
	switch key {
	case "-3z7rDjrBdyLg-9s5xtZ3qrppdYnLagfox74-qzH1i0":
		fmt.Fprint(w, "-3z7rDjrBdyLg-9s5xtZ3qrppdYnLagfox74-qzH1i0.CwUo3jqD5wgVTw9oysF22fhJ2kel89llDMCGOjGHoH0")
	case "5os0VSOkF6PpcSAKuG7_sneaNaM336jxWv3QPkD7OdA":
		fmt.Fprint(w, "5os0VSOkF6PpcSAKuG7_sneaNaM336jxWv3QPkD7OdA.CwUo3jqD5wgVTw9oysF22fhJ2kel89llDMCGOjGHoH0")
	case "xVQhNG6qDi3Osf90fQ0-OXp69GV3PHVDKP2SvenLRN0":
		fmt.Fprint(w, "xVQhNG6qDi3Osf90fQ0-OXp69GV3PHVDKP2SvenLRN0.CwUo3jqD5wgVTw9oysF22fhJ2kel89llDMCGOjGHoH0")
	case "UChFmDpPgDXW4TlrukX2a2IflhaCj4MmSUxn2eSgdAI":
		fmt.Fprint(w, "UChFmDpPgDXW4TlrukX2a2IflhaCj4MmSUxn2eSgdAI.CwUo3jqD5wgVTw9oysF22fhJ2kel89llDMCGOjGHoH0")
	default:
		fmt.Fprint(w, key)
	}
}
