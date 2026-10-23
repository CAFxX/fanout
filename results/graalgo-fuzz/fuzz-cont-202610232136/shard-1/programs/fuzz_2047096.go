package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    if (1.5 > (-2.25 * 0.0)) {
        var v6 int = -184
        v6 = (v6 % ((v6) | 1))
        var v7 bool = false
        _ = v6
        _ = v7
    }
    var v8 int8 = 0
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 42
    var v3 uint = 65535
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v9 float32 = 0.0
    if true {
        v9 = v9
        fmt.Println(v2)
        switch 42 {
        case 1:
            for fi10 := 0; fi10 < 16; fi10++ {
                v1[genFn0()] = 42
                _ = fi10
            }
        }
    }
    switch -107 {
    case 0:
        if (-32768 <= -869) {
            if false {
                if ((v2 * -927) < v2) {
                    fmt.Println(v9, v2, v2)
                    v3 = v3
                    var v11 int64 = -9223372036854775808
                    _ = v11
                } else {
                    var v12 float64 = 0.1
                    v13 := GenS1{F0: 2}
                    v14 := GenS2{F0: 128}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                fmt.Println(v2)
            }
            v9 = v9
            v15 := GenS2{F0: 0}
            _ = v15
        } else {
            var v16 uint8 = 42
            v17 := []int{int(v2)}
            v3 = v3
            _ = v16
            _ = v17
        }
    case -137:
        var v18 float64 = 1e10
        _ = v18
    default:
        switch "\n\t" {
        case "a":
            v3 += v3
        default:
            var v19 float32 = v9
            _ = v19
        }
    }
    fmt.Println(v2, v3)
    _ = len(v4)
    var v20 uint8 = 0
    var v21 int16 = -830
    v22 := []int{-314, 0, 374, -652}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
