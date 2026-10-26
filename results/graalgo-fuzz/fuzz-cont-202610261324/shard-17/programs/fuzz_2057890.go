package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 int8) uint64 {
    p0 += p0
    var v6 uint64 = 4294967295
    switch -1 {
    case -9223372036854775808:
        p0 -= p0
    case 32768:
        v6 -= 4294967295
    case 0:
        switch "\n\t" {
        case "\n\t":
            v6 = v6
        case "":
            var v7 int = -9223372036854775808
            _ = v7
        case "hello":
            v8 := []int{32767, 1, 959, 4294967295}
            _ = v8
        default:
            var v9 float64 = (((3.14159 + 0.1) * -2.25) + (3.14159 * 1.5))
            _ = v9
        }
    default:
        var v10 uint32 = 65535
        _ = v10
    }
    var v11 uint8 = 1
    var v12 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v6
    return (v6 - v6)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 1
    v3 := []uint32{0}
    var v4 int8 = 1
    v5 := []uint32{}
    var v13 int8 = (v4 ^ (v4 << (((v2 * v2)) & 63)))
    v4 = v13
    var v14 int16 = -32768
    var v15 int16 = v14
    var v16 int16 = (v14 | (v14 + v15))
    if false {
        v17 := GenS1{F0: true, F1: 1, F2: (^(v4))}
        if (v4 >= (v13 / (((v4 >> ((0) & 63))) | 1))) {
            _ = len(v3)
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 string = "hello"
                for fi20 := 0; fi20 < 17; fi20++ {
                    switch -1 {
                    case 32768:
                        fmt.Println(v15, fi18)
                    case 32767:
                        var v21 int = 2147483647
                        _ = v21
                    case 9223372036854775807:
                        v14 -= 739
                    default:
                        v4 -= v13
                        fmt.Println(v14)
                        var v22 uint16 = 255
                        _ = v22
                    }
                    v23 := GenS1{F0: false, F1: -386, F2: v13}
                    if (false && false) {
                        v13 -= (v4 / (((v4 & 1)) | 1))
                        if (genFn0(471, v4) <= 42) {
                            if false {
                                v4 -= (v4 ^ (v13 ^ v4))
                            } else {
                                _ = len(v5)
                                var v24 int8 = (v4 << ((760) & 63))
                                _ = v24
                            }
                            v25 := v1[0]
                            delete(v1, v25)
                            _ = v25
                        }
                    } else {
                        v15 = v14
                        if (4294967295 > 0) {
                            v13 = (v4 | 100)
                            v2 = (^(v2))
                        } else {
                            v5 = append(v5, 2147483648)
                            var v26 string = (v19 + (v19 + (v19 + (v19 + v19))))
                            _ = v26
                        }
                        var v27 uint32 = 42
                        delete(v1, v27)
                        if (1.5 != 0.0) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v28 int8 = -42
                        _ = v27
                        _ = v28
                    }
                    _ = fi20
                    _ = v23
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    v14 = (^(v15))
                    var v30 bool = (!(uint16(v2) < 209))
                    fmt.Println(fi18, fi18)
                    _ = fi29
                    _ = v30
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v16, v2, v13)
        } else {
            switch "" {
            case "":
                for fi31 := 0; fi31 < 7; fi31++ {
                    var v32 uint64 = 391
                    _ = fi31
                    _ = v32
                }
            }
        }
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            for fi33 := 0; fi33 < 11; fi33++ {
                v5 = append(v5, 0)
                _ = fi33
            }
        case "a":
            var v34 uint = v2
            _ = v34
        }
        _ = v17
    } else {
        v35 := GenS2{F0: ("a" + "a")}
        var v36 uint16 = 541
        v13 = v13
        _ = v35
        _ = v36
    }
    var v37 bool = (!((65535 <= 2147483648) || (-1 >= -261)))
    var v38 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
