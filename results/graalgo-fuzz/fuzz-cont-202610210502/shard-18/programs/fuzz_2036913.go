package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 int16 = -1
    var v2 bool = (!false)
    var v3 uint32 = 1
    var v4 int32 = 1
    var v5 uint = 877
    var v6 int = -986
    var v7 uint32 = v3
    v3 -= 65535
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v4)
        _ = fi8
    }
    fmt.Println(v6)
    if ("foo bar" != (("hello" + "") + (("a" + "") + "a"))) {
        v9 := GenS1{F0: 897, F1: ((v6 + -1) < v6)}
        var v10 float32 = (3.14159 / 1.5)
        if ((v4 % ((v4) | 1)) >= (v4 & 661)) {
            fmt.Println(v4)
            if true {
                v6 += (v6 >> ((uint(v7)) & 63))
            } else {
                v11 := GenS1{F0: 1, F1: true}
                v3 -= v3
                _ = v11
            }
            fmt.Println(v6, v2, v10)
        } else {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int8 = -1
                _ = fi12
                _ = v13
            }
            v3 -= (v3 >> ((v5) & 63))
            if false {
                v7 += (^(v7))
            }
        }
        _ = v10
        _ = v9
    } else {
        fmt.Println(v7, v6)
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := []int{int(v5)}
            fmt.Println(v4)
            _ = fi14
            _ = v15
        }
        v16 := []int{}
        _ = v16
    }
    v7 = (v7 << ((v5) & 63))
    v3 += (v3 * 0)
    v4 = int32(v1)
    v17 := []int{v6}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
