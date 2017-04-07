package main

//import "time"

type Resource struct {
	Resource	string	`json:"resource"`
	ContentType		string	`json:"content_type"`
	BrowserCached	bool	`json:"cached"`
	EncodedSize	int	`json:"encoded_size"`
	DecodedSize	int	`json:"decoded_size"`
	TransferSize	int	`json:"transfer_size"`
	HeaderSize	int	`json:"header_size"`
	RedirectTime	float32	`json:"redirect_time"`
	DnsLookupTime	float32	`json:"dns_lookup_time"`
	ConnectTime	float32	`json:"connect_time"`
	WaitTime	float32	`json:"wait_time"`
	SslTime		float32	`json:"ssl_time"`
	FirstByteTime	float32	`json:"first_byte_time"`
	ContentRespTime	float32	`json:"content_response_time"`
	WireTime	float32	`json:"wire_time"`
	RequestTime	float32	`json:"request_time"`
	ContentFetchTime	float32	`json:"fetch_time"`
	RequestRtt	float32	`json:"request_rtt"`
	FullRtt		float32	`json:"full_rtt"`
}

type Resources []Resource
