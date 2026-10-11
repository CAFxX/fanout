package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0(p0 bool, p1 uint32, p2 bool) uint64 {
    v7 := GenS1{F0: 42}
    fmt.Println(p1, p0)
    var v8 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v8 * (v8 + v8))
}

func main() {
    var v1 bool = (!(990 >= 779))
    var v2 int8 = 100
    var v3 int32 = -419
    var v4 uint64 = 449
    var v5 string = "a"
    var v6 int64 = -1
    v9 := GenS2{F0: v3, F1: (v2 ^ int8(v4))}
    v4 = (v4 | v4)
    var v10 int16 = 259
    v11 := []int{}
    v4 += v4
    fmt.Println(v10)
    switch 2147483647 {
    case -9223372036854775808:
        if len(v11) > 0 {
            v11[0] = 866
        }
    default:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    var v12 int16 = v10
    var v13 int16 = -1
    v6 = (v6 ^ (v6 | (v6 * -9223372036854775808)))
    var v14 bool = ((!(true || true)) || true)
    var v15 string = v5
    if ((v4 << ((4294967295) & 63)) <= genFn0((v1 && (!(v1 && ((v3 & -614) <= -1)))), 147, (2147483647 != 226))) {
        v2 = (v2 % ((v2) | 1))
    }
    fmt.Println(v14, v4, v14)
    v4 = genFn0(v14, 1, v1)
    var v16 float32 = -2.25
    v11 = append(v11, -205)
    v2 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
