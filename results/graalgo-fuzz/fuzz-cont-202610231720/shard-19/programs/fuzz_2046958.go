package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
func genFn0() int8 {
    v4 := GenS1{F0: 32768, F1: -9223372036854775808, F2: 13}
    v5 := []int{127, 32768}
    _ = v4
    _ = v5
    return 79
}
func genFn1(p0 string, p1 int16) uint32 {
    switch p0 {
    case "hello":
        var v6 float64 = (1.5 * (1.5 + ((float64((-2.25 * float32(1e10))) + 3.14159) * 0.0)))
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: 1, F1: 187, F2: 139}
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 19; fi8++ {
            switch (p0 + p0) {
            case "hello":
                var v9 float64 = -2.25
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: (fi8 % ((fi8) | 1)), F1: 765, F2: 127}
                _ = v10
            }
            for fi11 := 0; fi11 < 10; fi11++ {
                if true {
                    var v12 int = (fi8 * -427)
                    _ = v12
                } else {
                    p1 += p1
                }
                v13 := []int{fi11}
                p0 = "a"
                _ = fi11
                _ = v13
            }
            _ = fi8
        }
    default:
        p1 -= 42
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint32 = 884
        _ = fi14
        _ = v15
    }
    var v16 string = (("a" + "") + "\n\t")
    _ = p0
    _ = p1
    _ = v16
    return 1
}

func main() {
    var v1 uint = 700
    v2 := make(map[uint32]uint32)
    v3 := []uint32{uint32(v1), 42, 795}
    v1 -= v1
    var v17 uint64 = 547
    if false {
        var v18 float64 = (((((3.14159 / float64(((3.14159 * float32(-2.25)) + 0.1))) + 0.1) / 3.14159) + 0.0) / 3.14159)
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint16 = 533
            _ = fi19
            _ = v20
        }
        if (!true) {
            var v21 int64 = -1
            var v22 uint8 = 128
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(v21)
                _ = fi23
            }
            _ = v21
            _ = v22
        } else {
            var v24 int8 = 84
            _ = v24
        }
        _ = v18
    } else {
        fmt.Println(v17, v1)
        _ = len(v3)
        v25 := GenS1{F0: -505, F1: -433, F2: 46}
        _ = v25
    }
    var v26 int64 = 151
    v17 = (v17 | (v17 % ((0) | 1)))
    v27 := GenS1{F0: 128, F1: -9223372036854775808, F2: 0}
    var v28 uint64 = v17
    _ = len(v2)
    var v29 int64 = (^(v26))
    switch -2147483648 {
    case 127:
        var v30 int64 = 43
        _ = v30
    case -1:
        for fi31 := 0; fi31 < 7; fi31++ {
            switch "\n\t" {
            case "\n\t":
                delete(v2, 420)
            case "a":
                v26 = (v26 / ((-431) | 1))
            case "":
                fmt.Println(v28, v28)
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                switch 129 {
                case 2147483647:
                    fmt.Println(v1, v28)
                case -2147483648:
                    if true {
                        if ((0 != 0) && false) {
                            if ((^(v1)) != v1) {
                                v28 += (v28 * v17)
                            }
                            delete(v2, genFn1("hello", -1))
                        }
                        var v33 int16 = 1000
                        v29 = v29
                        if (!false) {
                            var v34 int = 2147483647
                            v35 := GenS1{F0: (fi32 >> (((v1 ^ v1)) & 63)), F1: v26, F2: 255}
                            _ = v34
                            _ = v35
                        } else {
                            if true {
                                var v36 float64 = (((0.1 / (-2.25 / (float64(3.14159) * (-2.25 / (1.5 / ((1.5 / 1e10) - 3.14159)))))) - -2.25) * 1e10)
                                v37 := GenS1{F0: 9223372036854775807, F1: v26, F2: 37}
                                _ = v36
                                _ = v37
                            }
                            v1 += (v1 ^ 255)
                        }
                        var v38 string = ("hello" + "")
                        _ = v33
                        _ = v38
                    }
                case 1:
                    switch "\n\t" {
                    case "hello":
                        var v39 uint32 = 65535
                        _ = v39
                    case "a":
                        v29 -= (v29 >> (((v1 % ((uint(v17)) | 1))) & 63))
                    }
                default:
                    var v40 uint16 = 42
                    _ = v40
                }
                _ = fi32
            }
            var v41 int32 = -896
            _ = fi31
            _ = v41
        }
    default:
        var v42 uint8 = 2
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
