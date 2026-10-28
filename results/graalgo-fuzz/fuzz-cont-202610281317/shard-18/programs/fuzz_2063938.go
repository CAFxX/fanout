package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 int; F2 string }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    v6 := GenS1{F0: 0, F1: -1, F2: 780}
    var v7 bool = false
    var v8 uint64 = 362
    v9 := GenS2{F0: "hello", F1: 127, F2: "foo bar"}
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{}
        v7 = (42 <= 1)
        _ = fi10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 34
}
func genFn1(p0 int) string {
    v12 := []int{}
    switch "" {
    case "hello":
        v13 := GenS2{F0: "", F1: p0, F2: ""}
        _ = v13
    case "a":
        var v14 uint64 = 42
        _ = v14
    case "":
        if ((true || (1000 >= 42)) && false) {
            v15 := GenS1{F0: 65535, F1: -128, F2: 1}
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint = 42
                var v18 int16 = 32767
                switch "\n\t" {
                case "a":
                    var v19 uint8 = 1
                    _ = v19
                case "hello":
                    fmt.Println(p0, v18)
                case "\n\t":
                    _ = len(v12)
                default:
                    fmt.Println(v17, v17)
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            v20 := GenS2{F0: (("\n\t" + "\n\t") + "hello"), F1: 966, F2: "a"}
            p0 = p0
            _ = v20
        }
    }
    var v21 bool = (true || (!false))
    _ = p0
    _ = v12
    _ = v21
    return "foo bar"
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{("a" + "foo bar"), "", ("foo bar" + ""), "foo bar"}
    var v3 uint8 = 2
    v4 := []int{}
    v5 := []int{961, 32768}
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            v3 = (v3 | (v3 | 67))
            _ = fi22
        }
        v1 = (v1 + 1e10)
    }
    var v23 float64 = 0.1
    _ = len(v2)
    for fi24 := 0; fi24 < 8; fi24++ {
        if (-55 >= 100) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v25 := GenS2{F0: ("a" + ("\n\t" + "")), F1: (^(fi24)), F2: (("a" + "foo bar") + "")}
        _ = fi24
        _ = v25
    }
    if false {
        var v26 uint = 42
        v23 += ((0.0 * (v23 / -2.25)) * (float64(0.1) + (0.1 + (v23 - 3.14159))))
        _ = v26
    } else {
        fmt.Println(v3, v3, v23)
    }
    switch ("a" + "a") {
    case "":
        var v27 uint = 18446744073709551615
        _ = v27
    case "a":
        var v28 bool = (209 == -876)
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        _ = len(v4)
        fmt.Println(v1)
        _ = fi29
    }
    v30 := GenS3{F0: 166}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
