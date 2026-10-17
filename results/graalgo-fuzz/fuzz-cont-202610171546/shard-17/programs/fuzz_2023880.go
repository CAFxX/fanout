package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0(p0 int, p1 uint64, p2 uint32) uint32 {
    var v4 int = p0
    if (2 >= 42) {
        fmt.Println(v4, p0)
        var v5 int16 = 851
        for fi6 := 0; fi6 < 10; fi6++ {
            v4 -= 4294967295
            v7 := GenS1{F0: 32767, F1: 0, F2: 304}
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    v1 := []int{-9223372036854775808, 42}
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    var v8 int = -2147483648
    _ = len(v2)
    v9 := GenS2{F0: (v3 * (v3 | v3))}
    v3 = v3
    switch 127 {
    case -921:
        var v10 float32 = (0.1 + -2.25)
        _ = v10
    default:
        v3 -= (v3 & (v3 - -1))
    }
    if false {
        var v11 string = ""
        _ = v11
    } else {
        v12 := GenS1{F0: int16(v8), F1: 100, F2: 196}
        v3 += (v3 / (((v3 ^ v3)) | 1))
        v13 := GenS3{F0: 610, F1: 2}
        _ = v12
        _ = v13
    }
    switch "foo bar" {
    case "foo bar":
        for fi14 := 0; fi14 < 12; fi14++ {
            v3 -= v3
            _ = fi14
        }
    case "\n\t":
        if (!true) {
            v1 = append(v1, 32768)
            if true {
                v15 := GenS2{F0: v3}
                if false {
                    fmt.Println(v8, v3)
                }
                v8 = v8
                _ = v15
            }
        }
    case "a":
        v8 = (v8 ^ (v8 << ((255) & 63)))
    default:
        fmt.Println(v3, v3)
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        switch ("hello" + ("a" + "hello")) {
        case "foo bar":
            v17 := GenS4{F0: -1, F1: -1, F2: 4294967295}
            _ = v17
        case "\n\t":
            if ((3.14159 == float64(float32((float64(float32(0.1)) * -2.25)))) && (2147483647 > 2147483647)) {
                for fi18 := 0; fi18 < 13; fi18++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi18
                }
                _ = len(v1)
            } else {
                v19 := v2[int(v3)]
                v19 += genFn0((fi16 ^ (^(v8))), 298, (v19 - (v19 + (v19 % ((4294967295) | 1)))))
                v20 := GenS5{F0: 42}
                _ = v19
                _ = v20
            }
        }
        v8 = (^(v8))
        _ = fi16
    }
    if (true && true) {
        v21 := GenS5{F0: 42}
        fmt.Println(v8)
        if (525 > 1) {
            switch "a" {
            case "\n\t":
                v22 := GenS1{F0: -1, F1: 0, F2: 51}
                _ = v22
            case "":
                if ((!(1 < 45)) && ((("foo bar" + ("foo bar" + "foo bar")) <= "\n\t") && false)) {
                    v3 = (v3 << ((0) & 63))
                } else {
                    v3 = (-(v3))
                    if false {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if (646 < 2) {
                                v3 -= v3
                                var v24 float64 = 3.14159
                                _ = v24
                            } else {
                                v25 := v2[(fi23 ^ fi23)]
                                _ = v25
                            }
                            _ = fi23
                        }
                        fmt.Println(v3, v3)
                        _ = len(v2)
                    }
                    v3 = -9223372036854775808
                }
            case "hello":
                v26 := GenS4{F0: 2147483647, F1: 598, F2: 2}
                _ = v26
            }
            switch 2147483647 {
            case -9223372036854775808:
                for fi27 := 0; fi27 < 15; fi27++ {
                    if (2147483648 != 756) {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v29 := GenS5{F0: 1}
                            if true {
                                _ = len(v2)
                            } else {
                                v3 -= (-(v3))
                                v30 := GenS2{F0: v3}
                                _ = v30
                            }
                            v3 -= (v3 | (v3 * v3))
                            _ = fi28
                            _ = v29
                        }
                    }
                    v3 = -1
                    fmt.Println(v8, v8, v3)
                    _ = fi27
                }
            default:
                v31 := GenS6{F0: 17, F1: 108, F2: 121}
                _ = v31
            }
            v3 -= v3
        } else {
            v32 := GenS5{F0: 284}
            _ = v32
        }
        _ = v21
    }
    var v33 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v8
    _ = v9
}
