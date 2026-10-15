package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
func genFn0() uint32 {
    v7 := GenS1{F0: false, F1: -128}
    var v8 string = ("foo bar" + "a")
    v8 = ""
    v9 := []int{-607, 1}
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint16 {
    var v10 uint8 = 54
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(fi11, fi11)
        var v12 uint8 = v10
        _ = fi11
        _ = v12
    }
    _ = v10
    return 271
}

func main() {
    v1 := []uint32{2, 883}
    var v2 float32 = 1.5
    var v3 float32 = (1e10 + (1.5 / (3.14159 + (float32(-2.25) - (3.14159 * v2)))))
    var v4 uint16 = 65535
    var v5 int16 = int16(v4)
    var v6 string = "a"
    if true {
        for fi13 := 0; fi13 < 20; fi13++ {
            if true {
                v2 = v3
                var v14 int64 = 387
                _ = v14
            }
            if true {
                v4 -= 283
            } else {
                v5 -= -683
            }
            _ = fi13
        }
    }
    var v15 uint16 = (^(v4))
    v4 -= v15
    if (true || (1.5 < (1.5 * (3.14159 + (0.1 + 0.0))))) {
        v15 += genFn1()
    } else {
        fmt.Println(v3)
        if false {
            switch 978 {
            case 128:
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            case 42:
                v5 -= 0
            default:
                fmt.Println(v4, v2, v4)
            }
            if true {
                v1 = append(v1, 65535)
                switch ("foo bar" + v6) {
                case "hello":
                    v15 = 226
                case "\n\t":
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v17 := GenS1{F0: (!(!(1 != 196))), F1: -1}
                        _ = fi16
                        _ = v17
                    }
                case "a":
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v19 := GenS1{F0: true, F1: 42}
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v20 int8 = -74
                    _ = v20
                }
            }
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            fmt.Println(v2, v2, v15)
            var v22 string = v6
            _ = fi21
            _ = v22
        }
    }
    v6 = v6
    if ((0.1 + float64(1e10)) > float64(-2.25)) {
        fmt.Println(v3, v15)
        v6 = ("foo bar" + (v6 + v6))
    } else {
        v15 -= genFn1()
        for fi23 := 0; fi23 < 1; fi23++ {
            if true {
                var v24 string = "hello"
                v3 = ((v3 - v3) + (1e10 + v2))
                _ = v24
            } else {
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 float64 = 1e10
                    if (-1 > 2147483647) {
                        v5 -= v5
                        var v27 int16 = (^(v5))
                        if false {
                            fmt.Println(v15)
                            _ = len(v1)
                            switch v6 {
                            case "a":
                                fmt.Println(v4)
                            case "\n\t":
                                if (76 >= 87) {
                                    if false {
                                        v4 -= (v4 / ((v15) | 1))
                                    }
                                    v28 := GenS1{F0: (int(v27) >= fi25), F1: 42}
                                    _ = v28
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if true {
                                    var v29 uint = 65535
                                    fmt.Println(v6, v2)
                                    _ = v29
                                }
                            default:
                                v30 := GenS2{F0: -1, F1: 63, F2: 2}
                                _ = v30
                            }
                        }
                        var v31 string = v6
                        v15 = 677
                        _ = v27
                        _ = v31
                    }
                    var v32 uint16 = (v4 << ((42) & 63))
                    _ = fi25
                    _ = v26
                    _ = v32
                }
                for fi33 := 0; fi33 < 14; fi33++ {
                    if ((!(4294967295 == 628)) && (((^(fi23)) > int(v15)) || false)) {
                        switch v6 {
                        case "hello":
                            var v34 uint8 = 37
                            _ = v34
                        default:
                            fmt.Println(fi33, v6, v3)
                        }
                        if false {
                            if true {
                                var v35 uint8 = 0
                                v35 = v35
                                _ = v35
                            }
                            var v36 uint16 = genFn1()
                            _ = v36
                        }
                        v37 := GenS3{F0: 42, F1: 724}
                        if len(v1) > 0 {
                            v1[0] = 143
                        }
                        var v38 bool = true
                        if false {
                            var v39 int64 = 0
                            var v40 int = fi23
                            v38 = ((false && true) || ((v5 | (v5 & v5)) <= (v5 & -1)))
                            v5 += v5
                            _ = v39
                            _ = v40
                        } else {
                            var v41 uint64 = 18446744073709551615
                            switch v6 {
                            case "":
                                fmt.Println(v38, v38)
                            }
                            if (!(-977 > -1)) {
                                switch v6 {
                                case "a":
                                    var v42 int8 = 126
                                    _ = v42
                                case "\n\t":
                                    v5 = v5
                                }
                                if (-2.25 > 1.5) {
                                    v5 -= (v5 | -1)
                                    v3 -= (v3 * v3)
                                    fmt.Println(v3)
                                } else {
                                    v1 = append(v1, 65535)
                                    v43 := GenS2{F0: -816, F1: uint8(v15), F2: 2147483647}
                                    _ = v43
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) > 0 {
                                    v1[0] = 65535
                                }
                            } else {
                                if ((v3 - (3.14159 * v3)) != float32((float64(v2) + 0.0))) {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                } else {
                                    var v44 uint = 4294967295
                                    _ = v44
                                }
                                v2 += v2
                                switch (fi23 * (fi33 % ((fi33) | 1))) {
                                case -740:
                                    v2 = (v3 - (v3 / v3))
                                case 32768:
                                    if len(v1) > 0 {
                                        v1[0] = 4294967295
                                    }
                                case 9223372036854775807:
                                    var v45 uint8 = 128
                                    _ = v45
                                }
                            }
                            _ = v41
                        }
                        var v46 bool = false
                        _ = v37
                        _ = v38
                        _ = v46
                    }
                    v1 = append(v1, 4294967295)
                    _ = fi33
                }
            }
            _ = fi23
        }
    }
    var v47 int64 = -14
    var v48 int64 = v47
    v3 += (v3 / ((((0.0 + -2.25) - ((float32((-2.25 / 1.0)) / v3) / (v3 * ((v2 - v3) - (v2 + 0.0))))) / ((v3 * (0.0 / (0.1 - 1.5))) / v3)) - ((v2 + (v3 / float32(1.5))) * (v2 + float32((0.1 / 3.14159))))))
    v1 = append(v1, 2147483647)
    if (232 < 128) {
        var v49 int32 = 42
        _ = v49
    } else {
        if true {
            var v50 int32 = int32(v5)
            v51 := GenS1{F0: (!(34 >= 0)), F1: 100}
            _ = v50
            _ = v51
        } else {
            for fi52 := 0; fi52 < 17; fi52++ {
                switch fi52 {
                case -538:
                    for fi53 := 0; fi53 < 19; fi53++ {
                        var v54 int8 = -1
                        _ = fi53
                        _ = v54
                    }
                default:
                    v55 := GenS2{F0: -554, F1: 255, F2: genFn0()}
                    _ = v55
                }
                v2 = v3
                _ = fi52
            }
            v48 = 42
        }
        for fi56 := 0; fi56 < 8; fi56++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi56
        }
    }
    _ = len(v1)
    var v57 bool = true
    v1 = append(v1, 423)
    if false {
        var v58 uint32 = genFn0()
        if (1 > (v47 % ((v48) | 1))) {
            v59 := GenS3{F0: 65535, F1: 0}
            v60 := GenS4{F0: 101, F1: (v4 << ((695) & 63)), F2: -385}
            v1 = append(v1, v58)
            _ = v59
            _ = v60
        }
        v3 += (((v3 + v3) * float32(1e10)) + 0.1)
        _ = v58
    } else {
        v4 = (v4 >> ((1) & 63))
        _ = len(v1)
        v61 := GenS4{F0: 4294967295, F1: v4, F2: -2147483648}
        _ = v61
    }
    var v62 uint32 = 723
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v57
    _ = v6
    _ = v62
}
