package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
func genFn0(p0 uint64) int {
    p0 -= (p0 >> ((721) & 63))
    var v6 uint16 = 554
    v7 := GenS1{F0: 42}
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int, p1 int32, p2 int) int16 {
    p1 -= p1
    p1 -= (p1 & -7)
    var v8 int16 = -32768
    v9 := GenS2{F0: 0}
    v10 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 3.14159
    var v3 uint8 = 127
    var v4 int64 = 1
    var v5 float32 = ((v2 / 1.0) + 1e10)
    v11 := []int{9223372036854775807, 42, 792}
    var v12 int64 = (v4 * v4)
    var v13 int32 = 0
    var v14 string = ""
    var v15 uint = 872
    if false {
        if true {
            var v16 int16 = 32767
            _ = v16
        } else {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
    } else {
        var v17 uint64 = (v1 ^ v1)
        v18 := GenS2{F0: (v3 % (((v3 + v3)) | 1))}
        v14 = ""
        _ = v17
        _ = v18
    }
    v12 = v12
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v5, v13)
    default:
        var v19 float64 = 0.0
        _ = v19
    }
    v20 := GenS3{F0: -969, F1: v15, F2: (v4 << (((^(v15))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
