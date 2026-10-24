package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32; F1 string; F2 int }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int32, p1 int8) string {
    p0 -= p0
    if (76 != 2) {
        v7 := []int{42, 127, 9223372036854775807}
        _ = v7
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1() int64 {
    var v8 uint16 = 65535
    v9 := GenS1{F0: 94, F1: 834}
    switch "\n\t" {
    case "a":
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int = (fi10 >> ((2) & 63))
            _ = fi10
            _ = v11
        }
    case "\n\t":
        var v12 int8 = 95
        _ = v12
    }
    v13 := GenS2{F0: 255}
    v8 += (v8 * 32767)
    _ = v13
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 42
    var v3 uint32 = v2
    v4 := []string{"\n\t", "a", "\n\t"}
    var v5 bool = (65535 == 255)
    v6 := []string{}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float64 = 0.1
        v16 := v1[(fi14 ^ fi14)]
        for fi17 := 0; fi17 < 20; fi17++ {
            switch fi14 {
            case -88:
                v2 = v2
            case 674:
                v4 = append(v4, v16)
            case 766:
                v18 := GenS3{F0: 2147483647, F1: 1}
                _ = v18
            default:
                fmt.Println(v16, v15)
            }
            switch fi14 {
            case -907:
                var v19 uint32 = 2147483647
                _ = v19
            case -2147483648:
                fmt.Println(fi14)
            default:
                v15 += v15
            }
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    if v5 {
        switch "" {
        case "a":
            v1[32768] = "hello"
        case "\n\t":
            v20 := GenS4{F0: 42, F1: -2147483648}
            _ = v20
        case "hello":
            if v5 {
                var v21 float32 = (0.1 + (0.1 * 3.14159))
                _ = v21
            } else {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                v3 = v3
            }
        default:
            switch ("" + "hello") {
            case "\n\t":
                v3 = v3
            case "hello":
                var v22 int8 = 127
                _ = v22
            default:
                var v23 uint64 = 4294967295
                _ = v23
            }
        }
    } else {
        if (4294967295 >= 743) {
            v1[32768] = (("hello" + "\n\t") + "foo bar")
            v24 := GenS1{F0: 0, F1: 4294967295}
            v5 = v5
            _ = v24
        } else {
            for fi25 := 0; fi25 < 16; fi25++ {
                for fi26 := 0; fi26 < 2; fi26++ {
                    delete(v1, fi26)
                    _ = fi26
                }
                _ = fi25
            }
            fmt.Println(v5, v5)
        }
        v2 -= 2
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS3{F0: 458, F1: 4294967295}
        v3 -= v3
        _ = fi27
        _ = v28
    }
    var v29 int16 = 1000
    var v30 uint8 = 173
    v29 = (v29 % ((0) | 1))
    fmt.Println(v29)
    if v5 {
        if (1 >= 65535) {
            if (252 <= 1) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 = 0
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                if (v5 && (true && v5)) {
                    var v31 bool = (false && v5)
                    for fi32 := 0; fi32 < 5; fi32++ {
                        v30 -= 0
                        v33 := GenS1{F0: uint64(v3), F1: 0}
                        _ = fi32
                        _ = v33
                    }
                    if v31 {
                        var v34 float32 = 0.0
                        var v35 int = 32768
                        var v36 int16 = (-(v29))
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = v31
                } else {
                    v5 = (!((v2 - (v2 | v3)) == v2))
                    switch 42 {
                    case -852:
                        var v37 int64 = 0
                        _ = v37
                    }
                }
            }
            var v38 float32 = 3.14159
            if v5 {
                v29 = -527
                var v39 int32 = -2147483648
                _ = v39
            }
            _ = v38
        } else {
            v40 := v1[856]
            for fi41 := 0; fi41 < 4; fi41++ {
                v30 += (v30 >> ((255) & 63))
                _ = fi41
            }
            var v42 string = v40
            _ = v40
            _ = v42
        }
        switch 1 {
        case 32767:
            var v43 int32 = 42
            _ = v43
        case 358:
            if v5 {
                switch 32768 {
                case 0:
                    v4 = append(v4, ("foo bar" + ("a" + "hello")))
                case 4294967295:
                    v44 := GenS1{F0: 468, F1: 421}
                    _ = v44
                case 1:
                    v4 = append(v4, "foo bar")
                default:
                    var v45 uint = 65535
                    _ = v45
                }
                if (false && (128 >= 1)) {
                    v46 := GenS4{F0: -128, F1: -337}
                    if false {
                        switch 32767 {
                        case -1:
                            v5 = v5
                        default:
                            for fi47 := 0; fi47 < 5; fi47++ {
                                v48 := GenS3{F0: int32(v30), F1: 147}
                                _ = fi47
                                _ = v48
                            }
                        }
                        var v49 int64 = 42
                        _ = v49
                    }
                    if v5 {
                        v50 := GenS2{F0: v30}
                        _ = v50
                    }
                    _ = v46
                } else {
                    if v5 {
                        var v51 uint64 = 1
                        _ = v51
                    } else {
                        v52 := GenS5{F0: -764}
                        v29 = 1
                        _ = v52
                    }
                    v53 := GenS6{F0: "", F1: (v30 - v30)}
                    _ = v53
                }
            }
        }
        var v54 int16 = 0
        _ = v54
    }
    v3 = v3
    fmt.Println(v5, v29)
    var v55 int64 = genFn1()
    v3 = v3
    v5 = (2 < 338)
    v56 := GenS5{F0: v55}
    if ((!(0 != 829)) || (255 <= 0)) {
        if (128 < -2147483648) {
            for fi57 := 0; fi57 < 16; fi57++ {
                v29 += (v29 & v29)
                v5 = v5
                switch (fi57 | (fi57 << ((445) & 63))) {
                case -1:
                    v58 := GenS6{F0: "a", F1: v30}
                    _ = v58
                default:
                    v4 = append(v4, ((("\n\t" + "hello") + "") + "hello"))
                }
                _ = fi57
            }
            for fi59 := 0; fi59 < 3; fi59++ {
                for fi60 := 0; fi60 < 12; fi60++ {
                    switch (("hello" + (("hello" + "a") + "hello")) + "hello") {
                    case "hello":
                        v61 := GenS7{F0: (v2 >> ((65535) & 63)), F1: ("foo bar" + "foo bar"), F2: fi60}
                        _ = v61
                    case "":
                        switch (fi60 * fi60) {
                        case -2147483648:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case -40:
                            if (4294967295 > 690) {
                                v29 = v29
                                v29 = 0
                                v30 = v30
                            } else {
                                if v5 {
                                    var v62 uint8 = v30
                                    v55 = (^(v55))
                                    _ = v62
                                } else {
                                    if (1 <= 892) {
                                        var v63 int64 = -1
                                        _ = v63
                                    }
                                    var v64 int64 = (v55 % ((-417) | 1))
                                    _ = v64
                                }
                                if (false || v5) {
                                    v55 -= -9223372036854775808
                                    if (((-2.25 <= -2.25) || (!(v5 && v5))) || (int(v30) <= fi60)) {
                                        var v65 float64 = 1.5
                                        if (109 >= 37) {
                                            var v66 float32 = 3.14159
                                            _ = v66
                                        }
                                        v67 := GenS8{F0: v5, F1: (v29 ^ (-(v29)))}
                                        v68 := GenS7{F0: 2147483647, F1: ("a" + "hello"), F2: (fi60 * fi60)}
                                        _ = v65
                                        _ = v67
                                        _ = v68
                                    } else {
                                        var v69 uint16 = 255
                                        v6 = append(v6, "a")
                                        fmt.Println(fi59, v5)
                                        switch ("\n\t" + "") {
                                        case "":
                                            var v70 int16 = -173
                                            _ = v70
                                        }
                                        var v71 string = (("hello" + "hello") + "\n\t")
                                        v72 := GenS9{F0: v3, F1: (v55 << ((0) & 63))}
                                        v73 := GenS5{F0: (v55 & (v55 + genFn1()))}
                                        _ = v69
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    }
                                    fmt.Println(v30, fi60)
                                    v2 = (v2 * (v2 * (v2 << ((1) & 63))))
                                    v74 := GenS9{F0: 42, F1: v55}
                                    _ = v74
                                } else {
                                    if false {
                                        switch ("" + "") {
                                        case "\n\t":
                                            if ((v55 % ((v55) | 1)) == (v55 >> ((255) & 63))) {
                                                fmt.Println(v2)
                                                switch ("\n\t" + ("\n\t" + "hello")) {
                                                case "a":
                                                    var v75 uint16 = 42
                                                    _ = v75
                                                default:
                                                    v29 = v29
                                                }
                                                fmt.Println(fi60, fi60)
                                                delete(v1, -2147483648)
                                            } else {
                                                v3 -= 820
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                        }
                                        fmt.Println(v55)
                                    } else {
                                        _ = len(v6)
                                    }
                                    if ((v5 && ((v55 - 376) <= (v55 & 487))) && (v5 || (((-128 == 60) && (42 <= -128)) && (v5 || true)))) {
                                        _ = len(v6)
                                    } else {
                                        v76 := GenS10{F0: v29, F1: 638}
                                        v29 = (v29 & v29)
                                        v77 := v1[32768]
                                        _ = v76
                                        _ = v77
                                    }
                                }
                            }
                        case 2147483647:
                            v5 = v5
                        }
                    default:
                        v30 += v30
                    }
                    if v5 {
                        v78 := GenS8{F0: v5, F1: int16(v55)}
                        _ = v78
                    }
                    _ = fi60
                }
                _ = fi59
            }
            var v79 int8 = 127
            _ = v79
        } else {
            v80 := GenS3{F0: 42, F1: 42}
            switch "foo bar" {
            case "":
                if (0.0 <= -2.25) {
                    fmt.Println(v29, v3)
                    fmt.Println(v29, v3)
                }
            case "a":
                v29 = 1000
            }
            if (!(v29 < 1)) {
                var v81 string = "\n\t"
                _ = v81
            }
            _ = v80
        }
        v2 += 4294967295
        for fi82 := 0; fi82 < 4; fi82++ {
            var v83 int32 = 1
            if v5 {
                delete(v1, fi82)
                fmt.Println(v2, v2, v30)
                v84 := GenS11{F0: v5, F1: v83, F2: 0}
                _ = v84
            }
            _ = fi82
            _ = v83
        }
    } else {
        var v85 uint16 = 68
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v85
    }
    for fi86 := 0; fi86 < 12; fi86++ {
        var v87 int = fi86
        var v88 int32 = -1
        var v89 int16 = v29
        _ = fi86
        _ = v87
        _ = v88
        _ = v89
    }
    var v90 uint16 = 42
    v29 = int16(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v90
}
