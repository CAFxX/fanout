package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int; F1 string; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 int16 = 871
    v4 := []uint32{42, 0}
    v2 = "hello"
    var v5 float32 = (0.0 - float32((1.5 + 0.1)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v2
    v2 = "foo bar"
    if false {
        v3 = v3
        for fi6 := 0; fi6 < 1; fi6++ {
            v4 = append(v4, 2147483647)
            var v7 int8 = -104
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: 100, F1: 2}
    fmt.Println(v5)
    var v9 float32 = 1e10
    if (true || true) {
        switch v2 {
        case "":
            if ((0 <= 0) || ((((!true) && ((false && (true && false)) && false)) || ((^(v3)) == v3)) && (false && false))) {
                _ = len(v4)
            }
        default:
            for fi10 := 0; fi10 < 17; fi10++ {
                if false {
                    if false {
                        v11 := GenS2{F0: 65535, F1: 2147483647}
                        _ = v11
                    } else {
                        v12 := GenS1{F0: -128, F1: 32767}
                        switch "\n\t" {
                        case "":
                            fmt.Println(v9, v9)
                        default:
                            fmt.Println(v5)
                        }
                        _ = v12
                    }
                    v5 = -2.25
                } else {
                    var v13 float32 = v5
                    var v14 int = (fi10 ^ fi10)
                    if (float32((3.14159 / -2.25)) < v9) {
                        v1[128] = ((("hello" + "a") + (v2 + v2)) + (((v2 + v2) + "a") + v2))
                        var v15 float64 = float64((v5 * (float32((1.5 + -2.25)) - v13)))
                        for fi16 := 0; fi16 < 5; fi16++ {
                            var v17 uint16 = 1
                            v18 := GenS3{F0: fi10, F1: ((((v2 + "foo bar") + v2) + (v2 + v2)) + v2), F2: 5}
                            var v19 int16 = v3
                            delete(v1, v14)
                            switch v14 {
                            case 128:
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                            default:
                                var v20 int = (v14 + 32768)
                                _ = v20
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = v15
                    }
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch 9223372036854775807 {
    case -2147483648:
        v9 -= v9
    }
    v9 += 0.1
    v5 = (v5 * ((v9 / 3.14159) * (v9 * v5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
