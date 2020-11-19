- [BENCHMARKING](#benchmarking)
  - [libs](#libs)
  - [talks](#talks)
- [GRAPHQL](#graphql)
  - [libs](#libs-1)
  - [talks](#talks-1)
- [CONCURRENCY](#concurrency)
  - [libs](#libs-2)
  - [talks](#talks-2)
- [DATA](#data)
  - [libs](#libs-3)
  - [talks](#talks-3)
- [DEPLOYMENT](#deployment)
  - [libs](#libs-4)
  - [talks](#talks-4)
- [TESTING](#testing)
  - [talks](#talks-5)
  - [libs](#libs-5)
- [OTHER](#other)
  - [libs](#libs-6)
  - [talks](#talks-6)
- [SECURITY](#security)
  - [talks](#talks-7)
  - [libs](#libs-7)


### BENCHMARKING

#### libs

[pprof - The Go Programming Language](https://golang.org/pkg/net/http/pprof/)

[aclements/perflock: Locking wrapper for running benchmarks on shared hosts](https://github.com/aclements/perflock)

[benchstat - GoDoc](https://godoc.org/golang.org/x/perf/cmd/benchstat)

[uber-archive/go-torch: Stochastic flame graph profiler for Go ... - GitHubhttps://github.com/uber-archive/go-torch](https://github.com/uber-archive/go-torch)

[julienschmidt/go-http-routing-benchmark: Go HTTP request router and web framework benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)

#### talks

[High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#profiling_benchmarks)

[How to optimize Go for really high performance - by Björn Rabenstein - YouTube](https://www.youtube.com/watch?v=ZuQcbqYK0BY)

[Golang UK Conference 2017 | Filippo Valsorda - Fighting latency: the CPU profiler is not your ally - YouTube](https://www.youtube.com/watch?v=Lxt8Vqn4JiQ)

[Profiling & Optimizing in Go / Brad Fitzpatrick - YouTube](https://www.youtube.com/watch?v=xxDZuPEgbBU)

[justforfunc #28: Basic Benchmarks - YouTube](https://www.youtube.com/watch?v=2AulMm-hsdI)

[davecheney/high-performance-go-workshop](https://github.com/davecheney/high-performance-go-workshop)

[GopherCon 2019: Chris Hines - Death by 3,000 Timers: Streaming Video-on-Demand for Cable TV - YouTube](https://www.youtube.com/watch?v=h0s8CWpIKdg)

### GRAPHQL

#### libs

[gqlgen](https://gqlgen.com/)

[graphql-go/handler: Golang HTTP.Handler for graphl-go](https://github.com/graphql-go/handler)

[machinebox/graphql: Simple low-level GraphQL HTTP client for Go](https://github.com/machinebox/graphql)
[go-chi/chi: lightweight, idiomatic and composable router for building Go HTTP services](https://github.com/go-chi/chi)

[samsarahq/thunder: ⚡️ A Go framework for rapidly building powerful graphql services](https://github.com/samsarahq/thunder)

#### talks

[Validating and Transforming Mutation Input | Graphcool Docs](https://www.graph.cool/docs/tutorials/functions/hook-functions-zeez7aiph3)

[Get Started With Persisted Queries in GraphQL](https://docs.scaphold.io/tutorials/persisted-queries/)

[An Absinthe GraphQL codec for better commits of JSON schema files https://medium.com/.../an-absinthe-graphql-codec-for-better-commits-of-json-schema...](https://medium.com/@hadfieldn/an-absinthe-graphql-codec-for-better-commits-of-json-schema-files-ea061b38676)

[GraphQL with Golang: A Deep Dive From Basics To Advanced](https://www.freecodecamp.org/news/deep-dive-into-graphql-with-golang-d3e02a429ac3/)

[GraphQL Schema Design: Building Evolvable Schemas - Apollo GraphQL](https://blog.apollographql.com/graphql-schema-design-building-evolvable-schemas-1501f3c59ed5)

### CONCURRENCY

#### libs

[Go - sync.Pool | go Tutorial](https://riptutorial.com/go/example/16314/sync-pool) \
[sync](https://golang.org/pkg/sync/)

#### talks

[How to Go Wrong with Concurrency - Speaker Deck](https://speakerdeck.com/aleksi/how-to-go-wrong-with-concurrency)

[Concurrency bugs](https://blog.acolyer.org/2019/05/17/understanding-real-world-concurrency-bugs-in-go/)

[GopherCon UK 2018: Roberto Clapis - Goroutines: The Dark Side of the Runtime - YouTube](https://www.youtube.com/watch?v=4CrL3Ygh7S0)

[Go: Understand the Design of Sync.Pool - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-understand-the-design-of-sync-pool-2dde3024e277)

### DATA

#### libs

[jonbodner/proteus: A simple tool for generating an application's data access layer.](https://github.com/jonbodner/proteus)

[soundcloud/roshi: Roshi is a large-scale CRDT set implementation for timestamped events.](https://github.com/soundcloud/roshi)

[Go Driver](https://docs.mongodb.com/ecosystem/drivers/go/)

[jmoiron/sqlx: general purpose extensions to golang's database/sql](https://github.com/jmoiron/sqlx)

[Davmuz/gqt: Go(lang) SQL Templates](https://github.com/Davmuz/gqt)

[gobuffalo/pop: A Tasty Treat For All Your Database Needs](https://github.com/gobuffalo/pop)

#### talks

[​​Our Go is fine but our SQL is great - Bumpers - Medium](https://medium.com/bumpers/our-go-is-fine-but-our-sql-is-great-b4857950a243)

[Inserting records into a PostgreSQL database with Go's database/sql package - Calhoun.io](https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/)

### DEPLOYMENT

#### libs

[Porter](https://porter.sh/) \
[facebookarchive/grace: Graceful restart & zero downtime deploy for Go servers.](https://github.com/facebookarchive/grace)

#### talks

[GopherCon UK 2018: Michael Hausenblas - Three Billy GOats Gruff: A Dev's Tale from VMs to Serverless - YouTube](https://www.youtube.com/watch?v=GUGYJup0xLc)

### TESTING

#### talks

[GopherCon UK 2018: Dmitry Matyukhin - Component & Integration Tests for Micro-services - YouTube](https://www.youtube.com/watch?v=Hm10y2tP5QQ)

[LondonGophers 20/03/2019: Dave Cheney - Absolute Unit (Test) - YouTube](https://www.youtube.com/watch?v=UKe5sX1dZ0k)

[Fuzzing Beyond Security: Automated Testing with go-fuzz by Filippo Valsorda at GothamGo 2015 - YouTube](https://www.youtube.com/watch?v=QEhPaj3vvPA)

[GopherCon 2017: Mitchell Hashimoto - Advanced Testing with Go - YouTube](https://www.youtube.com/watch?v=8hQG7QlcLBk)

[Go tooling essentials · Go, the unwritten partshttps://rakyll.org/go-tool-flags/](https://rakyll.org/go-tool-flags/)

[quii/learn-go-with-tests: Learn Go with test-driven development](https://github.com/quii/learn-go-with-tests)

[Golden Files — Why you should use them - Ibrahim Jarif - Medium](https://medium.com/@jarifibrahim/golden-files-why-you-should-use-them-47087ec994bf)

#### libs

[dvyukov/go-fuzz: Randomized testing for Go - GitHubhttps://github.com/dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz)

[matryer/is: Professional lightweight testing mini-framework for Go.](https://github.com/matryer/is)

[matryer/moq: Interface mocking tool for go generate](https://github.com/matryer/moq)

[google/go-cmp: Package for comparing Go values in tests](https://github.com/google/go-cmp)

[google/syzkaller: syzkaller is an unsupervised coverage-guided kernel fuzzer](https://github.com/google/syzkaller)

[spf13/afero: A FileSystem Abstraction System for Go](https://github.com/spf13/afero)

### OTHER

#### libs

[spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)

[go-kit/kit: A standard library for microservices.](https://github.com/go-kit/kit)

[gore](https://github.com/sriram-srinivasan/gore)

[uber-go/zap: Blazing fast, structured, leveled logging in Go.](https://github.com/uber-go/zap)

[pkg/rock: Create semantic version tags for your Go packages, search and discover new packages](https://github.com/pkg/rock)

[opentracing-contrib/go-zap: Integration with go.uber.org/zap](https://github.com/opentracing-contrib/go-zap)

[jaegertracing/jaeger-ui: Web UI for Jaeger](https://github.com/jaegertracing/jaeger-ui) \
[client9/misspell: Correct commonly misspelled English words in source files](https://github.com/client9/misspell) \
[gravityblast/fresh: Build and (re)start go web apps after saving/creating/deleting source files.](https://github.com/gravityblast/fresh) \
[maruel/panicparse: Crash your app in style (Golang)](https://github.com/maruel/panicparse) \
[dustin/go-humanize: Go Humans! (formatters for units to human friendly sizes)](https://github.com/dustin/go-humanize)

#### talks

[rustgo: calling Rust from Go with near-zero overhead](https://blog.filippo.io/rustgo/)

[GopherCon 2019: Gabbi Fisher - Socket to Me: Where do Sockets Live in Go? - YouTube](https://www.youtube.com/watch?v=pGR3r0UhoS8)

[A Complete Go Development Environment With Docker and VS Code](https://medium.com/better-programming/a-complete-go-development-environment-with-docker-and-vscode-a3e4410d27f7)

[A no-nonsense guide to environment variables in Go - DEV Community 👩‍💻👨‍💻](https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f)

[Go: Vet Command Is More Powerful Than You Think - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-vet-command-is-more-powerful-than-you-think-563e9fdec2f5)

[Process synchronization monitors in go - Angad Sharma - Medium](https://medium.com/@angadsharma1016/process-synchronization-monitors-in-go-d31f4c42fce7)

[How to correctly use context.Context in Go 1.7 - Jack Lindamood - Medium](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

[Pitfalls of context values and how to avoid or mitigate them in Go - Calhoun.io](https://www.calhoun.io/pitfalls-of-context-values-and-how-to-avoid-or-mitigate-them/)

[The anatomy of Slices in Go - Run Go - Medium](https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94)

[Golang UK Conference 2017 | Jack Lindamood - How to correctly use package context - YouTube](https://www.youtube.com/watch?v=-_B5uQ4UGi0)

[justforfunc #9: The Context Package - YouTube](https://www.youtube.com/watch?v=LSzR0VEraWw)

[cristaloleg/go-advices: List of advices and tricks for Go ʕ◔ϖ◔ʔ](https://github.com/cristaloleg/go-advices)

[Clearing the GOPATH to Go - Golab 2017 keynote on Vimeo](https://vimeo.com/200469720)

[Arrays vs Slices | Go Design Patterns](https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html)

[50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

[Do you make these Go coding mistakes? · YourBasic Go](https://yourbasic.org/golang/gotcha/)

[Take OpenTracing for a HotROD ride - OpenTracing - Medium](https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941)

[google/cel-go: Fast, portable, non-Turing complete expression evaluation with gradual typing (Go)](https://github.com/google/cel-go)

[google/cel-spec: Common Expression Language -- specification and binary representation](https://github.com/google/cel-spec)

### SECURITY

#### talks

[Go-SCP/src at master · OWASP/Go-SCP](https://github.com/OWASP/Go-SCP/tree/master/src)
[Clear is better than clever](https://www.youtube.com/watch?v=NwEuRO_w8HE)
[Best practices for securely storing API keys](https://www.freecodecamp.org/news/how-to-securely-store-api-keys-4ff3ea19ebda/)

[The Right Way To Manage Secrets](https://aws.amazon.com/blogs/mt/the-right-way-to-store-secrets-using-parameter-store/)

[Implementing OAuth 2.0 with Go(Golang) 🔐 - Soham's blog](https://www.sohamkamani.com/blog/golang/2018-06-24-oauth-with-golang/)

#### libs

[segmentio/chamber: CLI for managing secrets](https://github.com/segmentio/chamber) \
