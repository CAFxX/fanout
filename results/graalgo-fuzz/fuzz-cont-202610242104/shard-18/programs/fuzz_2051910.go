package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := []int{9223372036854775807}
    v5 := make(map[uint32]int)
    v6 := []int{}
    v7 := GenS1{F0: 128, F1: 356}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS1{F0: 172, F1: 839}
    if true {
        if true {
            v9 := GenS2{F0: 114}
            var v10 uint32 = 42
            _ = v10
            _ = v9
        } else {
            if false {
                if len(v6) > 0 {
                    v6[0] = 42
                }
            } else {
                var v11 int = 32768
                _ = v11
            }
            v1 = v1
            if true {
                var v12 int8 = 42
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint64 = 18446744073709551615
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                v1 = (v1 + v1)
                for fi15 := 0; fi15 < 3; fi15++ {
                    v6 = append(v6, 171)
                    _ = fi15
                }
            }
        }
        switch 32768 {
        case -351:
            var v16 uint32 = 4294967295
            _ = v16
        }
        switch v1 {
        case "\n\t":
            v1 = v1
        default:
            v17 := GenS3{F0: true}
            _ = v17
        }
    } else {
        var v18 int = -914
        if (-651 != 42) {
            v18 = 2147483647
            var v19 string = v1
            _ = v19
        }
        _ = v18
    }
    v1 = v1
    v1 = "a"
    switch v1 {
    case "":
        var v20 int = 42
        _ = v20
    }
    v21 := GenS4{F0: 797}
    v1 = "a"
    v22 := GenS4{F0: 42}
    var v23 int64 = -438
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
