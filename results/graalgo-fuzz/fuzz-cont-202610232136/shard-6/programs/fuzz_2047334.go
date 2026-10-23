package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint) int8 {
    if false {
        p0 += (p0 * 255)
        p0 = p0
        var v4 bool = true
        _ = v4
    }
    p0 = 106
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: -1, F1: 28, F2: 65535}
        v7 := GenS2{F0: -390, F1: 1, F2: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = (p0 | 65535)
    _ = p0
    return -128
}
func genFn1(p0 int32) uint8 {
    v8 := []int{}
    p0 = (p0 % (((^(p0))) | 1))
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 744
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v9 uint8 = genFn1(-1)
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint32 = 934
        var v12 int64 = 172
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v9, v2, v9)
    v2 = "foo bar"
    if ((true || false) || ((v9 & v9) <= (v9 >> ((uint(v9)) & 63)))) {
        switch 9223372036854775807 {
        case -297:
            v13 := GenS1{F0: 951, F1: 42, F2: 42}
            _ = v13
        case 1:
            v9 = (v9 & 8)
        case 872:
            if false {
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        fmt.Println(fi14)
                        if (!(685 > 1)) {
                            v9 = (v9 >> ((18446744073709551615) & 63))
                        } else {
                            var v16 uint32 = 65535
                            _ = v16
                        }
                        v2 = "a"
                        _ = fi15
                    }
                    _ = fi14
                }
                v3[294] = uint32(v9)
                v17 := v3[667]
                _ = v17
            } else {
                var v18 uint64 = 276
                v19 := v3[0]
                fmt.Println(v18)
                _ = v18
                _ = v19
            }
        default:
            v2 = v2
        }
        if (!true) {
            if false {
                v20 := GenS1{F0: 0, F1: -128, F2: 2}
                v21 := GenS1{F0: int64(v9), F1: 127, F2: 255}
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS1{F0: -113, F1: -128, F2: 255}
                    v9 += v9
                    _ = fi22
                    _ = v23
                }
                _ = v20
                _ = v21
            }
        } else {
            fmt.Println(v2)
            if ((0.1 != 0.0) && (517 != 42)) {
                v2 = v2
            } else {
                v24 := GenS1{F0: 9223372036854775807, F1: genFn0(255), F2: 0}
                var v25 int64 = 1
                _ = v24
                _ = v25
            }
        }
        if false {
            v2 = "a"
            switch v2 {
            case "\n\t":
                v9 += 0
            case "foo bar":
                if false {
                    v9 += genFn1(-629)
                    v9 += (v9 * v9)
                    v26 := GenS2{F0: 0, F1: 81, F2: 0}
                    _ = v26
                } else {
                    fmt.Println(v2)
                    switch ((v2 + v2) + v2) {
                    case "a":
                        v9 = (v9 + v9)
                    case "\n\t":
                        fmt.Println(v2)
                    case "foo bar":
                        v1 -= (v1 << ((4294967295) & 63))
                    }
                    v27 := GenS2{F0: 1, F1: 0, F2: 100}
                    _ = v27
                }
            default:
                v28 := v3[355]
                _ = v28
            }
            fmt.Println(v2, v2)
        } else {
            var v29 int = -240
            var v30 uint32 = 65535
            var v31 uint8 = v9
            _ = v29
            _ = v30
            _ = v31
        }
    } else {
        v2 = "foo bar"
    }
    if false {
        v32 := []int{-9223372036854775808, 42, 32768, 0}
        if len(v32) > 0 {
            v32[0] = -414
        }
        _ = v32
    }
    v9 -= v9
    v2 = "hello"
    var v33 int16 = v1
    var v34 int8 = 0
    if ((!true) && false) {
        v34 = (v34 + v34)
        v35 := GenS2{F0: 0, F1: 127, F2: 1}
        switch 0 {
        case 9223372036854775807:
            v36 := GenS3{F0: 0, F1: v34}
            _ = v36
        case 2147483647:
            var v37 string = "a"
            _ = v37
        case 1:
            var v38 uint8 = (v9 + (v9 + 128))
            _ = v38
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v9
}
