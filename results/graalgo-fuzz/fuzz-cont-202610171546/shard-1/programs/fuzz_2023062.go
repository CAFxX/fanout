package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
func genFn0() uint32 {
    var v4 int8 = -110
    switch "a" {
    case "a":
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v4 = 127
        }
    default:
        if false {
            for fi5 := 0; fi5 < 2; fi5++ {
                v4 += v4
                _ = fi5
            }
        } else {
            if (true && (0.1 != 0.1)) {
                var v6 uint32 = 2
                _ = v6
            }
        }
    }
    _ = v4
    return 65535
}

func main() {
    v1 := []int{-35, 2147483647, 9223372036854775807}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 1, F1: 32767, F2: -688}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            v10 := GenS1{F0: 297, F1: 0, F2: 1}
            fmt.Println(fi9)
            _ = v10
        }
        v11 := GenS1{F0: 61, F1: 791, F2: 0}
        _ = fi9
        _ = v11
    }
    var v12 uint64 = 42
    var v13 int8 = 0
    v12 -= 872
    var v14 float64 = -2.25
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint = 924
        _ = fi15
        _ = v16
    }
    if ((127 < -2147483648) && (!true)) {
        var v17 string = "a"
        if false {
            var v18 uint16 = 124
            fmt.Println(v18, v13, v18)
            _ = v18
        } else {
            fmt.Println(v14, v12, v12)
            switch ((("\n\t" + "\n\t") + ("a" + ((v17 + "a") + "hello"))) + ((((v17 + v17) + ((("" + v17) + v17) + ("hello" + "\n\t"))) + (((("foo bar" + v17) + v17) + v17) + "a")) + v17)) {
            case "foo bar":
                v14 += (0.1 + (v14 / (float64(0.0) - v14)))
            case "a":
                var v19 uint8 = 1
                _ = v19
            default:
                v14 = (v14 + ((3.14159 - float64(1e10)) + (1.5 + v14)))
            }
            v12 = 42
        }
        fmt.Println(v13, v13)
        _ = v17
    }
    fmt.Println(v12)
    var v20 uint = 42
    v21 := GenS1{F0: 32767, F1: 2, F2: int64(v13)}
    v14 -= float64(float32(float64(0.1)))
    var v22 int32 = 17
    var v23 int32 = (v22 ^ int32(v12))
    var v24 float32 = 0.0
    var v25 float32 = (v24 - v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
