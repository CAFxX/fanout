package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0() int {
    v4 := GenS1{F0: 1, F1: "a", F2: 1}
    var v5 uint32 = 2147483648
    v5 = v5
    v5 = v5
    fmt.Println(v5)
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1(p0 uint64, p1 uint16) int8 {
    var v6 uint8 = 2
    p1 = 32767
    if (true || false) {
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 = uint8(p0)
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: -458, F1: "", F2: 42}
            _ = fi9
            _ = v10
        }
        fmt.Println(p1, p1, p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    return 108
}

func main() {
    var v1 bool = (65535 == 40)
    var v2 float32 = (1.5 * 0.1)
    var v3 uint64 = 0
    v1 = v1
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint32 = 168
        _ = fi11
        _ = v12
    }
    v13 := []int{-480, 494, -1}
    var v14 float32 = (v2 + (v2 - (v2 + 0.1)))
    v14 -= v14
    v15 := GenS1{F0: -32768, F1: "hello", F2: 2}
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v1)
            v2 = 3.14159
            _ = fi16
        }
        v17 := GenS2{F0: genFn0(), F1: 42}
        var v18 float32 = v2
        _ = v17
        _ = v18
    }
    fmt.Println(v14, v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
