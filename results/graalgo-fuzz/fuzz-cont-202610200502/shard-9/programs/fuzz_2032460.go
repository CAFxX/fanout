package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }
type GenS5 struct { F0 int64; F1 uint64; F2 uint }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 int16, p1 uint16, p2 int8) uint64 {
    p1 += p1
    p0 -= 42
    if true {
        var v7 int = 32767
        _ = v7
    } else {
        var v8 uint64 = uint64(p1)
        _ = v8
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 436
}

func main() {
    var v1 int64 = 1
    var v2 int16 = 32767
    var v3 string = "hello"
    v4 := []string{(((v3 + v3) + (v3 + ("hello" + v3))) + ((((v3 + v3) + v3) + v3) + v3)), "a", (((v3 + v3) + "foo bar") + v3), v3}
    var v5 uint32 = 65535
    v6 := make(map[string]int)
    var v9 bool = (-99 < 956)
    v10 := GenS1{F0: -1, F1: 1, F2: 42}
    v5 -= v5
    for fi11 := 0; fi11 < 19; fi11++ {
        switch -2147483648 {
        case 9223372036854775807:
            v12 := GenS2{F0: (!v9), F1: v5, F2: 1}
            _ = v12
        }
        _ = fi11
    }
    if (-2.25 >= 3.14159) {
        fmt.Println(v9, v3)
        for fi13 := 0; fi13 < 6; fi13++ {
            switch ("\n\t" + ((v3 + ("" + v3)) + v3)) {
            case "\n\t":
                v14 := v6[("\n\t" + v3)]
                _ = v14
            case "hello":
                v15 := v6["hello"]
                _ = v15
            case "foo bar":
                v16 := GenS2{F0: v9, F1: 1, F2: 119}
                _ = v16
            }
            _ = fi13
        }
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS1{F0: 100, F1: uint16(v5), F2: 4294967295}
            fmt.Println(v1)
            if ((((255 != 2) && (!(!false))) || v9) || ((v9 && (v1 != (v1 & (v1 >> ((111) & 63))))) || true)) {
                _ = len(v4)
                v1 = (-(v1))
                var v19 int32 = 810
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    }
    if v9 {
        if (v3 == v3) {
            for fi20 := 0; fi20 < 5; fi20++ {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v1 -= (v1 / (((v1 / (((v1 << ((255) & 63))) | 1))) | 1))
                    if v9 {
                        v3 = ((v3 + v3) + v3)
                        fmt.Println(v2, fi20, v3)
                    }
                    fmt.Println(v9, fi20)
                    _ = fi21
                }
                fmt.Println(v2, v5)
                _ = fi20
            }
        } else {
            fmt.Println(v9, v5, v5)
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
        var v22 int32 = -532
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS3{F0: fi23, F1: v3}
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            if (0.1 == -2.25) {
                switch ("" + (v3 + (v3 + v3))) {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                var v26 uint = 469
                switch "hello" {
                case "foo bar":
                    if true {
                        var v27 uint64 = 865
                        _ = v27
                    } else {
                        var v28 string = ""
                        _ = v28
                    }
                case "\n\t":
                    if v9 {
                        var v29 int16 = 1000
                        var v30 int32 = int32(v1)
                        var v31 bool = (!v9)
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
                _ = v26
            }
            fmt.Println(v3, v1)
            switch 9223372036854775807 {
            case 0:
                if v9 {
                    if len(v4) > 0 {
                        v4[0] = ("hello" + ((v3 + v3) + v3))
                    }
                    var v32 int = (fi25 >> ((1) & 63))
                    if (v9 && (!v9)) {
                        v33 := GenS3{F0: -1, F1: "hello"}
                        _ = v33
                    }
                    _ = v32
                } else {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 int64 = v1
                        _ = fi34
                        _ = v35
                    }
                    var v36 int64 = (-(v1))
                    var v37 int16 = (v2 % ((0) | 1))
                    _ = v36
                    _ = v37
                }
            case 4294967295:
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            default:
                var v38 uint32 = (v5 >> ((492) & 63))
                _ = v38
            }
            _ = fi25
        }
    }
    v9 = (100 != -39)
    var v39 uint8 = 2
    fmt.Println(v9)
    v40 := GenS4{F0: (v39 + v39), F1: 42, F2: 120}
    v41 := GenS1{F0: -1, F1: 949, F2: 4294967295}
    var v42 uint = 84
    var v43 bool = (1000 != int16(v42))
    for fi44 := 0; fi44 < 7; fi44++ {
        if true {
            for fi45 := 0; fi45 < 19; fi45++ {
                delete(v6, ("\n\t" + (v3 + (("" + ("a" + v3)) + ((v3 + (("hello" + v3) + "\n\t")) + ("foo bar" + "\n\t"))))))
                switch ("\n\t" + "") {
                case "hello":
                    v9 = false
                case "a":
                    v1 += v1
                    v2 -= v2
                    switch v3 {
                    case "":
                        var v46 string = ""
                        _ = v46
                    case "\n\t":
                        if ((v5 % ((325) | 1)) >= v5) {
                            var v47 uint8 = v39
                            _ = v47
                        }
                    default:
                        switch fi45 {
                        case 203:
                            switch fi44 {
                            case 1:
                                fmt.Println(v1)
                            }
                        case 0:
                            if ((1000 >= (v2 | 592)) || true) {
                                if v9 {
                                    if v9 {
                                        var v48 bool = v43
                                        if ((-889 > 591) && ((v42 << ((v42) & 63)) < (v42 % ((26) | 1)))) {
                                            fmt.Println(fi45, fi44, fi45)
                                            var v49 int64 = v1
                                            if v43 {
                                                v49 += v49
                                                if true {
                                                    v1 = v1
                                                } else {
                                                    v39 = (v39 * 0)
                                                }
                                                var v50 int = 0
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v43 = (1.5 != 1e10)
                                                var v51 uint16 = 65535
                                                v2 = v2
                                                _ = v50
                                                _ = v51
                                            }
                                            _ = len(v6)
                                            switch (v3 + ((v3 + v3) + v3)) {
                                            case "a":
                                                var v52 int64 = (v49 ^ (v1 >> ((v42) & 63)))
                                                _ = v52
                                            case "foo bar":
                                                fmt.Println(v5, v3)
                                            }
                                            var v53 float32 = 0.1
                                            var v54 int32 = 715
                                            v55 := GenS3{F0: 4294967295, F1: "\n\t"}
                                            _ = v49
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            if (fi45 <= fi45) {
                                                if v48 {
                                                    v5 = v5
                                                } else {
                                                    var v56 int = (^(fi44))
                                                    v2 += 0
                                                    _ = v56
                                                }
                                            } else {
                                                v57 := GenS5{F0: v1, F1: 1, F2: (v42 - uint(v39))}
                                                _ = v57
                                            }
                                            switch 32768 {
                                            case -248:
                                                switch fi45 {
                                                case 9223372036854775807:
                                                    fmt.Println(v2, v2)
                                                default:
                                                    var v58 int16 = v2
                                                    _ = v58
                                                }
                                            case -488:
                                                switch v3 {
                                                case "a":
                                                    v3 = (v3 + (v3 + ("a" + v3)))
                                                }
                                            }
                                        }
                                        _ = v48
                                    } else {
                                        var v59 uint16 = uint16(v42)
                                        _ = len(v4)
                                        fmt.Println(v9, v9, v5)
                                        _ = v59
                                    }
                                    var v60 uint64 = 145
                                    _ = v60
                                } else {
                                    if v9 {
                                        var v61 float32 = 0.0
                                        _ = len(v6)
                                        v62 := GenS5{F0: v1, F1: 0, F2: 65535}
                                        _ = v61
                                        _ = v62
                                    }
                                    var v63 float64 = 1.5
                                    var v64 int64 = -222
                                    _ = v63
                                    _ = v64
                                }
                            } else {
                                var v65 float32 = 0.0
                                fmt.Println(fi45, fi45)
                                _ = v65
                            }
                        default:
                            v66 := GenS5{F0: (v1 % (((^(v1))) | 1)), F1: 4294967295, F2: (v42 & 187)}
                            _ = v66
                        }
                    }
                    v39 = 127
                    v42 += 2
                    if false {
                        var v67 int32 = 1
                        delete(v6, v3)
                        _ = v67
                    }
                    var v68 string = v3
                    _ = v68
                }
                v5 = 531
                v43 = false
                v6["foo bar"] = 4294967295
                _ = fi45
            }
            fmt.Println(fi44)
        } else {
            v1 = (v1 * 9223372036854775807)
            var v69 string = "foo bar"
            _ = v69
        }
        for fi70 := 0; fi70 < 4; fi70++ {
            v1 -= v1
            switch fi44 {
            case -29:
                if true {
                    if len(v4) > 0 {
                        v4[0] = v3
                    }
                    fmt.Println(v2, v42)
                    v9 = v9
                    v71 := GenS6{F0: -128, F1: v42}
                    _ = v71
                }
            }
            _ = fi70
        }
        switch "\n\t" {
        case "\n\t":
            if false {
                v72 := GenS4{F0: v39, F1: -128, F2: v42}
                fmt.Println(v3, v1, v1)
                v73 := GenS7{F0: 2147483647, F1: (v1 + -9223372036854775808), F2: v2}
                _ = v72
                _ = v73
            }
        case "":
            v74 := GenS5{F0: v1, F1: 469, F2: v42}
            _ = v74
        default:
            v3 = v3
        }
        _ = fi44
    }
    var v75 int64 = 622
    var v76 uint16 = 942
    v42 = v42
    var v77 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v9
}
