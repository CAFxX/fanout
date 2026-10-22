package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64) string {
    v6 := GenS1{F0: 1}
    v7 := GenS2{F0: 127, F1: p0}
    if (65535 >= 0) {
        p0 = (^(p0))
    }
    _ = p0
    _ = v6
    _ = v7
    return ("foo bar" + "")
}
func genFn1() int8 {
    v8 := GenS2{F0: 213, F1: 46}
    v9 := GenS2{F0: 197, F1: 42}
    v10 := []int{32768}
    _ = v10
    _ = v8
    _ = v9
    return -42
}

func main() {
    var v1 int16 = 700
    var v2 int64 = -932
    var v3 int = 9223372036854775807
    var v4 float32 = ((1.5 / 0.1) + (1.5 * (1.5 / ((-2.25 / 1e10) + (-2.25 - 1.5)))))
    var v5 uint16 = 2
    if true {
        var v11 uint16 = (v5 >> ((4294967295) & 63))
        v11 -= (v5 ^ v5)
        if (-86 != -1) {
            for fi12 := 0; fi12 < 1; fi12++ {
                if true {
                    v4 = v4
                    v13 := []int{-2147483648}
                    v14 := GenS3{F0: (0 >= 42)}
                    _ = v13
                    _ = v14
                }
                v15 := []int{}
                _ = fi12
                _ = v15
            }
            v1 = -901
            fmt.Println(v5, v5)
        } else {
            if true {
                var v16 float32 = v4
                fmt.Println(v3, v5, v5)
                _ = v16
            } else {
                var v17 string = "a"
                var v18 int = (v3 * (-(v3)))
                _ = v17
                _ = v18
            }
        }
        _ = v11
    }
    v2 -= (v2 ^ (v2 + (^(v2))))
    v5 = 1
    v19 := []int{(v3 | (^(v3))), v3}
    v3 -= (v3 - 2147483647)
    var v20 uint64 = 1
    fmt.Println(v20, v1, v5)
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(v20)
        var v22 uint8 = 173
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
