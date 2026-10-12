package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint64 = 1
    var v4 uint8 = 1
    var v5 uint16 = 42
    v6 := GenS1{F0: false}
    v4 = 1
    v3 += v3
    if true {
        var v7 uint = uint(v4)
        v2[(("foo bar" + "a") + ("" + ""))] = v1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := v2[""]
            v4 -= (v4 / ((v4) | 1))
            fmt.Println(v3, v1)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v3)
            _ = fi10
        }
        fmt.Println(v4, v4, v3)
    }
    v11 := GenS1{F0: ((v3 | (v3 * (^(v3)))) > 2)}
    switch "foo bar" {
    case "hello":
        v5 = 32767
    case "a":
        fmt.Println(v5)
    case "foo bar":
        v3 += v3
    }
    v12 := GenS1{F0: (((v3 - (v3 / ((v3) | 1))) != (v3 & (v3 & (^(v3))))) || true)}
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int = (fi13 / ((fi13) | 1))
        var v15 float64 = (-2.25 / 0.1)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint16 = v5
    v3 = v3
    var v17 uint32 = 2
    var v18 uint = 4294967295
    switch 32767 {
    case 9223372036854775807:
        var v19 float32 = 0.0
        _ = v19
    case 0:
        v4 = v4
    case -584:
        for fi20 := 0; fi20 < 7; fi20++ {
            v5 += v16
            v1 += v17
            _ = fi20
        }
    }
    switch "foo bar" {
    case "hello":
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := []int{fi21, (fi21 | (^(fi21)))}
            v23 := GenS2{F0: v3, F1: "\n\t", F2: -2147483648}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v18 = (v18 * v18)
    var v24 int16 = 42
    v25 := GenS3{F0: 2, F1: v24, F2: 9223372036854775807}
    v26 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
