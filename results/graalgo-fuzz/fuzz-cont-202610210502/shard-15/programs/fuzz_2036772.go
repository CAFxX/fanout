package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 int16 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int8) bool {
    v7 := GenS1{F0: -145, F1: 0, F2: -891}
    var v8 float32 = (1.5 - -2.25)
    _ = p0
    _ = v7
    _ = v8
    return (65535 < 0)
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = float64((3.14159 + 0.1))
    v3 := []int{828, 2147483647, 127}
    v4 := []int{9223372036854775807, 0}
    v5 := make(map[int]int)
    var v6 uint8 = uint8(v1)
    v9 := GenS1{F0: 6, F1: 1, F2: 42}
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 = (float64(1e10) - (v2 / 1.5))
        v11 := GenS2{F0: 2147483648, F1: v1}
        _ = fi10
        _ = v11
    }
    var v12 float32 = (1e10 - 1e10)
    if false {
        switch ((("a" + ("hello" + "hello")) + "") + "\n\t") {
        case "hello":
            var v13 uint32 = 645
            _ = v13
        case "\n\t":
            v14 := GenS2{F0: 2, F1: (^(v1))}
            _ = v14
        default:
            for fi15 := 0; fi15 < 5; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    if false {
                        var v17 uint = uint(v6)
                        _ = v17
                    }
                    v18 := GenS3{F0: (v1 % ((v1) | 1)), F1: ("" + ("foo bar" + "a"))}
                    switch (fi16 / (((fi15 | (fi16 << ((0) & 63)))) | 1)) {
                    case 992:
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        var v19 uint16 = (^(v1))
                        _ = v19
                    default:
                        if (32767 >= 1000) {
                            v20 := GenS4{F0: 1}
                            switch "a" {
                            case "hello":
                                var v21 int = (^(fi15))
                                _ = v21
                            case "\n\t":
                                fmt.Println(v1)
                            default:
                                v2 += (v2 - ((float64(float32(v2)) - (((3.14159 * 3.14159) / v2) - float64((v12 + v12)))) * 3.14159))
                            }
                            _ = v20
                        }
                    }
                    _ = fi16
                    _ = v18
                }
                if (4294967295 < 4294967295) {
                    v6 = (v6 >> ((385) & 63))
                } else {
                    var v22 bool = (!false)
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 float64 = 1e10
                        v2 -= v2
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS3{F0: v1, F1: ("hello" + "foo bar")}
                    _ = v22
                    _ = v25
                }
                _ = fi15
            }
        }
        var v26 uint64 = 989
        switch int(v1) {
        case -45:
            _ = len(v4)
        }
        _ = v26
    }
    if true {
        if (-723 != 42) {
            var v27 int = 128
            _ = v27
        }
    } else {
        if (!false) {
            v28 := GenS5{F0: -40, F1: 2147483647}
            v12 = ((float32((v2 + 0.1)) * (1e10 - ((v12 - v12) / v12))) + 1e10)
            if true {
                if ((true && ((-1 <= 352) || true)) || false) {
                    v2 = v2
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if (1000 != -32768) {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                v6 = v6
                                var v30 uint8 = 128
                                for fi31 := 0; fi31 < 18; fi31++ {
                                    v6 = 128
                                    _ = fi31
                                }
                                _ = fi29
                                _ = v30
                            }
                            var v32 float64 = (v2 - v2)
                            v33 := GenS5{F0: 72, F1: 4294967295}
                            _ = v32
                            _ = v33
                        } else {
                            var v34 float64 = 0.0
                            fmt.Println(v34, v6, v34)
                            if true {
                                if true {
                                    v2 -= v2
                                } else {
                                    v5[128] = 798
                                }
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    var v36 uint8 = (^(v6))
                                    v4 = append(v4, (fi35 / ((fi35) | 1)))
                                    var v37 uint = 371
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                v38 := GenS6{F0: -1, F1: -1, F2: 1}
                                _ = v38
                            }
                            _ = v34
                        }
                        v39 := GenS4{F0: 32767}
                        _ = v39
                    }
                } else {
                    var v40 uint = 464
                    _ = v40
                }
                if true {
                    if ((!(v6 <= v6)) && (!false)) {
                        var v41 string = ((("\n\t" + "") + "\n\t") + ("\n\t" + "foo bar"))
                        switch ("a" + v41) {
                        case "\n\t":
                            fmt.Println(v41, v1, v1)
                        case "a":
                            v1 += v1
                        case "":
                            var v42 uint32 = 0
                            _ = v42
                        }
                        _ = v41
                    } else {
                        v43 := GenS3{F0: 1, F1: "foo bar"}
                        var v44 int64 = -431
                        v5[32768] = 42
                        _ = v43
                        _ = v44
                    }
                    switch 1 {
                    case 128:
                        delete(v5, 32767)
                    case 32768:
                        switch 343 {
                        case 4294967295:
                            for fi45 := 0; fi45 < 15; fi45++ {
                                fmt.Println(fi45, v12)
                                v46 := v5[fi45]
                                _ = fi45
                                _ = v46
                            }
                        case 359:
                            switch "a" {
                            case "hello":
                                if false {
                                    v5[-1] = 627
                                } else {
                                    var v47 uint8 = 127
                                    var v48 bool = (!true)
                                    v12 -= v12
                                    _ = v47
                                    _ = v48
                                }
                            default:
                                v49 := v5[-253]
                                _ = v49
                            }
                        default:
                            v50 := v5[0]
                            _ = v50
                        }
                    case 9223372036854775807:
                        v2 -= v2
                    }
                }
            }
            _ = v28
        }
        v6 -= (v6 ^ 86)
    }
    if true {
        var v51 uint8 = (v6 & v6)
        _ = v51
    } else {
        if (0.0 >= 3.14159) {
            var v52 bool = (v6 == v6)
            _ = v52
        }
        var v53 int8 = 82
        fmt.Println(v53, v6, v2)
        _ = v53
    }
    v54 := GenS7{F0: 18446744073709551615}
    if (!true) {
        var v55 int32 = 1
        delete(v5, 354)
        v56 := GenS3{F0: (^(v1)), F1: "hello"}
        _ = v55
        _ = v56
    }
    v2 = 0.1
    for fi57 := 0; fi57 < 12; fi57++ {
        v6 = (v6 - (v6 - v6))
        fmt.Println(v1, v6, v6)
        v58 := GenS5{F0: -95, F1: 0}
        _ = fi57
        _ = v58
    }
    if (!(!false)) {
        delete(v5, -2147483648)
        var v59 string = ("a" + "\n\t")
        _ = v59
    }
    v60 := GenS6{F0: -2147483648, F1: 670, F2: 1000}
    fmt.Println(v1)
    if false {
        var v61 uint32 = 440
        var v62 uint64 = 18446744073709551615
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v61
        _ = v62
    } else {
        v1 = v1
        v6 = (^(v6))
    }
    _ = len(v5)
    if ((42 <= v6) || (!false)) {
        v6 = v6
        fmt.Println(v1, v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v60
    _ = v9
}
