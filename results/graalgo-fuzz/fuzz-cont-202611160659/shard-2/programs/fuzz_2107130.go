package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 bool; F2 uint16 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int32 = -101
    var v3 float32 = 0.0
    var v4 uint = 42
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 280
        var v9 int8 = -33
        v10 := GenS1{F0: v9, F1: v9, F2: (false || true)}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 float64 = (((1e10 / 1.0) + (0.0 - ((3.14159 + (3.14159 * 1e10)) - float64((((v3 * v3) / float32(0.1)) / (((v3 * -2.25) / v3) - v3)))))) / 0.1)
        v5[fi11] = (-(fi11))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v4 = 1
        fmt.Println(v3, v4, v3)
        fmt.Println(v4, v4)
        _ = fi13
    }
    v4 = v4
    var v14 int = 1
    if (198 >= 2147483648) {
        var v15 uint8 = 127
        _ = v15
    } else {
        var v16 int64 = -9223372036854775808
        v16 = (v16 / (((v16 - v16)) | 1))
        _ = v16
    }
    var v17 float64 = float64(v3)
    if (255 < 2) {
        v14 -= 32767
        if (((float32(1e10) / (((v3 * v3) * 0.0) / (v3 + 3.14159))) <= ((v3 - 3.14159) / float32(((v17 / v17) + -2.25)))) && false) {
            var v18 uint8 = 229
            delete(v6, (^(v14)))
            _ = v18
        } else {
            var v19 uint64 = 4294967295
            if (0 < 960) {
                v20 := GenS1{F0: -128, F1: 42, F2: true}
                _ = v20
            } else {
                if (42 <= -947) {
                    if true {
                        v4 += 636
                        if (false && (!true)) {
                            v14 += (^(v14))
                        }
                    } else {
                        fmt.Println(v14)
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            _ = len(v6)
                            var v22 uint = (v4 | 666)
                            _ = v22
                        } else {
                            if false {
                                var v23 uint64 = (v19 | (^(v19)))
                                _ = v23
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                fmt.Println(v4, v2)
                                var v24 int8 = -66
                                switch (fi21 >> (((v4 + (v4 * 4294967295))) & 63)) {
                                case 4294967295:
                                    fmt.Println(v17, v3)
                                default:
                                    for fi25 := 0; fi25 < 9; fi25++ {
                                        switch "\n\t" {
                                        case "a":
                                            if (("\n\t" + "\n\t") <= "foo bar") {
                                                var v26 uint16 = 302
                                                v27 := GenS2{F0: (v24 << ((2) & 63))}
                                                if true {
                                                    v24 += (v24 | (v24 << ((v4) & 63)))
                                                    v28 := v6[(fi21 & fi21)]
                                                    _ = v28
                                                } else {
                                                    v19 = uint64(v26)
                                                    var v29 int = v14
                                                    var v30 int = (fi21 / (((v29 % (((v14 & fi21)) | 1))) | 1))
                                                    _ = v29
                                                    _ = v30
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                        case "foo bar":
                                            if ("foo bar" < "\n\t") {
                                                fmt.Println(v14, v24)
                                            }
                                        default:
                                            v19 = v19
                                        }
                                        var v31 uint32 = 4294967295
                                        var v32 float64 = v17
                                        _ = fi25
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                _ = v24
                            }
                        }
                        _ = fi21
                    }
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if true {
                            var v34 bool = (18446744073709551615 < v19)
                            for fi35 := 0; fi35 < 2; fi35++ {
                                var v36 int16 = 1
                                v37 := GenS3{F0: (v36 / ((1) | 1)), F1: (-(v14)), F2: 9223372036854775807}
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            v2 = 1
                            _ = v34
                        } else {
                            var v38 int = fi33
                            switch (v38 * (-(v38))) {
                            case 128:
                                var v39 int64 = 1
                                _ = v39
                            case 326:
                                fmt.Println(v2)
                            case 9223372036854775807:
                                v14 += 127
                            default:
                                if false {
                                    fmt.Println(v17, v19)
                                    var v40 int8 = 127
                                    _ = v40
                                }
                            }
                            _ = v38
                        }
                        if (!(32767 >= 0)) {
                            switch fi33 {
                            case 32768:
                                v4 -= v4
                            case 4294967295:
                                if (((9223372036854775807 != 1) && false) && (float32(3.14159) <= 0.0)) {
                                    if false {
                                        var v41 float64 = (v17 * float64(((1.5 * v3) + 1.5)))
                                        _ = v41
                                    } else {
                                        var v42 float32 = float32((v17 * v17))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v19 -= (v19 - 2)
                                        _ = v42
                                    }
                                    fmt.Println(v19, v4)
                                } else {
                                    fmt.Println(v17, v19, fi33)
                                }
                            case 9223372036854775807:
                                v4 -= v4
                            }
                            if true {
                                v43 := GenS1{F0: -128, F1: 1, F2: (!true)}
                                v44 := GenS1{F0: -1, F1: 0, F2: false}
                                fmt.Println(v4)
                                _ = v43
                                _ = v44
                            } else {
                                v3 -= v3
                            }
                            if false {
                                v19 = v19
                            }
                        } else {
                            if true {
                                v45 := GenS4{F0: false, F1: true, F2: uint16(v4)}
                                _ = v45
                            } else {
                                v17 = float64((1.5 + v3))
                                v46 := GenS5{F0: -1, F1: 42}
                                _ = v46
                            }
                            _ = len(v1)
                            for fi47 := 0; fi47 < 18; fi47++ {
                                var v48 uint16 = 1
                                if (!(true && (-120 > 100))) {
                                    var v49 int16 = -32768
                                    _ = v49
                                }
                                _ = fi47
                                _ = v48
                            }
                        }
                        fmt.Println(v3)
                        _ = fi33
                    }
                }
            }
            _ = v19
        }
        v2 -= v2
    } else {
        if (v17 <= v17) {
            v50 := GenS6{F0: true, F1: 1, F2: -1}
            _ = v50
        } else {
            v2 += (^(v2))
        }
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        for fi52 := 0; fi52 < 15; fi52++ {
            if (42 >= 32767) {
                v14 += (v14 / ((fi52) | 1))
                v53 := GenS5{F0: 32767, F1: 42}
                var v54 int32 = v2
                _ = v53
                _ = v54
            }
            v4 -= v4
            _ = fi52
        }
        var v55 bool = (42 == 0)
        _ = len(v1)
        _ = fi51
        _ = v55
    }
    v56 := GenS4{F0: ((v4 << ((v4) & 63)) == (^(v4))), F1: true, F2: 255}
    v14 -= (v14 | 32767)
    v57 := GenS3{F0: -454, F1: -1, F2: 1}
    var v58 uint8 = 128
    v1 = append(v1, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
