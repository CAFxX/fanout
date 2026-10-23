package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int; F1 int8; F2 int }

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 83
    var v3 float64 = 0.0
    v4 := make(map[int]string)
    var v5 uint8 = 127
    var v6 int64 = -1
    fmt.Println(v2)
    v7 := GenS1{F0: (v5 ^ v5), F1: (42 > 4294967295)}
    fmt.Println(v5, v1)
    v8 := GenS1{F0: 159, F1: true}
    v9 := GenS2{F0: 65535, F1: 1}
    v10 := GenS3{F0: 749, F1: (v2 + v2), F2: 42}
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v3)
        delete(v4, (fi11 % (((fi11 / (((fi11 * int(v5))) | 1))) | 1)))
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v3)
            _ = fi12
        }
        _ = fi11
    }
    var v13 uint8 = v5
    var v14 float64 = v3
    v15 := []int{-1}
    switch v1 {
    case "a":
        fmt.Println(v1, v14, v6)
    case "":
        v13 = (v5 ^ (v13 / (((v5 ^ (v5 >> ((107) & 63)))) | 1)))
    default:
        switch 435 {
        case 32768:
            _ = len(v4)
        default:
            for fi16 := 0; fi16 < 13; fi16++ {
                delete(v4, 127)
                v3 = v14
                _ = fi16
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
