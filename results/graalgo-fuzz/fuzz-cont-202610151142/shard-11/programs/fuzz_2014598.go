package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 string }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int64; F1 int32; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS14 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS15 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 int, p1 int8) string {
    v7 := GenS1{F0: 42, F1: -760}
    p0 += 2147483647
    _ = p0
    _ = p1
    _ = v7
    return (("a" + ("\n\t" + ((("foo bar" + "a") + (("hello" + ("foo bar" + "hello")) + (("\n\t" + "a") + "hello"))) + "a"))) + "\n\t")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{("a" + "foo bar"), "", ""}
    var v4 uint32 = 65535
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]string)
    var v8 float32 = 0.1
    switch 32768 {
    case 127:
        v9 := v1[v4]
        _ = v9
    case 2147483647:
        switch 32767 {
        case 9223372036854775807:
            if (752 < 42) {
                var v10 uint = 469
                switch "foo bar" {
                case "a":
                    v6[v4] = (((("hello" + (("foo bar" + ("" + "hello")) + "\n\t")) + ("a" + "\n\t")) + "") + ("" + ("hello" + "foo bar")))
                case "foo bar":
                    v4 += v4
                case "":
                    v4 = v4
                }
                var v11 int8 = -69
                _ = v10
                _ = v11
            } else {
                var v12 uint32 = 270
                if len(v3) > 0 {
                    v3[0] = (("" + ("" + "foo bar")) + "hello")
                }
                var v13 int64 = -1
                _ = v12
                _ = v13
            }
        case 128:
            v14 := GenS2{F0: -658}
            _ = v14
        default:
            var v15 float32 = float32(1.5)
            _ = v15
        }
    case -448:
        var v16 uint = 255
        _ = v16
    }
    fmt.Println(v8, v4)
    if (!(9223372036854775807 <= 9223372036854775807)) {
        switch "foo bar" {
        case "a":
            if (65535 == 760) {
                switch "" {
                case "hello":
                    for fi17 := 0; fi17 < 13; fi17++ {
                        var v18 uint64 = 502
                        if (!(!true)) {
                            fmt.Println(v8)
                        }
                        switch 0 {
                        case 128:
                            v8 += 0.1
                        case -328:
                            v1[(v4 << ((uint(v4)) & 63))] = v4
                        default:
                            for fi19 := 0; fi19 < 7; fi19++ {
                                switch (fi19 ^ 32767) {
                                case -374:
                                    var v20 uint8 = 0
                                    _ = v20
                                case -9223372036854775808:
                                    v21 := GenS3{F0: (((("\n\t" + ("" + ("a" + (("foo bar" + "\n\t") + "hello")))) + "") + "a") < "\n\t"), F1: (9223372036854775807 == 9223372036854775807)}
                                    _ = v21
                                case 1000:
                                    v8 = v8
                                default:
                                    if false {
                                        var v22 int8 = 127
                                        _ = v22
                                    }
                                }
                                v8 = (v8 / (v8 / v8))
                                v23 := GenS4{F0: v4, F1: 1000, F2: 32767}
                                _ = fi19
                                _ = v23
                            }
                        }
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    _ = len(v6)
                }
                v24 := GenS1{F0: 32767, F1: 0}
                _ = v24
            } else {
                var v25 int64 = 9223372036854775807
                _ = v25
            }
        default:
            var v26 float64 = 1e10
            _ = v26
        }
        switch "" {
        case "foo bar":
            if true {
                var v27 bool = (!true)
                _ = v27
            } else {
                for fi28 := 0; fi28 < 9; fi28++ {
                    fmt.Println(v8, v4)
                    if ((v8 / (v8 * v8)) <= v8) {
                        v29 := GenS4{F0: 2, F1: 1000, F2: 42}
                        _ = v29
                    } else {
                        var v30 float32 = 1.5
                        v4 -= 2147483647
                        if (!(!((0.1 / 0.1) < (1e10 + -2.25)))) {
                            if false {
                                v8 -= (3.14159 / v8)
                                fmt.Println(v8, v30)
                            }
                        } else {
                            v8 = (v8 - 0.1)
                            if (!(!(-99 < 1))) {
                                v31 := GenS3{F0: ((!((!("a" != "\n\t")) && ((true && true) && true))) && (42 == 2)), F1: (!true)}
                                if (false && (!(!false))) {
                                    var v32 string = ("foo bar" + "hello")
                                    var v33 uint64 = 4294967295
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        switch ("foo bar" + v32) {
                                        case "hello":
                                            v35 := GenS5{F0: ((v32 + "hello") + v32), F1: uint(v4), F2: v32}
                                            _ = v35
                                        default:
                                            var v36 uint16 = 65535
                                            _ = v36
                                        }
                                        v32 = (("foo bar" + v32) + "foo bar")
                                        _ = fi34
                                    }
                                    _ = v32
                                    _ = v33
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v37 uint64 = 0
                                    _ = v37
                                }
                                _ = v31
                            }
                            for fi38 := 0; fi38 < 3; fi38++ {
                                v6[4294967295] = "\n\t"
                                v8 = v30
                                _ = fi38
                            }
                        }
                        _ = v30
                    }
                    var v39 int16 = 963
                    _ = fi28
                    _ = v39
                }
            }
        case "":
            if true {
                switch -192 {
                case 32767:
                    v4 -= 448
                case -242:
                    for fi40 := 0; fi40 < 13; fi40++ {
                        var v41 int64 = 42
                        v42 := GenS2{F0: 645}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
            } else {
                v43 := GenS3{F0: true, F1: (1 > 1)}
                if (true && false) {
                    var v44 uint16 = 923
                    _ = v44
                } else {
                    var v45 int32 = 2147483647
                    var v46 float32 = 1e10
                    _ = v45
                    _ = v46
                }
                fmt.Println(v8, v4)
                _ = v43
            }
        }
        if false {
            if (!true) {
                v8 = 3.14159
            }
            if (false && (1 < 424)) {
                v47 := GenS6{F0: 1000, F1: 1}
                _ = v47
            } else {
                if ((32767 <= 1000) || (true || false)) {
                    v2[9223372036854775807] = (v4 / ((358) | 1))
                }
                v48 := GenS7{F0: 4294967295, F1: -2147483648}
                var v49 int64 = -1
                _ = v48
                _ = v49
            }
        }
    }
    fmt.Println(v4, v4)
    for fi50 := 0; fi50 < 19; fi50++ {
        switch "foo bar" {
        case "hello":
            if ((((!(1 < 361)) && true) && true) && false) {
                v51 := v5[((("a" + ("\n\t" + "\n\t")) + "hello") + ("hello" + ""))]
                _ = v51
            }
        default:
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 int8 = -1
                v53 -= (v53 & (v53 ^ v53))
                _ = fi52
                _ = v53
            }
        }
        var v54 uint64 = 589
        _ = fi50
        _ = v54
    }
    v55 := GenS3{F0: (-2147483648 != -120), F1: (1e10 <= 0.0)}
    var v56 bool = (!true)
    v56 = true
    if v56 {
        for fi57 := 0; fi57 < 8; fi57++ {
            v56 = (v56 && v56)
            if (159 != 2) {
                v58 := GenS8{F0: 1, F1: int32(fi57), F2: -2147483648}
                if v56 {
                    for fi59 := 0; fi59 < 14; fi59++ {
                        var v60 int = (fi57 & (-(fi57)))
                        v60 = -15
                        v60 += (fi59 << ((255) & 63))
                        v4 += 2
                        v61 := GenS9{F0: v4, F1: 65535, F2: 42}
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                } else {
                    fmt.Println(v4, v8, v8)
                    var v62 uint64 = 0
                    _ = v62
                }
                var v63 uint32 = v4
                _ = v58
                _ = v63
            } else {
                v64 := GenS10{F0: 1}
                for fi65 := 0; fi65 < 17; fi65++ {
                    var v66 int16 = 0
                    if true {
                        var v67 int = 32767
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(fi57, v67, v8)
                        } else {
                            if (!(v8 < v8)) {
                                v56 = false
                                var v68 float64 = 1e10
                                var v69 uint = 65535
                                _ = len(v3)
                                v68 = ((0.0 + 1.5) - (float64(v8) / v68))
                                _ = v68
                                _ = v69
                            } else {
                                v56 = v56
                                var v70 string = ("\n\t" + "a")
                                v66 = (v66 - 1)
                                v8 += 0.0
                                _ = v70
                            }
                        }
                        switch (-(fi65)) {
                        case 821:
                            if true {
                                v71 := GenS5{F0: "a", F1: 720, F2: (("\n\t" + ("a" + "\n\t")) + "")}
                                _ = v71
                            }
                        default:
                            if (v56 && v56) {
                                switch (fi57 | (fi65 << ((1) & 63))) {
                                case 32768:
                                    v4 -= (v4 << ((1) & 63))
                                    v6[(v4 - v4)] = "a"
                                    var v72 uint64 = 0
                                    _ = v72
                                case -1:
                                    v4 = 847
                                }
                                var v73 uint64 = 18446744073709551615
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v73
                            }
                        }
                        _ = v67
                    } else {
                        fmt.Println(fi65, fi65, v66)
                        if v56 {
                            v8 = v8
                            v4 = (^(v4))
                            if v56 {
                                var v74 uint8 = 84
                                if v56 {
                                    var v75 uint32 = (v4 + v4)
                                    fmt.Println(v75, v66)
                                    _ = v75
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v76 int32 = -1
                                        _ = v76
                                    }
                                    if (!false) {
                                        v8 = 1.5
                                        fmt.Println(v74, fi57)
                                    } else {
                                        fmt.Println(v66)
                                        var v77 uint32 = v4
                                        _ = v77
                                    }
                                    if v56 {
                                        if v56 {
                                            switch ("\n\t" + "") {
                                            case "":
                                                v74 = 74
                                            }
                                            var v78 uint16 = 65535
                                            _ = len(v1)
                                            var v79 uint16 = 42
                                            _ = v78
                                            _ = v79
                                        } else {
                                            v8 += (-2.25 * float32(-2.25))
                                        }
                                    } else {
                                        if v56 {
                                            v80 := GenS9{F0: v4, F1: uint(v4), F2: 717}
                                            if (v56 || v56) {
                                                var v81 float64 = 1e10
                                                v82 := GenS4{F0: 42, F1: 1, F2: 0}
                                                if v56 {
                                                    var v83 uint16 = 32767
                                                    if ((v83 * (v83 >> ((4294967295) & 63))) < (^(v83))) {
                                                        v56 = (fi65 < (fi65 / ((fi65) | 1)))
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                    } else {
                                                        var v84 uint16 = v83
                                                        var v85 uint32 = 0
                                                        switch int(v4) {
                                                        case -913:
                                                            var v86 int8 = 42
                                                            _ = v86
                                                        case 9223372036854775807:
                                                            var v87 int = int(v84)
                                                            _ = v87
                                                        case 0:
                                                            if (0 == v4) {
                                                                if (0 <= -9223372036854775808) {
                                                                    switch fi57 {
                                                                    case -647:
                                                                        v74 += (v74 % ((v74) | 1))
                                                                    case 0:
                                                                        switch "" {
                                                                        case "a":
                                                                            if (!("\n\t" <= "foo bar")) {
                                                                                v81 -= -2.25
                                                                                var v88 uint32 = 2
                                                                                var v89 float32 = v8
                                                                                v83 = (v83 + 42)
                                                                                v90 := v5[""]
                                                                                if ((-9223372036854775808 == -9223372036854775808) && v56) {
                                                                                    v91 := GenS2{F0: 0}
                                                                                    if v56 {
                                                                                        v92 := GenS11{F0: 1, F1: (v56 || (!((int16(fi65) > v66) && ((!(2 == v85)) || ("foo bar" == ("" + ("foo bar" + "")))))))}
                                                                                        v85 += v90
                                                                                        var v93 bool = v56
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                    _ = v91
                                                                                }
                                                                                _ = v88
                                                                                _ = v89
                                                                                _ = v90
                                                                            } else {
                                                                                if (42 == 239) {
                                                                                    v83 = v83
                                                                                } else {
                                                                                    if (true && (!v56)) {
                                                                                        if v56 {
                                                                                            if v56 {
                                                                                                var v94 uint8 = v74
                                                                                                var v95 uint8 = (v94 >> ((0) & 63))
                                                                                                v96 := GenS12{F0: 2147483647, F1: (v83 & v83)}
                                                                                                var v97 uint8 = v95
                                                                                                v97 = 128
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            } else {
                                                                                                fmt.Println(v83, v81, v4)
                                                                                                if (18446744073709551615 < 139) {
                                                                                                    v3 = append(v3, "\n\t")
                                                                                                } else {
                                                                                                    v98 := GenS13{F0: 42, F1: 18446744073709551615, F2: (^(v74))}
                                                                                                    var v99 int16 = v66
                                                                                                    _ = v98
                                                                                                    _ = v99
                                                                                                }
                                                                                            }
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        v84 = (v84 | v84)
                                                                                        var v100 uint32 = v85
                                                                                        _ = v100
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "":
                                                                            switch int(v4) {
                                                                            case 42:
                                                                                var v101 uint16 = (v84 & 243)
                                                                                _ = v101
                                                                            default:
                                                                                v102 := GenS14{F0: (v74 / ((255) | 1)), F1: v84, F2: 44}
                                                                                _ = v102
                                                                            }
                                                                        default:
                                                                            v103 := GenS6{F0: (v66 % ((int16(v83)) | 1)), F1: 710}
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v81, fi57)
                                                                    }
                                                                    delete(v6, 633)
                                                                }
                                                            }
                                                        default:
                                                            var v104 string = "foo bar"
                                                            _ = v104
                                                        }
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    v105 := GenS10{F0: 1}
                                                    _ = v105
                                                    _ = v83
                                                } else {
                                                    switch 0 {
                                                    case 32767:
                                                        v106 := GenS15{F0: -567, F1: "\n\t", F2: v74}
                                                        _ = v106
                                                    }
                                                }
                                                _ = v81
                                                _ = v82
                                            } else {
                                                fmt.Println(v66, v56, v4)
                                            }
                                            _ = v80
                                        } else {
                                            switch fi65 {
                                            case 769:
                                                v4 += (v4 ^ (v4 & (v4 * (v4 / ((v4) | 1)))))
                                            case -1:
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(v56, fi65, v56)
                                                case "\n\t":
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                case "foo bar":
                                                    var v107 int64 = -1
                                                    _ = len(v1)
                                                    _ = v107
                                                }
                                            default:
                                                delete(v5, "\n\t")
                                            }
                                        }
                                    }
                                }
                                var v108 uint8 = (^(v74))
                                v8 = 0.0
                                _ = v108
                                _ = v74
                            }
                        }
                        v66 = int16(v4)
                    }
                    var v109 int16 = (v66 * v66)
                    _ = fi65
                    _ = v109
                    _ = v66
                }
                v8 += 0.1
                _ = v64
            }
            v56 = (127 >= 42)
            _ = fi57
        }
    }
    var v110 int64 = -1
    var v111 uint16 = 164
    fmt.Println(v56, v111, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v8
}
