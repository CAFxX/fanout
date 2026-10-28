package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
func genFn0(p0 string, p1 int8) int {
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            p0 = "hello"
            v7 := []int{-2147483648}
            p0 = p0
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 0, F1: 268, F2: "a"}
        var v9 uint = 42
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 bool = false
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 255, F1: 706, F2: p0}
        var v13 bool = false
        for fi14 := 0; fi14 < 20; fi14++ {
            v13 = v13
            var v15 uint16 = 460
            v16 := GenS2{F0: 127, F1: 255, F2: 34}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    return 32768
}

func main() {
    var v1 bool = (!false)
    var v2 uint16 = 42
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    if true {
        v17 := []int{}
        var v18 int = genFn0("hello", -128)
        v3 += ((1.5 - v3) * v3)
        _ = v17
        _ = v18
    }
    if v1 {
        switch 42 {
        case 4294967295:
            v1 = false
        default:
            var v19 uint64 = 4294967295
            _ = v19
        }
        var v20 uint = 65535
        _ = v20
    } else {
        fmt.Println(v2, v3, v1)
        var v21 bool = true
        v2 = (^(v2))
        _ = v21
    }
    if (4294967295 > 2147483648) {
        v22 := []int{-9223372036854775808}
        v3 = 0.1
        _ = v22
    }
    v3 -= v3
    v3 = 0.1
    var v23 uint = 4294967295
    v24 := GenS3{F0: v23, F1: 42, F2: 42}
    var v25 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
