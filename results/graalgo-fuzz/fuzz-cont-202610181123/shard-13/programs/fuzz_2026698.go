package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }

func main() {
    var v1 uint8 = 2
    v2 := []string{"a", "foo bar", "\n\t"}
    v3 := make(map[uint32]int)
    var v4 int8 = -59
    v5 := make(map[int]bool)
    v6 := []string{"hello", "hello", "", "hello"}
    fmt.Println(v4)
    if true {
        v4 -= v4
        v7 := GenS1{F0: 0}
        _ = v7
    }
    var v8 uint64 = 930
    v1 += 2
    v9 := GenS2{F0: (-(v4)), F1: 255}
    v4 += (v4 - v4)
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            switch fi11 {
            case 64:
                var v12 string = "a"
                _ = v12
            case 106:
                var v13 uint64 = v8
                _ = v13
            case 42:
                v14 := GenS3{F0: 157, F1: uint16(v8), F2: -286}
                _ = v14
            default:
                v1 = (v1 * v1)
            }
            var v15 uint8 = (v1 & v1)
            switch fi11 {
            case -1:
                v4 += v4
            case 32767:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi11
            _ = v15
        }
        _ = fi10
    }
    var v16 int32 = 2147483647
    var v17 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
