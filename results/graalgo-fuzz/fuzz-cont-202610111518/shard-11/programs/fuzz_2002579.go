package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 int16) int16 {
    v6 := []int{-796, 32768, -325, 128}
    p0 = (p0 | (p0 | -32768))
    fmt.Println(p0)
    fmt.Println(p0)
    var v7 int64 = -239
    _ = p0
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 int = -191
    var v2 uint64 = 0
    var v3 uint = 65535
    var v4 uint64 = 914
    var v5 bool = true
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v3, v8)
            v11 := GenS1{F0: "foo bar", F1: 0}
            _ = fi10
            _ = v11
        }
        v5 = true
        _ = fi9
    }
    v12 := GenS2{F0: 100, F1: (v1 << (((^(v3))) & 63))}
    v13 := GenS3{F0: 1000, F1: -2147483648, F2: (v3 - 4294967295)}
    v14 := []int{(v1 >> (((v3 * v3)) & 63)), (v1 >> ((v3) & 63))}
    var v15 bool = v5
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v3 = (^(v3))
    fmt.Println(v2, v3)
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint16 = 662
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 int32 = -2147483648
            fmt.Println(v4, v2)
            if len(v14) > 0 {
                v14[0] = v1
            }
            _ = fi18
            _ = v19
        }
        v17 += (v17 % (((v17 / ((0) | 1))) | 1))
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
