package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int32; F1 uint32; F2 int }
type GenS6 struct { F0 uint8; F1 string; F2 int16 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 uint32 = 72
    v4 := make(map[string]bool)
    switch 4294967295 {
    case 4294967295:
        var v5 bool = false
        _ = v5
    case 992:
        fmt.Println(v2)
    case -9223372036854775808:
        v6 := GenS1{F0: -161, F1: 9223372036854775807, F2: (!false)}
        _ = v6
    }
    v3 = 181
    var v7 bool = true
    if false {
        fmt.Println(v7, v2)
    } else {
        delete(v4, "")
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v1[v3] = (true || (v7 || (float32(v2) >= float32(v2))))
        for fi9 := 0; fi9 < 15; fi9++ {
            if (!(!(fi8 >= fi8))) {
                v10 := []int{fi9, -2147483648}
                _ = v10
            } else {
                v11 := v4[""]
                fmt.Println(v11)
                v3 = v3
                v2 = ((v2 - v2) / v2)
                _ = v11
            }
            v12 := GenS2{F0: (-9223372036854775808 != fi9), F1: 128}
            if true {
                _ = len(v4)
            } else {
                if v7 {
                    if ((!true) || ((!v7) && (v2 <= v2))) {
                        var v13 uint16 = 32767
                        v14 := GenS1{F0: -649, F1: -488, F2: v7}
                        _ = v13
                        _ = v14
                    } else {
                        switch (^(fi9)) {
                        case 127:
                            v15 := GenS1{F0: -2147483648, F1: 1, F2: v7}
                            _ = v15
                        case 1:
                            switch "a" {
                            case "":
                                v16 := GenS1{F0: -238, F1: -1, F2: (255 == 29)}
                                _ = v16
                            }
                        default:
                            if (fi9 < (fi8 * fi8)) {
                                v3 = v3
                                var v17 uint16 = 0
                                var v18 bool = v7
                                _ = v17
                                _ = v18
                            }
                        }
                        if false {
                            if v7 {
                                fmt.Println(fi8, v2, v7)
                                v1[v3] = v7
                                switch "\n\t" {
                                case "":
                                    if v7 {
                                        var v19 uint64 = 4294967295
                                        if v7 {
                                            if (((false || v7) || true) || v7) {
                                                fmt.Println(fi8)
                                                if true {
                                                    if false {
                                                        v20 := GenS2{F0: ((!(false || ((v19 * (^(v19))) < v19))) || v7), F1: 138}
                                                        _ = v20
                                                    }
                                                    if true {
                                                        v2 += 0.0
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                }
                                                v21 := v4["\n\t"]
                                                _ = v21
                                            } else {
                                                if false {
                                                    var v22 int32 = -858
                                                    _ = v22
                                                }
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            switch ("foo bar" + (("hello" + ("hello" + "hello")) + "hello")) {
                                            case "":
                                                v2 -= v2
                                            case "a":
                                                var v23 bool = ((1.5 / (-2.25 / 1.0)) > (1.5 - 1.5))
                                                _ = v23
                                            }
                                            v24 := GenS2{F0: ((1 >= -1) || false), F1: uint8(v3)}
                                            v25 := []int{}
                                            _ = v24
                                            _ = v25
                                        }
                                        if true {
                                            if v7 {
                                                fmt.Println(v19)
                                                v26 := GenS3{F0: -9223372036854775808, F1: 623}
                                                _ = len(v4)
                                                if (191 > 2147483647) {
                                                    var v27 int = fi9
                                                    if (!(65535 <= 4294967295)) {
                                                        v28 := []int{}
                                                        var v29 string = "\n\t"
                                                        var v30 uint32 = v3
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    _ = v27
                                                } else {
                                                    delete(v1, 24)
                                                    v31 := GenS2{F0: true, F1: 255}
                                                    _ = v31
                                                }
                                                var v32 uint16 = 2
                                                _ = v26
                                                _ = v32
                                            } else {
                                                var v33 int = (fi9 >> ((uint(v3)) & 63))
                                                v34 := GenS4{F0: ("hello" + ""), F1: "hello"}
                                                fmt.Println(v19, fi9)
                                                _ = v33
                                                _ = v34
                                            }
                                            v35 := GenS5{F0: int32(fi9), F1: (v3 << ((670) & 63)), F2: fi9}
                                            var v36 string = "\n\t"
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v19
                                    } else {
                                        fmt.Println(v3, v3, fi8)
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v3)
                                        case "hello":
                                            if v7 {
                                                var v37 int = fi8
                                                v38 := GenS4{F0: "foo bar", F1: ""}
                                                var v39 uint = 267
                                                v7 = v7
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v3 = v3
                                                v40 := []int{fi8, (^(fi9))}
                                                _ = v40
                                            }
                                        case "foo bar":
                                            fmt.Println(fi9, v7)
                                        }
                                        v7 = v7
                                    }
                                case "foo bar":
                                    var v41 int64 = 0
                                    _ = v41
                                default:
                                    v42 := []int{(fi9 - -9223372036854775808)}
                                    _ = v42
                                }
                            }
                            if (-1 == -128) {
                                var v43 float64 = (v2 * ((float64(0.0) - v2) - v2))
                                if ((^(v3)) > (v3 * (v3 >> ((65535) & 63)))) {
                                    var v44 string = "\n\t"
                                    v2 = 0.1
                                    var v45 uint8 = 36
                                    if true {
                                        if (-9223372036854775808 <= 1) {
                                            v2 = (v2 + ((v2 / 0.1) * (-2.25 - v2)))
                                            var v46 int8 = 100
                                            _ = v46
                                        } else {
                                            v47 := GenS6{F0: 0, F1: "", F2: int16(v45)}
                                            if (v7 && (-1 != 2147483647)) {
                                                v48 := []int{42, fi9, fi9, (fi9 / ((fi9) | 1))}
                                                if v7 {
                                                    _ = len(v4)
                                                    if v7 {
                                                        v3 = v3
                                                        var v49 uint8 = 19
                                                        _ = v49
                                                    }
                                                } else {
                                                    v43 -= 3.14159
                                                    if len(v48) > 0 {
                                                        v48[0] = fi9
                                                    }
                                                    if (255 != 255) {
                                                        if len(v48) >= 2 {
                                                            v48 = v48[1:len(v48)]
                                                        }
                                                        v50 := GenS7{F0: 65535}
                                                        var v51 int = fi9
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        fmt.Println(v44, v3, v3)
                                                        if (!(65535 == 1)) {
                                                            v44 = v44
                                                            v4[v44] = false
                                                        }
                                                        if (v7 && v7) {
                                                            fmt.Println(v3)
                                                        } else {
                                                            fmt.Println(v7, fi8)
                                                        }
                                                    }
                                                    if false {
                                                        if v7 {
                                                            v52 := v4[(v44 + v44)]
                                                            _ = v52
                                                        }
                                                        var v53 int64 = -9223372036854775808
                                                        _ = v53
                                                    }
                                                }
                                                _ = v48
                                            }
                                            var v54 float64 = (1e10 / v43)
                                            _ = v47
                                            _ = v54
                                        }
                                        v43 -= v43
                                    }
                                    _ = v44
                                    _ = v45
                                } else {
                                    var v55 int8 = 0
                                    _ = v55
                                }
                                if (3.14159 != 1e10) {
                                    v43 += (-2.25 + v43)
                                }
                                v56 := v1[904]
                                _ = v43
                                _ = v56
                            } else {
                                v7 = (!(18446744073709551615 >= 65535))
                            }
                        }
                        v2 += (v2 - (1.5 * (1.5 * (0.0 - v2))))
                        v2 = 0.1
                    }
                } else {
                    v57 := GenS6{F0: 127, F1: ("\n\t" + "\n\t"), F2: 1}
                    v58 := v4[("hello" + ("" + ((("a" + "\n\t") + ((("hello" + "") + "a") + "a")) + "hello")))]
                    if (784 <= 974) {
                        fmt.Println(fi9, fi9)
                    } else {
                        v58 = v58
                    }
                    _ = v57
                    _ = v58
                }
                v3 -= 4294967295
            }
            _ = fi9
            _ = v12
        }
        _ = fi8
    }
    if (834 < 421) {
        var v59 bool = (v7 && (0.0 > 0.1))
        _ = v59
    }
    var v60 bool = (127 > 255)
    var v61 uint64 = 744
    for fi62 := 0; fi62 < 1; fi62++ {
        var v63 int64 = int64(v61)
        var v64 uint64 = 42
        _ = fi62
        _ = v63
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v7
}
