package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (-2.25 - 0.1)
    var v3 int = 127
    var v4 int32 = -357
    var v5 bool = (42 <= 1000)
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 = v5
        _ = fi7
    }
    var v8 bool = v5
    switch 32768 {
    case 128:
        var v9 uint16 = 32767
        _ = v9
    case 1:
        fmt.Println(v3, v4, v5)
    case 2147483647:
        for fi10 := 0; fi10 < 16; fi10++ {
            v5 = v8
            _ = fi10
        }
    }
    if v5 {
        v11 := GenS1{F0: v4, F1: -9223372036854775808, F2: (v8 && v5)}
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := []int{(fi12 >> ((42) & 63))}
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v14 := GenS1{F0: (v4 >> ((2) & 63)), F1: 42, F2: (542 != 0)}
        v2 = v2
        _ = v14
    }
    v5 = v5
    fmt.Println(v8, v2)
    var v15 uint8 = 203
    if (!(!v8)) {
        v16 := GenS1{F0: v4, F1: -9223372036854775808, F2: (("a" <= "hello") && false)}
        var v17 int64 = 1
        _ = v16
        _ = v17
    }
    v18 := []int{v3, (v3 % ((v3) | 1)), v3}
    if len(v18) > 0 {
        v18[0] = (v3 % (((-(v3))) | 1))
    }
    if len(v18) > 0 {
        v18[0] = v3
    }
    v19 := GenS2{F0: 0, F1: 127}
    fmt.Println(v3, v8)
    v6[4294967295] = "\n\t"
    var v20 int32 = v4
    if (v3 == 4294967295) {
        var v21 uint16 = 255
        _ = v21
    }
    switch (v3 + 42) {
    case 32768:
        var v22 bool = (-1 < 232)
        _ = v22
    case -1:
        v3 = (v3 % ((32768) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
