package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 bool, p1 uint, p2 uint64) int8 {
    v5 := []int{35, -2147483648}
    for fi6 := 0; fi6 < 20; fi6++ {
        if p0 {
            p2 -= 0
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                if len(v5) > 0 {
                    v5[0] = (fi6 % ((128) | 1))
                }
                if p0 {
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                    p1 += (^(p1))
                    p2 = (p2 & 1)
                } else {
                    if (false || p0) {
                        if len(v5) > 0 {
                            v5[0] = (fi6 >> ((18446744073709551615) & 63))
                        }
                        v8 := GenS1{F0: 0, F1: p0, F2: (!(p0 || (!false)))}
                        fmt.Println(fi6)
                        _ = v8
                    } else {
                        if true {
                            var v9 bool = (127 > -103)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v9
                        }
                        p1 = (p1 | p1)
                    }
                    v10 := GenS1{F0: 0, F1: (!(uint16(p2) != 255)), F2: true}
                    _ = v10
                }
                v11 := GenS2{F0: -1, F1: ((((float64((((1.5 / 1.0) + (0.0 / -2.25)) / -2.25)) / 1e10) - 0.1) < 0.1) || (0 >= -665))}
                _ = fi7
                _ = v11
            }
        }
        v12 := GenS2{F0: 2147483647, F1: (!false)}
        var v13 int64 = 534
        _ = fi6
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 12
}
func genFn1(p0 int8) int {
    v14 := []int{}
    if false {
        var v15 int = 128
        _ = v15
    }
    var v16 bool = true
    _ = p0
    _ = v14
    _ = v16
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{860, 2, 2}
    var v3 uint32 = 189
    v4 := []string{"hello", ("foo bar" + (("hello" + "a") + "a")), "\n\t"}
    v2 = append(v2, 4294967295)
    v3 = (v3 + v3)
    if (65535 > 18446744073709551615) {
        var v17 uint8 = 2
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 uint = 4294967295
                if (v20 > v20) {
                    if true {
                        fmt.Println(v20, v20)
                        var v21 uint = v20
                        _ = v21
                    } else {
                        v22 := GenS3{F0: 2}
                        v23 := GenS2{F0: 679, F1: true}
                        v3 += 4294967295
                        _ = v22
                        _ = v23
                    }
                    v3 -= (v3 >> ((501) & 63))
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        v24 := GenS4{F0: -8}
        _ = v24
    }
    var v25 float32 = 3.14159
    fmt.Println(v3)
    v26 := GenS5{F0: 495, F1: -9223372036854775808}
    v3 = (v3 / (((v3 & (v3 | (v3 % ((2147483648) | 1))))) | 1))
    var v27 uint16 = 0
    v3 -= 1
    v28 := GenS2{F0: 0, F1: true}
    _ = len(v1)
    var v29 int = -2147483648
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v29 += v29
    if len(v2) > 0 {
        v2[0] = (v3 | 65535)
    }
    fmt.Println(v25)
    v29 = v29
    v25 += ((v25 / v25) + ((float32((-2.25 / 1.0)) * v25) * 3.14159))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
