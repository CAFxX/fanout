package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int; F1 string; F2 uint8 }
func genFn0() string {
    var v6 float64 = 3.14159
    if ((0.0 - 0.1) == 1.5) {
        v7 := GenS1{F0: 2147483647}
        _ = v7
    } else {
        if (2147483647 != 806) {
            fmt.Println(v6)
            v6 = v6
            v6 += (v6 - ((1e10 - (0.1 + v6)) * (v6 + 1.5)))
        } else {
            var v8 int16 = 42
            v8 += (v8 + v8)
            fmt.Println(v8, v8)
            _ = v8
        }
    }
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 127
    var v3 uint64 = 1
    v4 := []uint32{uint32(v3)}
    var v5 int = 2147483647
    v9 := GenS2{F0: 4294967295, F1: v2}
    v4 = append(v4, 2)
    v3 -= v3
    var v10 uint = 2
    fmt.Println(v10, v10)
    var v11 uint8 = 0
    var v12 int8 = -1
    v5 += (v5 + (v5 | v5))
    v10 = (v10 + v10)
    if true {
        var v13 uint16 = 496
        v14 := GenS3{F0: (-(v5)), F1: uint32(v13)}
        _ = v13
        _ = v14
    }
    v1[320] = ("\n\t" + (("" + ("\n\t" + ("a" + ("a" + "")))) + "a"))
    v15 := GenS4{F0: 9223372036854775807, F1: ("" + ""), F2: 82}
    v2 -= v2
    _ = len(v1)
    var v16 bool = true
    v2 -= v2
    v17 := GenS4{F0: v5, F1: "", F2: (v11 | (v11 >> (((v10 & (v10 | uint(v11)))) & 63)))}
    v18 := GenS1{F0: int32(v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
