## pathtree

A pathtree is a data structure for storing filesystem and API paths for
effective retrieval. 

Pathtrees are typically based on ideas from
[Trie](ahttps://en.wikipedia.org/wiki/Trie),
[Radix Tree](httaps://en.wikipedia.org/wiki/Radix_tree), and  [Suffix
Tree](https://en.w....ikipedia.org/wiki/Suffix_tree)


For example, a few of GitHub's APIs can be organized in to the following
pathtree:

```
/
|-- /gists
|   |-- /:id
|   |   |-- /forks
|   |   `-- /star
|   |-- /public
|   `-- /starred
|-- /repositories
|-- /user
|   |-- /emails
|   |-- /followers
|   |-- /following
|   |   `-- /:user
|   |-- /issues
|   |-- /keys
|   |   `-- /:id
|   |-- /orgs
|   |-- /repos
|   |-- /starred
|   |   `-- /:owner
|   |       `-- /:repo
|   |-- /subscriptions
|   |   `-- /:owner
|   |       `-- /:repo
|   `-- /teams
`-- /users
    `-- /:user
        |-- /events
        |   |-- /orgs
        |   |   `-- /:org
        |   `-- /public
        |-- /followers
        |-- /following
        |   `-- /:target_user
        |-- /gists
        |-- /keys
        |-- /orgs
        |-- /received_events
        |   `-- public
        |-- /repos
        |-- /starred
        `--/subscriptions
```

## Experiments

The following experiments (or version) were created:

0. Simple pathtree that can store static paths.
