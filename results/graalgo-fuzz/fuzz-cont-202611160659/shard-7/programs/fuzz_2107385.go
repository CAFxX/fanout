package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (1e10 + 0.0)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{263, 0}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            v2 -= -2.25
            var v7 int8 = 100
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 32767, F1: fi5}
        _ = fi5
        _ = v8
    }
    v2 = ((v2 + 3.14159) - 0.0)
    v2 += 3.14159
    if (!true) {
        v9 := GenS1{F0: 65535, F1: 9223372036854775807}
        _ = len(v4)
        _ = v9
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: false, F1: 0, F2: 1}
        var v12 int16 = -32768
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        delete(v3, 21)
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := v3[2]
            v15 += v15
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch 131 {
    case 0:
        v16 := GenS3{F0: 1, F1: 4294967295, F2: 563}
        _ = v16
    case -899:
        for fi17 := 0; fi17 < 18; fi17++ {
            switch (fi17 ^ (^(fi17))) {
            case 9223372036854775807:
                v18 := GenS4{F0: 127, F1: 406}
                _ = v18
            case 4294967295:
                delete(v3, 572)
            default:
                fmt.Println(fi17, fi17)
            }
            _ = fi17
        }
    default:
        var v19 int = 4294967295
        _ = v19
    }
    var v20 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
