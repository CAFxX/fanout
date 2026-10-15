package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 bool; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 int16 = 1
    var v2 uint16 = 920
    var v3 int32 = 1
    var v4 uint16 = (v2 % ((2) | 1))
    var v5 uint8 = 37
    v6 := make(map[uint32]int)
    fmt.Println(v3, v2)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = -2.25
        for fi9 := 0; fi9 < 6; fi9++ {
            v1 -= v1
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if (100 <= 0) {
        v10 := GenS1{F0: 255, F1: (v1 | v1), F2: -128}
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 1, F1: 32767, F2: 127}
                _ = fi12
                _ = v13
            }
            v5 -= (v5 & (^(v5)))
            for fi14 := 0; fi14 < 11; fi14++ {
                if (42 < 255) {
                    var v15 int8 = -1
                    v5 = v5
                    fmt.Println(v15)
                    _ = v15
                } else {
                    var v16 string = "foo bar"
                    var v17 uint8 = v5
                    _ = v16
                    _ = v17
                }
                fmt.Println(v4, fi14, v4)
                _ = fi14
            }
            _ = fi11
        }
        _ = v10
    }
    var v18 uint64 = 18446744073709551615
    v4 += v4
    v19 := []int{4294967295, -2147483648, 128}
    var v20 uint = 18446744073709551615
    fmt.Println(v20)
    switch int(v1) {
    case 4294967295:
        v1 = v1
    case 2147483647:
        v1 -= (v1 + v1)
    default:
        for fi21 := 0; fi21 < 6; fi21++ {
            if (v4 < (^(v4))) {
                var v22 int32 = 150
                if ((v18 | (v18 / (((v18 >> ((v20) & 63))) | 1))) == 455) {
                    if true {
                        switch ("hello" + (("foo bar" + ("hello" + ("" + ("hello" + "\n\t")))) + "")) {
                        case "a":
                            if (false && (true && (42 == 335))) {
                                v4 += v2
                                fmt.Println(v22)
                            }
                        case "foo bar":
                            v23 := GenS1{F0: v20, F1: (v1 & (v1 | (v1 - v1))), F2: 42}
                            _ = v23
                        default:
                            for fi24 := 0; fi24 < 3; fi24++ {
                                v25 := GenS1{F0: uint(v5), F1: v1, F2: -43}
                                var v26 int32 = (^(v3))
                                v3 += (v22 / (((v22 % ((v26) | 1))) | 1))
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        }
                        fmt.Println(v2)
                    } else {
                        var v27 int = -2147483648
                        if ((^(v20)) == v20) {
                            var v28 int = v27
                            for fi29 := 0; fi29 < 4; fi29++ {
                                var v30 float32 = ((3.14159 + ((((((1e10 - 0.0) / 0.1) * (1.5 - float32(1.5))) + 0.0) + 0.1) / -2.25)) - 1e10)
                                _ = fi29
                                _ = v30
                            }
                            v31 := GenS2{F0: 42, F1: "hello", F2: v27}
                            _ = v28
                            _ = v31
                        } else {
                            v32 := GenS1{F0: 691, F1: v1, F2: -128}
                            var v33 string = "hello"
                            fmt.Println(v4, fi21, v20)
                            _ = v32
                            _ = v33
                        }
                        _ = v27
                    }
                    var v34 string = "a"
                    for fi35 := 0; fi35 < 7; fi35++ {
                        if (v22 < (v3 % ((0) | 1))) {
                            if (false || false) {
                                var v36 uint32 = 42
                                _ = v36
                            }
                        } else {
                            if true {
                                v1 = (v1 * v1)
                                v37 := v6[42]
                                v3 = 396
                                _ = v37
                            }
                            if true {
                                if ((v18 & 498) <= (v18 % ((v18) | 1))) {
                                    v20 = (v20 << (((v20 + v20)) & 63))
                                    v38 := GenS3{F0: -48}
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = v38
                                }
                                var v39 uint64 = v18
                                v4 += 255
                                _ = v39
                            }
                        }
                        fmt.Println(v20)
                        if true {
                            v19 = append(v19, (fi21 / (((fi35 % ((fi35) | 1))) | 1)))
                            var v40 float32 = 0.1
                            if (v18 < (^(v18))) {
                                var v41 int32 = (v22 & (^(v22)))
                                _ = v41
                            } else {
                                if (fi21 <= int(v5)) {
                                    var v42 int16 = 1000
                                    _ = v42
                                } else {
                                    if false {
                                        if (!true) {
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            switch (^(fi35)) {
                                            case 810:
                                                v1 = v1
                                            case 1:
                                                var v43 int64 = -9223372036854775808
                                                _ = v43
                                            }
                                        }
                                    }
                                    v44 := GenS2{F0: (v4 & 32767), F1: v34, F2: (fi35 / ((4294967295) | 1))}
                                    v3 -= (v22 >> ((v20) & 63))
                                    _ = v44
                                }
                                v45 := GenS4{F0: "hello", F1: true, F2: -32768}
                                var v46 int = int(v20)
                                _ = v45
                                _ = v46
                            }
                            if (((-2.25 / 1.0) / (-2.25 * (((((-2.25 / 1.0) * 0.0) / float64((v40 / v40))) - (1e10 + ((3.14159 / 1.5) - -2.25))) / (((0.0 / (0.0 - float64(v40))) / 1e10) / (((0.1 - float64(v40)) / ((3.14159 - (float64(v40) - (0.1 + -2.25))) / float64((v40 - (((v40 + (v40 * (v40 - (v40 * 1.5)))) + v40) / (float32((3.14159 - 0.1)) * v40)))))) * float64(1.5)))))) < 1.5) {
                                var v47 uint16 = v2
                                v4 = (^(v4))
                                v3 += 2147483647
                                _ = v47
                            } else {
                                var v48 int = fi21
                                _ = v48
                            }
                            if false {
                                v22 = v22
                                v2 += v4
                                fmt.Println(v2)
                            } else {
                                switch v34 {
                                case "foo bar":
                                    v1 -= v1
                                }
                                switch 109 {
                                case 9223372036854775807:
                                    var v49 int8 = -71
                                    _ = v49
                                case -729:
                                    if true {
                                        fmt.Println(fi21, v40)
                                    }
                                case 1:
                                    var v50 uint16 = 2
                                    _ = v50
                                }
                                fmt.Println(fi35, v5, v1)
                            }
                            _ = v40
                        }
                        var v51 uint32 = uint32(v18)
                        _ = len(v19)
                        v1 = v1
                        if true {
                            if ((-(v1)) == 1000) {
                                switch (^(fi35)) {
                                case 865:
                                    v52 := GenS5{F0: (-2147483648 < v3)}
                                    _ = v52
                                case -128:
                                    v18 = v18
                                case -823:
                                    if false {
                                        _ = len(v19)
                                        v3 += v3
                                        fmt.Println(v34)
                                        var v53 int16 = v1
                                        if true {
                                            var v54 uint16 = 42
                                            _ = v54
                                        }
                                        v34 = v34
                                        fmt.Println(v5, v4, fi21)
                                        _ = v53
                                    } else {
                                        switch v34 {
                                        case "a":
                                            fmt.Println(v18)
                                        case "hello":
                                            var v55 int = -9223372036854775808
                                            _ = v55
                                        default:
                                            var v56 float32 = 0.1
                                            _ = v56
                                        }
                                        var v57 int = (fi21 ^ (^(fi21)))
                                        _ = v57
                                    }
                                }
                                v18 += v18
                                _ = len(v19)
                            } else {
                                var v58 float64 = 0.0
                                _ = v58
                            }
                        }
                        fmt.Println(fi35)
                        _ = fi35
                        _ = v51
                    }
                    _ = v34
                } else {
                    var v59 float64 = float64(3.14159)
                    _ = v59
                }
                v18 -= 42
                _ = v22
            }
            v2 = v2
            switch "" {
            case "hello":
                v60 := GenS6{F0: 2147483648}
                _ = v60
            case "a":
                fmt.Println(v1, v20)
            case "foo bar":
                if len(v19) > 0 {
                    v19[0] = (fi21 - 32768)
                }
            }
            _ = fi21
        }
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v61 := GenS2{F0: (v4 | (^(v4))), F1: "", F2: 850}
    v18 -= 0
    for fi62 := 0; fi62 < 14; fi62++ {
        var v63 uint64 = 127
        if (false && ((v3 / ((v3) | 1)) > (v3 / ((v3) | 1)))) {
            if false {
                var v64 int8 = 35
                for fi65 := 0; fi65 < 13; fi65++ {
                    v4 -= 2
                    fmt.Println(v3, v18)
                    v66 := GenS3{F0: (v64 % ((v64) | 1))}
                    if true {
                        fmt.Println(v20)
                    } else {
                        var v67 int16 = (v1 - (^(v1)))
                        _ = v67
                    }
                    v3 = v3
                    _ = fi65
                    _ = v66
                }
                _ = v64
            }
            var v68 uint8 = (v5 + 0)
            _ = v68
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
