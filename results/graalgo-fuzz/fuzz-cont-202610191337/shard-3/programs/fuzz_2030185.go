package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int16; F1 string; F2 int }
type GenS5 struct { F0 uint8; F1 int32; F2 int16 }

func main() {
    v1 := []uint32{2, 2}
    var v2 uint = 2
    var v3 uint32 = 2
    v4 := []uint32{}
    var v5 float32 = float32((1.5 - (-2.25 - 1.5)))
    v5 = 3.14159
    v6 := GenS1{F0: -45, F1: (v3 << (((v2 - (v2 / ((v2) | 1)))) & 63)), F2: (v2 ^ v2)}
    v2 = (v2 % ((uint(v3)) | 1))
    v7 := GenS1{F0: 823, F1: v3, F2: (^(v2))}
    var v8 uint8 = 19
    if (!(!true)) {
        v9 := GenS2{F0: -1, F1: 42}
        _ = v9
    } else {
        var v10 int64 = -655
        _ = v10
    }
    if (!false) {
        for fi11 := 0; fi11 < 7; fi11++ {
            if ((1 < -29) || true) {
                v5 -= (v5 + v5)
                for fi12 := 0; fi12 < 17; fi12++ {
                    if false {
                        fmt.Println(v8, v3)
                    }
                    v8 = (v8 << ((v2) & 63))
                    switch (fi11 * fi11) {
                    case 2147483647:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -1:
                        var v13 int8 = -2
                        _ = v13
                    default:
                        if true {
                            v2 += v2
                            v5 += float32(float64(v5))
                            fmt.Println(v8)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v14 uint = v2
                            _ = v14
                        } else {
                            switch "a" {
                            case "":
                                v5 += v5
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            case "hello":
                                v3 = (^(v3))
                            default:
                                var v15 uint16 = 1
                                v5 = (v5 + (float32((0.1 + 3.14159)) - ((v5 + ((v5 - (((v5 - 1.5) / ((float32((0.1 - 1.5)) + v5) - 0.0)) + 0.0)) / (0.1 - v5))) / v5)))
                                _ = len(v4)
                                _ = v15
                            }
                            var v16 uint32 = v3
                            _ = v16
                        }
                    }
                    _ = fi12
                }
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                switch "" {
                case "foo bar":
                    var v18 uint16 = 988
                    _ = v18
                case "hello":
                    v1 = append(v1, v3)
                default:
                    var v19 int16 = 32767
                    _ = v19
                }
                var v20 float32 = v5
                if ((true && ("" > "hello")) || false) {
                    if len(v1) > 0 {
                        v1[0] = (v3 ^ 1)
                    }
                    v5 = v20
                    if (false && true) {
                        if (true || true) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if (false && (!(((float32(1.5) + 0.0) / (v20 + v5)) == (v20 / v20)))) {
                                var v21 int8 = 0
                                var v22 int8 = (^(v21))
                                v8 -= v8
                                if (787 != 42) {
                                    v21 = v22
                                    var v23 uint8 = (v8 % ((v8) | 1))
                                    _ = v23
                                }
                                _ = v21
                                _ = v22
                            } else {
                                v2 += v2
                                v4 = append(v4, v3)
                                v24 := GenS3{F0: false, F1: (v8 * (^(v8)))}
                                _ = v24
                            }
                            v25 := GenS3{F0: false, F1: 255}
                            var v26 int8 = -128
                            v3 += 42
                            _ = v25
                            _ = v26
                        } else {
                            if ((v8 - (v8 & v8)) >= (v8 / ((127) | 1))) {
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                            } else {
                                v4 = append(v4, v3)
                            }
                            if true {
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch ("hello" + "") {
                                        case "a":
                                            if len(v4) > 0 {
                                                v4[0] = 2147483648
                                            }
                                        case "hello":
                                            if (-1 >= 682) {
                                                switch "" {
                                                case "hello":
                                                    var v27 string = "hello"
                                                    _ = v27
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = (^(v3))
                                                    }
                                                default:
                                                    var v28 string = "a"
                                                    _ = v28
                                                }
                                                v20 = ((v20 * (((v5 - float32(0.1)) / (1.5 / v20)) * (v5 * float32(((1e10 * 1e10) + 1.5))))) - v20)
                                            }
                                        case "":
                                            var v29 int64 = 1000
                                            _ = v29
                                        default:
                                            v20 += 1e10
                                        }
                                    case "\n\t":
                                        if (0 >= -1) {
                                            fmt.Println(v20, fi11, v20)
                                        }
                                    case "":
                                        var v30 int16 = 32767
                                        _ = v30
                                    default:
                                        _ = len(v1)
                                    }
                                    if true {
                                        v5 = (v20 / 3.14159)
                                        v8 -= (^(v8))
                                    }
                                    v8 = uint8(v3)
                                }
                            } else {
                                if ((false && true) || (v3 <= v3)) {
                                    v1 = append(v1, 861)
                                    fmt.Println(v5)
                                    v20 -= 1.5
                                }
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    v4 = append(v4, 0)
                                }
                                var v31 uint32 = 65535
                                v2 = v2
                                _ = v31
                            }
                        }
                        switch fi17 {
                        case 42:
                            v1 = append(v1, (v3 ^ (v3 | uint32(v2))))
                            v32 := GenS1{F0: 824, F1: v3, F2: (v2 << (((v2 | v2)) & 63))}
                            _ = v32
                        case -9223372036854775808:
                            if false {
                                var v33 uint = (v2 * (v2 | v2))
                                _ = v33
                            }
                        case 128:
                            if false {
                                switch (fi11 + (fi11 >> (((v2 << (((v2 / (((v2 & (v2 ^ (v2 & v2)))) | 1))) & 63))) & 63))) {
                                case -608:
                                    fmt.Println(v20, fi11)
                                default:
                                    var v34 int8 = 42
                                    _ = v34
                                }
                                var v35 uint32 = v3
                                v36 := GenS4{F0: 497, F1: "foo bar", F2: (fi17 + (fi11 >> ((v2) & 63)))}
                                switch fi17 {
                                case 128:
                                    v37 := GenS1{F0: 9223372036854775807, F1: 65535, F2: v2}
                                    _ = v37
                                case 32767:
                                    var v38 uint8 = (v8 * 2)
                                    _ = v38
                                }
                                v39 := GenS3{F0: (!(!false)), F1: 0}
                                v8 -= v8
                                fmt.Println(v5, v5, fi11)
                                var v40 bool = true
                                v41 := GenS1{F0: 1, F1: (v3 / ((v3) | 1)), F2: v2}
                                _ = v35
                                _ = v36
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        default:
                            if true {
                                var v42 float32 = (-2.25 / (3.14159 - 1.5))
                                _ = v42
                            }
                        }
                    } else {
                        var v43 int32 = 42
                        _ = v43
                    }
                    fmt.Println(v2, fi11, v3)
                    fmt.Println(fi17, v5, v5)
                }
                _ = fi17
                _ = v20
            }
            _ = fi11
        }
        if (!(0 >= -2147483648)) {
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 float64 = (1e10 * -2.25)
                _ = fi44
                _ = v45
            }
            fmt.Println(v2, v8)
        }
    }
    fmt.Println(v5)
    var v46 int32 = 201
    _ = len(v1)
    var v47 int8 = 0
    v48 := GenS1{F0: -495, F1: v3, F2: v2}
    for fi49 := 0; fi49 < 9; fi49++ {
        v50 := GenS5{F0: v8, F1: (v46 - 42), F2: 790}
        _ = fi49
        _ = v50
    }
    var v51 uint8 = (v8 & v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
