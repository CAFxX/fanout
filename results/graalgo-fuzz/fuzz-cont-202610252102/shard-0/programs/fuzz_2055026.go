package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
func genFn0() int16 {
    switch 32768 {
    case 32767:
        var v6 string = "hello"
        _ = v6
    default:
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 bool = true
            fmt.Println(fi7)
            _ = fi7
            _ = v8
        }
    }
    var v9 int16 = 42
    v9 = (v9 ^ v9)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint64 = 846
        _ = fi10
        _ = v11
    }
    _ = v9
    return (v9 * 32767)
}

func main() {
    var v1 float64 = ((float64(1e10) + -2.25) / 0.1)
    v2 := make(map[int]int)
    var v3 string = "hello"
    v4 := []string{v3}
    v5 := []int{-2147483648, 2147483647, 127}
    v1 += float64(1e10)
    fmt.Println(v3)
    v12 := GenS1{F0: v3, F1: false, F2: 0}
    var v13 bool = (false || (!true))
    var v14 float32 = 0.1
    switch -9223372036854775808 {
    case 363:
        var v15 float32 = -2.25
        _ = v15
    case 0:
        var v16 int64 = 9223372036854775807
        _ = v16
    case -590:
        switch (((v3 + (v3 + "")) + v3) + v3) {
        case "a":
            if ((86 != 42) && (278 < 42)) {
                var v17 int64 = 42
                switch 684 {
                case 0:
                    v14 -= (3.14159 / (0.1 + (float32(float64(1.5)) * ((v14 / 3.14159) * v14))))
                case 127:
                    var v18 bool = false
                    _ = v18
                }
                var v19 int = -1
                _ = v17
                _ = v19
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS1{F0: v3, F1: (65535 < 4294967295), F2: -117}
                    _ = fi20
                    _ = v21
                }
            }
        case "foo bar":
            for fi22 := 0; fi22 < 1; fi22++ {
                v14 += 0.0
                _ = fi22
            }
        }
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        switch (((v3 + v3) + (v3 + "")) + v3) {
        case "foo bar":
            v24 := GenS2{F0: 127, F1: 1, F2: v3}
            _ = v24
        default:
            var v25 bool = v13
            _ = v25
        }
        _ = fi23
    }
    var v26 uint16 = 0
    v13 = false
    v14 -= 3.14159
    if (float64(-2.25) <= float64(v14)) {
        v27 := GenS2{F0: 127, F1: 241, F2: v3}
        fmt.Println(v3, v14, v14)
        _ = v27
    }
    v3 = "foo bar"
    var v28 string = "a"
    var v29 int8 = 12
    var v30 int8 = (v29 >> ((404) & 63))
    if v13 {
        var v31 int16 = -32768
        var v32 uint32 = 810
        _ = v31
        _ = v32
    }
    fmt.Println(v29, v29, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
