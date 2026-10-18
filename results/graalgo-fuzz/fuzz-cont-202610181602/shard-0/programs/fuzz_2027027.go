package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 int) uint {
    var v6 int = p0
    var v7 uint8 = 61
    switch "" {
    case "foo bar":
        fmt.Println(p0, v7)
    case "\n\t":
        v8 := GenS1{F0: "a", F1: 0}
        _ = v8
    }
    if (true || true) {
        v9 := []int{}
        v7 = v7
        var v10 uint16 = 1
        _ = v10
        _ = v9
    } else {
        if true {
            var v11 uint64 = 42
            _ = v11
        }
        if true {
            var v12 int64 = 1
            _ = v12
        }
        var v13 int32 = 42
        _ = v13
    }
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 119
    var v2 uint8 = 0
    v3 := []string{""}
    var v4 uint32 = 65535
    var v5 int32 = 1
    var v14 uint8 = 207
    v1 = v1
    v14 = (v14 % ((v14) | 1))
    if (1000 >= 1000) {
        v2 = v2
        _ = len(v3)
        v15 := GenS1{F0: ("" + ((("foo bar" + "hello") + "hello") + "\n\t")), F1: 105}
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v5 = v5
        _ = fi16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint64 = 369
            fmt.Println(v14, fi18)
            _ = fi18
            _ = v19
        }
        var v20 uint16 = 610
        fmt.Println(v2, v2)
        _ = fi17
        _ = v20
    }
    if true {
        v3 = append(v3, ("\n\t" + "foo bar"))
    }
    if true {
        var v21 int = int(v5)
        _ = v21
    }
    var v22 float64 = (-2.25 + (0.0 + ((0.0 / 0.1) - (float64(-2.25) + (-2.25 / (1.5 * 1e10))))))
    v2 += v14
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    _ = len(v3)
    v4 += (v4 | 151)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
