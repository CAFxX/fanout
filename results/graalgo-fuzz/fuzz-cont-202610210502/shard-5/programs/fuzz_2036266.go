package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 int32 }
type GenS9 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 float32 = -2.25
    v4 := make(map[uint32]string)
    var v5 float32 = 0.0
    if true {
        switch (("a" + ("" + ("\n\t" + "a"))) + ("" + "")) {
        case "":
            for fi6 := 0; fi6 < 14; fi6++ {
                for fi7 := 0; fi7 < 6; fi7++ {
                    v5 -= v5
                    fmt.Println(fi6)
                    _ = fi7
                }
                _ = fi6
            }
        default:
            v8 := GenS1{F0: 592}
            _ = v8
        }
        var v9 int64 = 320
        _ = v9
    }
    fmt.Println(v5)
    var v10 int32 = 2147483647
    var v11 int64 = 884
    v10 -= v10
    v11 = (^(v11))
    switch "a" {
    case "foo bar":
        switch -236 {
        case 128:
            v3 -= 1e10
        case 32767:
            v5 += v3
        case -9223372036854775808:
            for fi12 := 0; fi12 < 8; fi12++ {
                v5 = 3.14159
                v11 += 1
                _ = fi12
            }
        }
    case "\n\t":
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                switch "hello" {
                case "a":
                    var v14 bool = false
                    _ = v14
                case "\n\t":
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        for fi15 := 0; fi15 < 4; fi15++ {
                            if true {
                                v16 := v1[(("hello" + (("a" + "a") + "a")) + ("" + (("\n\t" + "foo bar") + "")))]
                                _ = v16
                            }
                            fmt.Println(fi13)
                            _ = fi15
                        }
                    default:
                        if (!(!((!true) || false))) {
                            v11 = v11
                            var v17 int8 = 127
                            _ = v17
                        } else {
                            delete(v4, 42)
                            var v18 int32 = v10
                            switch "hello" {
                            case "hello":
                                if true {
                                    var v19 bool = false
                                    v20 := GenS1{F0: 42}
                                    for fi21 := 0; fi21 < 8; fi21++ {
                                        if (fi21 < fi21) {
                                            var v22 int64 = 42
                                            if (v19 && (true && false)) {
                                                fmt.Println(fi21, v22)
                                            } else {
                                                var v23 uint = 42
                                                var v24 uint32 = 0
                                                if (((-(fi21)) < 0) && (!((59 > 32767) && true))) {
                                                    if (v22 != 1) {
                                                        delete(v4, (v24 * v24))
                                                    }
                                                } else {
                                                    v25 := GenS2{F0: (-1 >= 1), F1: (v23 * (^(v23))), F2: "a"}
                                                    _ = v25
                                                }
                                                fmt.Println(fi21, v11)
                                                switch fi21 {
                                                case 1:
                                                    if ((v18 * (^(v18))) < (v10 ^ v10)) {
                                                        switch 127 {
                                                        case 127:
                                                            switch fi13 {
                                                            case 741:
                                                                v24 = (v24 / ((0) | 1))
                                                            }
                                                        case 0:
                                                            v26 := GenS3{F0: (v18 * 0), F1: 127}
                                                            _ = v26
                                                        case 32767:
                                                            var v27 uint16 = 255
                                                            _ = v27
                                                        default:
                                                            v28 := GenS2{F0: (211 >= 255), F1: 2, F2: "hello"}
                                                            _ = v28
                                                        }
                                                        v29 := GenS2{F0: (true && (-919 != (v11 - v11))), F1: (v23 | (v23 + (v23 * v23))), F2: "hello"}
                                                        _ = v29
                                                    }
                                                case 4294967295:
                                                    var v30 uint32 = (^(v24))
                                                    v31 := GenS4{F0: 128, F1: "foo bar", F2: (v23 - (^(v23)))}
                                                    _ = v30
                                                    _ = v31
                                                case 303:
                                                    v11 -= (v11 << ((2) & 63))
                                                default:
                                                    v32 := GenS4{F0: 2, F1: "foo bar", F2: (^(v23))}
                                                    _ = v32
                                                }
                                                switch ("" + ("" + ("\n\t" + ""))) {
                                                case "\n\t":
                                                    if false {
                                                        v24 = v24
                                                        switch fi21 {
                                                        case -238:
                                                            if true {
                                                                if v19 {
                                                                    v5 -= float32(-2.25)
                                                                    v23 += v23
                                                                }
                                                            } else {
                                                                var v33 uint16 = 65535
                                                                v18 = v18
                                                                _ = v33
                                                            }
                                                        }
                                                    } else {
                                                        v18 = v10
                                                        var v34 int8 = 1
                                                        _ = v34
                                                    }
                                                case "foo bar":
                                                    if v19 {
                                                        v35 := GenS2{F0: ((v18 >= (v18 >> ((v23) & 63))) || v19), F1: (v23 * 1), F2: "hello"}
                                                        _ = v35
                                                    } else {
                                                        if v19 {
                                                            v5 = (v5 / v3)
                                                            var v36 uint32 = 2147483647
                                                            _ = v36
                                                        } else {
                                                            var v37 uint64 = 1
                                                            var v38 int32 = v18
                                                            if (uint64(v23) == v37) {
                                                                v19 = v19
                                                            }
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        v39 := []int{(fi21 & 128), fi13, 69, (fi21 ^ (fi21 % ((1) | 1)))}
                                                        _ = v39
                                                    }
                                                case "a":
                                                    fmt.Println(fi21, fi13)
                                                default:
                                                    v40 := []int{fi13}
                                                    _ = v40
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                            var v41 bool = false
                                            _ = v22
                                            _ = v41
                                        }
                                        if false {
                                            if (4294967295 <= 4294967295) {
                                                var v42 int16 = 42
                                                fmt.Println(v19, v3)
                                                fmt.Println(v11, v5)
                                                _ = v42
                                            } else {
                                                switch (("hello" + "a") + ("\n\t" + ((("foo bar" + "a") + ("\n\t" + "hello")) + "a"))) {
                                                case "":
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "hello":
                                                            switch ("" + "\n\t") {
                                                            case "":
                                                                var v43 int64 = v11
                                                                _ = v43
                                                            case "a":
                                                                fmt.Println(fi13, fi21, v3)
                                                            case "\n\t":
                                                                v10 += (v18 / ((0) | 1))
                                                            }
                                                        case "foo bar":
                                                            v11 = 42
                                                        }
                                                    case "a":
                                                        v44 := GenS1{F0: 853}
                                                        _ = v44
                                                    case "\n\t":
                                                        v11 -= 9223372036854775807
                                                    default:
                                                        v3 += v5
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        var v45 float64 = 1e10
                                                        if (!false) {
                                                            switch fi21 {
                                                            case 127:
                                                                var v46 int32 = v10
                                                                _ = v46
                                                            case 42:
                                                                v47 := GenS5{F0: -104, F1: -41}
                                                                _ = v47
                                                            }
                                                        }
                                                        var v48 int64 = 0
                                                        _ = v45
                                                        _ = v48
                                                    } else {
                                                        fmt.Println(fi21, v11)
                                                        var v49 int64 = (v11 | (v11 - -1))
                                                        _ = v49
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi21
                                    }
                                    _ = v19
                                    _ = v20
                                }
                            case "foo bar":
                                v11 = (^(v11))
                            }
                            _ = v18
                        }
                    }
                case "foo bar":
                    var v50 float32 = v3
                    _ = v50
                }
                for fi51 := 0; fi51 < 10; fi51++ {
                    var v52 int = (-(fi13))
                    if true {
                        switch "\n\t" {
                        case "hello":
                            v53 := []int{fi51, fi51, (fi13 << ((42) & 63)), fi13}
                            v54 := GenS5{F0: 1, F1: 127}
                            var v55 bool = true
                            _ = v53
                            _ = v54
                            _ = v55
                        case "\n\t":
                            var v56 string = "hello"
                            _ = v56
                        case "a":
                            v57 := GenS2{F0: true, F1: 65535, F2: "hello"}
                            _ = v57
                        }
                        switch "" {
                        case "":
                            fmt.Println(fi13, v11)
                        case "foo bar":
                            var v58 int16 = -818
                            _ = v58
                        case "\n\t":
                            var v59 uint16 = 1
                            _ = v59
                        }
                        v60 := GenS1{F0: 118}
                        _ = v60
                    } else {
                        if false {
                            var v61 int = fi51
                            _ = v61
                        }
                        v1["hello"] = ""
                        if true {
                            var v62 uint64 = 2
                            _ = v62
                        } else {
                            v63 := []int{int(v10), (-(fi51)), (-(fi51)), (v52 | -2147483648)}
                            v64 := GenS4{F0: 1, F1: "\n\t", F2: 228}
                            v5 += (1e10 - ((v5 / v5) * (v5 - v3)))
                            _ = v63
                            _ = v64
                        }
                    }
                    _ = fi51
                    _ = v52
                }
                v65 := GenS3{F0: (v10 + (-(v10))), F1: 255}
                _ = v65
            }
            var v66 uint32 = 669
            for fi67 := 0; fi67 < 3; fi67++ {
                var v68 int = 923
                if false {
                    fmt.Println(v11, v5, v66)
                } else {
                    if (v10 == v10) {
                        switch fi13 {
                        case -9223372036854775808:
                            if false {
                                if (!(float64(((v3 / v3) + (-2.25 / 0.1))) > (1.5 / ((0.1 + 1e10) / 1e10)))) {
                                    if (127 < 127) {
                                        fmt.Println(v10, v10)
                                    } else {
                                        var v69 bool = (true && (!false))
                                        _ = v69
                                    }
                                    var v70 uint32 = v66
                                    _ = v70
                                }
                                if (0.0 > 1.5) {
                                    if ((v66 << ((1) & 63)) <= (v66 << ((65535) & 63))) {
                                        v68 = (fi13 / (((fi67 << ((42) & 63))) | 1))
                                        v71 := GenS6{F0: 21, F1: 1}
                                        var v72 float64 = 0.0
                                        _ = v71
                                        _ = v72
                                    } else {
                                        fmt.Println(fi13, fi67, v68)
                                        if ((0 > 18446744073709551615) || true) {
                                            v73 := GenS7{F0: -128}
                                            v4[65535] = "a"
                                            _ = v73
                                        } else {
                                            v3 = float32(1.5)
                                            v68 -= fi13
                                            v3 = float32((1.5 - ((float64(((3.14159 * v3) * v3)) - (-2.25 + 0.0)) - 0.1)))
                                        }
                                    }
                                }
                            } else {
                                if false {
                                    fmt.Println(v5)
                                    v68 -= 358
                                } else {
                                    fmt.Println(v5, v10, fi67)
                                    var v74 uint16 = 992
                                    _ = v74
                                }
                                v75 := GenS1{F0: 4294967295}
                                var v76 uint16 = 605
                                _ = v75
                                _ = v76
                            }
                            fmt.Println(fi13)
                        case 2147483647:
                            if (1 == 0) {
                                v11 = -1
                                switch ((("a" + "\n\t") + "\n\t") + (("\n\t" + "\n\t") + "")) {
                                case "a":
                                    v77 := GenS6{F0: 2, F1: 477}
                                    if (4294967295 <= (v66 + v66)) {
                                        var v78 uint32 = (v66 + 97)
                                        v79 := []int{42}
                                        _ = v78
                                        _ = v79
                                    } else {
                                        v10 -= v10
                                        v4[(^(v66))] = "hello"
                                        v5 += v5
                                    }
                                    _ = v77
                                case "":
                                    switch (fi67 + (fi67 & (v68 ^ v68))) {
                                    case 32768:
                                        v80 := GenS8{F0: 255, F1: -31, F2: v10}
                                        _ = v80
                                    case -747:
                                        delete(v4, (^(v66)))
                                    default:
                                        if false {
                                            if ((v11 >> ((255) & 63)) < v11) {
                                                v10 = (v10 - v10)
                                            }
                                            v81 := GenS9{F0: 1, F1: v66}
                                            _ = v81
                                        }
                                    }
                                }
                            }
                        case 9223372036854775807:
                            var v82 string = ""
                            _ = v82
                        default:
                            v5 += v3
                        }
                    }
                    var v83 string = ("" + (("hello" + ((("hello" + "a") + "foo bar") + ("hello" + (((("a" + ("\n\t" + "a")) + ("" + ("hello" + ("\n\t" + "hello")))) + "foo bar") + "foo bar")))) + "hello"))
                    _ = v83
                }
                v84 := v2[-9223372036854775808]
                _ = fi67
                _ = v68
                _ = v84
            }
            _ = fi13
            _ = v66
        }
    case "a":
        var v85 int8 = 100
        _ = v85
    }
    var v86 uint16 = 171
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
}
