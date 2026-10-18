package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
func genFn0(p0 int8, p1 uint16, p2 uint32) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch ("a" + "\n\t") {
            case "hello":
                p1 = p1
            case "":
                fmt.Println(p1, p2, fi8)
            default:
                var v9 float64 = (float64(-2.25) * (((-2.25 * ((-2.25 * 3.14159) + 1.5)) / 1.0) / 1.0))
                _ = v9
            }
            v10 := GenS1{F0: "\n\t", F1: 4294967295}
            _ = fi8
            _ = v10
        }
        p2 -= 735
        v11 := []int{-1}
        _ = fi7
        _ = v11
    }
    if false {
        v12 := []int{468}
        p1 += uint16(p2)
        fmt.Println(p1, p2, p2)
        _ = v12
    } else {
        v13 := []int{-2147483648}
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int16 = 42
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint64 {
    if true {
        var v16 bool = (65535 <= 255)
        _ = v16
    }
    v17 := GenS1{F0: "hello", F1: 0}
    var v18 int64 = -1
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 1000
    var v3 float32 = 0.1
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    var v6 float32 = ((3.14159 / (v3 / v3)) - v3)
    v3 = ((-2.25 - v3) - float32(((0.1 * 1.5) * float64(3.14159))))
    v3 += v3
    v3 += v6
    var v19 uint32 = 502
    if true {
        var v20 int32 = -2147483648
        v3 -= ((1e10 + (v6 + v6)) + (v3 / ((v3 / v3) + float32(0.0))))
        var v21 uint32 = v19
        _ = v20
        _ = v21
    }
    var v22 string = "hello"
    v23 := GenS1{F0: "a", F1: 29}
    var v24 int16 = (^(v2))
    var v25 int32 = 1
    if (127 >= -1) {
        var v26 uint = 2
        _ = v26
    } else {
        v22 = v22
        v1 -= v1
        v27 := GenS1{F0: (("hello" + "\n\t") + v22), F1: 342}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
