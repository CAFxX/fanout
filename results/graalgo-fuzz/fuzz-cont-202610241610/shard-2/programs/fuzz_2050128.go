package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 bool }

func main() {
    v1 := []uint32{2, 1, 569}
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := []int{1, -87, 2147483647, 127}
    var v5 int64 = -9223372036854775808
    var v6 uint = 205
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                switch fi7 {
                case -2147483648:
                    var v9 int32 = 1
                    _ = v9
                case 32768:
                    v10 := GenS1{F0: 38, F1: (("foo bar" + "a") + "\n\t"), F2: 42}
                    _ = v10
                case -9223372036854775808:
                    v5 += v5
                }
                _ = fi8
            }
            fmt.Println(v6, fi7)
        } else {
            v11 := v2[fi7]
            switch (fi7 >> (((^(v6))) & 63)) {
            case 1:
                if false {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v5 = v5
                        _ = fi12
                    }
                    v13 := GenS1{F0: -128, F1: ("\n\t" + (v11 + v11)), F2: 2147483647}
                    v11 = (v11 + (v11 + "a"))
                    _ = v13
                }
            case -68:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS2{F0: false, F1: -1, F2: 0}
                    v5 -= v5
                    v5 = (v5 % ((v5) | 1))
                    _ = fi14
                    _ = v15
                }
            case 225:
                var v16 float32 = float32((0.0 + float64(3.14159)))
                _ = v16
            default:
                v5 = 9223372036854775807
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 bool = false
                switch (fi17 >> (((v6 - (v6 | (v6 + (v6 + 0))))) & 63)) {
                case 127:
                    if (v18 || (241 < 42)) {
                        v4 = append(v4, (fi7 % (((fi17 + (^(fi17)))) | 1)))
                        var v19 float32 = (0.0 + 3.14159)
                        switch (fi7 % (((fi17 | fi17)) | 1)) {
                        case 9223372036854775807:
                            fmt.Println(v11, v6)
                        case -9223372036854775808:
                            fmt.Println(v11)
                            _ = len(v2)
                            v11 = "\n\t"
                            var v20 bool = v18
                            if (-1 <= 100) {
                                v21 := GenS1{F0: -9, F1: v11, F2: 0}
                                _ = v21
                            }
                            _ = v20
                        default:
                            if false {
                                if v18 {
                                    fmt.Println(v11, v5, v18)
                                    delete(v2, fi7)
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v22 := v2[fi17]
                                    switch (fi7 >> ((872) & 63)) {
                                    case 32768:
                                        if true {
                                            v19 = 3.14159
                                        } else {
                                            var v23 uint64 = 18446744073709551615
                                            _ = v23
                                        }
                                    case 2147483647:
                                        var v24 uint16 = 20
                                        _ = v24
                                    case 128:
                                        v5 = (^(v5))
                                    }
                                    v2[9223372036854775807] = (v22 + "")
                                    if false {
                                        var v25 string = "a"
                                        var v26 int8 = 34
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 float32 = 3.14159
                                        _ = v27
                                    }
                                    _ = v22
                                }
                            }
                        }
                        _ = v19
                    } else {
                        var v28 bool = (!((("\n\t" + v11) + (v11 + v11)) != "hello"))
                        var v29 int32 = 155
                        _ = v28
                        _ = v29
                    }
                case 978:
                    var v30 uint16 = 1
                    _ = v30
                case 862:
                    v31 := v3[2]
                    _ = v31
                default:
                    v32 := GenS3{F0: 127}
                    _ = v32
                }
                v33 := GenS1{F0: 113, F1: (v11 + "hello"), F2: -232}
                _ = fi17
                _ = v18
                _ = v33
            }
            _ = v11
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            v6 = (v6 << ((v6) & 63))
            v5 = v5
            v6 -= (v6 | (v6 / ((v6) | 1)))
            _ = fi34
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(v6)
            _ = fi35
        }
        _ = fi7
    }
    v3[378] = -2147483648
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 uint8 = 128
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 uint16 = 32767
            if (!true) {
                if (704 > -32768) {
                    var v40 int16 = 1000
                    if true {
                        var v41 uint8 = (v37 ^ 127)
                        var v42 int32 = 42
                        _ = v41
                        _ = v42
                    } else {
                        var v43 int16 = (v40 % (((v40 >> ((4294967295) & 63))) | 1))
                        _ = v43
                    }
                    _ = v40
                } else {
                    var v44 bool = (!(2147483648 == 1))
                    _ = v44
                }
                if ((fi38 >> ((v6) & 63)) < fi36) {
                    if true {
                        if false {
                            if false {
                                if (!(((fi38 * 32767) == fi36) || false)) {
                                    switch fi36 {
                                    case 127:
                                        v45 := GenS4{F0: (fi36 | (fi36 - fi36)), F1: 1, F2: false}
                                        _ = v45
                                    }
                                    if ((2 <= 412) || (!("hello" >= ("foo bar" + "")))) {
                                        if len(v1) > 0 {
                                            v1[0] = 334
                                        }
                                    } else {
                                        v39 = 31
                                    }
                                    v5 -= (v5 % ((v5) | 1))
                                    var v46 uint = v6
                                    v39 = 255
                                    _ = v46
                                }
                            } else {
                                v5 -= (-(v5))
                                v37 -= v37
                                if true {
                                    var v47 uint32 = uint32(v6)
                                    fmt.Println(v5)
                                    switch "hello" {
                                    case "":
                                        v5 = int64(v39)
                                    case "\n\t":
                                        v37 = v37
                                    }
                                    _ = v47
                                }
                            }
                        } else {
                            delete(v3, 65535)
                        }
                    } else {
                        v37 += v37
                        v6 += 686
                        switch "hello" {
                        case "":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v4 = append(v4, (fi38 - 1))
                            v6 -= 42
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            if (false || false) {
                                fmt.Println(v37)
                                v5 += (v5 << ((0) & 63))
                            } else {
                                var v48 int = (fi38 / ((2147483647) | 1))
                                v5 -= v5
                                _ = v48
                            }
                        default:
                            var v49 int64 = 516
                            _ = v49
                        }
                    }
                }
                v37 = 128
            } else {
                v50 := GenS2{F0: (-23 > 0), F1: 100, F2: uint32(v39)}
                fmt.Println(fi36, fi38)
                v51 := GenS5{F0: true}
                v1 = append(v1, 2)
                v6 = (^(v6))
                _ = v50
                _ = v51
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi38
            _ = v39
        }
        if (!true) {
            fmt.Println(fi36)
        } else {
            for fi52 := 0; fi52 < 6; fi52++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi52
            }
            switch fi36 {
            case 9223372036854775807:
                fmt.Println(fi36, fi36, fi36)
            case -1:
                v5 += (v5 / ((-88) | 1))
            case -926:
                v5 = (v5 / ((-1) | 1))
            default:
                if ((fi36 / (((fi36 / ((int(v5)) | 1))) | 1)) == 42) {
                    var v53 int32 = 659
                    if true {
                        var v54 uint64 = 1
                        _ = v54
                    }
                    _ = v53
                }
            }
            var v55 uint = v6
            _ = v55
        }
        _ = fi36
        _ = v37
    }
    v5 = (v5 + (v5 - (-(v5))))
    var v56 uint32 = 277
    v57 := GenS2{F0: ((v6 & (v6 + uint(v56))) <= (v6 >> (((^(v6))) & 63))), F1: 127, F2: 2147483648}
    var v58 uint64 = 806
    v59 := GenS2{F0: (!false), F1: 1, F2: 42}
    var v60 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
