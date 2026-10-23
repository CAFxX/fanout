package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 42
    var v2 bool = true
    var v3 uint64 = 0
    v4 := make(map[int]uint32)
    fmt.Println(v1)
    v5 := []int{-1, -2147483648, 32767}
    delete(v4, -1)
    v5 = append(v5, -9223372036854775808)
    var v6 uint64 = v1
    fmt.Println(v2, v6)
    var v7 int = -52
    if (!(!((v6 ^ v1) == v6))) {
        var v8 uint32 = 403
        if v2 {
            if ((v8 | (v8 + v8)) == 1) {
                fmt.Println(v2)
                var v9 uint = 255
                v7 += (v7 << ((v9) & 63))
                _ = v9
            } else {
                var v10 uint64 = (^(v6))
                var v11 uint32 = v8
                _ = v10
                _ = v11
            }
            v3 = v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v12 uint32 = 671
        _ = v12
        _ = v8
    }
    if v2 {
        fmt.Println(v7, v3, v7)
        if v2 {
            v13 := GenS1{F0: -9223372036854775808, F1: 65535}
            v14 := GenS2{F0: 38}
            var v15 bool = v2
            _ = v13
            _ = v14
            _ = v15
        } else {
            var v16 uint32 = 2147483648
            _ = v16
        }
    }
    if v2 {
        var v17 int32 = -193
        var v18 int8 = 100
        _ = v17
        _ = v18
    }
    _ = len(v4)
    var v19 string = ("\n\t" + ("foo bar" + ("\n\t" + "")))
    v2 = true
    v4[v7] = 755
    v1 = v6
    v19 = ""
    var v20 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
