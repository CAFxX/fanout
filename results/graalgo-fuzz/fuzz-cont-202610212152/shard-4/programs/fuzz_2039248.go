package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 bool) int {
    v7 := GenS1{F0: 4294967295, F1: 42}
    p0 = false
    if p0 {
        fmt.Println(p0)
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 2147483647, F1: 56}
            var v10 int = fi8
            v10 += -1
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = v7
    return 404
}
func genFn1(p0 uint64, p1 uint8) bool {
    fmt.Println(p0, p1)
    v11 := []int{-973, 32768, genFn0(("hello" == ((("hello" + "foo bar") + "a") + "a"))), -2147483648}
    p1 = 75
    _ = p0
    _ = p1
    _ = v11
    return true
}

func main() {
    var v1 int8 = -128
    var v2 bool = ((1.5 / ((float32(float64((0.0 + 1.5))) / (1e10 + (-2.25 - float32(0.1)))) * 1.5)) != -2.25)
    v3 := make(map[int]bool)
    var v4 int16 = 210
    var v5 int16 = (v4 ^ (v4 ^ (v4 / (((v4 + 0)) | 1))))
    var v6 uint8 = 244
    v4 -= (v5 ^ v5)
    var v12 uint32 = 848
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := []int{fi13, -112}
        _ = fi13
        _ = v14
    }
    v15 := []int{genFn0((1 <= -128))}
    v16 := v3[int(v5)]
    switch 1 {
    case 32767:
        v17 := GenS1{F0: (v12 & (v12 & (^(v12)))), F1: 667}
        _ = v17
    }
    if ((v2 && true) && (v16 && v2)) {
        fmt.Println(v16, v1)
        fmt.Println(v2, v16, v2)
        v15 = append(v15, -604)
    }
    var v18 uint8 = 2
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            v18 = 127
            _ = fi20
        }
        _ = fi19
    }
    v3[9223372036854775807] = v16
    v21 := GenS1{F0: 65535, F1: v4}
    fmt.Println(v12, v12, v4)
    v4 += 294
    fmt.Println(v18, v12)
    v22 := GenS1{F0: (v12 >> ((869) & 63)), F1: (v4 / ((v5) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
