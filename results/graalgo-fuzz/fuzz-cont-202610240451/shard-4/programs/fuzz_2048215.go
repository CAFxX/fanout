package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
func genFn0() uint {
    var v4 uint = 984
    switch 98 {
    case 32768:
        if ((true || false) && false) {
            var v5 uint32 = 2
            _ = v5
        }
    case 128:
        v4 = (^(v4))
    case 32767:
        if false {
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := []int{(fi6 << ((42) & 63)), (-(fi6)), (fi6 * fi6), (fi6 & int(v4))}
                if (42 >= 77) {
                    var v8 int16 = 0
                    _ = v8
                } else {
                    switch "hello" {
                    case "":
                        if (0 != 32767) {
                            for fi9 := 0; fi9 < 5; fi9++ {
                                v4 -= v4
                                v4 -= (v4 | (v4 >> ((v4) & 63)))
                                _ = fi9
                            }
                            var v10 int16 = 32767
                            _ = v10
                        }
                    case "a":
                        fmt.Println(v4, fi6)
                    }
                }
                _ = fi6
                _ = v7
            }
            if (2147483647 >= 42) {
                var v11 uint = (v4 - v4)
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: 42}
    _ = v12
    _ = v4
    return (v4 << ((18446744073709551615) & 63))
}
func genFn1(p0 uint64, p1 uint, p2 uint8) int32 {
    var v13 uint = (p1 << (((^(p1))) & 63))
    var v14 uint16 = 1
    v14 += uint16(p0)
    v15 := GenS1{F0: 112}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = -2.25
    v3 := make(map[uint32]bool)
    if (true || (!(42 < 127))) {
        for fi16 := 0; fi16 < 16; fi16++ {
            if ((1 <= 1) && (-1 != -128)) {
                if ((-2147483648 <= -2147483648) && false) {
                    v2 += v2
                } else {
                    v1 += (fi16 & -2147483648)
                }
                var v17 uint64 = 563
                _ = v17
            }
            _ = fi16
        }
    }
    v18 := GenS2{F0: 42, F1: 255, F2: 799}
    v19 := GenS3{F0: 0, F1: 4294967295}
    v1 -= v1
    switch "hello" {
    case "\n\t":
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint16 = 2
            var v22 int16 = -1
            _ = fi20
            _ = v21
            _ = v22
        }
    case "hello":
        v23 := GenS2{F0: 649, F1: 255, F2: 32767}
        _ = v23
    }
    v24 := GenS3{F0: -1, F1: 0}
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v3[593] = false
            v2 = v2
            _ = fi26
        }
        var v27 uint64 = 1
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS4{F0: 2, F1: 1, F2: 390}
            _ = len(v3)
            _ = fi28
            _ = v29
        }
        _ = fi25
        _ = v27
    }
    v2 = float64(1.5)
    var v30 uint64 = 18446744073709551615
    v2 -= v2
    fmt.Println(v30)
    v30 -= v30
    var v31 int = (v1 / ((-2147483648) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
}
