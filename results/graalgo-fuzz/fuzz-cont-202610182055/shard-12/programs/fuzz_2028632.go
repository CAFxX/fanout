package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 2147483647
    v3 := make(map[int]string)
    if true {
        if false {
            v1[65535] = v2
        } else {
            var v4 uint16 = 2
            fmt.Println(v2, v4)
            var v5 int8 = 7
            _ = v4
            _ = v5
        }
        v6 := GenS1{F0: uint8(v2)}
        _ = v6
    } else {
        v2 = v2
        fmt.Println(v2)
    }
    var v7 uint32 = v2
    var v8 bool = true
    if v8 {
        v9 := GenS1{F0: 255}
        v10 := v1[v2]
        delete(v1, 183)
        _ = v10
        _ = v9
    } else {
        if false {
            var v11 float64 = -2.25
            if true {
                if (9 <= -12) {
                    switch "foo bar" {
                    case "a":
                        var v12 int8 = 0
                        _ = v12
                    default:
                        var v13 uint = 65535
                        _ = v13
                    }
                }
            } else {
                v7 = (v2 & (v7 + v2))
                v14 := GenS1{F0: 127}
                var v15 uint16 = 32767
                _ = v14
                _ = v15
            }
            delete(v1, v2)
            _ = v11
        }
        if true {
            var v16 uint8 = 255
            _ = v16
        }
    }
    var v17 uint16 = 1
    v8 = (false && v8)
    if v8 {
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(fi18, v17, fi18)
            _ = fi18
        }
    }
    v19 := GenS2{F0: ((-1 <= -32768) || (!((0.0 + 0.1) == (0.0 / (1.5 - 3.14159))))), F1: 255}
    if v8 {
        fmt.Println(v7, v8, v2)
        var v20 uint16 = v17
        v21 := []int{-1}
        _ = v20
        _ = v21
    }
    var v22 int32 = -253
    v23 := []int{32768}
    v24 := GenS3{F0: 2147483648, F1: "foo bar"}
    var v25 uint64 = 1
    var v26 int = -322
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v7
    _ = v8
}
