package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int16; F2 uint }

func main() {
    var v1 bool = (false && false)
    v2 := []int{1, 9223372036854775807, 397, 32768}
    v3 := make(map[uint32]int)
    v1 = v1
    v4 := GenS1{F0: 127, F1: 1}
    if v1 {
        fmt.Println(v1)
    } else {
        var v5 uint32 = 2147483647
        var v6 int8 = 0
        _ = v5
        _ = v6
    }
    var v7 uint32 = 2147483648
    if v1 {
        for fi8 := 0; fi8 < 9; fi8++ {
            _ = len(v2)
            v1 = false
            var v9 int16 = 675
            _ = fi8
            _ = v9
        }
        fmt.Println(v1)
    }
    v3[v7] = 128
    v7 += (v7 - v7)
    var v10 int8 = 42
    for fi11 := 0; fi11 < 18; fi11++ {
        switch "\n\t" {
        case "a":
            switch fi11 {
            case 392:
                var v12 uint8 = 136
                _ = v12
            }
        case "\n\t":
            var v13 int8 = (v10 * v10)
            _ = v13
        }
        _ = fi11
    }
    switch "\n\t" {
    case "\n\t":
        _ = len(v3)
    }
    _ = len(v2)
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: 1}
        for fi16 := 0; fi16 < 15; fi16++ {
            if (v1 && v1) {
                var v17 uint = 962
                v18 := GenS1{F0: (^(v10)), F1: v10}
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 1, F1: v10}
                var v20 uint16 = uint16(v7)
                fmt.Println(fi14, v10, v1)
                _ = v19
                _ = v20
            }
            _ = fi16
        }
        var v21 uint16 = 1
        _ = fi14
        _ = v15
        _ = v21
    }
    if v1 {
        if v1 {
            if (255 != 961) {
                delete(v3, (^(v7)))
                v22 := GenS3{F0: 993, F1: 42, F2: -618}
                switch 1 {
                case -2147483648:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 195:
                    v1 = ((float64(float32(0.0)) + float64((((0.0 - 0.0) / -2.25) - ((0.1 + 0.1) / (1e10 + (1e10 + (0.1 + 3.14159))))))) >= 1e10)
                case 9223372036854775807:
                    var v23 uint = 868
                    _ = v23
                }
                _ = v22
            } else {
                var v24 float32 = 3.14159
                fmt.Println(v1)
                for fi25 := 0; fi25 < 11; fi25++ {
                    if len(v2) > 0 {
                        v2[0] = (fi25 / ((fi25) | 1))
                    }
                    if (false || (((-1 != 1) || v1) && (!true))) {
                        var v26 uint16 = 255
                        _ = v26
                    } else {
                        if false {
                            v10 = v10
                            if (4294967295 != uint64(v7)) {
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    var v28 uint8 = 128
                                    v10 -= v10
                                    if (((!v1) || (134 >= uint(v7))) || v1) {
                                        var v29 int32 = -23
                                        v30 := GenS2{F0: 595}
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v28 = v28
                                        var v31 float64 = 1.5
                                        _ = v31
                                    }
                                    var v32 uint16 = 42
                                    _ = len(v3)
                                    var v33 int8 = (^(v10))
                                    _ = fi27
                                    _ = v28
                                    _ = v32
                                    _ = v33
                                }
                                var v34 float64 = (3.14159 / (-2.25 * float64(v24)))
                                _ = v34
                            } else {
                                var v35 uint = 204
                                var v36 uint8 = 255
                                v36 += 0
                                _ = v35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 12; fi37++ {
                                v24 = (float32((((0.1 * float64(v24)) + (1e10 * 0.0)) * 0.1)) / (v24 / 3.14159))
                                if (true || v1) {
                                    var v38 int = fi25
                                    var v39 uint32 = 2
                                    var v40 bool = v1
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int16 = 42
                                    v42 := GenS3{F0: 553, F1: v10, F2: 46}
                                    if ((-9223372036854775808 != 42) || ((^(fi25)) != (fi25 & (fi37 + fi25)))) {
                                        v43 := v3[v7]
                                        v41 += v41
                                        _ = v43
                                    } else {
                                        v7 = 306
                                    }
                                    _ = v41
                                    _ = v42
                                }
                                _ = fi37
                            }
                        }
                        if (false && v1) {
                            switch "" {
                            case "a":
                                var v44 int32 = -2147483648
                                _ = v44
                            case "foo bar":
                                v45 := GenS2{F0: 1}
                                _ = v45
                            }
                            v46 := GenS1{F0: v10, F1: v10}
                            for fi47 := 0; fi47 < 9; fi47++ {
                                v1 = v1
                                v48 := GenS1{F0: v10, F1: v10}
                                var v49 uint8 = uint8(v7)
                                switch (fi47 & 3) {
                                case 66:
                                    var v50 int = fi47
                                    _ = v50
                                case 1:
                                    v49 = (v49 << ((567) & 63))
                                case 9223372036854775807:
                                    v24 = v24
                                default:
                                    switch (("hello" + "foo bar") + ((("\n\t" + "foo bar") + ("foo bar" + ((("\n\t" + "foo bar") + "") + ("foo bar" + "foo bar")))) + ("" + "\n\t"))) {
                                    case "\n\t":
                                        var v51 uint64 = 2
                                        _ = v51
                                    case "a":
                                        v49 -= (v49 / ((12) | 1))
                                    }
                                }
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            _ = v46
                        } else {
                            v7 += 39
                        }
                    }
                    fmt.Println(v24)
                    _ = fi25
                }
                _ = v24
            }
            var v52 string = "\n\t"
            var v53 uint = 524
            _ = v52
            _ = v53
        }
        switch ("foo bar" + "\n\t") {
        case "":
            _ = len(v3)
        case "hello":
            var v54 uint8 = 128
            _ = v54
        case "a":
            v55 := GenS4{F0: 21}
            _ = v55
        }
        v1 = (-242 != 32767)
    } else {
        switch "\n\t" {
        case "a":
            v3[(v7 - (v7 | v7))] = 9223372036854775807
        case "foo bar":
            v1 = ((!(-468 >= -787)) && v1)
        }
        var v56 float32 = (3.14159 + float32(3.14159))
        v10 = v10
        _ = v56
    }
    if (true && (!(4294967295 < -1))) {
        v57 := GenS5{F0: v1, F1: -32768, F2: 18446744073709551615}
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
