package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    var v4 string = "foo bar"
    var v5 bool = false
    var v6 uint16 = 0
    var v7 string = v4
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 7; fi9++ {
        v2[2] = (-(fi9))
        var v10 uint = 0
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: (4294967295 <= 999), F1: 1}
    v7 = (v7 + "hello")
    var v12 int16 = -32768
    if ((v8 % ((1) | 1)) != v8) {
        var v13 string = v4
        v14 := GenS1{F0: true, F1: v8}
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 927}
    if false {
        v16 := GenS1{F0: v5, F1: 1}
        _ = v16
    }
    if v5 {
        v17 := GenS3{F0: 226, F1: 1, F2: 127}
        for fi18 := 0; fi18 < 2; fi18++ {
            if (fi18 < (fi18 >> ((2) & 63))) {
                v3 -= v3
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := []int{(fi19 << ((4294967295) & 63))}
                var v21 int64 = 9223372036854775807
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    }
    if v5 {
        v5 = v5
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 float64 = 3.14159
            _ = fi22
            _ = v23
        }
        var v24 float32 = (0.0 * ((((1.5 - (0.1 / (3.14159 + float32(float64(0.0))))) * 0.1) + 3.14159) * (3.14159 + (3.14159 / float32(1e10)))))
        _ = v24
    } else {
        v6 = 570
        var v25 uint8 = v8
        fmt.Println(v25)
        _ = v25
    }
    v8 = 42
    v26 := []int{32768, -2147483648, -903}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
