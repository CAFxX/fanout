package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8) int16 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = ("" + "a")
        v6 = v6
        _ = fi5
        _ = v6
    }
    p0 += p0
    p0 -= p0
    _ = p0
    return 1
}
func genFn1(p0 string, p1 string, p2 uint32) int32 {
    var v7 uint = 1
    p0 = p0
    if false {
        v8 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v8
    }
    switch -320 {
    case 32767:
        var v9 uint64 = 924
        _ = v9
    case 32768:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint32 = (p2 * p2)
            fmt.Println(v11)
            v12 := []int{fi10}
            _ = fi10
            _ = v11
            _ = v12
        }
    case 2147483647:
        v13 := GenS2{F0: 9223372036854775807}
        _ = v13
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p2)
    case "foo bar":
        var v14 uint32 = (p2 & (p2 * (p2 / ((p2) | 1))))
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 0
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    fmt.Println(v3, v3)
    var v15 int64 = 1
    fmt.Println(v15)
    var v16 float64 = 1e10
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 int64 = (v15 % (((^(v15))) | 1))
        v16 -= 0.1
        _ = fi17
        _ = v18
    }
    var v19 uint = 673
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            v15 += v15
            v15 += v15
            fmt.Println(fi20)
            _ = fi21
        }
        v16 = v16
        _ = fi20
    }
    fmt.Println(v2)
    var v22 uint16 = 441
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
