package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint64, p1 uint, p2 int8) uint64 {
    fmt.Println(p0)
    p1 -= 356
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int32 = -2147483648
    var v2 int = 32767
    v3 := make(map[int]bool)
    v4 := []int{v2, 128, v2, (v2 * -2147483648)}
    var v5 uint32 = 4294967295
    v6 := make(map[uint32]uint32)
    v5 = v5
    fmt.Println(v1)
    var v7 string = ("a" + "foo bar")
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint32 = (^(v5))
        switch v2 {
        case 32768:
            var v11 float32 = 1e10
            _ = v11
        case 712:
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: (v8 * (v8 % ((128) | 1)))}
                v10 = uint32(v8)
                _ = fi12
                _ = v13
            }
        default:
            if true {
                for fi14 := 0; fi14 < 7; fi14++ {
                    fmt.Println(v10, v1)
                    switch -9223372036854775808 {
                    case -960:
                        fmt.Println(v5, v1, v7)
                    case -2147483648:
                        var v15 float32 = ((float32(3.14159) + -2.25) - 1e10)
                        _ = v15
                    case 2147483647:
                        v10 = (v5 - v10)
                    default:
                        v16 := GenS2{F0: v10, F1: v10, F2: 0}
                        _ = v16
                    }
                    var v17 int32 = (v1 & (v1 ^ (^(v1))))
                    v1 = v17
                    _ = fi14
                    _ = v17
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 float32 = -2.25
                    v5 -= (v5 << ((18446744073709551615) & 63))
                    v5 += v10
                    var v20 uint16 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v5 = (v10 ^ v5)
            } else {
                fmt.Println(v10, v1)
                fmt.Println(v7, v10, v8)
            }
        }
        _ = fi9
        _ = v10
    }
    v7 = ((v7 + "hello") + v7)
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v6[v5] = v5
            v23 := v6[4294967295]
            v23 += (v5 - (v23 * (v23 << ((481) & 63))))
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: 379, F1: 1}
        if true {
            var v25 int32 = (v1 * (v1 | (v1 << ((998) & 63))))
            if false {
                v1 = 0
                v7 = v7
                var v26 uint8 = 255
                _ = v26
            } else {
                var v27 uint64 = uint64(v8)
                if true {
                    fmt.Println(v1, v27)
                    fmt.Println(v25, v1, v2)
                }
                delete(v6, uint32(v27))
                _ = v27
            }
            v8 = v8
            _ = v25
        } else {
            if len(v4) > 0 {
                v4[0] = v2
            }
            var v28 float64 = -2.25
            _ = v28
        }
        _ = fi21
        _ = v24
    }
    var v29 int32 = v1
    if true {
        var v30 uint32 = (v5 >> ((1) & 63))
        v4 = append(v4, v2)
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 string = v7
            v30 = (v5 & v30)
            var v33 int8 = 0
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    } else {
        v7 = "\n\t"
    }
    switch v7 {
    case "\n\t":
        v3[-312] = true
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
