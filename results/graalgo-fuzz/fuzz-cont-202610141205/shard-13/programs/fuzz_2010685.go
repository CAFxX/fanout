package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint16 }
func genFn0(p0 uint16, p1 int32, p2 bool) uint {
    var v4 uint16 = p0
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 12; fi6++ {
        v4 = 0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint8 {
    v7 := GenS1{F0: 32767, F1: 4294967295}
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint = genFn0(1, -1, (true && (1 != 127)))
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -94
    var v2 int = 1
    var v3 float64 = (float64((-2.25 + float32((-2.25 / -2.25)))) - 3.14159)
    v3 = (v3 * v3)
    switch v2 {
    case -9223372036854775808:
        var v11 bool = false
        _ = v11
    case 32768:
        v3 += ((3.14159 + 0.1) / (-2.25 + v3))
    default:
        var v12 float32 = 1e10
        _ = v12
    }
    var v13 uint32 = 2147483647
    v14 := GenS2{F0: (v13 << ((42) & 63)), F1: v2}
    v13 -= 2147483648
    if (!true) {
        switch v2 {
        case 127:
            for fi15 := 0; fi15 < 17; fi15++ {
                v2 = (v2 >> ((76) & 63))
                _ = fi15
            }
        case -2147483648:
            if true {
                var v16 bool = false
                var v17 uint8 = 127
                _ = v16
                _ = v17
            } else {
                if true {
                    v18 := GenS1{F0: 1000, F1: 42}
                    var v19 uint16 = 1
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v19)
                    }
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        if ((true && ((v1 & v1) > (v1 * 2147483647))) || ((0.0 / 1e10) <= float32(v3))) {
                            v22 := GenS2{F0: (^(v13)), F1: (fi20 >> ((316) & 63))}
                            _ = v22
                        }
                        var v23 float64 = v3
                        var v24 bool = (65535 < 2)
                        v25 := []int{v2, 128, (fi20 << ((uint(v13)) & 63)), 32767}
                        _ = fi21
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    if true {
                        v1 = -2147483648
                        var v26 int16 = 1000
                        fmt.Println(fi20, v3, v26)
                        _ = v26
                    } else {
                        v2 -= v2
                        v3 -= v3
                    }
                    _ = fi20
                }
                v3 = 3.14159
            }
        case 9223372036854775807:
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS2{F0: (v13 - (v13 - v13)), F1: (v2 >> ((genFn0(255, 501, ((fi27 << ((4294967295) & 63)) < fi27))) & 63))}
                _ = fi27
                _ = v28
            }
        }
        v29 := GenS3{F0: 0}
        var v30 int = v2
        _ = v29
        _ = v30
    }
    fmt.Println(v1, v13, v13)
    for fi31 := 0; fi31 < 2; fi31++ {
        if true {
            v2 += v2
            v2 = 128
            v2 += (fi31 | (v2 << ((603) & 63)))
        }
        if true {
            fmt.Println(v3, v2, v3)
            v32 := GenS4{F0: 414, F1: 32767, F2: 42}
            _ = v32
        }
        _ = fi31
    }
    v33 := GenS1{F0: 389, F1: 42}
    v34 := GenS5{F0: 697, F1: 373, F2: 508}
    fmt.Println(v13)
    fmt.Println(v1, v2)
    var v35 bool = (!true)
    var v36 uint32 = v13
    v13 -= (v36 / ((51) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
