package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool }

func main() {
    var v1 int8 = 1
    var v2 float64 = (1.5 / (0.1 / 0.1))
    var v3 int64 = 42
    switch ("" + ("" + "a")) {
    case "a":
        var v4 uint64 = 165
        _ = v4
    case "":
        v1 = 102
    case "\n\t":
        if (false || (!true)) {
            fmt.Println(v1, v2, v3)
            v5 := GenS1{F0: 42, F1: 360, F2: 127}
            _ = v5
        }
    default:
        fmt.Println(v1)
    }
    v1 -= (v1 << ((255) & 63))
    var v6 int16 = 136
    v2 += ((1.5 - v2) - (v2 + v2))
    var v7 int16 = 683
    var v8 float64 = -2.25
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint64 = 42
        v8 = (((((v2 + 3.14159) - (v8 - v8)) + (v8 - (float64(3.14159) + -2.25))) / 1.0) / v8)
        var v11 int8 = v1
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 += v8
    var v12 uint = 2
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            if (383 <= 1) {
                var v14 string = "hello"
                v15 := GenS1{F0: 720, F1: v12, F2: 42}
                _ = v14
                _ = v15
            } else {
                if false {
                    if true {
                        v16 := GenS1{F0: 497, F1: (v12 % ((42) | 1)), F2: 1}
                        v17 := []int{127, 848, (fi13 * (fi13 - fi13)), (^(fi13))}
                        v6 -= v7
                        _ = v16
                        _ = v17
                    }
                }
            }
            _ = fi13
        }
        v6 = v7
    } else {
        v3 += (v3 * (^(v3)))
        if (false || (!(v1 > v1))) {
            for fi18 := 0; fi18 < 14; fi18++ {
                if false {
                    v19 := []int{(fi18 & -343), 4294967295, fi18}
                    _ = v19
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    v7 += v7
                    if (!true) {
                        if ((v1 / (((v1 << ((189) & 63))) | 1)) > (-(v1))) {
                            v3 -= (v3 / (((v3 ^ v3)) | 1))
                        } else {
                            if true {
                                fmt.Println(v3)
                            }
                            var v21 uint = (v12 % ((v12) | 1))
                            v22 := GenS2{F0: 4294967295}
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{}
                        _ = v23
                    } else {
                        v3 = -394
                        v7 -= (^(v6))
                    }
                    var v24 int8 = -128
                    v25 := GenS2{F0: 1}
                    _ = fi20
                    _ = v24
                    _ = v25
                }
                _ = fi18
            }
            var v26 bool = false
            var v27 uint = 65535
            _ = v26
            _ = v27
        }
    }
    v8 -= v2
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int64 = v3
        if false {
            var v30 uint8 = uint8(v12)
            if true {
                v12 = (v12 & (v12 << (((v12 * v12)) & 63)))
                if false {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        switch "foo bar" {
                        case "hello":
                            var v32 uint8 = 0
                            _ = v32
                        case "a":
                            v6 = (^(v7))
                        case "\n\t":
                            v33 := []int{}
                            _ = v33
                        default:
                            if (!(645 > 850)) {
                                fmt.Println(v2)
                            }
                        }
                        if true {
                            v3 += v3
                            switch (fi31 ^ (fi28 >> (((^(v12))) & 63))) {
                            case 0:
                                var v34 int64 = 1
                                _ = v34
                            case -635:
                                if true {
                                    var v35 int64 = -1
                                    v36 := GenS3{F0: (v30 % ((v30) | 1))}
                                    _ = v35
                                    _ = v36
                                } else {
                                    if ((v29 + (v29 % (((v3 ^ v29)) | 1))) >= 42) {
                                        var v37 float64 = (float64(float32(-2.25)) + v8)
                                        v1 -= int8(v12)
                                        _ = v37
                                    } else {
                                        if ("foo bar" < "hello") {
                                            v8 = 1.5
                                        } else {
                                            v2 -= float64(((-2.25 / -2.25) + -2.25))
                                            v38 := GenS1{F0: 65535, F1: (v12 & (v12 + 120)), F2: 42}
                                            _ = v38
                                        }
                                    }
                                    var v39 uint16 = 641
                                    _ = v39
                                }
                            }
                            v40 := []int{(fi31 & fi28), (fi31 & (^(fi28)))}
                            _ = v40
                        }
                        switch "hello" {
                        case "foo bar":
                            switch (("foo bar" + ("hello" + ("\n\t" + ("hello" + "hello")))) + "hello") {
                            case "":
                                var v41 int8 = v1
                                _ = v41
                            case "foo bar":
                                v42 := GenS4{F0: (fi28 >> (((v12 & v12)) & 63)), F1: 556, F2: int32(v6)}
                                _ = v42
                            case "hello":
                                v43 := []int{}
                                _ = v43
                            }
                        case "hello":
                            fmt.Println(v3, v29)
                        case "":
                            v1 -= v1
                        }
                        _ = fi31
                    }
                    var v44 float64 = (v2 - 0.0)
                    switch (fi28 + (fi28 - fi28)) {
                    case 4294967295:
                        fmt.Println(v30)
                    case 302:
                        for fi45 := 0; fi45 < 17; fi45++ {
                            if false {
                                fmt.Println(v3)
                                v46 := []int{-158, -569, fi28, (fi45 % ((9223372036854775807) | 1))}
                                v46 = append(v46, 1)
                                _ = v46
                            } else {
                                var v47 float32 = 1e10
                                v48 := []int{fi28, (fi28 - (fi45 - (fi28 % (((-(fi28))) | 1)))), fi45}
                                var v49 int = (fi28 + (fi28 ^ (fi45 % ((fi28) | 1))))
                                fmt.Println(fi45, v8)
                                fmt.Println(v1, v1, v3)
                                _ = len(v48)
                                v12 += v12
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            switch (("" + ("\n\t" + "\n\t")) + "") {
                            case "":
                                if (int64(v7) == (v3 ^ v29)) {
                                    switch ("" + "") {
                                    case "a":
                                        var v50 uint32 = 781
                                        _ = v50
                                    case "hello":
                                        if true {
                                            if false {
                                                var v51 int = (^(fi45))
                                                var v52 int8 = (^(v1))
                                                v53 := GenS2{F0: 4294967295}
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            }
                                            fmt.Println(v2, v29)
                                            var v54 uint8 = v30
                                            _ = v54
                                        }
                                    case "foo bar":
                                        var v55 int16 = (v6 | v6)
                                        _ = v55
                                    }
                                    var v56 uint64 = 4294967295
                                    v57 := []int{fi28, int(v6), int(v3)}
                                    _ = v56
                                    _ = v57
                                } else {
                                    var v58 int = (fi28 - fi28)
                                    _ = v58
                                }
                            case "foo bar":
                                var v59 uint64 = 58
                                _ = v59
                            default:
                                var v60 bool = true
                                _ = v60
                            }
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v1, v6, v6)
                            case "":
                                var v61 uint64 = 0
                                _ = v61
                            case "a":
                                switch 128 {
                                case 42:
                                    fmt.Println(v29)
                                case 32767:
                                    switch (^(fi28)) {
                                    case 9223372036854775807:
                                        switch fi28 {
                                        case -9223372036854775808:
                                            v7 -= v7
                                        case 614:
                                            var v62 int8 = (v1 - (v1 | -128))
                                            _ = v62
                                        case 127:
                                            switch "foo bar" {
                                            case "a":
                                                v63 := GenS1{F0: 65535, F1: v12, F2: v30}
                                                _ = v63
                                            case "foo bar":
                                                v3 += (-(v29))
                                            }
                                        }
                                    case -1:
                                        var v64 string = "hello"
                                        _ = v64
                                    case 42:
                                        v29 += (^(v3))
                                    default:
                                        var v65 float32 = 1e10
                                        _ = v65
                                    }
                                }
                            }
                            _ = fi45
                        }
                    }
                    _ = v44
                }
                for fi66 := 0; fi66 < 18; fi66++ {
                    v67 := []int{}
                    _ = fi66
                    _ = v67
                }
            }
            var v68 uint64 = uint64(v30)
            _ = v30
            _ = v68
        } else {
            var v69 bool = true
            _ = v69
        }
        for fi70 := 0; fi70 < 2; fi70++ {
            if (!((fi70 & 4294967295) != 102)) {
                if (0 >= 738) {
                    if false {
                        v71 := []int{}
                        var v72 int8 = v1
                        if (!true) {
                            if (-1 != (v6 % ((v6) | 1))) {
                                switch (("" + "") + ("hello" + "hello")) {
                                case "foo bar":
                                    var v73 int16 = v7
                                    _ = v73
                                case "hello":
                                    var v74 string = ("a" + "hello")
                                    _ = v74
                                default:
                                    v75 := GenS5{F0: false}
                                    _ = v75
                                }
                            } else {
                                v12 = (^(v12))
                            }
                        }
                        _ = v71
                        _ = v72
                    }
                    var v76 uint32 = 2147483648
                    _ = v76
                }
                var v77 int = (fi28 * -9223372036854775808)
                _ = v77
            }
            if (true || true) {
                fmt.Println(v3)
            } else {
                var v78 int32 = 42
                if true {
                    var v79 uint32 = 65535
                    var v80 uint8 = 1
                    fmt.Println(v80)
                    _ = v79
                    _ = v80
                }
                fmt.Println(v3, v1)
                var v81 int64 = v29
                if true {
                    switch "a" {
                    case "\n\t":
                        v82 := []int{}
                        _ = v82
                    default:
                        fmt.Println(v8, fi70)
                    }
                }
                _ = v78
                _ = v81
            }
            var v83 float64 = v8
            _ = fi70
            _ = v83
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
