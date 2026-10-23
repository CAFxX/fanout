package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 42
    var v3 uint64 = 2
    for fi4 := 0; fi4 < 17; fi4++ {
        if true {
            var v5 int64 = 42
            v2 -= (v2 | (^(v2)))
            for fi6 := 0; fi6 < 3; fi6++ {
                fmt.Println(fi4)
                v3 += (v3 - v3)
                _ = fi6
            }
            _ = v5
        }
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: -1, F1: -32768}
            var v9 int32 = -1
            var v10 int32 = v9
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        if (true && true) {
            var v11 int32 = 42
            for fi12 := 0; fi12 < 7; fi12++ {
                switch "hello" {
                case "foo bar":
                    v13 := []int{fi12, fi4, 128, fi12}
                    v14 := GenS1{F0: 42, F1: -1}
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            v15 := GenS2{F0: 100}
            _ = v11
            _ = v15
        } else {
            if true {
                fmt.Println(v2, v2, fi4)
            }
            v2 -= v2
        }
        _ = fi4
    }
    _ = len(v1)
    var v16 uint8 = uint8(v2)
    v2 = 32767
    if true {
        var v17 bool = true
        _ = v17
    }
    var v18 bool = true
    v19 := []int{-1, 32767}
    var v20 float32 = (0.0 + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
