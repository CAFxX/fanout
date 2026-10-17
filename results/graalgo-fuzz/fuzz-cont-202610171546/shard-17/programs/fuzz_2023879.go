package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 9223372036854775807
    var v3 string = "\n\t"
    var v4 uint8 = 42
    var v5 uint8 = v4
    v6 := make(map[uint32]uint32)
    var v7 int8 = 127
    if true {
        fmt.Println(v3)
        var v8 int64 = 42
        _ = v8
    }
    var v9 bool = false
    fmt.Println(v3, v4, v2)
    v10 := GenS1{F0: -2147483648, F1: 0}
    fmt.Println(v4)
    v11 := []int{42}
    v3 = "a"
    v12 := GenS2{F0: (v5 ^ (^(v4))), F1: -761}
    v13 := GenS3{F0: 25}
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            var v15 uint16 = 32767
            _ = v15
        }
        _ = len(v11)
        var v16 int64 = int64(fi14)
        _ = fi14
        _ = v16
    }
    if v9 {
        var v17 int = 32768
        v18 := GenS1{F0: -931, F1: v17}
        v19 := GenS4{F0: 984, F1: v9}
        _ = v17
        _ = v18
        _ = v19
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v7 = (v7 % (((v7 & (v7 | v7))) | 1))
    v7 = (v7 * 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
