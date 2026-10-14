package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
func genFn0() bool {
    var v7 bool = (1000 < -581)
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v7)
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = v7
    return v7
}
func genFn1(p0 uint8, p1 int, p2 int16) int8 {
    switch (("a" + "\n\t") + "\n\t") {
    case "\n\t":
        var v10 int32 = 855
        _ = v10
    case "foo bar":
        fmt.Println(p0, p0, p2)
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(p0, p2, fi12)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 bool = false
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 0
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    v6 := []int{-2147483648, 32767}
    v3 += v3
    var v13 float32 = 1.5
    if ((1 > 42) && ((0 <= 128) || (824 > 65535))) {
        v14 := GenS1{F0: 431, F1: 326}
        v1 = v1
        _ = v14
    }
    v3 = -660
    fmt.Println(v2, v2)
    switch (("a" + "hello") + ("" + ("\n\t" + "\n\t"))) {
    case "foo bar":
        v6 = append(v6, 0)
    case "hello":
        v15 := GenS2{F0: v1}
        _ = v15
    case "":
        if ((v13 - -2.25) == float32(((((3.14159 * ((0.1 - (0.1 + -2.25)) * float64(((3.14159 + v13) / v13)))) - (((0.1 + (-2.25 - 0.1)) + 0.1) * float64(((v13 / 1.0) - (-2.25 / 1.5))))) * float64(((1e10 / 1.0) / (v13 + (v13 + float32((0.1 * (3.14159 + 1e10)))))))) + ((1.5 * 3.14159) + ((1e10 / (3.14159 * 1.5)) / (3.14159 / 1.5)))))) {
            for fi16 := 0; fi16 < 13; fi16++ {
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v1, fi16, fi17)
                    case "a":
                        var v18 uint16 = 32767
                        _ = v18
                    default:
                        switch ("" + "hello") {
                        case "a":
                            v19 := GenS3{F0: 127, F1: 744, F2: v3}
                            _ = v19
                        }
                        var v20 int = fi17
                        var v21 uint32 = 562
                        _ = v20
                        _ = v21
                    }
                    var v22 uint = 346
                    _ = fi17
                    _ = v22
                }
                switch "a" {
                case "foo bar":
                    for fi23 := 0; fi23 < 13; fi23++ {
                        switch "hello" {
                        case "":
                            if (0.0 > 3.14159) {
                                v3 = v3
                                v1 = (0 <= 9223372036854775807)
                                fmt.Println(v3)
                            }
                        case "foo bar":
                            fmt.Println(v13)
                        default:
                            fmt.Println(v13, v13, v1)
                        }
                        v3 = v3
                        _ = fi23
                    }
                case "\n\t":
                    fmt.Println(v3)
                default:
                    if v1 {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if (int16(v2) > -32768) {
                                switch fi16 {
                                case 307:
                                    v25 := GenS4{F0: 165, F1: 0, F2: 1}
                                    _ = v25
                                }
                            }
                            switch -371 {
                            case -618:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if true {
                                    delete(v4, fi16)
                                    v26 := v5["foo bar"]
                                    v3 = (^(v3))
                                    _ = v26
                                }
                                delete(v4, fi16)
                            case -809:
                                fmt.Println(v3, v1, fi16)
                            case 9223372036854775807:
                                if v1 {
                                    v13 -= float32(3.14159)
                                    v27 := GenS5{F0: v3}
                                    _ = len(v5)
                                    v28 := GenS2{F0: v1}
                                    if v1 {
                                        var v29 int16 = -448
                                        v2 = v2
                                        _ = v29
                                    }
                                    _ = v27
                                    _ = v28
                                }
                            }
                            var v30 uint16 = 0
                            _ = fi24
                            _ = v30
                        }
                        var v31 float64 = 3.14159
                        _ = v31
                    }
                }
                _ = fi16
            }
            var v32 uint32 = 2147483648
            _ = v32
        }
    }
    v1 = v1
    v3 += (v3 ^ v3)
    for fi33 := 0; fi33 < 12; fi33++ {
        if v1 {
            _ = len(v5)
            v34 := GenS6{F0: 447}
            v13 += v13
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            if v1 {
                v1 = v1
                v6 = append(v6, fi33)
                switch "a" {
                case "\n\t":
                    v36 := GenS6{F0: 107}
                    _ = v36
                default:
                    v2 = (^(v2))
                }
                var v37 uint = 744
                v38 := v5[("" + "a")]
                _ = v37
                _ = v38
            } else {
                fmt.Println(fi33, v3)
            }
            if true {
                v1 = v1
            }
            fmt.Println(fi35, v2, fi33)
            var v39 uint32 = 42
            switch -1 {
            case -2147483648:
                delete(v4, fi35)
            default:
                var v40 int16 = 42
                _ = v40
            }
            fmt.Println(v1)
            v2 = 4294967295
            _ = fi35
            _ = v39
        }
        _ = fi33
    }
    v41 := GenS6{F0: 42}
    var v42 uint32 = 237
    var v43 uint16 = 193
    switch 32767 {
    case 32767:
        var v44 float64 = ((0.1 / 1.5) / 1.5)
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
