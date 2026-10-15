package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 65535
    var v3 int = 127
    v4 := []int{4294967295}
    switch ("" + "foo bar") {
    case "hello":
        fmt.Println(v2, v3)
    default:
        v5 := GenS1{F0: 127, F1: 0, F2: 0}
        _ = v5
    }
    var v6 bool = (!((("\n\t" + "foo bar") != ("hello" + "")) && false))
    v7 := GenS2{F0: "a", F1: 42}
    var v8 uint32 = 17
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 float64 = -2.25
        var v11 int64 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = 0
    v3 += (^(v3))
    var v13 bool = (true || false)
    if (253 < 255) {
        if v13 {
            v14 := v1[(^(v8))]
            _ = v14
        } else {
            fmt.Println(v13, v2, v6)
            var v15 float64 = (1e10 + 1e10)
            _ = v15
        }
        if v13 {
            for fi16 := 0; fi16 < 14; fi16++ {
                v4 = append(v4, fi16)
                _ = fi16
            }
        }
        fmt.Println(v12, v2, v8)
    } else {
        var v17 int16 = -121
        var v18 int16 = (^(v17))
        v19 := GenS2{F0: ("\n\t" + ""), F1: -381}
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int32 = -569
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
