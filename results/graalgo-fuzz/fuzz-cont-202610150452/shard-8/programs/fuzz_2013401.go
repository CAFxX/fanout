package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 32767
    var v3 string = "a"
    var v4 string = ((v3 + v3) + "")
    var v5 int64 = -1
    var v6 int = 42
    var v7 float32 = 1e10
    var v8 int16 = -1
    var v9 float32 = 3.14159
    v10 := GenS1{F0: 804}
    var v11 uint32 = 272
    var v12 int = -2147483648
    switch (v4 + "foo bar") {
    case "":
        v13 := []int{-2147483648, 9223372036854775807, v6, v12}
        _ = v13
    default:
        v14 := []int{9223372036854775807}
        _ = v14
    }
    v5 += v5
    v1 -= 2
    v15 := GenS1{F0: 679}
    if false {
        if (0 >= 242) {
            v16 := GenS2{F0: (v2 * v2), F1: 1}
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := []int{2147483647}
                    var v20 int16 = 683
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
            _ = v16
        }
        v21 := []int{v12, v6}
        v22 := GenS3{F0: (190 <= 127), F1: 0, F2: -1}
        _ = v21
        _ = v22
    } else {
        v23 := []int{(v12 - 126), v6, 9223372036854775807, (v6 / ((v6) | 1))}
        _ = v23
    }
    if false {
        v24 := GenS4{F0: 128}
        v25 := GenS4{F0: 155}
        _ = v24
        _ = v25
    }
    if ((false || false) && (-745 != 42)) {
        var v26 uint = v1
        fmt.Println(v12, v1, v7)
        _ = v26
    } else {
        v27 := []int{127, v12, (v12 << ((v1) & 63))}
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := []int{(v6 % ((42) | 1)), fi28, (v6 ^ fi28), v6}
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
