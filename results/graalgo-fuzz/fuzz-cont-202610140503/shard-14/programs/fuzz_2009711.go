package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0() int {
    var v4 uint32 = 65535
    var v5 int64 = 9223372036854775807
    var v6 int = 32768
    var v7 int = -955
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch 0 {
            case 2147483647:
                switch (fi9 - 32768) {
                case 2147483647:
                    var v10 string = ("" + "")
                    _ = v10
                case 0:
                    v11 := []int{v6, v6, fi9, (fi9 + v6)}
                    _ = v11
                case 1:
                    v12 := []int{fi8}
                    _ = v12
                }
            case -904:
                v13 := GenS1{F0: 174, F1: 1}
                _ = v13
            case 389:
                v14 := GenS2{F0: -1, F1: 77, F2: "hello"}
                _ = v14
            }
            v5 = v5
            var v15 int = 4294967295
            _ = fi9
            _ = v15
        }
        v16 := GenS1{F0: 0, F1: 1}
        _ = fi8
        _ = v16
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint16, p1 uint64) uint64 {
    if ((!((205 < uint32(p1)) && (-2147483648 == 0))) || (!(true || true))) {
        if ("a" >= "hello") {
            v17 := GenS2{F0: -68, F1: 2147483648, F2: ("foo bar" + "")}
            if (true || true) {
                v18 := []int{127, -25, -2147483648}
                if true {
                    _ = len(v18)
                    p1 -= 4294967295
                } else {
                    var v19 uint16 = (p0 + p0)
                    if true {
                        if (127 <= -9223372036854775808) {
                            v20 := GenS3{F0: 42, F1: "\n\t"}
                            var v21 int16 = -1
                            var v22 int32 = 695
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    } else {
                        fmt.Println(p0, v19, p0)
                        for fi23 := 0; fi23 < 13; fi23++ {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                p0 = (^(p0))
                                if ((((-2.25 + (-2.25 + 3.14159)) - (1.5 / ((float64(3.14159) / (((-2.25 * float64(1e10)) / 1e10) / 1.5)) - 0.1))) / 1.5) < 3.14159) {
                                    var v25 float32 = (0.1 + (3.14159 / 1.0))
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: int8(p1), F1: 4294967295, F2: "\n\t"}
                                    v19 -= p0
                                    _ = v26
                                }
                                var v27 uint = 1
                                _ = fi24
                                _ = v27
                            }
                            fmt.Println(p0)
                            _ = fi23
                        }
                    }
                    _ = v19
                }
                var v28 int16 = 830
                _ = v18
                _ = v28
            } else {
                if (!false) {
                    p0 = (p0 - 662)
                }
                var v29 uint32 = 2
                _ = v29
            }
            _ = v17
        }
    } else {
        v30 := GenS4{F0: (p1 % (((p1 | (p1 ^ p1))) | 1)), F1: 740, F2: -1}
        v31 := GenS3{F0: 255, F1: "\n\t"}
        var v32 uint32 = 617
        _ = v30
        _ = v31
        _ = v32
    }
    if true {
        v33 := GenS5{F0: 744, F1: p0}
        var v34 int64 = -9223372036854775808
        _ = v33
        _ = v34
    }
    var v35 float32 = 1e10
    var v36 bool = true
    fmt.Println(v35, p1, v36)
    _ = p0
    _ = p1
    _ = v35
    _ = v36
    return p1
}

func main() {
    v1 := []string{"foo bar", "\n\t", ("\n\t" + ("foo bar" + ("" + "foo bar"))), "hello"}
    var v2 float64 = float64(0.0)
    var v3 int8 = 100
    if true {
        switch "" {
        case "foo bar":
            switch -749 {
            case 379:
                var v37 string = ("\n\t" + ("\n\t" + (("" + "\n\t") + "hello")))
                _ = v37
            }
        default:
            if ((true && (9223372036854775807 == 0)) && (1 != -1)) {
                v38 := GenS1{F0: 2, F1: 1}
                fmt.Println(v2, v3)
                fmt.Println(v2, v2)
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            _ = len(v1)
            for fi40 := 0; fi40 < 11; fi40++ {
                var v41 int8 = 1
                if true {
                    fmt.Println(fi39, v3, fi39)
                    fmt.Println(fi39)
                }
                if true {
                    if true {
                        var v42 int8 = v41
                        _ = v42
                    }
                } else {
                    var v43 int16 = 1000
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
        for fi44 := 0; fi44 < 18; fi44++ {
            var v45 uint64 = 339
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 float32 = 0.1
                _ = fi46
                _ = v47
            }
            _ = fi44
            _ = v45
        }
    } else {
        v48 := GenS3{F0: 104, F1: ""}
        _ = v48
    }
    if false {
        switch ((("" + "a") + ("a" + "a")) + "") {
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = ""
            }
        default:
            switch -1 {
            case 1:
                fmt.Println(v3, v2)
            case 32768:
                v49 := GenS6{F0: "foo bar", F1: 1}
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 9; fi50++ {
            v2 = 1e10
            v3 = -30
            for fi51 := 0; fi51 < 18; fi51++ {
                v52 := GenS4{F0: 2, F1: 724, F2: fi51}
                var v53 uint = 2
                if false {
                    fmt.Println(v2, fi50)
                    var v54 bool = true
                    _ = v54
                } else {
                    v55 := GenS5{F0: 65535, F1: 65535}
                    _ = v55
                }
                _ = fi51
                _ = v52
                _ = v53
            }
            _ = fi50
        }
    }
    var v56 int8 = v3
    v56 = (-(v3))
    v57 := GenS2{F0: (v56 & v56), F1: 1, F2: ("a" + "\n\t")}
    if false {
        if false {
            v56 -= v56
            for fi58 := 0; fi58 < 20; fi58++ {
                var v59 int16 = 32767
                v2 = float64(0.0)
                v60 := GenS7{F0: 733}
                _ = fi58
                _ = v59
                _ = v60
            }
            var v61 uint = 573
            _ = v61
        } else {
            for fi62 := 0; fi62 < 6; fi62++ {
                v56 -= -1
                var v63 float64 = v2
                _ = fi62
                _ = v63
            }
        }
        v64 := GenS8{F0: 185, F1: v3, F2: v56}
        v56 = -128
        _ = v64
    } else {
        if false {
            var v65 int8 = v56
            fmt.Println(v56)
            _ = v65
        }
        if false {
            var v66 int16 = 1000
            _ = v66
        } else {
            fmt.Println(v56, v56)
        }
        v3 = (v3 ^ 26)
    }
    var v67 bool = (float64(0.1) > 3.14159)
    var v68 uint64 = 18446744073709551615
    if (2147483647 <= 338) {
        if (v2 != 0.0) {
            for fi69 := 0; fi69 < 14; fi69++ {
                for fi70 := 0; fi70 < 12; fi70++ {
                    var v71 uint = 21
                    v67 = (uint16(v71) == 2)
                    var v72 int64 = 9223372036854775807
                    if v67 {
                        var v73 string = "hello"
                        v67 = v67
                        if true {
                            v56 = v56
                            switch "\n\t" {
                            case "foo bar":
                                v72 += -113
                            case "\n\t":
                                if false {
                                    var v74 int32 = 178
                                    switch (fi69 >> ((v71) & 63)) {
                                    case -2147483648:
                                        switch genFn0() {
                                        case -985:
                                            if v67 {
                                                var v75 int8 = (v56 >> ((v71) & 63))
                                                v1 = append(v1, v73)
                                                v71 = (^(v71))
                                                switch fi70 {
                                                case 4294967295:
                                                    var v76 int64 = 242
                                                    _ = v76
                                                }
                                                _ = v75
                                            }
                                        default:
                                            fmt.Println(v74, v74)
                                        }
                                    case 0:
                                        v72 += 0
                                    }
                                    var v77 uint = v71
                                    var v78 uint16 = 888
                                    _ = v74
                                    _ = v77
                                    _ = v78
                                }
                            case "a":
                                if len(v1) > 0 {
                                    v1[0] = v73
                                }
                            default:
                                var v79 uint = (v71 - 65535)
                                _ = v79
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                var v80 float32 = -2.25
                                _ = v80
                            case "":
                                switch v73 {
                                case "foo bar":
                                    fmt.Println(v71, v68, v2)
                                case "":
                                    var v81 float64 = (-2.25 - 0.1)
                                    _ = v81
                                }
                            }
                            if (v67 && (v67 || (((v2 + (v2 + 1.5)) / (((-2.25 * v2) + (v2 / v2)) * 0.1)) <= v2))) {
                                v71 = (v71 + v71)
                                var v82 float64 = v2
                                var v83 uint8 = 2
                                _ = v82
                                _ = v83
                            } else {
                                v68 += v68
                                if (v67 || ("a" < v73)) {
                                    v73 = "a"
                                } else {
                                    fmt.Println(v56, fi70)
                                }
                                var v84 uint8 = 1
                                _ = v84
                            }
                        }
                        _ = v73
                    } else {
                        fmt.Println(v56, v68, fi69)
                        v72 = (-(v72))
                        if (v72 > (^(v72))) {
                            v2 = (v2 + (v2 + v2))
                            var v85 int64 = (-(v72))
                            _ = v85
                        }
                    }
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                fmt.Println(v56, v67)
                fmt.Println(v3)
                _ = fi69
            }
            var v86 string = (("foo bar" + "a") + ("" + (("\n\t" + ("hello" + "foo bar")) + "foo bar")))
            _ = v86
        }
        var v87 uint16 = 20
        fmt.Println(v3, v3, v2)
        _ = v87
    }
    var v88 int16 = -950
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v67
    _ = v68
    _ = v88
}
