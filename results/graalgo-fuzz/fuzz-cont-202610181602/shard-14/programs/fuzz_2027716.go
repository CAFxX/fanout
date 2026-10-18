package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }

func main() {
    var v1 uint8 = 0
    var v2 uint8 = v1
    var v3 uint64 = 18446744073709551615
    var v4 int8 = 1
    var v5 int8 = 1
    switch 32767 {
    case 0:
        v6 := GenS1{F0: v3, F1: (v2 % ((v2) | 1))}
        _ = v6
    default:
        fmt.Println(v3)
    }
    v7 := []int{1, 403, 416}
    var v8 float32 = -2.25
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint = 255
        v4 += v5
        _ = fi9
        _ = v10
    }
    switch -9223372036854775808 {
    case 128:
        switch 972 {
        case 9223372036854775807:
            v4 = (v4 / ((v4) | 1))
        case 127:
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v8 = 3.14159
                    var v13 float32 = v8
                    switch -797 {
                    case 4294967295:
                        v14 := GenS2{F0: 0}
                        _ = v14
                    case -830:
                        switch "a" {
                        case "foo bar":
                            var v15 int = (fi11 ^ fi12)
                            _ = v15
                        case "a":
                            v13 = 0.0
                        }
                    default:
                        v16 := GenS1{F0: (v3 << ((42) & 63)), F1: 42}
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        default:
            v5 = v5
        }
    case -2147483648:
        switch "\n\t" {
        case "hello":
            v17 := GenS2{F0: 2147483647}
            _ = v17
        }
    case 9223372036854775807:
        v5 = (-(v5))
    default:
        if len(v7) > 0 {
            v7[0] = int(v3)
        }
    }
    v18 := GenS3{F0: 32767, F1: true, F2: -202}
    v19 := GenS3{F0: 789, F1: ((-(v4)) > v5), F2: 857}
    var v20 uint = 67
    v21 := GenS2{F0: -1}
    var v22 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
