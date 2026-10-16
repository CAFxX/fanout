package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    v1 := []uint32{1, 1, 219}
    var v2 uint = 255
    var v3 uint8 = 42
    var v4 int = 32767
    v5 := make(map[uint32]string)
    v6 := []int{2147483647}
    var v7 float32 = (1e10 - -2.25)
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            if (true && true) {
                if true {
                    v7 += (-2.25 / v7)
                } else {
                    var v10 uint8 = (^(v3))
                    v6 = append(v6, v4)
                    v10 += (^(v3))
                    _ = v10
                }
                v11 := GenS1{F0: 729, F1: true}
                var v12 uint8 = (v3 | (v3 & v3))
                _ = v11
                _ = v12
            }
            var v13 uint = 1
            _ = fi9
            _ = v13
        }
        _ = fi8
    }
    var v14 uint8 = v3
    v5[2147483648] = ("foo bar" + "")
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 -= v2
        _ = len(v5)
        fmt.Println(fi15)
        _ = fi15
    }
    v7 = 0.1
    var v16 int16 = 32767
    fmt.Println(v7, v3, v14)
    fmt.Println(v16, v3)
    var v17 uint8 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
