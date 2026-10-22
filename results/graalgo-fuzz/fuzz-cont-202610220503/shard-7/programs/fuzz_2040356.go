package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 uint16) int32 {
    var v4 int16 = 222
    v4 -= (-(v4))
    var v5 uint = 146
    v5 = 2
    v6 := []int{9223372036854775807, 333, 32767}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []int{947, -559}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    var v7 int32 = -688
    v3[328] = 32767
    v1 = append(v1, 335)
    var v8 string = ("foo bar" + "hello")
    v8 = v8
    if (false && false) {
        var v9 float32 = 0.1
        _ = v9
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: (^(v7)), F1: 0}
            v2 = 9223372036854775807
            if (74 != 42) {
                v7 -= v7
            } else {
                v8 = "foo bar"
            }
            _ = fi10
            _ = v11
        }
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    }
    if false {
        v12 := v3[-2147483648]
        _ = v12
    } else {
        var v13 int64 = 42
        _ = v13
    }
    var v14 bool = (563 == 0)
    for fi15 := 0; fi15 < 9; fi15++ {
        v8 = "a"
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
