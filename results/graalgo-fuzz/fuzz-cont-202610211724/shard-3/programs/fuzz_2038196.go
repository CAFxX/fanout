package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 2147483648
    v3 := make(map[int]bool)
    var v4 int = 127
    v5 := GenS1{F0: 0, F1: 1, F2: 42}
    if (false || (3.14159 <= 1.5)) {
        fmt.Println(v1)
        v6 := GenS2{F0: uint(v2), F1: v4, F2: -733}
        var v7 float32 = -2.25
        _ = v6
        _ = v7
    }
    v2 = v2
    v1 += 4294967295
    v1 = 1
    for fi8 := 0; fi8 < 14; fi8++ {
        delete(v3, (fi8 + (fi8 / ((32768) | 1))))
        var v9 string = "\n\t"
        v9 = ((v9 + v9) + (((v9 + ((v9 + v9) + "\n\t")) + "hello") + ((("foo bar" + "\n\t") + "foo bar") + ((v9 + (("\n\t" + "\n\t") + "")) + ""))))
        _ = fi8
        _ = v9
    }
    v4 = v4
    v1 -= (v1 | v1)
    v1 = (v1 - (^(v1)))
    var v10 int64 = 9223372036854775807
    v11 := GenS3{F0: int16(v4)}
    v12 := []int{}
    v13 := GenS3{F0: -32768}
    var v14 uint64 = 786
    var v15 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
