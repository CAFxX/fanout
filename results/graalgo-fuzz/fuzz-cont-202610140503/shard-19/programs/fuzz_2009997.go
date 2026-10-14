package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 42
    var v2 bool = (0 <= 1)
    var v3 uint8 = 1
    v4 := make(map[int]int)
    if v2 {
        fmt.Println(v3)
    }
    fmt.Println(v2, v1, v3)
    var v5 uint16 = 255
    var v6 int16 = 843
    v7 := GenS1{F0: 53, F1: true}
    var v8 string = "a"
    switch v8 {
    case "\n\t":
        var v9 uint16 = uint16(v1)
        _ = v9
    case "hello":
        var v10 int16 = -783
        _ = v10
    case "a":
        v5 = (v5 << ((0) & 63))
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint32 = 2147483648
                _ = fi12
                _ = v13
            }
            v6 += v6
            switch v8 {
            case "a":
                if v2 {
                    var v14 float32 = (((0.1 + (-2.25 / (-2.25 + 1.5))) * 1.5) + 1e10)
                    _ = v14
                } else {
                    var v15 bool = (79 == 255)
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v6 -= (v6 % ((1000) | 1))
                        if (65535 > 0) {
                            v17 := GenS1{F0: 42, F1: ((v8 + (v8 + (v8 + "\n\t"))) > v8)}
                            v1 = (v1 - (v1 % (((v1 ^ v1)) | 1)))
                            var v18 uint8 = v3
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS1{F0: 18, F1: true}
                        _ = fi16
                        _ = v19
                    }
                    _ = v15
                }
            default:
                var v20 bool = (v1 <= v1)
                _ = v20
            }
            _ = fi11
        }
    }
    v21 := []int{}
    delete(v4, -1)
    v2 = v2
    if false {
        v8 = v8
    } else {
        v22 := GenS2{F0: 0}
        var v23 int32 = -1
        var v24 float64 = (((1e10 / 0.1) + 1e10) - 1.5)
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
