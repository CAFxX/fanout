package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 42
    var v3 int = 32768
    var v4 uint64 = 387
    v5 := make(map[int]int)
    var v6 float64 = 1.5
    v7 := GenS1{F0: 2}
    var v8 uint64 = v4
    var v9 uint64 = v4
    switch "a" {
    case "\n\t":
        var v10 int8 = -1
        _ = v10
    default:
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                if true {
                    var v13 int = fi11
                    _ = v13
                }
                var v14 int64 = 1
                v15 := GenS1{F0: 2147483647}
                _ = fi12
                _ = v14
                _ = v15
            }
            switch "a" {
            case "hello":
                var v16 int64 = 1
                _ = v16
            case "a":
                var v17 float32 = (0.1 / 1e10)
                _ = v17
            case "foo bar":
                v18 := []int{32767, fi11, v3}
                _ = v18
            default:
                v4 = (v8 ^ (v9 & (^(v8))))
            }
            _ = fi11
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v8 += (v8 - v4)
            var v21 uint16 = 32767
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v22 := []int{v3, (v3 & 1), v3}
    if ((0.1 - 1.5) >= 0.1) {
        var v23 uint64 = (v9 - (v8 & (v4 * (v9 >> ((91) & 63)))))
        if true {
            if (9223372036854775807 > 1) {
                v8 = (v8 - 643)
                var v24 uint = 255
                _ = v24
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v3 = fi25
                    _ = fi25
                }
            }
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v4, v9, v2)
            fmt.Println(fi26)
            v27 := GenS1{F0: 1}
            _ = fi26
            _ = v27
        }
        _ = v23
    }
    var v28 uint = 260
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
