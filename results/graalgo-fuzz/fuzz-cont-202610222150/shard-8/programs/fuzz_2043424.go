package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64 }
func genFn0() int {
    switch 2147483647 {
    case 128:
        if false {
            var v4 int16 = 1000
            var v5 float64 = 1.5
            _ = v4
            _ = v5
        } else {
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := GenS1{F0: "a", F1: 2, F2: 62}
                if true {
                    var v8 float32 = 0.1
                    _ = v8
                } else {
                    var v9 int = fi6
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
            v10 := []int{-1}
            if (253 < 1) {
                var v11 int = 32767
                fmt.Println(v11)
                _ = v11
            } else {
                var v12 uint32 = 2147483647
                switch -571 {
                case -9223372036854775808:
                    for fi13 := 0; fi13 < 15; fi13++ {
                        if (0.1 > 0.0) {
                            _ = len(v10)
                            v14 := GenS2{F0: 42, F1: -128}
                            v12 = (v12 | (v12 % ((2147483647) | 1)))
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 19; fi15++ {
                                var v16 int16 = 1
                                switch (fi15 + 42) {
                                case 32767:
                                    fmt.Println(fi13, v16, fi13)
                                case 4294967295:
                                    if ((!(((1e10 - -2.25) - 0.1) < float64(0.0))) || (false || true)) {
                                        var v17 int32 = -1
                                        v12 = v12
                                        _ = v17
                                    }
                                case -2147483648:
                                    switch fi15 {
                                    case 413:
                                        var v18 int16 = -803
                                        _ = v18
                                    case -154:
                                        fmt.Println(fi15, fi15, fi13)
                                    case -152:
                                        var v19 uint8 = 42
                                        switch 2147483647 {
                                        case 4294967295:
                                            if (1e10 == -2.25) {
                                                v20 := GenS3{F0: (v16 / ((v16) | 1)), F1: 0}
                                                switch fi15 {
                                                case -423:
                                                    v16 -= v16
                                                }
                                                v19 -= v19
                                                _ = v20
                                            } else {
                                                var v21 uint32 = (v12 | uint32(v19))
                                                _ = v21
                                            }
                                        case 32767:
                                            switch "" {
                                            case "foo bar":
                                                if true {
                                                    var v22 float32 = -2.25
                                                    _ = v22
                                                }
                                            default:
                                                var v23 bool = (-1 == -258)
                                                v12 = (v12 % ((972) | 1))
                                                var v24 string = ""
                                                _ = v23
                                                _ = v24
                                            }
                                        default:
                                            if (!(42 < 1)) {
                                                v19 = v19
                                            }
                                        }
                                        v25 := GenS4{F0: 1}
                                        _ = v19
                                        _ = v25
                                    }
                                }
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(fi13, fi13)
                            v12 = v12
                        }
                        switch fi13 {
                        case -2147483648:
                            var v26 int16 = -32768
                            _ = v26
                        case 0:
                            if false {
                                var v27 bool = (362 <= 0)
                                if true {
                                    fmt.Println(v27, v12)
                                    var v28 uint32 = v12
                                    _ = v28
                                }
                                var v29 uint16 = 2
                                _ = v27
                                _ = v29
                            }
                        case 32768:
                            _ = len(v10)
                        }
                        _ = fi13
                    }
                case 32768:
                    if (!true) {
                        v12 = v12
                    } else {
                        if (!((2 > 252) || true)) {
                            v12 -= v12
                            v12 += 4294967295
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        } else {
                            var v30 string = (("" + "\n\t") + (("a" + "") + "\n\t"))
                            v31 := GenS2{F0: 42, F1: 42}
                            v30 = ""
                            _ = v30
                            _ = v31
                        }
                        v12 = (v12 % ((v12) | 1))
                    }
                case 42:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 int16 = 1000
                    _ = fi32
                    _ = v33
                }
                _ = v12
            }
            _ = v10
        }
    case 538:
        var v34 uint32 = 2
        _ = v34
    case 9223372036854775807:
        var v35 bool = (0 >= 0)
        _ = v35
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int = fi36
            _ = fi36
            _ = v37
        }
    }
    var v38 uint = 186
    v39 := GenS1{F0: "hello", F1: v38, F2: -11}
    for fi40 := 0; fi40 < 1; fi40++ {
        fmt.Println(fi40)
        _ = fi40
    }
    var v41 int16 = 649
    _ = v38
    _ = v39
    _ = v41
    return 2147483647
}
func genFn1(p0 int32, p1 int16, p2 int8) int16 {
    v42 := []int{32768, 2147483647}
    for fi43 := 0; fi43 < 14; fi43++ {
        p0 += p0
        _ = fi43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -1
    var v3 uint64 = 0
    switch -1 {
    case 187:
        var v44 int = 32767
        _ = v44
    case -809:
        v2 -= (v2 & 937)
    case 4294967295:
        if false {
            for fi45 := 0; fi45 < 4; fi45++ {
                if true {
                    switch fi45 {
                    case -977:
                        var v46 int = (fi45 * (^(fi45)))
                        _ = v46
                    case 2147483647:
                        v47 := GenS3{F0: v2, F1: (^(v3))}
                        _ = v47
                    case 128:
                        v48 := []int{687, fi45, fi45}
                        _ = v48
                    }
                    v2 += 109
                    for fi49 := 0; fi49 < 8; fi49++ {
                        var v50 uint8 = 1
                        if (1 != 42) {
                            v51 := GenS3{F0: genFn1(42, v2, 127), F1: v3}
                            _ = v51
                        }
                        var v52 int32 = -1
                        _ = fi49
                        _ = v50
                        _ = v52
                    }
                }
                v53 := []int{fi45, (fi45 + (fi45 - fi45)), genFn0()}
                switch "" {
                case "foo bar":
                    var v54 string = "hello"
                    _ = v54
                case "\n\t":
                    v3 -= (v3 | 808)
                }
                _ = fi45
                _ = v53
            }
        }
    default:
        fmt.Println(v2, v2)
    }
    var v55 uint = 42
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 uint64 = (^(v3))
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        v1[(fi58 & fi58)] = uint32(v55)
        switch -1 {
        case 127:
            switch "" {
            case "\n\t":
                v59 := GenS2{F0: 534, F1: -1}
                _ = v59
            case "a":
                v2 -= v2
            case "":
                switch (("a" + "hello") + ("foo bar" + (("foo bar" + "hello") + ""))) {
                case "":
                    v2 = (v2 * (^(v2)))
                }
            default:
                v60 := []int{int(v2)}
                _ = v60
            }
        case -784:
            delete(v1, (fi58 / (((fi58 & int(v55))) | 1)))
        }
        _ = fi58
    }
    v55 = 42
    var v61 uint = 18446744073709551615
    v62 := GenS2{F0: (^(v3)), F1: -1}
    var v63 uint16 = 255
    v64 := []int{66}
    v65 := GenS2{F0: 42, F1: -1}
    for fi66 := 0; fi66 < 5; fi66++ {
        var v67 float64 = ((-2.25 * 0.1) + 0.1)
        _ = fi66
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
