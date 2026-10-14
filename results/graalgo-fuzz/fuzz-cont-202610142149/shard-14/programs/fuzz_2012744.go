package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("foo bar" + ("\n\t" + ((("\n\t" + "a") + "hello") + ("" + "\n\t")))), ("" + ("hello" + "")), "hello"}
    v3 := []int{0}
    v4 := []int{}
    switch 127 {
    case 2147483647:
        var v5 int8 = 1
        _ = v5
    case 1:
        v6 := GenS1{F0: -198, F1: 42, F2: 1}
        _ = v6
    case -935:
        var v7 int64 = -9223372036854775808
        _ = v7
    default:
        var v8 uint64 = 1
        _ = v8
    }
    var v9 float64 = float64((-2.25 - (1e10 + -2.25)))
    var v10 uint = 4294967295
    var v11 int8 = 1
    v12 := GenS1{F0: 128, F1: 42, F2: -1}
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS2{F0: 65535}
        v15 := v1[42]
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int16 = 0
    var v17 float64 = v9
    if false {
        v18 := GenS3{F0: (v16 / (((v16 >> (((v10 | (v10 - v10))) & 63))) | 1)), F1: -1, F2: 0}
        _ = v18
    } else {
        var v19 float64 = 1.5
        _ = v19
    }
    v20 := GenS4{F0: ("\n\t" + "")}
    fmt.Println(v17, v17, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
