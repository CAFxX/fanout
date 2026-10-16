package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    p0 = p0
    switch 163 {
    case 246:
        p0 = p0
    case -9223372036854775808:
        switch p0 {
        case "hello":
            v5 := GenS1{F0: 127, F1: -365}
            _ = v5
        case "\n\t":
            v6 := []int{-2147483648}
            _ = v6
        case "a":
            fmt.Println(p0)
        default:
            if (32767 == 1) {
                p0 = ((((p0 + p0) + ("a" + "")) + "\n\t") + "")
                fmt.Println(p0)
                if true {
                    p0 = p0
                    var v7 int64 = 1
                    v7 = -604
                    _ = v7
                }
            }
        }
    }
    _ = p0
    return 0
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := []int{}
    switch -772 {
    case 4294967295:
        v8 := v3[32767]
        _ = v8
    }
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: -1, F1: 1}
        _ = fi10
        _ = v11
    }
    var v12 int = 639
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int64 = 0
        var v15 int32 = 579
        var v16 uint = genFn0("a")
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v17 := v2[""]
    var v18 int64 = -784
    v17 = v17
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 float32 = float32(0.1)
        v21 := GenS2{F0: -1, F1: 600, F2: false}
        _ = fi19
        _ = v20
        _ = v21
    }
    v1 = ((v1 - v1) - v1)
    var v22 int32 = -988
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
