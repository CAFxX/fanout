package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 int64 = 9223372036854775807
    var v4 uint = 1
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    v7 := []int{32767}
    v3 = 1
    var v8 float64 = 3.14159
    var v9 uint8 = 2
    switch "hello" {
    case "foo bar":
        if false {
            switch 1 {
            case 32768:
                switch "" {
                case "foo bar":
                    fmt.Println(v9, v4, v9)
                case "a":
                    fmt.Println(v4, v8, v2)
                case "hello":
                    v3 -= 1
                default:
                    if (((v4 >> ((v4) & 63)) > (v4 + (v4 << ((858) & 63)))) || (((18446744073709551615 <= 42) && (!((v3 + v3) != (v3 >> ((v4) & 63))))) && (3.14159 > float32(v8)))) {
                        v10 := GenS1{F0: 2, F1: 11, F2: v3}
                        v11 := v1[-9223372036854775808]
                        _ = v10
                        _ = v11
                    } else {
                        v2 = (!v2)
                        v12 := GenS1{F0: 758, F1: 699, F2: 42}
                        if (!true) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        _ = v12
                    }
                }
            default:
                var v13 float64 = v8
                _ = v13
            }
            v2 = false
        }
    case "a":
        switch 0 {
        case 42:
            var v14 float64 = v8
            _ = v14
        case -1:
            v2 = ((v4 >> ((1) & 63)) > v4)
        case 32767:
            v15 := GenS2{F0: v3}
            _ = v15
        }
    }
    switch 127 {
    case 2147483647:
        v1[int(v9)] = 42
    default:
        var v16 uint8 = (v9 % ((v9) | 1))
        _ = v16
    }
    if (-63 <= 17) {
        v17 := GenS1{F0: 678, F1: 18446744073709551615, F2: (-(v3))}
        if v2 {
            fmt.Println(v4, v4, v3)
            if true {
                v18 := GenS3{F0: (true || (127 <= -9223372036854775808))}
                for fi19 := 0; fi19 < 1; fi19++ {
                    if (((^(fi19)) < int(v3)) && v2) {
                        fmt.Println(v2)
                        v7 = append(v7, 2147483647)
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v3 = (v3 * v3)
                            fmt.Println(v3)
                            var v21 int16 = 1000
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 uint = (^(v4))
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 int16 = 0
                        v3 = (v3 * (v3 / (((^(v3))) | 1)))
                        _ = fi23
                        _ = v24
                    }
                    _ = fi19
                    _ = v22
                }
                _ = v18
            } else {
                for fi25 := 0; fi25 < 9; fi25++ {
                    var v26 bool = (32767 != 255)
                    if v2 {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            if ((0.1 * (v8 + 1.5)) > 1e10) {
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        v3 += v3
                                        v7 = append(v7, fi27)
                                        if (("" >= ("hello" + "a")) || v2) {
                                            switch (fi27 + 2147483647) {
                                            case 2147483647:
                                                var v28 int64 = -743
                                                _ = v28
                                            default:
                                                fmt.Println(v3)
                                            }
                                            v2 = (float32(float64(1e10)) <= (3.14159 * (-2.25 / (1.5 * (0.0 - float32(float64(0.1)))))))
                                            v29 := GenS1{F0: 42, F1: 872, F2: (v3 - -1)}
                                            _ = v29
                                        } else {
                                            v8 = v8
                                            if v26 {
                                                var v30 uint32 = 42
                                                var v31 int8 = 42
                                                _ = v30
                                                _ = v31
                                            } else {
                                                if true {
                                                    delete(v5, (("\n\t" + "foo bar") + ("\n\t" + "a")))
                                                    var v32 int64 = (^(v3))
                                                    _ = v32
                                                } else {
                                                    var v33 int16 = 1000
                                                    v4 = (v4 ^ (v4 + 340))
                                                    _ = v33
                                                }
                                            }
                                        }
                                        if v2 {
                                            switch "" {
                                            case "a":
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            v2 = true
                                        }
                                    } else {
                                        if ("foo bar" < "") {
                                            var v34 int = (fi25 - fi25)
                                            var v35 bool = (87 < 35)
                                            _ = v34
                                            _ = v35
                                        }
                                        if v2 {
                                            if (-444 > fi25) {
                                                v4 += (v4 * (v4 & v4))
                                            } else {
                                                v5[(("\n\t" + ("\n\t" + "")) + ("hello" + "\n\t"))] = (("a" + (("hello" + "foo bar") + "foo bar")) + "hello")
                                            }
                                            v36 := GenS1{F0: 42, F1: 705, F2: 9223372036854775807}
                                            if v26 {
                                                v4 = (^(v4))
                                                var v37 int8 = int8(fi27)
                                                _ = v37
                                            }
                                            var v38 string = "hello"
                                            var v39 uint32 = 65535
                                            _ = v36
                                            _ = v38
                                            _ = v39
                                        } else {
                                            var v40 float64 = v8
                                            _ = v40
                                        }
                                    }
                                case "a":
                                    if false {
                                        var v41 int16 = -586
                                        _ = v41
                                    } else {
                                        v42 := v1[fi27]
                                        _ = v42
                                    }
                                case "":
                                    var v43 float64 = (0.0 - ((v8 - (v8 / -2.25)) - v8))
                                    var v44 uint16 = 1
                                    _ = v43
                                    _ = v44
                                default:
                                    v8 = v8
                                }
                                switch (fi27 * (fi25 ^ (fi27 - (fi27 + 0)))) {
                                case 0:
                                    var v45 uint8 = 180
                                    _ = v45
                                default:
                                    _ = len(v1)
                                }
                            }
                            _ = fi27
                        }
                        var v46 float64 = float64(-2.25)
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            if ((fi47 % (((fi25 * (fi25 + fi47))) | 1)) != 6) {
                                v48 := GenS1{F0: 1, F1: 299, F2: (v3 - v3)}
                                if false {
                                    v4 = v4
                                    fmt.Println(fi25, fi25)
                                    if v2 {
                                        v9 = v9
                                        v9 += (^(v9))
                                        var v49 uint8 = v9
                                        _ = v49
                                    }
                                    var v50 uint64 = 1
                                    fmt.Println(v9, v2)
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "hello":
                                            if (v50 < v50) {
                                                switch ("" + "") {
                                                case "\n\t":
                                                    if (true || true) {
                                                        v26 = false
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi25 / ((fi25) | 1))
                                                        }
                                                        v50 = (v50 - 4294967295)
                                                    }
                                                case "":
                                                    var v51 float32 = (-2.25 / (3.14159 * (-2.25 / (0.1 / 1e10))))
                                                    _ = v51
                                                default:
                                                    fmt.Println(v8)
                                                }
                                                switch ("foo bar" + (((("a" + "foo bar") + "") + ("" + ("hello" + (("hello" + "hello") + ("a" + ("hello" + "foo bar")))))) + ("foo bar" + "hello"))) {
                                                case "foo bar":
                                                    var v52 string = ""
                                                    _ = v52
                                                case "\n\t":
                                                    v53 := GenS3{F0: v26}
                                                    _ = v53
                                                case "":
                                                    var v54 int64 = (v3 & v3)
                                                    _ = v54
                                                default:
                                                    if len(v7) > 0 {
                                                        v7[0] = (fi47 * -9223372036854775808)
                                                    }
                                                }
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v2 = ("hello" == ("\n\t" + ""))
                                                }
                                            }
                                        default:
                                            v9 -= (v9 | 207)
                                        }
                                        _ = len(v7)
                                        var v55 uint = v4
                                        v56 := GenS3{F0: v2}
                                        switch 2147483647 {
                                        case -826:
                                            var v57 uint8 = 1
                                            _ = v57
                                        case -1:
                                            var v58 int64 = (v3 >> ((v55) & 63))
                                            var v59 int8 = 100
                                            var v60 uint32 = 393
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        case 32768:
                                            switch ("foo bar" + ("" + (("foo bar" + "\n\t") + "foo bar"))) {
                                            case "foo bar":
                                                v2 = ((!(!true)) && v2)
                                                v26 = (((-2.25 - 1e10) / 0.1) == 1.5)
                                            }
                                        }
                                        _ = v55
                                        _ = v56
                                    case "a":
                                        v7 = append(v7, fi25)
                                    case "":
                                        var v61 uint = v4
                                        _ = v61
                                    }
                                    var v62 int16 = 1000
                                    _ = v50
                                    _ = v62
                                }
                                var v63 uint8 = v9
                                v64 := GenS4{F0: 8, F1: 18446744073709551615}
                                var v65 uint64 = 1
                                _ = v48
                                _ = v63
                                _ = v64
                                _ = v65
                            } else {
                                var v66 string = "hello"
                                var v67 bool = (65535 <= 690)
                                var v68 int = (fi25 - (fi47 + fi25))
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            var v69 uint16 = 42
                            _ = fi47
                            _ = v69
                        }
                        if v26 {
                            var v70 int8 = 1
                            _ = v70
                        } else {
                            v71 := GenS4{F0: 100, F1: 42}
                            var v72 int = fi25
                            v7 = append(v7, (fi25 << ((v4) & 63)))
                            _ = v71
                            _ = v72
                        }
                        for fi73 := 0; fi73 < 1; fi73++ {
                            _ = len(v7)
                            if v26 {
                                fmt.Println(fi73)
                            }
                            var v74 int64 = int64(fi73)
                            _ = fi73
                            _ = v74
                        }
                    }
                    v75 := GenS5{F0: 1000}
                    _ = fi25
                    _ = v26
                    _ = v75
                }
                if v2 {
                    delete(v1, -514)
                    _ = len(v7)
                }
                var v76 float32 = 0.0
                _ = v76
            }
        } else {
            v9 = (v9 ^ 127)
            v3 = v3
        }
        var v77 string = "\n\t"
        _ = v17
        _ = v77
    }
    var v78 int32 = 0
    fmt.Println(v8)
    var v79 uint16 = uint16(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
    _ = v79
    _ = v8
    _ = v9
}
