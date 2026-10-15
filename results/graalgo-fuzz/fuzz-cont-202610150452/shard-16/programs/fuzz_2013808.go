package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
type GenS4 struct { F0 int16; F1 uint }
func genFn0() uint32 {
    v7 := []int{4294967295}
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    var v8 uint16 = 2
    _ = v7
    _ = v8
    return 299
}
func genFn1(p0 int8, p1 int16, p2 uint64) int32 {
    fmt.Println(p1)
    var v9 float32 = 0.1
    var v10 float32 = v9
    v11 := GenS1{F0: (p0 >> ((65535) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    v2 := []uint32{}
    var v3 uint32 = 1
    var v4 int16 = 1
    v5 := []string{"foo bar", "\n\t"}
    var v6 float32 = float32(0.1)
    if (255 <= 987) {
        v3 = (v3 * (^(v3)))
    }
    v1 += 42
    if len(v2) > 0 {
        v2[0] = genFn0()
    }
    var v12 uint8 = 255
    fmt.Println(v6)
    v13 := GenS2{F0: false, F1: uint64(v3)}
    for fi14 := 0; fi14 < 16; fi14++ {
        v4 = (v4 << ((4294967295) & 63))
        _ = fi14
    }
    v15 := GenS2{F0: false, F1: 2}
    if true {
        var v16 float32 = 3.14159
        if false {
            var v17 int16 = v4
            fmt.Println(v12, v17)
            v18 := GenS3{F0: -1, F1: 18446744073709551615, F2: 0}
            _ = v17
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS4{F0: v4, F1: 1}
        var v20 int32 = 42
        switch -451 {
        case 32767:
            v20 = genFn1(42, v4, 199)
        case 128:
            var v21 int32 = (-(v20))
            _ = v21
        }
        _ = v19
        _ = v20
    }
    if true {
        for fi22 := 0; fi22 < 5; fi22++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 -= (v1 << ((0) & 63))
            _ = fi22
        }
    } else {
        var v23 int16 = (-(v4))
        switch 128 {
        case 9223372036854775807:
            if true {
                v6 += ((1e10 - v6) * v6)
                v12 = 128
            } else {
                v2 = append(v2, v3)
            }
        case -554:
            var v24 float64 = 3.14159
            _ = v24
        default:
            v25 := GenS4{F0: v23, F1: 4294967295}
            _ = v25
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        switch fi26 {
        case 2147483647:
            var v27 uint16 = v1
            _ = v27
        default:
            if false {
                v12 = v12
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 uint16 = v1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v30 uint64 = 0
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                var v31 float64 = (3.14159 - 0.0)
                _ = v31
            } else {
                var v32 int8 = 42
                for fi33 := 0; fi33 < 9; fi33++ {
                    var v34 int8 = v32
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
        }
        if false {
            if true {
                fmt.Println(v4, v12, v3)
                if (false && false) {
                    var v35 uint = 2
                    _ = v35
                }
            } else {
                v36 := GenS1{F0: 100}
                v4 = v4
                _ = v36
            }
            v1 = 2
            if false {
                if len(v2) > 0 {
                    v2[0] = (v3 >> ((278) & 63))
                }
                _ = len(v2)
            }
        }
        _ = fi26
    }
    var v37 uint16 = (v1 - (v1 - (v1 % ((42) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
