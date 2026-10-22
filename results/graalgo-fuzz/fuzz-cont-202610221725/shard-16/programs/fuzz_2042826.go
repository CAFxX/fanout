package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int8) int {
    fmt.Println(p1, p1)
    switch "\n\t" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            v7 := GenS1{F0: 0, F1: 42, F2: 937}
            _ = v7
        }
    case "foo bar":
        var v8 uint16 = 0
        _ = v8
    case "":
        var v9 string = "\n\t"
        _ = v9
    default:
        var v10 int16 = 1
        _ = v10
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 string) int16 {
    v11 := []int{32767, 228, genFn0(251, 42), -773}
    fmt.Println(p1, p0)
    p0 = (p0 % ((2) | 1))
    var v12 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535, 663, 22, 2}
    v3 := []string{}
    var v4 int8 = -1
    var v5 int16 = -369
    v6 := []uint32{42, 4294967295}
    if false {
        switch ("hello" + (("foo bar" + ("foo bar" + "hello")) + "foo bar")) {
        case "a":
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(fi13)
                v4 -= v4
                v14 := GenS2{F0: v5, F1: 100}
                _ = fi13
                _ = v14
            }
        }
    }
    var v15 float64 = float64((3.14159 / 3.14159))
    v16 := GenS3{F0: 85}
    fmt.Println(v5, v1, v15)
    fmt.Println(v15, v4)
    v1 = (0.0 >= (1e10 / 1e10))
    var v17 uint32 = 2
    for fi18 := 0; fi18 < 16; fi18++ {
        if (((fi18 % ((fi18) | 1)) >= genFn0(2, (^(v4)))) && (!v1)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = v1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi18
    }
    v4 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
