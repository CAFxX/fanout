package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    v1 := []int{-2147483648, 154}
    var v2 float64 = (3.14159 + (-2.25 + (((0.0 - (1e10 * ((3.14159 - 1.5) * -2.25))) / (1e10 / -2.25)) - 3.14159)))
    var v3 string = (((("hello" + ("" + "\n\t")) + ("\n\t" + "hello")) + "") + "foo bar")
    v4 := make(map[uint32]int)
    var v5 int8 = -128
    switch 42 {
    case 245:
        var v6 int8 = (-(v5))
        _ = v6
    case 0:
        var v7 float64 = v2
        _ = v7
    case -9223372036854775808:
        var v8 uint16 = 821
        _ = v8
    }
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            v12 := GenS1{F0: 0, F1: -171}
            var v13 uint32 = 2
            var v14 float32 = (0.0 - 1e10)
            _ = v12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            if true {
                if (124 < v5) {
                    var v16 int8 = (v5 & v5)
                    _ = v16
                }
            } else {
                var v17 int64 = 9223372036854775807
                if len(v1) > 0 {
                    v1[0] = fi15
                }
                _ = v17
            }
            _ = fi15
        }
        _ = fi11
    }
    switch -917 {
    case -2147483648:
        var v18 float32 = float32(((1e10 / 3.14159) + v2))
        _ = v18
    }
    var v19 int64 = 1
    v5 += 40
    v20 := GenS1{F0: 0, F1: -221}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
