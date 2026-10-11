package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 int16 }
func genFn0() int32 {
    v6 := GenS1{F0: 1, F1: 253}
    if false {
        var v7 int = 9223372036854775807
        _ = v7
    } else {
        var v8 float64 = 1.5
        v8 = (v8 / v8)
        fmt.Println(v8)
        _ = v8
    }
    _ = v6
    return -1
}
func genFn1(p0 int32, p1 uint8) int {
    var v9 uint = 65535
    p0 = p0
    _ = p0
    _ = p1
    _ = v9
    return -567
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 uint64 = 42
    v4 := make(map[string]string)
    v5 := []string{"hello", "", ("" + ("a" + "hello"))}
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v10 float32 = (((float32(3.14159) - (-2.25 + ((1.5 - 1.5) + 1e10))) * 3.14159) + 0.0)
    v11 := GenS2{F0: 0, F1: 127}
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := v4["foo bar"]
            fmt.Println(v3)
            v3 = (v3 & 0)
            var v15 float64 = float64(((3.14159 * -2.25) / ((float32(float64(v10)) * 0.0) * v10)))
            if false {
                switch ("\n\t" + "") {
                case "a":
                    var v16 uint32 = 65535
                    _ = v16
                case "":
                    switch (((v14 + (v14 + v14)) + (v14 + v14)) + v14) {
                    case "a":
                        if false {
                            v17 := GenS2{F0: 47, F1: -1}
                            v10 -= 1e10
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v18 int64 = 1
                            fmt.Println(v10, fi12, v18)
                            v3 -= (v3 % (((v3 >> ((1) & 63))) | 1))
                            _ = v17
                            _ = v18
                        } else {
                            v3 -= 1
                        }
                    default:
                        v14 = "hello"
                    }
                case "\n\t":
                    v4[v14] = (v14 + "")
                default:
                    fmt.Println(v10, v15)
                }
                var v19 int8 = -55
                _ = v19
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v10, v3, v10)
        _ = fi12
    }
    v20 := GenS1{F0: -1, F1: 18446744073709551615}
    v2[128] = "\n\t"
    var v21 uint8 = 29
    if len(v5) > 0 {
        v5[0] = ("a" + "a")
    }
    fmt.Println(v3, v3)
    for fi22 := 0; fi22 < 8; fi22++ {
        v23 := GenS3{F0: 2147483647}
        for fi24 := 0; fi24 < 9; fi24++ {
            if true {
                var v25 int = (^(fi24))
                var v26 int32 = 755
                var v27 int16 = -18
                _ = v25
                _ = v26
                _ = v27
            } else {
                switch "hello" {
                case "foo bar":
                    if true {
                        var v28 int = fi24
                        if (v3 >= v3) {
                            fmt.Println(v3, v28)
                            v5 = append(v5, "")
                            v29 := v2[(v28 % ((42) | 1))]
                            _ = v29
                        } else {
                            _ = len(v5)
                            v30 := v1[v28]
                            v31 := GenS2{F0: (^(v21)), F1: 109}
                            _ = len(v2)
                            if true {
                                fmt.Println(fi24, v30)
                                var v32 string = v30
                                v33 := GenS3{F0: 42}
                                _ = v32
                                _ = v33
                            } else {
                                v3 = (v3 >> ((2) & 63))
                                var v34 uint8 = v21
                                v35 := GenS3{F0: 1}
                                _ = v34
                                _ = v35
                            }
                            var v36 uint8 = v21
                            _ = v30
                            _ = v31
                            _ = v36
                        }
                        v21 = v21
                        if ("\n\t" != "foo bar") {
                            if ((0.0 <= float32((1.5 / float64(v10)))) || false) {
                                var v37 int32 = -1
                                _ = v37
                            } else {
                                var v38 float64 = -2.25
                                v38 = v38
                                if (("foo bar" + ("" + "")) >= "hello") {
                                    if (true || false) {
                                        var v39 int8 = 1
                                        var v40 int16 = -845
                                        var v41 bool = false
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        switch 42 {
                                        case 32768:
                                            v42 := GenS1{F0: -1, F1: 18446744073709551615}
                                            _ = v42
                                        case 128:
                                            fmt.Println(v21)
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = ""
                                        }
                                    }
                                }
                                _ = v38
                            }
                            v21 -= v21
                            var v43 int8 = -7
                            _ = v43
                        }
                        v28 += (fi24 >> ((417) & 63))
                        v1[fi24] = ("" + "\n\t")
                        _ = v28
                    } else {
                        var v44 int64 = 878
                        var v45 uint64 = v3
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v46 := GenS1{F0: -792, F1: 42}
                            _ = v46
                        }
                        _ = v44
                        _ = v45
                    }
                default:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v3, v10)
                    default:
                        if (!false) {
                            v21 -= (v21 / (((v21 - (v21 % ((v21) | 1)))) | 1))
                        }
                    }
                }
                var v47 float64 = float64(((v10 / v10) - ((float32((1e10 - float64(v10))) - (v10 / float32(0.1))) * 3.14159)))
                v48 := GenS3{F0: -1}
                _ = v47
                _ = v48
            }
            v21 = v21
            _ = fi24
        }
        switch 32767 {
        case 127:
            switch (fi22 / ((int(v21)) | 1)) {
            case 4294967295:
                var v49 float64 = ((0.0 * (0.0 - ((1.5 / (1e10 + 0.0)) - float64((v10 / (v10 * 0.1)))))) / float64((((v10 - v10) + float32((1.5 + (1.5 / (0.1 * (0.0 + (((0.1 + (0.1 * 0.0)) / ((1e10 - 0.0) * -2.25)) * 0.1))))))) / v10)))
                _ = v49
            }
        default:
            var v50 uint8 = 42
            _ = v50
        }
        _ = fi22
        _ = v23
    }
    if false {
        v3 = v3
        v21 += v21
        fmt.Println(v21, v21)
    }
    if true {
        for fi51 := 0; fi51 < 19; fi51++ {
            _ = len(v5)
            for fi52 := 0; fi52 < 18; fi52++ {
                var v53 uint32 = 565
                _ = fi52
                _ = v53
            }
            var v54 float64 = 1e10
            _ = fi51
            _ = v54
        }
        var v55 string = ""
        _ = v55
    }
    v56 := v2[42]
    v3 += v3
    if (!(!(!true))) {
        for fi57 := 0; fi57 < 2; fi57++ {
            _ = len(v5)
            v10 -= 1.5
            _ = fi57
        }
        v10 += (float32(0.1) / v10)
        var v58 int8 = 1
        _ = v58
    } else {
        var v59 uint64 = v3
        v21 = (^(v21))
        if (786 == 2147483647) {
            for fi60 := 0; fi60 < 13; fi60++ {
                v21 -= 17
                _ = fi60
            }
            if false {
                switch (v56 + "a") {
                case "":
                    v61 := GenS4{F0: v21}
                    _ = v61
                case "foo bar":
                    if true {
                        var v62 int32 = 42
                        _ = v62
                    } else {
                        v63 := GenS1{F0: 1, F1: uint(v59)}
                        v10 = -2.25
                        fmt.Println(v3, v59)
                        _ = v63
                    }
                case "a":
                    if false {
                        var v64 uint = 12
                        _ = v64
                    }
                default:
                    var v65 uint32 = 92
                    _ = v65
                }
                v10 -= (v10 / 1e10)
                v66 := GenS5{F0: -1, F1: 4294967295}
                _ = v66
            } else {
                fmt.Println(v3)
                var v67 uint32 = 2147483647
                for fi68 := 0; fi68 < 14; fi68++ {
                    switch genFn1(genFn0(), (v21 << ((924) & 63))) {
                    case 1:
                        v5 = append(v5, v56)
                    case 2147483647:
                        v59 += 42
                    default:
                        if true {
                            if (!(v21 < (v21 / ((v21) | 1)))) {
                                fmt.Println(v3, fi68, v67)
                                v69 := GenS6{F0: (v56 + v56), F1: 823, F2: 1}
                                v70 := GenS2{F0: (v21 - v21), F1: 42}
                                _ = v69
                                _ = v70
                            } else {
                                delete(v4, "")
                            }
                            v4[v56] = ((("a" + "foo bar") + v56) + v56)
                        } else {
                            v71 := GenS4{F0: v21}
                            var v72 int16 = 234
                            _ = v71
                            _ = v72
                        }
                    }
                    _ = fi68
                }
                _ = v67
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v73 uint16 = 2
            var v74 int64 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v73
            _ = v74
        }
        _ = v59
    }
    _ = len(v5)
    if (42 >= 102) {
        if (100 != 42) {
            var v75 int64 = 609
            v10 += v10
            _ = v75
        }
        v76 := GenS2{F0: v21, F1: 100}
        v56 = v56
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
