package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 string }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 uint16; F1 uint; F2 int }
type GenS9 struct { F0 uint8 }
func genFn0(p0 int, p1 int32) string {
    if true {
        v6 := GenS1{F0: 42}
        p1 -= (p1 * 194)
        _ = v6
    } else {
        var v7 uint = 1
        if false {
            p1 -= (p1 << ((v7) & 63))
            if (-1 >= 32767) {
                if false {
                    v7 = (v7 & 4294967295)
                    v7 = 2
                }
            } else {
                p0 = (p0 >> ((v7) & 63))
            }
            v8 := []int{p0, p0, 32767, (p0 << (((v7 | 136)) & 63))}
            _ = v8
        }
        v9 := []int{p0, (^(p0))}
        _ = v7
        _ = v9
    }
    p1 += p1
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1(p0 string) int {
    var v10 int32 = 0
    switch p0 {
    case "\n\t":
        if ((!false) || (417 == 32768)) {
            v11 := GenS2{F0: true, F1: (p0 + p0)}
            if ((65535 == 2) && ((^(v10)) < v10)) {
                v12 := []int{788, 32768}
                _ = v12
            } else {
                if (true || (p0 == p0)) {
                    var v13 float64 = 1e10
                    v14 := []int{1}
                    for fi15 := 0; fi15 < 4; fi15++ {
                        _ = len(v14)
                        _ = fi15
                    }
                    _ = v13
                    _ = v14
                } else {
                    v10 -= 42
                }
                p0 = p0
                p0 = (p0 + p0)
            }
            var v16 string = p0
            _ = v11
            _ = v16
        }
    }
    switch -824 {
    case 127:
        var v17 string = p0
        _ = v17
    case 32767:
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := []int{(fi18 ^ (fi18 & fi18)), fi18, 32768, 128}
            _ = fi18
            _ = v19
        }
    case 4294967295:
        if (true || (!false)) {
            if (((v10 < v10) && (false && false)) || (416 != v10)) {
                v20 := []int{2147483647, 128, 9223372036854775807, 4294967295}
                v20 = append(v20, 127)
                _ = v20
            } else {
                var v21 uint32 = 65535
                _ = v21
            }
        } else {
            var v22 uint16 = 1
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS2{F0: true, F1: ("hello" + "a")}
        v25 := GenS2{F0: (!false), F1: "foo bar"}
        _ = fi23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = v10
    return 32767
}

func main() {
    var v1 int16 = 0
    v2 := []int{-1, 687}
    var v3 bool = true
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    if false {
        v1 = v1
        v3 = true
    } else {
        if v3 {
            for fi26 := 0; fi26 < 13; fi26++ {
                switch fi26 {
                case 408:
                    var v27 uint16 = 510
                    _ = v27
                case -9223372036854775808:
                    if (!false) {
                        var v28 uint64 = 42
                        if (!(2 > (^(v28)))) {
                            switch 128 {
                            case 9223372036854775807:
                                v2 = append(v2, (^(fi26)))
                            case 4294967295:
                                if v3 {
                                    var v29 int16 = v1
                                    if true {
                                        v30 := GenS3{F0: 0, F1: fi26}
                                        _ = v30
                                    }
                                    _ = v29
                                }
                            default:
                                if v3 {
                                    var v31 uint32 = 2147483647
                                    var v32 int8 = -1
                                    _ = v31
                                    _ = v32
                                }
                            }
                            var v33 int32 = -1
                            _ = v33
                        } else {
                            if (0 > 100) {
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    if (!true) {
                                        _ = len(v2)
                                        v1 += (v1 + (v1 + v1))
                                        var v35 float64 = (-2.25 - 3.14159)
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                                var v36 float32 = float32(3.14159)
                                _ = v36
                            } else {
                                v28 = v28
                                var v37 uint = 619
                                _ = v37
                            }
                        }
                        _ = v28
                    }
                case 880:
                    fmt.Println(v3)
                }
                if (!v3) {
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                }
                _ = len(v2)
                _ = fi26
            }
            if v3 {
                switch "hello" {
                case "\n\t":
                    v38 := GenS4{F0: 42}
                    _ = v38
                case "a":
                    var v39 int = 2147483647
                    _ = v39
                case "hello":
                    fmt.Println(v3, v1)
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v40 string = "hello"
            _ = v40
        }
        var v41 uint8 = 204
        _ = v41
    }
    _ = len(v4)
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 9; fi43++ {
            if (-2147483648 < 2147483647) {
                fmt.Println(v3, fi42)
            }
            v1 -= (v1 | (v1 | (v1 >> ((4294967295) & 63))))
            var v44 bool = (!false)
            if v3 {
                var v45 int = int(v1)
                v46 := GenS1{F0: 10}
                v4["\n\t"] = (4294967295 < 561)
                v1 += (v1 | (v1 << ((0) & 63)))
                var v47 float64 = (3.14159 + (-2.25 + 1e10))
                _ = v45
                _ = v46
                _ = v47
            }
            var v48 uint16 = 1
            var v49 uint32 = 0
            _ = len(v2)
            var v50 uint8 = 0
            var v51 uint8 = 1
            v51 -= (v51 & v51)
            _ = fi43
            _ = v44
            _ = v48
            _ = v49
            _ = v50
            _ = v51
        }
        fmt.Println(fi42)
        var v52 int = genFn1("")
        _ = fi42
        _ = v52
    }
    v1 -= v1
    v1 = 1
    var v53 string = "foo bar"
    if true {
        switch (v53 + "\n\t") {
        case "\n\t":
            v5[32767] = 156
        }
        var v54 uint8 = 141
        switch v53 {
        case "":
            if (v3 && (("a" + (v53 + (v53 + v53))) >= "foo bar")) {
                switch genFn1("a") {
                case -9223372036854775808:
                    if ((-695 != v1) && v3) {
                        if true {
                            var v55 int32 = 775
                            v55 += 444
                            _ = v55
                        } else {
                            v54 = 2
                        }
                    } else {
                        switch "a" {
                        case "a":
                            v53 = (v53 + ("a" + ("foo bar" + "")))
                        default:
                            for fi56 := 0; fi56 < 5; fi56++ {
                                var v57 float64 = 3.14159
                                v1 -= (-(v1))
                                switch (fi56 ^ 228) {
                                case -901:
                                    v1 = (v1 & (v1 & -1))
                                case -2147483648:
                                    v58 := GenS5{F0: 1, F1: -827, F2: (("\n\t" + v53) + v53)}
                                    _ = v58
                                case 673:
                                    if v3 {
                                        v53 = ((("hello" + "") + "\n\t") + "\n\t")
                                        v3 = (!false)
                                        var v59 string = v53
                                        _ = v59
                                    }
                                }
                                _ = fi56
                                _ = v57
                            }
                        }
                        switch 128 {
                        case 311:
                            v53 = "\n\t"
                        case 1:
                            if v3 {
                                v53 = (((v53 + v53) + (((((v53 + (v53 + v53)) + "a") + v53) + v53) + v53)) + (v53 + v53))
                            } else {
                                v2 = append(v2, -1)
                                var v60 uint32 = 65535
                                _ = len(v2)
                                _ = v60
                            }
                        }
                    }
                }
                for fi61 := 0; fi61 < 1; fi61++ {
                    v3 = (!true)
                    _ = fi61
                }
                if (!v3) {
                    v53 = "\n\t"
                    v1 += (v1 | -640)
                    fmt.Println(v54, v3, v54)
                }
            }
        case "hello":
            fmt.Println(v53, v3, v54)
        case "a":
            v62 := GenS1{F0: 602}
            _ = v62
        default:
            v63 := v4[v53]
            _ = v63
        }
        _ = v54
    }
    if (65535 <= 0) {
        switch genFn1(v53) {
        case 9223372036854775807:
            for fi64 := 0; fi64 < 9; fi64++ {
                var v65 int64 = -1
                _ = fi64
                _ = v65
            }
        case 42:
            var v66 int = 727
            _ = v66
        case -9223372036854775808:
            v53 = (("a" + ((v53 + "foo bar") + v53)) + v53)
        }
    } else {
        if v3 {
            if false {
                v67 := GenS6{F0: 407, F1: int64(v1)}
                for fi68 := 0; fi68 < 6; fi68++ {
                    for fi69 := 0; fi69 < 17; fi69++ {
                        switch ((v53 + v53) + v53) {
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = (fi69 ^ (fi69 << ((4294967295) & 63)))
                            }
                        }
                        var v70 float64 = -2.25
                        var v71 uint16 = 255
                        var v72 uint32 = 150
                        v73 := GenS1{F0: 1}
                        _ = fi69
                        _ = v70
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                    var v74 int64 = 42
                    switch ("" + v53) {
                    case "hello":
                        v53 = (v53 + ("" + v53))
                    case "a":
                        v5[int(v1)] = 65535
                    case "":
                        switch (^(fi68)) {
                        case -9223372036854775808:
                            if (229 < -627) {
                                for fi75 := 0; fi75 < 4; fi75++ {
                                    v76 := v5[-354]
                                    fmt.Println(fi68)
                                    _ = fi75
                                    _ = v76
                                }
                                switch 0 {
                                case 2147483647:
                                    v74 += int64(fi68)
                                default:
                                    var v77 uint64 = 42
                                    _ = v77
                                }
                            } else {
                                v74 = (v74 % ((776) | 1))
                                v74 += (v74 - (v74 << ((626) & 63)))
                                for fi78 := 0; fi78 < 6; fi78++ {
                                    if len(v2) > 0 {
                                        v2[0] = int(v74)
                                    }
                                    var v79 int16 = v1
                                    _ = fi78
                                    _ = v79
                                }
                            }
                        case -21:
                            var v80 uint16 = 42
                            _ = v80
                        case 32768:
                            var v81 int64 = -962
                            _ = v81
                        }
                    default:
                        if ((fi68 / ((fi68) | 1)) >= genFn1(v53)) {
                            switch fi68 {
                            case 128:
                                var v82 uint = 2
                                _ = v82
                            case -1:
                                v83 := GenS7{F0: (v53 + v53), F1: 222}
                                _ = v83
                            case -1000:
                                for fi84 := 0; fi84 < 5; fi84++ {
                                    v85 := GenS8{F0: 65535, F1: 43, F2: 0}
                                    _ = fi84
                                    _ = v85
                                }
                            default:
                                for fi86 := 0; fi86 < 13; fi86++ {
                                    fmt.Println(v53)
                                    _ = fi86
                                }
                            }
                        } else {
                            v1 -= -333
                            var v87 int16 = v1
                            _ = v87
                        }
                    }
                    _ = fi68
                    _ = v74
                }
                _ = v67
            } else {
                if false {
                    var v88 uint32 = 1
                    fmt.Println(v1)
                    v89 := GenS4{F0: 0}
                    _ = v88
                    _ = v89
                }
                var v90 uint32 = 2147483647
                v91 := GenS3{F0: 1, F1: 2147483647}
                _ = v90
                _ = v91
            }
        }
        switch v53 {
        case "":
            for fi92 := 0; fi92 < 9; fi92++ {
                var v93 int = 0
                var v94 int64 = -9223372036854775808
                _ = fi92
                _ = v93
                _ = v94
            }
        case "hello":
            v95 := GenS9{F0: 2}
            _ = v95
        default:
            var v96 string = v53
            _ = v96
        }
    }
    var v97 float32 = 1e10
    var v98 uint8 = 37
    var v99 uint32 = 1
    var v100 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v97
    _ = v98
    _ = v99
}
