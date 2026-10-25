package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 uint; F2 bool }
type GenS5 struct { F0 uint8 }
func genFn0() int {
    var v5 uint8 = 208
    if false {
        v5 -= v5
        if true {
            var v6 uint = 561
            var v7 int8 = -116
            _ = v6
            _ = v7
        }
    } else {
        v5 -= (v5 >> ((1) & 63))
    }
    var v8 uint = 0
    _ = v5
    _ = v8
    return -209
}
func genFn1(p0 int, p1 uint8) uint64 {
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{(fi9 & 308), (-(p0))}
        var v11 uint64 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = uint16(p1)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint8 = 255
        v14 = (v14 << ((999) & 63))
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 879, F1: -128, F2: 755}
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return 2
}

func main() {
    v1 := []int{1}
    var v2 float32 = ((float32(-2.25) / 1.5) / 1e10)
    var v3 float64 = -2.25
    var v4 uint8 = 194
    var v16 int16 = 478
    if false {
        fmt.Println(v2)
        v4 += (v4 % (((^(v4))) | 1))
    } else {
        switch ("hello" + ("hello" + ("" + "hello"))) {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            v16 = v16
        }
        v17 := GenS2{F0: false, F1: 2147483647}
        _ = len(v1)
        _ = v17
    }
    fmt.Println(v4, v3, v2)
    if (247 >= v4) {
        v1 = append(v1, 128)
        for fi18 := 0; fi18 < 15; fi18++ {
            v4 = (v4 * 177)
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS1{F0: 0, F1: 42, F2: 65535}
                v4 += (^(v4))
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    } else {
        v2 += 1e10
        switch "\n\t" {
        case "\n\t":
            if (32767 < 1) {
                v16 = 32767
            }
        }
        var v21 float64 = v3
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        if true {
            if true {
                var v23 int8 = -1
                v16 += (v16 / ((-1) | 1))
                _ = v23
            } else {
                v24 := GenS1{F0: 1, F1: -1, F2: 29}
                switch fi22 {
                case 1:
                    fmt.Println(fi22, v4)
                case -9223372036854775808:
                    if (false || (!true)) {
                        var v25 int8 = 100
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    v2 = float32(((v3 / (v3 * (((float64(1e10) / (v3 * v3)) - (float64(-2.25) - 1e10)) - 3.14159))) + v3))
                    if (!((2 == 42) || (("hello" <= "a") && true))) {
                        switch "foo bar" {
                        case "a":
                            v3 = ((v3 - 0.0) + ((v3 + -2.25) * (((0.0 + 1e10) * v3) * float64((v2 / (3.14159 * ((v2 * 1e10) + v2)))))))
                        }
                        v2 -= (-2.25 + 3.14159)
                        var v27 uint8 = 255
                        _ = v27
                    } else {
                        v4 = (v4 & 1)
                        v3 -= v3
                    }
                    var v28 int32 = 2147483647
                    _ = fi26
                    _ = v28
                }
                _ = v24
            }
            if len(v1) > 0 {
                v1[0] = (fi22 | fi22)
            }
        }
        if false {
            var v29 uint = 65535
            _ = v29
        } else {
            for fi30 := 0; fi30 < 5; fi30++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 += v2
                _ = fi30
            }
            _ = len(v1)
        }
        if false {
            _ = len(v1)
            v31 := GenS1{F0: 1, F1: -35, F2: 0}
            switch (fi22 >> ((4294967295) & 63)) {
            case 680:
                for fi32 := 0; fi32 < 17; fi32++ {
                    v16 += v16
                    var v33 int = -2147483648
                    if false {
                        v34 := GenS3{F0: (v4 ^ v4), F1: 42, F2: ("a" + ("hello" + "a"))}
                        v3 = 1.5
                        if (65535 < 32767) {
                            if (float32(((0.1 * float64((v2 * float32((v3 * float64((v2 + v2))))))) + v3)) >= v2) {
                                switch fi22 {
                                case 0:
                                    var v35 uint = 291
                                    _ = v35
                                case 42:
                                    fmt.Println(v2, v3, fi22)
                                case 4294967295:
                                    v36 := GenS4{F0: 931, F1: 120, F2: true}
                                    v1 = append(v1, v33)
                                    _ = v36
                                default:
                                    v1 = append(v1, fi32)
                                }
                                fmt.Println(v2, fi22, v4)
                            }
                            var v37 uint64 = 1
                            _ = v37
                        }
                        _ = v34
                    }
                    _ = fi32
                    _ = v33
                }
            default:
                v16 = v16
            }
            _ = v31
        }
        _ = fi22
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        v2 = v2
        if false {
            var v39 bool = ((-2147483648 == -2147483648) || true)
            _ = v39
        } else {
            v1 = append(v1, fi38)
            v16 = v16
        }
        v16 -= v16
        _ = fi38
    }
    fmt.Println(v16, v4, v3)
    fmt.Println(v4, v16, v4)
    var v40 uint = 255
    var v41 int16 = 1000
    if (0 < 4294967295) {
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := GenS5{F0: v4}
            if len(v1) > 0 {
                v1[0] = (^(fi42))
            }
            if false {
                v1 = append(v1, (fi42 | genFn0()))
                switch (fi42 + fi42) {
                case -136:
                    v2 = 0.0
                }
            } else {
                if true {
                    for fi44 := 0; fi44 < 15; fi44++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v45 uint16 = 2
                        if false {
                            var v46 float64 = 1e10
                            _ = v46
                        }
                        v2 += (v2 * v2)
                        _ = fi44
                        _ = v45
                    }
                } else {
                    _ = len(v1)
                }
                v2 -= ((v2 / 1.0) - 3.14159)
                var v47 int16 = (v16 * v16)
                _ = v47
            }
            _ = fi42
            _ = v43
        }
        v16 += (v16 + v41)
        switch ("hello" + "\n\t") {
        case "hello":
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        default:
            v40 = (v40 - v40)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
