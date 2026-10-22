package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 uint16, p1 int64, p2 uint64) string {
    if false {
        v7 := GenS1{F0: 903, F1: 781}
        for fi8 := 0; fi8 < 5; fi8++ {
            p2 += 4294967295
            var v9 int64 = p1
            _ = fi8
            _ = v9
        }
        p1 = p1
        _ = v7
    } else {
        var v10 float64 = -2.25
        _ = v10
    }
    p2 = 860
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 255
    var v3 int16 = -808
    v4 := make(map[int]int)
    var v5 uint = 1
    v6 := make(map[string]string)
    var v11 bool = false
    v12 := GenS1{F0: -2147483648, F1: -79}
    v13 := GenS1{F0: -1, F1: 9223372036854775807}
    fmt.Println(v5, v2, v3)
    var v14 int32 = -391
    v15 := GenS1{F0: 0, F1: -778}
    var v16 uint8 = 127
    switch "hello" {
    case "\n\t":
        fmt.Println(v14)
    default:
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 bool = v11
            _ = fi17
            _ = v18
        }
    }
    v3 += v3
    v19 := GenS1{F0: v14, F1: 549}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
