package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8; F1 int }
func genFn0(p0 uint16, p1 string) bool {
    var v6 int32 = 1
    v7 := GenS1{F0: 2, F1: 195, F2: -1}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int32) int8 {
    var v8 uint16 = 2
    p0 = (p0 ^ 670)
    var v9 int8 = -63
    fmt.Println(p0, v8)
    _ = p0
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 bool = (((!false) && true) && true)
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 float64 = 0.0
    v5 := make(map[int]int)
    v4 += float64(3.14159)
    switch 32767 {
    case 9223372036854775807:
        fmt.Println(v4, v4)
    }
    v4 = 1e10
    switch 214 {
    case -2147483648:
        if false {
            if (true || false) {
                var v10 int16 = -1
                _ = v10
            } else {
                if (!true) {
                    var v11 bool = (-1 == -42)
                    for fi12 := 0; fi12 < 17; fi12++ {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 bool = (18446744073709551615 >= 18446744073709551615)
                            _ = fi13
                            _ = v14
                        }
                        _ = fi12
                    }
                    var v15 int8 = -1
                    _ = v11
                    _ = v15
                }
            }
        }
    case 1:
        if v1 {
            v1 = (2 == 214)
            v1 = v1
        }
    case 128:
        var v16 int = -72
        _ = v16
    default:
        fmt.Println(v1)
    }
    if true {
        switch 1 {
        case 32768:
            fmt.Println(v4, v4)
        default:
            var v17 float32 = ((((1.5 / 1.0) * 1.5) * (3.14159 - 0.1)) / 0.1)
            _ = v17
        }
        v1 = v1
    } else {
        v18 := GenS1{F0: 42, F1: 1, F2: 655}
        v19 := GenS2{F0: 100, F1: -1}
        var v20 bool = v1
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int64 = -1
    var v22 float64 = v4
    fmt.Println(v1, v22)
    v23 := []int{}
    v24 := GenS2{F0: 42, F1: -583}
    if (v1 && (3.14159 > 1e10)) {
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 int32 = 768
            switch fi25 {
            case 2147483647:
                var v27 int64 = v21
                _ = v27
            default:
                for fi28 := 0; fi28 < 2; fi28++ {
                    _ = len(v5)
                    _ = fi28
                }
            }
            v22 -= (((((((0.0 + ((float64(3.14159) - ((v4 * (v22 - v4)) * (((v22 - v4) + ((v22 + v4) / v4)) * v4))) * v4)) / (0.0 + 1e10)) + (v22 + float64(0.0))) * (v22 + 0.0)) + 0.0) + v4) * v4)
            _ = fi25
            _ = v26
        }
        v21 = (v21 / (((v21 >> ((0) & 63))) | 1))
    }
    if true {
        if (!(0.0 < 0.1)) {
            fmt.Println(v1)
        }
        var v29 int16 = 1000
        _ = v29
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "hello":
        switch -1 {
        case 0:
            v3[-2147483648] = "a"
        case 9223372036854775807:
            var v30 float32 = (float32(float64((float32(0.1) - (float32(v22) - (0.0 - 3.14159))))) - 0.0)
            _ = v30
        }
    case "\n\t":
        for fi31 := 0; fi31 < 8; fi31++ {
            if ((130 <= 32767) && v1) {
                var v32 bool = (552 == 233)
                v33 := GenS1{F0: 862, F1: 2, F2: -849}
                _ = v32
                _ = v33
            }
            if len(v23) > 0 {
                v23[0] = fi31
            }
            fmt.Println(v21)
            _ = fi31
        }
    default:
        var v34 int64 = 0
        _ = v34
    }
    v35 := GenS2{F0: 127, F1: -46}
    var v36 int16 = 1
    fmt.Println(v21, v36)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
