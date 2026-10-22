package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint16 }
func genFn0() bool {
    v5 := []int{4294967295, -9223372036854775808, -115, -2147483648}
    var v6 uint = 65535
    v5 = append(v5, 646)
    if (false && (128 < 2)) {
        v7 := GenS1{F0: 128, F1: 0}
        _ = len(v5)
        _ = v7
    } else {
        v6 += 42
        fmt.Println(v6)
        var v8 uint16 = 0
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = (v6 - v6)
        if false {
            if false {
                if true {
                    _ = len(v5)
                } else {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        switch 32768 {
                        case 2147483647:
                            var v12 string = "hello"
                            _ = v12
                        case -864:
                            var v13 int16 = 1000
                            _ = v13
                        default:
                            v6 = 42
                        }
                        _ = fi11
                    }
                    v14 := GenS2{F0: true, F1: 127}
                    _ = v14
                }
                var v15 float64 = (0.0 * ((0.0 + 0.0) + 1e10))
                _ = v15
            } else {
                switch "" {
                case "hello":
                    _ = len(v5)
                case "":
                    var v16 int64 = -377
                    _ = v16
                case "foo bar":
                    var v17 string = "a"
                    _ = v17
                default:
                    var v18 int8 = 100
                    _ = v18
                }
            }
            fmt.Println(fi9, fi9, v6)
            switch ("" + "foo bar") {
            case "\n\t":
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 bool = true
                    _ = fi19
                    _ = v20
                }
            case "a":
                if len(v5) > 0 {
                    v5[0] = fi9
                }
            default:
                if true {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        var v22 int = (fi21 >> (((v6 + (v6 >> ((0) & 63)))) & 63))
                        v23 := GenS1{F0: 127, F1: -1}
                        if len(v5) > 0 {
                            v5[0] = -658
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
        } else {
            fmt.Println(v10, fi9)
        }
        v5 = append(v5, 32767)
        _ = fi9
        _ = v10
    }
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint8) int {
    v24 := []int{-9223372036854775808, 668, 1}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    p0 -= (^(p0))
    p0 += 2
    var v25 int8 = 1
    _ = p0
    _ = v24
    _ = v25
    return 32768
}

func main() {
    var v1 uint64 = 176
    var v2 uint16 = 42
    var v3 int64 = 943
    v4 := make(map[uint32]bool)
    var v26 int32 = 0
    v26 += -2147483648
    for fi27 := 0; fi27 < 13; fi27++ {
        switch fi27 {
        case 1:
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int8 = 100
                _ = fi28
                _ = v29
            }
        case 128:
            v30 := GenS3{F0: v3, F1: 128}
            _ = v30
        default:
            var v31 float32 = 0.0
            _ = v31
        }
        _ = fi27
    }
    var v32 float32 = 0.1
    switch ("\n\t" + "") {
    case "":
        v33 := GenS3{F0: v3, F1: -1}
        _ = v33
    case "\n\t":
        v2 += (v2 << ((255) & 63))
    }
    v34 := GenS2{F0: false, F1: 42}
    v35 := []int{-2147483648}
    if true {
        var v36 uint16 = 32767
        switch "foo bar" {
        case "a":
            v37 := GenS4{F0: 0}
            _ = v37
        case "":
            v38 := GenS5{F0: "hello", F1: v2}
            _ = v38
        case "foo bar":
            if (1e10 <= float64((v32 * v32))) {
                switch -289 {
                case -1:
                    fmt.Println(v26, v36)
                case -9223372036854775808:
                    fmt.Println(v2)
                default:
                    v39 := GenS1{F0: uint8(v36), F1: 1}
                    _ = v39
                }
            }
        }
        v36 = v36
        _ = v36
    } else {
        var v40 uint8 = 128
        v1 -= v1
        var v41 int32 = (^(v26))
        _ = v40
        _ = v41
    }
    var v42 int32 = 400
    for fi43 := 0; fi43 < 18; fi43++ {
        for fi44 := 0; fi44 < 12; fi44++ {
            fmt.Println(v2)
            _ = fi44
        }
        _ = fi43
    }
    var v45 uint32 = 560
    for fi46 := 0; fi46 < 17; fi46++ {
        v42 = (^(v42))
        if true {
            if ((false && (!true)) || (!true)) {
                delete(v4, uint32(v1))
                switch "foo bar" {
                case "\n\t":
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v1 = 4294967295
                        v4[42] = true
                        v4[v45] = (!true)
                        _ = fi47
                    }
                case "foo bar":
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                case "hello":
                    var v48 int = fi46
                    _ = v48
                }
                if (true && (false && false)) {
                    switch (fi46 | (fi46 % ((42) | 1))) {
                    case 32767:
                        v45 = v45
                    case -592:
                        v3 += (v3 << ((960) & 63))
                    case 715:
                        if true {
                            if true {
                                if false {
                                    v3 = (-(v3))
                                    var v49 bool = (1 > (v26 % ((v26) | 1)))
                                    _ = v49
                                } else {
                                    v45 += 0
                                    var v50 uint16 = (v2 + (v2 % ((0) | 1)))
                                    _ = v50
                                }
                                v51 := GenS5{F0: (("a" + "\n\t") + "foo bar"), F1: (v2 / ((v2) | 1))}
                                v1 -= 42
                                _ = v51
                            }
                            v1 = 395
                            if false {
                                switch (-(fi46)) {
                                case 127:
                                    var v52 uint8 = 0
                                    _ = v52
                                default:
                                    _ = len(v4)
                                }
                            }
                        } else {
                            if false {
                                var v53 float32 = 3.14159
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 5; fi54++ {
                                switch -2147483648 {
                                case -2147483648:
                                    switch "a" {
                                    case "hello":
                                        v3 = v3
                                    }
                                case 128:
                                    var v55 string = ""
                                    _ = v55
                                case 4294967295:
                                    if (!false) {
                                        switch "foo bar" {
                                        case "a":
                                            v56 := GenS6{F0: v1, F1: ""}
                                            _ = v56
                                        case "":
                                            v57 := GenS4{F0: (v45 - (v45 % ((v45) | 1)))}
                                            _ = v57
                                        }
                                        if (255 >= uint8(v45)) {
                                            v3 = 9223372036854775807
                                            v35 = append(v35, (fi54 - int(v2)))
                                            switch "" {
                                            case "":
                                                var v58 uint = 354
                                                _ = v58
                                            case "\n\t":
                                                switch (("a" + "a") + (("foo bar" + (("\n\t" + ("foo bar" + ("a" + "a"))) + (("a" + (("\n\t" + "") + "\n\t")) + "hello"))) + "\n\t")) {
                                                case "a":
                                                    v45 += v45
                                                case "hello":
                                                    v2 = (v2 >> ((4294967295) & 63))
                                                default:
                                                    var v59 uint64 = 2
                                                    _ = v59
                                                }
                                            }
                                        }
                                    }
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (65535 == 0) {
                                            if true {
                                                var v60 string = "hello"
                                                _ = v60
                                            } else {
                                                var v61 uint8 = 42
                                                v62 := GenS7{F0: v2}
                                                if false {
                                                    fmt.Println(fi46)
                                                    switch (fi54 << ((314) & 63)) {
                                                    case 42:
                                                        var v63 int64 = 42
                                                        _ = v63
                                                    case 128:
                                                        v45 -= (v45 % (((v45 / ((v45) | 1))) | 1))
                                                    case 852:
                                                        switch (fi54 >> ((42) & 63)) {
                                                        case 9223372036854775807:
                                                            if (!(float32(0.0) >= v32)) {
                                                                fmt.Println(fi46)
                                                            } else {
                                                                if false {
                                                                    if (!(3.14159 <= float64(float32(1.5)))) {
                                                                        v2 -= (v2 / (((^(v2))) | 1))
                                                                        if (false && false) {
                                                                            fmt.Println(v61)
                                                                            v42 = v42
                                                                        } else {
                                                                            if len(v35) >= 2 {
                                                                                v35 = v35[1:len(v35)]
                                                                            }
                                                                            fmt.Println(v42, v45, v42)
                                                                            v1 = v1
                                                                        }
                                                                        v32 = v32
                                                                    } else {
                                                                        if false {
                                                                            v4[v45] = ("foo bar" != ("foo bar" + "foo bar"))
                                                                            v64 := GenS5{F0: "foo bar", F1: v2}
                                                                            v35 = append(v35, 1)
                                                                            var v65 uint = 876
                                                                            switch "" {
                                                                            case "":
                                                                                if len(v35) > 0 {
                                                                                    v35[0] = fi54
                                                                                }
                                                                            case "foo bar":
                                                                                if ((^(v61)) < 2) {
                                                                                    v3 += 0
                                                                                } else {
                                                                                    v45 = v45
                                                                                    var v66 uint8 = v61
                                                                                    var v67 uint8 = 2
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                            case "hello":
                                                                                if (!(true && false)) {
                                                                                    v32 = v32
                                                                                    var v68 uint8 = 101
                                                                                    _ = v68
                                                                                } else {
                                                                                    if (0 > -875) {
                                                                                        v2 = 0
                                                                                        var v69 uint8 = (^(v61))
                                                                                        switch (fi54 * int(v3)) {
                                                                                        case 1:
                                                                                            var v70 uint64 = (v1 * (v1 & v1))
                                                                                            _ = v70
                                                                                        case -712:
                                                                                            if ((!true) || false) {
                                                                                                var v71 int = (fi46 ^ int(v61))
                                                                                                var v72 uint64 = v1
                                                                                                if true {
                                                                                                    if true {
                                                                                                        delete(v4, (v45 / ((190) | 1)))
                                                                                                    }
                                                                                                    var v73 float64 = 3.14159
                                                                                                    _ = v73
                                                                                                } else {
                                                                                                    var v74 uint = 681
                                                                                                    v65 -= 244
                                                                                                    _ = v74
                                                                                                }
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            } else {
                                                                                                fmt.Println(v65, fi46, v61)
                                                                                                var v75 uint32 = (v45 - v45)
                                                                                                _ = v75
                                                                                            }
                                                                                        default:
                                                                                            var v76 uint64 = uint64(v65)
                                                                                            _ = v76
                                                                                        }
                                                                                        _ = v69
                                                                                    } else {
                                                                                        var v77 uint64 = uint64(v65)
                                                                                        var v78 uint = (v65 ^ uint(v77))
                                                                                        fmt.Println(v3, v45)
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v35 = append(v35, (fi54 / (((fi46 & (fi46 + (^(fi46))))) | 1)))
                                                                            }
                                                                            v79 := GenS8{F0: "foo bar"}
                                                                            if true {
                                                                                delete(v4, (v45 & 2147483648))
                                                                            }
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v79
                                                                        }
                                                                        var v80 float32 = ((1.5 * float32(1.5)) - v32)
                                                                        _ = v80
                                                                    }
                                                                    v32 = (((v32 / v32) + ((v32 - (float32(float64(v32)) + (v32 * 3.14159))) / ((1e10 * -2.25) * (float32(-2.25) - 1.5)))) * v32)
                                                                } else {
                                                                    v81 := GenS9{F0: true, F1: (v2 & (v2 % (((^(v2))) | 1)))}
                                                                    var v82 uint = 915
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            }
                                                            fmt.Println(v32, v2, fi46)
                                                            var v83 int8 = 100
                                                            var v84 uint16 = v2
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    }
                                                    var v85 uint8 = v61
                                                    v86 := v4[65535]
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    fmt.Println(v26, fi54, v42)
                                                }
                                                if (!(!false)) {
                                                    v87 := v4[(v45 << ((2) & 63))]
                                                    fmt.Println(v2, v61)
                                                    _ = v87
                                                } else {
                                                    var v88 uint64 = (v1 % (((v1 >> ((2) & 63))) | 1))
                                                    switch fi46 {
                                                    case 548:
                                                        fmt.Println(v61, v32, v32)
                                                    default:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if true {
                                                                v42 -= 0
                                                                delete(v4, (^(v45)))
                                                            } else {
                                                                var v89 uint8 = 2
                                                                _ = v89
                                                            }
                                                        case "":
                                                            fmt.Println(v32, v45)
                                                            var v90 uint16 = v2
                                                            v91 := GenS2{F0: ((v45 | (v45 ^ v45)) <= v45), F1: 1}
                                                            _ = v90
                                                            _ = v91
                                                        default:
                                                            fmt.Println(v3)
                                                        }
                                                    }
                                                    _ = v88
                                                }
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                    case "a":
                                        if false {
                                            v92 := GenS3{F0: -1, F1: 547}
                                            _ = v92
                                        }
                                    case "foo bar":
                                        if ("" != "a") {
                                            v93 := GenS8{F0: ("a" + "foo bar")}
                                            v94 := GenS5{F0: "\n\t", F1: v2}
                                            _ = v93
                                            _ = v94
                                        } else {
                                            v35 = append(v35, genFn1(127))
                                        }
                                    default:
                                        v45 += (^(v45))
                                    }
                                }
                                var v95 int64 = 9223372036854775807
                                _ = fi54
                                _ = v95
                            }
                            for fi96 := 0; fi96 < 11; fi96++ {
                                var v97 int64 = 947
                                _ = fi96
                                _ = v97
                            }
                        }
                    default:
                        v98 := GenS1{F0: 1, F1: 270}
                        _ = v98
                    }
                    _ = len(v35)
                    var v99 bool = (false && false)
                    _ = v99
                }
            }
        }
        for fi100 := 0; fi100 < 13; fi100++ {
            v1 = 2
            var v101 bool = false
            v2 += v2
            _ = fi100
            _ = v101
        }
        _ = fi46
    }
    v32 = float32(0.1)
    fmt.Println(v3, v45, v42)
    v102 := GenS8{F0: ("hello" + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v102
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v45
}
