package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 string }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 uint, p1 int, p2 uint64) uint32 {
    v6 := []int{(p1 - p1), p1}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float32 = float32(1.5)
        var v9 float32 = v8
        fmt.Println(p2, fi7, p1)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 487
}
func genFn1() int32 {
    switch "" {
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: 2147483647, F1: 42}
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: -2147483648, F1: 2147483647}
                fmt.Println(fi12)
                v14 := GenS1{F0: -1, F1: 105}
                v15 := GenS1{F0: -1, F1: 1}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint64 = 18446744073709551615
                v18 := GenS1{F0: int32(fi10), F1: 386}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi10
            _ = v11
        }
    case "\n\t":
        var v19 uint8 = 255
        _ = v19
    }
    var v20 int16 = -193
    if true {
        if (-128 != 42) {
            v21 := GenS2{F0: 0, F1: 2, F2: int32(v20)}
            var v22 float32 = (-2.25 + (1e10 / 1.0))
            if true {
                for fi23 := 0; fi23 < 16; fi23++ {
                    switch (fi23 << ((1) & 63)) {
                    case 394:
                        for fi24 := 0; fi24 < 12; fi24++ {
                            v22 = 0.1
                            var v25 int32 = int32(fi23)
                            if true {
                                if true {
                                    switch (fi24 ^ fi24) {
                                    case -9223372036854775808:
                                        switch (("\n\t" + "foo bar") + ("" + ("" + "foo bar"))) {
                                        case "hello":
                                            switch 32768 {
                                            case -681:
                                                v26 := GenS3{F0: (^(v25)), F1: 2, F2: "foo bar"}
                                                _ = v26
                                            case -9223372036854775808:
                                                v27 := GenS2{F0: 429, F1: 42, F2: (-(v25))}
                                                _ = v27
                                            default:
                                                var v28 int = (fi23 >> ((255) & 63))
                                                _ = v28
                                            }
                                        }
                                    }
                                } else {
                                    v29 := GenS1{F0: (v25 & v25), F1: v25}
                                    _ = v29
                                }
                                v30 := GenS1{F0: v25, F1: -1}
                                v20 -= 515
                                _ = v30
                            } else {
                                var v31 float64 = 3.14159
                                _ = v31
                            }
                            _ = fi24
                            _ = v25
                        }
                    case 127:
                        fmt.Println(fi23)
                    default:
                        var v32 uint32 = 794
                        _ = v32
                    }
                    v33 := GenS1{F0: -2147483648, F1: 1}
                    for fi34 := 0; fi34 < 5; fi34++ {
                        switch (-(fi23)) {
                        case 127:
                            v20 -= (v20 ^ 0)
                        case 0:
                            v35 := GenS1{F0: 263, F1: -2147483648}
                            _ = v35
                        case -9223372036854775808:
                            v20 = (v20 / (((^(v20))) | 1))
                        }
                        _ = fi34
                    }
                    _ = fi23
                    _ = v33
                }
                var v36 uint32 = 42
                _ = v36
            } else {
                v22 -= (((v22 - (v22 + ((((float32(0.1) + (((float32(3.14159) - float32(float64(float32(float64((float32((1e10 + 1.5)) + (v22 + v22))))))) + v22) * (v22 - -2.25))) * v22) - 1.5) / 1e10))) / (v22 / v22)) - 1.5)
                var v37 uint = 65535
                v22 = (float32(-2.25) / 1e10)
                _ = v37
            }
            _ = v21
            _ = v22
        } else {
            if false {
                switch 32768 {
                case 42:
                    if false {
                        v38 := GenS4{F0: -1, F1: 128, F2: 0}
                        for fi39 := 0; fi39 < 5; fi39++ {
                            var v40 uint32 = 2147483648
                            var v41 uint32 = genFn0(18446744073709551615, fi39, 4294967295)
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        v20 += v20
                        _ = v38
                    }
                case 127:
                    v42 := []int{-530, 2147483647, 2147483647, 9223372036854775807}
                    _ = v42
                case 2147483647:
                    fmt.Println(v20)
                }
                var v43 int64 = -1
                _ = v43
            } else {
                var v44 uint8 = 127
                var v45 int64 = 1
                fmt.Println(v20, v45)
                _ = v44
                _ = v45
            }
        }
        v46 := GenS5{F0: 2147483648, F1: v20}
        _ = v46
    } else {
        v20 = v20
    }
    var v47 uint = 2
    var v48 bool = false
    _ = v20
    _ = v47
    _ = v48
    return 2147483647
}

func main() {
    var v1 int8 = -43
    var v2 bool = false
    v3 := []string{"hello", "foo bar", ("hello" + ""), "a"}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    var v49 int32 = 985
    var v50 float32 = 3.14159
    var v51 int8 = v1
    v52 := GenS6{F0: 255, F1: 2147483647, F2: 1}
    var v53 uint32 = 768
    var v54 uint64 = 301
    var v55 uint = 390
    fmt.Println(v1, v51, v55)
    _ = len(v4)
    v49 += (v49 >> (((v55 | (v55 >> ((0) & 63)))) & 63))
    if v2 {
        v49 += v49
    }
    v55 = 18446744073709551615
    switch 509 {
    case -219:
        fmt.Println(v51, v51, v1)
    case 127:
        for fi56 := 0; fi56 < 4; fi56++ {
            for fi57 := 0; fi57 < 4; fi57++ {
                v49 = genFn1()
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (v50 > v50) {
                    var v58 string = ((("\n\t" + "\n\t") + "") + (("hello" + "\n\t") + "hello"))
                    _ = v58
                } else {
                    var v59 int64 = -249
                    var v60 float64 = 1e10
                    _ = v59
                    _ = v60
                }
                _ = fi57
            }
            for fi61 := 0; fi61 < 6; fi61++ {
                fmt.Println(fi56, v51, v54)
                var v62 float32 = v50
                fmt.Println(fi56, fi61, v2)
                _ = fi61
                _ = v62
            }
            _ = fi56
        }
    default:
        if true {
            v50 -= (1.5 - v50)
            switch -9223372036854775808 {
            case -9223372036854775808:
                fmt.Println(v54, v54)
            case -1:
                if true {
                    v51 = (-(v51))
                    for fi63 := 0; fi63 < 8; fi63++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 = 127
                        _ = fi63
                    }
                    v55 = v55
                } else {
                    delete(v5, 4294967295)
                }
            }
            if (v54 <= v54) {
                v50 += v50
                v55 -= (^(v55))
                v3 = append(v3, "\n\t")
            } else {
                switch "foo bar" {
                case "hello":
                    v64 := GenS1{F0: genFn1(), F1: int32(v53)}
                    _ = v64
                case "foo bar":
                    for fi65 := 0; fi65 < 11; fi65++ {
                        v66 := GenS2{F0: 2, F1: 0, F2: (v49 + (v49 - 1))}
                        _ = fi65
                        _ = v66
                    }
                case "\n\t":
                    v3 = append(v3, "a")
                default:
                    v49 += v49
                }
            }
        } else {
            v49 -= (v49 + 42)
            var v67 uint32 = (v53 | 2147483647)
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 5; fi68++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
