package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 string, p1 uint) bool {
    v5 := GenS1{F0: -1}
    fmt.Println(p0, p1)
    p1 += (^(p1))
    var v6 int64 = -298
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (513 > 1)
}
func genFn1(p0 uint8, p1 uint16) uint16 {
    p0 = (p0 + (p0 | p0))
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            switch (("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello") {
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    v8 := GenS1{F0: 1}
                    _ = v8
                }
            case "":
                v9 := GenS1{F0: 65}
                _ = v9
            case "a":
                p0 -= p0
            default:
                var v10 float64 = (1e10 - -2.25)
                _ = v10
            }
            switch "a" {
            case "":
                p1 += p1
            case "\n\t":
                var v11 uint = 4294967295
                _ = v11
            case "a":
                fmt.Println(fi7, fi7)
            default:
                fmt.Println(p0, p0)
            }
            _ = fi7
        }
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p1, p0)
    }
    v12 := []int{0, -1}
    _ = p0
    _ = p1
    _ = v12
    return (p1 >> ((2) & 63))
}

func main() {
    v1 := []uint32{0}
    var v2 int16 = 1000
    v3 := make(map[string]int)
    v4 := []string{"hello"}
    var v13 int64 = 1
    switch "a" {
    case "foo bar":
        var v14 uint64 = 500
        _ = v14
    }
    if (4294967295 == 2147483648) {
        var v15 float64 = ((0.0 / (0.1 * 3.14159)) - 0.1)
        switch 0 {
        case -9223372036854775808:
            var v16 int = -868
            _ = v16
        case 32768:
            var v17 uint32 = 0
            _ = v17
        case 714:
            switch 127 {
            case 406:
                v18 := GenS2{F0: -128, F1: false}
                _ = v18
            case -1:
                var v19 uint64 = 0
                _ = v19
            }
        default:
            v20 := v3["a"]
            _ = v20
        }
        v15 -= v15
        _ = v15
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            if false {
                if ((!true) || (false || (0 <= (-(v2))))) {
                    var v22 int64 = (v13 ^ v13)
                    _ = v22
                } else {
                    if true {
                        v13 = (^(v13))
                    }
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 uint32 = 65535
                        var v25 int8 = int8(v24)
                        var v26 float32 = ((3.14159 + 1e10) - float32(0.0))
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 uint16 = 166
                    _ = v27
                }
                v13 += 1
                for fi28 := 0; fi28 < 7; fi28++ {
                    v29 := GenS3{F0: 498}
                    v2 -= (v2 * (v2 - v2))
                    var v30 uint64 = 2
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            if false {
                if true {
                    v3["hello"] = fi21
                    var v31 int = fi21
                    _ = v31
                } else {
                    fmt.Println(v2, v2, fi21)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                if (float32(1e10) < 3.14159) {
                    var v32 int64 = -754
                    var v33 int32 = -90
                    if (1 <= 255) {
                        switch fi21 {
                        case 489:
                            switch fi21 {
                            case 42:
                                _ = len(v1)
                            case -9223372036854775808:
                                switch 128 {
                                case 4294967295:
                                    _ = len(v3)
                                case 32767:
                                    switch "a" {
                                    case "\n\t":
                                        _ = len(v3)
                                    }
                                }
                            case 4294967295:
                                if false {
                                    for fi34 := 0; fi34 < 13; fi34++ {
                                        v1 = append(v1, 42)
                                        var v35 uint32 = 1
                                        if (!(!false)) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(fi21)
                                        } else {
                                            v36 := GenS2{F0: 1, F1: true}
                                            var v37 int = (fi21 << ((2) & 63))
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := GenS4{F0: 65535}
                                        _ = fi34
                                        _ = v35
                                        _ = v38
                                    }
                                    switch fi21 {
                                    case 32767:
                                        var v39 uint8 = 159
                                        _ = v39
                                    case -2147483648:
                                        if true {
                                            if (18446744073709551615 > 42) {
                                                for fi40 := 0; fi40 < 20; fi40++ {
                                                    fmt.Println(v2, fi21)
                                                    v33 = 494
                                                    _ = fi40
                                                }
                                            }
                                        } else {
                                            v13 -= -249
                                            if true {
                                                var v41 uint16 = genFn1(42, genFn1(1, 666))
                                                fmt.Println(v33, v32, v2)
                                                _ = v41
                                            } else {
                                                v33 = (v33 ^ (v33 << ((210) & 63)))
                                                switch (fi21 / ((fi21) | 1)) {
                                                case 875:
                                                    fmt.Println(v13, v13)
                                                case 4294967295:
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch ("" + "foo bar") {
                                                        case "\n\t":
                                                            for fi42 := 0; fi42 < 1; fi42++ {
                                                                var v43 bool = true
                                                                _ = fi42
                                                                _ = v43
                                                            }
                                                        case "a":
                                                            switch -2147483648 {
                                                            case -1:
                                                                var v44 uint16 = 42
                                                                _ = v44
                                                            case 2147483647:
                                                                for fi45 := 0; fi45 < 1; fi45++ {
                                                                    var v46 float32 = float32(1.5)
                                                                    v13 -= v13
                                                                    v32 += v13
                                                                    _ = fi45
                                                                    _ = v46
                                                                }
                                                            case 661:
                                                                var v47 int = (fi21 ^ fi21)
                                                                _ = v47
                                                            default:
                                                                if (42 < 0) {
                                                                    v32 = v32
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 4294967295
                                                                    }
                                                                    fmt.Println(v32, v2, v32)
                                                                } else {
                                                                    if (-128 <= 0) {
                                                                        if (754 > (v13 | v32)) {
                                                                            switch (fi21 * 128) {
                                                                            case -190:
                                                                                delete(v3, "hello")
                                                                            }
                                                                            v13 = v32
                                                                            var v48 uint64 = 42
                                                                            _ = v48
                                                                        }
                                                                        v49 := GenS5{F0: 772, F1: false}
                                                                        _ = v49
                                                                    } else {
                                                                        v50 := GenS1{F0: -128}
                                                                        if false {
                                                                            v51 := GenS4{F0: 65535}
                                                                            _ = v51
                                                                        } else {
                                                                            v13 += v13
                                                                            if true {
                                                                                for fi52 := 0; fi52 < 13; fi52++ {
                                                                                    var v53 float32 = 1.5
                                                                                    var v54 uint = 255
                                                                                    v54 = 895
                                                                                    _ = fi52
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                }
                                                                                fmt.Println(fi21)
                                                                                if (!(4294967295 >= 65535)) {
                                                                                    v33 = (v33 | v33)
                                                                                }
                                                                            }
                                                                        }
                                                                        var v55 bool = true
                                                                        _ = v50
                                                                        _ = v55
                                                                    }
                                                                    var v56 bool = (false || (true && true))
                                                                    _ = v56
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                var v57 int8 = int8(v2)
                                                                v2 -= v2
                                                                v33 = (v33 ^ (v33 & -12))
                                                                _ = v57
                                                            }
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v2, fi21, fi21)
                                                    default:
                                                        var v58 float32 = (0.0 - 0.0)
                                                        _ = v58
                                                    }
                                                case 642:
                                                    var v59 uint8 = 127
                                                    _ = v59
                                                default:
                                                    v32 = v13
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(v33, v33)
                                    }
                                    delete(v3, "\n\t")
                                }
                            default:
                                v1 = append(v1, 739)
                            }
                        case -584:
                            switch "\n\t" {
                            case "\n\t":
                                var v60 uint8 = 128
                                _ = v60
                            }
                        case 2147483647:
                            v61 := v3["a"]
                            _ = v61
                        default:
                            fmt.Println(v2)
                        }
                    }
                    _ = v32
                    _ = v33
                }
            }
            _ = fi21
        }
    }
    v62 := GenS5{F0: 0, F1: (2 == 414)}
    for fi63 := 0; fi63 < 16; fi63++ {
        v2 = (v2 >> ((5) & 63))
        for fi64 := 0; fi64 < 10; fi64++ {
            if false {
                var v65 uint8 = 255
                if true {
                    v66 := GenS6{F0: (v65 / ((v65) | 1)), F1: -1, F2: 2}
                    _ = v66
                }
                v13 = -1
                _ = v65
            }
            v67 := GenS6{F0: 3, F1: 1, F2: 0}
            _ = fi64
            _ = v67
        }
        _ = fi63
    }
    fmt.Println(v2, v13)
    if false {
        if false {
            if true {
                v13 -= v13
                if false {
                    for fi68 := 0; fi68 < 8; fi68++ {
                        v2 = int16(fi68)
                        _ = fi68
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                } else {
                    var v69 int = 0
                    _ = v69
                }
                for fi70 := 0; fi70 < 13; fi70++ {
                    if false {
                        v2 -= v2
                    } else {
                        if (("a" == "\n\t") && (!(!(!(644 > 255))))) {
                            v2 = int16(v13)
                        }
                    }
                    _ = fi70
                }
            }
        } else {
            fmt.Println(v2)
            fmt.Println(v2, v13)
        }
        for fi71 := 0; fi71 < 9; fi71++ {
            if ((0.0 / ((1.5 / (3.14159 / 1.5)) * 3.14159)) != -2.25) {
                var v72 float64 = (0.1 + 0.0)
                delete(v3, "foo bar")
                _ = v72
            } else {
                v13 -= v13
            }
            v13 = (v13 ^ int64(fi71))
            _ = fi71
        }
        if (false || true) {
            var v73 float32 = -2.25
            v13 = (v13 % (((v13 + -9223372036854775808)) | 1))
            _ = v73
        }
    } else {
        for fi74 := 0; fi74 < 1; fi74++ {
            fmt.Println(fi74, v13)
            fmt.Println(v2, v2, v2)
            _ = fi74
        }
        v75 := GenS4{F0: 518}
        _ = v75
    }
    v76 := GenS4{F0: 0}
    var v77 uint8 = 255
    var v78 uint16 = 65535
    if (487 <= 0) {
        var v79 bool = true
        var v80 uint64 = 0
        _ = v79
        _ = v80
    } else {
        v78 = genFn1((v77 % ((3) | 1)), (v78 & v78))
    }
    v13 = v13
    v2 = v2
    v78 = v78
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v62
    _ = v76
    _ = v77
    _ = v78
}
