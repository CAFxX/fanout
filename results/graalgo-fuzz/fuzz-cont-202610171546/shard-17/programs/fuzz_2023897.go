package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int64; F1 string; F2 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32; F1 uint; F2 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 string; F2 uint64 }
type GenS12 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS13 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 uint8) int64 {
    if ((902 <= 332) || true) {
        for fi5 := 0; fi5 < 4; fi5++ {
            if (65535 != 58) {
                p0 += (p0 % ((p0) | 1))
            }
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint = 18446744073709551615
                switch (^(fi5)) {
                case 1:
                    v8 := []int{fi6, fi5, fi6, (fi6 + fi6)}
                    _ = v8
                case 128:
                    v9 := []int{(fi6 >> ((v7) & 63)), 554, (-(fi5)), fi6}
                    _ = v9
                }
                p1 = 109
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    } else {
        v10 := GenS1{F0: p1, F1: (p0 >> ((18446744073709551615) & 63)), F2: p0}
        _ = v10
    }
    p0 = p0
    var v11 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 450
    var v2 uint16 = 255
    var v3 int8 = 127
    var v4 uint = 2
    v12 := GenS2{F0: 985}
    var v13 int8 = (v3 % ((v3) | 1))
    var v14 uint = v4
    switch 735 {
    case 42:
        var v15 string = ("" + "")
        _ = v15
    case -818:
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := []int{(-(fi16)), 404}
            fmt.Println(v3, v13, v14)
            _ = fi16
            _ = v17
        }
    default:
        v18 := GenS3{F0: 151, F1: -1}
        _ = v18
    }
    var v19 int = -9223372036854775808
    v20 := []int{v19, (v19 * v19), (v19 ^ -742), v19}
    var v21 float64 = 0.1
    fmt.Println(v19, v3, v3)
    fmt.Println(v4, v3)
    fmt.Println(v19, v21)
    v22 := GenS4{F0: -9223372036854775808, F1: (v21 <= (0.0 / v21))}
    for fi23 := 0; fi23 < 18; fi23++ {
        fmt.Println(v21, v13, v2)
        _ = fi23
    }
    var v24 string = "a"
    v20 = append(v20, (v19 * v19))
    for fi25 := 0; fi25 < 14; fi25++ {
        if true {
            v14 -= 18446744073709551615
            var v26 uint8 = uint8(v2)
            v2 += v2
            _ = v26
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint64 = 4294967295
                if false {
                    v29 := GenS5{F0: 32767, F1: -32768}
                    fmt.Println(v4, v4, v4)
                    fmt.Println(v19, fi27)
                    _ = v29
                }
                fmt.Println(v13, fi27)
                _ = fi27
                _ = v28
            }
        }
        fmt.Println(v3, v13)
        for fi30 := 0; fi30 < 11; fi30++ {
            if true {
                var v31 int64 = 42
                _ = v31
            }
            var v32 string = ("\n\t" + "hello")
            if true {
                v33 := GenS6{F0: 42, F1: fi25}
                if (((1e10 * 0.1) > (((-2.25 * (0.1 / -2.25)) - (1e10 + 0.0)) * 0.1)) || false) {
                    if (1 != (v1 << (((v14 - 1)) & 63))) {
                        if (!((v21 >= (v21 - float64(-2.25))) && (v32 == v32))) {
                            var v34 uint8 = 128
                            fmt.Println(fi25)
                            v13 = (v13 - v13)
                            _ = v34
                        }
                        v14 += v4
                        v35 := GenS7{F0: 608, F1: v24, F2: 0}
                        var v36 uint = 583
                        _ = v35
                        _ = v36
                    }
                    v13 -= v13
                    if (v3 != (v3 - int8(fi30))) {
                        switch (v24 + v24) {
                        case "foo bar":
                            v37 := GenS2{F0: 1}
                            _ = v37
                        case "a":
                            if true {
                                v2 -= 42
                                if len(v20) > 0 {
                                    v20[0] = int(v2)
                                }
                                fmt.Println(v24)
                            } else {
                                var v38 uint8 = uint8(v2)
                                if (false && true) {
                                    var v39 float32 = 1.5
                                    v20 = append(v20, (fi30 * fi30))
                                    if false {
                                        switch fi30 {
                                        case 281:
                                            fmt.Println(v32, v1, fi25)
                                        case 32767:
                                            v24 = v32
                                        case -862:
                                            var v40 int64 = 1
                                            _ = v40
                                        }
                                        fmt.Println(v2, fi30)
                                    }
                                    v41 := GenS2{F0: 0}
                                    v42 := GenS8{F0: (v1 % ((65535) | 1))}
                                    _ = v39
                                    _ = v41
                                    _ = v42
                                }
                                _ = v38
                            }
                        }
                    } else {
                        if (v21 >= 3.14159) {
                            v43 := GenS9{F0: 42, F1: (v4 & v4), F2: v24}
                            v44 := GenS10{F0: -432}
                            if (1 >= 1) {
                                v13 += (^(v13))
                                fmt.Println(v14, v2, v2)
                                if (true || ((-32768 > 369) && false)) {
                                    v1 -= (v1 >> ((v14) & 63))
                                }
                                if (false && (v24 > "\n\t")) {
                                    v19 = -507
                                    var v45 int64 = 42
                                    switch (v24 + v24) {
                                    case "hello":
                                        if true {
                                            switch "" {
                                            case "":
                                                if (0 != 65535) {
                                                    switch fi25 {
                                                    case 127:
                                                        if false {
                                                            if len(v20) > 0 {
                                                                v20[0] = (v19 >> ((v14) & 63))
                                                            }
                                                        } else {
                                                            v46 := GenS6{F0: uint8(v2), F1: int(v2)}
                                                            if (true && true) {
                                                                if true {
                                                                    v47 := GenS11{F0: v24, F1: ("foo bar" + "hello"), F2: 2}
                                                                    if true {
                                                                        v20 = append(v20, (fi25 % (((fi30 + fi25)) | 1)))
                                                                        v48 := GenS4{F0: genFn0(42, 255), F1: true}
                                                                        v2 = v2
                                                                        switch v32 {
                                                                        case "":
                                                                            v49 := GenS6{F0: 225, F1: (fi30 / ((int(v2)) | 1))}
                                                                            _ = v49
                                                                        default:
                                                                            v50 := GenS12{F0: 18446744073709551615, F1: 1, F2: genFn0(0, 50)}
                                                                            _ = v50
                                                                        }
                                                                        switch v32 {
                                                                        case "a":
                                                                            if false {
                                                                                var v51 int64 = v45
                                                                                _ = v51
                                                                            } else {
                                                                                var v52 uint = (^(v4))
                                                                                _ = len(v20)
                                                                                _ = v52
                                                                            }
                                                                        case "foo bar":
                                                                            var v53 bool = ((v2 | v2) > (^(v2)))
                                                                            _ = v53
                                                                        default:
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                switch v32 {
                                                                                case "\n\t":
                                                                                    var v54 uint16 = v2
                                                                                    _ = v54
                                                                                case "":
                                                                                    v55 := GenS3{F0: 999, F1: 2147483647}
                                                                                    _ = v55
                                                                                }
                                                                                var v56 int64 = v45
                                                                                _ = v56
                                                                            case "hello":
                                                                                v1 -= v1
                                                                            default:
                                                                                var v57 uint16 = (v2 >> ((v14) & 63))
                                                                                _ = v57
                                                                            }
                                                                        }
                                                                        v3 += v13
                                                                        v20 = append(v20, 42)
                                                                        _ = v48
                                                                    }
                                                                    _ = v47
                                                                }
                                                            } else {
                                                                if len(v20) >= 2 {
                                                                    v20 = v20[1:len(v20)]
                                                                }
                                                                v58 := GenS4{F0: int64(v3), F1: ((0 <= 127) && false)}
                                                                _ = v58
                                                            }
                                                            if len(v20) > 0 {
                                                                v20[0] = (v19 << ((v14) & 63))
                                                            }
                                                            _ = v46
                                                        }
                                                    case -65:
                                                        var v59 uint8 = 255
                                                        _ = v59
                                                    case 42:
                                                        v20 = append(v20, v19)
                                                    }
                                                    if false {
                                                        v32 = v32
                                                        var v60 int8 = int8(v45)
                                                        var v61 float64 = v21
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                } else {
                                                    switch v32 {
                                                    case "a":
                                                        var v62 string = v24
                                                        _ = v62
                                                    }
                                                    v4 = v14
                                                    v63 := GenS2{F0: 32767}
                                                    _ = v63
                                                }
                                            }
                                            if true {
                                                switch v24 {
                                                case "a":
                                                    var v64 string = (v32 + "a")
                                                    _ = v64
                                                case "\n\t":
                                                    var v65 int8 = (-(v3))
                                                    _ = v65
                                                case "hello":
                                                    if true {
                                                        v45 = v45
                                                        v2 = v2
                                                        var v66 int8 = (v3 % ((100) | 1))
                                                        _ = v66
                                                    }
                                                }
                                            } else {
                                                var v67 uint = 61
                                                _ = v67
                                            }
                                            if (!(!(!(!true)))) {
                                                fmt.Println(v1)
                                                var v68 string = (v32 + v24)
                                                _ = v68
                                            } else {
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            v69 := GenS9{F0: int32(v13), F1: 2, F2: "foo bar"}
                                            _ = v69
                                        }
                                    case "a":
                                        v70 := GenS13{F0: (v45 & v45), F1: (v4 & 255)}
                                        _ = v70
                                    case "":
                                        v21 = v21
                                    default:
                                        v71 := GenS4{F0: v45, F1: (1 >= 255)}
                                        _ = v71
                                    }
                                    _ = v45
                                }
                                switch (fi30 % ((697) | 1)) {
                                case 0:
                                    v32 = v24
                                case 32767:
                                    fmt.Println(fi25, v24)
                                default:
                                    var v72 uint16 = (v2 / ((uint16(v14)) | 1))
                                    _ = v72
                                }
                            }
                            _ = v43
                            _ = v44
                        }
                    }
                }
                _ = v33
            }
            _ = fi30
            _ = v32
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
