package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 uint, p1 uint16, p2 int16) uint32 {
    p2 += p2
    p0 = p0
    var v4 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{}
    var v3 int = -2147483648
    switch v3 {
    case 1:
        v5 := GenS1{F0: 128, F1: 446, F2: "a"}
        _ = v5
    }
    v3 -= 32768
    if false {
        v3 -= v3
        v1[(^(v3))] = 813
        if (((0.1 + float32(((((0.1 + 0.0) + ((0.0 * float64(1.5)) * 3.14159)) * 1.5) / (1e10 * float64((float32(-2.25) / float32(float64(-2.25)))))))) < 1e10) || true) {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint16 = 2
                fmt.Println(fi6, v7, v3)
                switch "hello" {
                case "foo bar":
                    v3 += fi6
                case "a":
                    if true {
                        var v8 bool = (100 >= 42)
                        _ = v8
                    }
                default:
                    v3 -= (^(v3))
                }
                _ = fi6
                _ = v7
            }
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                v3 = 718
                v3 += (fi9 >> ((4294967295) & 63))
                if (1 > fi9) {
                    var v10 uint8 = 255
                    _ = v10
                }
                _ = fi9
            }
            v3 = v3
            v11 := GenS2{F0: v3, F1: (v3 - (v3 ^ v3))}
            _ = v11
        }
    } else {
        v3 = (v3 / ((-1) | 1))
        v2 = append(v2, 65535)
        fmt.Println(v3)
    }
    var v12 int8 = 0
    v13 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
    switch (^(v3)) {
    case -1:
        var v14 int64 = 42
        _ = v14
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v15 int8 = -1
    if (-528 != -32768) {
        var v16 string = "a"
        v12 -= -128
        switch v16 {
        case "":
            if len(v2) > 0 {
                v2[0] = 42
            }
        case "hello":
            if false {
                fmt.Println(v16)
            }
        case "foo bar":
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(v15, v15)
                v18 := GenS1{F0: 128, F1: 42, F2: (v16 + v16)}
                var v19 uint64 = 0
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            var v20 uint8 = 101
            _ = v20
        }
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
