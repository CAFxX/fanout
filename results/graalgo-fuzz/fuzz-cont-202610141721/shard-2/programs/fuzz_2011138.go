package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint8 = 0
    var v2 int64 = -457
    v3 := []int{32767, 128, 9223372036854775807}
    v4 := make(map[string]int)
    var v5 uint16 = 433
    var v6 string = "\n\t"
    v1 = 128
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            _ = len(v4)
            var v9 uint32 = 42
            _ = fi8
            _ = v9
        }
        v5 = v5
        _ = fi7
    }
    if (2147483647 == 32768) {
        if (-957 <= -2147483648) {
            fmt.Println(v5, v1, v5)
        } else {
            var v10 uint16 = 32767
            if true {
                switch -47 {
                case 128:
                    var v11 int64 = 122
                    _ = v11
                case 744:
                    v12 := GenS1{F0: uint64(v5), F1: -1}
                    _ = v12
                case 2147483647:
                    v13 := GenS1{F0: 820, F1: -121}
                    _ = v13
                default:
                    var v14 int32 = 1
                    _ = v14
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 uint32 = 42
                    _ = fi15
                    _ = v16
                }
            }
            v2 = (-(v2))
            _ = v10
        }
        delete(v4, "foo bar")
        v3 = append(v3, int(v5))
    }
    var v17 float32 = 0.0
    if true {
        var v18 uint64 = uint64(v5)
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                v17 = v17
                v2 = (v2 / (((-(v2))) | 1))
                _ = fi20
            }
            v21 := GenS1{F0: 466, F1: -128}
            _ = fi19
            _ = v21
        }
        _ = v18
    } else {
        v1 -= v1
        var v22 uint32 = 4294967295
        fmt.Println(v6, v5, v2)
        _ = v22
    }
    fmt.Println(v1)
    var v23 int = -2147483648
    v5 -= v5
    var v24 int32 = 2147483647
    fmt.Println(v6)
    switch ("\n\t" + "foo bar") {
    case "":
        v17 = 3.14159
    case "a":
        fmt.Println(v23, v2, v1)
    case "\n\t":
        for fi25 := 0; fi25 < 4; fi25++ {
            v6 = v6
            for fi26 := 0; fi26 < 18; fi26++ {
                v4[""] = 1
                _ = fi26
            }
            _ = fi25
        }
    default:
        if (65535 >= 255) {
            v27 := GenS2{F0: 42, F1: (0 == 65535), F2: (v24 == int32(v1))}
            fmt.Println(v23)
            _ = v27
        } else {
            var v28 uint = 255
            _ = v28
        }
    }
    v29 := GenS3{F0: 614, F1: 42, F2: (v2 + 0)}
    switch v6 {
    case "foo bar":
        v30 := GenS2{F0: 1, F1: ((v1 <= 127) && ((((float64((1e10 + v17)) / 1.0) - 1e10) != 3.14159) || false)), F2: ((!(false && true)) && true)}
        _ = v30
    case "a":
        var v31 int32 = 484
        _ = v31
    default:
        if (!true) {
            v24 -= (v24 % ((v24) | 1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v32 uint16 = v5
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        if (42 <= (v2 % (((v2 << ((183) & 63))) | 1))) {
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int8 = 0
                v2 -= v2
                v6 = "\n\t"
                var v36 int32 = (v24 - (v24 >> ((65535) & 63)))
                v37 := GenS3{F0: 42, F1: 1, F2: v2}
                fmt.Println(v23, v17, v2)
                var v38 int8 = v35
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
                _ = v38
            }
        } else {
            if (!(!false)) {
                var v39 uint8 = 2
                _ = v39
            } else {
                if true {
                    if (v23 < v23) {
                        v3 = append(v3, 32768)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi40 := 0; fi40 < 11; fi40++ {
                            var v41 uint32 = 2
                            _ = fi40
                            _ = v41
                        }
                    }
                    var v42 int32 = (v24 / (((v24 / ((-884) | 1))) | 1))
                    _ = v42
                } else {
                    v6 = ("\n\t" + v6)
                    v43 := v4[(("" + (v6 + v6)) + v6)]
                    var v44 uint64 = 42
                    _ = v43
                    _ = v44
                }
                v1 += v1
            }
            var v45 int64 = (v2 & v2)
            if (42 <= 4294967295) {
                fmt.Println(v6, fi33)
            } else {
                v46 := GenS4{F0: ((v17 * 1.5) < v17)}
                _ = v46
            }
            _ = v45
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
