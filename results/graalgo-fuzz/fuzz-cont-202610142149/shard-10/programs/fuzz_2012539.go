package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 uint }
func genFn0(p0 string, p1 int8, p2 uint) int32 {
    switch 57 {
    case 32768:
        var v5 uint16 = 1
        _ = v5
    default:
        fmt.Println(p2)
    }
    p2 = (^(p2))
    v6 := GenS1{F0: 2, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -578
}
func genFn1() uint {
    v7 := []int{195, 2147483647, 128, -1}
    var v8 string = ("" + ("hello" + (("\n\t" + "a") + "\n\t")))
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -128
    v3 := []string{("hello" + ("hello" + ("a" + "\n\t"))), ("hello" + ("foo bar" + "a"))}
    var v4 int16 = -32768
    fmt.Println(v4, v1)
    v9 := GenS2{F0: 2147483647, F1: (^(v4))}
    if true {
        v10 := GenS3{F0: 675}
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 float64 = float64((1e10 + 1e10))
        v13 := GenS1{F0: 2147483647, F1: 2147483648}
        if true {
            v3 = append(v3, (("\n\t" + "hello") + ("hello" + "\n\t")))
        } else {
            v14 := GenS4{F0: 0, F1: int8(v4), F2: 65535}
            _ = v14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v1 -= int64(v2)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v15 float64 = float64(1e10)
    var v16 int8 = (v2 % ((v2) | 1))
    v17 := GenS4{F0: 2147483648, F1: v2, F2: 65535}
    var v18 float64 = (v15 / (((v15 / v15) - float64(1.5)) / 1.0))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
