package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 int32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int32; F1 int8; F2 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint) int8 {
    v4 := []int{42, 128, 127, 32768}
    fmt.Println(p0)
    p0 = 1
    _ = len(v4)
    var v5 uint64 = 42
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 string = "a"
    v2 := []string{"a"}
    var v3 uint = 255
    v3 += v3
    if (!(1 <= 0)) {
        switch "" {
        case "foo bar":
            _ = len(v2)
        case "a":
            switch v1 {
            case "foo bar":
                var v6 uint16 = 42
                _ = v6
            }
        case "\n\t":
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(fi7, fi7, v1)
                v8 := GenS1{F0: v1}
                v9 := GenS1{F0: v1}
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            v10 := GenS2{F0: genFn0(0), F1: -297, F2: 255}
            _ = v10
        }
        v11 := GenS3{F0: 127, F1: 0}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v11
    }
    if (!false) {
        fmt.Println(v3)
        v1 = v1
    }
    v12 := GenS2{F0: 100, F1: 1000, F2: 180}
    if (1 == 65535) {
        for fi13 := 0; fi13 < 15; fi13++ {
            v1 = v1
            if (32767 > -32768) {
                v1 = v1
                v14 := GenS4{F0: 255}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 11; fi15++ {
                    switch 505 {
                    case 638:
                        v3 += 354
                        v16 := GenS3{F0: 2, F1: 513}
                        v1 = (v1 + "\n\t")
                        _ = v16
                    case 42:
                        var v17 int = fi13
                        _ = v17
                    case 127:
                        v3 += v3
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if (-128 != 127) {
                        if (0 > (v3 - (v3 * 2))) {
                            var v18 float64 = (float64((3.14159 + 1e10)) + (0.1 * 0.1))
                            switch (fi15 * (fi15 >> (((v3 & (v3 % (((^(v3))) | 1)))) & 63))) {
                            case -431:
                                v1 = ((v1 + "\n\t") + v1)
                            default:
                                if ((-97 < -93) || (true && true)) {
                                    if true {
                                        fmt.Println(v3)
                                    }
                                    if (true || (!false)) {
                                        v3 = v3
                                        switch v1 {
                                        case "hello":
                                            v18 += v18
                                        case "foo bar":
                                            var v19 uint16 = 938
                                            _ = v19
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        }
                                    } else {
                                        var v20 int = 9223372036854775807
                                        fmt.Println(v1, v1, fi15)
                                        _ = v20
                                    }
                                    v21 := GenS5{F0: 1, F1: -190, F2: -223}
                                    v22 := GenS3{F0: 1, F1: 98}
                                    _ = v21
                                    _ = v22
                                } else {
                                    if false {
                                        v23 := GenS6{F0: 521, F1: "foo bar"}
                                        fmt.Println(v3)
                                        v24 := GenS7{F0: -1, F1: genFn0((v3 | v3)), F2: 18446744073709551615}
                                        _ = v23
                                        _ = v24
                                    }
                                }
                            }
                            _ = v18
                        }
                        fmt.Println(v1, v3)
                        var v25 int16 = 32767
                        _ = v25
                    } else {
                        var v26 float32 = 0.1
                        v27 := GenS1{F0: v1}
                        switch (fi15 & fi13) {
                        case -208:
                            fmt.Println(v1, v26)
                        case -9223372036854775808:
                            if false {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v26)
                                default:
                                    v3 = (v3 % ((v3) | 1))
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v28 := GenS7{F0: -87, F1: 42, F2: 0}
                                fmt.Println(fi13, v1, v26)
                                _ = v28
                            } else {
                                var v29 uint64 = 18446744073709551615
                                _ = v29
                            }
                        }
                        _ = v26
                        _ = v27
                    }
                    _ = fi15
                }
                v1 = v1
                v3 = v3
            }
            var v30 int32 = 959
            _ = fi13
            _ = v30
        }
        v31 := GenS8{F0: 0}
        _ = v31
    } else {
        _ = len(v2)
        v3 = (v3 * v3)
    }
    if true {
        var v32 int = 0
        _ = v32
    }
    var v33 uint32 = 935
    if (!false) {
        v33 = (v33 >> ((v3) & 63))
        var v34 string = v1
        v3 = v3
        _ = v34
    } else {
        var v35 uint32 = v33
        v36 := GenS9{F0: 1000, F1: (^(v3))}
        for fi37 := 0; fi37 < 7; fi37++ {
            if len(v2) > 0 {
                v2[0] = (v1 + v1)
            }
            switch v1 {
            case "a":
                if (!(32767 != 0)) {
                    fmt.Println(v1)
                    if (1 >= -127) {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            v39 := GenS10{F0: 127}
                            v3 = 901
                            _ = fi38
                            _ = v39
                        }
                        if len(v2) > 0 {
                            v2[0] = v1
                        }
                        v40 := GenS5{F0: -1, F1: 32767, F2: 42}
                        _ = v40
                    }
                    var v41 int64 = 42
                    _ = v41
                } else {
                    _ = len(v2)
                    v33 = 871
                    v42 := GenS2{F0: -128, F1: 236, F2: 441}
                    _ = v42
                }
            default:
                v33 = (v35 << (((^(v3))) & 63))
            }
            _ = fi37
        }
        _ = v35
        _ = v36
    }
    var v43 uint = (v3 << ((v3) & 63))
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 bool = true
        switch fi44 {
        case 221:
            fmt.Println(v3, v43)
        case -9223372036854775808:
            var v46 uint8 = 128
            _ = v46
        default:
            v47 := GenS11{F0: 2}
            _ = v47
        }
        for fi48 := 0; fi48 < 7; fi48++ {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + v1)
            }
            v49 := GenS6{F0: 0, F1: ("hello" + v1)}
            _ = fi48
            _ = v49
        }
        _ = fi44
        _ = v45
    }
    var v50 uint64 = 18446744073709551615
    if (67 < 127) {
        if true {
            var v51 int = 32768
            var v52 float32 = -2.25
            _ = v51
            _ = v52
        }
        switch v1 {
        case "\n\t":
            if true {
                v50 += v50
                v43 = 65535
            }
        case "":
            if (!(true || (true || true))) {
                v53 := GenS7{F0: -746, F1: 50, F2: (v50 & (v50 & (v50 + (v50 / (((^(v50))) | 1)))))}
                switch v1 {
                case "hello":
                    v54 := GenS8{F0: 100}
                    _ = v54
                case "foo bar":
                    _ = len(v2)
                default:
                    for fi55 := 0; fi55 < 16; fi55++ {
                        v43 += v3
                        _ = fi55
                    }
                }
                _ = v53
            }
        case "hello":
            v1 = (v1 + ((("\n\t" + (v1 + "\n\t")) + v1) + v1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v43:", v43)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v43
    _ = v50
}
