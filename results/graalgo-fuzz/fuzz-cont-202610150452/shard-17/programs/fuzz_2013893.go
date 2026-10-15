package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint; F1 string; F2 int64 }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int64; F1 int; F2 uint }
type GenS10 struct { F0 int; F1 int16 }
func genFn0(p0 int64, p1 string) int {
    switch 122 {
    case 127:
        v4 := GenS1{F0: 42, F1: 255, F2: p1}
        _ = v4
    case -17:
        var v5 int16 = 414
        _ = v5
    default:
        switch p1 {
        case "foo bar":
            switch ("" + "\n\t") {
            case "":
                var v6 int64 = p0
                _ = v6
            case "a":
                if (42 != 784) {
                    var v7 uint8 = 0
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 int8 = -1
                        p1 = p1
                        v10 := []int{(fi8 << ((0) & 63)), fi8, (fi8 + fi8)}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            case "foo bar":
                for fi11 := 0; fi11 < 14; fi11++ {
                    fmt.Println(p1)
                    _ = fi11
                }
            default:
                switch 42 {
                case 2147483647:
                    p1 = "foo bar"
                default:
                    if false {
                        p0 = (p0 | p0)
                        v12 := []int{4294967295}
                        p1 = "foo bar"
                        _ = v12
                    } else {
                        var v13 int64 = 42
                        _ = v13
                    }
                }
            }
        }
    }
    v14 := GenS1{F0: 281, F1: 690, F2: "foo bar"}
    if (-1 == 127) {
        if false {
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 4; fi16++ {
                    if false {
                        v17 := GenS1{F0: -1, F1: 1, F2: "hello"}
                        v18 := GenS1{F0: 32767, F1: 1, F2: (p1 + ((p1 + p1) + (p1 + "a")))}
                        var v19 float32 = ((((float32(3.14159) + 1e10) + 1.5) * 1.5) - (0.1 - float32(1e10)))
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        if true {
                            v20 := GenS1{F0: 0, F1: 65535, F2: (((((p1 + ((p1 + "a") + p1)) + ("hello" + "foo bar")) + p1) + p1) + "foo bar")}
                            switch 4294967295 {
                            case -525:
                                p0 = p0
                            }
                            v21 := GenS2{F0: (fi15 << ((2) & 63)), F1: 255, F2: 77}
                            _ = v20
                            _ = v21
                        }
                        var v22 int = fi15
                        _ = v22
                    }
                    if (false || false) {
                        fmt.Println(p1)
                        var v23 int = fi16
                        if false {
                            var v24 uint16 = 0
                            v25 := GenS3{F0: 2}
                            _ = v24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        fmt.Println(fi15)
                    }
                    fmt.Println(p0, p0)
                    _ = fi16
                }
                if true {
                    var v26 string = "hello"
                    v27 := GenS4{F0: 2}
                    var v28 uint64 = 134
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi15
            }
            for fi29 := 0; fi29 < 4; fi29++ {
                p1 = "hello"
                p1 = p1
                v30 := GenS5{F0: 1}
                _ = fi29
                _ = v30
            }
        }
    }
    _ = p0
    _ = p1
    _ = v14
    return 32767
}
func genFn1() int64 {
    var v31 uint16 = 32767
    var v32 float32 = 3.14159
    v33 := GenS5{F0: 676}
    v34 := GenS5{F0: 512}
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    return 1
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    switch "foo bar" {
    case "foo bar":
        v35 := GenS6{F0: 1, F1: 42}
        _ = v35
    case "\n\t":
        v1 += 65535
    default:
        var v36 bool = (-187 < -1)
        _ = v36
    }
    if (false && false) {
        _ = len(v2)
        v37 := []int{genFn0(genFn1(), "\n\t"), 174}
        var v38 uint = (v1 ^ v1)
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        switch (fi39 | (fi39 + fi39)) {
        case 0:
            switch "foo bar" {
            case "foo bar":
                var v40 int8 = 127
                _ = v40
            }
        case 729:
            v41 := GenS2{F0: fi39, F1: 42, F2: 0}
            _ = v41
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            if (1000 == 377) {
                delete(v3, 65535)
                var v43 bool = (0.1 > 0.1)
                v44 := GenS7{F0: (v1 | v1), F1: "a", F2: -1}
                v1 -= (^(v1))
                v45 := []int{}
                fmt.Println(v1, v1, v1)
                _ = v43
                _ = v44
                _ = v45
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v46 bool = true
                    var v47 uint64 = 1
                    v1 = 255
                    _ = v46
                    _ = v47
                case "":
                    v48 := []int{fi39, fi39, 4294967295, (fi39 & genFn0(367, "a"))}
                    _ = v48
                default:
                    fmt.Println(v1)
                }
            }
            if (false || true) {
                var v49 float32 = 0.0
                v50 := GenS3{F0: 2}
                switch ("" + "") {
                case "a":
                    v1 -= (v1 << ((v1) & 63))
                default:
                    v49 = (v49 + v49)
                }
                _ = v49
                _ = v50
            }
            _ = fi42
        }
        fmt.Println(v1, fi39)
        _ = fi39
    }
    switch ("" + ("\n\t" + "hello")) {
    case "hello":
        if (-2.25 == float32(float64(1e10))) {
            if false {
                v51 := GenS5{F0: v1}
                var v52 int64 = -9223372036854775808
                _ = v51
                _ = v52
            }
            var v53 uint8 = 255
            _ = v53
        }
    case "":
        fmt.Println(v1)
    }
    var v54 string = "\n\t"
    v54 = "foo bar"
    _ = len(v3)
    if true {
        if (!(!false)) {
            v55 := []int{genFn0(678, (v54 + v54)), 4294967295, 958, genFn0(0, v54)}
            fmt.Println(v1, v1)
            _ = v55
        }
        v56 := v3[639]
        if (128 >= 127) {
            v56 -= -700
            for fi57 := 0; fi57 < 15; fi57++ {
                var v58 int32 = 0
                _ = fi57
                _ = v58
            }
            var v59 int = -746
            _ = v59
        } else {
            for fi60 := 0; fi60 < 3; fi60++ {
                v54 = ((v54 + (((v54 + v54) + ((v54 + ((v54 + "foo bar") + v54)) + "hello")) + v54)) + (("foo bar" + v54) + v54))
                v54 = "\n\t"
                var v61 int8 = 127
                _ = fi60
                _ = v61
            }
            var v62 uint8 = 42
            if false {
                v63 := GenS8{F0: -9223372036854775808, F1: genFn0(491, ("foo bar" + v54))}
                switch (v56 - -9223372036854775808) {
                case 0:
                    for fi64 := 0; fi64 < 8; fi64++ {
                        for fi65 := 0; fi65 < 13; fi65++ {
                            v66 := GenS9{F0: 42, F1: -496, F2: v1}
                            v56 = fi65
                            _ = fi65
                            _ = v66
                        }
                        var v67 uint = (v1 << (((v1 >> ((18446744073709551615) & 63))) & 63))
                        _ = fi64
                        _ = v67
                    }
                }
                _ = v63
            }
            _ = v62
        }
        _ = v56
    } else {
        v68 := GenS9{F0: 9223372036854775807, F1: -911, F2: (^(v1))}
        v1 = v1
        _ = v68
    }
    v1 = (v1 + v1)
    switch 32767 {
    case 127:
        for fi69 := 0; fi69 < 15; fi69++ {
            switch 32767 {
            case -2147483648:
                var v70 float64 = float64((float32(1.5) * 0.1))
                _ = v70
            case 9223372036854775807:
                v1 = v1
            case -1:
                switch "\n\t" {
                case "foo bar":
                    v71 := []int{(fi69 << ((v1) & 63))}
                    _ = v71
                case "a":
                    var v72 int32 = 0
                    _ = v72
                }
            default:
                var v73 uint = (v1 % (((v1 * v1)) | 1))
                _ = v73
            }
            v74 := GenS5{F0: (^(v1))}
            switch (fi69 << (((v1 ^ v1)) & 63)) {
            case 0:
                var v75 string = (v54 + ("hello" + v54))
                _ = v75
            case 32767:
                switch ((v54 + "foo bar") + "a") {
                case "foo bar":
                    var v76 uint16 = 0
                    _ = v76
                case "hello":
                    if true {
                        var v77 bool = false
                        for fi78 := 0; fi78 < 5; fi78++ {
                            v54 = (v54 + v54)
                            v79 := []int{4294967295, (fi69 >> (((v1 - (v1 / (((v1 | v1)) | 1)))) & 63)), (fi78 + genFn0(197, v54))}
                            _ = fi78
                            _ = v79
                        }
                        _ = v77
                    }
                case "":
                    if false {
                        fmt.Println(v54)
                    }
                }
            case 316:
                if true {
                    var v80 uint32 = 65535
                    for fi81 := 0; fi81 < 18; fi81++ {
                        var v82 int64 = int64(v80)
                        v54 = (("foo bar" + v54) + v54)
                        v54 = (v54 + "foo bar")
                        if false {
                            v54 = "\n\t"
                        } else {
                            v83 := []int{}
                            var v84 uint32 = uint32(v1)
                            if (!((genFn1() <= (v82 & (-(v82)))) && ((!true) || (((v1 * 665) > (v1 / (((v1 % ((795) | 1))) | 1))) && false)))) {
                                if (true || false) {
                                    var v85 uint64 = 42
                                    _ = v85
                                } else {
                                    if false {
                                        var v86 uint32 = (v80 >> ((v1) & 63))
                                        _ = v86
                                    } else {
                                        var v87 int32 = 1
                                        v84 += 145
                                        _ = v87
                                    }
                                    if false {
                                        fmt.Println(v54, v54, v54)
                                        var v88 int64 = (^(v82))
                                        v89 := GenS10{F0: fi81, F1: 802}
                                        v90 := GenS5{F0: (v1 - v1)}
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                    } else {
                                        fmt.Println(v1)
                                    }
                                }
                            } else {
                                v84 += 2147483648
                            }
                            _ = v83
                            _ = v84
                        }
                        _ = fi81
                        _ = v82
                    }
                    fmt.Println(fi69, v54)
                    _ = v80
                }
            }
            _ = fi69
            _ = v74
        }
    case -1:
        var v91 uint = v1
        _ = v91
    default:
        v1 -= 2
    }
    _ = len(v3)
    v54 = v54
    if ((!(0 >= 719)) && true) {
        v92 := []int{2147483647}
        v1 = v1
        v54 = (v54 + ("foo bar" + "a"))
        _ = v92
    } else {
        fmt.Println(v54, v54)
    }
    var v93 uint8 = 42
    v54 = v54
    var v94 string = v54
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v93
    _ = v94
}
