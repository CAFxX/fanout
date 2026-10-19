package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0() int {
    v7 := GenS1{F0: 100, F1: "foo bar"}
    v8 := []int{-9223372036854775808}
    var v9 int = -9223372036854775808
    if false {
        if len(v8) > 0 {
            v8[0] = v9
        }
        v9 = (v9 - 9223372036854775807)
    }
    fmt.Println(v9)
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]int)
    var v3 uint32 = 2
    v4 := []string{("a" + "a"), "hello", ("\n\t" + ("hello" + ("" + "foo bar"))), "foo bar"}
    v5 := make(map[string]string)
    var v6 int32 = 2147483647
    if len(v4) > 0 {
        v4[0] = "a"
    }
    v10 := GenS2{F0: (v6 * -269), F1: 42, F2: (v3 + (^(v3)))}
    v11 := GenS1{F0: 100, F1: "\n\t"}
    switch "foo bar" {
    case "\n\t":
        v6 = 1
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = "hello"
        }
    }
    fmt.Println(v3)
    switch genFn0() {
    case 460:
        var v12 uint8 = 42
        _ = v12
    default:
        var v13 int8 = -44
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: v6, F1: 77, F2: 900}
            _ = fi15
            _ = v16
        }
        v6 += -1
        v17 := v5[("foo bar" + ((("foo bar" + ((("" + "hello") + ("" + ("hello" + "hello"))) + "hello")) + ("hello" + "\n\t")) + "\n\t"))]
        _ = fi14
        _ = v17
    }
    if true {
        var v18 int8 = -128
        var v19 uint64 = 170
        _ = v18
        _ = v19
    }
    var v20 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
