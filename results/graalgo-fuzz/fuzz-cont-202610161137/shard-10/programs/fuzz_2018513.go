package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 int8) int32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = 2
            _ = fi8
            _ = v9
        }
        p2 -= -1
        p2 += (^(p2))
        _ = fi7
    }
    if true {
        v10 := GenS1{F0: 1, F1: -109, F2: true}
        _ = v10
    } else {
        var v11 uint32 = 84
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := []int{870}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 string = "a"
    var v5 int64 = -9223372036854775808
    var v6 float32 = 3.14159
    var v12 string = (((v4 + "foo bar") + "") + "a")
    var v13 int = 1
    v14 := GenS1{F0: genFn0(18, 42, -1), F1: 42, F2: (("foo bar" + (v4 + "")) != (v4 + (v12 + (v12 + v12))))}
    var v15 int8 = -128
    switch v4 {
    case "":
        v16 := GenS2{F0: 1, F1: 0, F2: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float64 = ((-2.25 * 0.0) * ((float64(v6) + ((0.1 / (((((0.1 / (((float64(v6) * (((0.1 * 1.5) * (1.5 / ((-2.25 + 0.0) / 1.5))) - 1.5)) / 1.0) * (-2.25 * 0.1))) - 0.0) + 1.5) + 0.1) + ((-2.25 - (-2.25 + 3.14159)) - 0.0))) + -2.25)) - 1e10))
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 int16 = 32767
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v21 := GenS2{F0: 58, F1: v5, F2: 1}
    v22 := GenS3{F0: 127}
    if (true || false) {
        var v23 uint16 = 65535
        _ = v23
    }
    _ = len(v1)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int64 = (v5 & (v5 % (((-(v5))) | 1)))
        v26 := GenS3{F0: 42}
        fmt.Println(v12, v25, v12)
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS3{F0: 1}
    var v28 uint64 = 1
    var v29 float32 = v6
    if (true || false) {
        if false {
            var v30 int64 = -1
            v31 := GenS2{F0: 202, F1: (v30 / ((42) | 1)), F2: int16(v5)}
            _ = v30
            _ = v31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
