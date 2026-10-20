package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 uint64, p1 int16) int {
    p1 = p1
    var v5 float32 = (((((3.14159 + (1e10 * -2.25)) * -2.25) / 1.5) / 1.0) + 0.0)
    var v6 uint64 = p0
    switch "foo bar" {
    case "hello":
        var v7 string = ""
        _ = v7
    default:
        fmt.Println(v5, v6, v5)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    var v4 string = v1
    v8 := GenS1{F0: 42, F1: 565}
    fmt.Println(v1)
    v9 := []int{}
    v3 = 42
    v3 += (v3 % ((v3) | 1))
    if false {
        if (657 != -924) {
            v3 -= 0
        } else {
            delete(v2, 42)
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v3 = 255
                v1 = (v4 + ("hello" + "\n\t"))
                _ = fi11
            }
            _ = fi10
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS2{F0: 2147483647, F1: 396, F2: -1}
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 2, F1: 832}
        _ = v14
    }
    v4 = "a"
    v4 = v4
    switch v4 {
    case "foo bar":
        v15 := v2[25]
        _ = v15
    case "hello":
        _ = len(v2)
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS3{F0: 100, F1: -1}
        _ = fi16
        _ = v17
    }
    if false {
        v1 = v4
    }
    var v18 int16 = -32768
    v19 := GenS3{F0: 0, F1: 69}
    var v20 string = "foo bar"
    v3 = v3
    v21 := v2[0]
    v22 := GenS1{F0: 2, F1: 895}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
