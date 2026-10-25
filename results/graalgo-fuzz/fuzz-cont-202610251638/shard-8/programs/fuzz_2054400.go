package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{2147483647, 472, 2}
    v3 := make(map[int]int)
    var v4 uint32 = 2
    _ = len(v2)
    _ = len(v2)
    v5 := GenS1{F0: -269, F1: 2}
    var v6 int8 = -1
    if ((^(v1)) > (^(v1))) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS2{F0: 42, F1: uint16(v4), F2: 286}
            v9 := GenS3{F0: 1}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint8 = 255
        _ = len(v2)
        _ = v10
    } else {
        if (95 <= (^(v6))) {
            var v11 int8 = v6
            v6 = (-(v6))
            _ = v11
        } else {
            v4 += v4
            var v12 uint8 = 56
            var v13 uint32 = 498
            _ = v12
            _ = v13
        }
        v14 := GenS4{F0: 1}
        _ = v14
    }
    if (!false) {
        v15 := GenS5{F0: ""}
        var v16 int32 = -1
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(v6, v1, v6)
            v18 := GenS5{F0: ("hello" + ("a" + "foo bar"))}
            _ = fi17
            _ = v18
        }
    }
    switch 32768 {
    case 447:
        v6 = v6
    case 319:
        v4 += 2147483647
    case 9223372036854775807:
        fmt.Println(v1, v6, v6)
    default:
        v19 := GenS3{F0: -690}
        _ = v19
    }
    v2 = append(v2, v4)
    var v20 uint16 = 0
    v21 := GenS5{F0: (("hello" + "") + "foo bar")}
    var v22 int64 = 42
    fmt.Println(v20, v1)
    var v23 bool = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
