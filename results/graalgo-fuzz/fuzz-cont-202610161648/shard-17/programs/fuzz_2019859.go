package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8) int16 {
    p0 += p0
    if true {
        p0 = (p0 / ((1) | 1))
    }
    _ = p0
    return -32768
}
func genFn1(p0 string, p1 uint8, p2 uint) int {
    p2 = uint(p1)
    v5 := GenS1{F0: -1, F1: false}
    switch 1 {
    case 32768:
        p0 = "foo bar"
    case 1:
        p2 -= 18446744073709551615
    case 935:
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := []int{}
            if len(v7) > 0 {
                v7[0] = (fi6 >> (((^(p2))) & 63))
            }
            p0 = p0
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: -2147483648, F1: false}
                p2 = 65535
                if (9223372036854775807 <= 42) {
                    v11 := []int{9223372036854775807, -2147483648, (fi9 * 42), (fi9 + (fi8 / ((128) | 1)))}
                    if (!(false || false)) {
                        fmt.Println(p0, fi8, fi9)
                        p2 = 42
                        var v12 uint64 = 42
                        _ = v12
                    }
                    if false {
                        fmt.Println(p1)
                    } else {
                        _ = len(v11)
                        switch 0 {
                        case 682:
                            p0 = ((p0 + (p0 + "\n\t")) + (p0 + "a"))
                        case 32768:
                            switch 0 {
                            case 203:
                                var v13 int8 = -128
                                _ = v13
                            case 983:
                                var v14 float32 = (3.14159 / ((float32(1.5) + (0.1 - 0.0)) - -2.25))
                                _ = v14
                            default:
                                var v15 uint = (p2 ^ 255)
                                _ = v15
                            }
                        }
                        v16 := GenS1{F0: fi9, F1: true}
                        _ = v16
                    }
                    var v17 int16 = -1
                    _ = v11
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(p0, p2, fi8)
                p2 += (p2 * (p2 % ((2) | 1)))
                v19 := GenS1{F0: fi8, F1: (("hello" + p0) == p0)}
                p0 = p0
                p1 = p1
                _ = fi18
                _ = v19
            }
            _ = fi8
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    var v1 int = -102
    var v2 bool = (18446744073709551615 < 0)
    v3 := []int{v1, (^(v1))}
    var v4 uint16 = 1
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v2)
    default:
        switch v1 {
        case 9223372036854775807:
            var v20 uint16 = 32767
            _ = v20
        case -9223372036854775808:
            var v21 uint8 = 127
            _ = v21
        }
    }
    var v22 uint = 1
    var v23 uint32 = 4294967295
    _ = len(v3)
    v22 += v22
    if len(v3) > 0 {
        v3[0] = 128
    }
    fmt.Println(v23)
    if ((-666 == 51) && v2) {
        _ = len(v3)
        var v24 bool = false
        v25 := GenS1{F0: (v1 << (((v22 << ((1) & 63))) & 63)), F1: false}
        _ = v24
        _ = v25
    } else {
        if v2 {
            fmt.Println(v23, v23, v2)
            switch (^(v1)) {
            case -1:
                v26 := GenS1{F0: v1, F1: (!(v2 || (v2 || false)))}
                _ = v26
            case 356:
                v4 = 171
            case 4294967295:
                var v27 uint32 = v23
                _ = v27
            default:
                if v2 {
                    fmt.Println(v1, v22)
                    v22 -= v22
                } else {
                    var v28 float32 = (-2.25 / (0.1 * (((0.1 + float32((float64(3.14159) * 0.1))) / 1e10) + (-2.25 / 1e10))))
                    v29 := GenS2{F0: (v2 && v2), F1: v23}
                    _ = v28
                    _ = v29
                }
            }
        } else {
            v30 := GenS1{F0: (v1 >> (((v22 & 2)) & 63)), F1: (4 >= 42)}
            var v31 int = (v1 % ((v1) | 1))
            if false {
                v23 = v23
                if ((1e10 * 3.14159) < (-2.25 * ((float32((1e10 + 1.5)) * float32((-2.25 + float64(1.5)))) * (1e10 - 1.5)))) {
                    v22 = v22
                    var v32 int64 = 9223372036854775807
                    var v33 uint16 = v4
                    _ = v32
                    _ = v33
                } else {
                    v1 = v31
                    switch "a" {
                    case "a":
                        if (!v2) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v3 = append(v3, 32767)
                        }
                    }
                    var v34 int64 = 42
                    _ = v34
                }
            }
            _ = v30
            _ = v31
        }
        v23 = 1
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
