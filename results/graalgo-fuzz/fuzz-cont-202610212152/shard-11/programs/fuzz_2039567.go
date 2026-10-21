package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 string; F2 string }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 string }
func genFn0() uint64 {
    if true {
        v6 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v6
    }
    v7 := []int{4294967295, -9223372036854775808, 1, 4294967295}
    v8 := GenS1{F0: 36, F1: 32767}
    if false {
        var v9 int16 = 231
        if false {
            var v10 int64 = -652
            _ = v10
        } else {
            if true {
                _ = len(v7)
                fmt.Println(v9)
            } else {
                var v11 bool = false
                if v11 {
                    v11 = (4294967295 > -1)
                    var v12 string = "foo bar"
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v9 -= 42
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint64 = 42
                            switch (v12 + (("a" + (v12 + v12)) + (("a" + v12) + v12))) {
                            case "foo bar":
                                v15 -= (v15 >> ((18446744073709551615) & 63))
                                v16 := GenS2{F0: int16(fi14)}
                                if v11 {
                                    if v11 {
                                        v9 -= (v9 | -243)
                                        v11 = (65535 <= 2147483647)
                                    } else {
                                        v17 := GenS1{F0: -1, F1: (-(fi14))}
                                        if (540 == 324) {
                                            v18 := GenS1{F0: 100, F1: (fi14 & fi13)}
                                            switch v12 {
                                            case "hello":
                                                fmt.Println(v11, v12, v12)
                                            default:
                                                switch v12 {
                                                case "hello":
                                                    v7 = append(v7, (fi13 - (-(fi13))))
                                                case "\n\t":
                                                    var v19 uint16 = 805
                                                    _ = len(v7)
                                                    if (v11 || false) {
                                                        var v20 int32 = 1
                                                        _ = v20
                                                    } else {
                                                        switch (v12 + ("hello" + v12)) {
                                                        case "":
                                                            v7 = append(v7, (^(fi14)))
                                                        case "a":
                                                            var v21 uint = 255
                                                            _ = v21
                                                        case "\n\t":
                                                            v11 = v11
                                                        default:
                                                            v7 = append(v7, 42)
                                                        }
                                                        v9 = -1
                                                        switch (-(fi14)) {
                                                        case 746:
                                                            var v22 bool = (!false)
                                                            _ = v22
                                                        case -8:
                                                            v23 := GenS3{F0: (-75 >= -1)}
                                                            _ = v23
                                                        case -996:
                                                            var v24 float64 = (float64(1.5) + 0.1)
                                                            _ = v24
                                                        }
                                                    }
                                                    _ = v19
                                                default:
                                                    var v25 int8 = 127
                                                    _ = v25
                                                }
                                                fmt.Println(v11, v12)
                                            }
                                            _ = v18
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = fi13
                                        }
                                        _ = v17
                                    }
                                } else {
                                    v11 = (1 > 110)
                                }
                                _ = v16
                            case "hello":
                                v9 -= -1
                            default:
                                var v26 bool = (!(v11 && v11))
                                _ = v26
                            }
                            if (1e10 >= (3.14159 / -2.25)) {
                                switch (fi13 * -65) {
                                case 32767:
                                    if len(v7) > 0 {
                                        v7[0] = 128
                                    }
                                }
                                if (0 <= 255) {
                                    if ((184 != -802) || v11) {
                                        v11 = (v11 || v11)
                                        v27 := GenS4{F0: 0, F1: 2}
                                        fmt.Println(fi14)
                                        _ = v27
                                    }
                                }
                            }
                            v15 -= (v15 >> ((1) & 63))
                            v28 := GenS3{F0: (false && (0 != 127))}
                            v29 := GenS4{F0: -1, F1: 255}
                            _ = fi14
                            _ = v15
                            _ = v28
                            _ = v29
                        }
                        _ = fi13
                    }
                    _ = v12
                }
                switch "hello" {
                case "a":
                    v11 = (-2147483648 <= -578)
                }
                _ = v11
            }
            if false {
                fmt.Println(v9)
                v7 = append(v7, 2147483647)
                if true {
                    for fi30 := 0; fi30 < 14; fi30++ {
                        switch "\n\t" {
                        case "a":
                            var v31 string = "hello"
                            _ = v31
                        case "foo bar":
                            v32 := GenS5{F0: 1}
                            _ = v32
                        case "hello":
                            v9 -= v9
                        }
                        _ = fi30
                    }
                } else {
                    var v33 float32 = float32(3.14159)
                    _ = v33
                }
            } else {
                if len(v7) > 0 {
                    v7[0] = 128
                }
                v9 += v9
                var v34 bool = false
                _ = v34
            }
            var v35 bool = false
            _ = v35
        }
        _ = v9
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = v7
    _ = v8
    return 2
}
func genFn1() int64 {
    switch 130 {
    case -9223372036854775808:
        var v36 int16 = 1
        _ = v36
    case 42:
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := []int{(^(fi37)), fi37, (-(fi37)), (fi37 % ((1) | 1))}
            if (179 > 128) {
                var v39 int8 = -1
                _ = v39
            } else {
                if true {
                    fmt.Println(fi37)
                    switch -9223372036854775808 {
                    case 857:
                        if (349 >= 341) {
                            v40 := GenS4{F0: 2147483647, F1: 1}
                            fmt.Println(fi37)
                            var v41 int32 = -2147483648
                            _ = v40
                            _ = v41
                        }
                    case 42:
                        v42 := GenS5{F0: 257}
                        _ = v42
                    case 2147483647:
                        for fi43 := 0; fi43 < 7; fi43++ {
                            var v44 uint64 = genFn0()
                            v44 -= (^(v44))
                            var v45 uint32 = 563
                            v45 = (v45 * v45)
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    default:
                        if (4294967295 > 4294967295) {
                            var v46 string = "foo bar"
                            if false {
                                if ((2 > 182) || true) {
                                    var v47 uint64 = 121
                                    fmt.Println(v46, v46)
                                    _ = v47
                                }
                                switch (("foo bar" + "a") + v46) {
                                case "foo bar":
                                    var v48 int8 = 0
                                    _ = v48
                                case "\n\t":
                                    v49 := GenS1{F0: -1, F1: (fi37 | (fi37 - 128))}
                                    _ = v49
                                default:
                                    for fi50 := 0; fi50 < 15; fi50++ {
                                        v46 = ("a" + (v46 + v46))
                                        if len(v38) > 0 {
                                            v38[0] = 32768
                                        }
                                        var v51 int16 = 1
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            } else {
                                var v52 int64 = 42
                                for fi53 := 0; fi53 < 17; fi53++ {
                                    v54 := GenS4{F0: 892, F1: 42}
                                    _ = fi53
                                    _ = v54
                                }
                                _ = len(v38)
                                _ = v52
                            }
                            _ = v46
                        }
                    }
                } else {
                    if (1.5 == (-2.25 - 1.5)) {
                        for fi55 := 0; fi55 < 2; fi55++ {
                            var v56 uint16 = 255
                            _ = fi55
                            _ = v56
                        }
                        var v57 int8 = 42
                        _ = v57
                    }
                    fmt.Println(fi37)
                    for fi58 := 0; fi58 < 14; fi58++ {
                        if true {
                            if (int64(fi37) < 0) {
                                var v59 bool = false
                                _ = v59
                            }
                            if len(v38) > 0 {
                                v38[0] = fi58
                            }
                            switch ("hello" + "a") {
                            case "foo bar":
                                v60 := GenS4{F0: -1, F1: 492}
                                _ = v60
                            default:
                                var v61 int32 = -484
                                _ = v61
                            }
                        } else {
                            if (-1 <= 9223372036854775807) {
                                var v62 int = (^(fi37))
                                var v63 int32 = 0
                                _ = v62
                                _ = v63
                            }
                            fmt.Println(fi37, fi37)
                        }
                        _ = fi58
                    }
                }
                switch fi37 {
                case 150:
                    _ = len(v38)
                default:
                    var v64 uint16 = 1
                    _ = v64
                }
                switch "a" {
                case "foo bar":
                    v38 = append(v38, 32768)
                case "hello":
                    var v65 bool = false
                    _ = v65
                case "a":
                    for fi66 := 0; fi66 < 16; fi66++ {
                        _ = len(v38)
                        var v67 uint64 = 18446744073709551615
                        fmt.Println(fi66, fi37, v67)
                        _ = fi66
                        _ = v67
                    }
                default:
                    var v68 int16 = 0
                    _ = v68
                }
            }
            var v69 int = (fi37 << ((972) & 63))
            _ = fi37
            _ = v38
            _ = v69
        }
    case 128:
        for fi70 := 0; fi70 < 2; fi70++ {
            switch "\n\t" {
            case "\n\t":
                v71 := GenS6{F0: 2, F1: "a", F2: ("hello" + "hello")}
                _ = v71
            }
            var v72 uint = 18446744073709551615
            v73 := []int{}
            _ = fi70
            _ = v72
            _ = v73
        }
    }
    switch -2147483648 {
    case 450:
        v74 := []int{32767, 1, -1, 127}
        _ = v74
    case 32767:
        var v75 uint16 = 550
        _ = v75
    default:
        var v76 uint16 = 744
        _ = v76
    }
    var v77 int32 = 1
    if (1 < 42) {
        var v78 float32 = ((0.0 / 1e10) * (1e10 - float32(0.0)))
        _ = v78
    }
    if true {
        if true {
            if (("foo bar" + "\n\t") >= "") {
                v77 = (v77 ^ v77)
                if ((0 <= 32767) || false) {
                    var v79 int32 = v77
                    fmt.Println(v77, v79)
                    var v80 int16 = int16(v79)
                    _ = v79
                    _ = v80
                }
            } else {
                var v81 int32 = (v77 & (v77 + (v77 >> ((705) & 63))))
                var v82 int16 = -1
                v82 = v82
                _ = v81
                _ = v82
            }
        }
        v83 := []int{32768}
        v83 = append(v83, 0)
        _ = v83
    }
    _ = v77
    return -806
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]uint32)
    var v3 int8 = -128
    var v4 string = "a"
    var v5 int16 = 1000
    fmt.Println(v4, v1)
    switch v4 {
    case "foo bar":
        v84 := GenS6{F0: 5, F1: "foo bar", F2: "a"}
        _ = v84
    }
    var v85 float64 = float64((float32(v1) * ((float32(-2.25) - 1.5) / ((((3.14159 * ((((3.14159 + (float32((v1 / (0.1 - (v1 - v1)))) / 1.0)) - float32(v1)) + 0.1) - 0.0)) * 0.1) / float32(-2.25)) - 1e10))))
    if (0 <= 2147483647) {
        v86 := GenS3{F0: true}
        if false {
            if false {
                _ = len(v2)
                switch 9223372036854775807 {
                case -9223372036854775808:
                    for fi87 := 0; fi87 < 13; fi87++ {
                        for fi88 := 0; fi88 < 10; fi88++ {
                            var v89 int8 = v3
                            v90 := []int{}
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                        var v91 uint8 = 128
                        v85 = v1
                        _ = fi87
                        _ = v91
                    }
                }
                for fi92 := 0; fi92 < 11; fi92++ {
                    for fi93 := 0; fi93 < 1; fi93++ {
                        _ = len(v2)
                        if false {
                            var v94 uint = 42
                            switch (v4 + "") {
                            case "hello":
                                var v95 uint = (^(v94))
                                _ = v95
                            case "foo bar":
                                _ = len(v2)
                            }
                            v96 := []int{(fi92 ^ (fi92 + 0)), fi92}
                            _ = v94
                            _ = v96
                        }
                        if (!true) {
                            fmt.Println(v85, v5, v3)
                            fmt.Println(v5)
                        } else {
                            fmt.Println(v4, v5, fi93)
                            v97 := GenS7{F0: fi93, F1: v5}
                            _ = v97
                        }
                        _ = fi93
                    }
                    _ = fi92
                }
            }
            v3 = -44
        }
        for fi98 := 0; fi98 < 12; fi98++ {
            v99 := GenS4{F0: -745, F1: 570}
            _ = fi98
            _ = v99
        }
        _ = v86
    } else {
        var v100 uint16 = 1
        _ = v100
    }
    var v101 bool = false
    if (!(!(-324 <= 42))) {
        var v102 uint8 = 42
        v103 := GenS8{F0: 1, F1: 4294967295}
        fmt.Println(v5)
        _ = v102
        _ = v103
    }
    fmt.Println(v5)
    v104 := GenS4{F0: 1, F1: 1}
    v105 := GenS9{F0: "a"}
    v5 = v5
    v101 = false
    for fi106 := 0; fi106 < 8; fi106++ {
        var v107 string = v4
        v4 = v4
        _ = fi106
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v101
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
