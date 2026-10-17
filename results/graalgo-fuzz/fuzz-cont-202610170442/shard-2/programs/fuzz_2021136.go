package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int64; F2 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint16; F2 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 uint32; F2 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64 }
func genFn0(p0 int16, p1 string) int64 {
    p0 = p0
    p0 = -32768
    p0 -= (p0 ^ (p0 * p0))
    _ = p0
    _ = p1
    return -546
}
func genFn1() string {
    var v4 uint8 = 1
    var v5 uint64 = 0
    v6 := GenS1{F0: 205, F1: 42}
    _ = v4
    _ = v5
    _ = v6
    return ("" + "hello")
}

func main() {
    var v1 int = 0
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    v3 = 1
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS2{F0: 105, F1: ""}
        if (true && true) {
            v3 -= (-(v3))
            _ = len(v7)
        } else {
            var v10 uint64 = 1
            fmt.Println(fi8, v3)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    var v11 string = ("\n\t" + "hello")
    if ((false && (true && (false && false))) && true) {
        fmt.Println(v1, v3, v1)
        if len(v7) > 0 {
            v7[0] = (v1 | v1)
        }
    }
    switch v11 {
    case "hello":
        var v12 int8 = -128
        _ = v12
    case "\n\t":
        var v13 string = v11
        _ = v13
    default:
        switch v11 {
        case "foo bar":
            switch v11 {
            case "\n\t":
                if false {
                    var v14 string = v11
                    _ = v14
                }
            case "hello":
                var v15 uint32 = 2
                _ = v15
            }
        case "hello":
            for fi16 := 0; fi16 < 14; fi16++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi16
            }
        default:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 bool = false
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 string = ""
                    switch fi17 {
                    case 9223372036854775807:
                        var v21 float64 = 1.5
                        _ = v21
                    case 32767:
                        if v18 {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        } else {
                            if (false || v18) {
                                var v22 int32 = 382
                                v23 := GenS3{F0: 4294967295}
                                switch fi17 {
                                case -253:
                                    if (689 >= -9223372036854775808) {
                                        switch (fi19 % ((v1) | 1)) {
                                        case -2147483648:
                                            var v24 int16 = -94
                                            _ = v24
                                        }
                                        v1 -= (fi17 | v1)
                                    }
                                case -9223372036854775808:
                                    v25 := GenS4{F0: 127, F1: -9223372036854775808, F2: 601}
                                    _ = v25
                                }
                                var v26 uint32 = 886
                                _ = v22
                                _ = v23
                                _ = v26
                            } else {
                                var v27 int = fi17
                                switch (v1 ^ -349) {
                                case 9223372036854775807:
                                    var v28 int64 = 42
                                    var v29 uint16 = 755
                                    fmt.Println(v29, v29)
                                    var v30 uint = 255
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                case 32768:
                                    v1 += (v27 | v27)
                                default:
                                    v27 -= (v1 & fi17)
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                var v31 int8 = 0
                                if (v11 != ((((((v11 + v20) + (v11 + v20)) + "a") + v11) + v11) + v11)) {
                                    v11 = v11
                                } else {
                                    v32 := GenS5{F0: 184, F1: v3, F2: 2}
                                    v33 := GenS3{F0: 452}
                                    v27 = (v27 * (^(fi17)))
                                    _ = v32
                                    _ = v33
                                }
                                _ = v27
                                _ = v31
                            }
                            switch (("a" + "") + v11) {
                            case "foo bar":
                                delete(v2, 956)
                            case "hello":
                                var v34 float64 = 1.5
                                _ = v34
                            }
                            v35 := GenS1{F0: 1, F1: -1}
                            var v36 uint8 = 42
                            var v37 uint = 921
                            var v38 uint32 = 1
                            var v39 int8 = -1
                            if (!true) {
                                var v40 uint32 = 234
                                fmt.Println(v38, v38)
                                _ = v40
                            }
                            var v41 int16 = -32768
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v41
                        }
                    default:
                        v1 = v1
                        switch fi17 {
                        case -557:
                            fmt.Println(v3, fi17)
                        }
                        _ = len(v2)
                    }
                    fmt.Println(v1, fi17)
                    _ = fi19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
        }
    }
    if false {
        for fi42 := 0; fi42 < 12; fi42++ {
            if (18446744073709551615 <= 0) {
                var v43 int64 = 42
                var v44 uint32 = 2147483648
                var v45 uint = 42
                _ = v43
                _ = v44
                _ = v45
            } else {
                if false {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        var v47 int8 = 0
                        _ = fi46
                        _ = v47
                    }
                } else {
                    switch v11 {
                    case "hello":
                        v3 = (-(v3))
                    case "\n\t":
                        var v48 int16 = 1000
                        _ = v48
                    default:
                        var v49 uint8 = 71
                        _ = v49
                    }
                    v3 = (v3 >> ((4294967295) & 63))
                }
                fmt.Println(v3)
            }
            if (float64(1.5) != 1.5) {
                if ((42 == 2) && false) {
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 uint16 = 65535
                        _ = fi50
                        _ = v51
                    }
                    if (919 <= 988) {
                        if false {
                            if true {
                                if (103 != 65535) {
                                    v11 = "a"
                                    var v52 int = v1
                                    _ = v52
                                }
                                if (v11 <= ((v11 + v11) + "\n\t")) {
                                    fmt.Println(v11)
                                }
                            } else {
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    v1 = (fi53 % (((fi42 - (^(v1)))) | 1))
                                    var v54 float64 = 0.1
                                    fmt.Println(fi42, v54, fi42)
                                    _ = fi53
                                    _ = v54
                                }
                                fmt.Println(fi42, v1)
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            var v55 uint = 65535
                            _ = v55
                        } else {
                            v56 := GenS6{F0: 32768}
                            var v57 uint8 = 1
                            var v58 int64 = -884
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        v1 -= (v1 + fi42)
                    } else {
                        var v59 string = ("hello" + (v11 + "hello"))
                        switch fi42 {
                        case 32768:
                            fmt.Println(v11, v59, v11)
                        case 339:
                            fmt.Println(v3)
                        }
                        _ = v59
                    }
                }
                switch v1 {
                case 32768:
                    var v60 int64 = 777
                    _ = v60
                case -2147483648:
                    if len(v7) > 0 {
                        v7[0] = v1
                    }
                default:
                    for fi61 := 0; fi61 < 7; fi61++ {
                        var v62 int32 = v3
                        var v63 float64 = (-2.25 + (((float64(3.14159) - (((float64((0.0 * -2.25)) - float64(1e10)) * -2.25) / ((3.14159 + 3.14159) - 0.0))) / ((1e10 * (0.1 - 3.14159)) + 1.5)) / -2.25))
                        v64 := GenS2{F0: 0, F1: ((v11 + ("a" + v11)) + "")}
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                }
            } else {
                if ((v3 - v3) == v3) {
                    var v65 bool = false
                    v1 += -850
                    _ = v65
                } else {
                    var v66 uint16 = 42
                    _ = v66
                }
                if true {
                    switch (v1 / (((-(v1))) | 1)) {
                    case 9223372036854775807:
                        v11 = "a"
                    }
                    var v67 uint16 = 65535
                    if false {
                        if true {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if false {
                                v68 := GenS7{F0: ("hello" + (("" + (((v11 + v11) + v11) + v11)) + "hello")), F1: v67, F2: 0}
                                _ = v68
                            } else {
                                fmt.Println(v67, v1)
                            }
                            v1 += (fi42 / ((fi42) | 1))
                        } else {
                            for fi69 := 0; fi69 < 3; fi69++ {
                                v11 = "foo bar"
                                v67 += v67
                                _ = fi69
                            }
                        }
                        fmt.Println(fi42)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        if false {
                            v70 := GenS8{F0: true}
                            _ = v70
                        }
                        v71 := GenS9{F0: 249, F1: -1}
                        _ = v71
                    }
                    _ = v67
                } else {
                    v11 = ""
                }
            }
            var v72 string = (((v11 + (v11 + (("" + "") + (v11 + v11)))) + ("a" + "a")) + "a")
            _ = fi42
            _ = v72
        }
        _ = len(v2)
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        fmt.Println(v3, v3, fi73)
        switch (v1 * (fi73 - -103)) {
        case 11:
            if true {
                for fi74 := 0; fi74 < 16; fi74++ {
                    if true {
                        if (false || (v3 != (-(v3)))) {
                            if (false || ("foo bar" > v11)) {
                                v75 := v2[628]
                                if false {
                                    var v76 uint32 = v75
                                    v77 := GenS10{F0: 66, F1: 631}
                                    var v78 float64 = (-2.25 / (1.5 / 1.0))
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                }
                                _ = v75
                            } else {
                                v79 := GenS11{F0: ("a" >= (v11 + v11))}
                                v1 = (fi73 + fi73)
                                fmt.Println(v11, v11)
                                v80 := GenS5{F0: 32767, F1: 1, F2: 4294967295}
                                _ = v79
                                _ = v80
                            }
                        } else {
                            v11 = "hello"
                        }
                    }
                    fmt.Println(fi73, v1)
                    if (4294967295 > 533) {
                        fmt.Println(fi73)
                        var v81 int16 = -32768
                        var v82 bool = (!(1 >= 345))
                        _ = v81
                        _ = v82
                    } else {
                        switch 32767 {
                        case -9223372036854775808:
                            if false {
                                fmt.Println(v11, v11, v3)
                                switch v11 {
                                case "a":
                                    if true {
                                        fmt.Println(v3, v3, v11)
                                        v3 = (v3 % (((v3 << ((65535) & 63))) | 1))
                                        delete(v2, 42)
                                    } else {
                                        v83 := GenS10{F0: 127, F1: 372}
                                        var v84 uint8 = 2
                                        v1 = fi74
                                        _ = v83
                                        _ = v84
                                    }
                                case "hello":
                                    v3 += v3
                                }
                                switch ("" + v11) {
                                case "foo bar":
                                    var v85 float64 = 0.1
                                    _ = v85
                                case "\n\t":
                                    var v86 int8 = -128
                                    v3 = (v3 % (((^(v3))) | 1))
                                    var v87 uint16 = 255
                                    switch (fi74 | 2147483647) {
                                    case 4294967295:
                                        v87 = v87
                                    }
                                    _ = v86
                                    _ = v87
                                case "hello":
                                    v88 := v2[42]
                                    _ = v88
                                default:
                                    v89 := GenS12{F0: (float64(1e10) != 1.5)}
                                    _ = v89
                                }
                                fmt.Println(v1, fi74, fi73)
                                var v90 int64 = 42
                                v91 := GenS13{F0: 81, F1: 2, F2: -637}
                                _ = v90
                                _ = v91
                            }
                        }
                    }
                    _ = fi74
                }
                v11 = "foo bar"
                var v92 uint8 = 127
                _ = v92
            }
        case -1:
            if true {
                v3 += v3
                for fi93 := 0; fi93 < 20; fi93++ {
                    var v94 float64 = (float64(float32((0.1 - float64(float32(1.5))))) * float64((-2.25 / (-2.25 / ((-2.25 / 0.1) * (0.1 / 0.1))))))
                    v95 := GenS14{F0: (((v11 + v11) + ((v11 + "a") + v11)) + ((v11 + (v11 + ("" + "hello"))) + v11))}
                    fmt.Println(fi73)
                    _ = fi93
                    _ = v94
                    _ = v95
                }
            }
        default:
            var v96 bool = true
            _ = v96
        }
        v11 = v11
        _ = fi73
    }
    switch ("a" + "a") {
    case "":
        switch v1 {
        case 32768:
            for fi97 := 0; fi97 < 3; fi97++ {
                v98 := GenS10{F0: 100, F1: 4294967295}
                _ = fi97
                _ = v98
            }
        case 4294967295:
            fmt.Println(v3)
        case 380:
            var v99 string = ("hello" + "foo bar")
            _ = v99
        default:
            if (!((0.1 + float32((float64(3.14159) + (1e10 / 1.0)))) > float32(1.5))) {
                var v100 string = (v11 + v11)
                var v101 bool = ((1 < (v3 << ((42) & 63))) && ((!false) || true))
                _ = v100
                _ = v101
            }
        }
    case "foo bar":
        v3 = v3
    case "a":
        _ = len(v2)
    default:
        switch 4294967295 {
        case 650:
            v3 += v3
        case -2147483648:
            v102 := GenS4{F0: 82, F1: 0, F2: -9223372036854775808}
            _ = v102
        default:
            var v103 int16 = 891
            _ = v103
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch v11 {
    case "\n\t":
        var v104 int8 = 54
        _ = v104
    case "hello":
        v11 = (v11 + ((v11 + v11) + v11))
    case "a":
        v105 := GenS15{F0: 1}
        _ = v105
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (2 > 128) {
        if len(v7) > 0 {
            v7[0] = 32768
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
