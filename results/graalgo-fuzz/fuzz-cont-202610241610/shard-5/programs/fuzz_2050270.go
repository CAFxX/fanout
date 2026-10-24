package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int16; F1 uint; F2 int }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 bool; F1 uint32; F2 bool }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string; F1 int32 }
func genFn0() int {
    if (false && (!false)) {
        v4 := []int{759, 32768, 2147483647}
        v5 := GenS1{F0: -2147483648, F1: 4294967295}
        if (42 <= 0) {
            var v6 uint64 = 42
            var v7 float32 = -2.25
            if (1 >= 4294967295) {
                var v8 uint32 = 10
                _ = v8
            }
            _ = v6
            _ = v7
        } else {
            v9 := GenS2{F0: 18, F1: "hello", F2: -1}
            _ = v9
        }
        _ = v4
        _ = v5
    }
    var v10 uint64 = 658
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            if true {
                if ("a" >= (("\n\t" + "") + "hello")) {
                    v10 = 262
                }
            } else {
                fmt.Println(fi11, v10)
                if (-2.25 < (((1.5 - float32(3.14159)) + 3.14159) * (1e10 - float32((1.5 + float64(float32(float64(1.5)))))))) {
                    v12 := GenS2{F0: fi11, F1: "", F2: 1}
                    var v13 uint32 = 42
                    _ = v12
                    _ = v13
                }
                v14 := GenS3{F0: int16(v10), F1: true}
                _ = v14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                if (1 <= fi11) {
                    v16 := []int{(fi11 / (((fi11 + (fi15 / (((fi11 << ((0) & 63))) | 1)))) | 1)), (fi11 & (fi11 << ((65535) & 63)))}
                    v17 := GenS4{F0: -592, F1: 1, F2: fi11}
                    switch "a" {
                    case "":
                        v18 := GenS1{F0: int(v10), F1: (v10 >> ((841) & 63))}
                        _ = v18
                    case "\n\t":
                        fmt.Println(fi15, fi11, v10)
                    case "hello":
                        v10 -= (^(v10))
                    default:
                        v16 = append(v16, fi11)
                    }
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
            var v19 float64 = ((float64(0.0) / (0.1 - 1.5)) / 1.0)
            _ = v19
        } else {
            switch (("\n\t" + "hello") + "") {
            case "":
                v20 := []int{(fi11 ^ fi11), fi11, int(v10), fi11}
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 16; fi21++ {
                    if (!false) {
                        v22 := []int{32767, fi21}
                        var v23 bool = (!("hello" != "hello"))
                        _ = v22
                        _ = v23
                    } else {
                        switch fi21 {
                        case 127:
                            var v24 uint8 = 255
                            v25 := []int{fi11, (fi11 - (fi11 / ((461) | 1))), fi21}
                            _ = v24
                            _ = v25
                        }
                        v26 := []int{(fi21 - 4294967295), fi21, 0}
                        if (-32768 > 1000) {
                            _ = len(v26)
                            fmt.Println(v10)
                        } else {
                            _ = len(v26)
                            v10 = v10
                        }
                        _ = v26
                    }
                    if (1 != 124) {
                        v10 = (v10 - 42)
                        if true {
                            var v27 uint32 = 155
                            v10 = v10
                            var v28 uint = 410
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int8 = 100
                            _ = v29
                        }
                    }
                    _ = fi21
                }
            case "a":
                v30 := []int{}
                _ = v30
            default:
                fmt.Println(v10)
            }
            switch -2147483648 {
            case 127:
                for fi31 := 0; fi31 < 11; fi31++ {
                    v10 += 42
                    if true {
                        var v32 int16 = 42
                        _ = v32
                    } else {
                        v10 = (v10 + 363)
                    }
                    if ("foo bar" >= "\n\t") {
                        if (2 == 2) {
                            var v33 uint8 = 143
                            v34 := GenS5{F0: 100, F1: ("foo bar" + "a")}
                            _ = v33
                            _ = v34
                        } else {
                            switch "a" {
                            case "\n\t":
                                v10 = v10
                            case "foo bar":
                                var v35 uint8 = 13
                                _ = v35
                            case "":
                                v36 := []int{-2147483648, fi31, fi31, 32767}
                                fmt.Println(v10, v10)
                                _ = v36
                            default:
                                if false {
                                    v37 := GenS3{F0: -741, F1: ((false || (false && (369 != 1))) && true)}
                                    var v38 int16 = 836
                                    v10 = v10
                                    var v39 float64 = 1e10
                                    if (42 < 703) {
                                        v10 = v10
                                    } else {
                                        v40 := []int{(fi31 << ((216) & 63)), (fi11 / ((fi11) | 1)), int(v38), 32767}
                                        v10 = 1
                                        _ = v40
                                    }
                                    if true {
                                        v41 := []int{(fi11 + fi31)}
                                        v38 = (v38 / ((42) | 1))
                                        _ = v41
                                    } else {
                                        v42 := GenS5{F0: 80, F1: ""}
                                        v43 := []int{}
                                        if (100 <= -128) {
                                            switch -776 {
                                            case 17:
                                                var v44 uint64 = (v10 % ((0) | 1))
                                                _ = v44
                                            default:
                                                var v45 uint64 = (v10 * 755)
                                                _ = v45
                                            }
                                            v43 = append(v43, fi31)
                                            v10 = v10
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                        if (false && false) {
                            switch "hello" {
                            case "\n\t":
                                v10 = (v10 ^ (v10 | (v10 << ((552) & 63))))
                            case "a":
                                v46 := []int{fi31, (fi31 * (fi11 >> ((222) & 63)))}
                                _ = v46
                            }
                            v47 := GenS5{F0: -128, F1: ("\n\t" + "\n\t")}
                            _ = v47
                        } else {
                            v10 -= 1
                            var v48 int32 = 0
                            _ = v48
                        }
                    }
                    _ = fi31
                }
            case 4294967295:
                var v49 float32 = 1.5
                _ = v49
            default:
                for fi50 := 0; fi50 < 4; fi50++ {
                    v10 -= (v10 | v10)
                    var v51 int16 = 32767
                    _ = fi50
                    _ = v51
                }
            }
            var v52 uint64 = v10
            _ = v52
        }
        switch (fi11 | 2147483647) {
        case -2147483648:
            var v53 uint8 = 42
            _ = v53
        }
        _ = fi11
    }
    _ = v10
    return 127
}
func genFn1() uint16 {
    switch "\n\t" {
    case "foo bar":
        v54 := GenS6{F0: -32768, F1: -1}
        _ = v54
    case "a":
        for fi55 := 0; fi55 < 6; fi55++ {
            switch "hello" {
            case "\n\t":
                var v56 uint = 4294967295
                _ = v56
            case "":
                var v57 int32 = 2147483647
                _ = v57
            default:
                v58 := []int{(fi55 << ((559) & 63)), fi55, 32768}
                _ = v58
            }
            var v59 uint32 = 2
            v60 := []int{(fi55 << ((347) & 63))}
            _ = fi55
            _ = v59
            _ = v60
        }
    default:
        v61 := []int{}
        _ = v61
    }
    var v62 string = "hello"
    v63 := GenS7{F0: v62, F1: 0}
    for fi64 := 0; fi64 < 10; fi64++ {
        if true {
            if false {
                if (1 < 1) {
                    v65 := []int{fi64, genFn0(), 9223372036854775807, fi64}
                    v62 = v62
                    _ = v65
                } else {
                    fmt.Println(v62, v62)
                    var v66 int8 = 1
                    _ = v66
                }
                if ((!false) || (67 != 1)) {
                    v62 = v62
                    switch genFn0() {
                    case -1:
                        if false {
                            for fi67 := 0; fi67 < 11; fi67++ {
                                v68 := []int{(fi64 >> ((4294967295) & 63)), fi67}
                                if true {
                                    var v69 uint = 906
                                    if len(v68) > 0 {
                                        v68[0] = fi67
                                    }
                                    _ = v69
                                }
                                _ = fi67
                                _ = v68
                            }
                        }
                    }
                    if (true || (false || (132 < 2))) {
                        for fi70 := 0; fi70 < 16; fi70++ {
                            if (0 != -276) {
                                v62 = "a"
                            } else {
                                switch -1 {
                                case 1:
                                    if (fi70 <= (^(fi70))) {
                                        v62 = (v62 + (v62 + v62))
                                    } else {
                                        fmt.Println(fi64, fi70, fi70)
                                        v71 := GenS5{F0: 66, F1: v62}
                                        if ((1 >= -1) || ((!true) || (127 == 100))) {
                                            var v72 bool = (2 > 4294967295)
                                            _ = v72
                                        } else {
                                            var v73 string = v62
                                            _ = v73
                                        }
                                        _ = v71
                                    }
                                case 617:
                                    v62 = ("hello" + v62)
                                default:
                                    if ((!(v62 <= "a")) || false) {
                                        v62 = v62
                                        v74 := GenS8{F0: false, F1: 42, F2: false}
                                        _ = v74
                                    }
                                }
                            }
                            fmt.Println(v62, fi64)
                            switch ((((((v62 + ("hello" + "a")) + "hello") + "hello") + v62) + v62) + v62) {
                            case "a":
                                var v75 int64 = -85
                                _ = v75
                            case "\n\t":
                                v62 = ((v62 + "hello") + "a")
                            default:
                                v76 := []int{fi64}
                                _ = v76
                            }
                            _ = fi70
                        }
                        v62 = v62
                        v62 = ""
                    }
                }
            }
            for fi77 := 0; fi77 < 2; fi77++ {
                v62 = (v62 + v62)
                _ = fi77
            }
            switch "\n\t" {
            case "foo bar":
                v78 := GenS7{F0: v62, F1: 2}
                _ = v78
            }
        }
        v79 := GenS9{F0: "\n\t", F1: 128}
        _ = fi64
        _ = v79
    }
    var v80 uint64 = 18446744073709551615
    _ = v62
    _ = v63
    _ = v80
    return 910
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 840
    var v3 int64 = -1
    if false {
        var v81 int = 32768
        if false {
            var v82 int = 2147483647
            var v83 float32 = (0.1 * float32(0.1))
            var v84 string = "hello"
            _ = v82
            _ = v83
            _ = v84
        }
        if ((float32(float64((float32(-2.25) + (1e10 / -2.25)))) >= 0.1) || false) {
            for fi85 := 0; fi85 < 7; fi85++ {
                var v86 string = "foo bar"
                _ = fi85
                _ = v86
            }
            for fi87 := 0; fi87 < 3; fi87++ {
                if (0.0 >= (float32(((float64(((0.1 / 1.0) / (1e10 * (0.1 + float32((0.1 * (0.1 + 0.0))))))) / 1e10) / 1.0)) + -2.25)) {
                    switch ("foo bar" + "") {
                    case "hello":
                        v3 -= (v3 * v3)
                    case "\n\t":
                        for fi88 := 0; fi88 < 9; fi88++ {
                            v89 := []int{(v81 >> ((32) & 63)), (v81 - (fi88 % ((fi88) | 1))), genFn0()}
                            v2 = 32767
                            var v90 int16 = v2
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                    case "a":
                        var v91 float64 = float64(0.0)
                        _ = v91
                    default:
                        fmt.Println(v2)
                    }
                    var v92 uint16 = 813
                    var v93 uint16 = (v92 * v92)
                    _ = v92
                    _ = v93
                } else {
                    v2 = (v2 << ((255) & 63))
                }
                if ((v3 + 42) != v3) {
                    var v94 uint16 = 65535
                    v95 := []int{(v81 % (((v81 % (((v81 | genFn0())) | 1))) | 1)), (v81 - fi87), v81, 42}
                    var v96 float32 = (1e10 * (0.1 + -2.25))
                    _ = v94
                    _ = v95
                    _ = v96
                } else {
                    v3 = 1
                    v97 := v1[0]
                    fmt.Println(fi87)
                    _ = v97
                }
                if false {
                    v2 -= (-(v2))
                    var v98 int32 = 42
                    _ = v98
                }
                _ = fi87
            }
        } else {
            for fi99 := 0; fi99 < 19; fi99++ {
                for fi100 := 0; fi100 < 11; fi100++ {
                    var v101 int32 = -1
                    v102 := []int{(fi100 >> ((2) & 63)), 2147483647, (^(v81))}
                    v103 := GenS3{F0: v2, F1: (2147483647 < 136)}
                    _ = fi100
                    _ = v101
                    _ = v102
                    _ = v103
                }
                fmt.Println(v3, v3, v81)
                _ = fi99
            }
            var v104 int = v81
            _ = v104
        }
        _ = v81
    }
    v2 -= v2
    for fi105 := 0; fi105 < 7; fi105++ {
        var v106 int32 = 42
        v106 += (^(v106))
        _ = fi105
        _ = v106
    }
    if false {
        v107 := GenS10{F0: 42}
        _ = v107
    } else {
        var v108 string = ("" + (("hello" + "hello") + "foo bar"))
        _ = v108
    }
    var v109 int32 = -284
    if true {
        if (!false) {
            if (2 <= 42) {
                if true {
                    switch 32767 {
                    case -2147483648:
                        var v110 uint8 = 42
                        _ = v110
                    case 32:
                        v2 = v2
                    }
                } else {
                    var v111 float32 = ((3.14159 + ((1.5 / 1.0) - -2.25)) - 1.5)
                    _ = v111
                }
                v2 -= v2
            } else {
                var v112 float32 = 3.14159
                _ = v112
            }
            v2 = int16(v3)
            v113 := GenS11{F0: "hello", F1: v109}
            _ = v113
        } else {
            switch ("\n\t" + "foo bar") {
            case "hello":
                var v114 uint64 = 18446744073709551615
                _ = v114
            case "a":
                v115 := GenS10{F0: (v109 | (-(v109)))}
                _ = v115
            case "\n\t":
                if (255 <= 42) {
                    v116 := []int{}
                    _ = v116
                }
            }
            v117 := GenS9{F0: ("" + ""), F1: 2}
            for fi118 := 0; fi118 < 10; fi118++ {
                var v119 string = ("a" + "a")
                v120 := GenS1{F0: fi118, F1: 0}
                var v121 int = fi118
                _ = fi118
                _ = v119
                _ = v120
                _ = v121
            }
            _ = v117
        }
    }
    v122 := []int{}
    switch 837 {
    case -1:
        if false {
            for fi123 := 0; fi123 < 4; fi123++ {
                for fi124 := 0; fi124 < 15; fi124++ {
                    var v125 uint8 = 2
                    _ = fi124
                    _ = v125
                }
                delete(v1, 1)
                v126 := GenS3{F0: -1, F1: false}
                _ = fi123
                _ = v126
            }
        }
    case 1:
        v2 -= v2
    case 540:
        if false {
            v3 += -9223372036854775808
        } else {
            var v127 uint8 = 0
            _ = v127
        }
    default:
        if false {
            if len(v122) >= 2 {
                v122 = v122[1:len(v122)]
            }
        } else {
            var v128 string = "\n\t"
            if true {
                v129 := v1[42]
                v3 = v3
                _ = v129
            }
            v130 := GenS8{F0: false, F1: 42, F2: (v3 > (v3 * v3))}
            _ = v128
            _ = v130
        }
    }
    var v131 int32 = v109
    var v132 int64 = v3
    var v133 uint8 = 1
    var v134 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v109
    _ = v122
    _ = v131
    _ = v132
    _ = v133
    _ = v134
    _ = v2
    _ = v3
}
