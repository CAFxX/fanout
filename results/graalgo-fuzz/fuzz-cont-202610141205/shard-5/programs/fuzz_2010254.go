package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 uint) int16 {
    v4 := []int{128, 0}
    if len(v4) > 0 {
        v4[0] = -1
    }
    _ = p0
    _ = v4
    return -444
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 987
    var v3 int16 = -32768
    if true {
        v5 := []int{1}
        v6 := GenS1{F0: 644, F1: 42, F2: genFn0(18446744073709551615)}
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = 42
            _ = len(v5)
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        v1 = (v1 ^ v1)
    }
    v3 -= (v3 + 1)
    var v9 int16 = genFn0(255)
    v2 -= (v2 >> ((18446744073709551615) & 63))
    fmt.Println(v2, v2)
    var v10 int16 = (v9 >> ((255) & 63))
    v3 -= v10
    v11 := []int{-1}
    if len(v11) > 0 {
        v11[0] = int(v1)
    }
    if false {
        v11 = append(v11, 42)
        v2 = v2
        var v12 int = 4294967295
        _ = v12
    }
    var v13 float64 = 0.0
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v1 = v1
    v14 := GenS1{F0: 695, F1: genFn0(18446744073709551615), F2: (v9 >> ((754) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
