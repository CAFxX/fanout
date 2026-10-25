package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 string = ("hello" + ("a" + ("" + "a")))
    var v2 float64 = 3.14159
    var v3 string = ""
    var v4 int16 = -32768
    var v5 uint8 = 21
    v6 := make(map[int]uint32)
    var v7 uint = 1
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(v3, v3)
            if ((3.14159 * ((-2.25 - (v2 + v2)) * -2.25)) <= (v2 * v2)) {
                var v9 int32 = 0
                v10 := []int{(fi8 & (fi8 | (fi8 | fi8)))}
                _ = v10
                _ = v9
            } else {
                _ = len(v6)
            }
            _ = fi8
        }
        var v11 float32 = float32(v2)
        _ = v11
    } else {
        v12 := GenS1{F0: 2147483648}
        _ = v12
    }
    if true {
        if true {
            v4 -= 990
        } else {
            if (32767 > 727) {
                v4 = v4
                v13 := []int{32767}
                var v14 uint64 = 362
                _ = v13
                _ = v14
            } else {
                fmt.Println(v4)
                var v15 float64 = v2
                _ = v15
            }
            var v16 int16 = v4
            v7 -= (^(v7))
            _ = v16
        }
        v2 = (v2 / -2.25)
    } else {
        var v17 uint = 618
        _ = v17
    }
    var v18 int64 = 1
    v2 += float64(1.5)
    var v19 uint64 = 42
    for fi20 := 0; fi20 < 20; fi20++ {
        v1 = v3
        v1 = ((v3 + "a") + v1)
        v21 := GenS1{F0: 2147483647}
        _ = fi20
        _ = v21
    }
    fmt.Println(v7, v7, v5)
    v22 := []int{9223372036854775807}
    var v23 int64 = v18
    v22 = append(v22, 32768)
    var v24 int64 = int64(v19)
    if true {
        v1 = "a"
        v23 = (v23 - -547)
    }
    delete(v6, 825)
    v25 := GenS2{F0: 255, F1: -2147483648}
    v18 -= 9223372036854775807
    switch "" {
    case "\n\t":
        var v26 uint = v7
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
