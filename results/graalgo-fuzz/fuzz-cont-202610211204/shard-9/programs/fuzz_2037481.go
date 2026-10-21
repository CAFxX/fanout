package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int16) bool {
    var v4 int8 = 0
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p0, v4)
        _ = fi5
    }
    _ = p0
    _ = v4
    return false
}

func main() {
    var v1 bool = false
    v2 := []string{"a", "hello", "hello", "\n\t"}
    var v3 string = "foo bar"
    v3 = "hello"
    fmt.Println(v1, v1)
    switch 32768 {
    case -2147483648:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int8 = 1
            var v8 string = "foo bar"
            var v9 bool = false
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    case 0:
        var v10 int64 = -1
        _ = v10
    default:
        for fi11 := 0; fi11 < 19; fi11++ {
            v2 = append(v2, "hello")
            var v12 int32 = 2147483647
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v1 = v1
        _ = fi13
    }
    var v14 string = "foo bar"
    fmt.Println(v3)
    if (v14 == v14) {
        var v15 bool = v1
        v3 = v3
        fmt.Println(v14, v1)
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    v1 = v1
    v16 := GenS1{F0: (!(-1 < -1))}
    var v17 uint32 = 2147483648
    var v18 uint64 = 839
    switch v14 {
    case "hello":
        var v19 float32 = 1.5
        _ = v19
    case "\n\t":
        if (((1.5 - float64(0.0)) - -2.25) <= 1e10) {
            v1 = (-32768 == 1)
            v18 = v18
            for fi20 := 0; fi20 < 11; fi20++ {
                v21 := GenS2{F0: v17, F1: 42}
                var v22 uint32 = (v17 / (((v17 & v17)) | 1))
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            v1 = v1
            v18 = (v18 | (v18 | 458))
            var v23 float32 = 0.0
            _ = v23
        }
    default:
        var v24 bool = v1
        _ = v24
    }
    v25 := GenS2{F0: (v17 * (^(v17))), F1: -969}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
