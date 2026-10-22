package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []int{165, 32767, 176, -939}
    v2 := []uint32{0}
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    var v6 int32 = 2147483647
    var v7 uint = (v5 >> ((18446744073709551615) & 63))
    var v8 string = ("\n\t" + ((((("" + ("foo bar" + "")) + (("\n\t" + "hello") + "")) + "a") + "hello") + "\n\t"))
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: false, F1: 579}
        _ = fi9
        _ = v10
    }
    _ = len(v2)
    if true {
        fmt.Println(v7, v7)
        v5 = (^(v7))
    }
    if (569 <= 389) {
        fmt.Println(v8, v7)
    }
    v6 -= (-(v6))
    v11 := GenS1{F0: true, F1: 574}
    delete(v3, v8)
    fmt.Println(v8, v5)
    if (255 != 70) {
        var v12 string = v8
        _ = v12
    } else {
        fmt.Println(v5, v6, v5)
    }
    var v13 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
