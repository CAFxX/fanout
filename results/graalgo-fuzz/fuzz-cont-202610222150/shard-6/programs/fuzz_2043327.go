package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 string; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 int32 = 0
    switch -1 {
    case 4294967295:
        v3 = 42
    case 0:
        for fi4 := 0; fi4 < 20; fi4++ {
            if (true && (2 >= 19)) {
                v5 := GenS1{F0: (v3 >> ((1) & 63)), F1: 18446744073709551615, F2: 900}
                v3 += (v3 % (((v3 - v3)) | 1))
                _ = v5
            } else {
                fmt.Println(fi4, v3)
            }
            v3 = v3
            fmt.Println(v3, fi4)
            _ = fi4
        }
    default:
        var v6 int64 = -1
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{fi7, fi7, (^(fi7))}
            v9 := GenS1{F0: (v3 ^ -2147483648), F1: 2, F2: 18446744073709551615}
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                if true {
                    if true {
                        if (false || ((1e10 / (1.5 - -2.25)) != 3.14159)) {
                            switch (fi10 + (fi10 % (((^(fi10))) | 1))) {
                            case 127:
                                var v11 int8 = -1
                                _ = v11
                            case 128:
                                var v12 int32 = 1
                                _ = v12
                            default:
                                var v13 int16 = -627
                                _ = v13
                            }
                            fmt.Println(fi10, fi10)
                        }
                        if false {
                            v14 := GenS2{F0: v3, F1: false}
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 4; fi15++ {
                                fmt.Println(v3, fi15, fi15)
                                var v16 int = (fi15 >> ((42) & 63))
                                var v17 int = (v16 & 128)
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                        if true {
                            for fi18 := 0; fi18 < 14; fi18++ {
                                if false {
                                    var v19 uint32 = 644
                                    _ = v19
                                } else {
                                    if (!false) {
                                        var v20 uint32 = 2147483647
                                        _ = v20
                                    } else {
                                        if (-548 == 32767) {
                                            switch "hello" {
                                            case "":
                                                fmt.Println(fi18, fi18)
                                            case "\n\t":
                                                fmt.Println(fi10)
                                            case "a":
                                                v21 := GenS1{F0: (v3 / ((v3) | 1)), F1: 18446744073709551615, F2: 255}
                                                _ = v21
                                            }
                                            switch fi10 {
                                            case -2147483648:
                                                v3 = int32(fi18)
                                            case 1:
                                                if (true && false) {
                                                    if true {
                                                        v22 := GenS2{F0: (v3 + (v3 >> ((255) & 63))), F1: (960 == 18446744073709551615)}
                                                        if false {
                                                            var v23 uint8 = 0
                                                            var v24 uint = 174
                                                            v25 := GenS1{F0: v3, F1: 42, F2: (v24 ^ (v24 + v24))}
                                                            _ = v23
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            if (!true) {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    switch (fi18 | (-(fi18))) {
                                                                    case -9223372036854775808:
                                                                        if (0 < 32767) {
                                                                            v26 := GenS1{F0: v3, F1: 0, F2: 498}
                                                                            v27 := GenS3{F0: ("" + "a")}
                                                                            _ = v26
                                                                            _ = v27
                                                                        } else {
                                                                            v3 += (v3 * (v3 % (((v3 * (v3 | -2147483648))) | 1)))
                                                                        }
                                                                    case 42:
                                                                        if (true || false) {
                                                                            var v28 uint16 = 0
                                                                            if false {
                                                                                v28 += (v28 & (^(v28)))
                                                                                var v29 uint8 = 42
                                                                                _ = v29
                                                                            } else {
                                                                                v30 := []int{(fi10 >> ((2) & 63))}
                                                                                v31 := GenS2{F0: v3, F1: false}
                                                                                var v32 uint32 = 2147483648
                                                                                _ = v30
                                                                                _ = v31
                                                                                _ = v32
                                                                            }
                                                                            _ = v28
                                                                        }
                                                                    case 1:
                                                                        if (32768 == (fi18 & (fi10 >> ((42) & 63)))) {
                                                                            var v33 uint = 800
                                                                            _ = v33
                                                                        } else {
                                                                            v3 += v3
                                                                        }
                                                                    default:
                                                                        v34 := GenS3{F0: "hello"}
                                                                        _ = v34
                                                                    }
                                                                default:
                                                                    v35 := GenS4{F0: 255}
                                                                    _ = v35
                                                                }
                                                                v36 := []int{}
                                                                _ = v36
                                                            }
                                                        }
                                                        _ = v22
                                                    } else {
                                                        v1[(fi10 / (((-(fi10))) | 1))] = (("a" + ("" + "a")) + "\n\t")
                                                    }
                                                    v37 := GenS2{F0: (v3 ^ v3), F1: (42 == 0)}
                                                    _ = v37
                                                } else {
                                                    v3 = -645
                                                    v38 := []int{(fi18 - (-(fi18))), fi10, (fi18 & (fi10 >> ((1) & 63)))}
                                                    _ = len(v38)
                                                    _ = v38
                                                }
                                                var v39 uint = 4294967295
                                                var v40 uint32 = 2
                                                _ = v39
                                                _ = v40
                                            }
                                        } else {
                                            fmt.Println(fi18, fi18, fi18)
                                        }
                                    }
                                    v41 := GenS2{F0: -1, F1: false}
                                    _ = v41
                                }
                                delete(v1, -9223372036854775808)
                                v42 := GenS3{F0: "a"}
                                _ = fi18
                                _ = v42
                            }
                            v43 := GenS3{F0: "\n\t"}
                            _ = v43
                        } else {
                            v44 := GenS2{F0: (v3 & v3), F1: false}
                            _ = v44
                        }
                    }
                    var v45 uint8 = 0
                    v3 += v3
                    _ = v45
                } else {
                    v3 = v3
                    for fi46 := 0; fi46 < 11; fi46++ {
                        v3 -= (v3 - v3)
                        if false {
                            fmt.Println(v3)
                            v47 := GenS5{F0: "", F1: -2147483648, F2: (v3 % (((v3 + v3)) | 1))}
                            _ = v47
                        } else {
                            v48 := GenS3{F0: "\n\t"}
                            v3 += v3
                            var v49 int8 = 42
                            _ = v48
                            _ = v49
                        }
                        var v50 int8 = 0
                        if (42 <= (fi10 / (((fi10 | -2147483648)) | 1))) {
                            fmt.Println(v50, v50, v50)
                            if false {
                                v50 = v50
                                v51 := GenS5{F0: "foo bar", F1: -2147483648, F2: (-(v3))}
                                _ = v51
                            } else {
                                if ((1e10 - 3.14159) > ((0.0 / 1.0) * 0.0)) {
                                    var v52 bool = true
                                    if (v52 || v52) {
                                        switch -1 {
                                        case -1:
                                            fmt.Println(v3)
                                        case 972:
                                            switch "hello" {
                                            case "foo bar":
                                                var v53 uint16 = 255
                                                _ = v53
                                            default:
                                                v54 := GenS6{F0: "foo bar", F1: 255, F2: (v50 % (((v50 | v50)) | 1))}
                                                _ = v54
                                            }
                                        }
                                        var v55 string = ((("a" + "") + "hello") + "hello")
                                        _ = v55
                                    } else {
                                        fmt.Println(v52)
                                    }
                                    _ = v52
                                } else {
                                    var v56 float64 = 0.0
                                    v57 := []int{-1, 32767, (fi10 + 32768)}
                                    _ = v56
                                    _ = v57
                                }
                                if true {
                                    if (!(0 < 0)) {
                                        v50 -= v50
                                    } else {
                                        if ((v3 & v3) != -2147483648) {
                                            if (0.1 < 1e10) {
                                                v58 := GenS7{F0: 2}
                                                _ = v58
                                            }
                                            v50 = (-(v50))
                                        }
                                        var v59 uint = 820
                                        var v60 string = "hello"
                                        var v61 float32 = (0.1 - 3.14159)
                                        var v62 int = fi46
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                    if false {
                                        var v63 uint16 = 299
                                        _ = v63
                                    }
                                    var v64 uint = 949
                                    _ = v64
                                }
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        switch "hello" {
                                        case "a":
                                            var v65 uint = 18446744073709551615
                                            var v66 bool = (!(v3 > (v3 + (v3 >> ((1) & 63)))))
                                            _ = v65
                                            _ = v66
                                        case "foo bar":
                                            v3 -= (v3 + 2147483647)
                                        case "hello":
                                            v67 := GenS1{F0: (v3 & v3), F1: 42, F2: 2}
                                            _ = v67
                                        default:
                                            if true {
                                                if (125 > 2) {
                                                    delete(v1, (fi46 * fi10))
                                                    v3 = -1
                                                    v68 := GenS3{F0: ""}
                                                    _ = v68
                                                }
                                                var v69 uint16 = 42
                                                _ = v69
                                            } else {
                                                v1[(fi46 / ((fi46) | 1))] = ("\n\t" + "hello")
                                                v70 := GenS6{F0: "hello", F1: 395, F2: v50}
                                                _ = v70
                                            }
                                        }
                                    }
                                default:
                                    if false {
                                        if (0.0 >= 3.14159) {
                                            fmt.Println(v50, v50, fi46)
                                        } else {
                                            fmt.Println(fi46, v3, v3)
                                        }
                                    } else {
                                        v71 := GenS3{F0: "\n\t"}
                                        fmt.Println(v50)
                                        _ = v71
                                    }
                                }
                            }
                            if (4294967295 <= 404) {
                                v3 = v3
                                if false {
                                    fmt.Println(fi10, v50)
                                    switch -413 {
                                    case 371:
                                        v50 = (v50 ^ 100)
                                    default:
                                        v3 -= -2147483648
                                    }
                                } else {
                                    v72 := GenS7{F0: 1}
                                    v73 := GenS3{F0: "foo bar"}
                                    var v74 int32 = (^(v3))
                                    v74 = int32(v50)
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                                _ = len(v1)
                            }
                        } else {
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                v50 = (v50 / ((v50) | 1))
                            case 42:
                                v75 := []int{fi46, (fi46 >> ((255) & 63))}
                                _ = v75
                            }
                            var v76 int = fi46
                            v76 = fi10
                            v76 = fi46
                            _ = v76
                        }
                        if (v50 == 42) {
                            v77 := GenS5{F0: (("foo bar" + "a") + ("foo bar" + "hello")), F1: (v3 % (((^(v3))) | 1)), F2: v3}
                            _ = v77
                        } else {
                            var v78 float64 = ((1.5 / 1.5) + 3.14159)
                            var v79 int16 = 1
                            fmt.Println(v50, v78)
                            if (fi10 == fi10) {
                                _ = len(v1)
                            }
                            v79 = 0
                            _ = v78
                            _ = v79
                        }
                        _ = fi46
                        _ = v50
                    }
                    switch "a" {
                    case "a":
                        for fi80 := 0; fi80 < 14; fi80++ {
                            var v81 uint32 = 2147483648
                            _ = fi80
                            _ = v81
                        }
                    case "\n\t":
                        v82 := GenS8{F0: 487}
                        _ = v82
                    }
                }
            } else {
                var v83 uint8 = 127
                _ = v83
            }
            v3 -= v3
            for fi84 := 0; fi84 < 12; fi84++ {
                var v85 string = "foo bar"
                v86 := GenS9{F0: fi10}
                v3 += (-(v3))
                _ = fi84
                _ = v85
                _ = v86
            }
            _ = fi10
        }
        var v87 uint64 = 4294967295
        _ = v87
    } else {
        if false {
            delete(v1, 9223372036854775807)
            v88 := GenS10{F0: "\n\t", F1: -730}
            v89 := GenS11{F0: 0}
            _ = v88
            _ = v89
        } else {
            v3 = (v3 - 26)
        }
        v90 := GenS9{F0: 128}
        var v91 bool = false
        _ = v90
        _ = v91
    }
    v92 := GenS7{F0: 65535}
    for fi93 := 0; fi93 < 19; fi93++ {
        delete(v1, fi93)
        _ = fi93
    }
    for fi94 := 0; fi94 < 4; fi94++ {
        if (377 <= 0) {
            v2[""] = 1
            for fi95 := 0; fi95 < 13; fi95++ {
                v3 += (v3 % ((1) | 1))
                switch ("\n\t" + "") {
                case "a":
                    var v96 bool = ("a" != "a")
                    _ = v96
                case "\n\t":
                    v3 = (v3 ^ v3)
                }
                var v97 int = 0
                _ = fi95
                _ = v97
            }
            v3 -= v3
        }
        v98 := GenS10{F0: ("" + ("" + ("" + "foo bar"))), F1: -316}
        var v99 uint = 2
        _ = fi94
        _ = v98
        _ = v99
    }
    v100 := GenS3{F0: (("foo bar" + (("foo bar" + (("" + ("hello" + "foo bar")) + "foo bar")) + "\n\t")) + ("a" + "hello"))}
    switch "" {
    case "":
        fmt.Println(v3)
    case "\n\t":
        switch "" {
        case "foo bar":
            delete(v2, "a")
        }
    case "foo bar":
        var v101 int64 = -1
        _ = v101
    }
    var v102 bool = true
    v103 := GenS12{F0: (v3 * -2147483648), F1: "\n\t"}
    v104 := []int{}
    var v105 int = -1
    var v106 uint8 = 121
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v100
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v92
}
