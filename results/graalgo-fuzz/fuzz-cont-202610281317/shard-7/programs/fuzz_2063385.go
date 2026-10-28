package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 int64) int16 {
    if false {
        var v6 uint = 0
        p0 -= (p0 + p0)
        _ = v6
    }
    var v7 float32 = 3.14159
    p0 -= 42
    _ = p0
    _ = v7
    return 32767
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = v1
    var v3 bool = false
    var v4 uint64 = 4294967295
    v5 := make(map[string]int)
    delete(v5, "\n\t")
    var v8 int32 = -1
    if v3 {
        v5["\n\t"] = -1
    }
    v9 := GenS1{F0: -2147483648, F1: -182}
    v10 := []int{-1, 32767, -9223372036854775808}
    v11 := GenS2{F0: 127, F1: "\n\t"}
    if v3 {
        v1 = 1e10
        if true {
            v2 = -2.25
        }
    }
    switch "hello" {
    case "a":
        fmt.Println(v8, v4)
    case "hello":
        v2 -= v1
    case "\n\t":
        v12 := GenS2{F0: 84, F1: "a"}
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint32 = 65535
            _ = fi14
            _ = v15
        }
        var v16 string = (("" + "\n\t") + "foo bar")
        _ = fi13
        _ = v16
    }
    var v17 bool = v3
    switch "hello" {
    case "a":
        v18 := GenS2{F0: 128, F1: "hello"}
        _ = v18
    case "hello":
        if len(v10) > 0 {
            v10[0] = 32768
        }
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
