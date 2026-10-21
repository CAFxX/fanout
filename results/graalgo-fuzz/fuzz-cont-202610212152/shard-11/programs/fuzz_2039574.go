package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint32 = 519
    v5 := []string{(v1 + v1)}
    v1 = ((v1 + v1) + v1)
    var v6 int8 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v4)
        _ = fi7
    }
    fmt.Println(v1, v6, v1)
    var v8 float64 = 0.1
    var v9 string = ("hello" + v1)
    if true {
        var v10 int = 42
        for fi11 := 0; fi11 < 5; fi11++ {
            switch 4294967295 {
            case 2147483647:
                v12 := GenS1{F0: 401, F1: uint16(v4)}
                _ = v12
            case 779:
                v1 = "hello"
            case 32768:
                v13 := v2[4294967295]
                _ = v13
            default:
                v8 += v8
            }
            v1 = v1
            var v14 int16 = -1
            _ = fi11
            _ = v14
        }
        _ = v10
    }
    v5 = append(v5, ("foo bar" + v9))
    var v15 int8 = (v6 | v6)
    var v16 int16 = 1000
    var v17 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
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
