package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint16, p1 int32, p2 int) bool {
    v5 := GenS1{F0: 449, F1: 610}
    v6 := []int{-132, (p2 & 2147483647)}
    var v7 int8 = -128
    v8 := GenS1{F0: -32768, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (!(0.1 < float64(-2.25)))
}
func genFn1(p0 string, p1 uint8, p2 uint8) uint64 {
    var v9 int32 = 1
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(fi10 >> ((65535) & 63))}
        fmt.Println(v9, p0, p0)
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 65535
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 81
    var v3 int8 = 104
    var v4 float32 = 3.14159
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint = 2
        v15 := GenS1{F0: 195, F1: int16(v3)}
        var v16 uint32 = 65535
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v1["hello"] = "hello"
    var v17 int = 32767
    var v18 float64 = 0.0
    var v19 string = "foo bar"
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v19)
        fmt.Println(v3, v3)
        fmt.Println(v3, v19, v19)
        _ = fi20
    }
    var v21 int8 = -1
    v22 := GenS2{F0: 113, F1: 801}
    if false {
        v23 := []int{v17, (v17 + -1)}
        fmt.Println(v18, v2, v19)
        var v24 string = ("" + ("foo bar" + v19))
        _ = v23
        _ = v24
    }
    v2 = v2
    switch (v17 * (v17 + -2147483648)) {
    case 181:
        var v25 uint64 = (v2 / ((v2) | 1))
        _ = v25
    case 128:
        if false {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 int32 = 2147483647
                _ = fi26
                _ = v27
            }
        } else {
            if (-9223372036854775808 >= 1) {
                var v28 int64 = 42
                _ = v28
            }
        }
    default:
        var v29 string = "\n\t"
        _ = v29
    }
    fmt.Println(v19, v4, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
