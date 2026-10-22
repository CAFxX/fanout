package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[int]bool)
    var v3 int32 = 2147483647
    var v4 int32 = -323
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(v4, v4)
        v3 += v4
        v3 -= v4
        _ = fi5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            v4 -= v3
            v7 := GenS1{F0: 856}
            switch "hello" {
            case "hello":
                var v8 float32 = 1e10
                _ = v8
            default:
                _ = len(v2)
            }
            _ = v7
        } else {
            switch ("hello" + "foo bar") {
            case "\n\t":
                var v9 uint = 572
                _ = v9
            default:
                v4 = v4
            }
        }
        var v10 bool = (2 == 2)
        _ = fi6
        _ = v10
    }
    v11 := GenS1{F0: 1}
    fmt.Println(v3, v4)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 string = "foo bar"
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 42
    v15 := v2[4294967295]
    v16 := GenS2{F0: 42, F1: 13, F2: uint16(v14)}
    var v17 int64 = 9223372036854775807
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS3{F0: v3, F1: v4, F2: 127}
        _ = fi18
        _ = v19
    }
    v2[1] = (!true)
    v4 += -130
    v15 = (139 != 32767)
    switch "\n\t" {
    case "":
        var v20 bool = v15
        _ = v20
    default:
        v3 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
