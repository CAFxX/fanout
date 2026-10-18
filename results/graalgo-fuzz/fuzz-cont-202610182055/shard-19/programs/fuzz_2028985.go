package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int64 {
    var v7 uint16 = 255
    v7 = v7
    fmt.Println(v7)
    _ = v7
    return -407
}

func main() {
    var v1 int16 = 42
    var v2 uint = 2
    v3 := []uint32{2147483648, 2147483647, 2147483647, 1}
    v4 := []int{-2147483648, 9223372036854775807, -735}
    var v5 int64 = 42
    var v6 uint64 = 4294967295
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint = v2
        fmt.Println(fi8)
        v10 := GenS1{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = -91
    var v12 int32 = 341
    if false {
        var v13 uint16 = 255
        if false {
            v14 := GenS1{F0: (v11 % ((v11) | 1))}
            var v15 uint64 = (v6 >> (((^(v2))) & 63))
            _ = v14
            _ = v15
        }
        v2 = (v2 + v2)
        _ = v13
    }
    var v16 int = 329
    v3 = append(v3, 2)
    var v17 int8 = (v11 | v11)
    fmt.Println(v11, v16)
    fmt.Println(v16, v11, v6)
    switch -316 {
    case 32767:
        fmt.Println(v1)
    case 42:
        var v18 int = v16
        _ = v18
    case 364:
        fmt.Println(v11, v16, v16)
    default:
        var v19 int8 = v17
        _ = v19
    }
    var v20 uint8 = 134
    switch "hello" {
    case "a":
        v21 := GenS1{F0: v11}
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v20 = v20
                fmt.Println(v12)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    fmt.Println(v12, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
