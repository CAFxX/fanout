package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
func genFn0(p0 int, p1 uint32, p2 uint) uint64 {
    p1 += 65535
    var v4 string = "a"
    fmt.Println(p0, p0, p2)
    var v5 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]uint32)
    var v3 int16 = 32767
    v6 := GenS1{F0: "", F1: 42}
    fmt.Println(v3)
    v7 := GenS1{F0: "hello", F1: 32767}
    v2[65535] = 2
    if (!(!(440 < 4294967295))) {
        v8 := GenS2{F0: (v3 >> ((477) & 63)), F1: ("" + "\n\t"), F2: (v1 + v1)}
        _ = v8
    }
    var v9 uint8 = 0
    v3 += -423
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint16 = 391
        var v12 uint8 = (^(v9))
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint8 = 61
    v13 = (v13 & (v13 - (v9 >> ((296) & 63))))
    v3 = (v3 / (((v3 << ((356) & 63))) | 1))
    switch "" {
    case "":
        fmt.Println(v13, v1, v9)
    case "foo bar":
        v13 = (v9 ^ (v13 ^ v9))
    default:
        var v14 float32 = 0.0
        _ = v14
    }
    v15 := GenS1{F0: "foo bar", F1: 0}
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int32 = (v1 & v1)
        fmt.Println(v9, v1)
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 42
    var v19 int8 = 117
    switch 128 {
    case 32767:
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                v19 = v19
                v22 := GenS1{F0: "foo bar", F1: 65535}
                v1 += v1
                switch "\n\t" {
                case "foo bar":
                    var v23 int32 = 0
                    _ = v23
                case "\n\t":
                    fmt.Println(v3, v19, v1)
                }
                v18 += 0
                _ = fi21
                _ = v22
            }
            v3 -= (v3 | v3)
            _ = fi20
        }
    case -1:
        v24 := []int{2147483647}
        _ = v24
    case 0:
        v19 = (v19 - (v19 + -128))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
