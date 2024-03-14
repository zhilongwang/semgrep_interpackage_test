# Introduction.

This folder contains two go package, and the packagea/test.go invoke functions from packagea and packageb. 

The rule.yaml is written to detected data flow from input to output. 

To test the rule on the code, I run the command:

```bash
semgrep --pro --config=./rule.yaml .
```

However, it does not able to detect the dataflow pattern that across the two packages. 



```bash

 packagea/test.go
    ❯❱ test_inter_analysis
          message: test_inter analysis.

           25┆ fmt.Println("Output:", output)


          Taint comes from:

           16┆ input, _ := reader.ReadString('\n')


          Taint flows through these intermediate variables:

           16┆ input, _ := reader.ReadString('\n')

           19┆ output := processInput(input)

           29┆ func processInput(input string) string {

           19┆ output := processInput(input)


                This is how taint reaches the sink:

           25┆ fmt.Println("Output:", output)


            ⋮┆----------------------------------------
           31┆ fmt.Println("Output:", input)


          Taint comes from:

           16┆ input, _ := reader.ReadString('\n')


          Taint flows through these intermediate variables:

           16┆ input, _ := reader.ReadString('\n')


                This is how taint reaches the sink:

           19┆ output := processInput(input)

          Taint flows through these intermediate variables:

           29┆ func processInput(input string) string {

          then reaches:

           31┆ fmt.Println("Output:", input)




┌──────────────┐
│ Scan Summary │
└──────────────┘
Some files were skipped or only partially analyzed.
  Scan was limited to files tracked by git.

Ran 1 rule on 3 files: 2 findings.
```
