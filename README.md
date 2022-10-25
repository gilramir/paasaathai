# Introduction

**paasaathai** is a Go module with linguistic functions for
parsing and analyzying written Thai language.
It is a work in progress and the API and functionality is constantly changing.

# Functionality

The module provides two major pieces of functionality.

## GraphmeStacks

This module can parse UTF-8 Thai text and produce objects called GraphemeStacks.
A single user-perceived Thai "character" can have multiple individual
graphemes in it. It might have "diacritic" vowels either above or below
the consonant. It might have a tone mark. And there a few other things too.

In this example, the word "bperd" ("to close") is shown as a single word:

![bperd as a word text](docs/example-bperd-word.jpg)

And here are the 3 individual GraphemeStacks.

![bperd as grapheme stacks](docs/example-bperd-grapheme-stacks.jpg)


## GStackClusters

There are Thai orthography rules that treat some sequences of Thai graphemes as
complete, unbreakable units. The article,
[Character Cluster Based Thai Information Retrieval](https://www.researchgate.net/profile/Virach-Sornlertlamvanich/publication/2853284_Character_Cluster_Based_Thai_Information_Retrieval/links/02e7e514db194bcb1f000000/Character-Cluster-Based-Thai-Information-Retrieval.pdf)
by Thanaruk Theeramunkong, Virach Sornlertlamvanich,
Thanasan Tanhermhong, and Wirat Chinnan, call these "Thai Character Clusters".

Most of these clusters are created from how some vowels are written as
multi-character sequences. In this example, you see the same word, "bperd",
broken up into 2 clusters of grapheme stacks, or GStackClusters.

![bperd as gstack clusters](docs/example-bperd-clusters.jpg)

The front character, a vowel, must be followed by a consonant.  They,
they form a cluster.  Note that these clusters are not syllables;
they are not units of sound. They are simply an artifact of the rules
of orthography. In this example, the final character stands alone
as a cluster by itself.
