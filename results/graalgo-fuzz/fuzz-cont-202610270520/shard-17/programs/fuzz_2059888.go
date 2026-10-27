package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0() int32 {
    v7 := GenS1{F0: -9223372036854775808}
    v8 := GenS1{F0: 1}
    var v9 int64 = 42
    v10 := GenS2{F0: 255, F1: 816}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int16, p1 int8, p2 uint64) int8 {
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            switch (-(fi12)) {
            case 1:
                fmt.Println(p1, fi11)
            case 9223372036854775807:
                if (0 > 0) {
                    v13 := GenS1{F0: 9223372036854775807}
                    switch 434 {
                    case 624:
                        var v14 uint8 = 0
                        _ = v14
                    case 0:
                        v15 := GenS3{F0: -9223372036854775808, F1: 1, F2: 381}
                        _ = v15
                    }
                    if (42 > 9223372036854775807) {
                        var v16 int = 128
                        var v17 int64 = 42
                        v18 := []int{v16, (fi12 ^ 127)}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        if (-2.25 != float64(1e10)) {
                            var v19 uint8 = 2
                            v20 := GenS3{F0: 0, F1: 122, F2: 617}
                            p0 -= (^(p0))
                            var v21 uint64 = (p2 << ((4294967295) & 63))
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(fi12, p1)
                            var v22 int32 = -1
                            v23 := []int{fi12, (fi11 << ((255) & 63))}
                            _ = v22
                            _ = v23
                        }
                        v24 := []int{fi11}
                        _ = v24
                    }
                    _ = v13
                } else {
                    p2 = 389
                }
            default:
                var v25 uint32 = 65535
                _ = v25
            }
            _ = fi12
        }
        _ = fi11
    }
    p0 -= int16(p2)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (p1 | p1)
}

func main() {
    var v1 float32 = float32(((0.1 / (1.5 * -2.25)) / 1.5))
    var v2 int64 = 42
    v3 := []int{}
    v4 := make(map[string]int)
    var v5 uint16 = 0
    var v6 int8 = -128
    for fi26 := 0; fi26 < 6; fi26++ {
        if (335 > 65535) {
            var v27 string = "foo bar"
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int64 = v2
                _ = fi28
                _ = v29
            }
            v1 += 3.14159
            _ = v27
        }
        _ = fi26
    }
    v2 = (^(v2))
    v3 = append(v3, 4294967295)
    v6 = -1
    v2 = -1
    for fi30 := 0; fi30 < 1; fi30++ {
        v5 -= (v5 << ((591) & 63))
        var v31 int = -2147483648
        _ = fi30
        _ = v31
    }
    v3 = append(v3, -591)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
