package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0() int {
    var v5 uint32 = 2147483648
    v6 := []int{-265, -100, 128, 32768}
    _ = v5
    _ = v6
    return -175
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 uint64 = 783
    switch v1 {
    case "\n\t":
        v2 = (42 < 42)
    default:
        var v7 int16 = -1
        _ = v7
    }
    if v2 {
        for fi8 := 0; fi8 < 6; fi8++ {
            _ = len(v3)
            _ = fi8
        }
    }
    var v9 int64 = -9223372036854775808
    v9 += v9
    switch "foo bar" {
    case "hello":
        for fi10 := 0; fi10 < 9; fi10++ {
            if v2 {
                v11 := GenS1{F0: fi10}
                _ = v11
            } else {
                switch genFn0() {
                case 128:
                    var v12 float32 = float32(float64(1e10))
                    _ = v12
                }
                var v13 int64 = v9
                _ = v13
            }
            var v14 uint = 577
            _ = fi10
            _ = v14
        }
    case "a":
        if v2 {
            v15 := GenS1{F0: 900}
            _ = v15
        } else {
            fmt.Println(v9, v2)
            v16 := GenS1{F0: 9223372036854775807}
            if true {
                v1 = ("\n\t" + v1)
                v17 := []int{32767, 9223372036854775807, -2147483648, genFn0()}
                _ = v17
            }
            _ = v16
        }
    default:
        var v18 int16 = 219
        _ = v18
    }
    var v19 int64 = -1
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(fi20, fi20)
        v21 := GenS2{F0: 2, F1: 2147483647}
        _ = fi20
        _ = v21
    }
    var v22 int = genFn0()
    if false {
        v4 -= 575
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
