package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    var v5 int16 = 1
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6, fi6, fi6, 4294967295}
        var v8 int32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := []string{"hello"}
    var v4 int32 = -1
    for fi9 := 0; fi9 < 13; fi9++ {
        if (!true) {
            v4 = v4
            v10 := GenS1{F0: 1, F1: -32768}
            if false {
                v4 += (v4 / ((v4) | 1))
            }
            _ = v10
        } else {
            var v11 float64 = ((1.5 - 0.0) / (-2.25 + 0.1))
            v12 := GenS2{F0: 42}
            switch "a" {
            case "\n\t":
                v4 = 1
            default:
                var v13 uint16 = 0
                _ = v13
            }
            _ = v11
            _ = v12
        }
        v14 := GenS1{F0: -2147483648, F1: -1}
        v15 := GenS1{F0: 0, F1: 464}
        _ = fi9
        _ = v14
        _ = v15
    }
    v4 = (v4 & 0)
    delete(v1, 1)
    v16 := GenS2{F0: 32767}
    fmt.Println(v4)
    var v17 int64 = -9223372036854775808
    var v18 int32 = (^(v4))
    v19 := GenS2{F0: 42}
    v4 = (v4 * (v4 * -372))
    switch "" {
    case "\n\t":
        var v20 int16 = -818
        _ = v20
    case "foo bar":
        v17 = 1
    case "":
        v18 = v4
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = 1
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
