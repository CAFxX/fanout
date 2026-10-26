package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    var v1 int8 = 1
    var v2 uint64 = 997
    var v3 uint16 = 65535
    var v4 uint16 = v3
    fmt.Println(v4, v2, v3)
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                v4 = v3
                _ = fi6
            }
            v1 = v1
            v1 += 6
            _ = fi5
        }
        v7 := []int{32768}
        _ = v7
    }
    var v8 int = 78
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    var v11 uint = 796
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint8 = 0
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 255
    var v16 uint8 = (v15 & 1)
    v2 = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint8 = (v15 * (v15 / ((v16) | 1)))
        _ = fi17
        _ = v18
    }
    fmt.Println(v8, v11, v11)
    v19 := GenS1{F0: (v8 << ((v11) & 63))}
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS2{F0: (v2 * (v2 | (v2 + 42))), F1: (!false)}
        fmt.Println(v16, v11, v15)
        v22 := []int{-918, fi20}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch -9223372036854775808 {
    case -696:
        var v23 uint64 = (v2 | 968)
        _ = v23
    case -2147483648:
        v4 += (v4 << ((941) & 63))
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v4, v4)
        fmt.Println(v16)
        var v25 uint64 = v2
        _ = fi24
        _ = v25
    }
    switch (v8 * v8) {
    case 801:
        var v26 uint16 = 0
        _ = v26
    case -251:
        v27 := []int{}
        _ = v27
    case 42:
        fmt.Println(v11, v16, v1)
    default:
        fmt.Println(v16, v2, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
