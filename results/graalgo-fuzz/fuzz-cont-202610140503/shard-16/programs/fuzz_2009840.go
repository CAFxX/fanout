package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    p0 -= p0
    v4 := []int{-404, -2147483648, -885, 2147483647}
    p0 += 468
    var v5 int8 = -42
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := []int{-1, -9223372036854775808, 916, -9223372036854775808}
    var v3 bool = (65535 < 2147483648)
    var v6 int = 32767
    v3 = (v6 <= (v6 ^ (v6 % ((791) | 1))))
    var v7 int64 = 593
    var v8 int = (v6 % ((int(v1)) | 1))
    if ((!((2 <= v1) || (!true))) || (185 <= 2)) {
        fmt.Println(v8, v7, v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v9 int64 = v7
        _ = v9
    }
    v10 := GenS1{F0: "", F1: (v3 && false)}
    for fi11 := 0; fi11 < 13; fi11++ {
        if (1 <= 442) {
            v3 = v3
        }
        v12 := GenS2{F0: (v7 % ((v7) | 1)), F1: 2}
        _ = fi11
        _ = v12
    }
    v3 = (v3 || (!(!(false && (false || false)))))
    var v13 uint8 = (v1 / ((255) | 1))
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
