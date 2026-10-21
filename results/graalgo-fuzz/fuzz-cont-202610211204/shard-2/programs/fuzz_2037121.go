package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
func genFn0(p0 int16) uint8 {
    v4 := []int{}
    p0 = -948
    var v5 uint8 = 1
    var v6 uint64 = 865
    var v7 uint16 = 255
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v5
}

func main() {
    var v1 uint32 = 993
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v8 int8 = -49
    v1 = v1
    v9 := v3[2147483647]
    v10 := v2[(v9 | (^(v9)))]
    v9 = v1
    v11 := []int{-9223372036854775808, 413, 319}
    if (255 <= genFn0(-32768)) {
        v12 := GenS1{F0: (v9 + 2147483648), F1: -1, F2: (v8 ^ -128)}
        fmt.Println(v8, v9)
        if true {
            v10 = ("\n\t" + "hello")
            if (0 >= 433) {
                var v13 uint = 901
                _ = v13
            }
            v8 = (v8 ^ (v8 >> ((563) & 63)))
        }
        _ = v12
    }
    v8 -= v8
    fmt.Println(v9, v9, v10)
    v14 := v3[(v9 % ((v1) | 1))]
    v15 := GenS1{F0: 4294967295, F1: -2147483648, F2: 100}
    if true {
        v10 = (v10 + ((("foo bar" + v10) + v10) + ("\n\t" + (("\n\t" + "") + ((v10 + ("" + (v10 + ""))) + v10)))))
        v11 = append(v11, -66)
    }
    v11 = append(v11, 32767)
    if true {
        var v16 uint8 = 78
        switch 127 {
        case -2147483648:
            v17 := GenS1{F0: v9, F1: -407, F2: 100}
            _ = v17
        case -886:
            v11 = append(v11, 0)
        case -554:
            var v18 int16 = 1000
            _ = v18
        default:
            var v19 bool = false
            _ = v19
        }
        switch 32767 {
        case 2147483647:
            _ = len(v11)
        case -867:
            v8 += (v8 | v8)
        default:
            if true {
                v10 = v10
                v11 = append(v11, -2147483648)
            }
        }
        _ = v16
    } else {
        switch -906 {
        case 4294967295:
            var v20 bool = false
            _ = v20
        default:
            fmt.Println(v8, v9)
        }
        var v21 int64 = 42
        v1 += (v1 / ((v14) | 1))
        _ = v21
    }
    fmt.Println(v9, v8)
    switch v10 {
    case "foo bar":
        fmt.Println(v10, v8)
    case "a":
        _ = len(v11)
    default:
        for fi22 := 0; fi22 < 17; fi22++ {
            fmt.Println(v8)
            var v23 uint = 18446744073709551615
            _ = fi22
            _ = v23
        }
    }
    var v24 uint8 = 127
    var v25 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v8
    _ = v9
}
