# Dev from scratch

Bajo cada tecnología se enlistan:

- Requisitos recomendados: se aconseja cumplir con estos previo a meterse de lleno en la misma.
- Contenido standard: contenido base de la tecnología, sin especialización en ninguna de sus áreas.

La idea es ir tomando de a uno (en orden) los cursos o workshops listados, y usar a modo de "descanso" los bookmarks, que en su mayoría son charlas.

- Apéndices: contenido recomendable a la hora de querer especializarse en alguna de las vertientes que tenga la tecnología. Son tomables de forma paralela al contenido standard.

- Prueba básica: prueba arbitraria para asegurarse de que hay conocimiento mínimo como para contar con experiencia en la tecnología.

- [Dev from scratch](#dev-from-scratch)
  - [Lenguajes](#lenguajes)
    - [Go](#go)
      - [Requisitos recomendados](#requisitos-recomendados)
      - [Contenido standard](#contenido-standard)
      - [Prueba básica](#prueba-básica)
      - [Apéndices](#apéndices)
  - [IDEs & editors](#ides--editors)
    - [VSCode](#vscode)
    - [Vim](#vim)
  - [Conceptos](#conceptos)
  - [Infra & tooling](#infra--tooling)
    - [Linux](#linux)
    - [Networking](#networking)
    - [Git](#git)
    - [Docker](#docker)

## Lenguajes

### Go

#### Requisitos recomendados

- Un IDE (recomendado: VSCode + Go extension).
- Conceptos: 0, 2, 3, 7

#### Contenido standard

Referencia:

- [Effective Go - The Go Programming Language](https://golang.org/doc/effective_go.html)
- [Go doc](https://pkg.go.dev/)

Tours, cursos:

0. [ ] [Golang tutorial - TechWithTim](https://www.youtube.com/watch?v=75lJDVT1h0s&list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q)
1. [ ] [Go tour](https://tour.golang.org/welcome/1).
2. [ ] [Workshop - Practical Go](https://www.youtube.com/watch?v=gi7t6Pl9rxE)
3. [ ] [JustForFunc - Campoy](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)

Bookmarks:

0.  [ ] [Twelve Go Best Practices](https://talks.golang.org/2013/bestpractices.slide#36)
1.  [ ] [Clear is better than clever](https://www.youtube.com/watch?v=NwEuRO_w8HE)
2.  [ ] [Concurrency bugs](https://blog.acolyer.org/2019/05/17/understanding-real-world-concurrency-bugs-in-go/)
3.  [ ] [Dont Just Check Errors Handle Them Gracefully](https://www.youtube.com/watch?v=lsBF58Q-DnY)
4.  [ ] [Do you make these Go coding mistakes? · YourBasic Go](https://yourbasic.org/golang/gotcha/)
5.  [ ] [Go tooling essentials · Go, the unwritten parts](https://rakyll.org/go-tool-flags/)
6.  [ ] [Absolute Unit Test](https://www.youtube.com/watch?v=UKe5sX1dZ0k&list=PL402FFseOKOxtB3yb8Ww5tkLaKx3cTiWh&index=6)
7.  [ ] [The Zen of Go | Dave Cheney](https://dave.cheney.net/2020/02/23/the-zen-of-go)
8.  [ ] [GopherCon 2017: Mitchell Hashimoto - Advanced Testing with Go - YouTube](https://www.youtube.com/watch?v=8hQG7QlcLBk)
9.  [ ] [Solid Go](https://www.youtube.com/watch?v=zzAdEt3xZ1M)
10. [ ] [Golden Files — Why you should use them - Ibrahim Jarif - Medium](https://medium.com/@jarifibrahim/golden-files-why-you-should-use-them-47087ec994bf)
11. [ ] [cristaloleg/go-advices: List of advices and tricks for Go ʕ◔ϖ◔ʔ](https://github.com/cristaloleg/go-advices)
12. [ ] [First class functions](https://www.youtube.com/watch?v=5buaPyJ0XeQ)
13. [ ] [Guía de estilo de Uber](https://github.com/uber-go/guide/blob/master/style.md)
14. [ ] [GopherCon 2019: Mat Ryer - How I Write HTTP Web Services after Eight Years - YouTube](https://www.youtube.com/watch?v=rWBSMsLG8po)
15. [ ] [GopherCon 2019: Jessica Lucci - You Can't Go Your Own Way - YouTube](https://www.youtube.com/watch?v=hIWQeTL4gpQ)
16. [ ] [justforfunc #9: The Context Package - YouTube](https://www.youtube.com/watch?v=LSzR0VEraWw)
17. [ ] [Golang UK Conference 2017 | Jack Lindamood - How to correctly use package context - YouTube](https://www.youtube.com/watch?v=-_B5uQ4UGi0)
18. [ ] [How to correctly use context.Context in Go 1.7 - Jack Lindamood - Medium](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)
19. [ ] [Pitfalls of context values and how to avoid or mitigate them in Go - Calhoun.io](https://www.calhoun.io/pitfalls-of-context-values-and-how-to-avoid-or-mitigate-them/)
20. [ ] [How to Go Wrong with Concurrency - Speaker Deck](https://speakerdeck.com/aleksi/how-to-go-wrong-with-concurrency)
21. [ ] [Inserting records into a PostgreSQL database with Go's database/sql package - Calhoun.io](https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/)
22. [ ] [​​Our Go is fine but our SQL is great - Bumpers - Medium](https://medium.com/bumpers/our-go-is-fine-but-our-sql-is-great-b4857950a243)
23. [ ] [GopherCon 2019: Carolyn Van Slyck - Design Command-Line Tools People Love](https://www.youtube.com/watch?v=eMz0vni6PAw)
24. [ ] Conceptos: 1, 4.

#### Prueba básica

- [ ] [Cache](https://github.com/sebach1/dev-from-scratch/blob/master/test/go-01.md) ~ introduction
- [ ] [Backend](https://github.com/sebach1/dev-from-scratch/blob/master/test/go-02.md) ~ intermediate
- [ ] [Testing](https://github.com/sebach1/dev-from-scratch/blob/master/test/go-03.md) ~ advanced

#### Apéndices

Ver [apéndices](https://github.com/sebach1/dev-from-scratch/blob/master/bookmarks-go.md).

## IDEs & editors

### VSCode

- [ ] [Basics & getting started](https://code.visualstudio.com/docs/introvideos/basics)
- [ ] [VSCode Go](https://www.youtube.com/watch?v=Fnnp6FBYSvo)

### Vim

- [ ] [Vim adventures](https://vim-adventures.com/)
- [ ] [Headache](ihatemyself.com)

## Conceptos

0. [ ] [Rob Pike - 'Concurrency Is Not Parallelism'](https://www.youtube.com/watch?v=oV9rvDllKEg)
1. [ ] [Let’s talk about logging | Dave Cheney](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
2. [ ] [Go Proverbs - Rob Pike - Gopherfest](https://www.youtube.com/watch?v=PAAkCSZUG1c)
3. [ ] [APIs: The Good, The Bad, The Ugly](https://www.youtube.com/watch?v=t0nl5aUfAvY)
4. [ ] [Rob Pike - Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM)
5. [ ] [GopherCon 2019 Lightning Talk: Jessica Xie - How Not to DDOS Servers](https://www.youtube.com/watch?v=dT223YYT940)
6. [ ] [Micro Frontends - extending the microservice idea to frontend development](https://micro-frontends.org)
7. [ ] [JSON isi](https://www.youtube.com/watch?v=iiADhChRriM)


## Infra & tooling
todo
### Linux
todo

### Networking
todo

### Git
todo

### Docker
todo

