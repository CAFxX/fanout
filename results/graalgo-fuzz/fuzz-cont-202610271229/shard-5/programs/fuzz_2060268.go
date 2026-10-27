package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0() string {
    v4 := GenS1{F0: "foo bar", F1: 42, F2: 42}
    if true {
        var v5 bool = (-455 > 1000)
        v5 = (!v5)
        _ = v5
    }
    var v6 float64 = 1e10
    _ = v4
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint16, p1 uint) int32 {
    v7 := []int{}
    if true {
        var v8 int = 127
        v9 := GenS1{F0: "hello", F1: 255, F2: uint64(p1)}
        _ = v8
        _ = v9
    } else {
        p0 += (p0 | p0)
    }
    var v10 uint = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 42
}

func main() {
    v1 := []string{"", "", ("hello" + "hello")}
    var v2 bool = true
    var v3 int32 = 0
    var v11 float32 = 0.1
    _ = len(v1)
    switch 0 {
    case 128:
        var v12 float32 = v11
        _ = v12
    case 9223372036854775807:
        switch 9223372036854775807 {
        case -896:
            v3 -= (v3 ^ v3)
        default:
            var v13 float64 = 1.5
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v11 += 3.14159
        _ = fi14
    }
    var v15 int = -903
    _ = len(v1)
    fmt.Println(v11)
    fmt.Println(v11, v15, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
