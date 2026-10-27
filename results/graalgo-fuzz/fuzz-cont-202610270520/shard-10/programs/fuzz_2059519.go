package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 uint8) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (-(fi6)) {
        case -9223372036854775808:
            fmt.Println(p0, fi6)
        case 1:
            if false {
                var v7 int64 = 42
                var v8 int32 = -2147483648
                v7 += -9223372036854775808
                _ = v7
                _ = v8
            }
        default:
            v9 := GenS1{F0: 800}
            _ = v9
        }
        _ = fi6
    }
    var v10 uint8 = (p0 ^ p0)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: -411}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 string = (((("\n\t" + "foo bar") + ("foo bar" + ("a" + ("a" + ("a" + "\n\t"))))) + (("" + "a") + "hello")) + "foo bar")
    v4 := []string{}
    var v5 float64 = (((float64(1.5) + 3.14159) / 1.0) + float64((0.1 + -2.25)))
    fmt.Println(v1, v3)
    v3 = v3
    var v13 float32 = 3.14159
    fmt.Println(v1, v1, v13)
    var v14 int = 42
    switch v3 {
    case "\n\t":
        v2[415] = ("\n\t" + v3)
    case "a":
        var v15 uint64 = 2
        _ = v15
    }
    fmt.Println(v3)
    v16 := GenS2{F0: true, F1: 344, F2: 2}
    v17 := GenS3{F0: 398, F1: 1, F2: (float64(float32(((v5 * 1e10) + 3.14159))) != (1.5 + (-2.25 * v5)))}
    if (!true) {
        v18 := GenS3{F0: 0, F1: -16, F2: (42 >= 1)}
        if true {
            var v19 int16 = (v1 / ((v1) | 1))
            _ = v19
        }
        _ = v18
    } else {
        if true {
            v20 := GenS4{F0: 647, F1: 1}
            if false {
                for fi21 := 0; fi21 < 5; fi21++ {
                    switch fi21 {
                    case 36:
                        var v22 float32 = float32(v5)
                        _ = v22
                    case 1:
                        v5 = float64((3.14159 + v13))
                    default:
                        v5 += v5
                    }
                    v23 := GenS1{F0: 9223372036854775807}
                    fmt.Println(v13, v3)
                    _ = fi21
                    _ = v23
                }
                var v24 float64 = -2.25
                _ = v24
            } else {
                v3 = "hello"
                switch (v3 + (v3 + v3)) {
                case "foo bar":
                    switch (v14 + 2147483647) {
                    case 1:
                        var v25 uint64 = 1
                        _ = v25
                    }
                }
            }
            var v26 int32 = 1
            _ = v20
            _ = v26
        }
        v14 = v14
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 string = (v3 + "\n\t")
            _ = fi27
            _ = v28
        }
    }
    fmt.Println(v13)
    for fi29 := 0; fi29 < 8; fi29++ {
        v14 = (^(fi29))
        var v30 float32 = ((1e10 * v13) - 1e10)
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
