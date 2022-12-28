# short-ulid
Short Form "ULIDs"

**Experemental**

The ULID is in quotes intentionally, as this uses that spec as inspiration but has no intention of implementing the rigour required for a truly UUID. The goals are much more limited:

* create a lexicographically sortable identifier
* short form to be more user friendly

This package does not make any guarantees with:

* monotonicity of ids generated within the same millisecond
* non-collision

So what is it good for? It generates a short (10 character) id that is easier for humans to grok, and for use cases that do not required thousands of transactions ever millisecond it provides sufficient uniqueness.

## Components

### Timestamp

Unix millisecond timestamp represented as a 64bit int.

### Randomness (Entropy)

16 bits of CSPRNG data