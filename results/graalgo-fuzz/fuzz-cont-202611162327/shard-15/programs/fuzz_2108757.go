package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int64 = 0
    var v2 int8 = -1
    var v3 int32 = 0
    v4 := GenS1{F0: ("\n\t" + ("" + "")), F1: v3}
    if true {
        if (!(!(2 <= 431))) {
            var v5 int32 = v3
            v6 := GenS1{F0: "foo bar", F1: 42}
            _ = v5
            _ = v6
        }
    } else {
        switch "a" {
        case "":
            v7 := GenS2{F0: 4294967295, F1: (v3 | (v3 / ((0) | 1))), F2: v1}
            _ = v7
        }
    }
    var v8 uint8 = 127
    v9 := []int{9223372036854775807}
    v1 = (-(v1))
    v2 += (v2 - 100)
    var v10 float32 = 1e10
    v9 = append(v9, 128)
    var v11 uint8 = (v8 << ((879) & 63))
    fmt.Println(v8)
    if (4294967295 > 4294967295) {
        var v12 uint = 2
        _ = v12
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v2)
            var v14 bool = false
            if (!(("hello" + "\n\t") == "a")) {
                var v15 uint = 4294967295
                v16 := GenS3{F0: v3, F1: 66}
                var v17 uint = v15
                _ = v15
                _ = v16
                _ = v17
            } else {
                v10 = (v10 - (1e10 / float32(1e10)))
            }
            _ = fi13
            _ = v14
        }
        switch "\n\t" {
        case "\n\t":
            switch "" {
            case "\n\t":
                fmt.Println(v10, v3)
            }
        default:
            var v18 uint64 = 18446744073709551615
            _ = v18
        }
    }
    v19 := GenS2{F0: 974, F1: v3, F2: v1}
    var v20 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
