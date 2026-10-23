package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
func genFn0(p0 uint32) int16 {
    fmt.Println(p0)
    var v6 int64 = -65
    if (!(!(108 <= 0))) {
        var v7 int32 = 42
        v7 -= v7
        _ = v7
    }
    v6 = 0
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 491
    var v2 int16 = 1
    var v3 float64 = -2.25
    var v4 float32 = -2.25
    var v5 int32 = v1
    v3 += ((v3 * v3) - (v3 + float64(v4)))
    var v8 int64 = -1
    switch ((("" + "a") + "foo bar") + "hello") {
    case "foo bar":
        switch "" {
        case "hello":
            switch -849 {
            case 2147483647:
                v1 = (-(v5))
            }
        default:
            v9 := []int{}
            _ = v9
        }
    default:
        v3 += ((1.5 + (1e10 / (3.14159 + v3))) * v3)
    }
    v3 = -2.25
    v4 += 0.1
    v10 := []int{128, 2147483647, 32768, -1}
    var v11 int16 = 32767
    switch 1 {
    case 127:
        fmt.Println(v8, v5)
    case -2147483648:
        if true {
            var v12 bool = ("hello" != ("" + "\n\t"))
            v11 -= 32767
            _ = v12
        } else {
            if (!((v3 / 0.1) >= 0.1)) {
                v8 = (v8 + v8)
            }
        }
    case 2147483647:
        if ((0 < 4294967295) || (v8 <= v8)) {
            if false {
                var v13 uint64 = 286
                v10 = append(v10, -2147483648)
                switch 2147483647 {
                case 127:
                    var v14 uint8 = 42
                    _ = v14
                case 42:
                    if ((v3 * v3) >= (v3 - float64(float32((1e10 * ((0.1 * (v3 - float64(float32((0.0 / v3))))) - 1.5)))))) {
                        if (v2 < v11) {
                            var v15 float64 = v3
                            _ = v15
                        } else {
                            v11 = v11
                            for fi16 := 0; fi16 < 13; fi16++ {
                                v17 := GenS1{F0: -9223372036854775808, F1: false, F2: v2}
                                for fi18 := 0; fi18 < 12; fi18++ {
                                    if (!((0.0 - v3) <= v3)) {
                                        var v19 int16 = int16(v8)
                                        _ = v19
                                    }
                                    v2 = v11
                                    _ = fi18
                                }
                                _ = fi16
                                _ = v17
                            }
                        }
                    }
                case -164:
                    v20 := GenS1{F0: v8, F1: false, F2: (-(v11))}
                    _ = v20
                default:
                    v4 = v4
                }
                _ = v13
            }
            var v21 int64 = v8
            _ = v21
        }
    }
    fmt.Println(v11, v1)
    v11 += v11
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
