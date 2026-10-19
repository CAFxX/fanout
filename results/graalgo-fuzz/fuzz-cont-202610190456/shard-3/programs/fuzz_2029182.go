package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 int32; F2 bool }
type GenS7 struct { F0 uint; F1 bool }
func genFn0(p0 string, p1 int, p2 uint64) int {
    if (!(!false)) {
        var v6 int = (p1 / (((-(p1))) | 1))
        var v7 uint16 = 1
        _ = v6
        _ = v7
    }
    v8 := []int{}
    v9 := GenS1{F0: (p2 | (^(p2))), F1: int16(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 bool = (-32768 < 42)
    var v2 int32 = 0
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 int16 = 1000
    v5 = v5
    var v10 bool = v1
    switch "" {
    case "":
        if (false && v1) {
            var v11 int64 = -1
            var v12 bool = v1
            _ = v11
            _ = v12
        } else {
            if true {
                if v10 {
                    v1 = ((v5 + (v5 % ((-1) | 1))) >= 1)
                    switch "hello" {
                    case "hello":
                        v13 := []int{-9, 128, -2147483648, 0}
                        _ = v13
                    case "a":
                        v14 := []int{32768}
                        _ = v14
                    }
                    var v15 float64 = -2.25
                    _ = v15
                } else {
                    var v16 string = "\n\t"
                    if (1 == 2) {
                        switch -1 {
                        case 128:
                            fmt.Println(v5, v16)
                        default:
                            for fi17 := 0; fi17 < 16; fi17++ {
                                v16 = v16
                                var v18 int64 = 1
                                _ = fi17
                                _ = v18
                            }
                        }
                        var v19 bool = false
                        var v20 uint32 = 0
                        _ = v19
                        _ = v20
                    } else {
                        v16 = "a"
                    }
                    _ = v16
                }
            }
            v21 := GenS1{F0: 42, F1: v5}
            _ = v21
        }
    }
    v22 := GenS1{F0: 18446744073709551615, F1: (v5 & 1000)}
    switch (("a" + "a") + "foo bar") {
    case "\n\t":
        var v23 uint32 = 2
        _ = v23
    case "hello":
        v24 := []int{1, 463, 127, -553}
        _ = v24
    case "":
        v25 := GenS1{F0: 54, F1: 42}
        _ = v25
    default:
        v5 += (v5 & v5)
    }
    if v10 {
        fmt.Println(v5, v1)
    }
    v2 -= (v2 / ((1) | 1))
    if true {
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS2{F0: "", F1: -948}
            _ = fi26
            _ = v27
        }
        if (!("hello" >= ("hello" + (("\n\t" + "foo bar") + "\n\t")))) {
            v10 = (v10 || v10)
        } else {
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 19; fi29++ {
                    switch fi28 {
                    case 2147483647:
                        v30 := GenS1{F0: 29, F1: v5}
                        _ = v30
                    case 32768:
                        v31 := []int{}
                        _ = v31
                    default:
                        var v32 int64 = 0
                        _ = v32
                    }
                    switch (((("hello" + "foo bar") + "\n\t") + "a") + "a") {
                    case "a":
                        var v33 float64 = (-2.25 + 3.14159)
                        _ = v33
                    }
                    if v10 {
                        v10 = (v1 || ((1e10 / 0.1) <= 1e10))
                        v34 := GenS2{F0: "\n\t", F1: -603}
                        _ = v34
                    }
                    _ = fi29
                }
                v35 := GenS3{F0: "a", F1: (^(v5))}
                var v36 int64 = 42
                _ = fi28
                _ = v35
                _ = v36
            }
            if v1 {
                v37 := GenS2{F0: "", F1: 1}
                v1 = false
                var v38 uint8 = 0
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 12; fi39++ {
                if true {
                    v1 = ((v5 ^ 29) >= 1000)
                    v40 := GenS2{F0: "foo bar", F1: 1}
                    _ = v40
                }
                if (!true) {
                    var v41 bool = v10
                    v42 := GenS2{F0: "a", F1: -1000}
                    v1 = true
                    _ = v41
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 10; fi43++ {
                        var v44 float64 = (-2.25 / float64((3.14159 / 1.0)))
                        var v45 int32 = v2
                        if false {
                            v46 := []int{(fi43 << ((40) & 63)), (fi43 * genFn0("\n\t", (fi39 >> ((34) & 63)), 826)), fi39}
                            v4[2147483648] = "a"
                            v1 = true
                            _ = v46
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    switch (((("foo bar" + "foo bar") + (("\n\t" + "foo bar") + "a")) + ((("hello" + "a") + "") + "foo bar")) + "\n\t") {
                    case "\n\t":
                        var v47 uint8 = 127
                        _ = v47
                    case "a":
                        v1 = v10
                    case "hello":
                        fmt.Println(v10, v1, v2)
                    default:
                        v48 := GenS4{F0: -1, F1: -65}
                        _ = v48
                    }
                    v49 := []int{}
                    _ = v49
                }
                _ = fi39
            }
        }
    } else {
        switch "\n\t" {
        case "a":
            v1 = false
        case "":
            if (v1 && (1 != 9223372036854775807)) {
                delete(v4, 883)
                v50 := []int{4294967295, 0, -355, 0}
                _ = len(v50)
                _ = v50
            } else {
                var v51 bool = false
                v2 = -2147483648
                v52 := GenS4{F0: -9223372036854775808, F1: -128}
                _ = v51
                _ = v52
            }
        case "foo bar":
            switch ((("" + "\n\t") + "foo bar") + "hello") {
            case "a":
                var v53 uint16 = 0
                _ = v53
            case "foo bar":
                if (v1 && (!(128 < genFn0("hello", -106, 0)))) {
                    v5 += (v5 * (v5 * -32768))
                    fmt.Println(v2, v5)
                    var v54 uint64 = 42
                    _ = v54
                } else {
                    var v55 uint64 = 4294967295
                    _ = v55
                }
            }
        }
        if ((v5 >> ((3) & 63)) == v5) {
            v56 := v4[65535]
            if v10 {
                v56 = ""
                for fi57 := 0; fi57 < 15; fi57++ {
                    for fi58 := 0; fi58 < 6; fi58++ {
                        switch v56 {
                        case "foo bar":
                            if (((v56 + v56) != ("" + v56)) || v10) {
                                var v59 int64 = -289
                                _ = v59
                            } else {
                                var v60 int8 = 1
                                v61 := GenS5{F0: genFn0("a", genFn0((v56 + v56), fi58, 794), 42), F1: 4294967295, F2: 42}
                                _ = v60
                                _ = v61
                            }
                        case "":
                            var v62 uint64 = 1
                            _ = v62
                        default:
                            var v63 bool = v10
                            _ = v63
                        }
                        var v64 float32 = 0.0
                        _ = fi58
                        _ = v64
                    }
                    _ = fi57
                }
                v5 = 1000
            }
            _ = v56
        }
        switch -95 {
        case 2147483647:
            fmt.Println(v2, v2)
        case -1:
            switch ("" + "\n\t") {
            case "hello":
                v65 := GenS6{F0: v10, F1: v2, F2: (32767 > 32767)}
                _ = v65
            case "foo bar":
                if (-2.25 == 3.14159) {
                    for fi66 := 0; fi66 < 1; fi66++ {
                        if (0 < 632) {
                            v1 = (v10 || v10)
                        }
                        _ = fi66
                    }
                    if (!((3.14159 / 1.0) >= (-2.25 * ((1.5 / 1e10) / 1.5)))) {
                        var v67 uint8 = 2
                        v2 -= -2147483648
                        _ = v67
                    } else {
                        var v68 int = genFn0("", -2147483648, 930)
                        _ = v68
                    }
                } else {
                    var v69 int = genFn0(("a" + ""), -1, 840)
                    var v70 string = ("foo bar" + "")
                    _ = v69
                    _ = v70
                }
            case "a":
                var v71 int32 = (v2 >> ((65535) & 63))
                _ = v71
            default:
                v5 = 548
            }
        }
    }
    fmt.Println(v10, v1)
    switch (("" + "a") + "foo bar") {
    case "":
        v2 = v2
    case "\n\t":
        switch "foo bar" {
        case "hello":
            fmt.Println(v1, v5)
        case "foo bar":
            v10 = (1 <= v5)
        case "":
            for fi72 := 0; fi72 < 11; fi72++ {
                var v73 uint16 = 255
                for fi74 := 0; fi74 < 15; fi74++ {
                    v75 := []int{}
                    switch (fi72 ^ 983) {
                    case -2147483648:
                        v1 = (true && v10)
                    case 2147483647:
                        v2 = v2
                    case -1:
                        v76 := GenS4{F0: 1, F1: 0}
                        _ = v76
                    }
                    _ = fi74
                    _ = v75
                }
                _ = fi72
                _ = v73
            }
        default:
            for fi77 := 0; fi77 < 11; fi77++ {
                v78 := GenS7{F0: 1, F1: v1}
                var v79 float32 = 0.0
                v2 = (v2 - (v2 >> ((473) & 63)))
                _ = fi77
                _ = v78
                _ = v79
            }
        }
    case "foo bar":
        v5 += v5
    }
    _ = len(v4)
    v80 := GenS5{F0: 300, F1: 2, F2: 884}
    var v81 uint64 = 42
    fmt.Println(v81, v81)
    var v82 int64 = 1
    fmt.Println(v1)
    var v83 bool = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
