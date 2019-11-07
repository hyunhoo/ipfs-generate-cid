Generate CID in IPFS
======================

# 1. Explanation
How to generate CID in ifps is:

<code>File -> Raw data -> Chuck -> Hashing -> Multihashing -> Merkle root -> CID</code>


Currently only v0 CID supported.

# 2. How to use
<pre><code>$ go build -o generate_cid
$ ./generate_cid {path}
</code></pre>
