package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }

func main() {
    v1 := []int{-1, -1, -2147483648, 32767}
    var v2 int8 = 100
    var v3 string = ("" + "hello")
    var v4 uint16 = 2
    var v5 int32 = 0
    v6 := make(map[int]uint32)
    v5 += 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 int16 = -787
    v8 := v6[-2147483648]
    fmt.Println(v5, v3, v4)
    var v9 int8 = int8(v7)
    var v10 bool = true
    switch 770 {
    case 4294967295:
        _ = len(v1)
    case 128:
        v8 = (v8 % ((v8) | 1))
    case 27:
        if (!v10) {
            v11 := GenS1{F0: (v4 / ((uint16(v8)) | 1)), F1: 2, F2: 65535}
            v6[32767] = 65535
            _ = v11
        }
    default:
        if (v10 && true) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v10 = v10
                _ = fi12
            }
            v2 = 125
        }
    }
    _ = len(v6)
    for fi13 := 0; fi13 < 14; fi13++ {
        if (v8 > (v8 >> ((785) & 63))) {
            var v14 uint8 = 99
            var v15 int8 = v2
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    var v16 int64 = -772
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -2147483648 {
    case 42:
        v8 = v8
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
