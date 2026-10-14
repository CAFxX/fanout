package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint16) int8 {
    p0 = p0
    var v7 int8 = 127
    if p0 {
        fmt.Println(p1)
        v8 := []int{0, 32767, -1}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{(fi9 / ((0) | 1))}
            for fi11 := 0; fi11 < 10; fi11++ {
                if (!(100 != 1)) {
                    fmt.Println(p1, fi11, v7)
                    if true {
                        fmt.Println(p1, fi9, v7)
                        fmt.Println(p0, fi11)
                        p1 -= 168
                        v12 := GenS1{F0: -1, F1: p0}
                        _ = v12
                    }
                    var v13 string = ("hello" + "a")
                    _ = v13
                } else {
                    var v14 int64 = 0
                    if false {
                        p0 = (-2.25 < ((float64(0.0) / (1.5 / (-2.25 + 1e10))) / 0.1))
                        switch "hello" {
                        case "a":
                            v15 := GenS1{F0: 0, F1: ((92 >= 30) && p0)}
                            _ = v15
                        }
                        fmt.Println(fi11, fi11)
                    } else {
                        fmt.Println(v14, v14, p0)
                        if p0 {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v14 += 747
                            var v16 float32 = 0.1
                            _ = v16
                        } else {
                            if len(v10) > 0 {
                                v10[0] = -1
                            }
                            var v17 float32 = 0.0
                            if true {
                                switch "\n\t" {
                                case "a":
                                    switch "hello" {
                                    case "foo bar":
                                        switch fi11 {
                                        case -9223372036854775808:
                                            v14 -= (v14 >> ((0) & 63))
                                        }
                                    case "\n\t":
                                        _ = len(v10)
                                    case "a":
                                        var v18 uint32 = 2147483647
                                        _ = v18
                                    }
                                default:
                                    fmt.Println(fi9)
                                }
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                var v19 uint16 = 32767
                                _ = v19
                            }
                            _ = v17
                        }
                    }
                    _ = v14
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        v20 := []int{}
        var v21 bool = (9223372036854775807 > 501)
        _ = v20
        _ = v21
    }
    var v22 uint16 = 607
    _ = p0
    _ = p1
    _ = v22
    _ = v7
    return 1
}

func main() {
    v1 := []string{"", "hello"}
    v2 := []uint32{2147483647, 42, 42, 131}
    v3 := make(map[int]bool)
    var v4 bool = (42 > 42)
    var v5 uint = 6
    var v6 int8 = -1
    for fi23 := 0; fi23 < 20; fi23++ {
        if (((!(!true)) && (v6 > v6)) && v4) {
            v1 = append(v1, ("" + ("foo bar" + ("hello" + "foo bar"))))
            v24 := GenS1{F0: 128, F1: (128 >= 127)}
            _ = v24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "a" {
        case "a":
            for fi26 := 0; fi26 < 17; fi26++ {
                v27 := GenS2{F0: 764, F1: 100, F2: v5}
                var v28 uint64 = uint64(v5)
                if true {
                    if false {
                        var v29 uint32 = 2147483648
                        v30 := GenS1{F0: 2147483647, F1: ("\n\t" <= "a")}
                        var v31 int = -9223372036854775808
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                } else {
                    if (!(!v4)) {
                        v3[(fi25 << (((^(v5))) & 63))] = (((fi26 - (fi26 ^ -369)) <= -1) || ((!(v4 && v4)) || (1 > 715)))
                        delete(v3, fi25)
                        var v32 int16 = -830
                        _ = v32
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            fmt.Println(v5)
                            v33 := GenS2{F0: 42, F1: -1, F2: (^(v5))}
                            _ = v33
                        } else {
                            v4 = ((^(v28)) > 0)
                            _ = len(v2)
                            var v34 int = (fi26 ^ (^(fi26)))
                            v35 := GenS1{F0: -94, F1: true}
                            _ = v34
                            _ = v35
                        }
                        switch int(v5) {
                        case 0:
                            var v36 float64 = 1.5
                            _ = v36
                        case 2147483647:
                            v28 = (v28 % (((v28 / ((0) | 1))) | 1))
                        case 1:
                            var v37 float64 = (0.1 - 1e10)
                            _ = v37
                        default:
                            v1 = append(v1, "foo bar")
                        }
                        var v38 float64 = 3.14159
                        v5 -= 4294967295
                        _ = v38
                    }
                    v4 = v4
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        case "foo bar":
            var v39 int64 = -1
            _ = v39
        case "":
            var v40 bool = ("foo bar" >= ("hello" + ("a" + "foo bar")))
            _ = v40
        }
        _ = fi25
    }
    _ = len(v2)
    if v4 {
        v41 := GenS2{F0: 255, F1: v6, F2: v5}
        var v42 uint = (v5 << ((v5) & 63))
        v43 := GenS2{F0: 2, F1: (^(v6)), F2: (v42 | (v5 | 703))}
        _ = v41
        _ = v42
        _ = v43
    }
    if (!v4) {
        fmt.Println(v6)
        for fi44 := 0; fi44 < 9; fi44++ {
            fmt.Println(v6, v6, v4)
            _ = fi44
        }
        for fi45 := 0; fi45 < 11; fi45++ {
            if v4 {
                for fi46 := 0; fi46 < 1; fi46++ {
                    if ((1e10 / 1.5) >= 0.0) {
                        var v47 int = 0
                        _ = v47
                    } else {
                        var v48 int64 = -9223372036854775808
                        if (!(((-128 <= (v6 << ((v5) & 63))) && (!(!((-228 != 925) || v4)))) || (!(!v4)))) {
                            var v49 int8 = (^(v6))
                            _ = v49
                        } else {
                            v48 = -115
                            switch (("\n\t" + "\n\t") + (("" + "foo bar") + "a")) {
                            case "\n\t":
                                v50 := GenS1{F0: 2147483647, F1: ((!v4) && (v4 && false))}
                                _ = v50
                            case "foo bar":
                                var v51 bool = (!(423 == 2))
                                var v52 uint = (^(v5))
                                _ = v51
                                _ = v52
                            case "a":
                                var v53 uint32 = 2147483648
                                _ = v53
                            default:
                                v4 = v4
                            }
                        }
                        var v54 int64 = 0
                        _ = v48
                        _ = v54
                    }
                    _ = fi46
                }
                if true {
                    if ((!v4) || ((((v4 && (2147483647 >= 42)) && (v4 && v4)) && (!(65535 > 2147483647))) || ((true || (-57 != 0)) || false))) {
                        if (32768 >= 0) {
                            var v55 uint8 = 0
                            fmt.Println(v5, fi45, v55)
                            fmt.Println(fi45, v55)
                            _ = v55
                        } else {
                            v6 += v6
                            var v56 uint64 = 4294967295
                            if v4 {
                                for fi57 := 0; fi57 < 18; fi57++ {
                                    var v58 int = int(v5)
                                    var v59 uint16 = 65535
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            } else {
                                v60 := GenS2{F0: 2, F1: genFn0(false, 462), F2: (v5 >> ((771) & 63))}
                                switch (fi45 % (((fi45 * (fi45 + int(v56)))) | 1)) {
                                case 0:
                                    for fi61 := 0; fi61 < 12; fi61++ {
                                        v4 = ((15 > 189) || v4)
                                        _ = fi61
                                    }
                                }
                                _ = v60
                            }
                            _ = v56
                        }
                        var v62 int = -2147483648
                        var v63 int16 = -32768
                        _ = v62
                        _ = v63
                    }
                    v5 = v5
                } else {
                    fmt.Println(v4, v4)
                }
                for fi64 := 0; fi64 < 4; fi64++ {
                    if (("a" == "hello") || (((-32768 > 42) && (0 >= (fi64 % (((fi64 % ((9223372036854775807) | 1))) | 1)))) && true)) {
                        if (!(42 != 4294967295)) {
                            var v65 string = "a"
                            var v66 uint8 = 42
                            _ = v65
                            _ = v66
                        }
                        _ = len(v3)
                    } else {
                        v6 -= (v6 << ((v5) & 63))
                    }
                    _ = fi64
                }
            } else {
                v4 = (-9223372036854775808 >= (fi45 % ((fi45) | 1)))
                for fi67 := 0; fi67 < 6; fi67++ {
                    var v68 int = (fi67 & 4294967295)
                    v6 += v6
                    v68 = fi67
                    _ = fi67
                    _ = v68
                }
                v69 := GenS3{F0: ((2147483647 < 858) || v4), F1: 520}
                _ = v69
            }
            _ = fi45
        }
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v6, v4)
        case "a":
            if (0 == 42) {
                v5 = (v5 - v5)
                switch "hello" {
                case "foo bar":
                    if v4 {
                        fmt.Println(v6, v4)
                    }
                }
            }
        default:
            v6 = (v6 * v6)
        }
        fmt.Println(v6, v6, v4)
    }
    if v4 {
        for fi70 := 0; fi70 < 20; fi70++ {
            var v71 uint32 = 0
            v4 = v4
            var v72 int8 = v6
            _ = fi70
            _ = v71
            _ = v72
        }
        if (-9223372036854775808 != 32767) {
            var v73 string = (("foo bar" + "\n\t") + "")
            v5 += (v5 & v5)
            _ = v73
        } else {
            fmt.Println(v5, v4)
        }
    }
    switch 128 {
    case 9223372036854775807:
        v4 = v4
    case -67:
        v5 -= v5
    case -1:
        delete(v3, 128)
    default:
        for fi74 := 0; fi74 < 10; fi74++ {
            if (0 >= 42) {
                var v75 uint = (v5 / ((v5) | 1))
                v4 = v4
                v76 := GenS4{F0: 65535}
                _ = v75
                _ = v76
            } else {
                switch "hello" {
                case "hello":
                    var v77 float32 = (-2.25 * 1.5)
                    _ = v77
                case "a":
                    fmt.Println(v5)
                default:
                    switch (("hello" + "foo bar") + "\n\t") {
                    case "\n\t":
                        for fi78 := 0; fi78 < 6; fi78++ {
                            if v4 {
                                if v4 {
                                    fmt.Println(v4, v6, v6)
                                }
                                v3[fi74] = true
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                v4 = (true || (!(2147483648 < 429)))
                                v6 = v6
                                v5 = v5
                            }
                            v4 = true
                            fmt.Println(fi74, v4, fi74)
                            _ = fi78
                        }
                    case "":
                        v79 := GenS4{F0: 1}
                        _ = v79
                    default:
                        fmt.Println(v6, v6, v5)
                    }
                }
                fmt.Println(v6, fi74, v4)
            }
            for fi80 := 0; fi80 < 12; fi80++ {
                v81 := GenS5{F0: 42, F1: 285, F2: "a"}
                _ = fi80
                _ = v81
            }
            var v82 string = ""
            _ = fi74
            _ = v82
        }
    }
    switch 128 {
    case 996:
        v83 := GenS6{F0: 2}
        _ = v83
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -605:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        fmt.Println(v6)
    }
    v4 = false
    for fi84 := 0; fi84 < 4; fi84++ {
        fmt.Println(v5)
        for fi85 := 0; fi85 < 17; fi85++ {
            v86 := GenS6{F0: 42}
            _ = fi85
            _ = v86
        }
        _ = fi84
    }
    v87 := GenS3{F0: (4294967295 < -9223372036854775808), F1: 1000}
    delete(v3, 9223372036854775807)
    var v88 int8 = v6
    v4 = true
    v1 = append(v1, ((("hello" + "foo bar") + "foo bar") + "\n\t"))
    var v89 bool = ((255 >= 2) || ((v4 || v4) || ((!(false && v4)) && false)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v87
    _ = v88
    _ = v89
}
