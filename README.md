# certhub

Track your ssl certificates in one click

Todo:

- CSV Import
- expiring soon based on user setting
- notify me before x days
- share domains (manage deletion, tracking activation)
- ability to check Now for one domain
- use a job queue for user events to check there certificates
  - https://github.com/rverton/pgjobs
  - redis
    - https://github.com/RHEnVision/provisioning-backend/blob/main/pkg/worker/redis.go
  - rabbimq
- Bounded Parallelism concurrency pattern
- transactions and migrations
  - create a user and send a worker payload for email validation

```dbml
Table ssl_certificates {
	id integer [primary key]
	domain varchar
	created_at timestamp
	status boolean
	issuer varchar
	expiry_date timestamp
	error varchar
}

Table domains {
	id int [primary key]
	name varchar
	up boolean
	last_checked timestamp
	latency int
}

Table users {
	id integer [primary key]
	email varchar
	created_at timestamp
	active boolean
	last_login timestamp
}

Table cert_tracking{
	certificate integer [primary key]
	user integer [primary key]
	notification_threshold integer
	notification_email varchar
	active boolean
}

Ref: cert_tracking.certificate > ssl_certificates.id
Ref: cert_tracking.user > users.id
```
