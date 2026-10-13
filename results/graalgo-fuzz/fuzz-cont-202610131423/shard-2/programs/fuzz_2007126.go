package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 bool; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 bool, p1 int16, p2 uint32) uint32 {
    p2 = p2
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 114, F1: 128, F2: -9223372036854775808}
        v6 := []int{-267, 9223372036854775807, 42, fi4}
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return (p2 & 4294967295)
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 4294967295
    var v3 uint64 = 42
    if true {
        v7 := []int{911}
        v2 = (v2 / (((v2 + 78)) | 1))
        _ = v7
    } else {
        var v8 uint16 = 32767
        _ = v8
    }
    switch -340 {
    case -721:
        if ((((0.0 * -2.25) - 1.5) + (3.14159 / (1e10 + (0.1 + (((-2.25 * -2.25) - -2.25) + 0.1))))) <= -2.25) {
            v2 = 940
            for fi9 := 0; fi9 < 19; fi9++ {
                v3 = 2
                if false {
                    if (false && ((!(((0.1 * 0.1) * (float32(float64(1.5)) / (-2.25 - ((1.5 + 3.14159) + 0.0)))) != (0.1 / -2.25))) || (float32(3.14159) <= 0.1))) {
                        if false {
                            if (!true) {
                                if true {
                                    if ("foo bar" <= ("hello" + "a")) {
                                        v3 = v3
                                    } else {
                                        fmt.Println(v3, v3)
                                    }
                                } else {
                                    if true {
                                        v1 = (-(v1))
                                    } else {
                                        v2 += (v2 * 0)
                                        var v10 int = fi9
                                        _ = v10
                                    }
                                    v11 := GenS1{F0: 2, F1: 128, F2: 42}
                                    _ = v11
                                }
                                var v12 uint16 = 42
                                var v13 int32 = 42
                                _ = v12
                                _ = v13
                            }
                            fmt.Println(fi9)
                            var v14 float32 = (float32(-2.25) / 0.1)
                            _ = v14
                        } else {
                            v1 -= (^(v1))
                            if true {
                                var v15 uint16 = 827
                                if ((!(!(v2 != (v2 + (v2 | v2))))) && (!true)) {
                                    v2 = v2
                                    var v16 string = ("hello" + ("" + "a"))
                                    _ = v16
                                } else {
                                    v17 := GenS2{F0: 251}
                                    switch "" {
                                    case "":
                                        v18 := []int{}
                                        _ = v18
                                    }
                                    _ = v17
                                }
                                var v19 int16 = v1
                                _ = v15
                                _ = v19
                            } else {
                                if true {
                                    fmt.Println(v1, v2, fi9)
                                    fmt.Println(fi9, v2, v2)
                                } else {
                                    var v20 uint16 = 42
                                    v21 := []int{}
                                    _ = v20
                                    _ = v21
                                }
                                var v22 int16 = v1
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    v24 := []int{fi23}
                                    switch "hello" {
                                    case "a":
                                        if false {
                                            if len(v24) > 0 {
                                                v24[0] = -461
                                            }
                                        }
                                    case "foo bar":
                                        switch fi9 {
                                        case 877:
                                            var v25 int64 = 9223372036854775807
                                            _ = v25
                                        case 713:
                                            v1 = v22
                                        }
                                    default:
                                        var v26 uint = 693
                                        _ = v26
                                    }
                                    var v27 int64 = -430
                                    switch "hello" {
                                    case "":
                                        if len(v24) >= 2 {
                                            v24 = v24[1:len(v24)]
                                        }
                                    case "\n\t":
                                        v28 := GenS2{F0: -1}
                                        _ = v28
                                    case "a":
                                        if (false || (3.14159 > -2.25)) {
                                            v27 -= v27
                                            var v29 uint8 = 255
                                            _ = v29
                                        }
                                    default:
                                        if true {
                                            fmt.Println(fi9, v22)
                                            v24 = append(v24, fi9)
                                            v1 += (v1 & v1)
                                        }
                                    }
                                    fmt.Println(v2, fi23)
                                    if (((-15 != -1) && (557 > (^(v2)))) && ((^(v2)) > v2)) {
                                        v1 = (^(v1))
                                        var v30 int16 = (v22 >> (((v2 - (v2 & (v2 % (((v2 - 4294967295)) | 1))))) & 63))
                                        var v31 uint16 = uint16(v2)
                                        _ = v30
                                        _ = v31
                                    }
                                    v22 = 1000
                                    if (!true) {
                                        if (0.0 < float64(0.0)) {
                                            v32 := GenS1{F0: 128, F1: 127, F2: int64(v1)}
                                            v27 = v27
                                            _ = v32
                                        } else {
                                            switch fi9 {
                                            case 565:
                                                v27 += (-(v27))
                                            default:
                                                if (int32(v2) == 2147483647) {
                                                    switch "" {
                                                    case "":
                                                        var v33 uint16 = 32767
                                                        _ = v33
                                                    case "\n\t":
                                                        if false {
                                                            var v34 uint8 = 42
                                                            if false {
                                                                var v35 string = "\n\t"
                                                                if false {
                                                                    var v36 int = (fi9 >> ((v2) & 63))
                                                                    _ = v36
                                                                }
                                                                _ = v35
                                                            } else {
                                                                fmt.Println(v34, fi23)
                                                                v37 := GenS3{F0: 378}
                                                                _ = v37
                                                            }
                                                            var v38 int8 = int8(v27)
                                                            if (0 == 42) {
                                                                v2 = (v2 ^ (v2 + (^(v2))))
                                                            }
                                                            _ = v34
                                                            _ = v38
                                                        }
                                                    }
                                                } else {
                                                    var v39 float64 = 1e10
                                                    v40 := GenS2{F0: -1}
                                                    var v41 uint64 = v3
                                                    v42 := GenS4{F0: ((v22 / ((1) | 1)) > 92), F1: true, F2: 42}
                                                    var v43 int16 = v1
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                switch fi23 {
                                                case -9223372036854775808:
                                                    switch 912 {
                                                    case 42:
                                                        var v44 uint8 = 255
                                                        _ = v44
                                                    default:
                                                        v3 = (v3 | 4294967295)
                                                    }
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                    fmt.Println(fi23)
                                                    switch (fi23 * (-(fi23))) {
                                                    case 127:
                                                        fmt.Println(v27)
                                                    }
                                                    if ((false || false) && (76 > 0)) {
                                                        var v45 uint64 = 4294967295
                                                        _ = v45
                                                    }
                                                case 1:
                                                    var v46 string = "hello"
                                                    _ = v46
                                                }
                                                var v47 string = "\n\t"
                                                _ = v47
                                            }
                                        }
                                        switch (fi9 / ((fi23) | 1)) {
                                        case 9223372036854775807:
                                            v27 -= v27
                                        case 128:
                                            fmt.Println(v1)
                                        }
                                    }
                                    _ = fi23
                                    _ = v24
                                    _ = v27
                                }
                                _ = v22
                            }
                        }
                        var v48 float32 = ((3.14159 + 1e10) * 1e10)
                        _ = v48
                    }
                    var v49 uint = v2
                    _ = v49
                }
                var v50 uint32 = 2147483648
                _ = fi9
                _ = v50
            }
        }
    }
    if (1e10 >= float64(3.14159)) {
        v1 = -473
        v3 = v3
    } else {
        fmt.Println(v1)
        if true {
            var v51 uint16 = 2
            _ = v51
        } else {
            switch (("a" + "") + (("hello" + "foo bar") + "\n\t")) {
            case "a":
                for fi52 := 0; fi52 < 1; fi52++ {
                    for fi53 := 0; fi53 < 16; fi53++ {
                        if false {
                            switch ((("a" + "a") + "") + "") {
                            case "a":
                                v54 := []int{fi52, fi53, (fi53 - (fi52 << ((v2) & 63))), fi53}
                                _ = v54
                            default:
                                var v55 string = ""
                                _ = v55
                            }
                            v1 = (-(v1))
                            var v56 int32 = -203
                            _ = v56
                        } else {
                            v1 = 748
                            v57 := []int{-984, fi52}
                            _ = v57
                        }
                        switch int(v1) {
                        case 127:
                            v58 := GenS1{F0: 202, F1: 128, F2: 296}
                            _ = v58
                        }
                        v59 := GenS5{F0: uint16(v2)}
                        v60 := GenS1{F0: 238, F1: 128, F2: 1}
                        v61 := []int{fi53}
                        switch "foo bar" {
                        case "a":
                            v3 -= v3
                        case "":
                            switch "hello" {
                            case "":
                                v2 += 965
                            case "foo bar":
                                switch "hello" {
                                case "hello":
                                    switch "\n\t" {
                                    case "hello":
                                        var v62 int16 = 32767
                                        _ = v62
                                    }
                                default:
                                    v63 := GenS3{F0: 42}
                                    _ = v63
                                }
                            case "a":
                                var v64 uint = v2
                                _ = v64
                            }
                        case "\n\t":
                            if (!true) {
                                var v65 int32 = -1
                                _ = v65
                            } else {
                                var v66 uint8 = 221
                                var v67 int8 = 100
                                fmt.Println(v1)
                                var v68 uint = v2
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                        default:
                            if len(v61) > 0 {
                                v61[0] = fi52
                            }
                            v2 = v2
                        }
                        var v69 int16 = (^(v1))
                        _ = fi53
                        _ = v59
                        _ = v60
                        _ = v61
                        _ = v69
                    }
                    v1 = (v1 * (^(v1)))
                    _ = fi52
                }
            case "foo bar":
                v3 -= (^(v3))
            }
            v70 := GenS6{F0: 32767, F1: true, F2: 0}
            v3 = (v3 >> ((4294967295) & 63))
            _ = v70
        }
        v3 += v3
    }
    if ((1e10 - float32(1e10)) >= (0.0 / 1.0)) {
        var v71 int64 = 42
        v3 = (v3 - (v3 % ((v3) | 1)))
        _ = v71
    } else {
        if (!false) {
            v1 -= v1
        }
    }
    v72 := []int{433, -1}
    v72 = append(v72, 32768)
    var v73 int = 32768
    if (2147483647 != 753) {
        for fi74 := 0; fi74 < 6; fi74++ {
            if ((uint(v3) != (^(v2))) || true) {
                v3 += v3
            } else {
                for fi75 := 0; fi75 < 20; fi75++ {
                    var v76 float64 = 0.0
                    var v77 uint = 1
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                v1 = v1
            }
            var v78 int = -437
            fmt.Println(v78, v3)
            _ = fi74
            _ = v78
        }
        v72 = append(v72, v73)
        fmt.Println(v2)
    }
    v2 = v2
    if len(v72) > 0 {
        v72[0] = v73
    }
    fmt.Println(v73)
    for fi79 := 0; fi79 < 11; fi79++ {
        var v80 float32 = (1e10 + 0.0)
        _ = fi79
        _ = v80
    }
    if len(v72) >= 2 {
        v72 = v72[1:len(v72)]
    }
    var v81 uint16 = 256
    v82 := GenS3{F0: int32(v81)}
    for fi83 := 0; fi83 < 19; fi83++ {
        fmt.Println(v73)
        var v84 uint = 4294967295
        v85 := GenS7{F0: 9223372036854775807, F1: 460, F2: "foo bar"}
        _ = fi83
        _ = v84
        _ = v85
    }
    var v86 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    fmt.Println("v81:", v81)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
    _ = v73
    _ = v81
    _ = v82
    _ = v86
}
