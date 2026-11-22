package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
func genFn0() int16 {
    var v7 uint16 = 2
    var v8 float64 = (((float64((1e10 / 1e10)) - 1.5) * (1.5 + (float64((1.5 - 0.0)) - (((1e10 * 3.14159) / 1.5) + 0.0)))) / (0.1 / 1.0))
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint8, p1 uint, p2 uint32) bool {
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{fi9, -9223372036854775808}
        _ = fi9
        _ = v10
    }
    if (!true) {
        v11 := []int{628, 128, 4294967295, -9223372036854775808}
        fmt.Println(p2)
        _ = v11
    }
    var v12 bool = (!(-2147483648 > 0))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return v12
}

func main() {
    var v1 string = ""
    var v2 int64 = -9223372036854775808
    var v3 uint = 0
    var v4 string = (v1 + v1)
    v5 := []int{32767, 0}
    var v6 uint8 = 1
    if false {
        var v13 int = -1
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int8 = int8(v6)
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v16 float64 = 0.1
    var v17 uint64 = 697
    fmt.Println(v6, v2)
    v18 := GenS1{F0: v6, F1: (v6 + (v6 - (v6 / ((2) | 1)))), F2: 164}
    v6 -= v6
    v2 = v2
    fmt.Println(v17, v6, v6)
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS1{F0: (v6 % (((v6 / (((^(v6))) | 1))) | 1)), F1: v6, F2: -123}
        _ = len(v5)
        _ = fi19
        _ = v20
    }
    v16 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
