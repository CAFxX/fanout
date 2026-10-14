package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 string; F2 uint64 }
type GenS6 struct { F0 int32; F1 int16; F2 uint16 }
type GenS7 struct { F0 int; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 uint32 }
func genFn0(p0 int16, p1 int32, p2 int8) int64 {
    p1 = p1
    p2 = 1
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int8, p1 uint16, p2 int16) uint32 {
    var v7 int32 = -377
    if ((-(p0)) == (-(p0))) {
        v8 := []int{719, 2147483647, 127, -2147483648}
        switch (("hello" + (("a" + "") + "hello")) + "\n\t") {
        case "foo bar":
            if false {
                v9 := GenS1{F0: 885, F1: 2, F2: p2}
                v10 := GenS2{F0: ""}
                _ = v10
                _ = v9
            }
        case "a":
            if false {
                var v11 uint32 = 4294967295
                _ = v11
            } else {
                v12 := GenS3{F0: -922, F1: "", F2: (p2 % ((p2) | 1))}
                v13 := GenS4{F0: v7}
                _ = v12
                _ = v13
            }
        case "hello":
            for fi14 := 0; fi14 < 17; fi14++ {
                for fi15 := 0; fi15 < 17; fi15++ {
                    if true {
                        var v16 int8 = (p0 >> ((571) & 63))
                        _ = v16
                    } else {
                        if ((v7 - v7) == (v7 + int32(p2))) {
                            v17 := GenS5{F0: 917, F1: "a", F2: 1}
                            var v18 string = ""
                            var v19 float64 = (0.1 - 0.1)
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        var v20 int32 = -2147483648
                        switch fi14 {
                        case 32768:
                            _ = len(v8)
                        case -36:
                            switch fi14 {
                            case 128:
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(p2, v7, v20)
                                }
                            }
                        }
                        _ = v20
                    }
                    _ = fi15
                }
                v7 = v7
                if false {
                    var v21 float64 = -2.25
                    fmt.Println(fi14)
                    if true {
                        var v22 bool = (42 < 0)
                        _ = v22
                    } else {
                        if (((false && true) && true) || (!false)) {
                            var v23 uint = 65535
                            v23 -= 2
                            if (-9223372036854775808 != -1) {
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    v23 += (v23 & (v23 & (v23 - v23)))
                                    var v25 bool = (-2.25 <= v21)
                                    switch "" {
                                    case "\n\t":
                                        if len(v8) > 0 {
                                            v8[0] = (fi14 | fi14)
                                        }
                                    case "":
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    default:
                                        v26 := GenS6{F0: (v7 ^ (v7 % ((0) | 1))), F1: (p2 - 32767), F2: (p1 + 32767)}
                                        _ = v26
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                            }
                            _ = v23
                        } else {
                            if (555 <= 280) {
                                switch fi14 {
                                case 765:
                                    if (128 > 127) {
                                        if false {
                                            v27 := GenS7{F0: (fi14 ^ 32767), F1: 2147483648, F2: (v7 << ((248) & 63))}
                                            if (953 == 2147483647) {
                                                fmt.Println(fi14, p0)
                                                for fi28 := 0; fi28 < 12; fi28++ {
                                                    v29 := GenS6{F0: (v7 * (v7 / ((v7) | 1))), F1: (p2 / (((p2 | p2)) | 1)), F2: (p1 ^ p1)}
                                                    _ = fi28
                                                    _ = v29
                                                }
                                            }
                                            _ = v27
                                        } else {
                                            p0 = int8(fi14)
                                            v21 -= (v21 / 1.5)
                                            p1 = p1
                                        }
                                    }
                                case -2147483648:
                                    v21 = v21
                                default:
                                    p0 -= 100
                                }
                            } else {
                                _ = len(v8)
                                for fi30 := 0; fi30 < 20; fi30++ {
                                    if true {
                                        var v31 int32 = 42
                                        _ = v31
                                    }
                                    _ = fi30
                                }
                            }
                            p1 += (p1 >> ((0) & 63))
                        }
                        v8 = append(v8, (fi14 % ((fi14) | 1)))
                        v32 := GenS8{F0: (p1 << ((255) & 63)), F1: (fi14 >> ((869) & 63))}
                        _ = v32
                    }
                    _ = v21
                } else {
                    _ = len(v8)
                }
                _ = fi14
            }
        }
        if (true || (-9223372036854775808 == 32767)) {
            for fi33 := 0; fi33 < 17; fi33++ {
                p1 += p1
                v7 = (^(v7))
                v34 := GenS9{F0: v7}
                _ = fi33
                _ = v34
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        } else {
            var v35 uint32 = 955
            v36 := GenS10{F0: 1}
            _ = v35
            _ = v36
        }
        _ = v8
    } else {
        v37 := []int{-987, 0, 2147483647}
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 375
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.1
    var v3 uint = 4294967295
    var v4 uint8 = 10
    v5 := make(map[uint32]uint32)
    var v6 int64 = 0
    if (v2 < 3.14159) {
        v6 = v6
        switch 127 {
        case 32768:
            switch 128 {
            case 2147483647:
                if (false && (v3 == 0)) {
                    fmt.Println(v6)
                }
            case -1:
                v2 += 3.14159
            case 453:
                v1["hello"] = 871
            }
        default:
            switch "\n\t" {
            case "\n\t":
                _ = len(v5)
            case "foo bar":
                var v38 uint8 = 0
                _ = v38
            case "a":
                v6 = genFn0(32767, 685, 1)
            }
        }
        v3 = 94
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 float64 = 1.5
        var v41 int64 = (v6 + (v6 << (((^(v3))) & 63)))
        v42 := GenS11{F0: -1, F1: 42}
        _ = fi39
        _ = v40
        _ = v41
        _ = v42
    }
    v1[("hello" + "a")] = 4294967295
    var v43 int16 = -1
    v1["foo bar"] = 2147483647
    fmt.Println(v3, v3, v2)
    switch 127 {
    case 32767:
        for fi44 := 0; fi44 < 10; fi44++ {
            var v45 uint64 = 157
            v46 := GenS10{F0: genFn0((v43 % (((v43 % (((v43 >> ((v3) & 63))) | 1))) | 1)), 396, -54)}
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    v47 := GenS11{F0: -39, F1: 931}
    v5[1] = 406
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v5
    _ = v6
}
