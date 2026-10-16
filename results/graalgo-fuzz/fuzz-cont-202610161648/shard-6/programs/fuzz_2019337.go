package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int16) int8 {
    var v5 string = ("hello" + "")
    v6 := GenS1{F0: 32767, F1: (p0 | 1), F2: 639}
    _ = p0
    _ = v5
    _ = v6
    return -127
}
func genFn1(p0 int8, p1 uint) uint16 {
    if (42 > 2) {
        switch 0 {
        case -1:
            var v7 uint16 = 427
            _ = v7
        default:
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 = (p0 / (((p0 >> ((74) & 63))) | 1))
                _ = fi8
            }
        }
    } else {
        v9 := GenS2{F0: 179}
        _ = v9
    }
    v10 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{"", "\n\t", "a", ((("\n\t" + "\n\t") + "") + "hello")}
    v3 := []int{}
    var v4 int32 = 2147483647
    switch "a" {
    case "a":
        fmt.Println(v1, v1)
    case "foo bar":
        v1 -= 0.0
    case "":
        _ = len(v2)
    default:
        v11 := GenS3{F0: 945, F1: 2}
        _ = v11
    }
    switch 42 {
    case 0:
        fmt.Println(v1, v4)
    }
    fmt.Println(v4, v1)
    for fi12 := 0; fi12 < 17; fi12++ {
        v1 -= (3.14159 / (0.1 * float64(0.1)))
        v1 -= v1
        var v13 float32 = (((0.0 * 3.14159) * (0.0 * -2.25)) + 0.0)
        _ = fi12
        _ = v13
    }
    var v14 int16 = 32767
    if (genFn1(genFn0(v14), 176) > 877) {
        if true {
            var v15 bool = true
            v14 += 671
            if v15 {
                v1 = (v1 - 0.1)
                var v16 float32 = (-2.25 + ((-2.25 * 3.14159) + (((1e10 - float32((v1 / (v1 - v1)))) / 1.0) * ((1e10 / 3.14159) * -2.25))))
                for fi17 := 0; fi17 < 14; fi17++ {
                    v1 = v1
                    _ = fi17
                }
                _ = v16
            } else {
                var v18 int = 5
                if v15 {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 int16 = v14
                        var v21 string = "\n\t"
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                _ = v18
            }
            _ = v15
        }
        if false {
            fmt.Println(v1)
        } else {
            v22 := GenS3{F0: 565, F1: 691}
            fmt.Println(v14, v4, v1)
            if true {
                for fi23 := 0; fi23 < 5; fi23++ {
                    v1 -= v1
                    if (2 == 4294967295) {
                        v1 = (v1 / v1)
                        v24 := GenS4{F0: 2}
                        _ = v24
                    } else {
                        if (!((-2.25 + float32(v1)) == 1.5)) {
                            v3 = append(v3, (fi23 % ((-1) | 1)))
                        } else {
                            if false {
                                if ((1e10 + -2.25) != (1.5 - 0.0)) {
                                    var v25 int16 = v14
                                    _ = v25
                                }
                            } else {
                                fmt.Println(v1, fi23, v4)
                            }
                            v1 = (v1 / 0.1)
                        }
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v14 = 891
                            if (float32(float64(float32(((float64(3.14159) * 3.14159) - float64((((1.5 * (float32(0.1) * 3.14159)) * float32((((v1 / v1) + 1e10) - (float64(1.5) * ((0.0 + float64(1e10)) + (0.0 - 3.14159)))))) * 0.1)))))) >= 1e10) {
                                v27 := GenS4{F0: 4294967295}
                                _ = v27
                            } else {
                                v1 -= -2.25
                                v2 = append(v2, "")
                            }
                            var v28 int = fi26
                            fmt.Println(v4, v14, v14)
                            var v29 uint8 = 2
                            _ = fi26
                            _ = v28
                            _ = v29
                        }
                        _ = len(v3)
                    }
                    fmt.Println(v14)
                    _ = fi23
                }
                var v30 uint16 = 255
                _ = v30
            } else {
                v31 := GenS4{F0: 42}
                v32 := GenS5{F0: 379, F1: 2, F2: 1}
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = v31
                _ = v32
            }
            _ = v22
        }
    }
    if false {
        if true {
            if len(v3) > 0 {
                v3[0] = 32767
            }
        } else {
            v1 = v1
            var v33 int8 = int8(v14)
            fmt.Println(v14, v33)
            _ = v33
        }
    }
    switch "foo bar" {
    case "hello":
        for fi34 := 0; fi34 < 10; fi34++ {
            if (!(!(1 < 1))) {
                v35 := GenS6{F0: 648, F1: 2147483647}
                if false {
                    var v36 float64 = 1e10
                    switch (fi34 - (fi34 << ((65535) & 63))) {
                    case 1:
                        switch fi34 {
                        case 437:
                            v3 = append(v3, -391)
                        case -920:
                            v14 -= v14
                        }
                    case 0:
                        v1 += 1.5
                    case -399:
                        v1 -= 0.1
                    }
                    v4 += -1
                    _ = v36
                } else {
                    v4 = v4
                }
                _ = v35
            }
            v1 = float64((3.14159 - 3.14159))
            var v37 int64 = 42
            _ = fi34
            _ = v37
        }
    default:
        if len(v3) > 0 {
            v3[0] = -548
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
