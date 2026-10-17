package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int) int16 {
    v4 := []int{p0, 128, p0}
    if (true || (false && false)) {
        if (0 != 883) {
            var v5 uint8 = 42
            _ = v5
        } else {
            v4 = append(v4, p0)
            fmt.Println(p0)
            v6 := GenS1{F0: -1}
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 -= fi7
            _ = fi7
        }
        var v8 uint64 = 942
        if true {
            v8 = 4294967295
            var v9 int16 = int16(p0)
            _ = v9
        } else {
            fmt.Println(v8)
            var v10 uint8 = 1
            _ = v10
        }
        _ = v8
    }
    if false {
        v11 := GenS2{F0: 18446744073709551615, F1: -32768, F2: (p0 & 851)}
        v12 := GenS1{F0: -32768}
        p0 += (p0 >> ((0) & 63))
        _ = v11
        _ = v12
    } else {
        var v13 uint64 = 733
        if (8 == -20) {
            fmt.Println(p0, p0)
            for fi14 := 0; fi14 < 1; fi14++ {
                if (18446744073709551615 != 384) {
                    fmt.Println(v13, p0)
                } else {
                    _ = len(v4)
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS1{F0: 211}
                    p0 = 4294967295
                    fmt.Println(fi15, v13, fi14)
                    _ = fi15
                    _ = v16
                }
                switch "foo bar" {
                case "hello":
                    for fi17 := 0; fi17 < 1; fi17++ {
                        p0 -= fi17
                        v13 += 745
                        fmt.Println(p0, fi14)
                        _ = fi17
                    }
                case "a":
                    var v18 int = -371
                    _ = v18
                case "":
                    var v19 int32 = 1
                    _ = v19
                }
                _ = fi14
            }
        } else {
            var v20 int = (^(p0))
            _ = v20
        }
        if (2147483647 >= int32(v13)) {
            if true {
                var v21 uint32 = 4294967295
                v22 := GenS2{F0: (v13 - (v13 | 4294967295)), F1: 0, F2: (-(p0))}
                for fi23 := 0; fi23 < 12; fi23++ {
                    if (0 <= uint16(v21)) {
                        var v24 uint8 = 255
                        v25 := GenS3{F0: 0}
                        _ = v24
                        _ = v25
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v26 string = "hello"
                            _ = v26
                        case "a":
                            if ((1 != 2) && (((100 != -105) && ((((!((v21 ^ v21) > (v21 % ((v21) | 1)))) || true) && (float64(0.0) >= -2.25)) || (0 > 450))) && (0 <= 4294967295))) {
                                v13 = (v13 % ((v13) | 1))
                                var v27 bool = (926 < 815)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v27
                            }
                        default:
                            var v28 string = "a"
                            _ = v28
                        }
                        v29 := GenS4{F0: v21}
                        _ = v29
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
            var v30 uint16 = 0
            _ = v30
        }
        _ = v13
    }
    p0 = (p0 << ((0) & 63))
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{790, 2}
    v3 := []uint32{0, 4294967295, 154}
    switch 837 {
    case 32768:
        switch 32768 {
        case -9223372036854775808:
            var v31 int8 = 100
            _ = v31
        }
    case 128:
        _ = len(v1)
    default:
        _ = len(v3)
    }
    var v32 uint64 = 846
    v32 += (v32 ^ v32)
    v32 = v32
    fmt.Println(v32)
    var v33 int = 127
    v34 := GenS5{F0: 0}
    fmt.Println(v33)
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS5{F0: 1}
        _ = fi35
        _ = v36
    }
    v37 := GenS5{F0: -1}
    v33 = (v33 * (v33 >> ((770) & 63)))
    var v38 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
}
