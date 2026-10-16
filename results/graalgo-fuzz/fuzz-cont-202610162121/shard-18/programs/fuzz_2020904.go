package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 int64, p1 uint16) int64 {
    v7 := []int{32767, 128}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    var v4 int32 = 197
    var v5 float64 = 0.0
    v6 := make(map[int]string)
    switch "\n\t" {
    case "":
        v4 += v4
    case "foo bar":
        switch 4294967295 {
        case -656:
            switch 804 {
            case 32768:
                var v8 uint32 = 42
                _ = v8
            case 1:
                if true {
                    v3 -= (1e10 - float32(float64((v3 / (float32(v5) + v3)))))
                    v4 -= v4
                }
            case 974:
                for fi9 := 0; fi9 < 20; fi9++ {
                    if (1 <= 4294967295) {
                        var v10 int32 = v4
                        _ = v10
                    } else {
                        v5 -= (v5 / (((v5 * v5) / (1e10 * v5)) + 0.0))
                        v11 := GenS1{F0: 42, F1: 32767, F2: 9223372036854775807}
                        var v12 uint = 4294967295
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
            default:
                v13 := GenS1{F0: 139, F1: 0, F2: -9223372036854775808}
                _ = v13
            }
        case 4294967295:
            for fi14 := 0; fi14 < 16; fi14++ {
                if true {
                    var v15 int8 = -65
                    delete(v2, 0)
                    var v16 int64 = -48
                    _ = v15
                    _ = v16
                } else {
                    v4 += (^(v4))
                    v3 = float32(v5)
                }
                switch "a" {
                case "\n\t":
                    v17 := []int{4294967295, fi14, fi14}
                    _ = v17
                case "a":
                    var v18 int32 = (v4 / (((-(v4))) | 1))
                    _ = v18
                }
                v3 += (float32(0.1) + (3.14159 - float32(((1.5 + float64(0.1)) - v5))))
                _ = fi14
            }
        }
    case "\n\t":
        v19 := v1["hello"]
        _ = v19
    }
    delete(v2, -647)
    v3 = ((v3 + ((v3 * v3) + ((1e10 - -2.25) + ((float32(v5) / (0.0 + (v3 + v3))) - (v3 / (v3 * (float32(v5) - float32(v5)))))))) / 1.5)
    v4 -= (v4 & (-(v4)))
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int32 = 42
        _ = fi20
        _ = v21
    }
    v22 := []int{-2147483648}
    if true {
        for fi23 := 0; fi23 < 17; fi23++ {
            v5 += ((((float64(((float32(((v5 * ((-2.25 - v5) / ((3.14159 + v5) / 1.5))) * (v5 + v5))) - v3) * ((v3 - v3) * (v3 / 1.0)))) / 1.0) - ((0.0 / (v5 / v5)) * v5)) / 3.14159) + v5)
            v24 := GenS1{F0: 127, F1: 32767, F2: 1}
            var v25 uint8 = 79
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint = 5
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        _ = fi28
    }
    v29 := GenS1{F0: 191, F1: 581, F2: 9223372036854775807}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v30 := GenS2{F0: ("" + ("\n\t" + "\n\t")), F1: 0, F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
