package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 string; F1 uint; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32) uint {
    p0 = p0
    p0 = p0
    _ = p0
    return 65535
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 42
    var v3 int16 = (v1 % ((0) | 1))
    var v4 int16 = v1
    fmt.Println(v3, v4, v3)
    v5 := []int{}
    if (!true) {
        fmt.Println(v3, v2)
        v6 := GenS1{F0: "a"}
        _ = v6
    }
    if true {
        v7 := GenS2{F0: 18446744073709551615, F1: -2147483648}
        v8 := GenS3{F0: 2147483647}
        switch "\n\t" {
        case "\n\t":
            v4 = v4
        case "":
            var v9 float64 = 1e10
            _ = v9
        case "a":
            if true {
                v2 += (v2 * 42)
            } else {
                v4 += 1
                fmt.Println(v3)
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 int64 = -9223372036854775808
                    if true {
                        v3 -= int16(fi10)
                        v12 := GenS3{F0: 504}
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v14 := GenS4{F0: 119, F1: 635}
                        var v15 uint64 = 0
                        v16 := GenS5{F0: "\n\t", F1: genFn0(4294967295), F2: 42}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = v7
        _ = v8
    }
    var v17 int = 252
    fmt.Println(v1, v2)
    var v18 float32 = float32(3.14159)
    v18 = v18
    if false {
        _ = len(v5)
        v17 = -57
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v4, v2)
        _ = fi19
    }
    v1 -= (v3 & (v1 & (v1 - (^(v1)))))
    v18 = v18
    if true {
        v20 := GenS6{F0: 100, F1: 18446744073709551615}
        fmt.Println(v3, v18, v18)
        _ = v20
    } else {
        fmt.Println(v2)
        v3 += (v3 & (v3 - (^(v1))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
