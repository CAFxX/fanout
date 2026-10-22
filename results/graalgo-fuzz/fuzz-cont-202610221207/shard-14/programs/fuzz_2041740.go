package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 int32; F2 int8 }
type GenS8 struct { F0 uint8; F1 string; F2 int32 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint32; F1 int; F2 int16 }
func genFn0(p0 string) uint {
    p0 = p0
    var v4 uint64 = 4294967295
    v5 := []int{-867, -2147483648}
    _ = p0
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 int64, p1 uint) uint64 {
    v6 := []int{32767, -875, -891, 32767}
    if false {
        v7 := GenS1{F0: 0, F1: 847, F2: 42}
        var v8 bool = true
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 float64 = (-2.25 - 3.14159)
    var v3 uint32 = 2147483648
    v9 := GenS1{F0: 42, F1: 0, F2: 42}
    v2 = (v2 - 0.1)
    if true {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: 100, F1: 138, F2: 0}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = (^(v3))
            _ = fi10
            _ = v11
        }
        if false {
            var v12 bool = false
            _ = v12
        }
        if (-1 != -1) {
            v3 += (^(v3))
            fmt.Println(v3, v2)
            var v13 uint8 = 2
            _ = v13
        }
    }
    switch "" {
    case "":
        var v14 int16 = 42
        _ = v14
    default:
        var v15 uint16 = 42
        _ = v15
    }
    v1 = append(v1, (^(v3)))
    var v16 int = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        fmt.Println(v3, v16, v3)
        _ = fi17
    }
    if (false && ((726 > 0) || ((true || (629 != 0)) || true))) {
        for fi18 := 0; fi18 < 8; fi18++ {
            switch (("hello" + ("" + ("hello" + "a"))) + "foo bar") {
            case "\n\t":
                for fi19 := 0; fi19 < 8; fi19++ {
                    switch fi19 {
                    case -2147483648:
                        _ = len(v1)
                    default:
                        v1 = append(v1, v3)
                    }
                    v2 -= ((((v2 * v2) * (v2 / 1.0)) * (v2 * 3.14159)) - (v2 / 1.0))
                    _ = fi19
                }
            case "foo bar":
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch ("" + ("\n\t" + "a")) {
                    case "a":
                        switch (fi18 >> ((0) & 63)) {
                        case 32767:
                            v16 = fi18
                        case 4294967295:
                            fmt.Println(v2, v16)
                        }
                    case "":
                        if (!(false || (42 <= 1))) {
                            v21 := GenS2{F0: 42}
                            _ = v21
                        } else {
                            var v22 uint8 = 42
                            if (42 > 9223372036854775807) {
                                v3 = v3
                                switch "" {
                                case "a":
                                    switch "a" {
                                    case "\n\t":
                                        v2 = v2
                                    default:
                                        var v23 uint32 = 1
                                        _ = v23
                                    }
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            v22 += (^(v22))
                            _ = v22
                        }
                    case "\n\t":
                        var v24 int = (fi18 / ((-847) | 1))
                        _ = v24
                    default:
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                fmt.Println(v2, v2, v16)
                                v25 := GenS3{F0: false}
                                v16 = v16
                                _ = v25
                            } else {
                                v26 := GenS3{F0: false}
                                if true {
                                    switch 32768 {
                                    case 42:
                                        v27 := GenS4{F0: 9223372036854775807, F1: true}
                                        _ = v27
                                    case 460:
                                        var v28 float32 = (1e10 / (0.1 * 1e10))
                                        _ = v28
                                    }
                                    var v29 uint = 388
                                    _ = v29
                                }
                                _ = v26
                            }
                            v30 := GenS5{F0: v3, F1: 42}
                            _ = v30
                        }
                        if true {
                            v31 := GenS5{F0: v3, F1: 1}
                            var v32 int8 = 1
                            fmt.Println(v2, v16)
                            _ = v31
                            _ = v32
                        } else {
                            if (!(18446744073709551615 == genFn0("\n\t"))) {
                                var v33 bool = (1 <= 9223372036854775807)
                                v2 = 0.1
                                v2 = 0.1
                                _ = v33
                            }
                        }
                        switch ("" + "\n\t") {
                        case "foo bar":
                            v3 = 65535
                        default:
                            v16 = 127
                        }
                        var v34 uint16 = 1
                        _ = v34
                    }
                    var v35 uint16 = 231
                    _ = fi20
                    _ = v35
                }
            }
            if (true && (110 >= -9223372036854775808)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                for fi36 := 0; fi36 < 2; fi36++ {
                    var v37 uint8 = 127
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi18
        }
        v38 := GenS2{F0: 42}
        _ = v38
    }
    switch (v16 << ((0) & 63)) {
    case 128:
        v1 = append(v1, v3)
    case 0:
        for fi39 := 0; fi39 < 6; fi39++ {
            for fi40 := 0; fi40 < 6; fi40++ {
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                v2 = 0.0
                var v41 float32 = ((0.1 * float32(0.0)) + (float32(v2) / 1.5))
                _ = fi40
                _ = v41
            }
            if (("hello" + "foo bar") < "foo bar") {
                v2 = (((1.5 / (1.5 - ((v2 - ((float64((3.14159 + 0.0)) + ((v2 - 0.0) / v2)) + v2)) * v2))) - v2) / (v2 - ((v2 / (0.1 * v2)) * v2)))
                fmt.Println(v2)
                v1 = append(v1, (v3 >> ((42) & 63)))
            }
            v42 := GenS6{F0: 32767}
            _ = fi39
            _ = v42
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (("a" + "a") >= "foo bar") {
        v43 := GenS7{F0: ("a" + "foo bar"), F1: -2147483648, F2: 41}
        _ = v43
    } else {
        for fi44 := 0; fi44 < 3; fi44++ {
            switch "foo bar" {
            case "":
                v45 := GenS8{F0: 101, F1: "hello", F2: 1}
                _ = v45
            case "\n\t":
                v3 = v3
            case "hello":
                for fi46 := 0; fi46 < 11; fi46++ {
                    v47 := GenS9{F0: "a", F1: 42}
                    if false {
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        v1 = append(v1, v3)
                        switch (fi46 / (((fi44 % (((v16 * v16)) | 1))) | 1)) {
                        case -9223372036854775808:
                            if false {
                                var v48 int64 = -700
                                _ = v48
                            }
                        case -1:
                            var v49 uint64 = 934
                            _ = v49
                        case 2147483647:
                            v16 -= v16
                        }
                    }
                    v2 = v2
                    v16 += -780
                    _ = fi46
                    _ = v47
                }
            }
            v2 -= (v2 * ((v2 - v2) * v2))
            if (29 <= 42) {
                var v50 uint8 = 189
                v3 = (v3 - v3)
                var v51 int = fi44
                _ = v50
                _ = v51
            } else {
                if true {
                    fmt.Println(v2, fi44, v2)
                    if ((((float64(1.5) / v2) - float64(0.1)) * v2) > (1.5 + (3.14159 / ((v2 - (v2 * (float64((((((0.0 / -2.25) * -2.25) - 0.0) / -2.25) * 1.5)) / ((0.0 / (v2 - v2)) - -2.25)))) * (((v2 / (v2 + v2)) - (v2 - v2)) - ((v2 - (v2 - float64((1.5 + float32(0.1))))) * v2)))))) {
                        fmt.Println(v3, v3)
                        v1 = append(v1, 32)
                        var v52 uint16 = 32767
                        _ = v52
                    } else {
                        if (((("" + "") < "foo bar") && false) && true) {
                            var v53 string = ("a" + "hello")
                            _ = v53
                        } else {
                            v54 := GenS10{F0: 75}
                            var v55 uint32 = v3
                            v3 -= (^(v3))
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS11{F0: 0, F1: (fi44 << ((42) & 63)), F2: -176}
                        for fi57 := 0; fi57 < 13; fi57++ {
                            var v58 uint32 = (v3 >> ((2) & 63))
                            v58 += v3
                            _ = fi57
                            _ = v58
                        }
                        _ = v56
                    }
                } else {
                    _ = len(v1)
                }
                var v59 bool = false
                _ = v59
            }
            _ = fi44
        }
        if false {
            v16 = -537
            switch (-(v16)) {
            case 42:
                v2 = v2
            default:
                _ = len(v1)
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v60 uint64 = 488
    var v61 uint64 = uint64(v3)
    v3 = 65535
    if ((127 > 48) || (v16 > v16)) {
        fmt.Println(v61, v61)
    }
    v1 = append(v1, (v3 / (((v3 ^ v3)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v9
}
