package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 int8 = 50
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 375
    var v4 uint8 = 42
    v1 += v1
    var v5 int8 = (v1 - v1)
    v6 := GenS1{F0: 255, F1: 32768}
    v5 += v1
    v7 := GenS1{F0: uint16(v2), F1: 477}
    fmt.Println(v4, v4)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS2{F0: (!false)}
        var v10 uint16 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: (!false)}
    v12 := GenS1{F0: 42, F1: 127}
    switch ("hello" + "") {
    case "\n\t":
        if (!false) {
            v3 = (v3 + v3)
            v13 := []int{}
            v14 := GenS3{F0: -2147483648, F1: -1}
            _ = v13
            _ = v14
        } else {
            var v15 int16 = (v3 >> ((18446744073709551615) & 63))
            _ = v15
        }
    case "hello":
        var v16 int8 = -128
        _ = v16
    }
    var v17 uint8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
