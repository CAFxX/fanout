package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 uint8, p2 int64) uint16 {
    p1 -= p1
    fmt.Println(p1)
    v6 := GenS1{F0: p0}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 string) uint32 {
    fmt.Println(p1, p0)
    var v8 int32 = -126
    fmt.Println(p1, v8)
    p1 = (p1 + p1)
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 2147483648
}

func main() {
    v1 := []uint32{2147483647, 319, 2}
    var v2 uint8 = 160
    var v3 float32 = 0.0
    var v4 uint64 = 18446744073709551615
    var v5 string = (("hello" + "foo bar") + "hello")
    v10 := GenS1{F0: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if (1 > 2147483647) {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint8 = v2
                v14 := GenS1{F0: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(fi11, v5)
            var v15 bool = (!(4294967295 <= 0))
            _ = v15
        } else {
            v1 = append(v1, 4294967295)
        }
        v3 += v3
        _ = fi11
    }
    v16 := GenS1{F0: 42}
    v1 = append(v1, 42)
    for fi17 := 0; fi17 < 17; fi17++ {
        if len(v1) > 0 {
            v1[0] = 1
        }
        var v18 uint = 663
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v5 = v5
        if true {
            v4 += (v4 / ((v4) | 1))
            v3 = 1e10
            var v20 float64 = 0.0
            _ = v20
        } else {
            var v21 string = v5
            fmt.Println(fi19, v21, v21)
            v22 := GenS1{F0: 0}
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: 0}
        _ = fi19
        _ = v23
    }
    switch v5 {
    case "":
        for fi24 := 0; fi24 < 9; fi24++ {
            if ((float64(v3) * 0.1) != 3.14159) {
                switch fi24 {
                case 128:
                    v25 := GenS1{F0: 42}
                    _ = v25
                default:
                    if (false && true) {
                        switch v5 {
                        case "hello":
                            var v26 uint16 = 715
                            _ = v26
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 388
                        }
                        if false {
                            fmt.Println(v5, fi24, v5)
                            fmt.Println(v3, v4, v3)
                            v4 -= (v4 / ((0) | 1))
                        }
                    }
                }
                fmt.Println(v5, v3, v3)
                var v27 int8 = -1
                _ = v27
            }
            if (0.0 == 1.5) {
                fmt.Println(v3)
                fmt.Println(v3)
            } else {
                if true {
                    var v28 uint = 2
                    _ = v28
                }
                var v29 int32 = 26
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                _ = v29
            }
            _ = fi24
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v4 = (v4 % ((4294967295) | 1))
        _ = fi30
    }
    var v31 uint8 = 131
    v3 -= v3
    var v32 float64 = 3.14159
    v33 := GenS1{F0: 18446744073709551615}
    var v34 int16 = -1
    v32 = 0.0
    if true {
        switch "foo bar" {
        case "foo bar":
            v4 += v4
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            v34 = v34
            v36 := GenS1{F0: 2}
            if false {
                switch "a" {
                case "\n\t":
                    v37 := GenS2{F0: 1, F1: v5}
                    _ = v37
                default:
                    v31 = v31
                }
                v1 = append(v1, 65535)
            } else {
                if len(v1) > 0 {
                    v1[0] = 955
                }
                if (!true) {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        var v39 float64 = v32
                        var v40 string = "a"
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v2 += v31
                        _ = fi41
                    }
                    for fi42 := 0; fi42 < 14; fi42++ {
                        var v43 int16 = (v34 / (((v34 % ((v34) | 1))) | 1))
                        if false {
                            switch (fi35 * 62) {
                            case -9223372036854775808:
                                var v44 uint16 = uint16(v4)
                                _ = v44
                            case 42:
                                if true {
                                    switch ("\n\t" + (v5 + (("foo bar" + v5) + (v5 + "hello")))) {
                                    case "foo bar":
                                        fmt.Println(v32, v4)
                                    case "\n\t":
                                        v45 := GenS1{F0: 0}
                                        _ = v45
                                    default:
                                        v1 = append(v1, 0)
                                    }
                                    fmt.Println(v43, v31)
                                    _ = len(v1)
                                }
                            default:
                                v2 -= v31
                            }
                            switch (v5 + "\n\t") {
                            case "a":
                                var v46 uint8 = v2
                                _ = v46
                            case "hello":
                                v47 := GenS2{F0: 0, F1: "\n\t"}
                                var v48 uint64 = v4
                                _ = v47
                                _ = v48
                            default:
                                var v49 string = ""
                                _ = v49
                            }
                        }
                        _ = fi42
                        _ = v43
                    }
                } else {
                    for fi50 := 0; fi50 < 19; fi50++ {
                        if ((v4 % (((v4 + v4)) | 1)) >= (v4 & 18446744073709551615)) {
                            v51 := GenS1{F0: 65535}
                            fmt.Println(fi35)
                            if (937 == 0) {
                                var v52 int16 = (v34 << ((0) & 63))
                                _ = v52
                            } else {
                                if true {
                                    switch 648 {
                                    case 1:
                                        switch 158 {
                                        case -1:
                                            v32 -= (v32 + (v32 / v32))
                                        case 9223372036854775807:
                                            if false {
                                                v31 += v2
                                            } else {
                                                v53 := GenS2{F0: 9223372036854775807, F1: (v5 + v5)}
                                                _ = v53
                                            }
                                        }
                                    case 127:
                                        switch (v5 + v5) {
                                        case "hello":
                                            v2 -= v31
                                        case "":
                                            v1 = append(v1, 2147483647)
                                        case "a":
                                            var v54 int = (fi50 / ((2147483647) | 1))
                                            _ = v54
                                        default:
                                            v1 = append(v1, 0)
                                        }
                                    default:
                                        if true {
                                            v55 := GenS2{F0: int64(v4), F1: ""}
                                            if true {
                                                v32 += v32
                                            } else {
                                                v3 += v3
                                            }
                                            _ = v55
                                        }
                                    }
                                    switch 9223372036854775807 {
                                    case 32767:
                                        v56 := GenS1{F0: uint(v2)}
                                        v3 += (((float32(v32) * (float32(float64(float32(v32))) - v3)) - float32(float64(v3))) - ((float32((v32 - v32)) / v3) - float32(v32)))
                                        v31 += v31
                                        _ = v56
                                    case 1:
                                        var v57 string = "\n\t"
                                        _ = v57
                                    default:
                                        switch (fi50 | fi35) {
                                        case 42:
                                            var v58 uint64 = v4
                                            _ = v58
                                        case 32767:
                                            v59 := GenS1{F0: 255}
                                            _ = v59
                                        case 775:
                                            v1 = append(v1, 833)
                                        default:
                                            var v60 uint8 = v31
                                            _ = v60
                                        }
                                    }
                                    var v61 int16 = (v34 + 164)
                                    switch v5 {
                                    case "\n\t":
                                        var v62 int = 42
                                        _ = v62
                                    case "a":
                                        v3 = 1.5
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = 65535
                                        }
                                    }
                                    fmt.Println(v5, v3)
                                    _ = v61
                                }
                                v63 := GenS1{F0: 18446744073709551615}
                                var v64 uint8 = 239
                                _ = v63
                                _ = v64
                            }
                            _ = v51
                        }
                        v65 := GenS3{F0: 42}
                        _ = fi50
                        _ = v65
                    }
                    if true {
                        fmt.Println(v3, v5)
                    }
                }
                for fi66 := 0; fi66 < 10; fi66++ {
                    v34 += v34
                    _ = fi66
                }
            }
            _ = fi35
            _ = v36
        }
        var v67 int32 = 52
        _ = v67
    } else {
        v4 = 534
    }
    var v68 uint = 42
    if false {
        for fi69 := 0; fi69 < 7; fi69++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v5 = v5
            if true {
                if ((^(v68)) >= v68) {
                    var v70 uint32 = 4294967295
                    _ = v70
                } else {
                    if (-1 >= 9223372036854775807) {
                        v34 = v34
                        var v71 int16 = (v34 >> ((v68) & 63))
                        _ = v71
                    } else {
                        var v72 uint64 = (v4 / ((328) | 1))
                        if true {
                            v31 = (^(v31))
                            v5 = "foo bar"
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            switch fi69 {
                            case 32768:
                                v72 += (v72 >> ((uint(v4)) & 63))
                            case -665:
                                var v73 uint16 = 255
                                _ = v73
                            case -9223372036854775808:
                                var v74 uint32 = 985
                                _ = v74
                            }
                            var v75 int16 = (v34 << ((v68) & 63))
                            var v76 int16 = v75
                            _ = v75
                            _ = v76
                        }
                        if (1 >= genFn1((v2 * (^(v2))), "hello")) {
                            if false {
                                fmt.Println(v34, fi69, v31)
                            } else {
                                var v77 string = v5
                                _ = v77
                            }
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        }
                        _ = v72
                    }
                    v32 += v32
                    v78 := GenS1{F0: 42}
                    _ = v78
                }
                var v79 string = "\n\t"
                if true {
                    fmt.Println(v31, fi69, v2)
                }
                _ = v79
            } else {
                v32 += v32
                v31 += v2
                if len(v1) > 0 {
                    v1[0] = 2
                }
            }
            _ = fi69
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v68
}
