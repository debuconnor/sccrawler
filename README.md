# spaceCloudCrawler

This is for Space Cloud Hosts.

Expect Workflow:
1. Crawling reservations of Space Cloud as HTML
2. Parsing reservations to store into DB
3. Sending SMS to the reserved users

Requirements & Test spec:
- GCP e2-micro
- Go 1.17
- Sqlite3 3.37
- Chrome
