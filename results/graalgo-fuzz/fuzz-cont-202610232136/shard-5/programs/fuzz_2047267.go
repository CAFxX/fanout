package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 uint16, p2 uint32) uint16 {
    var v4 int8 = -128
    v5 := []int{}
    var v6 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (p1 - p1)
}

func main() {
    var v1 int64 = 679
    v2 := []int{4294967295, -998, 0}
    var v3 int8 = -1
    var v7 float32 = -2.25
    var v8 uint16 = genFn0(-938, 204, 2)
    if (42 >= 1) {
        var v9 uint16 = (v8 / (((v8 | (v8 & (v8 | v8)))) | 1))
        var v10 int16 = 1000
        _ = v10
        _ = v9
    } else {
        var v11 float32 = 3.14159
        var v12 int32 = 923
        _ = len(v2)
        _ = v11
        _ = v12
    }
    if false {
        if (("foo bar" + ("\n\t" + "hello")) >= "\n\t") {
            v8 += v8
            for fi13 := 0; fi13 < 13; fi13++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v7 -= v7
                var v14 uint = 42
                _ = fi13
                _ = v14
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        v2 = append(v2, 63)
        v2 = append(v2, -2147483648)
        if (v1 == v1) {
            fmt.Println(v7, v7, v3)
        }
    }
    v8 += genFn0(int(v3), v8, 659)
    switch 4294967295 {
    case 605:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = (v3 + 84)
    v15 := GenS1{F0: (v3 + (^(v3))), F1: 170}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
