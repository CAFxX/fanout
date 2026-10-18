package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int }

func main() {
    var v1 int = -874
    var v2 int8 = 1
    var v3 bool = false
    var v4 uint = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 900}
        _ = fi5
        _ = v6
    }
    var v7 string = ("" + ("\n\t" + (("\n\t" + ("a" + "hello")) + "a")))
    v8 := GenS1{F0: 32767}
    if (2 != 621) {
        var v9 uint64 = 18446744073709551615
        var v10 int8 = (^(v2))
        _ = v10
        _ = v9
    } else {
        if v3 {
            if false {
                v11 := GenS1{F0: 1}
                if v3 {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch v7 {
                        case "a":
                            v13 := GenS2{F0: 4294967295, F1: -32768}
                            _ = v13
                        }
                        v14 := []int{9223372036854775807, v1}
                        var v15 int16 = 366
                        _ = fi12
                        _ = v14
                        _ = v15
                    }
                } else {
                    switch 644 {
                    case 873:
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 uint32 = 664
                            v18 := []int{(v1 % (((fi16 >> (((v4 & v4)) & 63))) | 1)), v1, fi16, (v1 * -1)}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    default:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 int = -9223372036854775808
                            _ = fi19
                            _ = v20
                        }
                    }
                    v21 := GenS1{F0: 1}
                    _ = v21
                }
                _ = v11
            }
            fmt.Println(v2, v1, v3)
            if ((!v3) && (!(!(-1 <= -1)))) {
                v1 = v1
                var v22 string = ("foo bar" + ("foo bar" + ("\n\t" + v7)))
                _ = v22
            } else {
                fmt.Println(v1)
                switch v1 {
                case 312:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v3 = v3
                        switch v7 {
                        case "hello":
                            for fi24 := 0; fi24 < 18; fi24++ {
                                v25 := []int{v1, 2147483647, (fi24 << (((v4 - 0)) & 63))}
                                _ = fi24
                                _ = v25
                            }
                        }
                        _ = fi23
                    }
                case 32767:
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v2 = (v2 & (v2 - v2))
                        _ = fi26
                    }
                default:
                    v27 := GenS2{F0: 2147483647, F1: 1000}
                    _ = v27
                }
                var v28 uint16 = 0
                _ = v28
            }
        }
        var v29 bool = (-1 <= 28)
        v30 := GenS3{F0: v1, F1: v7, F2: 0}
        _ = v29
        _ = v30
    }
    v2 += 31
    switch v7 {
    case "hello":
        fmt.Println(v7)
    case "":
        if v3 {
            var v31 uint16 = 255
            v32 := GenS1{F0: (v31 | (v31 & v31))}
            if false {
                v4 = (^(v4))
                v33 := []int{}
                _ = v33
            }
            _ = v31
            _ = v32
        } else {
            v34 := []int{}
            _ = len(v34)
            _ = v34
        }
    }
    if v3 {
        var v35 uint32 = 435
        if v3 {
            if v3 {
                v36 := []int{v1, (v1 ^ 225), v1}
                for fi37 := 0; fi37 < 10; fi37++ {
                    if true {
                        switch v7 {
                        case "a":
                            if v3 {
                                var v38 uint16 = 0
                                var v39 uint = 909
                                _ = v38
                                _ = v39
                            }
                        case "\n\t":
                            v35 = 2
                        }
                    }
                    fmt.Println(v3, v1, v4)
                    switch -1 {
                    case 1:
                        switch 1 {
                        case 32767:
                            var v40 float32 = float32(((1e10 - (0.1 + 1e10)) / 3.14159))
                            _ = v40
                        default:
                            switch -1 {
                            case 32767:
                                v2 -= (v2 | (v2 | v2))
                            case 755:
                                var v41 bool = v3
                                _ = v41
                            default:
                                v4 -= (v4 + v4)
                            }
                        }
                    case 128:
                        if len(v36) > 0 {
                            v36[0] = fi37
                        }
                    case 127:
                        fmt.Println(v1, v4)
                    default:
                        v42 := GenS4{F0: -9223372036854775808, F1: 321, F2: v4}
                        _ = v42
                    }
                    _ = fi37
                }
                v1 = 32767
                _ = v36
            }
            var v43 int32 = 715
            v44 := []int{(v1 * v1), (v1 << ((v4) & 63)), 9223372036854775807, v1}
            _ = v43
            _ = v44
        }
        if v3 {
            fmt.Println(v35, v35, v7)
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := GenS2{F0: 1, F1: 67}
                var v47 int64 = 1
                _ = fi45
                _ = v46
                _ = v47
            }
        } else {
            fmt.Println(v35, v35)
            for fi48 := 0; fi48 < 5; fi48++ {
                v49 := []int{(-(v1)), (-(v1)), v1, fi48}
                for fi50 := 0; fi50 < 13; fi50++ {
                    var v51 float64 = -2.25
                    var v52 int16 = 1000
                    v53 := GenS5{F0: (!false), F1: 18446744073709551615, F2: 42}
                    _ = fi50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                _ = fi48
                _ = v49
            }
            fmt.Println(v7, v2)
        }
        _ = v35
    } else {
        v2 -= v2
        v4 = (^(v4))
        v7 = (((v7 + v7) + v7) + v7)
    }
    if (0.0 < 0.0) {
        var v54 uint64 = 0
        _ = v54
    }
    fmt.Println(v4, v4)
    v1 += 78
    for fi55 := 0; fi55 < 17; fi55++ {
        v7 = v7
        for fi56 := 0; fi56 < 9; fi56++ {
            var v57 uint8 = 123
            v58 := GenS4{F0: 42, F1: (^(v4)), F2: v4}
            if v3 {
                if v3 {
                    v59 := []int{(v1 << ((v4) & 63)), v1, (v1 % ((fi56) | 1))}
                    if true {
                        v1 = fi55
                    }
                    _ = v59
                } else {
                    v60 := []int{(^(v1))}
                    v3 = true
                    _ = v60
                }
            } else {
                v61 := []int{127, 4294967295, v1, (-(fi55))}
                v4 += v4
                _ = v61
            }
            _ = fi56
            _ = v57
            _ = v58
        }
        _ = fi55
    }
    switch v7 {
    case "hello":
        v4 += 65535
    default:
        var v62 int16 = int16(v1)
        _ = v62
    }
    v63 := []int{9223372036854775807, v1}
    v64 := GenS6{F0: v7}
    fmt.Println(v1)
    var v65 uint8 = 42
    if ((((((v4 ^ v4) > 42) && true) || v3) && (false && v3)) && ((0.0 + 0.0) >= 0.1)) {
        if len(v63) > 0 {
            v63[0] = 0
        }
        if len(v63) > 0 {
            v63[0] = v1
        }
        for fi66 := 0; fi66 < 19; fi66++ {
            var v67 float32 = 0.0
            switch (fi66 >> ((0) & 63)) {
            case 250:
                _ = len(v63)
            case -1:
                v7 = (((("hello" + ("\n\t" + ("foo bar" + v7))) + "a") + v7) + "")
            default:
                v68 := GenS7{F0: (v1 + (fi66 - -1))}
                _ = v68
            }
            var v69 bool = v3
            _ = fi66
            _ = v67
            _ = v69
        }
    } else {
        if v3 {
            var v70 uint32 = 1
            if ((1 < v65) && (3.14159 < float64(3.14159))) {
                fmt.Println(v2, v70)
                var v71 uint32 = v70
                _ = v71
            } else {
                v72 := GenS7{F0: (-(v1))}
                v3 = (!false)
                var v73 uint32 = v70
                _ = v72
                _ = v73
            }
            _ = v70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v65
    _ = v7
    _ = v8
}
