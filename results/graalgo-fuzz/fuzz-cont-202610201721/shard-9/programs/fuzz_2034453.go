package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
func genFn0(p0 int, p1 string, p2 string) uint16 {
    p0 += p0
    fmt.Println(p0, p1)
    v4 := []int{(p0 + p0), p0, p0, (-(p0))}
    switch 4294967295 {
    case -192:
        if (0 <= 253) {
            var v5 uint = 1
            _ = v5
        }
    case -2147483648:
        p2 = ("\n\t" + "hello")
    case 2147483647:
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint = 0
            _ = fi6
            _ = v7
        }
    }
    p2 = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 943
}
func genFn1() uint32 {
    switch -9223372036854775808 {
    case -407:
        var v8 uint64 = 0
        _ = v8
    default:
        var v9 uint32 = 2147483647
        _ = v9
    }
    var v10 uint = 169
    v10 = (v10 ^ v10)
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 uint32 = 0
            var v13 uint = v10
            v14 := GenS1{F0: 0, F1: v10, F2: fi11}
            _ = v12
            _ = v13
            _ = v14
        }
        v10 = v10
        _ = fi11
    }
    _ = v10
    return 790
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 42
    v3 := make(map[int]int)
    v15 := v1[("\n\t" + "")]
    v16 := []int{}
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            var v18 uint32 = 1
            var v19 int64 = 42
            _ = v18
            _ = v19
        }
        var v20 int8 = v2
        _ = fi17
        _ = v20
    }
    v21 := GenS1{F0: 0, F1: 65535, F2: 128}
    var v22 bool = (!(v15 == ((("hello" + "foo bar") + v15) + v15)))
    v15 = "a"
    var v23 uint32 = genFn1()
    var v24 float32 = ((0.1 - 0.1) * float32(0.0))
    v24 = v24
    fmt.Println(v23, v15)
    if (int(v2) <= 0) {
        v2 -= (v2 % (((-(v2))) | 1))
        var v25 uint8 = 116
        _ = v25
    } else {
        v15 = ("" + ((v15 + v15) + (v15 + v15)))
        var v26 float64 = 1e10
        var v27 uint64 = 0
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
