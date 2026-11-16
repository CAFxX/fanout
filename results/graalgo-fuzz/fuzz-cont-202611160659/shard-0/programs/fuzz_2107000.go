package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0(p0 uint32) bool {
    v5 := []int{-654, -1}
    if (((!((1 < 18446744073709551615) && false)) || false) && false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v6 int32 = -2147483648
        _ = len(v5)
        _ = v6
    } else {
        v5 = append(v5, -287)
        v7 := GenS1{F0: 31, F1: 32768}
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: 385}
    _ = p0
    _ = v5
    _ = v8
    return (3.14159 > (1e10 * (1e10 + 0.1)))
}
func genFn1(p0 uint32, p1 int32) uint8 {
    var v9 int32 = (p1 | (p1 ^ -142))
    switch "a" {
    case "":
        fmt.Println(p1, p1)
    case "a":
        for fi10 := 0; fi10 < 12; fi10++ {
            switch (^(fi10)) {
            case 385:
                var v11 int8 = 57
                _ = v11
            case 128:
                var v12 float32 = 1.5
                _ = v12
            }
            v13 := []int{(fi10 % ((-1) | 1))}
            _ = fi10
            _ = v13
        }
    }
    fmt.Println(p1)
    v14 := []int{-834, -2147483648}
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return 2
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{4294967295}
    var v3 int8 = 98
    v4 := make(map[uint32]string)
    _ = len(v2)
    v2 = append(v2, 2147483648)
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 -= v3
        _ = fi15
    }
    _ = len(v2)
    fmt.Println(v3, v3)
    switch -19 {
    case -9223372036854775808:
        var v16 int = 726
        _ = v16
    case 32768:
        var v17 uint16 = 255
        _ = v17
    }
    v18 := GenS1{F0: int8(v1), F1: 32768}
    v19 := GenS1{F0: v3, F1: 190}
    v4[2] = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
