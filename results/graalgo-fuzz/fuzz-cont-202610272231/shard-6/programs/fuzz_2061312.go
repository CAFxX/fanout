package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0() uint8 {
    var v5 float32 = 1.5
    v5 -= (((3.14159 * v5) / v5) / 1.5)
    fmt.Println(v5)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int64 = 732
        v5 = v5
        _ = fi6
        _ = v7
    }
    _ = v5
    return 187
}

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 1
    var v3 float32 = (1e10 * 1.5)
    var v4 uint = 65535
    v8 := []int{-9223372036854775808}
    var v9 uint = (v4 * uint(v2))
    v10 := GenS1{F0: 535, F1: 221}
    switch "" {
    case "hello":
        fmt.Println(v2, v4)
    case "":
        var v11 int64 = 0
        _ = v11
    default:
        var v12 uint32 = 0
        _ = v12
    }
    switch "\n\t" {
    case "foo bar":
        var v13 float64 = ((0.1 + -2.25) - (1.5 / 0.1))
        _ = v13
    default:
        switch -9223372036854775808 {
        case 4294967295:
            var v14 int64 = 42
            _ = v14
        }
    }
    switch "foo bar" {
    case "hello":
        switch "\n\t" {
        case "hello":
            if false {
                v15 := GenS1{F0: (v1 - v1), F1: 1000}
                v1 -= v1
                v16 := GenS1{F0: (^(v1)), F1: -32768}
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    fmt.Println(v4)
                    _ = fi17
                }
                v3 -= ((v3 + v3) / v3)
                for fi18 := 0; fi18 < 8; fi18++ {
                    v1 -= 65535
                    _ = fi18
                }
            }
        }
    }
    v19 := GenS2{F0: 483, F1: (v1 >> ((83) & 63))}
    var v20 int64 = 1
    var v21 int = 9223372036854775807
    for fi22 := 0; fi22 < 15; fi22++ {
        v8 = append(v8, (-(fi22)))
        _ = fi22
    }
    v23 := GenS3{F0: 65535, F1: 32767}
    fmt.Println(v21, v20, v4)
    v2 = (v2 ^ (v2 & v2))
    var v24 int8 = 1
    var v25 int32 = -2147483648
    var v26 uint64 = 624
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
