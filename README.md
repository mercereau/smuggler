# smuggler
Simple HTTP tool for consuming JSON, specifically RUM data

The need for a super lightweight HTTP server that could consume and grok simple JSON data without the overhead of logstash reaquired a custom HTTP server in go.

Eventually, this will support any JSON thrown at it. For now, the json is unmarshaled only to verify proper json format.

No Changes
