package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16; F1 uint }
func genFn0(p0 uint64, p1 uint8, p2 uint8) int {
    var v6 int = 4294967295
    var v7 float64 = float64(1e10)
    v8 := GenS1{F0: v6, F1: 32767, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v6
}

func main() {
    var v1 uint16 = 1
    var v2 int32 = 1
    var v3 int8 = 127
    var v4 uint32 = 33
    v5 := []int{990, -367, 576, -9223372036854775808}
    v3 = 42
    v4 += 331
    v3 = v3
    v9 := GenS2{F0: -1, F1: 751}
    var v10 uint64 = 405
    if false {
        var v11 string = "\n\t"
        _ = v11
    } else {
        switch 42 {
        case -73:
            var v12 float32 = 3.14159
            _ = v12
        case 32768:
            switch 2147483647 {
            case 0:
                switch ("\n\t" + "") {
                case "foo bar":
                    v4 += v4
                case "":
                    if ((v1 >> ((18446744073709551615) & 63)) != v1) {
                        switch genFn0(v10, 128, 212) {
                        case 32767:
                            for fi13 := 0; fi13 < 14; fi13++ {
                                v5 = append(v5, fi13)
                                if (true && true) {
                                    if len(v5) > 0 {
                                        v5[0] = genFn0((^(v10)), 0, 0)
                                    }
                                    if false {
                                        v14 := GenS1{F0: (fi13 / ((fi13) | 1)), F1: v1, F2: uint16(v4)}
                                        _ = v14
                                    }
                                    var v15 int16 = -451
                                    _ = v15
                                }
                                _ = fi13
                            }
                        }
                        fmt.Println(v1)
                        var v16 string = "a"
                        _ = v16
                    } else {
                        var v17 uint64 = v10
                        v2 -= v2
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v19 int64 = int64(v10)
                            var v20 uint = 647
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    }
                }
            case 849:
                if ("" == "hello") {
                    v21 := GenS2{F0: 793, F1: 354}
                    _ = v21
                }
            default:
                if ((false && (2147483647 != -734)) || (255 < 4294967295)) {
                    var v22 float32 = -2.25
                    for fi23 := 0; fi23 < 15; fi23++ {
                        switch int(v2) {
                        case -9223372036854775808:
                            var v24 int64 = -364
                            _ = v24
                        case 42:
                            v25 := GenS2{F0: 755, F1: 0}
                            _ = v25
                        }
                        if true {
                            for fi26 := 0; fi26 < 5; fi26++ {
                                v5 = append(v5, fi26)
                                v27 := GenS1{F0: genFn0(v10, 78, 1), F1: (v1 % ((v1) | 1)), F2: v1}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi26
                                _ = v27
                            }
                            if true {
                                v28 := GenS1{F0: genFn0(1, 42, 127), F1: v1, F2: v1}
                                var v29 int64 = 9223372036854775807
                                _ = v28
                                _ = v29
                            } else {
                                if ("\n\t" < ((("foo bar" + "foo bar") + "a") + "hello")) {
                                    v22 += 0.1
                                    var v30 uint32 = v4
                                    if len(v5) > 0 {
                                        v5[0] = -68
                                    }
                                    _ = v30
                                } else {
                                    v10 -= (v10 + v10)
                                    if true {
                                        for fi31 := 0; fi31 < 7; fi31++ {
                                            if (("hello" + "hello") <= "") {
                                                v5 = append(v5, (fi23 % ((int(v10)) | 1)))
                                                var v32 string = ("a" + (("foo bar" + "\n\t") + (("hello" + "") + "")))
                                                _ = v32
                                            } else {
                                                switch "a" {
                                                case "":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v2 = 1
                                                    case "hello":
                                                        v33 := GenS2{F0: int16(v10), F1: 42}
                                                        _ = v33
                                                    default:
                                                        v34 := GenS2{F0: -32768, F1: 255}
                                                        _ = v34
                                                    }
                                                default:
                                                    v1 += v1
                                                }
                                                switch (^(fi23)) {
                                                case 15:
                                                    v10 += v10
                                                case -1:
                                                    v10 = 1
                                                case 0:
                                                    v4 = (v4 / ((v4) | 1))
                                                }
                                            }
                                            _ = fi31
                                        }
                                        fmt.Println(v4)
                                        if true {
                                            if len(v5) > 0 {
                                                v5[0] = genFn0(471, 93, 255)
                                            }
                                        }
                                    }
                                    for fi35 := 0; fi35 < 5; fi35++ {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if (float32((1.5 + 3.14159)) > (v22 - v22)) {
                                            var v36 int8 = int8(v1)
                                            _ = v36
                                        }
                                        v10 = (v10 & (v10 - v10))
                                        var v37 int32 = v2
                                        if ("" == "") {
                                            v1 = (v1 | v1)
                                            if true {
                                                v38 := GenS2{F0: 1, F1: 238}
                                                if ((v1 - v1) != 0) {
                                                    v39 := GenS3{F0: fi23, F1: (!false)}
                                                    _ = v39
                                                }
                                                fmt.Println(v4, fi35, v2)
                                                var v40 uint16 = v1
                                                _ = v38
                                                _ = v40
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v37 = v2
                                                switch (fi35 << ((uint(v1)) & 63)) {
                                                case -9223372036854775808:
                                                    v4 -= (v4 >> ((18446744073709551615) & 63))
                                                case 1:
                                                    v3 += (-(v3))
                                                }
                                                if (1000 == int16(v4)) {
                                                    var v41 int32 = -1
                                                    v5 = append(v5, (fi23 - (^(fi23))))
                                                    _ = v41
                                                } else {
                                                    fmt.Println(v22)
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v42 int16 = 330
                                                        _ = v42
                                                    case "hello":
                                                        v43 := GenS3{F0: 32768, F1: false}
                                                        _ = v43
                                                    default:
                                                        var v44 int = fi35
                                                        _ = v44
                                                    }
                                                    if (((!("" == "hello")) && ("a" == "")) && true) {
                                                        v2 += (v2 % ((-1) | 1))
                                                        v45 := GenS4{F0: -93, F1: 4294967295}
                                                        _ = v45
                                                    } else {
                                                        var v46 int32 = -52
                                                        fmt.Println(fi23)
                                                        switch fi35 {
                                                        case 473:
                                                            v2 = v37
                                                        case 0:
                                                            v22 = ((1e10 * 3.14159) / v22)
                                                        case 511:
                                                            fmt.Println(v22, v10)
                                                        }
                                                        _ = v46
                                                    }
                                                }
                                            }
                                            _ = len(v5)
                                        } else {
                                            fmt.Println(v10, v3)
                                            if true {
                                                v4 += v4
                                            } else {
                                                v4 = v4
                                            }
                                            var v47 int8 = -58
                                            var v48 float32 = float32(1.5)
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = fi35
                                        _ = v37
                                    }
                                }
                                for fi49 := 0; fi49 < 18; fi49++ {
                                    var v50 int32 = (v2 >> ((2) & 63))
                                    _ = fi49
                                    _ = v50
                                }
                                fmt.Println(v22)
                            }
                            var v51 uint8 = 174
                            _ = v51
                        }
                        _ = fi23
                    }
                    _ = v22
                } else {
                    v52 := GenS2{F0: 1, F1: 2}
                    fmt.Println(v4, v3)
                    var v53 int64 = -9223372036854775808
                    _ = v52
                    _ = v53
                }
            }
        }
        fmt.Println(v3, v3)
    }
    if false {
        v4 = v4
    }
    v5 = append(v5, 1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 += (v10 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
