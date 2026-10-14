package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int) string {
    v6 := []int{p0}
    if true {
        var v7 uint64 = 42
        var v8 bool = (v7 > v7)
        _ = v7
        _ = v8
    }
    p0 = 42
    _ = p0
    _ = v6
    return "a"
}
func genFn1() bool {
    var v9 string = "hello"
    fmt.Println(v9)
    _ = v9
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 217
    var v3 bool = false
    var v4 uint16 = 0
    var v5 float64 = float64(0.0)
    if v3 {
        if v3 {
            switch ("" + ("a" + ("foo bar" + ("\n\t" + "hello")))) {
            case "hello":
                _ = len(v1)
            case "":
                if v3 {
                    v10 := GenS1{F0: 32767}
                    _ = v10
                } else {
                    switch 1 {
                    case 0:
                        for fi11 := 0; fi11 < 1; fi11++ {
                            var v12 bool = false
                            for fi13 := 0; fi13 < 2; fi13++ {
                                switch (^(fi11)) {
                                case 4294967295:
                                    delete(v1, "foo bar")
                                }
                                if v3 {
                                    var v14 int8 = 118
                                    var v15 int8 = v14
                                    v16 := []int{(fi11 - (fi13 ^ -885)), fi11, fi11, fi13}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                v4 -= 255
                                _ = fi13
                            }
                            _ = fi11
                            _ = v12
                        }
                    case -382:
                        v17 := []int{1}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        switch 189 {
                        case 9223372036854775807:
                            var v19 string = "a"
                            _ = v19
                        case 2147483647:
                            v3 = (v3 || (false && (1.5 <= 1.5)))
                        default:
                            fmt.Println(fi18)
                        }
                        _ = fi18
                    }
                    v20 := GenS1{F0: (v4 | (v4 / ((v4) | 1)))}
                    _ = v20
                }
            default:
                v21 := []int{}
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int64 = -219
            v24 := []int{}
            _ = fi22
            _ = v23
            _ = v24
        }
        fmt.Println(v3, v3)
    }
    var v25 int8 = 78
    if v3 {
        v3 = true
        v26 := GenS1{F0: (^(v4))}
        var v27 bool = true
        _ = v26
        _ = v27
    }
    switch "" {
    case "\n\t":
        var v28 int = 127
        _ = v28
    case "a":
        var v29 uint32 = 65535
        _ = v29
    case "foo bar":
        var v30 float64 = (v5 / v5)
        _ = v30
    }
    var v31 int64 = 0
    v32 := v1["a"]
    v33 := GenS1{F0: (^(v4))}
    var v34 int = (v32 | 502)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
