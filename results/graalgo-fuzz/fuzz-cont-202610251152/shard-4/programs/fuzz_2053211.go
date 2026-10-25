package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int; F1 string; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 int; F2 int64 }
type GenS10 struct { F0 uint8; F1 int; F2 uint16 }
type GenS11 struct { F0 uint; F1 uint64; F2 uint }
type GenS12 struct { F0 string; F1 int16; F2 int8 }
type GenS13 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS14 struct { F0 int }
func genFn0(p0 int64, p1 string, p2 int32) int32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch (fi4 / (((fi4 | (fi4 & (^(fi4))))) | 1)) {
        case 1:
            for fi5 := 0; fi5 < 2; fi5++ {
                p0 = p0
                var v6 bool = false
                var v7 float32 = -2.25
                if ((fi4 >> ((230) & 63)) != (fi4 ^ -934)) {
                    p2 += (p2 % ((709) | 1))
                } else {
                    var v8 uint32 = 65535
                    if v6 {
                        if v6 {
                            p1 = p1
                            switch ("" + "foo bar") {
                            case "foo bar":
                                var v9 uint16 = 1
                                _ = v9
                            case "hello":
                                var v10 int16 = -1
                                _ = v10
                            case "a":
                                v11 := GenS1{F0: 42, F1: 2, F2: p1}
                                _ = v11
                            }
                            v12 := GenS2{F0: ("" + "hello")}
                            _ = v12
                        } else {
                            v13 := GenS2{F0: p1}
                            v14 := []int{-1, 333, (fi5 % ((fi4) | 1))}
                            _ = v13
                            _ = v14
                        }
                    } else {
                        var v15 int = (fi5 >> ((4294967295) & 63))
                        _ = v15
                    }
                    var v16 float32 = (0.1 / -2.25)
                    _ = v16
                    _ = v8
                }
                var v17 uint16 = 32767
                if (-1 == 42) {
                    if false {
                        v18 := []int{-9223372036854775808, fi5, (-(fi4))}
                        if (!(32767 < 0)) {
                            var v19 int16 = -1
                            v20 := GenS3{F0: 65535, F1: 128, F2: v6}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint16 = v17
                            _ = v21
                        }
                        _ = v18
                    }
                    v7 -= (((v7 / 1.0) - (v7 / float32(0.1))) * (3.14159 / v7))
                    var v22 float32 = v7
                    if v6 {
                        v23 := GenS4{F0: 768, F1: 0}
                        var v24 uint = 2
                        if ((9223372036854775807 <= (-(p0))) && true) {
                            v24 += (v24 - v24)
                        } else {
                            v25 := GenS5{F0: v17, F1: 9223372036854775807}
                            switch fi4 {
                            case 128:
                                v26 := []int{}
                                _ = v26
                            case -765:
                                v27 := []int{}
                                _ = v27
                            }
                            v28 := []int{}
                            _ = v25
                            _ = v28
                        }
                        p0 = 42
                        _ = v23
                        _ = v24
                    } else {
                        if (v6 || true) {
                            if false {
                                if (((2 >= 42) && (true || false)) || (v6 || (!(999 == 2)))) {
                                    v29 := []int{}
                                    v30 := GenS1{F0: 42, F1: 255, F2: p1}
                                    switch -2147483648 {
                                    case -267:
                                        var v31 uint16 = 269
                                        _ = v31
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    p0 += p0
                                    if (32767 != int16(fi5)) {
                                        fmt.Println(fi5)
                                    } else {
                                        if (!(!true)) {
                                            if v6 {
                                                var v32 float32 = v22
                                                if v6 {
                                                    v32 += v7
                                                    v33 := GenS6{F0: (fi5 << ((42) & 63)), F1: ((p1 + p1) + ""), F2: (-(fi4))}
                                                    var v34 uint16 = v17
                                                    var v35 uint16 = v34
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v7 += (float32(-2.25) * v32)
                                                }
                                                var v36 int16 = 1000
                                                _ = v32
                                                _ = v36
                                            } else {
                                                var v37 uint8 = 223
                                                _ = v37
                                            }
                                            v38 := GenS2{F0: p1}
                                            _ = v38
                                        }
                                    }
                                    var v39 float64 = 3.14159
                                    _ = v39
                                }
                            } else {
                                switch (p1 + ((p1 + p1) + p1)) {
                                case "":
                                    v40 := []int{(fi4 & fi5), (fi5 * fi4), fi5}
                                    _ = v40
                                case "foo bar":
                                    v41 := []int{}
                                    v6 = true
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                    _ = v41
                                case "hello":
                                    v42 := GenS7{F0: fi5}
                                    _ = v42
                                default:
                                    v43 := []int{fi4, fi4, fi5}
                                    _ = v43
                                }
                            }
                        }
                    }
                    v17 -= 0
                    _ = v22
                }
                _ = fi5
                _ = v17
                _ = v6
                _ = v7
            }
        case 9223372036854775807:
            v44 := GenS7{F0: (fi4 >> ((4294967295) & 63))}
            _ = v44
        case 32767:
            p1 = ("a" + ("" + p1))
        }
        _ = fi4
    }
    v45 := GenS3{F0: 2, F1: 1, F2: (158 == 42)}
    _ = p0
    _ = p1
    _ = p2
    _ = v45
    return (p2 >> ((258) & 63))
}

func main() {
    var v1 int32 = 289
    var v2 int32 = 1
    v3 := make(map[string]string)
    delete(v3, "\n\t")
    switch 32767 {
    case -455:
        if (!false) {
            v3[("hello" + "")] = "foo bar"
            for fi46 := 0; fi46 < 1; fi46++ {
                var v47 uint8 = 0
                fmt.Println(v2, v47)
                _ = fi46
                _ = v47
            }
        }
    default:
        v1 = (^(v1))
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        fmt.Println(fi48, v1)
        if true {
            v1 = 42
            v2 = (v1 >> ((2) & 63))
            for fi49 := 0; fi49 < 12; fi49++ {
                v3["hello"] = ("a" + "")
                _ = fi49
            }
        } else {
            v50 := v3["hello"]
            v50 = (((v50 + "foo bar") + (v50 + "a")) + v50)
            _ = v50
        }
        v2 -= v1
        _ = fi48
    }
    v2 -= v1
    v51 := GenS4{F0: 390, F1: 921}
    v52 := GenS8{F0: "foo bar"}
    v53 := []int{-271, 42}
    fmt.Println(v2)
    v1 = genFn0(-539, ((("\n\t" + (("" + (("\n\t" + ("a" + ((("\n\t" + "foo bar") + "\n\t") + "foo bar"))) + "a")) + ("\n\t" + ""))) + ("\n\t" + (("" + ("" + "foo bar")) + "hello"))) + ""), v1)
    if len(v53) > 0 {
        v53[0] = 42
    }
    if true {
        switch -1 {
        case 300:
            if false {
                fmt.Println(v2)
                if (4294967295 >= 255) {
                    for fi54 := 0; fi54 < 18; fi54++ {
                        if false {
                            v55 := GenS5{F0: 2, F1: 128}
                            fmt.Println(fi54, fi54)
                            _ = len(v53)
                            _ = v55
                        }
                        v1 -= (^(v2))
                        var v56 string = "\n\t"
                        _ = fi54
                        _ = v56
                    }
                    if false {
                        _ = len(v3)
                        var v57 int32 = v1
                        var v58 uint = 65535
                        _ = v57
                        _ = v58
                    } else {
                        var v59 float32 = 0.1
                        _ = v59
                    }
                }
                for fi60 := 0; fi60 < 2; fi60++ {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        if true {
                            var v62 int32 = v2
                            _ = v62
                        } else {
                            switch (^(fi61)) {
                            case 32767:
                                var v63 string = "hello"
                                _ = v63
                            case 0:
                                v1 += genFn0(0, "hello", (-(v1)))
                            case 127:
                                v64 := GenS9{F0: 42, F1: (fi60 & (fi61 - (fi61 | fi60))), F2: 446}
                                _ = v64
                            }
                            var v65 bool = (127 <= 72)
                            v66 := GenS1{F0: 4294967295, F1: 2, F2: "a"}
                            _ = v65
                            _ = v66
                        }
                        _ = len(v53)
                        var v67 uint8 = 2
                        _ = fi61
                        _ = v67
                    }
                    var v68 bool = false
                    v1 -= v2
                    _ = fi60
                    _ = v68
                }
            } else {
                var v69 bool = true
                _ = v69
            }
        case -818:
            switch "\n\t" {
            case "foo bar":
                if (127 < 128) {
                    v2 = (-(v2))
                    v70 := GenS4{F0: 1, F1: 305}
                    _ = v70
                }
            case "hello":
                switch ("a" + ("hello" + ((("a" + "") + "hello") + "\n\t"))) {
                case "hello":
                    v2 = 42
                case "a":
                    var v71 string = "foo bar"
                    _ = v71
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        switch 954 {
                        case 573:
                            v1 += v1
                        case -785:
                            fmt.Println(v1)
                        case 466:
                            for fi72 := 0; fi72 < 12; fi72++ {
                                v2 = v2
                                v73 := GenS3{F0: 0, F1: 127, F2: (!true)}
                                _ = fi72
                                _ = v73
                            }
                        default:
                            var v74 int32 = genFn0(-1, "a", v1)
                            _ = v74
                        }
                    case "a":
                        v75 := GenS10{F0: 127, F1: 900, F2: 1}
                        _ = v75
                    case "hello":
                        fmt.Println(v2, v2)
                    }
                }
            case "":
                _ = len(v53)
            default:
                v76 := GenS11{F0: 4294967295, F1: 42, F2: 4294967295}
                _ = v76
            }
        default:
            switch "foo bar" {
            case "foo bar":
                for fi77 := 0; fi77 < 6; fi77++ {
                    for fi78 := 0; fi78 < 17; fi78++ {
                        switch (fi77 | fi78) {
                        case 32767:
                            v1 -= v1
                        case 1:
                            v79 := v3[(("a" + "a") + "a")]
                            _ = v79
                        }
                        _ = fi78
                    }
                    for fi80 := 0; fi80 < 7; fi80++ {
                        v1 -= (v2 << ((4294967295) & 63))
                        _ = fi80
                    }
                    var v81 uint16 = 741
                    _ = fi77
                    _ = v81
                }
            case "a":
                for fi82 := 0; fi82 < 7; fi82++ {
                    var v83 int = (fi82 % ((fi82) | 1))
                    var v84 uint = 2
                    v85 := GenS12{F0: "foo bar", F1: 766, F2: 103}
                    _ = fi82
                    _ = v83
                    _ = v84
                    _ = v85
                }
            case "hello":
                if len(v53) >= 2 {
                    v53 = v53[1:len(v53)]
                }
            default:
                switch "hello" {
                case "hello":
                    v2 += (v2 ^ v1)
                }
            }
        }
    }
    v86 := GenS11{F0: 881, F1: 1, F2: 666}
    if false {
        v87 := GenS13{F0: 1, F1: 255, F2: 749}
        _ = v87
    } else {
        var v88 bool = (0 <= 2)
        _ = v88
    }
    if len(v53) >= 2 {
        v53 = v53[1:len(v53)]
    }
    for fi89 := 0; fi89 < 4; fi89++ {
        fmt.Println(v1, v2, v1)
        _ = fi89
    }
    var v90 int8 = 42
    if true {
        v91 := GenS14{F0: -153}
        _ = v91
    } else {
        var v92 uint32 = 4294967295
        _ = v92
    }
    delete(v3, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v86
    _ = v90
}
