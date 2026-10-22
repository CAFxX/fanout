package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0() int64 {
    var v6 int8 = 0
    fmt.Println(v6)
    v6 -= v6
    _ = v6
    return -1
}
func genFn1() uint64 {
    var v7 float32 = float32(float64((float32(float64(((((float32((0.1 + ((float64(0.0) - (((1e10 + 1.5) * (0.1 - (0.1 / ((0.0 * 1.5) + float64((0.1 / 0.1)))))) / 3.14159)) / 1.0))) / 1e10) + (0.1 - 3.14159)) * ((1e10 + 0.1) * ((3.14159 / 3.14159) * 1e10))) * (-2.25 - (3.14159 * 0.0))))) + 0.1)))
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int64 = genFn0()
        _ = fi8
        _ = v9
    }
    if (!true) {
        v10 := GenS1{F0: 0, F1: 1, F2: 335}
        v7 += 3.14159
        var v11 bool = (2147483647 != 42)
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        switch fi12 {
        case 0:
            if ("\n\t" < "hello") {
                for fi13 := 0; fi13 < 15; fi13++ {
                    fmt.Println(fi12)
                    _ = fi13
                }
            } else {
                var v14 int64 = -1
                if (279 <= 613) {
                    var v15 string = ("foo bar" + "hello")
                    _ = v15
                }
                _ = v14
            }
        case -727:
            v16 := []int{fi12}
            _ = v16
        case 907:
            switch -2147483648 {
            case 522:
                v17 := []int{(^(fi12))}
                _ = v17
            }
        }
        var v18 uint64 = 42
        _ = fi12
        _ = v18
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 float32 = (0.0 * 0.1)
    v2 := []int{406, -411, 0, 9223372036854775807}
    v3 := []int{-446, 4294967295, 9223372036854775807, -435}
    v4 := []int{9223372036854775807, 9223372036854775807, 32768, 32768}
    var v5 uint16 = 42
    v5 = (v5 + v5)
    v3 = append(v3, -2147483648)
    _ = len(v2)
    var v19 uint16 = v5
    var v20 int32 = -178
    switch 851 {
    case 2147483647:
        if false {
            var v21 uint8 = 225
            v22 := GenS1{F0: v20, F1: -9223372036854775808, F2: 1}
            if true {
                switch 32768 {
                case 0:
                    var v23 float32 = (1e10 + 0.0)
                    _ = v23
                case 29:
                    if (false || true) {
                        switch "a" {
                        case "foo bar":
                            _ = len(v2)
                        case "hello":
                            v1 = v1
                        default:
                            var v24 uint64 = 42
                            _ = v24
                        }
                        var v25 int16 = 1000
                        _ = v25
                    }
                }
            } else {
                var v26 float64 = 1e10
                _ = v26
            }
            _ = v21
            _ = v22
        } else {
            v1 -= float32(3.14159)
        }
    default:
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = 651
    }
    var v27 uint64 = 1
    if len(v4) > 0 {
        v4[0] = -937
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
