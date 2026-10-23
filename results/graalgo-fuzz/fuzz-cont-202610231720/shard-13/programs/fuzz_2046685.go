package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int16; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int }
func genFn0(p0 uint8, p1 string, p2 bool) uint {
    fmt.Println(p0, p1)
    var v4 uint64 = 42
    v5 := GenS1{F0: 281}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 931
}

func main() {
    var v1 uint8 = 2
    var v2 bool = true
    var v3 int16 = -127
    if ((("foo bar" + "\n\t") < ("hello" + "hello")) || true) {
        v6 := GenS1{F0: 900}
        v1 = v1
        var v7 int8 = -107
        _ = v6
        _ = v7
    }
    var v8 uint32 = 42
    if (!false) {
        fmt.Println(v3, v3, v8)
        var v9 float32 = 3.14159
        switch ("\n\t" + "\n\t") {
        case "":
            var v10 uint = 1
            _ = v10
        case "hello":
            if (!(v3 >= 1000)) {
                var v11 uint64 = 839
                _ = v11
            }
        default:
            v1 += (v1 + 42)
        }
        _ = v9
    }
    if v2 {
        switch "foo bar" {
        case "a":
            switch 913 {
            case 2147483647:
                v12 := []int{}
                _ = v12
            case -355:
                fmt.Println(v2, v8, v1)
            case 42:
                var v13 float64 = 1e10
                _ = v13
            default:
                v14 := GenS1{F0: v8}
                _ = v14
            }
        case "":
            v15 := GenS2{F0: 2147483647, F1: (v8 * v8)}
            _ = v15
        default:
            fmt.Println(v2, v8, v3)
        }
        v16 := []int{9223372036854775807, 127, 0, -1}
        if false {
            if ("hello" != "") {
                var v17 float64 = 0.0
                v18 := GenS3{F0: "\n\t"}
                if (65535 <= (v8 << ((660) & 63))) {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if ((65535 > 32767) && ((v2 && v2) || false)) {
                            v20 := GenS2{F0: -1, F1: v8}
                            v21 := GenS4{F0: int32(v3)}
                            v17 = float64(1.5)
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                } else {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        switch fi22 {
                        case 128:
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        case 2147483647:
                            v8 -= (^(v8))
                        case 32768:
                            fmt.Println(v3, fi22, v3)
                        }
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 17; fi23++ {
                        if len(v16) > 0 {
                            v16[0] = fi23
                        }
                        v17 += ((float64(0.1) - (v17 / 0.1)) / float64((-2.25 + (0.0 - (1.5 - 1e10)))))
                        if (42 == -1) {
                            if ((1 < -1) && v2) {
                                v17 += (float64(1e10) - 1e10)
                                if v2 {
                                    if v2 {
                                        v17 = (v17 + ((1.5 / float64(3.14159)) + (((float64(3.14159) - v17) + ((float64(1.5) + ((v17 / 0.1) / v17)) / v17)) * (v17 + 0.1))))
                                    }
                                    var v24 int8 = -107
                                    if v2 {
                                        if true {
                                            v25 := GenS5{F0: (fi23 & (fi23 | fi23)), F1: (v3 & (v3 * v3)), F2: (v24 << ((42) & 63))}
                                            if len(v16) > 0 {
                                                v16[0] = (fi23 % (((fi23 % (((-(fi23))) | 1))) | 1))
                                            }
                                            if len(v16) > 0 {
                                                v16[0] = (fi23 - (fi23 ^ -1))
                                            }
                                            _ = v25
                                        } else {
                                            v24 += (v24 | (^(v24)))
                                            if (((3.14159 * (0.0 + (0.0 + ((-2.25 + float32(3.14159)) - float32((0.1 + ((1.5 + (float64(0.1) / ((3.14159 * (float64(3.14159) - -2.25)) * v17))) + 1e10))))))) - -2.25) <= 1.5) {
                                                v1 -= (^(v1))
                                                v26 := GenS3{F0: "\n\t"}
                                                for fi27 := 0; fi27 < 3; fi27++ {
                                                    v28 := GenS4{F0: -2147483648}
                                                    switch "" {
                                                    case "a":
                                                        if v2 {
                                                            switch fi27 {
                                                            case 4294967295:
                                                                var v29 bool = v2
                                                                _ = v29
                                                            }
                                                            if true {
                                                                var v30 int = (fi27 - (fi27 & fi23))
                                                                switch ("hello" + ("" + "foo bar")) {
                                                                case "\n\t":
                                                                    v31 := GenS6{F0: -9223372036854775808}
                                                                    _ = v31
                                                                }
                                                                _ = v30
                                                            }
                                                        } else {
                                                            if len(v16) > 0 {
                                                                v16[0] = 625
                                                            }
                                                        }
                                                    }
                                                    var v32 int64 = 789
                                                    _ = fi27
                                                    _ = v28
                                                    _ = v32
                                                }
                                                _ = v26
                                            } else {
                                                if len(v16) > 0 {
                                                    v16[0] = fi23
                                                }
                                                for fi33 := 0; fi33 < 11; fi33++ {
                                                    v8 -= v8
                                                    if false {
                                                        v34 := GenS7{F0: v2, F1: int16(v24)}
                                                        _ = v34
                                                    }
                                                    if len(v16) > 0 {
                                                        v16[0] = (^(fi23))
                                                    }
                                                    _ = fi33
                                                }
                                            }
                                            v8 = 357
                                        }
                                        if len(v16) > 0 {
                                            v16[0] = (fi23 + (fi23 << ((65535) & 63)))
                                        }
                                        for fi35 := 0; fi35 < 9; fi35++ {
                                            var v36 float32 = (((((1e10 + 1e10) + 0.1) - 1.5) / 1.0) - ((0.1 * 3.14159) + ((float32(v17) / -2.25) + ((0.1 + float32(0.0)) + (0.0 / 1.5)))))
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                    }
                                    _ = v24
                                } else {
                                    if len(v16) > 0 {
                                        v16[0] = fi23
                                    }
                                }
                            }
                            var v37 uint64 = 465
                            _ = v37
                        }
                        _ = fi23
                    }
                }
                _ = v17
                _ = v18
            }
        }
        _ = v16
    }
    v8 += v8
    for fi38 := 0; fi38 < 18; fi38++ {
        if (892 > 42) {
            if (-110 == 1) {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    if v2 {
                        switch ("" + ("\n\t" + "foo bar")) {
                        case "hello":
                            if ((!v2) || (v3 > (v3 - -32768))) {
                                v39 := []int{fi38}
                                v39 = append(v39, 0)
                                _ = v39
                            } else {
                                fmt.Println(v2)
                            }
                        case "\n\t":
                            var v40 float32 = 0.0
                            _ = v40
                        default:
                            for fi41 := 0; fi41 < 7; fi41++ {
                                v8 -= 4294967295
                                _ = fi41
                            }
                        }
                        v8 = v8
                        var v42 uint64 = 0
                        _ = v42
                    }
                case "a":
                    var v43 int32 = int32(v1)
                    _ = v43
                case "\n\t":
                    if v2 {
                        if (42 != -1) {
                            for fi44 := 0; fi44 < 17; fi44++ {
                                v45 := []int{fi44}
                                var v46 uint32 = (v8 ^ v8)
                                if v2 {
                                    v47 := GenS8{F0: 2}
                                    v8 = (v46 >> ((42) & 63))
                                    _ = v47
                                } else {
                                    v3 = (v3 - -1)
                                    v48 := GenS9{F0: 1}
                                    fmt.Println(fi44, v1)
                                    _ = v48
                                }
                                if (!true) {
                                    var v49 int8 = -1
                                    if false {
                                        var v50 float64 = (1e10 + ((1.5 - (0.1 - (-2.25 + -2.25))) - -2.25))
                                        var v51 int32 = -381
                                        var v52 uint8 = uint8(v8)
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    } else {
                                        if ((true || v2) && false) {
                                            var v53 string = ((("hello" + "hello") + "\n\t") + ("" + "a"))
                                            fmt.Println(v46)
                                            _ = v53
                                        }
                                    }
                                    v49 = (v49 / ((v49) | 1))
                                    _ = v49
                                } else {
                                    v3 -= (-(v3))
                                    v46 = 231
                                }
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                            if (888 >= 32767) {
                                var v54 uint8 = 49
                                fmt.Println(fi38, v2)
                                var v55 int = (fi38 % ((fi38) | 1))
                                _ = v54
                                _ = v55
                            }
                        }
                    } else {
                        v56 := GenS2{F0: -80, F1: v8}
                        v8 += (v8 / (((v8 >> ((4294967295) & 63))) | 1))
                        switch (fi38 * fi38) {
                        case 0:
                            if v2 {
                                v1 = (^(v1))
                                v2 = (-873 == -12)
                            } else {
                                fmt.Println(v2, v3)
                            }
                        case -2147483648:
                            fmt.Println(fi38)
                        case 127:
                            v57 := []int{}
                            _ = v57
                        default:
                            v1 = 1
                        }
                        _ = v56
                    }
                }
            } else {
                var v58 int32 = 0
                _ = v58
            }
            if (0.0 >= (0.0 / float32(1e10))) {
                for fi59 := 0; fi59 < 1; fi59++ {
                    v60 := []int{(fi38 | (fi38 / ((int(v1)) | 1))), (^(fi38)), fi59}
                    _ = fi59
                    _ = v60
                }
                for fi61 := 0; fi61 < 2; fi61++ {
                    var v62 uint = 0
                    v2 = v2
                    v63 := GenS10{F0: (fi38 / ((fi61) | 1))}
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                v2 = v2
            }
            fmt.Println(v2, v2, v2)
        } else {
            for fi64 := 0; fi64 < 19; fi64++ {
                v3 -= v3
                _ = fi64
            }
            for fi65 := 0; fi65 < 12; fi65++ {
                v8 -= uint32(v1)
                v66 := GenS3{F0: ""}
                _ = fi65
                _ = v66
            }
            switch 32767 {
            case -9223372036854775808:
                fmt.Println(v1, v1)
            }
        }
        _ = fi38
    }
    if (!(((!true) && false) && (!(42 > 461)))) {
        var v67 bool = false
        fmt.Println(v67, v8)
        _ = v67
    } else {
        for fi68 := 0; fi68 < 9; fi68++ {
            var v69 float32 = ((0.1 - 3.14159) * (0.0 / 1.0))
            var v70 float64 = 0.0
            _ = fi68
            _ = v69
            _ = v70
        }
    }
    v71 := GenS6{F0: -359}
    v1 = 255
    var v72 float32 = -2.25
    v73 := GenS1{F0: (v8 * 2147483647)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v72:", v72)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v73
    _ = v8
}
