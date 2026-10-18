package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0() uint64 {
    v5 := []int{301, 32767, 384}
    var v6 int16 = 32767
    v6 += v6
    var v7 int8 = 0
    fmt.Println(v7, v7)
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    v3 := make(map[int]bool)
    v4 := []int{4294967295, -367}
    v2 = v2
    fmt.Println(v2)
    v2 = "hello"
    for fi8 := 0; fi8 < 5; fi8++ {
        switch 32767 {
        case 42:
            for fi9 := 0; fi9 < 8; fi9++ {
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v10 string = "a"
                    _ = v10
                } else {
                    v11 := GenS1{F0: 2, F1: 858, F2: 42}
                    var v12 string = v2
                    v12 = v2
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        case -1:
            var v13 uint64 = 18446744073709551615
            _ = v13
        default:
            var v14 float64 = 3.14159
            _ = v14
        }
        _ = fi8
    }
    fmt.Println(v2)
    var v15 uint32 = 373
    v16 := GenS1{F0: 42, F1: 0, F2: 1000}
    if (true && (true || true)) {
        v17 := GenS2{F0: 213}
        _ = v17
    } else {
        v2 = "foo bar"
        v18 := GenS1{F0: 2, F1: 42, F2: -32768}
        v2 = v2
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
