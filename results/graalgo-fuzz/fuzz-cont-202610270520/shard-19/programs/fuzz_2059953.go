package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 int16; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint32; F1 string; F2 bool }
func genFn0(p0 uint16, p1 int8, p2 string) int {
    if true {
        v6 := []int{}
        _ = v6
    } else {
        fmt.Println(p1, p0)
        fmt.Println(p1, p0)
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int32 = 0
            switch fi7 {
            case 32768:
                p2 = "\n\t"
            default:
                if (!true) {
                    v9 := []int{fi7, -9223372036854775808, fi7, (fi7 * (fi7 << ((2) & 63)))}
                    for fi10 := 0; fi10 < 3; fi10++ {
                        p0 = (p0 << ((4294967295) & 63))
                        _ = fi10
                    }
                    _ = v9
                } else {
                    var v11 uint16 = (^(p0))
                    fmt.Println(p2, p0, p0)
                    _ = v11
                }
            }
            var v12 string = ""
            _ = fi7
            _ = v12
            _ = v8
        }
    }
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{("foo bar" + ""), "a"}
    v3 := []int{167, 361}
    var v4 string = ("\n\t" + "a")
    var v5 float32 = float32(1e10)
    v2 = append(v2, ("hello" + ("a" + v4)))
    var v13 string = v4
    var v14 uint32 = 4294967295
    for fi15 := 0; fi15 < 19; fi15++ {
        if (-1 <= 32) {
            for fi16 := 0; fi16 < 13; fi16++ {
                if (((0.0 - 3.14159) - 0.1) != (float64((((((v5 + (v5 - (float32(3.14159) - -2.25))) / 1e10) - (v5 + (-2.25 - 3.14159))) / 1.0) / v5)) / 1e10)) {
                    var v17 int32 = 0
                    fmt.Println(fi15)
                    _ = v17
                }
                var v18 int32 = -1
                switch (fi15 | 407) {
                case 2147483647:
                    delete(v1, v13)
                    v4 = v13
                    v4 = v4
                case 128:
                    if (198 <= 32767) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if false {
                            fmt.Println(fi15, v4, v4)
                        }
                        v19 := GenS1{F0: 1, F1: 4294967295}
                        _ = v19
                    } else {
                        fmt.Println(v18)
                        if len(v3) > 0 {
                            v3[0] = (fi15 | (fi15 ^ fi16))
                        }
                    }
                case 32768:
                    v20 := GenS2{F0: 32767, F1: 0, F2: 42}
                    var v21 uint32 = v14
                    _ = v20
                    _ = v21
                default:
                    delete(v1, v13)
                }
                switch "foo bar" {
                case "":
                    var v22 uint32 = 2
                    _ = v22
                }
                v18 += int32(fi15)
                _ = fi16
                _ = v18
            }
            if false {
                if (!true) {
                    v23 := GenS1{F0: 4294967295, F1: 776}
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v4 = v13
                        if (!(false || (42 != -1))) {
                            var v25 float32 = v5
                            if (42 <= 698) {
                                v14 = (v14 ^ (v14 | (v14 & v14)))
                            } else {
                                var v26 int8 = -128
                                _ = v26
                            }
                            switch ("hello" + ((v13 + v4) + "hello")) {
                            case "a":
                                var v27 int8 = -1
                                _ = v27
                            case "foo bar":
                                v28 := GenS1{F0: 0, F1: 0}
                                _ = v28
                            case "hello":
                                v29 := v1[v4]
                                _ = v29
                            default:
                                delete(v1, v4)
                            }
                            _ = v25
                        } else {
                            if (278 > 2147483647) {
                                v30 := GenS3{F0: (fi15 + fi15), F1: 181, F2: -1}
                                if false {
                                    fmt.Println(v13)
                                    v13 = (v13 + v4)
                                    if ((42 > v14) || (true && true)) {
                                        v4 = ((v4 + v13) + v13)
                                        var v31 string = ("hello" + "foo bar")
                                        _ = v31
                                    }
                                    v13 = "a"
                                    switch "foo bar" {
                                    case "":
                                        var v32 int = fi24
                                        _ = v32
                                    case "\n\t":
                                        v14 -= 0
                                    case "hello":
                                        switch (v13 + "a") {
                                        case "foo bar":
                                            if true {
                                                var v33 int = fi24
                                                _ = v33
                                            }
                                        case "a":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        v34 := GenS2{F0: 65535, F1: 32767, F2: 928}
                                        _ = v34
                                    }
                                    if true {
                                        v35 := v1[("foo bar" + v13)]
                                        _ = v35
                                    } else {
                                        switch (fi24 << ((858) & 63)) {
                                        case 32768:
                                            if ((false || (1e10 < -2.25)) && true) {
                                                if true {
                                                    if (3 >= 32) {
                                                        v13 = "hello"
                                                        var v36 uint32 = v14
                                                        var v37 int32 = 292
                                                        if (!false) {
                                                            switch genFn0(255, -127, "") {
                                                            case -2147483648:
                                                                if (42 != 1) {
                                                                    var v38 uint = 42
                                                                    if false {
                                                                        if (42 > -32768) {
                                                                            v39 := GenS1{F0: uint64(v38), F1: v38}
                                                                            _ = v39
                                                                        }
                                                                        v14 -= 2
                                                                    }
                                                                    _ = v38
                                                                }
                                                            case 730:
                                                                v5 -= v5
                                                            case 32767:
                                                                var v40 float64 = (3.14159 / 1e10)
                                                                _ = v40
                                                            default:
                                                                v41 := v1[(v13 + v4)]
                                                                _ = v41
                                                            }
                                                            var v42 int16 = -785
                                                            v37 = (v37 - v37)
                                                            v14 -= (^(v14))
                                                            _ = v42
                                                        } else {
                                                            var v43 string = (v13 + (v13 + v4))
                                                            var v44 float64 = (1e10 + 0.1)
                                                            var v45 float64 = v44
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        var v46 uint16 = 679
                                                        switch v4 {
                                                        case "foo bar":
                                                            fmt.Println(v5, v4, v13)
                                                        case "a":
                                                            var v47 uint16 = (^(v46))
                                                            _ = v47
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                        _ = v46
                                                    } else {
                                                        v5 = v5
                                                    }
                                                    if (2 == 2) {
                                                        if (!(true && (4 >= 4294967295))) {
                                                            switch fi15 {
                                                            case 1:
                                                                var v48 uint16 = 42
                                                                _ = v48
                                                            case -2147483648:
                                                                switch (fi15 ^ fi15) {
                                                                case -717:
                                                                    v5 -= -2.25
                                                                }
                                                            case 4294967295:
                                                                switch v4 {
                                                                case "":
                                                                    var v49 int32 = 1
                                                                    _ = v49
                                                                case "\n\t":
                                                                    fmt.Println(v5, v5)
                                                                case "foo bar":
                                                                    v5 -= ((v5 + float32((-2.25 + ((3.14159 / -2.25) * float64(v5))))) * ((float32(0.1) + ((v5 + v5) * v5)) / v5))
                                                                default:
                                                                    if (!false) {
                                                                        var v50 int32 = 1
                                                                        fmt.Println(v4, v5)
                                                                        v51 := GenS2{F0: 32767, F1: 1, F2: v50}
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                }
                                                            default:
                                                                delete(v1, ("hello" + v13))
                                                            }
                                                            var v52 int8 = 59
                                                            _ = v52
                                                        } else {
                                                            _ = len(v2)
                                                            if false {
                                                                v14 = v14
                                                                var v53 uint32 = v14
                                                                if ((!false) && (64 > 100)) {
                                                                    if (255 < 114) {
                                                                        if ((((fi15 < (fi15 << ((42) & 63))) && false) && true) || ("a" == "")) {
                                                                            v4 = ("a" + ("\n\t" + v13))
                                                                            var v54 uint64 = 328
                                                                            _ = v54
                                                                        }
                                                                        if (445 <= 1) {
                                                                            v55 := GenS4{F0: 100}
                                                                            var v56 int8 = 0
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v55
                                                                            _ = v56
                                                                        } else {
                                                                            var v57 int64 = 1
                                                                            _ = v57
                                                                        }
                                                                        v53 -= (^(v53))
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        switch v13 {
                                                                        case "foo bar":
                                                                            var v58 int64 = 42
                                                                            _ = v58
                                                                        case "hello":
                                                                            v4 = v4
                                                                        case "":
                                                                            v59 := GenS5{F0: 582, F1: 814}
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                    var v60 float64 = float64(((1e10 - (v5 - 0.0)) / v5))
                                                                    _ = v60
                                                                } else {
                                                                    v61 := v1[(v4 + v4)]
                                                                    var v62 int = fi15
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                _ = v53
                                                            } else {
                                                                var v63 string = (v13 + ((v4 + v4) + (("a" + v13) + "hello")))
                                                                _ = v63
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 4294967295
                                                            }
                                                        }
                                                    } else {
                                                        v13 = v13
                                                        _ = len(v2)
                                                        if false {
                                                            v4 = ("foo bar" + v13)
                                                            v13 = "a"
                                                        } else {
                                                            delete(v1, "\n\t")
                                                        }
                                                    }
                                                }
                                            }
                                            v64 := GenS3{F0: fi24, F1: 1000, F2: 1}
                                            _ = v64
                                        case 42:
                                            _ = len(v2)
                                        case 2147483647:
                                            var v65 int = genFn0(986, 115, v4)
                                            _ = v65
                                        }
                                        fmt.Println(v13, fi15)
                                    }
                                } else {
                                    if false {
                                        v5 += v5
                                    } else {
                                        if false {
                                            if ((-(fi15)) <= genFn0(uint16(v14), 127, v4)) {
                                                var v66 int16 = 1
                                                fmt.Println(v13, v5, v14)
                                                _ = v66
                                            }
                                            if true {
                                                v5 -= -2.25
                                                v67 := GenS3{F0: (fi15 ^ (fi15 - fi15)), F1: -1, F2: 42}
                                                switch v13 {
                                                case "a":
                                                    var v68 int = (fi15 % ((fi15) | 1))
                                                    _ = v68
                                                case "hello":
                                                    fmt.Println(v14)
                                                case "":
                                                    v69 := GenS1{F0: 2, F1: 456}
                                                    _ = v69
                                                }
                                                v4 = v4
                                                v4 = "hello"
                                                var v70 bool = true
                                                v14 = (v14 * v14)
                                                _ = v67
                                                _ = v70
                                            }
                                        } else {
                                            var v71 uint8 = 0
                                            fmt.Println(v14)
                                            _ = v71
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = fi24
                                    }
                                }
                                _ = v30
                            } else {
                                var v72 int16 = 932
                                fmt.Println(v5, v4, fi24)
                                var v73 bool = false
                                _ = v72
                                _ = v73
                            }
                            if true {
                                var v74 bool = true
                                switch -9223372036854775808 {
                                case 984:
                                    var v75 int8 = -1
                                    _ = v75
                                case -9223372036854775808:
                                    var v76 int16 = -32768
                                    _ = v76
                                default:
                                    v74 = ((v5 > v5) && (0 >= 42))
                                }
                                fmt.Println(v5)
                                _ = v74
                            } else {
                                v77 := GenS4{F0: 18}
                                v78 := GenS6{F0: 0, F1: (^(v14)), F2: 1}
                                _ = v77
                                _ = v78
                            }
                            v14 = (v14 ^ v14)
                        }
                        var v79 uint16 = 32767
                        _ = fi24
                        _ = v79
                    }
                    _ = v23
                }
                v13 = v13
            } else {
                v5 = float32((0.0 / (0.1 / -2.25)))
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v3)
        _ = fi15
    }
    if true {
        v80 := GenS7{F0: -452, F1: "foo bar"}
        _ = v80
    } else {
        var v81 int = 9223372036854775807
        for fi82 := 0; fi82 < 17; fi82++ {
            var v83 int8 = 0
            _ = fi82
            _ = v83
        }
        switch (v81 / (((v81 ^ -1)) | 1)) {
        case 32767:
            switch genFn0(2, 0, "foo bar") {
            case -409:
                v5 += (((v5 * (((0.1 * float32(1e10)) * (1.5 + v5)) - (0.0 / v5))) * float32(-2.25)) / (v5 + v5))
            default:
                v14 += 289
            }
        case 127:
            fmt.Println(v13)
        default:
            var v84 uint32 = 2147483648
            _ = v84
        }
        _ = v81
    }
    var v85 int = 803
    delete(v1, v13)
    v86 := GenS4{F0: 127}
    for fi87 := 0; fi87 < 19; fi87++ {
        var v88 int = (^(fi87))
        _ = fi87
        _ = v88
    }
    for fi89 := 0; fi89 < 6; fi89++ {
        for fi90 := 0; fi90 < 1; fi90++ {
            if false {
                var v91 int32 = 2147483647
                v92 := GenS5{F0: uint16(v14), F1: -321}
                var v93 uint64 = 422
                _ = v91
                _ = v92
                _ = v93
            }
            v85 -= 4294967295
            _ = fi90
        }
        v5 += 1e10
        v14 = v14
        _ = fi89
    }
    switch v13 {
    case "\n\t":
        v85 += (v85 - (^(v85)))
    case "foo bar":
        v85 = v85
    case "a":
        if true {
            if (1 <= -1) {
                for fi94 := 0; fi94 < 6; fi94++ {
                    if true {
                        fmt.Println(v13, v14, v14)
                        fmt.Println(v4, v14)
                    }
                    _ = fi94
                }
                _ = len(v3)
                if (42 != 127) {
                    for fi95 := 0; fi95 < 12; fi95++ {
                        for fi96 := 0; fi96 < 1; fi96++ {
                            var v97 int64 = 42
                            var v98 int64 = 9223372036854775807
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                        var v99 uint32 = 4294967295
                        var v100 uint = 42
                        _ = fi95
                        _ = v100
                        _ = v99
                    }
                    v101 := GenS6{F0: 1, F1: v14, F2: 915}
                    if len(v2) > 0 {
                        v2[0] = ((v13 + (v13 + v13)) + ((v4 + (v4 + ("" + v13))) + v4))
                    }
                    _ = v101
                } else {
                    var v102 int32 = 42
                    v85 = genFn0(255, 29, v13)
                    if (!false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v103 := GenS8{F0: v14, F1: v13, F2: (65535 <= 311)}
                        if true {
                            fmt.Println(v4, v4)
                            var v104 int = (-(v85))
                            v4 = v4
                            _ = v104
                        }
                        v14 = 1
                        _ = v103
                    }
                    _ = v102
                }
            } else {
                var v105 float32 = 3.14159
                _ = v105
            }
            if false {
                v106 := GenS8{F0: (v14 & 2), F1: v13, F2: true}
                v13 = v4
                _ = v106
            }
            _ = len(v2)
        } else {
            fmt.Println(v14, v14)
        }
    default:
        var v107 float32 = v5
        _ = v107
    }
    v108 := GenS7{F0: v85, F1: ("foo bar" + "a")}
    var v109 uint = 317
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v108
    _ = v109
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
}
