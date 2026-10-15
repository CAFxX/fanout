package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32 }
func genFn0() string {
    var v5 uint = 226
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: 42, F1: 42}
    v8 := GenS1{F0: v5, F1: 758}
    var v9 int64 = v6
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 0
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1.5
    var v10 uint8 = 42
    var v11 float32 = float32(1e10)
    _ = len(v3)
    var v12 int8 = (v2 % ((-1) | 1))
    _ = len(v3)
    fmt.Println(v1, v4)
    v1 = v1
    v13 := GenS2{F0: 306}
    switch 349 {
    case -288:
        switch ("\n\t" + ("\n\t" + "a")) {
        case "":
            for fi14 := 0; fi14 < 8; fi14++ {
                v4 -= v4
                fmt.Println(v2)
                v15 := []int{4294967295}
                _ = fi14
                _ = v15
            }
        case "foo bar":
            v12 += 1
        }
    default:
        v12 += (v12 & v2)
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        _ = len(v3)
        var v17 uint16 = 32767
        _ = fi16
        _ = v17
    }
    var v18 int32 = int32(v10)
    v19 := GenS2{F0: v18}
    fmt.Println(v11)
    v20 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
