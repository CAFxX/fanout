package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint32, p2 bool) uint32 {
    v5 := []int{128, 128}
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 uint16 = 600
    var v2 uint = uint(v1)
    var v3 int64 = 1
    var v4 string = ("foo bar" + "\n\t")
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = 1e10
        _ = fi6
        _ = v7
    }
    switch 943 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{fi8, (fi8 * fi8), (fi8 | (fi8 * fi8))}
            if (v3 == v3) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint8 = 255
                    v12 := GenS1{F0: v11, F1: (v2 & (v2 % ((65535) | 1))), F2: 42}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            fmt.Println(v1, v2)
            _ = fi8
            _ = v9
        }
    case 128:
        v13 := []int{-719, -2147483648, -421}
        _ = v13
    case 1:
        v4 = "hello"
    default:
        v14 := GenS2{F0: (v3 * v3)}
        _ = v14
    }
    v2 = (v2 % ((0) | 1))
    var v15 int64 = 42
    fmt.Println(v3, v4, v1)
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int32 = 2147483647
            v19 := []int{fi16}
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := []int{4294967295}
        v20 = append(v20, 2147483647)
        _ = fi16
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint = 1
        _ = fi21
        _ = v22
    }
    fmt.Println(v3)
    v3 = (v3 % ((v15) | 1))
    if false {
        var v23 uint16 = v1
        fmt.Println(v2)
        _ = v23
    } else {
        var v24 uint32 = 42
        var v25 bool = (0 == (v3 / ((-34) | 1)))
        var v26 float64 = 0.1
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
