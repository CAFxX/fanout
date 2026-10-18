package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint16; F1 bool; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int32; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint; F1 string; F2 int16 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 int32; F2 uint }
type GenS12 struct { F0 uint64; F1 uint }
type GenS13 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8, p1 int32, p2 uint16) int64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        switch -2147483648 {
        case 1:
            v7 := GenS1{F0: p2}
            _ = v7
        case -9223372036854775808:
            p1 -= 2147483647
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = 127
            p0 += 0
            switch "\n\t" {
            case "a":
                v9 := []int{(fi8 ^ (fi6 - fi6)), (fi8 + fi6), (fi6 & -2147483648)}
                _ = v9
            case "":
                v10 := GenS2{F0: 914}
                _ = v10
            case "hello":
                switch "foo bar" {
                case "hello":
                    p0 += 127
                case "":
                    fmt.Println(p2, fi8, fi6)
                case "a":
                    var v11 int32 = int32(fi6)
                    _ = v11
                }
            default:
                fmt.Println(fi8)
                if (true && (65535 > 812)) {
                    v12 := GenS3{F0: 756, F1: fi8}
                    var v13 int = (fi8 / ((fi8) | 1))
                    _ = v12
                    _ = v13
                }
                if (true && (!false)) {
                    p1 = p1
                    var v14 int16 = 1000
                    v15 := GenS3{F0: (v14 / ((v14) | 1)), F1: (-(fi6))}
                    _ = v14
                    _ = v15
                } else {
                    p1 -= p1
                    p1 = -249
                    v16 := GenS4{F0: (p2 >> ((756) & 63)), F1: (uint8(p2) > 162), F2: ("a" + "")}
                    _ = v16
                }
            }
            _ = fi8
        }
        p1 = (p1 - p1)
        _ = fi6
    }
    var v17 int = 128
    v18 := GenS2{F0: 291}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 1
}
func genFn1() uint32 {
    v19 := GenS2{F0: 4294967295}
    var v20 int16 = 42
    if true {
        v21 := GenS5{F0: v20}
        _ = v21
    }
    v20 = -32768
    v20 = (v20 >> ((172) & 63))
    _ = v19
    _ = v20
    return 65535
}

func main() {
    var v1 int8 = -22
    var v2 int8 = 70
    var v3 uint = 0
    var v4 uint = (v3 % ((65535) | 1))
    v5 := make(map[string]string)
    var v22 int64 = -9223372036854775808
    v4 += v3
    v23 := GenS3{F0: -22, F1: -243}
    v24 := GenS6{F0: 704, F1: 1}
    fmt.Println(v4, v22)
    if (true || (1e10 > (-2.25 + (0.1 - 1.5)))) {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS7{F0: -469, F1: (v22 % ((v22) | 1)), F2: 39}
            fmt.Println(v2)
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS3{F0: -32768, F1: fi27}
                v29 := GenS3{F0: -978, F1: (fi27 >> ((2) & 63))}
                var v30 int8 = (-(v1))
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi25
            _ = v26
        }
        fmt.Println(v4)
    }
    v2 -= (v2 % (((v1 + (-(v1)))) | 1))
    var v31 bool = ((false || false) && true)
    v32 := GenS8{F0: v3, F1: "\n\t", F2: -231}
    var v33 float32 = 0.1
    var v34 float64 = 3.14159
    if (v4 <= (v4 << (((v3 >> (((^(v4))) & 63))) & 63))) {
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 int32 = 42
            _ = fi35
            _ = v36
        }
    } else {
        if v31 {
            fmt.Println(v1, v2, v2)
            var v37 string = "a"
            if (32767 >= 572) {
                fmt.Println(v31)
            } else {
                v34 = 1e10
            }
            _ = v37
        } else {
            _ = len(v5)
            var v38 uint32 = 0
            _ = v38
        }
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            if false {
                if (437 >= genFn1()) {
                    var v39 uint32 = 286
                    v40 := GenS9{F0: 78, F1: (-(v22))}
                    for fi41 := 0; fi41 < 9; fi41++ {
                        if ("\n\t" >= "\n\t") {
                            var v42 uint32 = (^(v39))
                            fmt.Println(v39)
                            v43 := GenS10{F0: -2147483648}
                            _ = v42
                            _ = v43
                        } else {
                            v44 := GenS1{F0: 255}
                            v5[""] = "foo bar"
                            if true {
                                var v45 uint8 = 42
                                for fi46 := 0; fi46 < 8; fi46++ {
                                    var v47 float64 = v34
                                    _ = fi46
                                    _ = v47
                                }
                                _ = v45
                            } else {
                                for fi48 := 0; fi48 < 9; fi48++ {
                                    v3 = v4
                                    v49 := GenS8{F0: v4, F1: "foo bar", F2: 32767}
                                    var v50 string = (("\n\t" + "foo bar") + (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + ("" + ("\n\t" + ("\n\t" + ("" + "hello"))))))
                                    _ = fi48
                                    _ = v49
                                    _ = v50
                                }
                                switch "" {
                                case "hello":
                                    for fi51 := 0; fi51 < 11; fi51++ {
                                        switch fi41 {
                                        case 127:
                                            if ((v2 % (((v2 / ((57) | 1))) | 1)) == v1) {
                                                v52 := GenS5{F0: 1000}
                                                _ = v52
                                            } else {
                                                if (fi41 == (-(fi51))) {
                                                    switch 128 {
                                                    case 128:
                                                        v53 := []int{(fi41 % ((601) | 1))}
                                                        _ = v53
                                                    }
                                                    v31 = (!v31)
                                                }
                                                v54 := GenS11{F0: ("a" + ("a" + "\n\t")), F1: -900, F2: v4}
                                                fmt.Println(v39)
                                                _ = v54
                                            }
                                        }
                                        fmt.Println(v4)
                                        switch "foo bar" {
                                        case "hello":
                                            v31 = (0 >= -2147483648)
                                        case "\n\t":
                                            var v55 uint32 = (^(v39))
                                            _ = v55
                                        case "a":
                                            delete(v5, "\n\t")
                                        }
                                        _ = fi51
                                    }
                                default:
                                    v56 := GenS12{F0: 18446744073709551615, F1: (v4 + (v3 / ((v4) | 1)))}
                                    _ = v56
                                }
                            }
                            _ = v44
                        }
                        var v57 uint8 = 42
                        fmt.Println(v33, v2, v2)
                        _ = fi41
                        _ = v57
                    }
                    _ = v39
                    _ = v40
                } else {
                    var v58 uint64 = 4294967295
                    _ = v58
                }
                v33 += v33
                v59 := GenS13{F0: 42, F1: 308}
                _ = v59
            } else {
                fmt.Println(v2, v22, v4)
                v4 -= v4
                var v60 string = ("hello" + (("foo bar" + ((("" + "a") + ("a" + "a")) + "hello")) + ((("a" + "foo bar") + (("foo bar" + "\n\t") + "a")) + "\n\t")))
                _ = v60
            }
        default:
            for fi61 := 0; fi61 < 6; fi61++ {
                for fi62 := 0; fi62 < 15; fi62++ {
                    v3 += v3
                    _ = fi62
                }
                v63 := GenS2{F0: 2147483648}
                _ = fi61
                _ = v63
            }
        }
        v64 := GenS8{F0: 651, F1: "", F2: 568}
        _ = v64
    }
    var v65 int16 = 42
    delete(v5, "")
    v1 -= v1
    switch 32768 {
    case 42:
        var v66 string = "hello"
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v65
}
