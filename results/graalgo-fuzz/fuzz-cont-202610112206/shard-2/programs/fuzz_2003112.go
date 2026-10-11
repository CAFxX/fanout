package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    v4 := []uint32{2147483648}
    var v5 int16 = -32768
    var v6 string = "foo bar"
    _ = len(v4)
    var v7 uint64 = 4294967295
    fmt.Println(v2, v2, v6)
    for fi8 := 0; fi8 < 18; fi8++ {
        v6 = "a"
        _ = fi8
    }
    v9 := GenS1{F0: v7}
    switch v6 {
    case "a":
        _ = len(v1)
    case "\n\t":
        for fi10 := 0; fi10 < 11; fi10++ {
            switch (fi10 | (fi10 << ((42) & 63))) {
            case -86:
                var v11 float64 = float64(((v2 - v2) + 3.14159))
                _ = v11
            case -815:
                switch fi10 {
                case 0:
                    delete(v3, (v6 + "foo bar"))
                case 42:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 float32 = (1e10 + float32(0.0))
                        if false {
                            v2 = (v13 / (1.5 + 0.0))
                            v14 := GenS2{F0: (v6 + (v6 + v6)), F1: 42}
                            _ = v14
                        } else {
                            if true {
                                var v15 uint16 = 32767
                                _ = v15
                            } else {
                                if false {
                                    var v16 uint32 = 2147483647
                                    _ = v16
                                }
                                switch -827 {
                                case 742:
                                    if (fi10 <= (fi10 - fi12)) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        v4 = append(v4, 42)
                                        v7 = 97
                                    }
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v17 int64 = 42
                                v5 -= -576
                                v18 := GenS1{F0: (v7 >> ((4294967295) & 63))}
                                v13 = float32(1.5)
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS1{F0: (v7 * v7)}
                            _ = v19
                        }
                        _ = fi12
                        _ = v13
                    }
                case -756:
                    v20 := GenS3{F0: -128}
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 8; fi21++ {
                        fmt.Println(fi10)
                        _ = fi21
                    }
                }
            }
            v5 = -398
            _ = fi10
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v6 = v6
        _ = fi22
    }
    var v23 uint8 = 0
    v5 = v5
    var v24 int16 = (v5 % (((v5 >> ((65535) & 63))) | 1))
    if (!false) {
        v1[(((v6 + v6) + v6) + v6)] = v6
        if ((v23 * v23) >= 36) {
            var v25 uint64 = v7
            if (true && (v24 != v5)) {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v27 := GenS4{F0: true, F1: 32767, F2: (v24 >> ((18446744073709551615) & 63))}
                    var v28 float64 = float64(float32(0.1))
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                var v29 bool = true
                delete(v3, (("hello" + "foo bar") + v6))
                if v29 {
                    v30 := GenS5{F0: 1, F1: 520, F2: -1}
                    _ = v30
                } else {
                    v23 = (v23 | (^(v23)))
                    v24 += 559
                }
                _ = v29
            }
            _ = v25
        }
    } else {
        var v31 bool = (v6 == ("foo bar" + (v6 + "foo bar")))
        _ = v31
    }
    var v32 float32 = (v2 + ((((v2 + v2) * 0.0) - (0.1 * 3.14159)) / v2))
    var v33 uint = 42
    fmt.Println(v5, v5, v23)
    if len(v4) > 0 {
        v4[0] = 2
    }
    switch v6 {
    case "\n\t":
        var v34 int16 = (^(v24))
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            v2 -= v32
        } else {
            v7 += (v7 ^ 0)
            v36 := GenS1{F0: 4294967295}
            switch fi35 {
            case 16:
                var v37 string = ((v6 + v6) + v6)
                _ = v37
            case 2147483647:
                v32 -= ((3.14159 - ((v2 * (v2 + v32)) / 1.5)) * 1.5)
            default:
                v5 = v24
            }
            _ = v36
        }
        if (v23 != (v23 << ((0) & 63))) {
            var v38 int64 = -9223372036854775808
            _ = v38
        }
        v39 := GenS2{F0: (v6 + "\n\t"), F1: v23}
        _ = fi35
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
