# 🔐 Security Policy

Security is a priority. We maintain a proactive stance to identify and fix vulnerabilities in **go-chaincfg**.

<br/>

## 🛠️ Supported & Maintained Versions

| Version | Status               |
|---------|----------------------|
| 1.x.x   | ✅ Supported & Active |

<br/>

## 📨 Reporting a Vulnerability

If you’ve found a security issue, **please don’t open a public issue or PR**.

Instead, send a private email to:
📧 [security@bsvassocation.org](mailto:security@bsvassociation.org)

Include the following:

* 🕵️ Description of the issue and its impact
* 🧪 Steps to reproduce or a working PoC
* 🔧 Any known workarounds or mitigations

We welcome responsible disclosures from researchers, vendors, users, and curious tinkerers alike.

<br/>

## 📅 What to Expect

* 🧾 **Acknowledgment** within 72 hours
* 📢 **Status updates** every 5 business days
* ✅ **Resolution target** of 30 days (for confirmed vulnerabilities)

Prefer encrypted comms? Let us know in your initial email—we’ll reply with our PGP public key. 
All official security responses are signed with it.

<br/>

## 🧪 Security Tooling

We regularly scan for known vulnerabilities using:

* [`govulncheck`](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck): Checks Go code and dependencies for known vulnerabilities using the Go vulnerability database.
* [`ask nancy`](https://github.com/sonatype-nexus-community/nancy): As part of our CI (see `fortress.yml`), we run [nancy](https://github.com/sonatype-nexus-community/nancy) to check Go dependencies for vulnerabilities against the OSS Index. This helps us catch issues in third-party packages early.

Want to run these yourself?

```sh
make govulncheck
# or run nancy via the CI workflow
```

This will check your local build for known issues in Go modules and dependencies.

<br/>

## 🛡️ Security Standards

We follow the [OpenSSF](https://openssf.org) best practices to ensure this repository remains compliant with industry‑standard open source security guidelines.

<br/>

## 🛠️ GitHub Security Workflows

To proactively protect this repository, we use several automated GitHub workflows:

- **[CodeQL Analysis](./workflows/codeql-analysis.yml)**: Scans the codebase for security vulnerabilities and coding errors using GitHub's CodeQL engine on every push and pull request to the `master` branch.
- **[OpenSSF Scorecard](./workflows/scorecard.yml)**: Periodically evaluates the repository against OpenSSF Scorecard checks, providing insights and recommendations for improving supply chain security and best practices.

These workflows help us identify, remediate, and prevent security issues as early as possible in the development lifecycle. For more details, see the workflow files in the [`.github/workflows/`](https://github.com/bsv-blockchain/go-chaincfg/tree/master/.github/workflows) directory.

<br/>
