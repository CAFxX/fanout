package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0() uint64 {
    var v7 uint32 = 581
    v8 := []int{-881, 2147483647, -128}
    v7 = v7
    _ = len(v8)
    var v9 uint16 = 255
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 bool, p1 int32) uint8 {
    var v10 float32 = float32((1.5 - 1e10))
    var v11 float64 = 1e10
    v12 := []int{1, 9223372036854775807, 42}
    p0 = true
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 75
}

func main() {
    var v1 int16 = 0
    var v2 uint = 2
    var v3 string = "\n\t"
    var v4 uint = (^(v2))
    var v5 uint32 = 2147483648
    var v6 uint16 = 2
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float64 = 0.1
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: -9223372036854775808, F1: 32768}
    var v16 int32 = 467
    v6 = (v6 - (v6 / ((v6) | 1)))
    v17 := GenS1{F0: 231, F1: -1}
    fmt.Println(v5)
    switch 42 {
    case 32768:
        if true {
            v18 := []int{-9223372036854775808}
            _ = v18
        }
    case -802:
        if (120 <= 0) {
            if false {
                fmt.Println(v6, v1)
                if false {
                    v1 += 804
                }
            }
            v19 := []int{-238, -733}
            var v20 int64 = 629
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v1, v1)
    var v21 int8 = -85
    var v22 int8 = (v21 & (v21 - (v21 ^ (v21 >> ((445) & 63)))))
    fmt.Println(v1, v21)
    for fi23 := 0; fi23 < 20; fi23++ {
        v16 = 2147483647
        v24 := []int{fi23}
        var v25 uint = v2
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
