package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 bool }
type GenS4 struct { F0 bool; F1 uint8; F2 uint }
type GenS5 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int8; F1 uint32; F2 bool }
type GenS8 struct { F0 uint16; F1 uint; F2 int8 }
type GenS9 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 uint64 }
func genFn0(p0 int) uint8 {
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 101
    v6 := GenS1{F0: ("" + "hello"), F1: 1, F2: uint32(v4)}
    v7 := GenS1{F0: "hello", F1: 2147483648, F2: 42}
    v5 -= (^(v5))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v5 - v5)
}
func genFn1(p0 int8, p1 int16, p2 int8) uint {
    v8 := GenS2{F0: -1, F1: 65535}
    v9 := GenS3{F0: 2147483647, F1: 0, F2: (!false)}
    v10 := []int{-518}
    var v11 string = "a"
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            if (1 == 18446744073709551615) {
                var v13 int = fi12
                _ = v13
            } else {
                p0 += p2
            }
            if len(v10) > 0 {
                v10[0] = fi12
            }
            if false {
                var v14 uint = 255
                for fi15 := 0; fi15 < 17; fi15++ {
                    v14 -= 65535
                    _ = fi15
                }
                _ = v14
            } else {
                v16 := GenS3{F0: 811, F1: 2, F2: (!(!(543 >= 42)))}
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int32 = -890
                switch "hello" {
                case "foo bar":
                    switch v11 {
                    case "hello":
                        switch v11 {
                        case "foo bar":
                            var v19 string = (("foo bar" + v11) + v11)
                            if (!false) {
                                var v20 int8 = (p2 & 100)
                                _ = v20
                            }
                            v21 := GenS2{F0: v18, F1: 32767}
                            _ = v19
                            _ = v21
                        default:
                            if (false || false) {
                                switch (v11 + ("\n\t" + (v11 + v11))) {
                                case "hello":
                                    v10 = append(v10, 32767)
                                case "a":
                                    fmt.Println(p2, v11, p0)
                                case "\n\t":
                                    var v22 float64 = 0.0
                                    _ = v22
                                default:
                                    var v23 float32 = 0.1
                                    p1 -= 1
                                    switch v11 {
                                    case "hello":
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                    _ = v23
                                }
                                var v24 string = v11
                                _ = v24
                            } else {
                                if len(v10) > 0 {
                                    v10[0] = fi12
                                }
                                fmt.Println(p2, fi12, p0)
                                v10 = append(v10, 128)
                            }
                        }
                    }
                }
                _ = fi17
                _ = v18
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []int{128}
    var v2 int16 = -522
    var v3 float32 = (1.5 - -2.25)
    if (338 == 255) {
        for fi25 := 0; fi25 < 4; fi25++ {
            v1 = append(v1, fi25)
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := GenS3{F0: 2147483648, F1: 2147483648, F2: (583 != (fi26 % ((183) | 1)))}
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: true, F1: 255, F2: 140}
    var v29 bool = (0 < -2147483648)
    fmt.Println(v2)
    switch 42 {
    case 751:
        v3 = v3
    default:
        var v30 uint8 = 2
        _ = v30
    }
    fmt.Println(v2, v3)
    switch -330 {
    case 128:
        if (-1 > 0) {
            v31 := GenS5{F0: 598, F1: 0, F2: 2}
            switch 2147483647 {
            case 32767:
                if len(v1) > 0 {
                    v1[0] = 127
                }
            default:
                v29 = (4294967295 > 4294967295)
            }
            v2 += v2
            _ = v31
        }
    default:
        if true {
            v29 = v29
            fmt.Println(v29)
            _ = len(v1)
        } else {
            v32 := GenS1{F0: "a", F1: 65535, F2: 2}
            if len(v1) > 0 {
                v1[0] = int(v2)
            }
            _ = v32
        }
    }
    var v33 string = (("" + "foo bar") + ("foo bar" + ("hello" + "foo bar")))
    if false {
        if (!(3.14159 < (3.14159 - float64(float32(float64(v3)))))) {
            switch v33 {
            case "foo bar":
                var v34 uint8 = 102
                _ = v34
            case "\n\t":
                v35 := GenS3{F0: 42, F1: 65535, F2: ((0.0 < 1e10) && ((1.5 * 0.0) > 0.1))}
                _ = v35
            case "hello":
                switch 32767 {
                case 0:
                    for fi36 := 0; fi36 < 4; fi36++ {
                        var v37 uint8 = 255
                        switch "foo bar" {
                        case "":
                            for fi38 := 0; fi38 < 17; fi38++ {
                                fmt.Println(fi38, v3)
                                _ = fi38
                            }
                        case "a":
                            if true {
                                var v39 float32 = 3.14159
                                fmt.Println(v39, v2, v2)
                                _ = v39
                            } else {
                                if (90 != 42) {
                                    if true {
                                        v33 = v33
                                        v40 := GenS3{F0: 504, F1: 830, F2: (0.0 == -2.25)}
                                        _ = v40
                                    }
                                    for fi41 := 0; fi41 < 18; fi41++ {
                                        var v42 float64 = 3.14159
                                        _ = fi41
                                        _ = v42
                                    }
                                }
                                if false {
                                    v3 += (1e10 + v3)
                                } else {
                                    v1 = append(v1, (fi36 >> ((65535) & 63)))
                                    for fi43 := 0; fi43 < 17; fi43++ {
                                        if len(v1) > 0 {
                                            v1[0] = (fi36 + (fi43 | (-(fi43))))
                                        }
                                        var v44 uint8 = genFn0((fi36 >> ((0) & 63)))
                                        _ = fi43
                                        _ = v44
                                    }
                                    v45 := GenS5{F0: 32767, F1: 0, F2: 730}
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 14; fi46++ {
                                    v29 = v29
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch fi46 {
                                    case 127:
                                        switch fi46 {
                                        case 226:
                                            v3 += -2.25
                                        case -25:
                                            v47 := GenS6{F0: 309, F1: ("foo bar" + v33)}
                                            _ = v47
                                        case 4294967295:
                                            var v48 uint64 = 42
                                            _ = v48
                                        default:
                                            var v49 int8 = -128
                                            _ = v49
                                        }
                                    }
                                    v33 = v33
                                    _ = fi46
                                }
                            }
                        }
                        if false {
                            v29 = (1 == 2)
                            v50 := GenS1{F0: (("a" + "") + v33), F1: 317, F2: 2147483648}
                            _ = v50
                        }
                        _ = fi36
                        _ = v37
                    }
                }
            }
            switch (((v33 + "\n\t") + ("" + "")) + "") {
            case "foo bar":
                v29 = v29
            case "a":
                v29 = (-671 <= 9223372036854775807)
            default:
                fmt.Println(v29)
            }
            for fi51 := 0; fi51 < 17; fi51++ {
                _ = len(v1)
                if (!(1 > 1)) {
                    v52 := GenS3{F0: 2147483648, F1: 2147483648, F2: v29}
                    _ = len(v1)
                    _ = v52
                } else {
                    fmt.Println(v33, fi51)
                }
                _ = fi51
            }
        }
        v53 := GenS7{F0: 127, F1: 131, F2: v29}
        _ = v53
    } else {
        if false {
            var v54 float64 = ((-2.25 / (1e10 * 3.14159)) * 1.5)
            v2 = v2
            _ = v54
        } else {
            fmt.Println(v29, v29)
            v55 := GenS8{F0: 571, F1: 18446744073709551615, F2: -1}
            if v29 {
                v56 := GenS9{F0: 0, F1: 42, F2: v2}
                _ = v56
            } else {
                v57 := GenS3{F0: 185, F1: 2, F2: v29}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v57
            }
            _ = v55
        }
    }
    v2 = v2
    if ((!true) && v29) {
        fmt.Println(v2, v2, v2)
        v29 = v29
    } else {
        v3 -= ((1.5 - -2.25) / (v3 + v3))
    }
    var v58 bool = (128 <= 72)
    v59 := GenS10{F0: 0, F1: 2147483647}
    fmt.Println(v3, v58, v33)
    if true {
        switch "foo bar" {
        case "\n\t":
            v60 := GenS3{F0: 607, F1: 0, F2: false}
            _ = v60
        case "":
            fmt.Println(v29, v58, v29)
        }
        v61 := GenS11{F0: 1}
        _ = v61
    } else {
        var v62 int8 = 100
        _ = v62
    }
    switch 0 {
    case 322:
        fmt.Println(v3, v2)
    case 748:
        v58 = v29
    case 2147483647:
        if (v58 || (v29 || (1 >= 0))) {
            if true {
                for fi63 := 0; fi63 < 2; fi63++ {
                    fmt.Println(v29)
                    _ = fi63
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v58
    _ = v59
}
