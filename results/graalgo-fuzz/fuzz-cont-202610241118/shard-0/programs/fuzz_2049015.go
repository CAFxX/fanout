package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int64; F2 int64 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 int8; F1 bool; F2 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 int16; F2 bool }
type GenS12 struct { F0 bool; F1 uint8 }
type GenS13 struct { F0 uint8 }
func genFn0(p0 uint64) uint8 {
    p0 = p0
    var v4 bool = (("hello" + ("\n\t" + "hello")) != "hello")
    var v5 bool = (-1 > 228)
    var v6 int32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v7 := GenS1{F0: 0, F1: 0}
    v8 := v2[128]
    v1 = -265
    v1 = 1
    var v9 int8 = 100
    switch "foo bar" {
    case "a":
        var v10 float32 = 0.0
        _ = v10
    case "":
        var v11 float32 = 0.1
        _ = v11
    default:
        v12 := []int{(v8 * (v8 & (v8 & 42))), int(v9), 323, v8}
        _ = v12
    }
    var v13 float32 = -2.25
    if (2 >= 255) {
        v14 := GenS1{F0: 1, F1: 42}
        var v15 uint16 = 42
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint16 = v15
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := []int{fi18}
                v1 += 2147483647
                v13 = v13
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(fi20, v17)
                var v21 float32 = (0.0 + 1.5)
                if false {
                    v22 := []int{fi16}
                    _ = v22
                } else {
                    if false {
                        switch (("foo bar" + "\n\t") + "foo bar") {
                        case "a":
                            v23 := GenS1{F0: 4294967295, F1: v17}
                            _ = v23
                        case "hello":
                            if false {
                                if (v15 == v15) {
                                    v24 := GenS2{F0: 143, F1: 1, F2: (-794 != v1)}
                                    _ = v24
                                } else {
                                    if false {
                                        if false {
                                            v1 += (v1 & (v1 - v1))
                                        } else {
                                            if true {
                                                var v25 uint8 = 42
                                                _ = v25
                                            }
                                            fmt.Println(v21, v13)
                                            _ = len(v2)
                                            v26 := v2[(fi20 + (v8 & 1))]
                                            _ = v26
                                        }
                                        if true {
                                            v3["foo bar"] = "\n\t"
                                            var v27 int64 = 1
                                            v28 := GenS2{F0: 2, F1: 1, F2: true}
                                            _ = v27
                                            _ = v28
                                        }
                                        switch (-(fi20)) {
                                        case -657:
                                            var v29 uint = 42
                                            _ = v29
                                        case 1:
                                            v21 += ((v21 + v13) * float32((1e10 * (0.1 - 3.14159))))
                                        case -111:
                                            delete(v2, int(v9))
                                        default:
                                            v15 += v17
                                        }
                                        v15 = v17
                                    } else {
                                        fmt.Println(v17)
                                        v30 := []int{v8}
                                        if (633 != 940) {
                                            if len(v30) > 0 {
                                                v30[0] = fi16
                                            }
                                            v8 = 1
                                        } else {
                                            if true {
                                                v31 := GenS3{F0: v8}
                                                _ = v31
                                            } else {
                                                v21 -= v13
                                                fmt.Println(v15, fi20)
                                                var v32 uint32 = 698
                                                var v33 uint32 = v32
                                                if (32767 != -32768) {
                                                    switch fi16 {
                                                    case 32768:
                                                        v30 = append(v30, (fi16 * 42))
                                                    }
                                                    if true {
                                                        if true {
                                                            if false {
                                                                switch v8 {
                                                                case 127:
                                                                    if (v1 <= v1) {
                                                                        v9 = (v9 / ((v9) | 1))
                                                                        var v34 int = (v8 | -593)
                                                                        if false {
                                                                            switch fi20 {
                                                                            case -236:
                                                                                v34 = (v8 / (((v34 | (^(fi20)))) | 1))
                                                                            default:
                                                                                var v35 uint16 = v17
                                                                                _ = v35
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                if ((("foo bar" + ("" + "a")) < ("foo bar" + "a")) && false) {
                                                                                    v33 += v33
                                                                                    var v36 uint64 = 42
                                                                                    _ = v36
                                                                                } else {
                                                                                    v37 := GenS1{F0: 2, F1: (^(v15))}
                                                                                    var v38 float64 = (-2.25 - (0.0 + 0.1))
                                                                                    if (!false) {
                                                                                        v39 := GenS3{F0: (v34 - (v34 / ((2147483647) | 1)))}
                                                                                        _ = v39
                                                                                    }
                                                                                    _ = v37
                                                                                    _ = v38
                                                                                }
                                                                                v40 := GenS4{F0: 2}
                                                                                _ = v40
                                                                            }
                                                                            if true {
                                                                                if (((("\n\t" + (("" + "foo bar") + "")) + ("hello" + ("hello" + "hello"))) + "foo bar") < ("foo bar" + ("hello" + "foo bar"))) {
                                                                                    v41 := GenS3{F0: 4294967295}
                                                                                    fmt.Println(v1, v33, v32)
                                                                                    var v42 int16 = 42
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                            var v43 string = "\n\t"
                                                                            _ = v43
                                                                        }
                                                                        fmt.Println(fi20, v17)
                                                                        _ = v34
                                                                    } else {
                                                                        v44 := GenS2{F0: 0, F1: int64(v32), F2: (genFn0(42) != 255)}
                                                                        _ = v44
                                                                    }
                                                                    v2[v8] = fi16
                                                                    v45 := GenS5{F0: 42, F1: ("" + ("" + "\n\t"))}
                                                                    _ = v45
                                                                default:
                                                                    fmt.Println(fi20, v21, v15)
                                                                }
                                                            } else {
                                                                fmt.Println(v1, v9, v33)
                                                            }
                                                            v46 := GenS5{F0: 0, F1: ""}
                                                            if ((float64(v13) - 0.0) < (0.1 + 3.14159)) {
                                                                var v47 int64 = -152
                                                                _ = v47
                                                            }
                                                            if (true && (false && (!(false || (-1 <= 2147483647))))) {
                                                                var v48 int64 = 0
                                                                v49 := v3["foo bar"]
                                                                v17 = 42
                                                                var v50 uint32 = v32
                                                                fmt.Println(v15, v50, v8)
                                                                fmt.Println(v13, v13, v9)
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                switch 128 {
                                                                case 32767:
                                                                    v13 = 1.5
                                                                case 2147483647:
                                                                    v51 := GenS2{F0: 127, F1: 42, F2: false}
                                                                    _ = v51
                                                                case 27:
                                                                    v15 -= (v17 - 2)
                                                                default:
                                                                    delete(v2, (-(fi20)))
                                                                }
                                                                switch 1 {
                                                                case 127:
                                                                    var v52 float64 = (3.14159 + 1e10)
                                                                    _ = v52
                                                                case 910:
                                                                    if ((false && true) || false) {
                                                                        v13 = v21
                                                                        fmt.Println(v8, v17)
                                                                        v13 = 0.1
                                                                    }
                                                                }
                                                            }
                                                            _ = v46
                                                        }
                                                        v3[""] = ""
                                                    } else {
                                                        var v53 uint = 878
                                                        var v54 uint8 = 167
                                                        fmt.Println(v15, v21, fi20)
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                } else {
                                                    v55 := GenS6{F0: 65535}
                                                    if (9223372036854775807 != 679) {
                                                        if (!(!true)) {
                                                            if false {
                                                                if (0.1 != 3.14159) {
                                                                    v32 = 0
                                                                }
                                                                v9 = v9
                                                                var v56 float32 = v21
                                                                _ = v56
                                                            }
                                                            var v57 uint8 = 1
                                                            fmt.Println(v15, v1, v33)
                                                            v58 := GenS3{F0: 57}
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v59 := GenS7{F0: 435, F1: -9223372036854775808, F2: 740}
                                                        v13 += float32(0.1)
                                                        _ = v59
                                                    } else {
                                                        v17 = (v17 - v15)
                                                    }
                                                    var v60 uint64 = 42
                                                    _ = v55
                                                    _ = v60
                                                }
                                                _ = v32
                                                _ = v33
                                            }
                                            if (42 > 1) {
                                                var v61 string = ("\n\t" + "a")
                                                if true {
                                                    switch fi20 {
                                                    case 465:
                                                        fmt.Println(v61)
                                                    }
                                                } else {
                                                    var v62 int = fi16
                                                    fmt.Println(v9, v8, v13)
                                                    var v63 int = (fi16 & 2147483647)
                                                    _ = v62
                                                    _ = v63
                                                }
                                                _ = v61
                                            }
                                            v64 := GenS8{F0: 1, F1: false}
                                            _ = v64
                                        }
                                        _ = v30
                                    }
                                    fmt.Println(fi20, v13)
                                    v9 -= (v9 / ((7) | 1))
                                    var v65 uint32 = 4294967295
                                    var v66 string = (("hello" + "hello") + "hello")
                                    _ = v65
                                    _ = v66
                                }
                            } else {
                                v8 += (fi20 << ((2) & 63))
                                if true {
                                    fmt.Println(fi16)
                                    v67 := GenS1{F0: 18446744073709551615, F1: (v17 | (v17 % (((v17 + (v17 & 255))) | 1)))}
                                    if (42 <= 42) {
                                        v68 := GenS9{F0: 127, F1: false, F2: int64(v9)}
                                        var v69 uint8 = 2
                                        _ = v68
                                        _ = v69
                                    } else {
                                        var v70 int8 = v9
                                        _ = v70
                                    }
                                    _ = v67
                                } else {
                                    v9 += -122
                                }
                                var v71 int64 = -1
                                _ = v71
                            }
                        default:
                            switch 9223372036854775807 {
                            case 2147483647:
                                _ = len(v2)
                            default:
                                v13 -= float32(0.1)
                            }
                        }
                        v72 := GenS10{F0: (^(v9))}
                        _ = v72
                    } else {
                        var v73 int64 = -612
                        fmt.Println(v17, fi20, v13)
                        v74 := GenS2{F0: 200, F1: (v73 >> ((255) & 63)), F2: (-1 >= 1)}
                        _ = v73
                        _ = v74
                    }
                    var v75 uint = 955
                    var v76 uint16 = v17
                    _ = v75
                    _ = v76
                }
                v15 -= v15
                delete(v2, 4294967295)
                switch 42 {
                case 32767:
                    if (v9 < v9) {
                        var v77 string = "a"
                        fmt.Println(v1)
                        v17 += (v15 ^ (v15 * 2))
                        _ = v77
                    } else {
                        v78 := GenS11{F0: 42, F1: int16(v9), F2: (4294967295 >= 1)}
                        v9 -= (v9 % (((v9 << ((4294967295) & 63))) | 1))
                        var v79 uint8 = 0
                        _ = v78
                        _ = v79
                    }
                case 42:
                    var v80 int64 = 1
                    _ = v80
                }
                _ = fi20
                _ = v21
            }
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    fmt.Println(v8, v9, v13)
    v9 -= (v9 / (((v9 | 0)) | 1))
    switch "hello" {
    case "\n\t":
        v81 := GenS12{F0: true, F1: 1}
        _ = v81
    case "hello":
        for fi82 := 0; fi82 < 17; fi82++ {
            var v83 uint8 = 128
            for fi84 := 0; fi84 < 17; fi84++ {
                v83 = v83
                v85 := GenS13{F0: (^(v83))}
                _ = fi84
                _ = v85
            }
            if false {
                if ((-2 >= -1) && (42 != -487)) {
                    if (((4294967295 == 18446744073709551615) && false) && false) {
                        v13 = (v13 - -2.25)
                        switch "a" {
                        case "hello":
                            v1 += (^(v1))
                        default:
                            v86 := []int{v8}
                            _ = v86
                        }
                    }
                }
            }
            _ = fi82
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
