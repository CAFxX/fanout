package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 uint16) uint16 {
    v5 := []int{32767, 2147483647, 128}
    var v6 uint16 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 uint8) uint {
    v7 := []int{-869, -45, 0}
    _ = len(v7)
    _ = len(v7)
    var v8 int = -2147483648
    v9 := GenS1{F0: (-(v8))}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = (("\n\t" + ("" + ("" + "\n\t"))) + "hello")
    v3 := make(map[int]bool)
    var v4 uint = 0
    v10 := GenS1{F0: 32767}
    var v11 uint8 = 128
    switch "\n\t" {
    case "\n\t":
        v12 := GenS1{F0: -2147483648}
        _ = v12
    default:
        var v13 int32 = 2147483647
        _ = v13
    }
    v2 = (v2 + "")
    v2 = v2
    v14 := []int{-83}
    v15 := GenS1{F0: 4294967295}
    v16 := GenS1{F0: 244}
    var v17 float64 = ((1.5 / 1.0) / (-2.25 * 0.1))
    var v18 float32 = 0.1
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int32 = 0
        v17 -= ((((1.5 * -2.25) * v17) + (float64(v18) + (v17 - (v17 + 0.0)))) * (float64((((v18 * 0.0) + v18) * v18)) * (0.1 + (((3.14159 - (v17 - (((v17 + v17) + v17) - v17))) + v17) + 3.14159))))
        _ = fi19
        _ = v20
    }
    delete(v3, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
