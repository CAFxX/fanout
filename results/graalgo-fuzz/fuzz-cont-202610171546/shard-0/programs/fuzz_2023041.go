package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 int) int32 {
    v5 := GenS1{F0: false, F1: 76}
    p0 = p0
    switch "" {
    case "a":
        var v6 uint32 = 0
        _ = v6
    }
    _ = p0
    _ = v5
    return -706
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := make(map[uint32]int)
    var v7 bool = false
    _ = len(v2)
    v8 := []int{-1, 42, 2147483647}
    _ = len(v8)
    if (128 <= 1) {
        var v9 float32 = 1.5
        v3 += v3
        switch v1 {
        case "\n\t":
            fmt.Println(v7, v1)
        case "":
            switch -2147483648 {
            case 32768:
                v10 := GenS1{F0: (genFn0(32767) != -1), F1: 835}
                _ = v10
            case 127:
                var v11 uint = 2
                _ = v11
            default:
                var v12 float64 = float64(3.14159)
                _ = v12
            }
        default:
            v1 = "hello"
        }
        _ = v9
    }
    if (!v7) {
        var v13 string = v1
        var v14 int16 = 32767
        v2[2147483648] = 4294967295
        _ = v13
        _ = v14
    } else {
        _ = len(v8)
        for fi15 := 0; fi15 < 4; fi15++ {
            if (!v7) {
                v3 = (^(v3))
                v7 = false
            } else {
                if false {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        fmt.Println(v7)
                        _ = fi16
                    }
                    _ = len(v8)
                    for fi17 := 0; fi17 < 9; fi17++ {
                        if v7 {
                            var v18 uint32 = 2147483648
                            v1 = v1
                            _ = v18
                        }
                        fmt.Println(fi15)
                        _ = fi17
                    }
                }
                _ = len(v8)
            }
            _ = fi15
        }
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = -167
        _ = fi19
        _ = v20
    }
    var v21 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
