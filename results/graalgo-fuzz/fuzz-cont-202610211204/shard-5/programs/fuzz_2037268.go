package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 uint8 }
func genFn0() int8 {
    var v7 int16 = 0
    var v8 uint64 = 42
    var v9 float32 = 3.14159
    _ = v7
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 uint = 323
    v2 := make(map[string]bool)
    v3 := []string{"foo bar"}
    var v4 float32 = 1e10
    var v5 uint8 = 104
    v6 := []uint32{2, uint32(v5), 611}
    v3 = append(v3, "\n\t")
    v10 := GenS1{F0: 2147483648, F1: int16(v5)}
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint16 = 727
        delete(v2, ("" + ("hello" + "hello")))
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS2{F0: 2147483648}
            switch fi11 {
            case 9223372036854775807:
                switch ("" + "hello") {
                case "hello":
                    var v15 int8 = 1
                    _ = v15
                case "a":
                    fmt.Println(fi13, v12)
                }
            case 42:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 571:
                switch "" {
                case "\n\t":
                    v12 += (^(v12))
                case "a":
                    var v16 uint32 = 2147483647
                    _ = v16
                default:
                    var v17 uint8 = (v5 / (((v5 / ((77) | 1))) | 1))
                    _ = v17
                }
            default:
                var v18 uint32 = 4294967295
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v1 = 241
    if (0 == 183) {
        if (true && (-32768 <= -32768)) {
            if (float64(((0.1 * (((((v4 / 0.1) * (((3.14159 - float32(3.14159)) * v4) + float32(-2.25))) / (float32(3.14159) - -2.25)) - float32((3.14159 - 1e10))) - 0.0)) + v4)) < 1.5) {
                var v19 uint16 = 732
                var v20 int32 = 0
                _ = v19
                _ = v20
            }
        }
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        var v21 uint64 = 0
        _ = v21
    } else {
        if len(v6) > 0 {
            v6[0] = 977
        }
        switch "\n\t" {
        case "foo bar":
            var v22 uint64 = 0
            _ = v22
        case "hello":
            for fi23 := 0; fi23 < 12; fi23++ {
                v4 += ((-2.25 / v4) * v4)
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 int64 = -9223372036854775808
                    _ = fi24
                    _ = v25
                }
                switch -2147483648 {
                case 788:
                    v26 := GenS3{F0: 916}
                    _ = v26
                case 127:
                    switch "a" {
                    case "\n\t":
                        if ((65535 > uint16(v1)) || true) {
                            v4 = ((0.0 * v4) * float32(((-2.25 * 0.1) * 0.0)))
                            v27 := GenS4{F0: (v1 * v1)}
                            _ = v27
                        }
                    case "":
                        if false {
                            if len(v6) > 0 {
                                v6[0] = 2147483647
                            }
                        }
                    default:
                        v28 := GenS5{F0: 255}
                        _ = v28
                    }
                }
                _ = fi23
            }
        default:
            fmt.Println(v1, v5, v1)
        }
        var v29 int16 = -32768
        _ = v29
    }
    delete(v2, "foo bar")
    v5 -= v5
    var v30 float64 = (0.1 - -2.25)
    fmt.Println(v4, v5)
    v1 -= v1
    switch 128 {
    case -1:
        for fi31 := 0; fi31 < 20; fi31++ {
            if (!(180 <= -673)) {
                var v32 float32 = v4
                var v33 float32 = 1e10
                if (true && (true && (213 < 85))) {
                    var v34 float32 = float32(3.14159)
                    _ = v34
                } else {
                    switch (fi31 ^ (fi31 - (fi31 ^ (fi31 % ((fi31) | 1))))) {
                    case 32767:
                        if (4294967295 == 4294967295) {
                            v1 -= v1
                        } else {
                            v5 = (v5 | (^(v5)))
                            v30 += (0.0 + 1.5)
                        }
                    }
                }
                _ = v32
                _ = v33
            } else {
                v35 := GenS4{F0: v1}
                _ = v35
            }
            var v36 uint = 255
            fmt.Println(v1, v5, v4)
            _ = fi31
            _ = v36
        }
    case 9223372036854775807:
        v37 := GenS1{F0: 65535, F1: 42}
        _ = v37
    }
    switch "a" {
    case "\n\t":
        if ((v5 / ((v5) | 1)) >= v5) {
            fmt.Println(v30, v1, v4)
            var v38 uint = v1
            v39 := GenS6{F0: 979, F1: (v5 - v5)}
            _ = v38
            _ = v39
        }
    case "":
        fmt.Println(v4, v30)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
