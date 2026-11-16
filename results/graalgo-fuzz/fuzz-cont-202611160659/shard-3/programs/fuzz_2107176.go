package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }
type GenS5 struct { F0 string; F1 string; F2 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int64 }

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    var v3 int8 = 73
    var v4 float32 = 1e10
    var v5 uint = 0
    var v6 int8 = -14
    v1 = (!v1)
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    switch "" {
    case "a":
        if (42 == 1) {
            v1 = (131 >= 255)
            _ = len(v2)
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9, v1)
                _ = fi9
            }
        }
    }
    v1 = (128 > 128)
    v10 := GenS1{F0: 128, F1: 607, F2: 1}
    v11 := GenS2{F0: 4294967295}
    v12 := v2[-9223372036854775808]
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int = (fi13 ^ (-(fi13)))
        switch (("hello" + "\n\t") + ((((("\n\t" + ("" + "foo bar")) + "a") + "hello") + "a") + "a")) {
        case "":
            if (-2.25 != 1e10) {
                v15 := GenS2{F0: v5}
                v7 = -509
                _ = v15
            } else {
                var v16 int8 = (v3 % ((v6) | 1))
                var v17 int64 = 788
                var v18 int16 = -751
                _ = v16
                _ = v17
                _ = v18
            }
        case "foo bar":
            var v19 float32 = (0.0 / (0.1 * v4))
            _ = v19
        case "\n\t":
            var v20 int64 = 1
            _ = v20
        default:
            if ((v4 - 0.0) > (v4 * v4)) {
                v21 := GenS3{F0: 42}
                fmt.Println(v1)
                if true {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        switch int(v5) {
                        case 42:
                            var v23 uint = v5
                            _ = v23
                        case 9223372036854775807:
                            v24 := []int{v14, v14, fi13, (fi22 << (((v5 >> ((v5) & 63))) & 63))}
                            _ = v24
                        case -60:
                            if v12 {
                                v25 := GenS1{F0: 233, F1: 725, F2: 2147483648}
                                var v26 uint32 = 170
                                v27 := []int{fi13, (v14 - 531), (^(fi13))}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                v3 += (v3 % (((v6 & -128)) | 1))
                                var v28 float64 = (3.14159 * (1.5 - 1.5))
                                _ = v28
                            }
                        }
                        _ = fi22
                    }
                    _ = len(v2)
                    v7 -= v7
                } else {
                    for fi29 := 0; fi29 < 1; fi29++ {
                        if ((v4 + float32(1e10)) > v4) {
                            if (fi13 != 32768) {
                                if (!v1) {
                                    v7 = v7
                                    v1 = false
                                    fmt.Println(v6, v12)
                                    switch 32768 {
                                    case 9223372036854775807:
                                        fmt.Println(v6, v12)
                                    case 1:
                                        if ((fi13 >> (((v5 - v5)) & 63)) < (v14 % ((fi29) | 1))) {
                                            _ = len(v2)
                                            v30 := GenS1{F0: 2, F1: 1, F2: 825}
                                            switch -1 {
                                            case 87:
                                                var v31 uint64 = 1
                                                _ = v31
                                            case -9223372036854775808:
                                                if (v4 <= v4) {
                                                    var v32 uint32 = 0
                                                    var v33 int8 = 42
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    if (v12 && ("" <= "hello")) {
                                                        v14 = (v14 << ((v5) & 63))
                                                        if v1 {
                                                            var v34 float64 = 0.1
                                                            v3 -= v3
                                                            var v35 uint8 = 255
                                                            _ = v34
                                                            _ = v35
                                                        } else {
                                                            v14 = v14
                                                        }
                                                    }
                                                    _ = len(v2)
                                                }
                                            case 1:
                                                var v36 string = "foo bar"
                                                _ = v36
                                            default:
                                                if (0 == 128) {
                                                    v14 -= -9223372036854775808
                                                    v37 := GenS1{F0: 136, F1: 686, F2: 2147483648}
                                                    if (288 != 18446744073709551615) {
                                                        if v1 {
                                                            v38 := GenS2{F0: (v5 & (^(v5)))}
                                                            fmt.Println(v6)
                                                            _ = v38
                                                        }
                                                        if (1e10 >= ((v4 - v4) / 3.14159)) {
                                                            v5 = (v5 % ((244) | 1))
                                                            switch ((("\n\t" + "") + "hello") + "foo bar") {
                                                            case "\n\t":
                                                                var v39 uint16 = 812
                                                                _ = v39
                                                            case "":
                                                                v40 := v2[(^(fi13))]
                                                                _ = v40
                                                            case "hello":
                                                                v7 = -195
                                                            }
                                                        }
                                                    }
                                                    v41 := GenS3{F0: 205}
                                                    _ = v37
                                                    _ = v41
                                                } else {
                                                    if (true || false) {
                                                        if ((v1 || true) && (!((!v1) && ((v7 | (-(v7))) == (v7 ^ (v7 ^ (v7 - v7))))))) {
                                                            v42 := GenS4{F0: (fi29 % (((^(fi13))) | 1)), F1: 1, F2: 42}
                                                            v43 := GenS3{F0: 255}
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            v44 := GenS5{F0: "foo bar", F1: "\n\t", F2: 4294967295}
                                                            _ = v44
                                                        }
                                                        fmt.Println(v5, v12, v1)
                                                    }
                                                }
                                            }
                                            delete(v2, fi13)
                                            fmt.Println(v5, v7, v6)
                                            _ = v30
                                        }
                                    }
                                }
                            }
                            v45 := GenS5{F0: ("foo bar" + (("a" + "hello") + "hello")), F1: ("a" + ("" + "foo bar")), F2: 0}
                            _ = v45
                        }
                        v12 = true
                        _ = fi29
                    }
                }
                _ = v21
            }
        }
        if (2147483647 <= 783) {
            var v46 int8 = (v6 % ((v6) | 1))
            for fi47 := 0; fi47 < 18; fi47++ {
                switch ("a" + "foo bar") {
                case "a":
                    v12 = ((v4 - 3.14159) > (v4 * v4))
                case "":
                    if (v1 && (true || true)) {
                        if (((571 == 2147483648) || (255 < v5)) || true) {
                            var v48 uint = v5
                            _ = v48
                        }
                        var v49 float64 = 1.5
                        var v50 uint16 = 511
                        v5 += 18446744073709551615
                        _ = v49
                        _ = v50
                    } else {
                        if true {
                            var v51 float32 = float32(0.1)
                            _ = v51
                        }
                        if false {
                            v52 := v2[(fi13 - 42)]
                            switch fi13 {
                            case 128:
                                v6 -= (^(v6))
                            case 4294967295:
                                v12 = false
                            }
                            fmt.Println(v7, v46)
                            v12 = (!false)
                            if false {
                                var v53 int = fi47
                                fmt.Println(v52, v5)
                                _ = v53
                            }
                            var v54 int16 = 32767
                            _ = v52
                            _ = v54
                        }
                        v3 -= v3
                        v55 := []int{(fi47 | (fi13 - (fi47 & fi13))), (-(v14))}
                        _ = v55
                    }
                case "foo bar":
                    v12 = (2 > 413)
                default:
                    var v56 uint = 4294967295
                    _ = len(v2)
                    v1 = true
                    _ = v56
                }
                _ = len(v2)
                if v1 {
                    switch (fi13 | fi13) {
                    case 127:
                        fmt.Println(v6)
                    case -9223372036854775808:
                        v7 = (v7 * (v7 + 42))
                    case 4294967295:
                        var v57 int8 = v3
                        _ = v57
                    }
                }
                _ = fi47
            }
            _ = v46
        }
        _ = fi13
        _ = v14
    }
    v58 := GenS5{F0: "hello", F1: "", F2: 589}
    var v59 int16 = 148
    v59 -= (v59 - v59)
    var v60 string = "foo bar"
    var v61 int64 = 42
    if (233 <= 0) {
        if (false || v1) {
            v62 := GenS6{F0: (v59 % ((v59) | 1))}
            if v12 {
                v63 := []int{32767, 128, -1, -1}
                v6 = -8
                v60 = v60
                _ = v63
            } else {
                var v64 uint8 = 42
                var v65 uint32 = 377
                switch 32768 {
                case 2147483647:
                    v1 = (v7 <= 1)
                case 32767:
                    v66 := []int{2147483647}
                    _ = v66
                }
                _ = v64
                _ = v65
            }
            v67 := GenS7{F0: 290, F1: -36}
            _ = v62
            _ = v67
        } else {
            if (!(1 > 127)) {
                var v68 uint16 = 255
                fmt.Println(v3)
                _ = v68
            } else {
                if false {
                    switch int(v5) {
                    case 9223372036854775807:
                        if v12 {
                            var v69 string = ("" + "a")
                            _ = v69
                        } else {
                            v59 += v59
                            var v70 uint8 = 8
                            _ = v70
                        }
                    case -2147483648:
                        fmt.Println(v61)
                    case -855:
                        v71 := []int{4294967295, 0}
                        _ = v71
                    default:
                        v6 += v3
                    }
                } else {
                    var v72 uint = (v5 * v5)
                    var v73 bool = ((-(v7)) < v7)
                    for fi74 := 0; fi74 < 11; fi74++ {
                        if (2 != 1) {
                            switch "" {
                            case "a":
                                if (v4 != 0.0) {
                                    if (100 != 101) {
                                        fmt.Println(v60, v7, v12)
                                    }
                                    var v75 uint = (v5 & v5)
                                    _ = v75
                                } else {
                                    fmt.Println(v1, fi74)
                                    v76 := GenS5{F0: (v60 + (v60 + v60)), F1: v60, F2: uint64(v5)}
                                    _ = v76
                                }
                            default:
                                v4 -= v4
                            }
                            for fi77 := 0; fi77 < 6; fi77++ {
                                v72 += v72
                                _ = fi77
                            }
                        } else {
                            v78 := GenS3{F0: 102}
                            for fi79 := 0; fi79 < 9; fi79++ {
                                if (-9223372036854775808 != (v61 & v61)) {
                                    v60 = v60
                                    v2[1] = v12
                                    switch (v60 + "foo bar") {
                                    case "":
                                        fmt.Println(v5, v60)
                                    default:
                                        switch v60 {
                                        case "hello":
                                            v72 += v5
                                        case "foo bar":
                                            var v80 uint16 = 1
                                            _ = v80
                                        case "a":
                                            var v81 uint8 = 2
                                            _ = v81
                                        default:
                                            v82 := []int{}
                                            _ = v82
                                        }
                                    }
                                } else {
                                    v83 := []int{}
                                    _ = v83
                                }
                                fmt.Println(v1, v1)
                                if v12 {
                                    if (!(!v1)) {
                                        fmt.Println(v59)
                                    } else {
                                        var v84 string = "foo bar"
                                        switch (-(fi79)) {
                                        case 565:
                                            var v85 uint8 = 1
                                            _ = v85
                                        case -2147483648:
                                            v86 := GenS4{F0: fi74, F1: 2, F2: -113}
                                            _ = v86
                                        case 127:
                                            if v12 {
                                                v87 := GenS2{F0: 558}
                                                if (-9223372036854775808 == (fi79 ^ 2147483647)) {
                                                    var v88 string = "hello"
                                                    _ = len(v2)
                                                    v89 := GenS2{F0: (^(v5))}
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    var v90 uint32 = 400
                                                    _ = v90
                                                }
                                                _ = v87
                                            } else {
                                                var v91 uint = (v72 * v5)
                                                v92 := GenS5{F0: ("hello" + v84), F1: v84, F2: 1}
                                                _ = v91
                                                _ = v92
                                            }
                                        default:
                                            v60 = "a"
                                        }
                                        v93 := GenS7{F0: fi74, F1: -786}
                                        _ = v84
                                        _ = v93
                                    }
                                    switch v60 {
                                    case "a":
                                        v12 = false
                                    }
                                    v94 := []int{int(v61), (fi74 - fi74), (fi74 >> ((1) & 63))}
                                    _ = v94
                                } else {
                                    if true {
                                        fmt.Println(v6, fi74)
                                        v12 = (v60 != ("" + "hello"))
                                    } else {
                                        v1 = v73
                                    }
                                }
                                _ = fi79
                            }
                            var v95 uint8 = 128
                            _ = v78
                            _ = v95
                        }
                        v5 = (v72 / ((v72) | 1))
                        var v96 int = fi74
                        _ = fi74
                        _ = v96
                    }
                    _ = v72
                    _ = v73
                }
                fmt.Println(v6, v7, v6)
                for fi97 := 0; fi97 < 3; fi97++ {
                    var v98 float64 = (3.14159 + (((1e10 * 1e10) - 1.5) * (-2.25 - float64(float32(0.0)))))
                    v99 := []int{fi97, 32767, int(v6)}
                    v61 -= (v61 % ((v61) | 1))
                    _ = fi97
                    _ = v98
                    _ = v99
                }
            }
            switch ((v60 + "") + v60) {
            case "foo bar":
                var v100 uint32 = 516
                _ = v100
            }
            fmt.Println(v12, v59, v7)
        }
    } else {
        if (!(v60 == ("" + v60))) {
            fmt.Println(v7)
            var v101 uint16 = 253
            _ = v101
        } else {
            if v1 {
                var v102 int16 = (-(v59))
                var v103 bool = ((v59 << ((4294967295) & 63)) >= v102)
                var v104 int = int(v6)
                _ = v102
                _ = v103
                _ = v104
            } else {
                v105 := []int{}
                _ = v105
            }
        }
        fmt.Println(v4, v60)
        v106 := []int{127, 459}
        _ = v106
    }
    v2[1] = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v7
}
