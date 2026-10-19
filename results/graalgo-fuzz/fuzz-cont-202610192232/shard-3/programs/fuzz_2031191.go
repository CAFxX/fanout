package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
func genFn0(p0 int) uint32 {
    var v7 int64 = -818
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: "foo bar", F1: "a"}
        for fi10 := 0; fi10 < 6; fi10++ {
            v7 -= -1
            fmt.Println(fi8, fi8)
            v7 = v7
            _ = fi10
        }
        v11 := []int{}
        _ = fi8
        _ = v11
        _ = v9
    }
    fmt.Println(v7, v7)
    var v12 int64 = (v7 >> ((109) & 63))
    _ = p0
    _ = v12
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16) int32 {
    v13 := GenS2{F0: ("a" + ""), F1: -689, F2: 1}
    switch "hello" {
    case "foo bar":
        var v14 float64 = (float64(-2.25) / 1e10)
        _ = v14
    case "hello":
        fmt.Println(p0)
    default:
        if true {
            var v15 int32 = 42
            switch 128 {
            case 58:
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 uint8 = 128
                    p0 = 717
                    _ = fi16
                    _ = v17
                }
            case 882:
                if true {
                    v15 = 1
                    var v18 int32 = v15
                    _ = v18
                } else {
                    if true {
                        var v19 uint64 = 45
                        _ = v19
                    }
                }
            case 990:
                fmt.Println(p0, v15)
            default:
                v15 += (-(v15))
            }
            _ = v15
        }
    }
    switch 2147483647 {
    case 42:
        v20 := GenS3{F0: 42, F1: -396, F2: 1}
        _ = v20
    default:
        fmt.Println(p0)
    }
    v21 := []int{937}
    var v22 string = "\n\t"
    _ = p0
    _ = v13
    _ = v21
    _ = v22
    return 554
}

func main() {
    var v1 float32 = 0.1
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    var v5 uint = 709
    v6 := make(map[string]string)
    var v23 int8 = -55
    v3[685] = 95
    fmt.Println(v4, v1)
    var v24 int64 = v2
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := []int{-2147483648, fi25, 2147483647}
        _ = fi25
        _ = v26
    }
    delete(v3, genFn0(1))
    v27 := v3[2147483647]
    var v28 float64 = -2.25
    var v29 uint = (v5 % ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
