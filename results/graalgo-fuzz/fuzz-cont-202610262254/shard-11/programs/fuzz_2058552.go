package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    v1 := []int{}
    var v2 float64 = (3.14159 - (1.5 / 1.0))
    var v3 bool = true
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v4["a"] = 364
    var v6 int64 = -331
    fmt.Println(v2, v3, v3)
    v2 = (v2 + v2)
    if ((-2.25 + -2.25) == float32(((-2.25 / -2.25) - -2.25))) {
        v6 -= (^(v6))
    } else {
        if (997 < 42) {
            var v7 int16 = -1
            _ = v7
        } else {
            for fi8 := 0; fi8 < 11; fi8++ {
                fmt.Println(v6, v6)
                v1 = append(v1, fi8)
                var v9 int32 = -2147483648
                _ = fi8
                _ = v9
            }
            v3 = false
        }
        v6 -= 42
    }
    var v10 int8 = 127
    switch "\n\t" {
    case "foo bar":
        v2 = v2
    case "hello":
        v3 = v3
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int16 = -32768
        fmt.Println(v6, v6, v6)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint = 770
        if ((v14 % (((^(v14))) | 1)) <= (v14 >> (((v14 + v14)) & 63))) {
            _ = len(v4)
            fmt.Println(fi13, v3)
        }
        if true {
            var v15 uint32 = 359
            v16 := GenS1{F0: -1, F1: v3, F2: 622}
            v17 := GenS2{F0: "\n\t", F1: 127, F2: 42}
            _ = v15
            _ = v16
            _ = v17
        } else {
            if ((v14 + v14) > 90) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = fi13
        _ = v14
    }
    var v18 int16 = -32768
    var v19 bool = false
    v18 -= -32768
    v10 += (-(v10))
    v20 := v5[880]
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
