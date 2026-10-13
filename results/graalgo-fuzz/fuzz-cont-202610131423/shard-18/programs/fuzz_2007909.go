package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float32 = 1e10
    var v2 uint8 = 42
    var v3 float32 = float32(1e10)
    var v4 int64 = 42
    v2 = (v2 ^ (^(v2)))
    fmt.Println(v4, v3)
    v3 += -2.25
    fmt.Println(v1, v2)
    fmt.Println(v3, v2)
    if (1000 > -32768) {
        v1 = float32(1.5)
    } else {
        if false {
            var v5 int16 = -32768
            for fi6 := 0; fi6 < 14; fi6++ {
                for fi7 := 0; fi7 < 10; fi7++ {
                    v8 := GenS1{F0: 1, F1: (((("\n\t" + "foo bar") + ("a" + "hello")) + "\n\t") + "hello")}
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
            _ = v5
        }
        v9 := GenS2{F0: (v4 == -328), F1: 0}
        _ = v9
    }
    switch 42 {
    case 55:
        switch 32768 {
        case 0:
            v3 -= v3
        case 1:
            v10 := []int{9223372036854775807, 4294967295}
            _ = v10
        }
    }
    v1 = -2.25
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{fi11}
        var v13 int16 = 32767
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 128 {
    case 343:
        var v14 bool = true
        _ = v14
    case -9223372036854775808:
        var v15 uint8 = 233
        _ = v15
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v1 += v1
    case -665:
        v2 += v2
    default:
        v16 := GenS3{F0: 32767}
        _ = v16
    }
    v4 += v4
    v17 := []int{2147483647, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
