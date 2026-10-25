package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 uint32, p1 uint8) string {
    v5 := GenS1{F0: (128 >= p1)}
    v6 := GenS2{F0: (p1 * 255), F1: -434, F2: p1}
    var v7 uint32 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 uint8) string {
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = true
        if (1000 != 42) {
            fmt.Println(p0, fi8, fi8)
            v10 := []int{}
            _ = v10
        } else {
            v11 := []int{fi8, fi8, (fi8 ^ fi8)}
            v12 := GenS3{F0: (p0 - 255), F1: 42}
            if false {
                var v13 string = ("\n\t" + (("a" + "") + "hello"))
                _ = v13
            }
            _ = v11
            _ = v12
        }
        v9 = (65535 <= 2)
        _ = fi8
        _ = v9
    }
    v14 := GenS2{F0: (p0 * 214), F1: -375, F2: (p0 & (p0 & 255))}
    _ = p0
    _ = v14
    return "a"
}

func main() {
    var v1 int16 = -176
    var v2 int16 = -1
    var v3 int = -387
    var v4 uint64 = 0
    var v15 bool = (((14 > 0) && false) && (v1 < 0))
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            fmt.Println(v2, v2)
            var v17 string = "\n\t"
            _ = v17
        }
        var v18 int = fi16
        var v19 uint64 = (v4 | (v4 ^ v4))
        _ = fi16
        _ = v18
        _ = v19
    }
    var v20 uint8 = 1
    v15 = (!v15)
    v4 -= v4
    var v21 string = "\n\t"
    v2 -= 1000
    v22 := GenS2{F0: v20, F1: -2147483648, F2: (v20 + (v20 - v20))}
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
