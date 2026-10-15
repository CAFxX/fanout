package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 float32 = 0.0
    var v3 uint = 166
    fmt.Println(v2)
    var v4 uint = v3
    var v5 uint8 = (v1 >> ((v3) & 63))
    switch 823 {
    case -862:
        v6 := []int{2147483647, 9223372036854775807}
        _ = v6
    default:
        v7 := []int{32767}
        _ = v7
    }
    var v8 int64 = -1
    switch 32768 {
    case 2147483647:
        v9 := []int{-1, 1}
        _ = v9
    case 642:
        if (v5 >= (v1 ^ 51)) {
            v10 := GenS1{F0: 32767, F1: false}
            var v11 float64 = -2.25
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: 1, F1: false}
            fmt.Println(v1, v3, v5)
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if false {
            fmt.Println(fi13, v5, fi13)
            var v14 float32 = ((float32((-2.25 + ((-2.25 + (1e10 * (-2.25 / (1.5 / float64(0.1))))) * 3.14159))) + float32(1e10)) / (((0.1 + v2) - v2) * v2))
            _ = v14
        }
        _ = fi13
    }
    v15 := GenS1{F0: 42, F1: (1 == -114)}
    v16 := GenS2{F0: (v5 % ((128) | 1)), F1: 194, F2: "foo bar"}
    v17 := GenS2{F0: (v1 ^ (v5 << (((v4 % (((v4 + 688)) | 1))) & 63))), F1: 0, F2: ("" + "a")}
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
