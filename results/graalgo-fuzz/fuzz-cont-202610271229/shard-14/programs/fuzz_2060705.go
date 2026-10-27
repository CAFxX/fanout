package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 bool = false
    var v2 uint32 = 0
    v3 := []string{"", "\n\t", ("hello" + ("" + ("hello" + (("foo bar" + "a") + "\n\t"))))}
    v4 := make(map[string]bool)
    var v5 float32 = ((1.5 + ((1.5 + (3.14159 * 3.14159)) * (0.1 + ((1e10 + 3.14159) - ((((0.1 + (((1e10 + 0.0) + -2.25) + (3.14159 * 3.14159))) * (float32(-2.25) + (-2.25 - float32(((3.14159 - 0.0) / float64(0.1)))))) * float32(((((1e10 / 1.0) / 1.0) + 0.1) / 0.1))) + (1.5 * (((((3.14159 + 1.5) - (0.1 + -2.25)) - 0.0) * float32((float64(-2.25) + (0.1 + 0.1)))) - 3.14159))))))) * 1e10)
    v6 := make(map[uint32]uint32)
    v5 += v5
    fmt.Println(v5, v1)
    var v7 int = 1
    v7 += v7
    v5 = v5
    var v8 uint8 = 0
    switch (v7 & v7) {
    case 1:
        fmt.Println(v8, v8)
    case 127:
        fmt.Println(v7, v1)
    default:
        var v9 int = 9223372036854775807
        _ = v9
    }
    var v10 int16 = int16(v7)
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint8 = (^(v8))
            switch ("hello" + ("\n\t" + ("hello" + "foo bar"))) {
            case "":
                if (v1 || true) {
                    v14 := v6[(v2 | v2)]
                    switch fi11 {
                    case -1:
                        var v15 float32 = 0.1
                        _ = v15
                    default:
                        if ((v13 / ((v8) | 1)) > v13) {
                            _ = len(v4)
                            switch (("foo bar" + "a") + "foo bar") {
                            case "":
                                var v16 int32 = -1
                                _ = v16
                            case "hello":
                                switch (-(fi11)) {
                                case 135:
                                    v17 := GenS1{F0: "\n\t", F1: 100}
                                    _ = v17
                                case -9223372036854775808:
                                    var v18 uint8 = (v13 % ((127) | 1))
                                    _ = v18
                                default:
                                    var v19 string = "a"
                                    _ = v19
                                }
                            case "\n\t":
                                v1 = v1
                            default:
                                if (!((4294967295 < 2) && (true || false))) {
                                    var v20 int8 = 100
                                    if ((v10 << ((0) & 63)) > -993) {
                                        _ = len(v4)
                                    } else {
                                        var v21 float64 = (-2.25 / 3.14159)
                                        v13 -= v8
                                        _ = v21
                                    }
                                    fmt.Println(v1)
                                    _ = v20
                                } else {
                                    var v22 uint8 = (v8 | 128)
                                    var v23 int32 = 1
                                    v23 = 807
                                    _ = v22
                                    _ = v23
                                }
                            }
                        } else {
                            if ((v8 >= 1) && v1) {
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                        }
                    }
                    v24 := GenS2{F0: (v2 ^ 4294967295), F1: (255 <= 255)}
                    _ = v14
                    _ = v24
                }
            case "a":
                v25 := GenS2{F0: v2, F1: false}
                _ = v25
            default:
                v26 := GenS1{F0: (("hello" + "hello") + ""), F1: -1}
                _ = v26
            }
            v13 = (^(v8))
            v27 := GenS1{F0: "", F1: -16}
            _ = len(v6)
            _ = fi12
            _ = v13
            _ = v27
        }
        _ = fi11
    }
    fmt.Println(v2, v1, v10)
    v28 := GenS1{F0: "foo bar", F1: 30}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
