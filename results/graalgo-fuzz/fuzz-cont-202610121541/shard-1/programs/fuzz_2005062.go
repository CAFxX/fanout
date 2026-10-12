package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int64; F1 bool; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0(p0 uint32) uint16 {
    v7 := []int{-923, 1, -748, 4294967295}
    p0 = (^(p0))
    p0 = p0
    _ = p0
    _ = v7
    return 849
}
func genFn1(p0 uint8) int16 {
    fmt.Println(p0)
    var v8 uint = 1
    fmt.Println(v8)
    if (60 != 42) {
        if false {
            var v9 int = 32767
            v10 := GenS1{F0: v9, F1: 42, F2: 255}
            v8 = v8
            _ = v10
            _ = v9
        }
        v11 := []int{1, -9223372036854775808}
        var v12 int8 = 42
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v8
    return 1000
}

func main() {
    v1 := []int{}
    v2 := []uint32{968}
    v3 := []uint32{}
    var v4 float64 = float64((1.5 + 1.5))
    v5 := []int{1, -2147483648, 9223372036854775807}
    var v6 string = ("" + "a")
    switch -2147483648 {
    case 32768:
        v13 := GenS1{F0: -9223372036854775808, F1: 317, F2: 127}
        _ = v13
    }
    var v14 float64 = (v4 - (v4 * (v4 + v4)))
    var v15 uint = 42
    v16 := GenS2{F0: 0, F1: 4294967295, F2: 71}
    var v17 float32 = (3.14159 - 0.0)
    v14 = (-2.25 + 1.5)
    for fi18 := 0; fi18 < 11; fi18++ {
        if len(v3) > 0 {
            v3[0] = 2
        }
        switch "a" {
        case "\n\t":
            if (!true) {
                if true {
                    fmt.Println(v15, v6)
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if len(v1) > 0 {
                            v1[0] = fi18
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi18 << ((v15) & 63))
                        }
                        v20 := GenS2{F0: -1, F1: 983, F2: 74}
                        if false {
                            v15 += v15
                            _ = len(v5)
                            switch (v6 + ("\n\t" + v6)) {
                            case "a":
                                switch fi18 {
                                case -408:
                                    if ((1000 <= 842) && (0 < 771)) {
                                        v21 := GenS3{F0: -1, F1: (fi19 != 128), F2: 0}
                                        var v22 bool = false
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            v23 := GenS3{F0: 774, F1: (!(fi18 < (fi19 % ((690) | 1)))), F2: 9223372036854775807}
                                            _ = v23
                                        case "foo bar":
                                            if false {
                                                var v24 uint32 = 2147483648
                                                _ = len(v3)
                                                if len(v3) > 0 {
                                                    v3[0] = (v24 & 863)
                                                }
                                                _ = v24
                                            }
                                        case "":
                                            if false {
                                                var v25 float64 = ((-2.25 - (v4 - v4)) + float64(v17))
                                                if ((-134 >= -434) && false) {
                                                    fmt.Println(v25, fi18)
                                                    var v26 int32 = -2147483648
                                                    if (3.14159 == (float32(v25) / ((3.14159 + v17) * (0.1 - (v17 + v17))))) {
                                                        v25 = -2.25
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v26 = (v26 >> ((0) & 63))
                                                    } else {
                                                        v27 := GenS2{F0: (v26 % ((1) | 1)), F1: 2, F2: -128}
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                } else {
                                                    if false {
                                                        fmt.Println(v14, v6)
                                                        if len(v2) > 0 {
                                                            v2[0] = 203
                                                        }
                                                    } else {
                                                        var v28 int64 = 1
                                                        var v29 uint8 = 78
                                                        v28 -= (v28 + v28)
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                v30 := GenS2{F0: 1, F1: 42, F2: -1}
                                                _ = v25
                                                _ = v30
                                            }
                                        }
                                        switch ("" + v6) {
                                        case "":
                                            fmt.Println(v17, v17)
                                        default:
                                            v31 := GenS3{F0: 0, F1: true, F2: 1}
                                            _ = v31
                                        }
                                    }
                                default:
                                    fmt.Println(fi19, v4)
                                }
                            case "\n\t":
                                var v32 uint16 = genFn0(4294967295)
                                _ = v32
                            case "":
                                var v33 float64 = float64(3.14159)
                                _ = v33
                            }
                        } else {
                            v17 += v17
                        }
                        v14 = 0.1
                        _ = fi19
                        _ = v20
                    }
                }
            } else {
                _ = len(v1)
                switch "" {
                case "":
                    var v34 int8 = -128
                    _ = v34
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    v35 := GenS3{F0: 9223372036854775807, F1: (!(0 != -740)), F2: 1}
                    _ = v35
                default:
                    var v36 uint8 = 42
                    _ = v36
                }
                v37 := GenS4{F0: int8(fi18)}
                _ = v37
            }
        case "foo bar":
            _ = len(v2)
        default:
            fmt.Println(v17)
        }
        _ = fi18
    }
    v17 += v17
    switch (v6 + (((v6 + ("hello" + ((v6 + (v6 + (v6 + v6))) + ("" + (("foo bar" + (v6 + "")) + "foo bar"))))) + (("a" + (((v6 + (("" + v6) + "")) + v6) + "a")) + "")) + v6)) {
    case "":
        if true {
            if true {
                var v38 float32 = ((((0.0 / ((v17 * float32(v14)) - v17)) * v17) * v17) / -2.25)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch -9223372036854775808 {
                case -648:
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if (32767 >= 834) {
                            v40 := GenS2{F0: 1, F1: 119, F2: -118}
                            v14 = float64(0.1)
                            v41 := GenS1{F0: fi39, F1: uint16(v15), F2: 239}
                            _ = v40
                            _ = v41
                        }
                        _ = fi39
                    }
                case 9223372036854775807:
                    v42 := GenS3{F0: -9223372036854775808, F1: false, F2: 42}
                    _ = v42
                case -1:
                    v6 = v6
                default:
                    if false {
                        v3 = append(v3, 4294967295)
                        var v43 int8 = 127
                        v44 := GenS1{F0: 32767, F1: 1000, F2: 148}
                        _ = v43
                        _ = v44
                    } else {
                        var v45 int32 = 2147483647
                        if (float64(v17) > v4) {
                            v46 := GenS2{F0: v45, F1: 674, F2: 127}
                            var v47 uint16 = 0
                            _ = v46
                            _ = v47
                        } else {
                            var v48 float64 = (0.1 * v4)
                            var v49 uint32 = 4294967295
                            _ = len(v1)
                            _ = v48
                            _ = v49
                        }
                        _ = v45
                    }
                }
                _ = v38
            } else {
                v6 = v6
                if (float64(v17) <= (v14 / (3.14159 - float64(-2.25)))) {
                    var v50 int8 = 42
                    var v51 string = v6
                    for fi52 := 0; fi52 < 3; fi52++ {
                        if ((!true) && true) {
                            if false {
                                if false {
                                    for fi53 := 0; fi53 < 7; fi53++ {
                                        switch ("\n\t" + v51) {
                                        case "":
                                            switch fi52 {
                                            case 79:
                                                switch int(v15) {
                                                case 476:
                                                    v14 += v14
                                                }
                                            case 1:
                                                v14 -= v4
                                            }
                                        case "foo bar":
                                            v1 = append(v1, (fi53 ^ fi52))
                                            v54 := GenS1{F0: 4294967295, F1: 230, F2: 128}
                                            v15 += (v15 % ((v15) | 1))
                                            _ = v54
                                        case "hello":
                                            v15 += v15
                                        }
                                        if (!false) {
                                            v55 := GenS3{F0: 9223372036854775807, F1: true, F2: 9223372036854775807}
                                            var v56 uint = 4294967295
                                            fmt.Println(fi52, v56, v56)
                                            _ = v55
                                            _ = v56
                                        } else {
                                            if (231 != 255) {
                                                v57 := GenS3{F0: 153, F1: false, F2: 1}
                                                fmt.Println(fi52)
                                                _ = v57
                                            }
                                            v58 := GenS5{F0: v50, F1: 1, F2: 444}
                                            _ = v58
                                        }
                                        _ = fi53
                                    }
                                    v14 = v14
                                }
                            }
                            _ = len(v2)
                            v17 -= 1e10
                        }
                        v6 = v6
                        var v59 int32 = -26
                        _ = fi52
                        _ = v59
                    }
                    _ = v50
                    _ = v51
                }
            }
        }
    case "a":
        v4 = 3.14159
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
