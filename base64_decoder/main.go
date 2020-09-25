package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {

	encodedString := "LS0tLS0tLS0tLSBGb3J3YXJkZWQgbWVzc2FnZSAtLS0tLS0tLS0NCkZyb206IEFsYWJhbWEgRGVudGlzdCA8bm9yZXBseUAxMDBtYXJrZXRlcnMuY29tPg0KRGF0ZTogVHVlLCBTZXAgMTUsIDIwMjAgYXQgMTI6NDUgQU0NClN1YmplY3Q6IE5ldyBzdWJtaXNzaW9uIGZyb20gQ29udGFjdCBGb3JtDQpUbzogPHJlcG9ydGluZ19hbGFiYW1hLWRlbnRpc3RAMTAwbWFya2V0ZXJzLmNvbT4sIDxkYXJ5bEB3bTRkLmNvbT4NCg0KDQpIZWFkcyBVcCEgQSBOZXcgUGF0aWVudCBIYXMgQ29udGFjdGVkIFlvdQ0KPCNtXy00NjQ4NTY1NzEzNTIwOTY1NzUwXz4NCkRBVEUgU0VOVA0KMDkvMTQvMjAyMA0KDQpORVcgUEFUSUVOVCBMRUFEIENPTlRBQ1RFRCBZT1UNCg0KQSBuZXcgcGF0aWVudCBsZWFkIGZyb20geW91ciB3ZWIgbWFya2V0aW5nIHNpdGUsIGhhcyBjb250YWN0ZWQgeW91IGZvcg0KaW5mb3JtYXRpb24gcmVnYXJkaW5nICpDb250YWN0IFVzKi4gUGxlYXNlIHNlZSB0aGUgbWVzc2FnZSB0aGV5IHNlbnQgeW91DQpiZWxvdy4gUGxlYXNlIGNhbGwgdGhlIHBhdGllbnQgb24gdGhlIG51bWJlciBwcm92aWRlZCwgdGV4dCB0aGUgcGF0aWVudA0KdGhyb3VnaCBHb29nbGUgVm9pY2Ugb3IgYW5vdGhlciBtb2RhbGl0eSBhbmQgc2VuZCB0aGUgcGF0aWVudCBhbiBlbWFpbA0KcmVzcG9uc2UuIFdlIGVuY291cmFnZSB5b3UgdG8gY29udGludWUgdHJ5aW5nIHRvIGdldCBpbiB0b3VjaCB3aXRoIHRoZQ0KcGF0aWVudCByZXBlYXRlZGx5IHVudGlsIHlvdSBoYXZlIG1ldCB3aXRoIHN1Y2Nlc3MuIFJlbWVtYmVyLCBZT1UgYXJlIHRoZQ0Ka2V5IHRvIHRoZSBzdWNjZXNzIG9mIHRoZSBtYXJrZXRpbmcgcHJvZ3JhbS4NCg0KSGVyZSBpcyB0aGUgcGFnZSB0aGUgcGF0aWVudCBsYW5kZWQgb24gYmVmb3JlIGNvbnRhY3RpbmcgeW91Og0KaHR0cHM6Ly9hbGFiYW1hLWRlbnRpc3QuY29tL2NvbnRhY3QtdXMvLg0KUEFUSUVOVCdTIElORk9STUFUSU9ODQpEYXRlOiAwOS8xNC8yMDIwDQpOYW1lOiBTaGVuYXZpYSBSZWVkZXINClBob25lOiArMTI1MTYyMjA0OTINCkVtYWlsOiBzbHI3MDFAamFnbWFpbC5zb3V0aGFsYWJhbWEuZWR1DQpNZXNzYWdlOg0KaG93IG11Y2ggZm9yIGEgYnJpZGdlIGZvciBhIG1pc3NpbmcgdG9vdGg_DQpQcm9jZWR1cmU6KkNvbnRhY3QgVXMqDQpodHRwczovL2FsYWJhbWEtZGVudGlzdC5jb20vY29udGFjdC11cy8NCipFTUFJTCBQQVRJRU5UKg0KPHNscjcwMUBqYWdtYWlsLnNvdXRoYWxhYmFtYS5lZHU_c3ViamVjdD1BK1Jlc3BvbnNlK2Zyb20rU3ByaW5naGlsbCtEZW50YWwrSGVhbHRoK0NlbnRlcj4NCipDQUxMDQpQQVRJRU5UKiA8KzEyNTE2MjIwNDkyPg0KR09PRCBMVUNLISBHTyBHRVQgJ0VNISBSZW1lbWJlcjogUGVyc2lzdGVuY2UgaXMgdGhlIGtleS4NCg0KVGhlIGludGVuZGVkIHJlY2lwaWVudCBvZiB0aGlzIGVtYWlsIGlzIFNwcmluZ2hpbGwgRGVudGFsIEhlYWx0aCBDZW50ZXIgYW5kDQppcyBhIHNlcnZpY2Ugb2YgV2ViIE1hcmtldGluZyBGb3IgRGVudGlzdHMgSWYgeW91IGhhdmUgcmVjZWl2ZWQgdGhpcyBlbWFpbA0KYnkgYWNjaWRlbnQscGxlYXNlIG5vdGlmeSB1cyBieSBjYWxsaW5nIDg3Ny02NjEtNTgyNS4NCg0KwqkgQ29weXJpZ2h0IDIwMjAgV2ViIE1hcmtldGluZyBGb3IgRGVudGlzdHMuQWxsIHJpZ2h0cyByZXNlcnZlZC4NCg=="

	// --- Decoding: convert encS from ShiftJIS to UTF8
	// declare a decoder which reads from the string we have just encoded
	rInUTF8 := transform.NewReader(strings.NewReader(encodedString), japanese.ShiftJIS.NewDecoder())
	// decode our string
	decBytes, _ := ioutil.ReadAll(rInUTF8)

	var decodedByte, _ = base64.StdEncoding.DecodeString(string(decBytes))
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

}