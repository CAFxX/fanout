package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    var v3 uint = 424
    var v4 float32 = ((1.5 * (1.5 + -2.25)) * 1.5)
    var v5 int8 = -105
    fmt.Println(v5, v5)
    var v6 bool = false
    fmt.Println(v5, v3)
    switch 0 {
    case 2147483647:
        switch 899 {
        case -7:
            if ("foo bar" == "a") {
                switch 9223372036854775807 {
                case -116:
                    var v7 bool = false
                    _ = v7
                case 32767:
                    if (486 < 4294967295) {
                        if ((9223372036854775807 > 0) || (!v6)) {
                            var v8 uint8 = 127
                            for fi9 := 0; fi9 < 5; fi9++ {
                                v10 := GenS1{F0: (fi9 * (^(fi9))), F1: v8}
                                fmt.Println(v8, v3, v6)
                                _ = fi9
                                _ = v10
                            }
                            for fi11 := 0; fi11 < 17; fi11++ {
                                v8 = v8
                                if v6 {
                                    var v12 uint64 = 18446744073709551615
                                    _ = v12
                                } else {
                                    if v6 {
                                        v13 := GenS1{F0: -9223372036854775808, F1: v8}
                                        _ = v13
                                    }
                                    switch ("hello" + (((((("foo bar" + ("\n\t" + "")) + (((("\n\t" + "\n\t") + ("" + ("" + (("a" + (("foo bar" + "\n\t") + "a")) + "hello")))) + "") + "")) + "foo bar") + (("foo bar" + "a") + "hello")) + ("foo bar" + "a")) + "foo bar")) {
                                    case "foo bar":
                                        for fi14 := 0; fi14 < 9; fi14++ {
                                            switch fi14 {
                                            case -928:
                                                if (!(78 <= 2147483647)) {
                                                    var v15 float64 = 0.0
                                                    if false {
                                                        if v6 {
                                                            var v16 uint32 = 2147483648
                                                            _ = len(v2)
                                                            _ = v16
                                                        } else {
                                                            v2[(fi14 / (((fi11 | (fi14 >> (((v3 << (((v3 % (((^(v3))) | 1))) & 63))) & 63)))) | 1))] = 1
                                                            v4 = (((v4 / v4) / (v4 / (v4 / v4))) / v4)
                                                        }
                                                        v17 := GenS1{F0: fi11, F1: (v8 % (((v8 - (v8 + 2))) | 1))}
                                                        _ = v17
                                                    }
                                                    v18 := GenS1{F0: (fi14 << (((v3 / ((712) | 1))) & 63)), F1: (v8 ^ 166)}
                                                    _ = v15
                                                    _ = v18
                                                } else {
                                                    switch fi11 {
                                                    case 206:
                                                        v8 -= v8
                                                    case 954:
                                                        v19 := v2[fi11]
                                                        _ = v19
                                                    case 0:
                                                        v5 = (v5 & v5)
                                                    }
                                                    var v20 uint32 = 2147483647
                                                    _ = v20
                                                }
                                                v1 = v1
                                            case -1:
                                                v21 := v2[fi14]
                                                _ = v21
                                            default:
                                                if v6 {
                                                    v22 := GenS1{F0: int(v3), F1: (v8 << ((v3) & 63))}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if false {
                                                            fmt.Println(fi11, v1, v5)
                                                        } else {
                                                            v6 = (fi11 <= fi14)
                                                            v4 -= (1.5 + 0.1)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v23 := []int{}
                                                                _ = v23
                                                            }
                                                        }
                                                    case "a":
                                                        v24 := GenS1{F0: 32767, F1: v8}
                                                        switch ("" + "hello") {
                                                        case "a":
                                                            v5 = v5
                                                            v1 += v5
                                                        case "\n\t":
                                                            if (!(((1 == 42) && ((v3 < v3) && v6)) || (0 < 2))) {
                                                                var v25 uint32 = 2
                                                                _ = v25
                                                            } else {
                                                                if (!v6) {
                                                                    var v26 int8 = (v1 - 0)
                                                                    _ = v26
                                                                } else {
                                                                    v4 = float32(0.1)
                                                                    var v27 uint = (v3 ^ 0)
                                                                    if (!((v6 && true) || (!false))) {
                                                                        var v28 uint64 = 18446744073709551615
                                                                        v27 = 255
                                                                        _ = v28
                                                                    }
                                                                    if v6 {
                                                                        var v29 float32 = float32(0.0)
                                                                        v27 += (^(v3))
                                                                        var v30 uint8 = v8
                                                                        var v31 string = "foo bar"
                                                                        _ = v29
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                    v32 := []int{(fi11 / ((fi11) | 1)), (fi14 / (((^(fi11))) | 1)), (fi11 * fi14)}
                                                                    _ = v27
                                                                    _ = v32
                                                                }
                                                                var v33 uint8 = 245
                                                                _ = v33
                                                            }
                                                        case "hello":
                                                            v4 = v4
                                                            v3 = 488
                                                            v3 += 65535
                                                        default:
                                                            var v34 string = ""
                                                            _ = v34
                                                        }
                                                        _ = v24
                                                    default:
                                                        v35 := GenS1{F0: fi11, F1: (v8 - (^(v8)))}
                                                        _ = v35
                                                    }
                                                    v36 := []int{(fi11 - 32768), (fi14 + (fi14 * fi14))}
                                                    v5 -= (^(v5))
                                                    _ = v22
                                                    _ = v36
                                                }
                                            }
                                            delete(v2, (fi14 + (fi11 * fi11)))
                                            v37 := []int{fi11, -921, (-(fi14)), fi11}
                                            _ = fi14
                                            _ = v37
                                        }
                                    case "\n\t":
                                        fmt.Println(v1, v3)
                                    default:
                                        v38 := []int{}
                                        _ = v38
                                    }
                                }
                                _ = fi11
                            }
                            _ = v8
                        }
                        fmt.Println(v6)
                        for fi39 := 0; fi39 < 3; fi39++ {
                            v40 := GenS2{F0: 516, F1: -9223372036854775808}
                            switch -45 {
                            case -2147483648:
                                var v41 uint32 = 42
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                    } else {
                        v42 := GenS3{F0: 42, F1: 905}
                        v1 += (v1 >> (((v3 | (v3 - v3))) & 63))
                        _ = v42
                    }
                default:
                    fmt.Println(v4, v6)
                }
                v3 = (v3 - 255)
            } else {
                if false {
                    v43 := GenS2{F0: 255, F1: 1}
                    var v44 uint = v3
                    _ = v43
                    _ = v44
                } else {
                    if (127 > 95) {
                        var v45 uint = 65535
                        var v46 int32 = -1
                        _ = v45
                        _ = v46
                    } else {
                        v6 = true
                        if (852 < v3) {
                            for fi47 := 0; fi47 < 12; fi47++ {
                                switch fi47 {
                                case 1:
                                    fmt.Println(v5, fi47)
                                }
                                var v48 uint8 = 127
                                _ = fi47
                                _ = v48
                            }
                        }
                    }
                    v49 := GenS3{F0: 0, F1: 4294967295}
                    fmt.Println(v6, v4, v3)
                    _ = v49
                }
                v4 += -2.25
                v4 += (float32((0.1 / 1.0)) * v4)
            }
        }
    case -46:
        var v50 int16 = int16(v5)
        _ = v50
    default:
        v51 := GenS2{F0: 750, F1: 0}
        _ = v51
    }
    var v52 uint32 = 2147483647
    if v6 {
        v4 = 1.5
        if true {
            fmt.Println(v3)
            var v53 float32 = v4
            _ = v53
        }
        v54 := GenS1{F0: 0, F1: 127}
        _ = v54
    }
    if ((810 >= 2147483647) && v6) {
        v52 = 0
        for fi55 := 0; fi55 < 17; fi55++ {
            switch ("foo bar" + "a") {
            case "hello":
                for fi56 := 0; fi56 < 7; fi56++ {
                    if false {
                        if (false || true) {
                            fmt.Println(v3, v5, v1)
                        } else {
                            v57 := []int{(fi55 << (((v3 & (^(v3)))) & 63)), fi56, fi55, (fi56 % ((fi56) | 1))}
                            var v58 float32 = ((3.14159 + 0.0) - v4)
                            _ = v57
                            _ = v58
                        }
                        var v59 int8 = v5
                        fmt.Println(fi56)
                        _ = v59
                    }
                    if true {
                        v1 += v1
                        v60 := GenS2{F0: 595, F1: 1}
                        _ = v60
                    }
                    var v61 int16 = 32767
                    _ = fi56
                    _ = v61
                }
            case "":
                for fi62 := 0; fi62 < 3; fi62++ {
                    var v63 bool = false
                    _ = fi62
                    _ = v63
                }
            }
            fmt.Println(v5, v1, v6)
            _ = fi55
        }
    } else {
        var v64 string = "a"
        if (!((3.14159 + 0.0) < (0.1 / (0.1 + 1e10)))) {
            v65 := GenS4{F0: 32767}
            _ = v65
        } else {
            v66 := []int{1}
            for fi67 := 0; fi67 < 6; fi67++ {
                var v68 int32 = 2147483647
                _ = fi67
                _ = v68
            }
            _ = v66
        }
        v69 := GenS3{F0: -887, F1: 42}
        _ = v64
        _ = v69
    }
    if false {
        v70 := GenS3{F0: 400, F1: (v52 % ((188) | 1))}
        v5 += (^(v5))
        fmt.Println(v4, v1, v6)
        _ = v70
    }
    fmt.Println(v1, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
