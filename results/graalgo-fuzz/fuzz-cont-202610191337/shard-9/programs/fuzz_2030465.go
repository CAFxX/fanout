package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0() int {
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    v5 += v5
    _ = v5
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 126
    v3 := make(map[int]uint32)
    var v4 float64 = 3.14159
    v8 := GenS1{F0: -174, F1: 127}
    if (!true) {
        if (int32(v2) != 2147483647) {
            v4 = ((1.5 + v4) + (3.14159 * (1e10 / (float64(0.1) / ((1e10 - 0.1) + -2.25)))))
            var v9 uint32 = 1
            _ = v9
        } else {
            v4 = (v4 + (3.14159 / 1.0))
        }
        var v10 int8 = 1
        _ = v10
    }
    v2 = (^(v2))
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := []int{127, genFn0(), fi11}
        fmt.Println(v2)
        _ = fi11
        _ = v12
    }
    fmt.Println(v4, v2)
    var v13 int8 = -37
    if (("\n\t" + "foo bar") >= "a") {
        switch 42 {
        case -1:
            var v14 int8 = 127
            _ = v14
        default:
            var v15 int8 = v13
            _ = v15
        }
        v4 += float64(-2.25)
    }
    if false {
        var v16 int16 = 1000
        if true {
            fmt.Println(v2)
        }
        fmt.Println(v4, v2, v4)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 2; fi17++ {
            v13 -= v13
            _ = fi17
        }
    }
    if true {
        v3[genFn0()] = 42
        var v18 int32 = -173
        if true {
            v4 -= ((v4 / -2.25) * -2.25)
        } else {
            v13 = 100
        }
        _ = v18
    } else {
        v2 = v2
        v13 = 1
        var v19 float64 = (v4 - v4)
        _ = v19
    }
    v2 = v2
    v1["\n\t"] = 1
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    switch ((("" + "") + "a") + "a") {
    case "\n\t":
        switch "\n\t" {
        case "":
            v4 += v4
        case "\n\t":
            v22 := v1["hello"]
            _ = v22
        }
    case "hello":
        var v23 uint64 = 18446744073709551615
        _ = v23
    case "":
        var v24 int64 = -1
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
