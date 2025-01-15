# Lookup MX Tool

Just a tool I wrote for a more specific task since MX Toolbox doesn't allow bulk by default and their API cost money to use so I wrote my own after learning the net library in Go has LookupMX. This isn't my best work, but I'm busy with a bunch of other projects and will try to write a shell script to parse the JSON for what I want with JQ.

It by default outputs to JSON so I can parse the output a bit easier.

You can compile this program by running:

```
git clone https://github.com/dleto614/lookup-mx
cd lookup-mx && go build
```

-----

#### Usage:

```
Usage of ./lookupmx: 

  -d string
        Specify domain.
  -i string
        Specify input file with a list of domains.
  -o string
        Specify the output file to write results into.
```
----

#### Domain:

```
./lookupmx -d google.com
./lookupmx -d google.com -o test-mx.json 
```

----

#### Input file:

```
./lookupmx -i test-domains.txt
./lookupmx -i test-domains.txt -o test-domains-mx.json
```
