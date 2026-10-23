package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int }

func main() {
    var v1 uint8 = 0
    var v2 int32 = 2147483647
    var v3 int8 = 1
    v4 := []string{"\n\t", ("hello" + "foo bar"), ((("a" + "\n\t") + "\n\t") + "")}
    v5 := make(map[uint32]uint32)
    _ = len(v4)
    v6 := GenS1{F0: 581, F1: 127, F2: 1}
    v2 = v2
    if (!false) {
        var v7 string = ("hello" + "a")
        _ = v7
    } else {
        var v8 uint8 = (v1 >> ((42) & 63))
        fmt.Println(v1)
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if true {
            if (239 >= v1) {
                var v9 int64 = 0
                v10 := GenS2{F0: 0, F1: -1, F2: (v3 + (v3 >> ((4294967295) & 63)))}
                v3 = (-(v3))
                _ = v10
                _ = v9
            }
            var v11 uint16 = 900
            fmt.Println(v1, v3, v3)
            _ = v11
        }
        v12 := GenS3{F0: 572, F1: 255}
        v2 = v2
        _ = v12
    } else {
        v2 -= v2
        v2 = v2
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v3, v3)
        fmt.Println(v1, v2, v3)
        _ = fi13
    }
    fmt.Println(v1)
    fmt.Println(v1, v3, v3)
    v14 := GenS3{F0: 65535, F1: 265}
    v15 := GenS4{F0: 255, F1: false}
    var v16 uint64 = 18446744073709551615
    v17 := GenS5{F0: 127}
    v18 := GenS4{F0: uint16(v16), F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
