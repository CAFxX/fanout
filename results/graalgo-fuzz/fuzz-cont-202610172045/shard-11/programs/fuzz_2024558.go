package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 string; F2 int }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int; F1 int; F2 int32 }
type GenS5 struct { F0 string; F1 int; F2 uint }
type GenS6 struct { F0 int64; F1 bool; F2 int64 }
type GenS7 struct { F0 uint16; F1 int }
func genFn0(p0 uint32, p1 uint8) uint32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 bool = false
        p1 = p1
        var v7 uint = 889
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 42 {
    case 127:
        p0 = p0
    }
    if false {
        var v8 float32 = 0.0
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
        v8 -= (v8 - v8)
        _ = v8
    }
    p0 = 494
    v11 := []int{-2147483648, 650, -2147483648}
    _ = p0
    _ = p1
    _ = v11
    return (^(p0))
}
func genFn1(p0 uint) uint8 {
    v12 := GenS1{F0: 424, F1: 42, F2: false}
    v13 := GenS1{F0: 788, F1: 42, F2: (true || true)}
    var v14 uint32 = 0
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v14 -= v14
            switch "a" {
            case "hello":
                v17 := GenS2{F0: 1, F1: "a", F2: -577}
                v14 = 562
                if true {
                    v18 := GenS1{F0: 0, F1: 1, F2: true}
                    var v19 uint64 = 860
                    v19 -= 2
                    var v20 uint16 = 523
                    v21 := GenS2{F0: -1, F1: "a", F2: (fi15 - fi15)}
                    v22 := []int{(-(fi16)), (fi16 | (^(fi15)))}
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if (false && ((3.14159 * 1.5) > 0.0)) {
                    var v23 int16 = 1
                    fmt.Println(p0, fi16, v23)
                    _ = v23
                } else {
                    v24 := GenS1{F0: 42, F1: uint8(v14), F2: false}
                    if (false && (0 >= 825)) {
                        var v25 bool = false
                        v26 := []int{(fi16 >> (((p0 % ((0) | 1))) & 63))}
                        var v27 uint8 = 243
                        if v25 {
                            fmt.Println(fi15, v27, v14)
                        } else {
                            switch -2147483648 {
                            case -762:
                                v28 := GenS2{F0: 0, F1: "", F2: (-(fi16))}
                                _ = v28
                            default:
                                var v29 float64 = 3.14159
                                _ = v29
                            }
                        }
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        switch fi16 {
                        case 1:
                            v14 -= v14
                        case -1:
                            switch "\n\t" {
                            case "\n\t":
                                switch ("" + (("\n\t" + (("hello" + ("\n\t" + ("\n\t" + ""))) + "")) + ("hello" + "a"))) {
                                case "\n\t":
                                    var v30 uint = 1
                                    _ = v30
                                case "a":
                                    var v31 uint8 = 0
                                    _ = v31
                                default:
                                    if (false || (0 >= 42)) {
                                        v32 := []int{fi16, 32768}
                                        var v33 int32 = 467
                                        v33 = (v33 & int32(fi16))
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v34 := GenS3{F0: "hello", F1: 598}
                                        v35 := GenS1{F0: 0, F1: 0, F2: false}
                                        p0 += p0
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            case "":
                                if (654 < 2) {
                                    var v36 int32 = 406
                                    p0 -= p0
                                    _ = v36
                                }
                            default:
                                fmt.Println(fi16, v14)
                            }
                        case 127:
                            v14 -= genFn0(138, uint8(v14))
                        default:
                            switch (("\n\t" + "hello") + "a") {
                            case "":
                                v37 := GenS1{F0: 2, F1: uint8(v14), F2: ((^(v14)) != v14)}
                                _ = v37
                            case "a":
                                fmt.Println(fi15, p0, v14)
                            }
                        }
                        v38 := GenS2{F0: -1, F1: "\n\t", F2: fi16}
                        _ = v38
                    }
                    _ = v24
                }
                var v39 uint = p0
                _ = v17
                _ = v39
            default:
                switch fi15 {
                case 32767:
                    fmt.Println(fi15, fi16, v14)
                case -521:
                    var v40 int32 = 42
                    _ = v40
                }
            }
            fmt.Println(fi16, fi15)
            var v41 float32 = (-2.25 / -2.25)
            _ = fi16
            _ = v41
        }
        v42 := GenS4{F0: (fi15 >> ((p0) & 63)), F1: 32768, F2: -32}
        _ = fi15
        _ = v42
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    var v3 uint = 42
    v4 := make(map[string]bool)
    delete(v1, -2147483648)
    switch 128 {
    case 1:
        v3 = 0
    default:
        v3 -= (^(v3))
    }
    if (32767 <= 1) {
        var v43 int = 32768
        for fi44 := 0; fi44 < 10; fi44++ {
            v45 := GenS4{F0: fi44, F1: 32768, F2: -1}
            if false {
                v2 = 1e10
                v2 += float64(0.1)
            }
            _ = fi44
            _ = v45
        }
        _ = v43
    } else {
        switch "" {
        case "\n\t":
            if (-1 == 1) {
                v3 = v3
                if true {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        for fi47 := 0; fi47 < 5; fi47++ {
                            switch "" {
                            case "foo bar":
                                v3 = (^(v3))
                            default:
                                v48 := GenS5{F0: ("a" + ("a" + ("" + ("" + ("" + "foo bar"))))), F1: (fi47 << ((v3) & 63)), F2: v3}
                                v49 := []int{128, (fi46 | 4294967295), (fi46 * (fi47 / (((fi46 >> (((v3 >> ((v3) & 63))) & 63))) | 1)))}
                                if false {
                                    var v50 float64 = float64(1.5)
                                    _ = v50
                                } else {
                                    v51 := GenS4{F0: 578, F1: fi46, F2: 331}
                                    _ = v51
                                }
                                _ = v48
                                _ = v49
                            }
                            _ = fi47
                        }
                        v2 -= (0.1 - (1e10 + v2))
                        switch fi46 {
                        case 32767:
                            v2 -= (((0.1 + ((0.1 + v2) * v2)) - (float64((1e10 / (0.0 - 1e10))) + (0.1 - v2))) * 3.14159)
                        case 581:
                            v52 := []int{(fi46 << ((v3) & 63)), fi46}
                            _ = v52
                        }
                        _ = fi46
                    }
                } else {
                    var v53 bool = (394 > 78)
                    var v54 int32 = 1
                    _ = v53
                    _ = v54
                }
            } else {
                var v55 int16 = 32767
                fmt.Println(v2, v2)
                _ = v55
            }
        case "a":
            v56 := GenS6{F0: 1, F1: false, F2: -944}
            _ = v56
        case "foo bar":
            var v57 int = 1
            _ = v57
        default:
            var v58 int8 = 127
            _ = v58
        }
        for fi59 := 0; fi59 < 20; fi59++ {
            switch ("hello" + ("\n\t" + "")) {
            case "\n\t":
                v60 := GenS3{F0: ("" + "a"), F1: -819}
                _ = v60
            case "foo bar":
                v3 -= (v3 & v3)
            }
            _ = fi59
        }
    }
    var v61 uint = (v3 ^ (^(v3)))
    _ = len(v1)
    if (v3 != v61) {
        v62 := GenS3{F0: "a", F1: int64(v3)}
        v2 += v2
        _ = v62
    } else {
        v63 := []int{}
        v61 = 2
        for fi64 := 0; fi64 < 12; fi64++ {
            var v65 int32 = -2147483648
            if false {
                switch fi64 {
                case 0:
                    var v66 uint64 = 42
                    _ = v66
                case -1:
                    if (v65 > v65) {
                        v65 = (v65 + (-(v65)))
                    }
                case 4294967295:
                    v67 := GenS2{F0: 1, F1: "", F2: 369}
                    _ = v67
                }
                for fi68 := 0; fi68 < 5; fi68++ {
                    if false {
                        v3 = (v3 << (((v61 * v3)) & 63))
                        v69 := GenS2{F0: 42, F1: "foo bar", F2: fi64}
                        if true {
                            v70 := GenS1{F0: 368, F1: 201, F2: (3.14159 == 0.1)}
                            fmt.Println(v65)
                            _ = v70
                        } else {
                            if false {
                                if len(v63) >= 2 {
                                    v63 = v63[1:len(v63)]
                                }
                                var v71 uint16 = 42
                                _ = v71
                            } else {
                                var v72 float32 = (0.0 * ((0.1 - (1.5 - 1e10)) * 1e10))
                                _ = len(v63)
                                delete(v4, "\n\t")
                                _ = v72
                            }
                            v73 := v4[("hello" + ("foo bar" + "a"))]
                            var v74 uint64 = 0
                            fmt.Println(v74)
                            _ = v73
                            _ = v74
                        }
                        switch (fi68 * 338) {
                        case 128:
                            if false {
                                var v75 string = ""
                                _ = v75
                            } else {
                                if false {
                                    fmt.Println(v65)
                                } else {
                                    v76 := GenS5{F0: ("\n\t" + ("hello" + "\n\t")), F1: -1, F2: (v61 * v61)}
                                    if (((3.14159 + v2) >= 1.5) || (true && (false && (65535 <= 2147483647)))) {
                                        var v77 int8 = -82
                                        v2 = 1e10
                                        if (!((true && (!(false || (fi68 > (fi68 << (((v3 + (^(v3)))) & 63)))))) && true)) {
                                            v65 -= -1
                                            var v78 uint = v3
                                            if false {
                                                v77 = v77
                                            } else {
                                                _ = len(v63)
                                                var v79 float32 = (float32((((v2 + float64(1e10)) - -2.25) - v2)) * (float32((v2 / (1e10 * 1e10))) * ((1e10 + 0.0) - float32(1e10))))
                                                _ = v79
                                            }
                                            _ = v78
                                        } else {
                                            var v80 uint = (v3 * 994)
                                            var v81 uint16 = 42
                                            var v82 uint16 = v81
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        }
                                        if false {
                                            var v83 int64 = 1
                                            v3 += (^(v61))
                                            v83 = int64(v77)
                                            _ = v83
                                        } else {
                                            v3 -= (v3 & (^(v3)))
                                        }
                                        fmt.Println(v65)
                                        if false {
                                            var v84 bool = true
                                            _ = v84
                                        }
                                        var v85 uint8 = 0
                                        _ = v77
                                        _ = v85
                                    }
                                    _ = v76
                                }
                            }
                        default:
                            fmt.Println(fi68)
                        }
                        v86 := GenS7{F0: 2, F1: (fi68 * (fi64 - 42))}
                        delete(v4, "\n\t")
                        v61 -= v61
                        _ = v69
                        _ = v86
                    } else {
                        var v87 bool = false
                        fmt.Println(v3)
                        _ = v87
                    }
                    _ = fi68
                }
            }
            _ = fi64
            _ = v65
        }
        _ = v63
    }
    for fi88 := 0; fi88 < 14; fi88++ {
        if ((0.1 < v2) && false) {
            v61 -= v61
            for fi89 := 0; fi89 < 10; fi89++ {
                v90 := []int{(-(fi89)), (fi89 + fi89)}
                _ = fi89
                _ = v90
            }
            var v91 int64 = 0
            _ = v91
        }
        _ = fi88
    }
    v2 = float64((((1e10 + 3.14159) / -2.25) * (1e10 * 3.14159)))
    v92 := GenS4{F0: -9223372036854775808, F1: 42, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v92
}
