package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint16 }

func main() {
    var v1 int8 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := []uint32{2147483648}
    v1 -= 88
    var v5 int16 = 871
    var v6 uint16 = 0
    if (32767 == v5) {
        v7 := GenS1{F0: false, F1: -1, F2: 4294967295}
        var v8 int = 42
        v4 = append(v4, 2)
        _ = v7
        _ = v8
    } else {
        v9 := GenS2{F0: "a"}
        v10 := GenS2{F0: "foo bar"}
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: "a"}
    switch 127 {
    case 32768:
        if (true || false) {
            var v12 int16 = -313
            if false {
                fmt.Println(v5)
                if false {
                    var v13 int32 = 0
                    v1 -= v1
                    var v14 uint = 18446744073709551615
                    _ = v13
                    _ = v14
                }
                if len(v4) > 0 {
                    v4[0] = 137
                }
            } else {
                switch 1 {
                case -1:
                    v2[0] = 2147483647
                case -2147483648:
                    v12 = (v12 & -32768)
                case 0:
                    var v15 uint16 = (v6 | (^(v6)))
                    _ = v15
                }
                v2[4294967295] = 592
                v12 += (v5 * v12)
            }
            v16 := GenS2{F0: ""}
            _ = v12
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS2{F0: "a"}
                if true {
                    v20 := v2[(fi17 & (fi18 >> ((1) & 63)))]
                    v21 := GenS3{F0: (v20 ^ v20), F1: -2147483648}
                    v5 -= 1000
                    _ = v20
                    _ = v21
                } else {
                    v1 = (-(v1))
                }
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    var v22 uint = 2
    var v23 uint = (v22 + 255)
    v5 = -210
    v24 := GenS4{F0: 1, F1: 1, F2: 2}
    var v25 uint = (v23 ^ (v23 / ((v22) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
