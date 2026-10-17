package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    var v3 uint8 = 127
    v4 := []uint32{}
    var v5 float32 = (0.0 / -2.25)
    v6 := []uint32{2147483647}
    v3 += v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 := GenS1{F0: "foo bar", F1: (v1 * v1), F2: 42}
    var v8 string = "\n\t"
    v9 := GenS2{F0: v8}
    var v10 string = (v8 + v8)
    switch -529 {
    case 127:
        if true {
            v10 = v10
            v10 = v8
            switch -876 {
            case 42:
                v11 := GenS2{F0: (v10 + (v10 + ""))}
                _ = v11
            case 1:
                v1 += (v1 + v1)
            case 202:
                var v12 uint = 255
                _ = v12
            default:
                if (100 <= 100) {
                    fmt.Println(v5)
                    v13 := GenS3{F0: 501, F1: 127}
                    _ = v13
                } else {
                    _ = len(v2)
                }
            }
        } else {
            var v14 int = 16
            var v15 int8 = 0
            v14 += v14
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v3)
    var v16 uint8 = v3
    var v17 int16 = v1
    var v18 uint16 = 186
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
