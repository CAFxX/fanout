package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int16) uint8 {
    if (("a" + "foo bar") >= "a") {
        v6 := []int{42, 0}
        var v7 uint8 = 128
        _ = v6
        _ = v7
    }
    p1 = 1
    fmt.Println(p1, p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int8 = -124
    var v2 string = ("\n\t" + ((("a" + "foo bar") + "foo bar") + "foo bar"))
    var v3 float32 = 3.14159
    v4 := make(map[int]int)
    var v5 int8 = v1
    v5 = v5
    var v8 uint32 = 556
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(fi9, v2)
        _ = fi9
    }
    v2 = v2
    fmt.Println(v2)
    if true {
        if false {
            v10 := v4[32768]
            _ = v10
        } else {
            v11 := GenS1{F0: (false || (((((!false) || false) && true) || (false || (!false))) || false))}
            fmt.Println(v1, v8)
            if false {
                v12 := GenS2{F0: false, F1: 42, F2: 88}
                for fi13 := 0; fi13 < 18; fi13++ {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v1 = -128
                        v15 := []int{fi14, (fi13 ^ fi13), fi14, fi13}
                        fmt.Println(v2, v2, v1)
                        _ = fi14
                        _ = v15
                    }
                    v1 = -44
                    v16 := []int{fi13, (fi13 << ((779) & 63)), -2147483648, 0}
                    _ = fi13
                    _ = v16
                }
                var v17 uint8 = genFn0(2, 1)
                _ = v12
                _ = v17
            }
            _ = v11
        }
    }
    var v18 uint64 = 931
    var v19 bool = (217 <= 26)
    var v20 uint64 = v18
    v21 := []int{-1, 9223372036854775807, 127, -9223372036854775808}
    v22 := GenS3{F0: 147}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
