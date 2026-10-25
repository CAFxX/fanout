package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint8 }
func genFn0() uint {
    var v5 int = 1
    v6 := []int{(v5 & v5), (v5 << ((193) & 63)), v5, (v5 % ((v5) | 1))}
    v5 = v5
    v7 := GenS1{F0: 1, F1: v5, F2: -127}
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 143
    v2 := make(map[int]int)
    v3 := []string{}
    var v4 bool = (42 > -32768)
    v1 = (v1 * v1)
    v1 -= (v1 >> ((7) & 63))
    v1 += (v1 << ((18446744073709551615) & 63))
    var v8 uint = 4294967295
    if ((-84 > 4294967295) && false) {
        var v9 int32 = 42
        _ = v9
    } else {
        v10 := GenS2{F0: 1, F1: 9223372036854775807}
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint32 = 889
            if true {
                fmt.Println(v12)
            } else {
                v13 := GenS1{F0: 65535, F1: (fi11 | -479), F2: 1}
                _ = v13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                v4 = false
                _ = fi14
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v15 bool = v4
    v3 = append(v3, "\n\t")
    var v16 uint64 = 911
    if (v1 > v1) {
        var v17 uint16 = 4
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS3{F0: (v15 && v15)}
            v3 = append(v3, ("a" + ("foo bar" + ("hello" + "foo bar"))))
            if v15 {
                fmt.Println(v15)
                fmt.Println(v15, v17)
            }
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        var v20 uint16 = 65535
        _ = v20
    }
    if v15 {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 string = ((("a" + ("foo bar" + (((("\n\t" + "foo bar") + ("hello" + "")) + ("foo bar" + "\n\t")) + ("hello" + "\n\t")))) + "foo bar") + (("hello" + "hello") + "foo bar"))
            _ = fi21
            _ = v22
        }
        if true {
            switch 977 {
            case 32767:
                if ((v1 * 120) > v1) {
                    v15 = (v8 <= 255)
                    switch 0 {
                    case -898:
                        v3 = append(v3, (("foo bar" + "hello") + "\n\t"))
                    case -4:
                        if (true || true) {
                            _ = len(v2)
                        } else {
                            v23 := GenS4{F0: 2, F1: -588, F2: 0}
                            v24 := GenS4{F0: v16, F1: 372, F2: 2}
                            _ = v23
                            _ = v24
                        }
                    case 32768:
                        v4 = (0 <= 2147483647)
                    }
                    _ = len(v2)
                } else {
                    v8 = (v8 ^ (v8 % (((v8 << (((v8 * (v8 + uint(v1)))) & 63))) | 1)))
                    var v25 bool = (32767 == 42)
                    _ = v25
                }
            case 32768:
                var v26 string = "hello"
                _ = v26
            case -441:
                var v27 uint32 = 2147483648
                _ = v27
            default:
                switch ("hello" + ("foo bar" + "hello")) {
                case "foo bar":
                    if (87 == 697) {
                        v28 := GenS3{F0: (v4 && (9223372036854775807 > -1))}
                        fmt.Println(v15, v8, v1)
                        if v15 {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                var v30 int16 = int16(v16)
                                _ = fi29
                                _ = v30
                            }
                        } else {
                            if v4 {
                                fmt.Println(v4, v15)
                                if (-2147483648 >= 1) {
                                    var v31 int8 = 1
                                    _ = v31
                                } else {
                                    for fi32 := 0; fi32 < 17; fi32++ {
                                        if v15 {
                                            var v33 uint64 = (v16 % ((1) | 1))
                                            var v34 uint = genFn0()
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v1 = v1
                                            v16 += 18446744073709551615
                                        }
                                        if ((-237 >= 42) || (!v15)) {
                                            fmt.Println(v16, v8)
                                            for fi35 := 0; fi35 < 14; fi35++ {
                                                v3 = append(v3, "hello")
                                                _ = fi35
                                            }
                                            for fi36 := 0; fi36 < 3; fi36++ {
                                                fmt.Println(fi36, v4)
                                                var v37 uint32 = 2147483648
                                                _ = fi36
                                                _ = v37
                                            }
                                        }
                                        _ = fi32
                                    }
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        for fi39 := 0; fi39 < 16; fi39++ {
                                            fmt.Println(fi39, v16)
                                            var v40 int16 = -32768
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v16, fi39)
                                                v4 = (v16 < v16)
                                            default:
                                                v1 = v1
                                            }
                                            _ = fi39
                                            _ = v40
                                        }
                                        var v41 uint = (v8 % ((1) | 1))
                                        _ = fi38
                                        _ = v41
                                    }
                                    fmt.Println(v1)
                                }
                            }
                            for fi42 := 0; fi42 < 2; fi42++ {
                                for fi43 := 0; fi43 < 16; fi43++ {
                                    if v4 {
                                        var v44 uint64 = v16
                                        v16 -= 4294967295
                                        _ = v44
                                    }
                                    if true {
                                        _ = len(v3)
                                        var v45 int32 = -473
                                        _ = len(v2)
                                        _ = v45
                                    } else {
                                        v15 = ((1e10 <= ((((float64(3.14159) - 3.14159) * 0.0) / 1e10) * 3.14159)) || (v4 || v15))
                                        var v46 uint = v8
                                        _ = v46
                                    }
                                    v47 := v2[-644]
                                    var v48 int8 = -35
                                    var v49 uint16 = 0
                                    _ = fi43
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                }
                                v15 = ((!v4) || v15)
                                _ = fi42
                            }
                            v16 = 1
                        }
                        _ = v28
                    }
                case "":
                    var v50 int32 = 478
                    _ = v50
                }
            }
            if (!(8 == 0)) {
                if false {
                    var v51 float64 = (((float64(0.1) * 1e10) / 0.1) / 3.14159)
                    _ = v51
                }
                fmt.Println(v1)
                var v52 uint = (v8 ^ genFn0())
                _ = v52
            }
            v53 := GenS5{F0: (^(v1)), F1: (v16 & 771), F2: int16(v1)}
            _ = v53
        }
        if false {
            var v54 int = 128
            for fi55 := 0; fi55 < 13; fi55++ {
                if len(v3) > 0 {
                    v3[0] = ((("\n\t" + ("hello" + "\n\t")) + "") + "hello")
                }
                _ = fi55
            }
            delete(v2, v54)
            _ = v54
        }
    } else {
        if (v1 <= 209) {
            for fi56 := 0; fi56 < 8; fi56++ {
                for fi57 := 0; fi57 < 11; fi57++ {
                    switch fi57 {
                    case 128:
                        v8 -= genFn0()
                    }
                    if ((true || (183 > -1)) || v4) {
                        var v58 uint8 = v1
                        _ = v58
                    } else {
                        fmt.Println(v1)
                        if (fi56 >= -9223372036854775808) {
                            v8 = (v8 >> ((v8) & 63))
                            var v59 uint = 110
                            switch ("\n\t" + ("\n\t" + ("" + "\n\t"))) {
                            case "hello":
                                if v15 {
                                    v15 = (v4 && false)
                                }
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + (("" + ("foo bar" + "hello")) + "\n\t"))
                                }
                            case "foo bar":
                                switch fi57 {
                                case 765:
                                    fmt.Println(fi56, v1)
                                default:
                                    var v60 float32 = 0.0
                                    _ = v60
                                }
                            }
                            _ = v59
                        }
                        var v61 uint32 = 972
                        v3 = append(v3, (("a" + "hello") + (("\n\t" + "hello") + ((("foo bar" + "") + "hello") + "\n\t"))))
                        _ = v61
                    }
                    _ = fi57
                }
                var v62 int16 = 42
                _ = fi56
                _ = v62
            }
            if (2147483647 <= 4294967295) {
                _ = len(v3)
            }
        }
        v63 := GenS1{F0: 283, F1: 32767, F2: -1}
        v8 -= genFn0()
        _ = v63
    }
    v64 := GenS6{F0: 1000, F1: "\n\t"}
    fmt.Println(v4, v15)
    v65 := v2[-158]
    if v4 {
        var v66 int64 = -9223372036854775808
        fmt.Println(v1, v15, v1)
        v3 = append(v3, "hello")
        _ = v66
    } else {
        v67 := GenS3{F0: (-309 > -998)}
        _ = len(v3)
        _ = v67
    }
    v68 := GenS7{F0: uint8(v16)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v68
    _ = v8
}
