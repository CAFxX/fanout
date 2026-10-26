package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 int8, p1 int8, p2 uint32) int32 {
    p1 = -29
    if true {
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := []int{}
            var v7 string = ("foo bar" + "")
            var v8 uint8 = 128
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        if false {
            var v9 string = (("" + "\n\t") + "")
            switch 32767 {
            case 42:
                if (!(1.5 >= 0.0)) {
                    p1 += (^(p0))
                    switch -2147483648 {
                    case -351:
                        fmt.Println(p1, p0)
                    }
                    v9 = (v9 + "\n\t")
                } else {
                    v10 := GenS1{F0: false, F1: -1}
                    if false {
                        v11 := GenS2{F0: 1000}
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 9; fi12++ {
                        fmt.Println(p0, v9, v9)
                        _ = fi12
                    }
                    _ = v10
                }
            case 0:
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := []int{(fi13 >> ((42) & 63)), (fi13 - 127), 60, fi13}
                    _ = fi13
                    _ = v14
                }
            default:
                v15 := GenS3{F0: 32767, F1: 162}
                _ = v15
            }
            _ = v9
        } else {
            var v16 uint8 = 189
            _ = v16
        }
    }
    v17 := []int{950, 309, -2147483648, 838}
    p2 = 2
    var v18 uint8 = 219
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return -1
}

func main() {
    var v1 int64 = 0
    var v2 int64 = v1
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v19 uint32 = 983
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{(fi20 << ((42) & 63))}
        for fi22 := 0; fi22 < 7; fi22++ {
            v2 = v1
            v23 := v4[1]
            v1 = (v2 & (v2 + v1))
            _ = fi22
            _ = v23
        }
        fmt.Println(v2, v19)
        _ = fi20
        _ = v21
    }
    fmt.Println(v1, v1)
    var v24 int8 = 0
    delete(v4, 900)
    var v25 int16 = 32767
    v19 = (v19 + 4294967295)
    _ = len(v4)
    var v26 int64 = 852
    switch "a" {
    case "hello":
        if true {
            v25 -= (-(v25))
        }
    default:
        var v27 uint64 = 42
        _ = v27
    }
    v28 := GenS3{F0: v25, F1: (v19 ^ (v19 / ((277) | 1)))}
    v19 -= v19
    v29 := GenS3{F0: 0, F1: (^(v19))}
    _ = len(v3)
    for fi30 := 0; fi30 < 11; fi30++ {
        fmt.Println(v25)
        _ = fi30
    }
    v1 += v26
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
