package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int64; F1 int64; F2 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 int32; F2 uint16 }
type GenS11 struct { F0 uint8; F1 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int16; F1 int16 }
func genFn0() string {
    v4 := GenS1{F0: 9223372036854775807}
    var v5 uint = 1
    v5 -= v5
    v6 := []int{9223372036854775807, -964, 42, 32768}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint64 = 1000
        var v9 int64 = 9223372036854775807
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 int = -639
    var v2 int32 = -204
    v3 := []uint32{42, 721, 2147483648}
    fmt.Println(v1, v1)
    v10 := GenS1{F0: v1}
    v11 := GenS2{F0: -32768, F1: v1, F2: 18446744073709551615}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint = 42
            switch "foo bar" {
            case "a":
                v15 := GenS1{F0: 32767}
                _ = v15
            case "hello":
                var v16 uint32 = 4294967295
                _ = v16
            case "":
                var v17 uint8 = 0
                _ = v17
            default:
                v18 := GenS3{F0: (("\n\t" + ("hello" + "")) + "hello"), F1: 57}
                _ = v18
            }
            v19 := GenS4{F0: false}
            if false {
                if len(v3) > 0 {
                    v3[0] = 513
                }
                _ = len(v3)
                var v20 int32 = v2
                _ = v20
            }
            switch (fi12 & fi12) {
            case 32767:
                fmt.Println(fi12, v2)
            case 1:
                var v21 uint8 = 0
                _ = v21
            }
            fmt.Println(fi13)
            _ = fi13
            _ = v14
            _ = v19
        }
        v22 := GenS5{F0: -527, F1: ("" < "\n\t")}
        var v23 int32 = (v2 - v2)
        _ = fi12
        _ = v22
        _ = v23
    }
    v2 -= (v2 - (v2 << ((1) & 63)))
    var v24 int = v1
    if (0 <= 1) {
        v1 = 982
        fmt.Println(v24)
    } else {
        v2 += v2
        if (true || true) {
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 uint64 = 0
                fmt.Println(v2, v1)
                v27 := GenS6{F0: 4294967295, F1: -9223372036854775808}
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    }
    var v28 string = "hello"
    v28 = v28
    switch "" {
    case "hello":
        if true {
            v28 = "\n\t"
            var v29 int64 = -1
            _ = v29
        }
    case "\n\t":
        switch -1 {
        case 239:
            if ((("\n\t" + v28) + "a") <= (v28 + ("foo bar" + "hello"))) {
                v28 = ((v28 + "") + "hello")
                v28 = v28
            } else {
                if false {
                    fmt.Println(v2, v1, v1)
                    v30 := GenS7{F0: -9223372036854775808, F1: -571, F2: 2}
                    _ = v30
                } else {
                    v2 += 2147483647
                }
                var v31 uint8 = 255
                _ = v31
            }
        case 32767:
            var v32 uint64 = 1
            _ = v32
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    _ = len(v3)
    v2 = v2
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint8 = 40
        var v35 uint = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 8; fi36++ {
            v37 := GenS6{F0: 2, F1: 87}
            if true {
                for fi38 := 0; fi38 < 10; fi38++ {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                    } else {
                        if (1 >= 32767) {
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            switch ("foo bar" + (v28 + v28)) {
                            case "foo bar":
                                v39 := GenS8{F0: (true || true)}
                                _ = v39
                            case "":
                                v24 -= v24
                            case "a":
                                var v40 float32 = 1e10
                                switch 865 {
                                case 843:
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                case 9223372036854775807:
                                    v28 = "foo bar"
                                    _ = len(v3)
                                case 127:
                                    v41 := GenS9{F0: 824}
                                    _ = v41
                                }
                                if (1 != 42) {
                                    if (165 <= 2) {
                                        v1 = v1
                                    }
                                    v24 = fi38
                                    var v42 float32 = (v40 + v40)
                                    _ = v42
                                } else {
                                    v24 = -604
                                    var v43 float32 = 1e10
                                    v40 = (v40 + (v43 + 0.0))
                                    _ = v43
                                }
                                v28 = v28
                                v2 -= (v2 % ((v2) | 1))
                                _ = v40
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v44 int = (v1 << ((65535) & 63))
                            fmt.Println(fi36, fi38)
                            _ = v44
                        }
                        if true {
                            v1 -= v24
                        }
                    }
                    switch v28 {
                    case "\n\t":
                        var v45 int16 = -1
                        _ = v45
                    case "hello":
                        v3 = append(v3, 65535)
                    default:
                        v2 += v2
                    }
                    _ = fi38
                }
                _ = len(v3)
            } else {
                if (4294967295 < 4294967295) {
                    if (42 < 810) {
                        if false {
                            v1 += v24
                        } else {
                            var v46 bool = true
                            v3 = append(v3, 74)
                            _ = v46
                        }
                        var v47 uint64 = 42
                        for fi48 := 0; fi48 < 3; fi48++ {
                            v49 := GenS6{F0: 927, F1: 0}
                            _ = fi48
                            _ = v49
                        }
                        _ = v47
                    }
                }
                for fi50 := 0; fi50 < 17; fi50++ {
                    v24 = 1
                    _ = fi50
                }
                v2 = v2
            }
            v2 = (-(v2))
            _ = fi36
            _ = v37
        }
        if true {
            if (-2.25 == 3.14159) {
                switch ("hello" + (v28 + "a")) {
                case "":
                    var v51 int64 = 42
                    _ = v51
                case "a":
                    fmt.Println(v2, v24)
                default:
                    v2 = (v2 >> ((42) & 63))
                }
            } else {
                switch -1 {
                case 1:
                    fmt.Println(v28)
                default:
                    switch 128 {
                    case 32767:
                        fmt.Println(v24, v24)
                    case -305:
                        for fi52 := 0; fi52 < 13; fi52++ {
                            switch "" {
                            case "hello":
                                var v53 uint32 = 2147483648
                                _ = v53
                            case "a":
                                var v54 uint = 269
                                _ = v54
                            case "foo bar":
                                v2 += (v2 + v2)
                            default:
                                var v55 int32 = 2147483647
                                _ = v55
                            }
                            if (!(!(2 == 618))) {
                                v56 := GenS8{F0: (!false)}
                                if true {
                                    var v57 uint32 = 0
                                    _ = v57
                                }
                                v24 += (fi52 - (fi52 << ((1) & 63)))
                                _ = v56
                            } else {
                                if (2147483648 != 573) {
                                    if false {
                                        v24 -= 4294967295
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 288
                                        }
                                        if false {
                                            v2 = 0
                                            fmt.Println(v2)
                                            switch v28 {
                                            case "foo bar":
                                                v58 := GenS2{F0: -32768, F1: fi52, F2: 119}
                                                _ = v58
                                            default:
                                                v1 = 748
                                            }
                                        } else {
                                            if ("hello" > v28) {
                                                for fi59 := 0; fi59 < 1; fi59++ {
                                                    v3 = append(v3, 42)
                                                    switch fi52 {
                                                    case 0:
                                                        v28 = "\n\t"
                                                    default:
                                                        var v60 uint8 = 255
                                                        v61 := GenS10{F0: 127, F1: -1, F2: 2}
                                                        v62 := GenS4{F0: false}
                                                        _ = v60
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    _ = fi59
                                                }
                                                v28 = v28
                                                v28 = "a"
                                            } else {
                                                for fi63 := 0; fi63 < 17; fi63++ {
                                                    switch fi63 {
                                                    case 2147483647:
                                                        var v64 string = ""
                                                        _ = v64
                                                    case -15:
                                                        var v65 float32 = -2.25
                                                        _ = v65
                                                    case 0:
                                                        switch (v24 & v24) {
                                                        case 0:
                                                            v28 = v28
                                                        case 42:
                                                            v2 = v2
                                                        }
                                                    }
                                                    var v66 float64 = ((1.5 - 1.5) + 3.14159)
                                                    if false {
                                                        var v67 uint8 = 80
                                                        v68 := GenS11{F0: v67, F1: v2}
                                                        v2 += 42
                                                        var v69 float32 = 1.5
                                                        var v70 bool = true
                                                        if v70 {
                                                            v71 := GenS12{F0: (v67 % ((v67) | 1))}
                                                            _ = v71
                                                        } else {
                                                            var v72 int32 = v2
                                                            v69 = 0.1
                                                            if v70 {
                                                                v1 -= v1
                                                                var v73 string = v28
                                                                var v74 int = (v1 >> ((65535) & 63))
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                switch (((v28 + v28) + v28) + ("" + v28)) {
                                                                case "hello":
                                                                    switch v1 {
                                                                    case 127:
                                                                        _ = len(v3)
                                                                        if false {
                                                                            v75 := GenS2{F0: 863, F1: v24, F2: 0}
                                                                            v1 += -741
                                                                            var v76 int64 = -446
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v77 := GenS11{F0: 181, F1: (v72 / ((v72) | 1))}
                                                                                _ = v77
                                                                            case "hello":
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 42
                                                                                }
                                                                            }
                                                                            var v78 int16 = 0
                                                                            _ = v75
                                                                            _ = v76
                                                                            _ = v78
                                                                        }
                                                                        v79 := GenS6{F0: 1, F1: int64(v67)}
                                                                        _ = v79
                                                                    case 0:
                                                                        fmt.Println(v1, v1)
                                                                        v80 := GenS13{F0: 1}
                                                                        _ = v80
                                                                    }
                                                                default:
                                                                    v81 := GenS14{F0: 2}
                                                                    _ = v81
                                                                }
                                                                var v82 float32 = v69
                                                                _ = v82
                                                            }
                                                            _ = v72
                                                        }
                                                        v83 := GenS2{F0: -159, F1: (fi52 / ((fi52) | 1)), F2: 221}
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                        _ = v83
                                                    }
                                                    _ = fi63
                                                    _ = v66
                                                }
                                                var v84 string = "hello"
                                                _ = v84
                                            }
                                            var v85 bool = (-877 <= 9223372036854775807)
                                            _ = v85
                                        }
                                    }
                                    switch (v28 + v28) {
                                    case "a":
                                        _ = len(v3)
                                    case "hello":
                                        if (-1 > 100) {
                                            v3 = append(v3, 4294967295)
                                        } else {
                                            if true {
                                                v24 += fi52
                                                var v86 float32 = (float32(-2.25) / float32(float64(1.5)))
                                                _ = v86
                                            }
                                            var v87 int = v24
                                            switch v24 {
                                            case 9223372036854775807:
                                                for fi88 := 0; fi88 < 14; fi88++ {
                                                    fmt.Println(v2)
                                                    v1 -= v1
                                                    if true {
                                                        v28 = (v28 + "hello")
                                                    } else {
                                                        v89 := GenS15{F0: 873}
                                                        if false {
                                                            v87 = 127
                                                            switch v28 {
                                                            case "hello":
                                                                _ = len(v3)
                                                            case "":
                                                                v90 := GenS12{F0: 128}
                                                                _ = v90
                                                            case "foo bar":
                                                                v91 := GenS7{F0: 42, F1: -1, F2: 42}
                                                                fmt.Println(v1, fi88, v2)
                                                                if (true && false) {
                                                                    var v92 uint = 474
                                                                    v93 := GenS10{F0: 34, F1: -1, F2: 42}
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                                _ = v91
                                                            default:
                                                                var v94 int = -759
                                                                _ = v94
                                                            }
                                                        } else {
                                                            v24 = (fi52 - 32768)
                                                            v2 = v2
                                                        }
                                                        fmt.Println(v87, fi52, v1)
                                                        _ = v89
                                                    }
                                                    _ = fi88
                                                }
                                            }
                                            _ = v87
                                        }
                                    }
                                } else {
                                    v3 = append(v3, 427)
                                    var v95 bool = (65535 == 65535)
                                    _ = v95
                                }
                                if (!true) {
                                    if ((557 <= -1) || false) {
                                        if (0 < 0) {
                                            v96 := GenS8{F0: true}
                                            _ = v96
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        v97 := GenS16{F0: 1000, F1: 32767}
                                        var v98 int64 = -9223372036854775808
                                        _ = v97
                                        _ = v98
                                    }
                                    v2 = (v2 >> ((0) & 63))
                                }
                                v1 += (^(fi52))
                            }
                            _ = fi52
                        }
                    case 836:
                        fmt.Println(v1, v1, v28)
                    }
                }
            }
            switch (^(v1)) {
            case 32768:
                if false {
                    var v99 string = ("hello" + v28)
                    if false {
                        if false {
                            switch v24 {
                            case -412:
                                for fi100 := 0; fi100 < 1; fi100++ {
                                    if (!false) {
                                        v24 = (fi100 / ((-9223372036854775808) | 1))
                                        var v101 string = ((v28 + v28) + v99)
                                        _ = v101
                                    }
                                    v102 := GenS13{F0: 2}
                                    _ = fi100
                                    _ = v102
                                }
                            }
                        }
                        for fi103 := 0; fi103 < 17; fi103++ {
                            fmt.Println(v1)
                            _ = fi103
                        }
                        var v104 uint32 = 1
                        _ = v104
                    } else {
                        v3 = append(v3, 2147483648)
                        v105 := GenS8{F0: false}
                        _ = v105
                    }
                    if false {
                        if false {
                            v24 += (v1 + -9223372036854775808)
                        }
                    } else {
                        var v106 int16 = 181
                        _ = len(v3)
                        v107 := GenS1{F0: -232}
                        _ = v106
                        _ = v107
                    }
                    _ = v99
                }
            case -638:
                var v108 uint16 = 538
                _ = v108
            case 42:
                v109 := GenS6{F0: 649, F1: 1}
                _ = v109
            default:
                fmt.Println(v2)
            }
            var v110 uint64 = 1
            _ = v110
        }
    } else {
        _ = len(v3)
        var v111 int32 = v2
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v28
    _ = v3
}
