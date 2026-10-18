package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }

func main() {
    var v1 int8 = -1
    v2 := []int{127, -2147483648, 985, 32767}
    v3 := make(map[int]int)
    if (-9223372036854775808 == 643) {
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
        v1 -= v1
    } else {
        var v4 int8 = (^(v1))
        _ = v4
    }
    v1 -= v1
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = -476
        var v7 bool = false
        if false {
            v1 -= v1
            v7 = false
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int16 = 32767
                _ = fi8
                _ = v9
            }
            var v10 int16 = 32767
            _ = len(v2)
            _ = v10
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v11 int8 = 100
    fmt.Println(v11, v11)
    _ = len(v3)
    fmt.Println(v11, v1)
    switch 9223372036854775807 {
    case 19:
        var v12 bool = false
        _ = v12
    case 42:
        for fi13 := 0; fi13 < 13; fi13++ {
            delete(v3, fi13)
            _ = fi13
        }
    case 128:
        var v14 string = ("hello" + ("foo bar" + "a"))
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint16 = 65535
            v17 -= v17
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: -9223372036854775808, F1: 567, F2: 2}
        _ = len(v3)
        _ = fi15
        _ = v18
    }
    var v19 uint16 = 191
    v2 = append(v2, -700)
    switch -2147483648 {
    case 2147483647:
        var v20 uint = 318
        _ = v20
    case -1:
        v19 = (v19 % (((v19 % ((255) | 1))) | 1))
    case -512:
        v1 += v1
    default:
        if false {
            v19 = v19
        }
    }
    var v21 int8 = (v11 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
