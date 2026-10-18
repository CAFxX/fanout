package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 int; F2 uint32 }
type GenS3 struct { F0 int; F1 int; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int8; F2 string }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int16, p1 uint16) int {
    if true {
        v4 := GenS1{F0: 2147483647, F1: 2147483647, F2: p0}
        _ = v4
    } else {
        var v5 int = 127
        v6 := GenS2{F0: 60, F1: (v5 + -925), F2: 4294967295}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: -1, F1: -9223372036854775808, F2: -1}
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "":
        switch "" {
        case "a":
            v10 := []int{2147483647, 644, 32767}
            _ = v10
        case "foo bar":
            switch "hello" {
            case "\n\t":
                switch ((("foo bar" + "") + "") + "foo bar") {
                case "foo bar":
                    v11 := []int{42}
                    _ = v11
                case "":
                    v12 := []int{-945, 32767, 42}
                    _ = v12
                default:
                    var v13 uint = 2
                    _ = v13
                }
            case "a":
                v14 := []int{}
                _ = v14
            }
        case "hello":
            var v15 uint8 = 2
            _ = v15
        }
    case "\n\t":
        if (p1 < 2) {
            fmt.Println(p0, p1)
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "" {
                case "\n\t":
                    if true {
                        var v17 float32 = ((float32(float64(((((3.14159 * -2.25) - 1.5) - -2.25) / -2.25))) / 1e10) + 0.1)
                        p1 += p1
                        _ = v17
                    }
                }
                _ = fi16
            }
        } else {
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS3{F0: (fi18 >> ((42) & 63)), F1: -9223372036854775808, F2: "hello"}
                _ = fi18
                _ = v19
            }
            v20 := GenS2{F0: 0, F1: 128, F2: 42}
            for fi21 := 0; fi21 < 8; fi21++ {
                switch "\n\t" {
                case "":
                    v22 := []int{(fi21 >> ((255) & 63)), fi21}
                    _ = v22
                case "\n\t":
                    switch (fi21 + (fi21 / (((fi21 / ((fi21) | 1))) | 1))) {
                    case 266:
                        fmt.Println(p0, p0)
                    case 9223372036854775807:
                        var v23 string = (("" + "hello") + "")
                        _ = v23
                    default:
                        var v24 int32 = 0
                        _ = v24
                    }
                case "foo bar":
                    var v25 int = (fi21 & -196)
                    _ = v25
                default:
                    v26 := []int{(fi21 >> ((587) & 63)), fi21, 9223372036854775807}
                    _ = v26
                }
                _ = fi21
            }
            _ = v20
        }
    default:
        p1 += 32767
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]bool)
    if (0 > 0) {
        var v27 uint8 = 127
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(fi28)
            fmt.Println(fi28, v1)
            _ = fi29
        }
        v30 := GenS2{F0: 54, F1: fi28, F2: 2147483648}
        _ = fi28
        _ = v30
    }
    var v31 uint16 = 734
    v31 += 0
    v2 -= v2
    if (1.5 == 0.1) {
        if false {
            for fi32 := 0; fi32 < 16; fi32++ {
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := GenS3{F0: genFn0(1, 1), F1: -796, F2: ""}
                    _ = fi33
                    _ = v34
                }
                v35 := GenS2{F0: 1, F1: fi32, F2: 2}
                _ = fi32
                _ = v35
            }
        }
        v2 += v2
    } else {
        var v36 float32 = -2.25
        switch "\n\t" {
        case "a":
            var v37 uint64 = (v1 << ((78) & 63))
            _ = v37
        case "":
            if true {
                fmt.Println(v31, v36, v1)
                switch "foo bar" {
                case "":
                    for fi38 := 0; fi38 < 16; fi38++ {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v36, v36, v2)
                        case "a":
                            switch 128 {
                            case 400:
                                var v39 int8 = -1
                                _ = v39
                            case 0:
                                for fi40 := 0; fi40 < 11; fi40++ {
                                    fmt.Println(v1, v2)
                                    _ = fi40
                                }
                            default:
                                if (!true) {
                                    var v41 uint64 = v1
                                    v2 += v2
                                    _ = v41
                                } else {
                                    delete(v3, genFn0(0, (^(v31))))
                                    v42 := GenS2{F0: 1, F1: 104, F2: 311}
                                    v2 = (v2 * 515)
                                    _ = v42
                                }
                            }
                        case "foo bar":
                            fmt.Println(v2, v31, fi38)
                        }
                        _ = fi38
                    }
                case "\n\t":
                    if ((v31 * 255) == v31) {
                        v43 := GenS4{F0: (v31 | (v31 * (v31 >> ((369) & 63))))}
                        var v44 int64 = -959
                        _ = v43
                        _ = v44
                    } else {
                        if true {
                            var v45 uint32 = 2147483647
                            v3[genFn0(1000, 65535)] = (42 < (v45 | v45))
                            var v46 bool = false
                            _ = v45
                            _ = v46
                        }
                        fmt.Println(v2, v31, v31)
                        if true {
                            fmt.Println(v1, v2, v2)
                        } else {
                            v1 = 740
                            v47 := []int{-2147483648, -1, 2147483647}
                            fmt.Println(v36)
                            _ = v47
                        }
                    }
                default:
                    var v48 float32 = (v36 - v36)
                    _ = v48
                }
                switch ("foo bar" + ("foo bar" + "a")) {
                case "foo bar":
                    _ = len(v3)
                case "":
                    if (0 >= genFn0(-1, v31)) {
                        fmt.Println(v31, v2, v36)
                    }
                default:
                    switch 892 {
                    case -2147483648:
                        v49 := GenS5{F0: 1, F1: 14}
                        _ = v49
                    case 788:
                        v50 := GenS6{F0: 42, F1: v2}
                        _ = v50
                    default:
                        if (1000 <= -32768) {
                            if false {
                                fmt.Println(v31, v2)
                                if false {
                                    switch -362 {
                                    case 2147483647:
                                        var v51 float32 = 1e10
                                        _ = v51
                                    case -388:
                                        fmt.Println(v2, v31, v36)
                                    default:
                                        v2 = v2
                                    }
                                } else {
                                    switch genFn0(1000, v31) {
                                    case 32768:
                                        for fi52 := 0; fi52 < 16; fi52++ {
                                            fmt.Println(v31, v31)
                                            v36 -= float32(0.0)
                                            v36 = v36
                                            _ = fi52
                                        }
                                    case 0:
                                        var v53 int = 0
                                        _ = v53
                                    case 636:
                                        if (577 < 65535) {
                                            v54 := []int{32768, genFn0(549, v31), 2147483647}
                                            v55 := GenS7{F0: false}
                                            v31 += 42
                                            _ = v54
                                            _ = v55
                                        } else {
                                            if (!true) {
                                                v1 += (v1 % ((v1) | 1))
                                                if false {
                                                    if (!(true && (-32768 == 655))) {
                                                        fmt.Println(v31, v1, v1)
                                                        if false {
                                                            var v56 int16 = 1
                                                            fmt.Println(v36)
                                                            switch "" {
                                                            case "\n\t":
                                                                v56 = (-(v56))
                                                            case "a":
                                                                v2 = (v2 * (v2 * v2))
                                                            case "foo bar":
                                                                var v57 uint16 = 0
                                                                _ = v57
                                                            default:
                                                                v58 := GenS5{F0: -61, F1: v56}
                                                                _ = v58
                                                            }
                                                            _ = v56
                                                        } else {
                                                            v59 := GenS4{F0: 32767}
                                                            _ = v59
                                                        }
                                                    } else {
                                                        for fi60 := 0; fi60 < 10; fi60++ {
                                                            v31 = v31
                                                            v31 += (v31 >> ((1) & 63))
                                                            switch (fi60 << ((18446744073709551615) & 63)) {
                                                            case -1:
                                                                for fi61 := 0; fi61 < 18; fi61++ {
                                                                    v62 := []int{459, genFn0(42, v31)}
                                                                    v2 = (v2 * v2)
                                                                    switch -906 {
                                                                    case -2147483648:
                                                                        fmt.Println(v2, v36, fi60)
                                                                    }
                                                                    var v63 uint64 = 18446744073709551615
                                                                    fmt.Println(v36, fi61)
                                                                    v63 = 0
                                                                    _ = fi61
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            case -752:
                                                                if (405 != v2) {
                                                                    v64 := GenS3{F0: fi60, F1: (fi60 << ((104) & 63)), F2: ("foo bar" + (("a" + "") + (("a" + ("" + "hello")) + "")))}
                                                                    _ = v64
                                                                } else {
                                                                    var v65 uint = 0
                                                                    if (!false) {
                                                                        v1 -= v1
                                                                    } else {
                                                                        if (!false) {
                                                                            var v66 uint32 = 274
                                                                            fmt.Println(fi60, v36)
                                                                            fmt.Println(v36, v31, v36)
                                                                            _ = v66
                                                                        } else {
                                                                            v67 := GenS8{F0: 1, F1: 45, F2: ("a" + (("hello" + "foo bar") + ("hello" + "foo bar")))}
                                                                            _ = v67
                                                                        }
                                                                        fmt.Println(v65, v31)
                                                                        v68 := []int{}
                                                                        _ = v68
                                                                    }
                                                                    v69 := GenS9{F0: (((false && (v36 <= (((v36 * (0.1 / ((v36 + 1.5) * v36))) - v36) * v36))) || true) || true), F1: (v31 >> ((2) & 63))}
                                                                    _ = v65
                                                                    _ = v69
                                                                }
                                                            case -309:
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v70 uint16 = v31
                                                                    _ = v70
                                                                case "a":
                                                                    v71 := GenS2{F0: 100, F1: fi60, F2: 65535}
                                                                    _ = v71
                                                                case "":
                                                                    fmt.Println(v1)
                                                                default:
                                                                    var v72 int8 = 37
                                                                    _ = v72
                                                                }
                                                            }
                                                            _ = fi60
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            var v73 float64 = 0.1
                                                            _ = v73
                                                        case "hello":
                                                            v1 = (^(v1))
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    if (true && ("\n\t" == "")) {
                                                        for fi74 := 0; fi74 < 9; fi74++ {
                                                            fmt.Println(v1)
                                                            v75 := GenS5{F0: 0, F1: -32768}
                                                            _ = fi74
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                                v76 := []int{9223372036854775807, 2147483647}
                                                _ = v76
                                            }
                                        }
                                    }
                                    v36 = v36
                                }
                            } else {
                                var v77 uint8 = 127
                                fmt.Println(v2)
                                v78 := GenS4{F0: v31}
                                _ = v77
                                _ = v78
                            }
                        }
                    }
                }
            }
        case "foo bar":
            v1 += 44
        }
        v79 := GenS10{F0: 65535, F1: uint32(v31)}
        _ = v36
        _ = v79
    }
    v1 += (v1 + (v1 ^ (v1 >> ((255) & 63))))
    fmt.Println(v2)
    v31 -= (v31 & v31)
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
