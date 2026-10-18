package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 int64; F2 int }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint32) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        if true {
            if (((-2.25 * float64((-2.25 / 1.0))) >= ((-2.25 + 1e10) / 3.14159)) || false) {
                fmt.Println(p1, fi5, p0)
            }
            var v6 uint = 40
            _ = v6
        }
        var v7 int16 = (p0 >> ((958) & 63))
        _ = fi5
        _ = v7
    }
    if false {
        p0 -= p0
        v8 := GenS1{F0: (p1 << ((0) & 63)), F1: 944}
        if true {
            if ((!(-95 >= 42)) && true) {
                v9 := GenS2{F0: 497, F1: (p0 >> ((753) & 63)), F2: ("foo bar" + "a")}
                v10 := []int{9223372036854775807, 32768, 128, -698}
                var v11 int16 = (-(p0))
                _ = v10
                _ = v11
                _ = v9
            } else {
                var v12 int64 = 176
                v13 := GenS2{F0: 32767, F1: (p0 * (-(p0))), F2: ("a" + ("a" + (("\n\t" + "foo bar") + ("" + "foo bar"))))}
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 uint32 = 2
                        p1 += (^(p1))
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 18; fi17++ {
                        fmt.Println(fi17, p1)
                        _ = fi17
                    }
                    _ = fi14
                }
                _ = v12
                _ = v13
            }
            v18 := []int{2147483647, 0}
            _ = v18
        }
        _ = v8
    }
    v19 := []int{325, 42, 822, 1}
    var v20 float64 = float64(((1.5 - (3.14159 * 3.14159)) - 0.0))
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 127
}
func genFn1(p0 string, p1 uint16, p2 uint) uint {
    v21 := []int{}
    var v22 uint = (p2 % (((p2 << ((p2) & 63))) | 1))
    switch p0 {
    case "hello":
        fmt.Println(p2)
    case "\n\t":
        var v23 int16 = 32767
        _ = v23
    case "foo bar":
        v24 := GenS2{F0: 1000, F1: 1, F2: "\n\t"}
        _ = v24
    default:
        for fi25 := 0; fi25 < 9; fi25++ {
            p0 = p0
            _ = fi25
        }
    }
    if ((p2 > v22) && (((^(p1)) == (p1 + 65535)) || false)) {
        v26 := GenS3{F0: true}
        _ = v26
    } else {
        switch (p0 + "hello") {
        case "":
            v22 = 2
        case "foo bar":
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint = (v22 & p2)
                if len(v21) > 0 {
                    v21[0] = (fi27 % ((127) | 1))
                }
                switch fi27 {
                case 4294967295:
                    var v29 int8 = 0
                    _ = v29
                case 774:
                    fmt.Println(p1)
                default:
                    v30 := GenS4{F0: 403, F1: 42}
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        default:
            if false {
                switch int(v22) {
                case 16:
                    p0 = ""
                case -9223372036854775808:
                    if false {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            fmt.Println(p1, p0)
                            for fi32 := 0; fi32 < 19; fi32++ {
                                switch (fi32 - 945) {
                                case 879:
                                    v22 = (v22 << ((v22) & 63))
                                default:
                                    var v33 uint8 = 39
                                    p1 = p1
                                    _ = v33
                                }
                                p0 = (p0 + "foo bar")
                                var v34 uint64 = 42
                                _ = fi32
                                _ = v34
                            }
                            _ = fi31
                        }
                    } else {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            var v36 uint = 650
                            _ = fi35
                            _ = v36
                        }
                    }
                }
                var v37 int = 32767
                if true {
                    v38 := GenS2{F0: 1000, F1: 0, F2: p0}
                    _ = v38
                } else {
                    var v39 uint64 = 4294967295
                    for fi40 := 0; fi40 < 1; fi40++ {
                        v41 := GenS5{F0: 1000, F1: -9223372036854775808, F2: (v37 * fi40)}
                        var v42 uint8 = 2
                        switch p0 {
                        case "\n\t":
                            p2 += (^(v22))
                        default:
                            v37 = (v37 >> ((p2) & 63))
                        }
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v43 int = (v37 * v37)
                    _ = v39
                    _ = v43
                }
                _ = v37
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return v22
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 int32 = 2147483647
    v4 := []int{}
    for fi44 := 0; fi44 < 6; fi44++ {
        if true {
            var v45 uint16 = 119
            for fi46 := 0; fi46 < 7; fi46++ {
                v45 += (^(v45))
                v47 := GenS1{F0: 4, F1: -65}
                v48 := GenS3{F0: false}
                _ = fi46
                _ = v47
                _ = v48
            }
            var v49 float64 = 0.0
            _ = v45
            _ = v49
        } else {
            v3 += v3
        }
        for fi50 := 0; fi50 < 14; fi50++ {
            switch 128 {
            case 4294967295:
                fmt.Println(fi44, v3, v3)
            case 127:
                if (65535 >= 2) {
                    var v51 uint16 = 42
                    var v52 uint8 = 42
                    v51 = v51
                    _ = v51
                    _ = v52
                }
                v2 = append(v2, "a")
                var v53 uint16 = 2
                _ = v53
            }
            v3 += -2147483648
            var v54 uint64 = 1
            var v55 float32 = 1e10
            _ = fi50
            _ = v54
            _ = v55
        }
        _ = fi44
    }
    if false {
        fmt.Println(v3)
        v3 = v3
        var v56 int = -1
        _ = v56
    }
    if ((false || false) || false) {
        v3 = v3
        var v57 uint = genFn1("hello", 32767, genFn1("", 65535, 147))
        fmt.Println(v57)
        _ = v57
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v58 := GenS6{F0: 2}
    for fi59 := 0; fi59 < 2; fi59++ {
        switch (fi59 ^ 0) {
        case 848:
            if true {
                for fi60 := 0; fi60 < 7; fi60++ {
                    var v61 uint = 2
                    var v62 uint64 = 4294967295
                    var v63 int16 = -32768
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
                if false {
                    v64 := GenS2{F0: 1, F1: 130, F2: ((("foo bar" + "") + "a") + "")}
                    _ = v64
                }
                v3 -= v3
            }
        case 9223372036854775807:
            v65 := GenS2{F0: -511, F1: -32768, F2: "\n\t"}
            _ = v65
        case 42:
            v3 = -1
        default:
            var v66 uint16 = 978
            _ = v66
        }
        var v67 bool = (false && (!(65535 < 1)))
        var v68 uint32 = 65535
        _ = fi59
        _ = v67
        _ = v68
    }
    var v69 uint8 = 0
    fmt.Println(v3, v3)
    if true {
        if true {
            if (2 <= 1) {
                fmt.Println(v3, v69)
            } else {
                v3 = (v3 + v3)
                var v70 bool = true
                _ = v70
            }
            switch 2147483647 {
            case 127:
                switch 4294967295 {
                case 9223372036854775807:
                    var v71 uint64 = 18446744073709551615
                    _ = v71
                case 4294967295:
                    v69 = v69
                }
            case 4294967295:
                if ("foo bar" > "a") {
                    v72 := GenS5{F0: 42, F1: -1, F2: 497}
                    var v73 uint16 = 179
                    _ = v72
                    _ = v73
                } else {
                    var v74 float32 = (0.1 * 0.1)
                    _ = v74
                }
            }
        } else {
            var v75 string = (("a" + "hello") + "foo bar")
            v69 = 255
            fmt.Println(v75, v3, v75)
            _ = v75
        }
        if len(v4) > 0 {
            v4[0] = 128
        }
        v69 += v69
    } else {
        var v76 bool = (4294967295 > 1)
        for fi77 := 0; fi77 < 17; fi77++ {
            v69 += (v69 % (((v69 << ((255) & 63))) | 1))
            _ = fi77
        }
        _ = v76
    }
    v78 := v1[((("hello" + ("" + ("hello" + "a"))) + "foo bar") + "hello")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v69:", v69)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v69
    _ = v78
}
