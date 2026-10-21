package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 string; F1 string; F2 int }
type GenS6 struct { F0 uint8; F1 bool; F2 int }
type GenS7 struct { F0 uint64; F1 int32; F2 int8 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 string; F1 bool }
func genFn0() uint {
    var v4 uint16 = 42
    fmt.Println(v4)
    var v5 int32 = 42
    var v6 uint32 = 42
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint16, p1 int32, p2 uint16) int16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if (18446744073709551615 > 1) {
            v8 := []int{fi7, (fi7 ^ -9223372036854775808), (fi7 / ((fi7) | 1))}
            _ = v8
        }
        _ = fi7
    }
    p0 = 579
    _ = p0
    _ = p1
    _ = p2
    return -672
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{}
    v3 := []uint32{}
    if (((1.5 / v1) * (((0.1 / float64(3.14159)) / (1e10 * ((1e10 + v1) * v1))) * v1)) > (v1 * 3.14159)) {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 830, F1: 65535, F2: "\n\t"}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi9
            _ = v10
        }
        fmt.Println(v1)
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    var v11 bool = (9223372036854775807 > 834)
    fmt.Println(v1, v11)
    var v12 int = 32768
    v1 += 1.5
    fmt.Println(v11)
    v12 += (v12 + v12)
    switch v12 {
    case 2147483647:
        fmt.Println(v12, v12, v12)
    case -392:
        v11 = v11
    default:
        v13 := GenS1{F0: 276, F1: 0, F2: "foo bar"}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v11)
        if (v11 || true) {
            v15 := GenS1{F0: 2, F1: 2147483648, F2: "foo bar"}
            var v16 int16 = -625
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint64 = 1
    v3 = append(v3, 782)
    var v18 float32 = 1e10
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            v18 = (-2.25 / (v18 / (float32(v1) - 3.14159)))
            switch 127 {
            case -2147483648:
                fmt.Println(v1, fi19, v12)
            case 2147483647:
                var v21 uint64 = v17
                _ = v21
            case 0:
                if true {
                    v11 = v11
                    switch (v12 + (fi20 - (^(fi19)))) {
                    case 54:
                        v22 := GenS1{F0: 65535, F1: 705, F2: (("foo bar" + ("\n\t" + ("\n\t" + "\n\t"))) + "foo bar")}
                        fmt.Println(fi19)
                        _ = v22
                    case 128:
                        var v23 string = ""
                        _ = v23
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    if (0 <= 509) {
                        v12 += (v12 ^ 127)
                    }
                    fmt.Println(v1, fi19)
                }
            default:
                if true {
                    v12 = fi20
                    v18 += 1.5
                    v12 -= (v12 / (((v12 ^ 9223372036854775807)) | 1))
                } else {
                    v17 = v17
                    v24 := GenS2{F0: 42, F1: 143, F2: 998}
                    _ = v24
                }
            }
            _ = fi20
        }
        if true {
            var v25 float32 = v18
            v26 := GenS3{F0: 0}
            if (v11 || true) {
                _ = len(v2)
            }
            _ = v25
            _ = v26
        } else {
            v18 = 0.0
            var v27 uint64 = v17
            switch v12 {
            case 32767:
                for fi28 := 0; fi28 < 5; fi28++ {
                    v29 := GenS4{F0: 9223372036854775807, F1: -1}
                    fmt.Println(v27)
                    if v11 {
                        var v30 uint16 = uint16(v17)
                        _ = v30
                    }
                    _ = fi28
                    _ = v29
                }
            case 970:
                switch ("\n\t" + "") {
                case "a":
                    if v11 {
                        v31 := GenS5{F0: "", F1: "\n\t", F2: 127}
                        var v32 string = "foo bar"
                        if true {
                            v32 = ("hello" + v32)
                        } else {
                            for fi33 := 0; fi33 < 7; fi33++ {
                                if ((!(((v11 || v11) || v11) && v11)) || ((!(v11 || (v1 > v1))) && ((fi33 >> ((27) & 63)) == 922))) {
                                    if (562 < 899) {
                                        var v34 float64 = v1
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + "foo bar")
                                        }
                                        _ = v34
                                    } else {
                                        v35 := GenS5{F0: "hello", F1: "hello", F2: fi33}
                                        var v36 uint64 = 2
                                        v37 := GenS6{F0: 255, F1: ((v27 * (v27 + (v17 ^ (^(v17))))) > v17), F2: -390}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 uint64 = v17
                                    if (v38 >= v17) {
                                        v2 = append(v2, v32)
                                    }
                                    _ = v38
                                }
                                _ = fi33
                            }
                            fmt.Println(v27)
                        }
                        _ = v31
                        _ = v32
                    } else {
                        for fi39 := 0; fi39 < 3; fi39++ {
                            if v11 {
                                fmt.Println(fi19)
                                if ((!v11) || v11) {
                                    fmt.Println(v18, v27, v27)
                                    var v40 string = "hello"
                                    _ = v40
                                } else {
                                    var v41 uint = 18446744073709551615
                                    _ = v41
                                }
                            } else {
                                v11 = true
                                v1 = 3.14159
                                if false {
                                    v42 := GenS7{F0: v27, F1: -793, F2: 42}
                                    var v43 int32 = -2147483648
                                    _ = v42
                                    _ = v43
                                } else {
                                    if ((!(4294967295 <= 65535)) && (v27 <= 4294967295)) {
                                        if false {
                                            v1 -= (v1 * float64((0.0 * v18)))
                                        }
                                        v27 = (v27 / ((v17) | 1))
                                    }
                                    v3 = append(v3, 779)
                                    v1 += float64(1.5)
                                }
                                v44 := GenS6{F0: 127, F1: ((2147483647 == 42) || (("a" != "foo bar") || (42 >= 162))), F2: (fi19 * (fi19 % (((fi19 | fi19)) | 1)))}
                                var v45 int = 4294967295
                                var v46 uint8 = 255
                                v47 := GenS8{F0: 255, F1: "\n\t"}
                                _ = v44
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi39
                        }
                        if true {
                            v48 := GenS8{F0: 4294967295, F1: ("hello" + "foo bar")}
                            v1 = v1
                            _ = v48
                        } else {
                            switch v12 {
                            case 2147483647:
                                var v49 uint = 2
                                _ = v49
                            case -1:
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                            case 4294967295:
                                switch (^(v12)) {
                                case -9223372036854775808:
                                    fmt.Println(v12, v27, v12)
                                }
                            default:
                                fmt.Println(v17)
                            }
                            var v50 float64 = v1
                            _ = v50
                        }
                        var v51 int8 = 42
                        _ = v51
                    }
                case "\n\t":
                    v12 += v12
                case "foo bar":
                    var v52 uint8 = 0
                    _ = v52
                }
            }
            _ = v27
        }
        _ = fi19
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 int16 = -171
        v17 = v17
        v11 = (!v11)
        _ = fi53
        _ = v54
    }
    v17 = 18446744073709551615
    if true {
        var v55 uint64 = (v17 << ((2) & 63))
        _ = len(v2)
        _ = v55
    }
    fmt.Println(v12)
    for fi56 := 0; fi56 < 16; fi56++ {
        for fi57 := 0; fi57 < 16; fi57++ {
            switch "hello" {
            case "hello":
                var v58 uint32 = 2
                _ = v58
            case "foo bar":
                fmt.Println(v18)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch (("a" + ("foo bar" + "")) + "foo bar") {
            case "hello":
                var v59 bool = true
                _ = v59
            case "a":
                var v60 uint = 2
                _ = v60
            case "foo bar":
                v11 = (771 != 14)
            }
            var v61 int = (v12 | 168)
            switch "" {
            case "\n\t":
                v18 -= v18
            case "a":
                v1 += (float64(v18) * float64((0.0 * v18)))
                fmt.Println(fi57, v1)
            case "foo bar":
                if v11 {
                    if v11 {
                        v18 += (v18 / v18)
                    }
                    v17 = (v17 % (((v17 >> ((2) & 63))) | 1))
                    if (fi56 < (fi57 + v12)) {
                        fmt.Println(v18)
                        v62 := GenS4{F0: (fi56 * (v61 + fi57)), F1: fi57}
                        var v63 uint32 = 2
                        _ = v62
                        _ = v63
                    }
                } else {
                    var v64 uint32 = 2147483648
                    if ((v64 + 495) < v64) {
                        var v65 uint8 = 127
                        _ = v65
                    } else {
                        if v11 {
                            v66 := GenS9{F0: "foo bar", F1: (42 < 0)}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v11 = ((763 <= (v17 - (v17 % ((v17) | 1)))) || false)
                            _ = v66
                        }
                        var v67 int64 = 1
                        _ = v67
                    }
                    v18 += v18
                    _ = v64
                }
            }
            _ = fi57
            _ = v61
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
