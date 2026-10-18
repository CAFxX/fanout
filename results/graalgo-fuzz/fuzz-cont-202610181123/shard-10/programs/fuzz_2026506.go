package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 bool; F1 int8; F2 int8 }

func main() {
    var v1 float32 = 0.0
    var v2 uint = 0
    var v3 uint8 = 255
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    var v6 int8 = -128
    v7 := []int{-737}
    var v8 int32 = -230
    fmt.Println(v3, v1, v2)
    v1 = 0.0
    var v9 int16 = 42
    fmt.Println(v3, v3, v2)
    if true {
        delete(v4, 127)
        if false {
            var v10 int16 = v9
            _ = v10
        } else {
            v11 := GenS1{F0: 4294967295, F1: true, F2: (v9 - (^(v9)))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v11
        }
        v12 := GenS2{F0: 1, F1: v2, F2: 1}
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: (856 < v9), F2: (v9 | v9)}
    v1 += v1
    switch "\n\t" {
    case "a":
        if false {
            fmt.Println(v2)
        }
    default:
        if true {
            var v14 int16 = v9
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v6)
                    _ = fi16
                }
                v17 := GenS1{F0: 42, F1: (2 < 42), F2: int16(v3)}
                _ = fi15
                _ = v17
            }
            _ = v14
        } else {
            var v18 uint32 = 994
            _ = v18
        }
    }
    var v19 int16 = v9
    v1 = 3.14159
    v5[-235] = ((v6 / (((v6 ^ v6)) | 1)) >= 24)
    var v20 int16 = -32768
    fmt.Println(v1, v6, v20)
    v9 = v9
    if false {
        v21 := GenS3{F0: 166, F1: 0}
        if ((v6 ^ (v6 | (v6 + v6))) <= (^(v6))) {
            v3 += (v3 % (((^(v3))) | 1))
        } else {
            v22 := GenS4{F0: false, F1: (v6 & -114), F2: v6}
            if false {
                _ = len(v5)
                var v23 int8 = (v6 ^ 0)
                _ = v23
            } else {
                var v24 string = "a"
                _ = v24
            }
            v25 := GenS2{F0: -576, F1: v2, F2: 42}
            _ = v22
            _ = v25
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
