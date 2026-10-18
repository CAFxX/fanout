package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint = 179
    v2 := make(map[int]string)
    var v3 float32 = ((0.0 + 0.0) - 0.1)
    var v4 int16 = -1
    var v5 float64 = -2.25
    v6 := []int{1, 961, 32768, 9223372036854775807}
    var v7 uint64 = 4294967295
    switch "\n\t" {
    case "\n\t":
        var v8 int64 = -9223372036854775808
        _ = v8
    case "a":
        v9 := GenS1{F0: 325}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 42}
            if true {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v7 = (v7 * uint64(v1))
                    _ = fi12
                }
                if true {
                    switch ((("a" + ("hello" + (("" + "a") + "foo bar"))) + "hello") + (((("" + "foo bar") + "") + (("a" + ("foo bar" + "a")) + "a")) + "hello")) {
                    case "foo bar":
                        if (337 != fi10) {
                            var v13 int = -510
                            fmt.Println(v1, v4, fi10)
                            v4 = (^(v4))
                            _ = v13
                        } else {
                            v5 -= 1.5
                        }
                    case "\n\t":
                        var v14 uint32 = 1
                        _ = v14
                    case "":
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 uint8 = 42
                            if (false || (!((v5 == (v5 * (1.5 + v5))) || true))) {
                                var v17 int16 = 446
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                    default:
                        switch -1 {
                        case 959:
                            var v18 float32 = (3.14159 / (v3 / float32(v5)))
                            _ = v18
                        case -287:
                            var v19 uint16 = 42
                            _ = v19
                        default:
                            v3 -= v3
                        }
                    }
                }
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS1{F0: 198}
                    v22 := GenS2{F0: 1}
                    v23 := GenS3{F0: 2}
                    v7 = uint64(v1)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                if false {
                    var v24 int32 = 2147483647
                    v7 += v7
                    _ = v24
                } else {
                    fmt.Println(v4, v7, v7)
                    if len(v6) > 0 {
                        v6[0] = fi10
                    }
                }
            }
            var v25 int8 = 0
            _ = fi10
            _ = v11
            _ = v25
        }
    }
    var v26 int64 = 9223372036854775807
    v1 += uint(v7)
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(v1, v4, v1)
    case "hello":
        v6 = append(v6, 32767)
    }
    var v27 float64 = float64(float32(((v5 - 0.0) / float64((v3 * float32(v5))))))
    switch "\n\t" {
    case "foo bar":
        var v28 float32 = 3.14159
        _ = v28
    case "\n\t":
        for fi29 := 0; fi29 < 6; fi29++ {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int64 = (v26 | (v26 + v26))
                if len(v6) > 0 {
                    v6[0] = (fi29 & -2147483648)
                }
                _ = fi30
                _ = v31
            }
            v32 := GenS4{F0: 255, F1: v26}
            fmt.Println(fi29, v26)
            _ = fi29
            _ = v32
        }
    }
    fmt.Println(v3, v26, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
