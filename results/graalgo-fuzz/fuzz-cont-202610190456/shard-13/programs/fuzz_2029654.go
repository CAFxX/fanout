package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 uint16; F1 bool; F2 int64 }
type GenS9 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32767
    var v3 int = -374
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    if (!true) {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 248}
            var v8 uint64 = 474
            _ = fi6
            _ = v7
            _ = v8
        }
        if true {
            switch "foo bar" {
            case "\n\t":
                v9 := v5[""]
                _ = v9
            case "":
                var v10 uint8 = 76
                _ = v10
            case "foo bar":
                var v11 bool = true
                _ = v11
            }
            v12 := GenS1{F0: 894}
            var v13 int64 = -9223372036854775808
            _ = v12
            _ = v13
        } else {
            var v14 int = v3
            switch (v14 | (^(v14))) {
            case 2147483647:
                switch "" {
                case "foo bar":
                    v14 += -426
                default:
                    var v15 float64 = -2.25
                    _ = v15
                }
            }
            _ = v14
        }
        _ = len(v5)
    } else {
        switch v2 {
        case 1:
            v16 := GenS1{F0: 65535}
            _ = v16
        }
        switch ("a" + ("a" + "")) {
        case "hello":
            switch (v3 % ((2147483647) | 1)) {
            case 42:
                switch 32767 {
                case 790:
                    var v17 uint64 = 168
                    _ = v17
                default:
                    v18 := GenS2{F0: (!(509 < 255))}
                    _ = v18
                }
            case 1:
                if false {
                    if (255 >= 2) {
                        if false {
                            switch (v3 | (v3 + (v2 << ((232) & 63)))) {
                            case 2147483647:
                                var v19 int64 = -306
                                _ = v19
                            case 9223372036854775807:
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    var v21 float64 = ((0.1 * 0.0) - (1.5 * 0.1))
                                    _ = fi20
                                    _ = v21
                                }
                            case 0:
                                v22 := GenS3{F0: "hello", F1: 30}
                                _ = v22
                            }
                            v23 := GenS4{F0: ("" + "a"), F1: false}
                            _ = v23
                        } else {
                            if true {
                                if false {
                                    if true {
                                        v3 = v2
                                        if false {
                                            var v24 uint8 = 127
                                            for fi25 := 0; fi25 < 15; fi25++ {
                                                for fi26 := 0; fi26 < 11; fi26++ {
                                                    _ = len(v4)
                                                    if false {
                                                        fmt.Println(v24)
                                                        fmt.Println(v3, fi26)
                                                    } else {
                                                        if (false && true) {
                                                            if false {
                                                                v24 += v24
                                                            }
                                                        }
                                                        switch -1 {
                                                        case -1:
                                                            v27 := GenS2{F0: false}
                                                            _ = v27
                                                        case 127:
                                                            _ = len(v5)
                                                            switch ("hello" + "foo bar") {
                                                            case "a":
                                                                if (!false) {
                                                                    var v28 int8 = -93
                                                                    _ = v28
                                                                }
                                                                switch 42 {
                                                                case 1:
                                                                    var v29 int64 = 0
                                                                    _ = v29
                                                                }
                                                                switch ("hello" + "a") {
                                                                case "\n\t":
                                                                    v24 = 128
                                                                case "foo bar":
                                                                    if true {
                                                                        v24 = v24
                                                                        v24 -= (v24 >> ((65535) & 63))
                                                                    }
                                                                default:
                                                                    if true {
                                                                        var v30 int8 = -127
                                                                        v31 := GenS3{F0: "\n\t", F1: 286}
                                                                        if (true && (true || (42 <= 128))) {
                                                                            if false {
                                                                                var v32 uint32 = 436
                                                                                var v33 int8 = v30
                                                                                v34 := []int{-77, fi25, (fi25 * 9223372036854775807)}
                                                                                _ = v32
                                                                                _ = v33
                                                                                _ = v34
                                                                            }
                                                                            var v35 int8 = (-(v30))
                                                                            _ = v35
                                                                        }
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v36 := GenS4{F0: ("a" + ((("hello" + "a") + (("foo bar" + "hello") + ("" + ("" + "")))) + "a")), F1: (!false)}
                                                                _ = v36
                                                            case "foo bar":
                                                                v37 := GenS4{F0: "hello", F1: (false || (!(!(v24 > v24))))}
                                                                _ = v37
                                                            default:
                                                                v38 := v1["a"]
                                                                _ = v38
                                                            }
                                                            _ = len(v4)
                                                        default:
                                                            var v39 int8 = 0
                                                            _ = v39
                                                        }
                                                        v3 += 9223372036854775807
                                                    }
                                                    _ = fi26
                                                }
                                                v40 := GenS5{F0: 65535, F1: 18446744073709551615}
                                                _ = fi25
                                                _ = v40
                                            }
                                            var v41 int64 = 1
                                            _ = v24
                                            _ = v41
                                        } else {
                                            fmt.Println(v3, v3)
                                        }
                                    }
                                    var v42 int64 = 421
                                    _ = v42
                                }
                                v2 = v3
                            }
                        }
                        var v43 int = v2
                        _ = v43
                    }
                    if true {
                        v44 := GenS4{F0: "\n\t", F1: false}
                        for fi45 := 0; fi45 < 8; fi45++ {
                            if true {
                                v3 = -1
                                var v46 float64 = (1e10 * -2.25)
                                var v47 uint8 = 63
                                _ = v46
                                _ = v47
                            }
                            var v48 uint32 = 256
                            _ = fi45
                            _ = v48
                        }
                        _ = v44
                    }
                } else {
                    v2 = v3
                }
            }
        }
        fmt.Println(v3)
    }
    if ((-966 != 1000) || (!(!(42 >= 46)))) {
        fmt.Println(v2, v2)
        v49 := GenS1{F0: 65535}
        _ = v49
    } else {
        _ = len(v4)
        for fi50 := 0; fi50 < 3; fi50++ {
            v2 -= (v3 << ((2) & 63))
            var v51 int8 = 127
            fmt.Println(v2, v51, fi50)
            _ = fi50
            _ = v51
        }
        v52 := []int{(v2 ^ v2)}
        _ = v52
    }
    _ = len(v4)
    var v53 uint32 = 839
    v54 := GenS4{F0: "", F1: true}
    v53 = 65535
    if false {
        switch (v3 + v3) {
        case 776:
            v3 += v3
        }
        v55 := []int{(v2 << ((18446744073709551615) & 63)), 32768, (v2 ^ v2), v3}
        _ = v55
    } else {
        v56 := GenS2{F0: (-113 != -111)}
        for fi57 := 0; fi57 < 20; fi57++ {
            if (703 == 8) {
                var v58 string = "hello"
                _ = v58
            } else {
                for fi59 := 0; fi59 < 19; fi59++ {
                    v60 := GenS4{F0: "a", F1: ((!(-9223372036854775808 != 154)) || (128 >= 209))}
                    v61 := []int{v2, 128, fi57}
                    switch "foo bar" {
                    case "foo bar":
                        v62 := GenS5{F0: 162, F1: 18446744073709551615}
                        _ = v62
                    case "":
                        v3 -= fi57
                    case "a":
                        switch (("" + "\n\t") + "") {
                        case "hello":
                            v63 := GenS5{F0: 701, F1: 4294967295}
                            _ = v63
                        case "a":
                            v3 = (v2 % ((fi59) | 1))
                        case "":
                            switch (fi59 >> ((140) & 63)) {
                            case 42:
                                var v64 float64 = (-2.25 * 1e10)
                                _ = v64
                            }
                        default:
                            var v65 uint8 = 128
                            _ = v65
                        }
                    }
                    switch v3 {
                    case -1:
                        if ((96 >= 255) || true) {
                            _ = len(v61)
                            switch -2147483648 {
                            case 989:
                                if (648 != 549) {
                                    fmt.Println(v3, fi57, v2)
                                    var v66 uint32 = 2147483648
                                    var v67 string = ("\n\t" + "a")
                                    if false {
                                        var v68 uint64 = 1
                                        var v69 uint8 = 2
                                        var v70 int64 = 397
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            v3 = v3
                                        }
                                    }
                                    _ = v66
                                    _ = v67
                                } else {
                                    var v71 uint8 = uint8(v53)
                                    _ = v71
                                }
                            }
                            v72 := GenS6{F0: 42}
                            _ = v72
                        } else {
                            v4[(("" + "") + (("foo bar" + ((("hello" + "hello") + ("foo bar" + "foo bar")) + "\n\t")) + "a"))] = (v53 * v53)
                            fmt.Println(v3, v53)
                        }
                        var v73 int64 = 1
                        v61 = append(v61, (-(fi59)))
                        _ = v73
                    default:
                        delete(v1, "foo bar")
                    }
                    var v74 int = -9223372036854775808
                    v4[("hello" + "\n\t")] = (v53 / (((v53 ^ v53)) | 1))
                    var v75 float64 = 1e10
                    fmt.Println(v75, v74)
                    var v76 string = ""
                    fmt.Println(v75, v2, v76)
                    if false {
                        v2 = (fi57 * (v74 ^ 296))
                        if (1 <= 0) {
                            delete(v4, v76)
                            v77 := v5[v76]
                            v75 = v75
                            _ = v77
                        }
                    } else {
                        switch v2 {
                        case -2147483648:
                            fmt.Println(v53, v3)
                        case 0:
                            var v78 uint8 = 1
                            _ = v78
                        }
                    }
                    v61 = append(v61, int(v53))
                    if false {
                        fmt.Println(v75)
                    } else {
                        var v79 float64 = (1.5 * 1e10)
                        v80 := GenS4{F0: ((v76 + (v76 + ((v76 + "hello") + ("hello" + (((v76 + (v76 + v76)) + "a") + v76))))) + ("a" + v76)), F1: true}
                        _ = v79
                        _ = v80
                    }
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v74
                    _ = v75
                    _ = v76
                }
                v81 := GenS2{F0: false}
                v82 := []int{fi57}
                _ = v81
                _ = v82
            }
            v53 = 0
            _ = fi57
        }
        fmt.Println(v2, v3, v53)
        _ = v56
    }
    v83 := []int{v3, -9223372036854775808, (v3 * (v3 >> ((18446744073709551615) & 63)))}
    if true {
        v84 := GenS7{F0: 32767, F1: 100}
        _ = v84
    }
    if true {
        fmt.Println(v53, v53, v3)
        var v85 int64 = -1
        _ = v85
    } else {
        v86 := GenS8{F0: 432, F1: false, F2: 42}
        v83 = append(v83, -650)
        _ = v86
    }
    if len(v83) > 0 {
        v83[0] = (v2 / ((32768) | 1))
    }
    for fi87 := 0; fi87 < 4; fi87++ {
        for fi88 := 0; fi88 < 19; fi88++ {
            var v89 uint = 65535
            _ = fi88
            _ = v89
        }
        _ = fi87
    }
    fmt.Println(v53, v3)
    for fi90 := 0; fi90 < 2; fi90++ {
        v91 := GenS9{F0: 65535, F1: 128, F2: 1}
        var v92 float32 = (-2.25 + 0.1)
        _ = fi90
        _ = v91
        _ = v92
    }
    fmt.Println(v53, v53)
    fmt.Println(v53, v3, v53)
    if (1 > 1) {
        if len(v83) >= 2 {
            v83 = v83[1:len(v83)]
        }
        v93 := GenS7{F0: 887, F1: 1}
        _ = v93
    } else {
        v94 := GenS8{F0: 42, F1: false, F2: 551}
        _ = v94
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v83:", len(v83), v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v83
}
