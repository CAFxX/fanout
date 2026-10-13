package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int16 = 839
    var v2 int8 = -1
    var v3 int = -1
    v4 := []int{(v3 * 42), (v3 | 1), (v3 * (v3 % ((v3) | 1))), (v3 * (v3 ^ 127))}
    v5 := []string{(("\n\t" + "foo bar") + ""), "hello", ("hello" + "a"), "foo bar"}
    v6 := make(map[string]bool)
    delete(v6, ("hello" + ("" + "\n\t")))
    var v7 int16 = -171
    var v8 int64 = 604
    for fi9 := 0; fi9 < 18; fi9++ {
        v7 = v7
        var v10 bool = (!true)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int = -1
        v1 -= 32767
        v6["a"] = ((^(v1)) < v1)
        _ = fi11
        _ = v12
    }
    v2 += v2
    v8 += v8
    if true {
        if false {
            v13 := GenS1{F0: 1}
            _ = v13
        }
        if len(v5) > 0 {
            v5[0] = ("" + ("a" + "foo bar"))
        }
        v8 += (v8 / ((v8) | 1))
    } else {
        v1 = v1
    }
    if ("a" != "") {
        if len(v4) > 0 {
            v4[0] = 32767
        }
        var v14 bool = (!false)
        _ = v14
    } else {
        var v15 float64 = 1e10
        _ = v15
    }
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
