package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 int8) bool {
    v5 := GenS1{F0: 2, F1: 851}
    if (1.5 <= 3.14159) {
        var v6 bool = (9223372036854775807 == -508)
        v6 = (2147483647 != 9223372036854775807)
        v7 := GenS1{F0: 42, F1: 967}
        _ = v6
        _ = v7
    }
    var v8 bool = (988 != 65535)
    if v8 {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return true
}
func genFn1(p0 bool, p1 int16, p2 uint8) uint16 {
    switch 2147483647 {
    case 81:
        v9 := []int{-820, -691}
        _ = v9
    case 128:
        v10 := []int{1, 2147483647, 1, -687}
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 968
    var v3 bool = (32767 < -32768)
    var v4 string = "hello"
    var v13 uint8 = 0
    v14 := GenS2{F0: 0, F1: 408, F2: 100}
    var v15 int = 561
    fmt.Println(v13)
    if (!((v13 > (v13 + v13)) || false)) {
        if (!false) {
            _ = len(v1)
        } else {
            var v16 int8 = 39
            var v17 uint = 1
            var v18 string = v4
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 int64 = (v2 + -42)
        var v20 uint32 = 585
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS1{F0: 65535, F1: 65535}
            v23 := []int{v15, fi21, fi21}
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS2{F0: 18446744073709551615, F1: 42, F2: -16}
        _ = v24
    }
    v1[v4] = 178
    v15 = v15
    var v25 float32 = (0.0 / ((1.5 - -2.25) * float32((-2.25 * (-2.25 * float64(3.14159))))))
    var v26 int8 = -101
    fmt.Println(v4)
    v1["hello"] = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
