package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 bool = true
    v2 := []int{42}
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 1
    v2 = append(v2, 9223372036854775807)
    var v6 uint64 = 837
    if len(v2) > 0 {
        v2[0] = 127
    }
    v6 += v6
    var v7 int16 = -25
    var v8 float32 = -2.25
    v9 := GenS1{F0: 0}
    switch "hello" {
    case "a":
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS2{F0: 255, F1: 42}
            if (fi10 >= (fi10 - fi10)) {
                fmt.Println(fi10, fi10, fi10)
            } else {
                v12 := GenS3{F0: 150, F1: fi10}
                if v1 {
                    var v13 string = ""
                    v14 := v4[v13]
                    _ = v13
                    _ = v14
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    case "foo bar":
        if (-1 >= -1) {
            var v15 float32 = (1e10 - v8)
            var v16 string = "\n\t"
            fmt.Println(v7, v5, v16)
            _ = v15
            _ = v16
        }
    case "\n\t":
        switch 32767 {
        case -2147483648:
            fmt.Println(v6)
        case 1:
            v17 := GenS4{F0: 128, F1: 471, F2: 32}
            _ = v17
        default:
            if v1 {
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v8, v7)
                    v1 = (!(!false))
                    if v1 {
                        var v19 uint32 = 322
                        var v20 float32 = v8
                        _ = v19
                        _ = v20
                    }
                    _ = fi18
                }
                v21 := GenS5{F0: v6}
                _ = v21
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
