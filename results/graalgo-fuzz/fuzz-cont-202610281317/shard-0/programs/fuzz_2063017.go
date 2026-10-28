package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8) uint {
    if true {
        p0 = (p0 * p0)
        var v7 int16 = 42
        v8 := []int{-508, 619}
        _ = v7
        _ = v8
    }
    p0 = (p0 / ((p0) | 1))
    _ = p0
    return 42
}
func genFn1(p0 string, p1 int32, p2 int) uint16 {
    p1 += -1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 325
}

func main() {
    var v1 bool = (!true)
    var v2 uint16 = 42
    var v3 int32 = 1
    var v4 string = "a"
    var v5 uint32 = 42
    var v6 bool = true
    var v9 int32 = 0
    v10 := GenS1{F0: 240, F1: (v4 + "hello")}
    v11 := GenS1{F0: 105, F1: (v4 + "hello")}
    switch 32768 {
    case 32767:
        switch -852 {
        case 4294967295:
            var v12 int8 = 36
            _ = v12
        }
    }
    var v13 int16 = 0
    v1 = (v1 || v6)
    fmt.Println(v6, v3, v4)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int8 = 48
        _ = fi14
        _ = v15
    }
    switch 947 {
    case -9223372036854775808:
        v16 := []int{127, 2147483647, -9223372036854775808, 9223372036854775807}
        _ = v16
    }
    v17 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
