package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }

func main() {
    var v1 string = ""
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    v6 := make(map[int]uint32)
    var v7 int = 127
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 float32 = float32(1.5)
        v10 := GenS1{F0: 1, F1: false, F2: 18446744073709551615}
        var v11 int16 = -32768
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int8 = -128
    v13 := GenS1{F0: int16(v7), F1: true, F2: 2}
    v14 := GenS1{F0: 42, F1: (true || ((-(v2)) == v2)), F2: 0}
    var v15 uint = 472
    var v16 int64 = v2
    v4 = ((v4 + float32(1.5)) / 1.0)
    var v17 int16 = 32767
    var v18 uint = (v15 + (v15 | v15))
    v19 := GenS1{F0: 1000, F1: ((v12 % (((-(v12))) | 1)) == (v12 | (v12 ^ (v12 ^ (v12 & (v12 << ((42) & 63))))))), F2: 18446744073709551615}
    v20 := GenS1{F0: 42, F1: false, F2: 2}
    var v21 float64 = 1.5
    fmt.Println(v17)
    switch v1 {
    case "":
        v22 := []int{v7, (v7 ^ v7), (v7 & v7), (v7 / ((v7) | 1))}
        _ = v22
    case "foo bar":
        var v23 uint32 = 2147483647
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
