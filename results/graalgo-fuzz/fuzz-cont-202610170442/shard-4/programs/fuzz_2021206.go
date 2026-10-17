package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int32, p1 int16, p2 int16) int32 {
    if true {
        var v6 uint8 = 119
        var v7 int8 = 114
        var v8 uint = 4294967295
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        p2 = 42
        _ = fi9
    }
    if (214 <= 9223372036854775807) {
        var v10 uint = 4294967295
        var v11 uint8 = uint8(v10)
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 890
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -404
    v3 := make(map[string]uint32)
    var v4 string = ("\n\t" + "foo bar")
    var v5 int32 = -195
    switch (v2 / ((-1) | 1)) {
    case 906:
        var v12 uint16 = 2
        _ = v12
    case 128:
        v5 = v5
    case 955:
        v13 := GenS1{F0: v4}
        _ = v13
    }
    v4 = "\n\t"
    var v14 int16 = -520
    fmt.Println(v4, v2)
    fmt.Println(v2, v14, v2)
    v15 := v3[((("foo bar" + v4) + ("hello" + ((v4 + "hello") + "foo bar"))) + "\n\t")]
    v16 := []int{v2, (^(v2)), (v2 / ((v2) | 1)), v2}
    v14 -= (v14 - v14)
    if false {
        v2 = v2
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int16 = v14
            v14 = v18
            _ = fi17
            _ = v18
        }
        var v19 int64 = 1
        _ = v19
    } else {
        if (true || true) {
            var v20 string = v4
            var v21 uint8 = 2
            switch (v2 * (v2 & 128)) {
            case 443:
                v22 := GenS2{F0: (v5 ^ genFn0((v5 >> ((4294967295) & 63)), 333, (v14 / ((v14) | 1)))), F1: int8(v5)}
                _ = v22
            case -910:
                v23 := GenS3{F0: v21, F1: 350}
                _ = v23
            default:
                v15 += v15
            }
            _ = v20
            _ = v21
        }
        v5 = v5
        switch v2 {
        case 127:
            v2 = v2
        case 2147483647:
            if (!(!false)) {
                switch 0 {
                case -869:
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v15 += (v15 << ((968) & 63))
                        var v25 float64 = 1.5
                        _ = fi24
                        _ = v25
                    }
                default:
                    fmt.Println(v2, v14, v5)
                }
                v4 = "a"
                var v26 bool = false
                _ = v26
            } else {
                v4 = "foo bar"
                var v27 bool = true
                _ = v27
            }
        case 281:
            var v28 float32 = 0.1
            _ = v28
        default:
            var v29 float32 = (((-2.25 - ((0.1 + 3.14159) / 1.0)) * 0.0) / 1e10)
            _ = v29
        }
    }
    fmt.Println(v2, v14, v14)
    switch ((((v4 + v4) + "a") + v4) + (v4 + v4)) {
    case "foo bar":
        fmt.Println(v4)
    }
    _ = len(v16)
    var v30 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
