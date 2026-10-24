package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.1
    v3 := []string{("hello" + "\n\t")}
    var v4 bool = false
    delete(v1, 42)
    var v5 uint32 = 1
    v2 += v2
    fmt.Println(v2, v5)
    v6 := GenS1{F0: 1, F1: (v5 | v5)}
    var v7 uint16 = 42
    v8 := GenS1{F0: 1, F1: v5}
    var v9 uint16 = v7
    if v4 {
        if ((!false) && (1e10 != float64((v2 - (-2.25 / v2))))) {
            switch ("a" + ("\n\t" + "hello")) {
            case "a":
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(v4, v2)
                    _ = fi10
                }
            case "foo bar":
                var v11 uint32 = 338
                _ = v11
            case "hello":
                fmt.Println(v2, v5)
            }
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            if v4 {
                for fi13 := 0; fi13 < 20; fi13++ {
                    v9 += 2
                    if (v4 || v4) {
                        var v14 uint64 = 2
                        v4 = (601 <= -2147483648)
                        _ = v14
                    }
                    _ = fi13
                }
                v3 = append(v3, "\n\t")
            }
            fmt.Println(v5, fi12)
            v9 -= 2
            _ = fi12
        }
    } else {
        fmt.Println(v7, v9)
        v15 := v1[(v5 - (v5 * (v5 * (v5 | (v5 - v5)))))]
        var v16 int = 42
        _ = v15
        _ = v16
    }
    _ = len(v3)
    var v17 int8 = -25
    v18 := GenS2{F0: 722, F1: -1, F2: false}
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 uint64 = 1
        v5 = (v5 % (((v5 >> ((0) & 63))) | 1))
        _ = fi19
        _ = v20
    }
    fmt.Println(v5, v9, v17)
    var v21 float64 = (1e10 - ((1e10 * float64(0.0)) - float64(v2)))
    var v22 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
