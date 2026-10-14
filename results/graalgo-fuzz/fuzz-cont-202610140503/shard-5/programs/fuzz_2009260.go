package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }

func main() {
    var v1 uint64 = 212
    var v2 uint16 = 65535
    v3 := []string{"foo bar", ("hello" + ((("\n\t" + (("a" + "a") + "foo bar")) + "") + ("a" + ("" + ("foo bar" + "foo bar"))))), "\n\t", "foo bar"}
    if true {
        switch 4294967295 {
        case 693:
            _ = len(v3)
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 42:
            switch ("foo bar" + ("a" + (("a" + "a") + (("" + "hello") + (("hello" + ((("\n\t" + "hello") + "a") + ("a" + "hello"))) + ""))))) {
            case "foo bar":
                var v4 uint16 = v2
                _ = v4
            case "hello":
                fmt.Println(v1, v2)
            case "":
                var v5 uint64 = (v1 * v1)
                _ = v5
            }
        }
        var v6 int16 = 736
        v7 := GenS1{F0: int16(v1), F1: 0, F2: 1}
        _ = v6
        _ = v7
    }
    _ = len(v3)
    var v8 uint32 = 2147483647
    v8 += v8
    for fi9 := 0; fi9 < 20; fi9++ {
        if (false && true) {
            fmt.Println(v1)
        }
        v10 := GenS1{F0: 32767, F1: 78, F2: 1}
        var v11 int8 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 9223372036854775807, F1: 221, F2: (^(v1))}
    var v13 int = 4294967295
    if (!true) {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint64 = (v1 ^ v1)
            if (-1 != -691) {
                switch v13 {
                case -2147483648:
                    v3 = append(v3, ("\n\t" + "foo bar"))
                case 127:
                    switch "foo bar" {
                    case "\n\t":
                        v15 = v1
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        if true {
                            v1 = uint64(v2)
                            var v16 int64 = 0
                            _ = v16
                        }
                    }
                case -479:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v1 += (^(v15))
                        if true {
                            if false {
                                v3 = append(v3, "hello")
                                var v18 bool = false
                                _ = v18
                            } else {
                                var v19 float64 = (0.0 * 0.0)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v20 uint32 = v8
                                switch "a" {
                                case "foo bar":
                                    if (uint64(v2) != (v1 % ((v1) | 1))) {
                                        var v21 uint64 = (v15 + (^(v15)))
                                        v1 -= (v21 | uint64(v2))
                                        _ = v21
                                    }
                                case "a":
                                    fmt.Println(fi14, v20)
                                case "":
                                    if true {
                                        var v22 int = (-(fi14))
                                        var v23 uint64 = (v1 + v15)
                                        var v24 string = ("foo bar" + "\n\t")
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        if (!(!false)) {
                                            v3 = append(v3, ("\n\t" + "foo bar"))
                                        } else {
                                            var v25 string = (("\n\t" + "") + "")
                                            v26 := GenS2{F0: -729, F1: 1, F2: (v1 << ((4294967295) & 63))}
                                            _ = v25
                                            _ = v26
                                        }
                                        v8 = 1
                                        fmt.Println(fi14, v8, v15)
                                        var v27 int16 = 32767
                                        _ = len(v3)
                                        _ = v27
                                    }
                                }
                                _ = v19
                                _ = v20
                            }
                            var v28 int = (^(fi14))
                            _ = v28
                        }
                        var v29 string = "foo bar"
                        fmt.Println(v1)
                        _ = fi17
                        _ = v29
                    }
                default:
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 int64 = 42
                        if (-1 < 2147483647) {
                            if (65535 >= 521) {
                                var v32 int16 = 42
                                var v33 uint8 = 42
                                var v34 uint32 = v8
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(fi30, v2)
                            v35 := GenS3{F0: v8, F1: 255}
                            _ = v35
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                for fi36 := 0; fi36 < 13; fi36++ {
                    v1 += (v15 ^ v1)
                    if true {
                        fmt.Println(v13, v8, v2)
                        _ = len(v3)
                    }
                    v37 := GenS3{F0: (v8 - (^(v8))), F1: (v2 / (((^(v2))) | 1))}
                    _ = fi36
                    _ = v37
                }
                if false {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    v38 := GenS3{F0: (v8 ^ v8), F1: v2}
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if (2 == v15) {
                            var v40 int8 = 32
                            _ = v40
                        }
                        _ = fi39
                    }
                    _ = v38
                } else {
                    v41 := GenS3{F0: (v8 << ((255) & 63)), F1: 0}
                    for fi42 := 0; fi42 < 2; fi42++ {
                        v2 = 255
                        _ = fi42
                    }
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v1 = (v15 & (v1 ^ 42))
                        fmt.Println(fi43)
                        v13 += fi43
                        _ = fi43
                    }
                    _ = v41
                }
            }
            if (255 > (v2 << ((18446744073709551615) & 63))) {
                v44 := GenS4{F0: 935, F1: (fi14 - (fi14 - (-(v13))))}
                v45 := GenS5{F0: v15}
                v15 = (^(v15))
                _ = v44
                _ = v45
            }
            _ = fi14
            _ = v15
        }
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            for fi46 := 0; fi46 < 13; fi46++ {
                v2 += v2
                switch 2147483647 {
                case 127:
                    v2 -= v2
                case 9223372036854775807:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                case 42:
                    v47 := GenS4{F0: 1, F1: (v13 + (fi46 ^ (v13 - 128)))}
                    _ = v47
                default:
                    if (!(v2 != v2)) {
                        switch 1 {
                        case 32767:
                            switch (("foo bar" + "hello") + "hello") {
                            case "\n\t":
                                var v48 uint32 = v8
                                _ = v48
                            case "foo bar":
                                var v49 uint16 = (v2 ^ v2)
                                _ = v49
                            case "a":
                                switch (fi46 | v13) {
                                case 32768:
                                    var v50 float32 = 0.0
                                    _ = v50
                                case 128:
                                    for fi51 := 0; fi51 < 6; fi51++ {
                                        var v52 uint = 42
                                        v1 = v1
                                        if false {
                                            v52 -= v52
                                            if (true || ((^(v52)) > v52)) {
                                                v53 := GenS4{F0: (^(v1)), F1: (v13 << (((v52 - uint(v2))) & 63))}
                                                _ = v53
                                            }
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                case -2147483648:
                                    var v54 int32 = -935
                                    _ = v54
                                default:
                                    v3 = append(v3, "")
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case 127:
                            v8 = 2147483648
                        case -591:
                            v3 = append(v3, "a")
                        }
                    } else {
                        fmt.Println(v13)
                        var v55 uint16 = v2
                        v3 = append(v3, "hello")
                        _ = v55
                    }
                }
                v2 -= v2
                _ = fi46
            }
        case "":
            v56 := GenS6{F0: (v13 << ((0) & 63))}
            _ = v56
        }
    }
    var v57 string = "\n\t"
    var v58 int8 = -98
    if false {
        fmt.Println(v8)
        v59 := GenS3{F0: (v8 % ((4294967295) | 1)), F1: 902}
        switch v13 {
        case -1:
            var v60 uint64 = (^(v1))
            _ = v60
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v59
    }
    var v61 uint8 = 42
    _ = len(v3)
    var v62 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v61
    _ = v62
    _ = v8
}
