package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 uint16) uint32 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 20; fi5++ {
        if ((1 <= 42) && false) {
            if (float32((1.5 / 1.5)) > 1.5) {
                var v6 uint16 = p0
                var v7 uint16 = v6
                _ = v6
                _ = v7
            } else {
                var v8 int = fi5
                v9 := GenS1{F0: 219, F1: true, F2: 1}
                var v10 uint64 = 42
                _ = v10
                _ = v8
                _ = v9
            }
        }
        _ = fi5
    }
    var v11 int16 = 956
    _ = p0
    _ = v11
    return 2147483648
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint16 = 255
        _ = fi12
        _ = v13
    }
    v3["\n\t"] = 282
    v14 := GenS2{F0: 42}
    v1 = v1
    v1 = (v1 >> ((4294967295) & 63))
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS3{F0: 1, F1: 4294967295, F2: fi15}
        var v17 float32 = 0.1
        var v18 uint64 = 48
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint8 = 1
    var v20 int32 = (v1 * v1)
    switch "foo bar" {
    case "a":
        var v21 int = -544
        _ = v21
    case "hello":
        delete(v3, "\n\t")
    }
    delete(v4, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
