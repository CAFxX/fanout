package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS11 struct { F0 int; F1 bool; F2 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 bool; F1 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 16; fi4++ {
        fmt.Println(fi4)
        fmt.Println(fi4)
        var v5 int8 = 100
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 1
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint32 = 42
        var v10 string = ("" + "foo bar")
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint64 = 2
    _ = v11
    return 0
}
func genFn1(p0 int) int64 {
    v12 := GenS1{F0: 9223372036854775807}
    if (true && true) {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int64 = -89
            v15 := []int{(p0 ^ (-(fi13)))}
            if false {
                var v16 uint64 = 0
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v16
            } else {
                var v17 float32 = float32(3.14159)
                var v18 int64 = v14
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(p0)
        v19 := GenS1{F0: 0}
        _ = v19
    } else {
        var v20 uint8 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS2{F0: -128}
        var v23 float64 = 0.0
        if (true || false) {
            var v24 int8 = int8(fi21)
            var v25 int8 = 127
            _ = v24
            _ = v25
        } else {
            var v26 uint32 = 590
            fmt.Println(p0)
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := []int{fi27}
                _ = len(v28)
                v29 := GenS1{F0: 9223372036854775807}
                if (-1 == 42) {
                    v30 := GenS1{F0: -9223372036854775808}
                    v31 := GenS2{F0: -1}
                    switch p0 {
                    case 128:
                        v32 := GenS3{F0: -75}
                        v26 = v26
                        _ = v32
                    case 1:
                        var v33 bool = ((fi27 >> ((genFn0()) & 63)) > fi21)
                        _ = v33
                    case 127:
                        p0 += p0
                        v34 := GenS1{F0: 0}
                        var v35 float64 = (v23 - -2.25)
                        _ = v34
                        _ = v35
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(p0)
    if true {
        p0 = p0
        var v36 string = ("\n\t" + "hello")
        if (!(!(!(1e10 == (((float64(float32((-2.25 / 1.5))) - (1.5 / 1e10)) * ((3.14159 * 0.0) + ((-2.25 * 0.0) * 1e10))) - float64(1.5)))))) {
            if (false || (9223372036854775807 != (p0 % ((p0) | 1)))) {
                var v37 int16 = -803
                if (false && ((^(v37)) >= (^(v37)))) {
                    var v38 uint8 = 1
                    v39 := GenS4{F0: 42, F1: 0, F2: 50}
                    switch (v36 + v36) {
                    case "hello":
                        v40 := []int{(p0 + 32768)}
                        _ = v40
                    }
                    _ = v38
                    _ = v39
                }
                v36 = "foo bar"
                _ = v37
            }
            fmt.Println(p0)
        }
        _ = v36
    }
    _ = p0
    _ = v12
    return 829
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = 0.1
    var v3 uint64 = 4294967295
    v1 = v1
    v41 := []int{-9223372036854775808, 0, -1}
    if len(v41) > 0 {
        v41[0] = 42
    }
    v42 := GenS3{F0: -10}
    if true {
        var v43 int32 = -266
        _ = v43
    } else {
        if true {
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 float32 = ((v2 * (0.0 / v2)) * v2)
                if true {
                    v46 := GenS4{F0: 1000, F1: 541, F2: 450}
                    for fi47 := 0; fi47 < 12; fi47++ {
                        if false {
                            if len(v41) >= 2 {
                                v41 = v41[1:len(v41)]
                            }
                        } else {
                            var v48 int32 = -47
                            _ = v48
                        }
                        _ = len(v41)
                        _ = fi47
                    }
                    fmt.Println(v2)
                    _ = v46
                } else {
                    var v49 uint8 = (v1 << ((2) & 63))
                    var v50 int8 = int8(v1)
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    _ = v49
                    _ = v50
                }
                v3 = uint64(v1)
                _ = fi44
                _ = v45
            }
            fmt.Println(v2, v3, v2)
        } else {
            if false {
                if false {
                    switch ("\n\t" + "a") {
                    case "":
                        fmt.Println(v3, v1)
                    }
                } else {
                    v3 += (^(v3))
                    if (true && true) {
                        if len(v41) >= 2 {
                            v41 = v41[1:len(v41)]
                        }
                    } else {
                        v51 := GenS1{F0: 42}
                        v2 = v2
                        var v52 int8 = 127
                        _ = v51
                        _ = v52
                    }
                }
                if ((v3 == v3) || false) {
                    fmt.Println(v2)
                    if false {
                        v41 = append(v41, -476)
                        if false {
                            if (true && true) {
                                switch -2147483648 {
                                case 0:
                                    v2 -= v2
                                case 33:
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                case 1:
                                    v53 := GenS2{F0: 0}
                                    _ = v53
                                }
                                if len(v41) > 0 {
                                    v41[0] = 32768
                                }
                                for fi54 := 0; fi54 < 9; fi54++ {
                                    var v55 uint8 = (v1 * v1)
                                    v3 += 1
                                    _ = fi54
                                    _ = v55
                                }
                            }
                            switch ((("a" + "hello") + "\n\t") + "\n\t") {
                            case "\n\t":
                                v41 = append(v41, 127)
                            case "hello":
                                var v56 uint = 778
                                _ = v56
                            case "a":
                                v3 += 1
                            default:
                                for fi57 := 0; fi57 < 11; fi57++ {
                                    for fi58 := 0; fi58 < 1; fi58++ {
                                        v59 := GenS5{F0: fi57}
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch (^(fi57)) {
                                            case 0:
                                                v60 := GenS6{F0: 787, F1: (^(v1))}
                                                _ = v60
                                            case 42:
                                                if ((556 < genFn0()) || (!(false || (true || false)))) {
                                                    if len(v41) >= 2 {
                                                        v41 = v41[1:len(v41)]
                                                    }
                                                    var v61 float64 = 1e10
                                                    _ = v61
                                                } else {
                                                    v62 := GenS7{F0: v3, F1: 0}
                                                    fmt.Println(v1, fi57, fi57)
                                                    _ = v62
                                                }
                                                if (0.0 <= 1.5) {
                                                    var v63 int64 = 1
                                                    v41 = append(v41, fi58)
                                                    v64 := GenS8{F0: int8(v63), F1: 0}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                if len(v41) > 0 {
                                                    v41[0] = fi57
                                                }
                                                var v65 float32 = (0.1 * float32(float64(float32((3.14159 - (0.0 + (3.14159 * 3.14159)))))))
                                                var v66 string = ""
                                                v2 += (0.0 * v2)
                                                var v67 uint32 = 344
                                                v68 := GenS5{F0: fi57}
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            case -426:
                                                var v69 uint32 = 4294967295
                                                var v70 int32 = -1
                                                _ = v69
                                                _ = v70
                                            default:
                                                var v71 uint16 = 255
                                                _ = v71
                                            }
                                        case "a":
                                            if ((true && false) && true) {
                                                var v72 bool = true
                                                if true {
                                                    var v73 int = (^(fi58))
                                                    var v74 float32 = 1e10
                                                    _ = v73
                                                    _ = v74
                                                }
                                                if false {
                                                    var v75 uint16 = 1
                                                    _ = v75
                                                } else {
                                                    v76 := GenS7{F0: (v3 >> ((1) & 63)), F1: 42}
                                                    _ = v76
                                                }
                                                _ = v72
                                            }
                                            v77 := GenS4{F0: 501, F1: uint(v3), F2: 253}
                                            _ = v77
                                        case "hello":
                                            var v78 int64 = -1
                                            _ = v78
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                    v79 := GenS8{F0: 127, F1: 2147483647}
                                    _ = fi57
                                    _ = v79
                                }
                            }
                        }
                        v80 := GenS9{F0: -197, F1: "a"}
                        _ = v80
                    } else {
                        var v81 uint32 = 4294967295
                        v41 = append(v41, 0)
                        v82 := GenS2{F0: 100}
                        _ = v81
                        _ = v82
                    }
                    v1 = v1
                } else {
                    var v83 uint64 = v3
                    v41 = append(v41, -2147483648)
                    _ = v83
                }
            } else {
                v3 = 294
            }
            if (3.14159 <= 1.5) {
                var v84 uint = 18446744073709551615
                var v85 uint32 = 360
                for fi86 := 0; fi86 < 19; fi86++ {
                    if true {
                        fmt.Println(fi86, v1)
                        var v87 int = 42
                        _ = v87
                    }
                    _ = fi86
                }
                _ = v84
                _ = v85
            } else {
                fmt.Println(v2, v1)
                fmt.Println(v2, v2)
            }
        }
        if false {
            v88 := GenS10{F0: 42, F1: -128, F2: v1}
            var v89 uint32 = 2
            _ = v88
            _ = v89
        }
    }
    if true {
        fmt.Println(v1)
        v90 := GenS4{F0: -32768, F1: 18446744073709551615, F2: 65535}
        _ = v90
    }
    v91 := GenS11{F0: 127, F1: (((-294 == -951) || false) && (218 < (v3 - 1))), F2: (119 < -1)}
    fmt.Println(v1, v1)
    if true {
        v1 = (v1 * 42)
        for fi92 := 0; fi92 < 9; fi92++ {
            if false {
                var v93 uint32 = uint32(v1)
                v1 -= (^(v1))
                v41 = append(v41, fi92)
                _ = v93
            }
            switch (^(fi92)) {
            case 32768:
                for fi94 := 0; fi94 < 13; fi94++ {
                    if (!(1 <= -32768)) {
                        fmt.Println(v1, fi92, fi92)
                        var v95 float64 = ((3.14159 - ((1e10 * (1.5 / -2.25)) / 3.14159)) * 1e10)
                        _ = v95
                    } else {
                        v1 += v1
                    }
                    v3 = v3
                    switch fi94 {
                    case 4294967295:
                        v96 := GenS12{F0: fi92}
                        _ = v96
                    }
                    if true {
                        v2 += 0.1
                        var v97 uint16 = 640
                        _ = v97
                    }
                    _ = fi94
                }
            case 9223372036854775807:
                var v98 uint = 1
                _ = v98
            }
            v2 = (v2 + (-2.25 * v2))
            _ = fi92
        }
        v2 = (3.14159 / v2)
    }
    for fi99 := 0; fi99 < 9; fi99++ {
        _ = len(v41)
        var v100 int32 = 1
        for fi101 := 0; fi101 < 16; fi101++ {
            if false {
                if (false && false) {
                    v100 = (v100 % (((v100 + 378)) | 1))
                    var v102 uint16 = 0
                    switch 128 {
                    case 883:
                        var v103 uint32 = 737
                        _ = v103
                    case 128:
                        if true {
                            _ = len(v41)
                            v104 := GenS13{F0: (uint(v3) <= 385), F1: false}
                            fmt.Println(v3, fi99)
                            _ = v104
                        }
                    }
                    v100 += (^(v100))
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    _ = v102
                } else {
                    if (false || (123 <= 0)) {
                        var v105 string = "foo bar"
                        var v106 uint8 = v1
                        _ = v105
                        _ = v106
                    }
                    var v107 int8 = 127
                    v1 = v1
                    _ = v107
                }
            } else {
                var v108 float64 = (3.14159 * (1.5 + -2.25))
                _ = v108
            }
            switch "\n\t" {
            case "foo bar":
                var v109 bool = true
                _ = v109
            }
            _ = fi101
        }
        _ = fi99
        _ = v100
    }
    if true {
        var v110 uint32 = 105
        _ = v110
    } else {
        if len(v41) > 0 {
            v41[0] = 128
        }
    }
    _ = len(v41)
    if len(v41) > 0 {
        v41[0] = 42
    }
    for fi111 := 0; fi111 < 15; fi111++ {
        switch "a" {
        case "hello":
            var v112 int8 = 100
            _ = v112
        default:
            if len(v41) > 0 {
                v41[0] = 1
            }
        }
        var v113 uint = 42
        _ = fi111
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v91
}
