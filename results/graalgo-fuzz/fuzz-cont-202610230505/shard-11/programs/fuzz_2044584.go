package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }

func main() {
    var v1 bool = true
    v2 := []uint32{0, 4294967295}
    v3 := make(map[uint32]bool)
    v4 := GenS1{F0: "a", F1: 2147483647, F2: 42}
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = v1
        _ = fi5
    }
    if (false || (false || v1)) {
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := v3[4294967295]
                v9 := GenS2{F0: 2, F1: 277, F2: 32767}
                var v10 uint8 = 0
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            var v11 int = fi6
            if (32767 == 8) {
                fmt.Println(v11, fi6, v11)
            } else {
                if (true || (9223372036854775807 >= 9223372036854775807)) {
                    if ((-4 <= 0) && v1) {
                        v11 = 684
                        v12 := GenS3{F0: 42, F1: (v1 || true), F2: 1}
                        _ = v12
                    }
                    var v13 float32 = float32(-2.25)
                    _ = v13
                }
                v1 = (659 != 470)
            }
            _ = fi6
            _ = v11
        }
    } else {
        v14 := GenS3{F0: 100, F1: v1, F2: 202}
        _ = v14
    }
    v1 = v1
    v15 := v3[42]
    switch 120 {
    case 42:
        fmt.Println(v1, v15)
    default:
        v15 = true
    }
    v1 = (v15 || true)
    fmt.Println(v15, v15)
    if (!v1) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    _ = len(v2)
    var v16 int64 = -158
    v2 = append(v2, 1)
    var v17 int64 = (^(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
