package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    var v5 float64 = 1.5
    var v6 float32 = (1e10 * float32(1.5))
    fmt.Println(v6)
    v7 := GenS1{F0: 1, F1: 42}
    var v8 int16 = 32767
    var v9 int64 = 0
    v10 := GenS2{F0: v9}
    v1 += v1
    v11 := GenS1{F0: (^(v1)), F1: 0}
    var v12 uint = 18446744073709551615
    var v13 int8 = 42
    if false {
        var v14 bool = false
        _ = v14
    } else {
        v15 := []int{-334, -493, 128}
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 string = v3
            fmt.Println(fi16)
            _ = fi17
            _ = v18
        }
        _ = len(v2)
        if true {
            v6 -= (v6 - v6)
            v8 = v8
        }
        _ = fi16
    }
    v6 += v6
    var v19 string = ""
    v20 := v4[-70]
    var v21 int8 = (v13 / (((v13 * v13)) | 1))
    fmt.Println(v21, v6, v9)
    v22 := GenS3{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
