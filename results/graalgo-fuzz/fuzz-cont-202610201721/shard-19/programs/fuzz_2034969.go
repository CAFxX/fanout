package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint64 {
    var v5 uint = 235
    v6 := []int{-2147483648, 0, -2147483648, 458}
    v5 -= v5
    var v7 uint64 = 42
    fmt.Println(v7, v7)
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 uint16) int32 {
    v8 := GenS1{F0: false}
    var v9 float64 = float64((-2.25 - float32(1.5)))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 0
    var v3 int64 = -9223372036854775808
    var v4 int32 = -2147483648
    v3 += v3
    v1 = v1
    v1 = ("" + v1)
    fmt.Println(v1)
    var v10 uint = 752
    var v11 float32 = -2.25
    var v12 uint8 = 42
    switch -1 {
    case 32768:
        var v13 uint16 = 2
        _ = v13
    case 1:
        var v14 int16 = 1
        _ = v14
    }
    var v15 int64 = 42
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint16 = 255
        v18 := []int{fi16}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
