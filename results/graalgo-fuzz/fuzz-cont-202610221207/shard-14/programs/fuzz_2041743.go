package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64) bool {
    var v6 string = "\n\t"
    v7 := GenS1{F0: 42, F1: 0}
    fmt.Println(p0)
    var v8 int = 32767
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint64, p1 uint16) int16 {
    switch -1 {
    case 2147483647:
        var v9 int8 = 115
        _ = v9
    default:
        var v10 string = "hello"
        _ = v10
    }
    v11 := GenS1{F0: 32767, F1: 0}
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 1; fi13++ {
        if (3.14159 <= 1.5) {
            v14 := []int{fi13, 4294967295, 1}
            v12 = 1
            _ = v14
        } else {
            if false {
                fmt.Println(p0)
                v15 := GenS1{F0: -832, F1: 1}
                _ = v15
            }
        }
        if false {
            fmt.Println(fi13, v12, v12)
            switch (fi13 ^ fi13) {
            case 127:
                v12 -= uint64(p1)
            case 32767:
                v16 := GenS2{F0: 907, F1: fi13, F2: 100}
                _ = v16
            default:
                var v17 uint64 = v12
                _ = v17
            }
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        switch "foo bar" {
        case "":
            v19 := []int{}
            _ = v19
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -304
}

func main() {
    v1 := []int{307}
    var v2 bool = true
    var v3 string = "hello"
    v4 := []int{828, -2147483648}
    var v5 uint64 = 0
    v3 = (v3 + "hello")
    if len(v4) > 0 {
        v4[0] = 32768
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        _ = len(v1)
        _ = fi20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS3{F0: 1}
        _ = fi21
        _ = v22
    }
    var v23 float32 = (0.0 / (1.5 + -2.25))
    _ = len(v4)
    v5 -= (v5 ^ v5)
    v5 = v5
    var v24 uint = 2
    v3 = v3
    v25 := GenS1{F0: 0, F1: 156}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
