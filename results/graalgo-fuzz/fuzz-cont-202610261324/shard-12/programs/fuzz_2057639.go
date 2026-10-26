package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int16; F2 uint }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 string; F1 int }
type GenS12 struct { F0 int8; F1 int8 }
func genFn0() int {
    if ((true || true) && true) {
        var v7 uint8 = 2
        var v8 float64 = (((3.14159 - ((0.1 + (3.14159 + -2.25)) / 1.0)) / (((((3.14159 / 1.5) - -2.25) / 1e10) / 1e10) / 1.0)) + 0.1)
        _ = v7
        _ = v8
    }
    var v9 uint16 = 65535
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := []int{}
    var v4 uint64 = 42
    var v5 float64 = (0.1 - (float64(0.1) + float64((3.14159 - -2.25))))
    var v6 uint8 = 211
    var v10 bool = false
    if (255 >= (v6 - (v6 * (^(v6))))) {
        if v2 {
            v10 = v2
            v4 += (^(v4))
            var v11 int8 = 0
            _ = v11
        } else {
            fmt.Println(v2)
            v12 := GenS1{F0: 0, F1: 42, F2: 2147483647}
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            if (uint16(v4) <= 579) {
                v10 = (("hello" + ("foo bar" + "")) > (("\n\t" + ("a" + "hello")) + "\n\t"))
            } else {
                _ = len(v1)
                _ = len(v1)
                for fi14 := 0; fi14 < 7; fi14++ {
                    if len(v3) > 0 {
                        v3[0] = (fi14 - (fi14 * fi14))
                    }
                    _ = fi14
                }
            }
            _ = fi13
        }
        var v15 int16 = 793
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        switch "" {
        case "hello":
            fmt.Println(v10)
        default:
            v5 = -2.25
        }
        v4 = (v4 & (v4 / ((879) | 1)))
        _ = fi16
    }
    v4 += v4
    fmt.Println(v4)
    var v17 int32 = -639
    switch "\n\t" {
    case "hello":
        v18 := GenS2{F0: 2}
        _ = v18
    case "a":
        v2 = (false || v10)
    default:
        if v10 {
            v4 = (v4 % ((733) | 1))
        } else {
            delete(v1, "a")
        }
    }
    switch "hello" {
    case "":
        delete(v1, "hello")
    case "hello":
        delete(v1, "\n\t")
    case "foo bar":
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint8 = v6
            _ = fi19
            _ = v20
        }
    default:
        v21 := GenS3{F0: 288, F1: "hello"}
        _ = v21
    }
    v22 := GenS4{F0: 18446744073709551615}
    v17 -= int32(v4)
    v5 = (v5 + float64(0.1))
    v5 = 0.0
    if false {
        for fi23 := 0; fi23 < 16; fi23++ {
            switch fi23 {
            case -1:
                fmt.Println(v5, v4, fi23)
            case 32768:
                if (190 == 2147483647) {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if (!(v4 == (v4 << ((919) & 63)))) {
                            _ = len(v3)
                        } else {
                            v4 += v4
                            v2 = true
                            switch (fi23 << ((4294967295) & 63)) {
                            case 15:
                                v5 += v5
                            case 9223372036854775807:
                                switch (fi23 | (fi23 * (-(fi24)))) {
                                case 128:
                                    var v25 uint32 = 2147483647
                                    _ = v25
                                case -1:
                                    if v2 {
                                        if v10 {
                                            if v10 {
                                                v26 := GenS5{F0: 2, F1: "foo bar"}
                                                v27 := GenS1{F0: 32767, F1: 59, F2: 53}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v2 = v10
                                                switch "a" {
                                                case "":
                                                    switch (fi23 + (fi23 * 127)) {
                                                    case -1:
                                                        v10 = v2
                                                    default:
                                                        v2 = (((!(v2 && v10)) && true) && false)
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v2, v5)
                                                case "\n\t":
                                                    fmt.Println(v2, v6)
                                                default:
                                                    var v28 int8 = -15
                                                    _ = v28
                                                }
                                                fmt.Println(v6, v10, v17)
                                            }
                                            v2 = v10
                                            if (!true) {
                                                var v29 string = "\n\t"
                                                var v30 int64 = 9223372036854775807
                                                var v31 uint = 4294967295
                                                v32 := GenS5{F0: uint16(v4), F1: (v29 + ((v29 + (((v29 + v29) + (("a" + v29) + "\n\t")) + ((v29 + "\n\t") + v29))) + v29))}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        } else {
                                            if v10 {
                                                var v33 uint16 = 42
                                                v34 := GenS1{F0: 0, F1: 1, F2: 476}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                if (false && (((!(1 <= 937)) || true) && (true || false))) {
                                                    delete(v1, "foo bar")
                                                } else {
                                                    if (v17 >= -2147483648) {
                                                        if true {
                                                            v35 := v1["\n\t"]
                                                            fmt.Println(v6)
                                                            _ = v35
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = len(v3)
                                                    } else {
                                                        var v36 uint32 = 0
                                                        v4 = (v4 << ((677) & 63))
                                                        v37 := GenS2{F0: 67}
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    var v38 int32 = -34
                                                    switch "foo bar" {
                                                    case "":
                                                        fmt.Println(v6, v10)
                                                    }
                                                    _ = v38
                                                }
                                                v2 = (1 != 4294967295)
                                                v17 = (v17 >> ((2) & 63))
                                            }
                                        }
                                    } else {
                                        v10 = false
                                        var v39 uint32 = 2
                                        _ = v39
                                    }
                                default:
                                    _ = len(v1)
                                }
                            case 628:
                                switch "" {
                                case "a":
                                    var v40 uint32 = 0
                                    _ = v40
                                case "foo bar":
                                    var v41 int32 = (v17 ^ (v17 + v17))
                                    _ = v41
                                }
                            default:
                                fmt.Println(v2)
                            }
                        }
                        if v2 {
                            var v42 float32 = (((0.1 * -2.25) - 0.1) * 1e10)
                            if false {
                                v43 := GenS1{F0: -1, F1: 100, F2: 4294967295}
                                _ = v43
                            } else {
                                var v44 float32 = ((v42 / 1.0) + v42)
                                _ = v44
                            }
                            if false {
                                switch fi24 {
                                case 32768:
                                    v3 = append(v3, 554)
                                case 1:
                                    switch "" {
                                    case "foo bar":
                                        if true {
                                            v45 := GenS6{F0: -9223372036854775808}
                                            fmt.Println(v42)
                                            _ = v45
                                        } else {
                                            switch fi23 {
                                            case 516:
                                                delete(v1, "a")
                                            default:
                                                var v46 int32 = v17
                                                _ = v46
                                            }
                                            v2 = v10
                                            switch ("foo bar" + ("foo bar" + "a")) {
                                            case "a":
                                                switch ("hello" + "hello") {
                                                case "\n\t":
                                                    var v47 int32 = -409
                                                    _ = v47
                                                case "":
                                                    v10 = v10
                                                case "a":
                                                    var v48 uint32 = 517
                                                    _ = v48
                                                }
                                            case "foo bar":
                                                fmt.Println(v2)
                                            default:
                                                v17 = v17
                                            }
                                        }
                                    }
                                case -105:
                                    v49 := GenS7{F0: 1, F1: 106, F2: 0}
                                    _ = v49
                                }
                                fmt.Println(v5)
                                v4 = v4
                            } else {
                                fmt.Println(v2, v17)
                            }
                            v50 := GenS5{F0: 2, F1: ""}
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            _ = v42
                            _ = v50
                        } else {
                            fmt.Println(v6)
                            var v51 int16 = 1
                            if v10 {
                                v2 = v10
                            } else {
                                if (!v10) {
                                    _ = len(v3)
                                }
                                switch (^(fi23)) {
                                case 621:
                                    v52 := GenS8{F0: 2}
                                    _ = v52
                                case 4294967295:
                                    v53 := GenS9{F0: 603, F1: -567, F2: 18446744073709551615}
                                    _ = v53
                                case -768:
                                    var v54 int32 = v17
                                    _ = v54
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v5 = v5
                            }
                            _ = v51
                        }
                        _ = fi24
                    }
                    fmt.Println(fi23, fi23, v10)
                    var v55 string = ("hello" + "a")
                    _ = v55
                } else {
                    v56 := GenS10{F0: "a", F1: 18446744073709551615}
                    switch genFn0() {
                    case 32768:
                        var v57 int32 = v17
                        _ = v57
                    case 2147483647:
                        v17 += (v17 >> ((18446744073709551615) & 63))
                    case -265:
                        if true {
                            if (v17 != -1) {
                                _ = len(v3)
                                var v58 float32 = 3.14159
                                _ = v58
                            }
                            if (1 >= 9223372036854775807) {
                                fmt.Println(v2)
                                v5 -= v5
                            } else {
                                fmt.Println(v10, v2, v6)
                                switch (fi23 % ((128) | 1)) {
                                case 127:
                                    if (!false) {
                                        v59 := GenS2{F0: v6}
                                        v60 := GenS6{F0: 0}
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 int = -462
                                        switch 4294967295 {
                                        case 127:
                                            if ((!((-2.25 / 3.14159) != 0.1)) && true) {
                                                var v62 uint32 = 2147483647
                                                v10 = v10
                                                _ = v62
                                            }
                                        case 128:
                                            v5 -= ((((1e10 * -2.25) * 1e10) - float64(float32(v5))) + v5)
                                        }
                                        _ = v61
                                    }
                                }
                                for fi63 := 0; fi63 < 9; fi63++ {
                                    if v2 {
                                        var v64 float64 = v5
                                        _ = v64
                                    }
                                    _ = fi63
                                }
                            }
                        } else {
                            v65 := GenS11{F0: "", F1: (fi23 | (^(fi23)))}
                            _ = v65
                        }
                    default:
                        var v66 uint32 = 2147483647
                        _ = v66
                    }
                    _ = v56
                }
            default:
                var v67 uint64 = uint64(v6)
                _ = v67
            }
            v68 := GenS12{F0: -1, F1: 100}
            for fi69 := 0; fi69 < 6; fi69++ {
                _ = len(v3)
                _ = fi69
            }
            _ = fi23
            _ = v68
        }
    } else {
        v2 = ((v6 & 2) == (v6 & (v6 >> ((42) & 63))))
        v10 = v2
        switch -605 {
        case 1:
            var v70 bool = (int16(v4) != -498)
            _ = v70
        }
    }
    v71 := v1[("" + "foo bar")]
    _ = len(v1)
    var v72 bool = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
