package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
func genFn0() int8 {
    var v6 string = "hello"
    fmt.Println(v6)
    v7 := []int{32767, 554}
    if (42 == 869) {
        v6 = ("hello" + (v6 + v6))
        _ = len(v7)
        var v8 int = 42
        _ = v8
    }
    if true {
        v7 = append(v7, -9223372036854775808)
    }
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 bool = true
    v3 := []int{1, 1}
    var v4 float32 = 1.5
    var v5 int64 = -143
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int16 = 32767
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 100}
    v4 += v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -9223372036854775808 {
    case 787:
        v4 += (0.0 - (1e10 / -2.25))
    case 9223372036854775807:
        if true {
            v12 := GenS1{F0: 0, F1: -1}
            var v13 uint8 = 156
            _ = v12
            _ = v13
        }
    case -2147483648:
        if (0 < -363) {
            var v14 float64 = (-2.25 - ((1e10 + float64((v4 * ((v4 * ((0.0 / float32(1e10)) * (3.14159 - v4))) - float32(0.0))))) * ((3.14159 / (-2.25 - 1e10)) + 0.1)))
            var v15 uint16 = 255
            _ = v14
            _ = v15
        }
    default:
        v4 = 3.14159
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS2{F0: 966, F1: 32767, F2: 2}
    if len(v3) > 0 {
        v3[0] = 1
    }
    if v2 {
        v17 := GenS2{F0: 950, F1: 42, F2: 275}
        v2 = ((-2.25 >= float64(v4)) && v2)
        v18 := GenS2{F0: 384, F1: 980, F2: 42}
        _ = v17
        _ = v18
    } else {
        fmt.Println(v2, v5)
        fmt.Println(v2, v2)
        var v19 int = 128
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = -532
    }
    v4 += (-2.25 - (float32(1.5) - v4))
    var v20 int32 = -2147483648
    if false {
        if ((v5 & 42) == v5) {
            var v21 uint = 42
            _ = v21
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v1, v4, v2)
    switch 32768 {
    case 9223372036854775807:
        v5 += v5
    case 990:
        v3 = append(v3, -9223372036854775808)
    default:
        if false {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 5; fi23++ {
                    fmt.Println(v4, v1)
                    _ = fi23
                }
                _ = fi22
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 uint16 = 255
                _ = fi24
                _ = v25
            }
            v20 -= (v20 * 2147483647)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
