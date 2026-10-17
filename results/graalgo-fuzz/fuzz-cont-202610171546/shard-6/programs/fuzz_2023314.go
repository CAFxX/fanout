package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 bool; F1 int64; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 741
    var v3 float32 = (0.0 - 1.5)
    v4 := make(map[uint32]string)
    if true {
        var v5 int64 = -9223372036854775808
        _ = v5
    } else {
        fmt.Println(v2, v3, v3)
        var v6 float32 = v3
        if (false && (v2 > (^(v2)))) {
            fmt.Println(v3, v6, v6)
        } else {
            v7 := GenS1{F0: 127, F1: (v1 <= v1)}
            _ = v7
        }
        _ = v6
    }
    switch 4294967295 {
    case 0:
        if true {
            var v8 float32 = (v3 / v3)
            _ = v8
        }
    default:
        v2 += v2
    }
    if (65535 != 255) {
        fmt.Println(v1)
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 float64 = 3.14159
            if true {
                v10 -= ((v10 + float64(float32((v10 + float64(v3))))) * v10)
                if false {
                    var v11 uint = 1
                    _ = v11
                }
            } else {
                if true {
                    switch "" {
                    case "":
                        v12 := GenS2{F0: -1, F1: 120, F2: 0}
                        _ = v12
                    default:
                        var v13 float32 = ((v3 + 1.5) / v3)
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS2{F0: -1, F1: 236, F2: 32767}
                        _ = fi14
                        _ = v15
                    }
                } else {
                    v10 = v10
                    if false {
                        if (false || true) {
                            var v16 bool = false
                            v17 := []int{}
                            for fi18 := 0; fi18 < 17; fi18++ {
                                var v19 bool = false
                                v20 := GenS3{F0: "hello", F1: -1, F2: "a"}
                                v4[2147483647] = ""
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = v16
                            _ = v17
                        }
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        var v21 int16 = 0
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        switch fi22 {
        case -695:
            v23 := GenS1{F0: 32768, F1: (false && true)}
            _ = v23
        case 1:
            var v24 float64 = 0.1
            _ = v24
        default:
            fmt.Println(v2, fi22)
        }
        var v25 float32 = v3
        var v26 int16 = 1
        _ = fi22
        _ = v25
        _ = v26
    }
    v27 := []int{55}
    v28 := GenS2{F0: 76, F1: v1, F2: uint16(v1)}
    v2 -= v2
    var v29 int8 = 48
    if (v29 <= -58) {
        if (!(true || (9223372036854775807 <= -1))) {
            switch -2147483648 {
            case 662:
                delete(v4, 2)
            case 1:
                switch "" {
                case "hello":
                    fmt.Println(v1, v3)
                case "":
                    if (!false) {
                        for fi30 := 0; fi30 < 18; fi30++ {
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            _ = fi30
                        }
                        var v31 int = -9223372036854775808
                        _ = v31
                    }
                }
            case 127:
                if true {
                    fmt.Println(v2, v29, v1)
                }
            default:
                var v32 uint64 = 18446744073709551615
                _ = v32
            }
            v2 = (v2 << ((4294967295) & 63))
        }
        v29 = (v29 / ((127) | 1))
        switch "\n\t" {
        case "":
            for fi33 := 0; fi33 < 4; fi33++ {
                if false {
                    v34 := v4[2147483647]
                    for fi35 := 0; fi35 < 1; fi35++ {
                        var v36 string = ("a" + v34)
                        var v37 float32 = (v3 - (v3 - 1.5))
                        v38 := GenS1{F0: (fi33 & fi35), F1: ((((1.5 * (3.14159 / 3.14159)) * (float64(v37) + -2.25)) + -2.25) <= (3.14159 + -2.25))}
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    v3 -= v3
                    _ = v34
                }
                v29 = v29
                if (-244 > 1) {
                    var v39 int32 = 0
                    v29 -= v29
                    var v40 uint = 4294967295
                    _ = v39
                    _ = v40
                }
                _ = fi33
            }
        case "a":
            for fi41 := 0; fi41 < 2; fi41++ {
                v42 := GenS4{F0: false, F1: -9223372036854775808, F2: "a"}
                if true {
                    for fi43 := 0; fi43 < 4; fi43++ {
                        fmt.Println(v1, v1)
                        v4[2147483647] = "foo bar"
                        var v44 bool = true
                        _ = fi43
                        _ = v44
                    }
                }
                _ = len(v4)
                _ = fi41
                _ = v42
            }
        case "hello":
            fmt.Println(v2)
        }
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        fmt.Println(v2, v2, v3)
        _ = fi45
    }
    var v46 uint32 = uint32(v2)
    switch "a" {
    case "\n\t":
        v1 = (^(v1))
    case "a":
        var v47 int16 = 0
        _ = v47
    }
    v29 = -1
    v3 += (float32(1e10) / v3)
    fmt.Println(v46, v46, v29)
    var v48 float32 = 0.1
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v46
    _ = v48
}
