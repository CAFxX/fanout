package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 string }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]string)
    var v3 uint8 = 96
    var v4 int16 = 1000
    var v5 uint16 = 42
    var v6 uint32 = 1
    var v7 bool = true
    v8 := GenS1{F0: (v5 >> ((2) & 63)), F1: (0 > 42), F2: -790}
    v9 := []int{2147483647, 0}
    if (!(("\n\t" + (("hello" + "") + (("a" + "hello") + (("a" + "") + (("\n\t" + "\n\t") + "hello"))))) >= "foo bar")) {
        if (1e10 >= (((3.14159 / ((((1.5 / 1e10) + 1.5) + (float32(3.14159) / -2.25)) - float32(3.14159))) * (3.14159 / ((1.5 - ((1.5 / ((1e10 * float32(1e10)) / 0.1)) + 3.14159)) * float32((3.14159 / 1.0))))) * float32(0.0))) {
            v10 := GenS2{F0: v1, F1: 70, F2: 1}
            fmt.Println(v7, v3, v4)
            _ = v10
        }
        var v11 uint = 65535
        _ = v11
    }
    v1 += v1
    var v12 int64 = 500
    var v13 uint8 = v1
    v14 := v2[127]
    fmt.Println(v7)
    var v15 float32 = 0.1
    v16 := GenS2{F0: 127, F1: 0, F2: v5}
    if true {
        var v17 int8 = 37
        if (v14 < v14) {
            fmt.Println(v12)
        }
        _ = v17
    }
    if true {
        var v18 int16 = -1
        _ = v18
    } else {
        switch -162 {
        case 0:
            v19 := GenS1{F0: (^(v5)), F1: true, F2: 127}
            _ = v19
        }
    }
    if false {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 float32 = v15
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint = 0
                var v24 int32 = int32(v5)
                v25 := GenS3{F0: v1, F1: v12, F2: (v14 + (v14 + v14))}
                v23 = 528
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        v5 -= 2
        var v26 uint = 65535
        _ = v26
    } else {
        v13 += v3
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
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
