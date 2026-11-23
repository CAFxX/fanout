package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8 }
func genFn0() bool {
    var v6 bool = (32767 <= 27)
    v7 := GenS1{F0: 582, F1: 1, F2: 255}
    var v8 uint16 = 255
    _ = v6
    _ = v7
    _ = v8
    return (32767 != (v8 / (((v8 & v8)) | 1)))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint8 = 128
    var v4 uint8 = v3
    var v5 int8 = 100
    switch "hello" {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        var v9 uint32 = 772
        _ = v9
    default:
        if (!(!false)) {
            v4 += 128
            v3 -= (v4 - (v4 & v4))
        }
    }
    if false {
        v10 := GenS1{F0: 4294967295, F1: 0, F2: (^(v4))}
        v4 += (v4 / ((0) | 1))
        v11 := GenS2{F0: (^(v5))}
        _ = v10
        _ = v11
    } else {
        v12 := []int{}
        _ = len(v12)
        v4 = v3
        _ = v12
    }
    v3 = v4
    if (3.14159 != 1e10) {
        var v13 uint32 = 2147483647
        _ = v13
    }
    v14 := []int{9223372036854775807, 1}
    fmt.Println(v5, v5)
    var v15 uint64 = 18446744073709551615
    v3 = (v4 - 127)
    fmt.Println(v15, v5, v3)
    var v16 string = ("foo bar" + ("" + ((("hello" + ("\n\t" + "a")) + "hello") + "foo bar")))
    v17 := v2[2]
    v18 := GenS1{F0: (v15 >> ((2) & 63)), F1: -9223372036854775808, F2: (v3 << ((18446744073709551615) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
