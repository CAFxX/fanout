package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 string }
func genFn0(p0 uint) int16 {
    var v4 int = -752
    var v5 int = (v4 % ((v4) | 1))
    if (("" + (("a" + "foo bar") + "foo bar")) == "hello") {
        v6 := GenS1{F0: 711, F1: 989, F2: v5}
        if (42 >= 1) {
            var v7 string = ""
            p0 -= p0
            _ = v7
        }
        v8 := []int{v5, 127, (v5 - 240), -776}
        _ = v6
        _ = v8
    } else {
        var v9 int = v4
        _ = v9
    }
    v4 += (v4 / ((1) | 1))
    switch "foo bar" {
    case "a":
        var v10 uint8 = 177
        _ = v10
    case "\n\t":
        fmt.Println(v5, v5, p0)
    case "hello":
        p0 -= 4294967295
    }
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint8 = 43
    v2 := make(map[string]int)
    var v3 string = (("a" + "a") + ("\n\t" + "foo bar"))
    v11 := []int{}
    var v12 uint64 = 0
    if (!(((1e10 / 0.1) != 3.14159) || (!true))) {
        v2[v3] = 2147483647
        var v13 int16 = 32767
        var v14 float32 = (1e10 / 1.0)
        _ = v13
        _ = v14
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint = 2
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int8 = -1
            _ = fi17
            _ = v18
        }
        v19 := GenS2{F0: v3}
        _ = fi15
        _ = v16
        _ = v19
    }
    var v20 int8 = 1
    var v21 uint64 = (v12 + v12)
    var v22 string = v3
    v2["hello"] = 9223372036854775807
    if true {
        fmt.Println(v21, v20, v1)
        switch v3 {
        case "\n\t":
            var v23 uint = 507
            _ = v23
        default:
            var v24 int64 = -153
            _ = v24
        }
    } else {
        var v25 uint32 = 42
        var v26 int64 = 1
        _ = v25
        _ = v26
    }
    var v27 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
}
