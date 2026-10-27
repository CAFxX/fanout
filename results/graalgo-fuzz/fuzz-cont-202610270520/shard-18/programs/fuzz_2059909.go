package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint16) int8 {
    p0 = (^(p0))
    v4 := GenS1{F0: 42}
    var v5 int32 = -58
    fmt.Println(p0, p0)
    var v6 bool = ("a" >= "a")
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 uint = 1
    var v7 float32 = ((1.5 + -2.25) * (0.0 - (((0.0 / ((1.5 * (((float32((float64(0.0) + float64(((1e10 + 0.1) - 1e10)))) * (1.5 + 0.0)) + 0.0) / ((0.1 * -2.25) - 3.14159))) + float32((0.1 + float64((0.0 / 0.1)))))) / 0.1) / float32(1e10))))
    var v8 string = "hello"
    fmt.Println(v8)
    switch 127 {
    case -153:
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if (100 == 31) {
                    switch v8 {
                    case "\n\t":
                        var v11 int = (fi10 & 4294967295)
                        _ = v11
                    default:
                        if (!true) {
                            var v12 int8 = 127
                            v7 = ((v7 - 1e10) + (v7 - v7))
                            v1[(v8 + ("a" + v8))] = (-(fi9))
                            _ = v12
                        } else {
                            switch fi9 {
                            case 0:
                                v8 = v8
                            case 9223372036854775807:
                                var v13 string = v8
                                _ = v13
                            default:
                                var v14 int64 = -1
                                _ = v14
                            }
                        }
                        v15 := GenS1{F0: 1}
                        if true {
                            v16 := GenS2{F0: -628, F1: ((v8 + v8) + "foo bar")}
                            if false {
                                v17 := GenS3{F0: v8, F1: "a"}
                                if (fi9 < fi9) {
                                    var v18 float64 = float64(((v7 / -2.25) / v7))
                                    var v19 int32 = -2147483648
                                    v3 = v3
                                    _ = v18
                                    _ = v19
                                }
                                v3 = (v3 << (((v3 >> (((v3 | v3)) & 63))) & 63))
                                _ = v17
                            }
                            _ = len(v2)
                            _ = v16
                        }
                        _ = v15
                    }
                    v7 -= (((((float32(((((-2.25 + (1.5 * (1e10 * (1.5 * 1e10)))) - float64(v7)) * 3.14159) - ((1e10 + (3.14159 * (0.0 - 3.14159))) / (0.1 + -2.25)))) + ((-2.25 - 1e10) - 1.5)) - ((float32(-2.25) + 0.1) * ((1.5 / ((1.5 / v7) * 1e10)) + (v7 / (3.14159 - v7))))) * v7) + 0.1) - 1.5)
                }
                _ = fi10
            }
            fmt.Println(v8, v8)
            _ = fi9
        }
    case -553:
        if (!true) {
            v20 := GenS1{F0: 1}
            v21 := GenS4{F0: 1000}
            if (!((("foo bar" + v8) + ("a" + v8)) == "hello")) {
                v3 += v3
                v3 = v3
            }
            _ = v20
            _ = v21
        }
    default:
        if (genFn0(2) >= 1) {
            if false {
                v3 += (^(v3))
                v22 := GenS3{F0: v8, F1: v8}
                _ = v22
            } else {
                v23 := GenS4{F0: -757}
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 14; fi24++ {
                switch (-(fi24)) {
                case -1:
                    var v25 int32 = -577
                    _ = v25
                }
                v8 = "hello"
                _ = fi24
            }
        }
    }
    fmt.Println(v3)
    var v26 string = ("\n\t" + v8)
    v8 = (v8 + "\n\t")
    var v27 float64 = (0.0 - 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
