package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    var v1 bool = (!false)
    var v2 uint64 = 173
    var v3 uint8 = 61
    v4 := make(map[uint32]uint32)
    var v5 uint = 881
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: "hello", F1: -32768}
    var v8 int8 = 42
    v8 += v8
    switch "" {
    case "":
        v9 := v6[127]
        _ = v9
    case "foo bar":
        var v10 int32 = 1
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v8, v8)
            var v12 int64 = 1
            v8 = v8
            _ = fi11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 11; fi13++ {
            v1 = v1
            if (((1.5 < 1.5) && (501 != 2)) || v1) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS1{F0: (("foo bar" + "\n\t") + ("a" + ("foo bar" + "hello"))), F1: 1}
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 string = "\n\t"
                    fmt.Println(v8, v2, v2)
                    var v18 uint32 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi13
        }
    }
    v19 := GenS2{F0: v1, F1: 42}
    v5 = v5
    if v1 {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := []int{(fi20 / (((fi20 >> (((v5 & uint(v3))) & 63))) | 1))}
            v1 = v1
            _ = fi20
            _ = v21
        }
        v22 := GenS3{F0: v2, F1: 0}
        var v23 bool = true
        _ = v22
        _ = v23
    } else {
        v24 := GenS1{F0: "foo bar", F1: -32768}
        v2 -= (^(v2))
        _ = v24
    }
    v3 += 1
    var v25 int64 = -70
    if v1 {
        var v26 float64 = -2.25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
