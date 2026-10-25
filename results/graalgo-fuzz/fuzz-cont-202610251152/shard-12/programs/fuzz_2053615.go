package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 bool; F1 uint16; F2 int16 }
func genFn0(p0 string, p1 int, p2 int) uint {
    fmt.Println(p2)
    fmt.Println(p1, p2, p0)
    v5 := []int{(p2 ^ 176), -905, (p1 & p2), p1}
    var v6 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 634
}
func genFn1(p0 int32) string {
    v7 := []int{0, 128, 32768}
    if len(v7) > 0 {
        v7[0] = -950
    }
    if (32767 >= 425) {
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                p0 = -1
                _ = fi9
            }
            v10 := GenS1{F0: 4294967295, F1: 314, F2: 32767}
            _ = fi8
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 bool = false
            _ = fi11
            _ = v12
        }
        var v13 uint64 = 42
        _ = v13
    }
    p0 += p0
    _ = p0
    _ = v7
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint16 = 42
    var v2 string = ("" + (("\n\t" + (((("a" + "\n\t") + ("a" + "foo bar")) + (("hello" + ("hello" + "\n\t")) + ((("" + "") + "foo bar") + ""))) + "a")) + ""))
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v14 := GenS1{F0: 42, F1: 4294967295, F2: 42}
    for fi15 := 0; fi15 < 9; fi15++ {
        v2 = v2
        if (!true) {
            var v16 int64 = -1
            _ = v16
        } else {
            var v17 int = fi15
            v3[v2] = false
            v2 = "a"
            _ = v17
        }
        if false {
            if true {
                v2 = ((v2 + v2) + ("foo bar" + v2))
                v1 = v1
                if ((false || (308 <= -822)) || false) {
                    fmt.Println(v1, fi15)
                    v18 := GenS1{F0: 2147483647, F1: 2, F2: fi15}
                    var v19 uint64 = 815
                    _ = v18
                    _ = v19
                } else {
                    v1 = 255
                    var v20 float64 = float64(0.1)
                    _ = v20
                }
            } else {
                var v21 int16 = 0
                v2 = v2
                _ = v21
            }
            v1 -= v1
        } else {
            fmt.Println(v2)
            v1 -= v1
        }
        _ = fi15
    }
    v22 := GenS2{F0: 4294967295}
    if true {
        fmt.Println(v2, v2)
    }
    v23 := GenS1{F0: 798, F1: 18446744073709551615, F2: 4294967295}
    v24 := []int{-2147483648}
    v25 := GenS1{F0: 1, F1: uint(v1), F2: -77}
    switch (v2 + v2) {
    case "":
        if true {
            switch v2 {
            case "foo bar":
                v1 += (v1 << ((155) & 63))
            default:
                fmt.Println(v2, v2)
            }
            switch 128 {
            case -305:
                v1 -= 255
            case 4294967295:
                v1 += 255
            case 127:
                if ((173 > -515) || ((3.14159 / -2.25) == float64(-2.25))) {
                    if false {
                        v1 += v1
                        v24 = append(v24, 2147483647)
                    }
                    var v26 int32 = 0
                    _ = v26
                } else {
                    v27 := v4[128]
                    switch v2 {
                    case "":
                        _ = len(v24)
                    case "a":
                        var v28 int = v27
                        _ = v28
                    case "foo bar":
                        var v29 uint16 = (v1 << ((2) & 63))
                        _ = v29
                    default:
                        v30 := GenS3{F0: -2147483648, F1: 18446744073709551615, F2: 90}
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v32 := GenS1{F0: 2, F1: 96, F2: fi31}
                        switch v27 {
                        case -863:
                            switch (v2 + "") {
                            case "\n\t":
                                v33 := GenS2{F0: 42}
                                _ = v33
                            default:
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    delete(v4, -526)
                                    if (fi34 <= 9223372036854775807) {
                                        var v35 uint = 1
                                        var v36 string = (v2 + v2)
                                        _ = v35
                                        _ = v36
                                    }
                                    if true {
                                        v2 = (v2 + v2)
                                        if false {
                                            v1 = v1
                                            v37 := GenS4{F0: ((((v2 + v2) + v2) + v2) + v2), F1: v2, F2: (-(fi34))}
                                            v27 += 1
                                            _ = v37
                                        } else {
                                            switch fi34 {
                                            case 32768:
                                                v27 -= fi31
                                            case 9223372036854775807:
                                                v27 = (fi34 | fi34)
                                            default:
                                                v38 := GenS1{F0: 815, F1: genFn0("a", (v27 + -1), 530), F2: v27}
                                                var v39 int8 = -24
                                                _ = v38
                                                _ = v39
                                            }
                                            v27 -= 4294967295
                                            if len(v24) > 0 {
                                                v24[0] = (fi34 | (fi34 - (fi34 | fi31)))
                                            }
                                            if (128 >= 255) {
                                                fmt.Println(v2, v2)
                                                v1 -= 255
                                                v1 -= v1
                                            } else {
                                                v24 = append(v24, (fi31 << ((907) & 63)))
                                            }
                                        }
                                        if (v1 >= (v1 / ((v1) | 1))) {
                                            fmt.Println(v27, fi34)
                                            v40 := GenS2{F0: 2}
                                            _ = v40
                                        }
                                        v2 = v2
                                        var v41 string = (("" + (v2 + v2)) + (("" + (((v2 + v2) + ("" + "hello")) + "a")) + v2))
                                        _ = v41
                                    }
                                    _ = fi34
                                }
                            }
                        case 32767:
                            var v42 float64 = (((1e10 - float64(((float32(1.5) - 1e10) - 0.0))) * -2.25) - (0.0 + 1.5))
                            _ = v42
                        default:
                            _ = len(v24)
                        }
                        _ = fi31
                        _ = v32
                    }
                    _ = v27
                }
            }
        }
    default:
        v43 := GenS5{F0: 0}
        _ = v43
    }
    var v44 int = 128
    if (255 <= 128) {
        fmt.Println(v2)
        switch v2 {
        case "a":
            v44 = (v44 + (v44 ^ -384))
        case "hello":
            v2 = "hello"
        default:
            v45 := GenS6{F0: 42, F1: 119}
            _ = v45
        }
        var v46 uint32 = 65535
        _ = v46
    } else {
        v1 = v1
    }
    if true {
        fmt.Println(v1)
        for fi47 := 0; fi47 < 13; fi47++ {
            v2 = v2
            var v48 float32 = 3.14159
            var v49 int64 = 42
            _ = fi47
            _ = v48
            _ = v49
        }
        var v50 float64 = 0.0
        _ = v50
    }
    if ((false && true) || (3.14159 == ((-2.25 + ((1e10 / 1.5) - 1.5)) * 3.14159))) {
        for fi51 := 0; fi51 < 3; fi51++ {
            var v52 uint8 = 41
            v1 = (^(v1))
            v52 = v52
            _ = fi51
            _ = v52
        }
    }
    v1 -= (^(v1))
    v1 += v1
    var v53 int16 = 42
    v1 = (v1 / ((42) | 1))
    switch "hello" {
    case "foo bar":
        switch (v44 % ((v44) | 1)) {
        case 1:
            v54 := GenS7{F0: false, F1: (v1 & v1), F2: (^(v53))}
            _ = v54
        }
    case "a":
        switch -2147483648 {
        case 627:
            for fi55 := 0; fi55 < 3; fi55++ {
                v24 = append(v24, (fi55 | (v44 ^ fi55)))
                _ = fi55
            }
        case 1:
            var v56 uint64 = 42
            _ = v56
        case 127:
            var v57 int64 = 0
            _ = v57
        default:
            v58 := GenS7{F0: (v53 != v53), F1: v1, F2: (v53 % ((v53) | 1))}
            _ = v58
        }
    }
    v59 := GenS1{F0: 489, F1: 255, F2: v44}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v53
    _ = v59
}
