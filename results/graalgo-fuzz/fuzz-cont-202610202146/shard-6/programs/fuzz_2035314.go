package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 int8, p2 int64) uint {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{-2147483648, (fi5 | (fi5 ^ (^(fi6)))), (fi5 + (^(fi6))), fi5}
            p0 = (p0 >> ((65535) & 63))
            if (3.14159 <= 0.0) {
                var v8 uint16 = 547
                var v9 int8 = (p1 + p1)
                var v10 int16 = -32768
                _ = v10
                _ = v8
                _ = v9
            } else {
                fmt.Println(fi6)
                v7 = append(v7, (fi6 << ((65535) & 63)))
            }
            _ = fi6
            _ = v7
        }
        var v11 int = (fi5 << ((18446744073709551615) & 63))
        _ = fi5
        _ = v11
    }
    v12 := []int{0, 42, -2147483648, 768}
    for fi13 := 0; fi13 < 14; fi13++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(fi14, p2, fi14)
            fmt.Println(p1, p1)
            p2 = p0
            _ = fi14
        }
        var v15 bool = true
        _ = fi13
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 255
}
func genFn1() int64 {
    if (true && ((65535 != 42) || true)) {
        var v16 uint = 628
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint32 = 1
            _ = fi17
            _ = v18
        }
        var v19 int32 = 0
        fmt.Println(v19)
        _ = v19
    }
    var v20 float64 = (-2.25 * -2.25)
    _ = v20
    return 708
}

func main() {
    var v1 uint64 = 389
    v2 := make(map[int]bool)
    var v3 uint32 = 22
    var v4 int = 42
    _ = len(v2)
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int64 = 1
        for fi23 := 0; fi23 < 6; fi23++ {
            switch "hello" {
            case "foo bar":
                if false {
                    fmt.Println(v3, v22, v22)
                }
            case "hello":
                switch (fi21 / ((fi23) | 1)) {
                case 127:
                    if false {
                        v4 -= (fi23 + (fi21 / ((-437) | 1)))
                    }
                case -9223372036854775808:
                    v3 = v3
                default:
                    v4 += 735
                }
            case "a":
                var v24 int32 = 905
                _ = v24
            default:
                fmt.Println(v3, v4, v3)
            }
            v25 := []int{fi23, (^(v4)), (fi23 >> ((57) & 63)), (fi21 + (fi23 + 2147483647))}
            var v26 uint64 = (v1 + (v1 | 18446744073709551615))
            _ = fi23
            _ = v25
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    if false {
        v27 := GenS1{F0: (v3 + (v3 * (v3 % ((286) | 1))))}
        if ((v4 > v4) && false) {
            v28 := GenS2{F0: 65535, F1: 128}
            v29 := GenS2{F0: 42, F1: 41}
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 int = (v4 * (fi30 * fi30))
            _ = fi30
            _ = v31
        }
        _ = v27
    }
    var v32 int64 = 0
    fmt.Println(v1)
    v33 := GenS3{F0: v3, F1: 42}
    v34 := GenS1{F0: (v3 | v3)}
    v35 := GenS4{F0: 42}
    var v36 uint64 = (v1 % ((v1) | 1))
    v32 += int64(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
