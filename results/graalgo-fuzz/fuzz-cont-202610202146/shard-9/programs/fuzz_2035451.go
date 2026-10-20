package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() uint16 {
    var v4 uint16 = 255
    var v5 uint32 = 2147483648
    v4 = v4
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    fmt.Println(v1)
    var v6 int64 = 0
    v7 := []int{-1, int(v6), 127}
    if ("hello" != ("hello" + "\n\t")) {
        switch -2147483648 {
        case 127:
            v8 := GenS1{F0: 42, F1: 255}
            _ = v8
        case 2147483647:
            v6 -= (v6 % ((v6) | 1))
        case -1:
            switch 42 {
            case -219:
                v9 := GenS1{F0: 1, F1: 42}
                _ = v9
            }
        }
        var v10 int8 = 48
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        switch fi11 {
        case 1:
            for fi12 := 0; fi12 < 8; fi12++ {
                delete(v3, fi11)
                v6 -= (v6 + -188)
                v13 := GenS1{F0: 2147483647, F1: 255}
                if true {
                    v14 := GenS2{F0: 118, F1: 255}
                    var v15 uint = 0
                    _ = v14
                    _ = v15
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                v16 := GenS1{F0: -2147483648, F1: 42}
                v1 = true
                _ = fi12
                _ = v13
                _ = v16
            }
        case 42:
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 int = (fi17 >> ((0) & 63))
                _ = len(v2)
                if v1 {
                    var v19 uint = 1
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
        case 769:
            v1 = (0 <= -10)
        default:
            v6 = 746
        }
        v20 := GenS2{F0: 42, F1: 936}
        _ = fi11
        _ = v20
    }
    if v1 {
        if true {
            v7 = append(v7, 42)
            fmt.Println(v6, v6)
            _ = len(v7)
        }
    }
    fmt.Println(v6, v1)
    v1 = (65535 != 510)
    v21 := GenS2{F0: 1, F1: 738}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
}
