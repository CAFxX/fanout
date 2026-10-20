package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 string; F2 int8 }
func genFn0() bool {
    var v4 float64 = float64(1.5)
    v4 = ((((v4 * (3.14159 - float64(-2.25))) - v4) * v4) - v4)
    v5 := GenS1{F0: "a", F1: 158, F2: 2147483647}
    v6 := []int{489, 128, 2147483647}
    _ = v4
    _ = v5
    _ = v6
    return (332 > 1)
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = -637
    var v3 int64 = -9223372036854775808
    var v7 int8 = -1
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int = (fi8 | 2147483647)
        var v10 int16 = 162
        v7 -= (v7 * v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 42
    v12 := GenS2{F0: 4294967295, F1: ("" + ("a" + ("\n\t" + "hello"))), F2: (v7 >> ((1) & 63))}
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v7, v1)
        _ = fi13
    }
    v3 += (v3 >> ((0) & 63))
    fmt.Println(v3, v3, v1)
    if true {
        v1 += v1
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint = 65535
                _ = fi15
                _ = v16
            }
            v2 = (v2 / ((-715) | 1))
            fmt.Println(v2)
            _ = fi14
        }
    }
    v3 += v3
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint = 873
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 2147483647
    var v20 int8 = (v7 / (((v7 ^ v7)) | 1))
    v11 += uint8(v1)
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := []int{(fi21 & fi21), (fi21 / (((fi21 + 32767)) | 1))}
        fmt.Println(v2, v3)
        v3 = 1
        _ = fi21
        _ = v22
    }
    if ((float64(-2.25) * float64(-2.25)) == -2.25) {
        v23 := []int{2147483647, 1, 32768}
        var v24 string = ("foo bar" + "\n\t")
        v25 := GenS2{F0: 32768, F1: v24, F2: v7}
        _ = v23
        _ = v24
        _ = v25
    } else {
        fmt.Println(v1, v19, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
