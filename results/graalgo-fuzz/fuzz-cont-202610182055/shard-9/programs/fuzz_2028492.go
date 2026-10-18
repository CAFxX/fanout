package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 136
    v3 := make(map[string]int)
    var v4 string = ""
    var v5 float32 = 0.0
    var v6 float32 = v5
    fmt.Println(v1, v5, v4)
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            v8 := GenS1{F0: (^(v1)), F1: (fi7 >> ((1) & 63)), F2: 42}
            _ = v8
        }
        switch "" {
        case "a":
            var v9 float32 = (((float32(-2.25) - 3.14159) + (v5 - 1e10)) / 1.0)
            _ = v9
        case "hello":
            switch 0 {
            case -188:
                v10 := GenS1{F0: (v1 + (v1 / (((v1 + v1)) | 1))), F1: -9223372036854775808, F2: -10}
                _ = v10
            case 2147483647:
                for fi11 := 0; fi11 < 15; fi11++ {
                    v5 = (v5 - v5)
                    var v12 int64 = 9223372036854775807
                    if false {
                        v1 -= 128
                        v12 = v12
                        v13 := []int{fi7, (-(fi11)), (fi7 | fi11), fi7}
                        _ = v13
                    } else {
                        v4 = (("hello" + "hello") + "foo bar")
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v6)
                        case "hello":
                            v14 := GenS2{F0: fi7, F1: -934}
                            _ = v14
                        }
                        var v15 float64 = (float64(v5) * 0.0)
                        var v16 int16 = 42
                        v5 -= float32(float64(float32(float64(float32(v15)))))
                        v5 = v6
                        v12 += 286
                        _ = v15
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                }
            }
        default:
            var v17 string = ((("" + v4) + v4) + v4)
            _ = v17
        }
        _ = fi7
    }
    switch -439 {
    case 4294967295:
        v5 += v5
    case 0:
        v18 := GenS1{F0: 0, F1: 557, F2: 42}
        _ = v18
    }
    v19 := GenS1{F0: 255, F1: 128, F2: -125}
    switch v4 {
    case "a":
        var v20 string = v4
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if (("" + "hello") <= v4) {
            v1 += 0
            fmt.Println(v1, v1, v6)
        }
        v22 := []int{(fi21 / ((0) | 1)), fi21, -855, fi21}
        var v23 uint8 = (v1 ^ (v1 * (v1 | v1)))
        _ = fi21
        _ = v22
        _ = v23
    }
    if ((v2 * 42) <= v2) {
        v4 = v4
    }
    fmt.Println(v1, v4, v4)
    if false {
        for fi24 := 0; fi24 < 9; fi24++ {
            switch 32767 {
            case 9223372036854775807:
                var v25 int8 = -1
                _ = v25
            case -503:
                v4 = v4
            case 32768:
                var v26 int16 = 32767
                _ = v26
            }
            _ = fi24
        }
    } else {
        v27 := GenS2{F0: 4294967295, F1: 369}
        _ = v27
    }
    v4 = v4
    v28 := []int{-2147483648, -614}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
