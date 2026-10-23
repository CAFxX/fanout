package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 4294967295
    var v3 uint16 = 32767
    var v4 int32 = -253
    var v5 uint = 4294967295
    var v6 uint64 = (v2 << ((v5) & 63))
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint16 = v3
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: (v4 % ((0) | 1)), F1: "a"}
            var v11 string = ("\n\t" + "")
            fmt.Println(v8)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint8 = 255
        _ = fi7
        _ = v12
        _ = v8
    }
    v13 := []int{}
    var v14 uint32 = 956
    var v15 int8 = -1
    v16 := GenS2{F0: -1, F1: (v4 * v4)}
    if true {
        var v17 bool = (true && false)
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint8 = 128
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS2{F0: 0, F1: v4}
                if ("\n\t" == "a") {
                    switch (fi20 - fi20) {
                    case 128:
                        _ = len(v1)
                    }
                    v5 = (v5 ^ 4294967295)
                    _ = len(v1)
                }
                v19 = v19
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        v15 = 1
        if false {
            if len(v13) > 0 {
                v13[0] = 4294967295
            }
            if (v4 <= v4) {
                if (!((0.0 - 0.1) != (1.5 - (0.1 + (0.0 * 0.1))))) {
                    if (196 == 42) {
                        v22 := GenS1{F0: -2147483648, F1: "\n\t"}
                        _ = v22
                    }
                }
                v4 = 0
            } else {
                v14 += (^(v14))
                var v23 int64 = 9223372036854775807
                _ = v23
            }
            var v24 int32 = (v4 & (v4 / ((v4) | 1)))
            _ = v24
        }
        v14 -= (v14 + 4294967295)
    }
    if true {
        var v25 float64 = 3.14159
        var v26 uint64 = 18446744073709551615
        v6 = uint64(v3)
        _ = v25
        _ = v26
    }
    _ = len(v1)
    v13 = append(v13, 513)
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 string = "\n\t"
        switch "" {
        case "foo bar":
            if false {
                v2 = 122
                if true {
                    fmt.Println(v2, fi27)
                } else {
                    v29 := GenS3{F0: v28, F1: 255}
                    for fi30 := 0; fi30 < 19; fi30++ {
                        _ = len(v1)
                        var v31 int = 837
                        _ = len(v13)
                        var v32 int8 = (v15 | v15)
                        switch (v28 + v28) {
                        case "":
                            v33 := GenS1{F0: 2147483647, F1: ((v28 + v28) + "")}
                            _ = v33
                        case "\n\t":
                            v34 := GenS2{F0: 32767, F1: -2147483648}
                            _ = v34
                        case "foo bar":
                            switch v28 {
                            case "hello":
                                v5 += (v5 * uint(v6))
                            case "\n\t":
                                var v35 uint8 = 126
                                _ = v35
                            case "foo bar":
                                var v36 float64 = 1e10
                                _ = v36
                            default:
                                v37 := GenS3{F0: v28, F1: 122}
                                _ = v37
                            }
                            if (v4 != int32(v15)) {
                                var v38 string = "foo bar"
                                var v39 uint8 = 127
                                var v40 float32 = float32(1.5)
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            v6 += (v2 - (^(v6)))
                        }
                        var v41 uint = (^(v5))
                        var v42 uint8 = 128
                        switch v31 {
                        case -1:
                            v43 := GenS4{F0: (v31 ^ (-(fi27))), F1: (^(v5)), F2: -590}
                            _ = v43
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v41
                        _ = v42
                    }
                    for fi44 := 0; fi44 < 7; fi44++ {
                        if true {
                            switch v28 {
                            case "\n\t":
                                fmt.Println(v3, v3, v28)
                            default:
                                v15 += -128
                            }
                            var v45 string = "hello"
                            _ = v45
                        } else {
                            v2 += (v6 | v6)
                        }
                        v46 := GenS5{F0: (v6 ^ (v6 << ((v5) & 63)))}
                        v47 := GenS5{F0: (v6 ^ (v2 ^ v6))}
                        v13 = append(v13, -732)
                        _ = fi44
                        _ = v46
                        _ = v47
                    }
                    _ = v29
                }
            }
        case "a":
            v28 = v28
        case "hello":
            if true {
                for fi48 := 0; fi48 < 20; fi48++ {
                    var v49 bool = (float32(1e10) < 1.5)
                    _ = fi48
                    _ = v49
                }
            } else {
                if false {
                    v50 := GenS6{F0: 9, F1: -797}
                    _ = v50
                } else {
                    var v51 uint32 = 156
                    _ = v51
                }
                v14 -= (v14 * v14)
                v5 = v5
            }
        default:
            v52 := GenS4{F0: (fi27 ^ 1), F1: v5, F2: 231}
            _ = v52
        }
        v28 = ((("hello" + "") + "foo bar") + (v28 + v28))
        _ = fi27
        _ = v28
    }
    v15 = (v15 ^ (^(v15)))
    v53 := GenS4{F0: 1, F1: (^(v5)), F2: -1}
    v14 -= v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
