package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0() int32 {
    switch "hello" {
    case "hello":
        var v4 uint16 = 778
        _ = v4
    case "foo bar":
        switch -9223372036854775808 {
        case -1:
            if false {
                var v5 int64 = -285
                v5 = (v5 / ((258) | 1))
                _ = v5
            } else {
                if (!false) {
                    var v6 uint8 = 209
                    var v7 string = "\n\t"
                    var v8 int64 = -1
                    _ = v6
                    _ = v7
                    _ = v8
                }
            }
        }
    case "\n\t":
        var v9 float32 = 0.0
        _ = v9
    }
    var v10 float32 = 3.14159
    fmt.Println(v10)
    _ = v10
    return -672
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{}
    var v3 float32 = 3.14159
    var v11 int16 = -1
    v1 += v1
    if len(v2) > 0 {
        v2[0] = 596
    }
    var v12 int8 = -121
    v13 := GenS1{F0: 9223372036854775807, F1: 316}
    switch -455 {
    case 4294967295:
        fmt.Println(v12, v3, v12)
    case 360:
        fmt.Println(v3)
    }
    var v14 uint64 = 0
    fmt.Println(v12)
    switch -191 {
    case 190:
        switch "\n\t" {
        case "":
            for fi15 := 0; fi15 < 18; fi15++ {
                _ = len(v2)
                var v16 int8 = (v12 % (((v12 / ((0) | 1))) | 1))
                v17 := GenS2{F0: 0}
                _ = fi15
                _ = v16
                _ = v17
            }
        case "hello":
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(v12, v12)
                for fi19 := 0; fi19 < 3; fi19++ {
                    v3 -= (v3 * v3)
                    v12 += 127
                    var v20 float32 = 0.0
                    _ = fi19
                    _ = v20
                }
                var v21 uint8 = 198
                _ = fi18
                _ = v21
            }
        }
    default:
        fmt.Println(v3, v1, v1)
    }
    v3 += v3
    var v22 bool = true
    var v23 uint64 = (v14 << ((42) & 63))
    var v24 float64 = 0.0
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
