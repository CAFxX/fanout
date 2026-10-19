package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 uint32 = 785
    v4 := make(map[string]uint32)
    var v5 float64 = 0.1
    var v6 uint32 = v3
    v7 := []int{-2147483648, -2147483648, -451}
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float64 = v5
        v11 := GenS1{F0: 631, F1: 427, F2: ((("foo bar" + "a") + "foo bar") + "hello")}
        v6 += v3
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 = v6
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            for fi13 := 0; fi13 < 15; fi13++ {
                v5 = ((v5 + ((1e10 - (v5 + float64(1.5))) / v5)) / 3.14159)
                var v14 int16 = -705
                v15 := GenS2{F0: 2, F1: (^(v14))}
                var v16 uint64 = 18446744073709551615
                v2 -= v2
                v17 := v4["foo bar"]
                if (v16 == (v8 >> ((uint(v8)) & 63))) {
                    v18 := v1["foo bar"]
                    _ = v18
                }
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            var v19 uint64 = (v8 << ((65535) & 63))
            _ = v19
        }
        _ = fi12
    }
    var v20 uint8 = 192
    switch "foo bar" {
    case "\n\t":
        if (v8 >= v8) {
            var v21 float32 = 1.5
            var v22 int8 = -12
            _ = v21
            _ = v22
        }
    default:
        switch "" {
        case "a":
            var v23 bool = (42 == -1)
            _ = v23
        case "foo bar":
            delete(v1, ("" + ("\n\t" + "hello")))
        default:
            v24 := v1["foo bar"]
            _ = v24
        }
    }
    var v25 int = int(v2)
    if (293 != v8) {
        var v26 uint16 = uint16(v8)
        fmt.Println(v26)
        var v27 uint64 = v8
        _ = v26
        _ = v27
    } else {
        v28 := GenS1{F0: 779, F1: 276, F2: ""}
        var v29 uint64 = v8
        _ = v28
        _ = v29
    }
    v25 -= 2147483647
    fmt.Println(v6, v5)
    fmt.Println(v20, v25, v6)
    var v30 uint16 = 65535
    v2 = v2
    fmt.Println(v8, v6, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
