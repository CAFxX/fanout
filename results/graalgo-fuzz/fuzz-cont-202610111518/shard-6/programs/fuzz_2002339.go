package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 uint = 42
    var v5 int16 = 32767
    switch "\n\t" {
    case "foo bar":
        if (1 > -1) {
            var v6 uint16 = 881
            fmt.Println(v5, v6)
            fmt.Println(v5)
            _ = v6
        } else {
            var v7 int8 = -83
            _ = v7
        }
    case "\n\t":
        if ("a" != ("a" + ("\n\t" + ("\n\t" + (("a" + ("" + "\n\t")) + (("foo bar" + "") + "")))))) {
            v8 := GenS1{F0: 269}
            v5 = v5
            v9 := []int{552, 128, -1, -94}
            _ = v8
            _ = v9
        }
    case "hello":
        v10 := []int{950, 32767, 42, 32767}
        _ = v10
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 float64 = (1.5 * -2.25)
            v13 := []int{(fi11 & 2147483647), 32768, (fi11 ^ (fi11 >> (((v4 >> ((255) & 63))) & 63)))}
            fmt.Println(v4, v12)
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v5, v5)
    for fi14 := 0; fi14 < 19; fi14++ {
        v4 = (v4 >> (((^(v4))) & 63))
        v5 = (^(v5))
        v15 := []int{2147483647, 42}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        _ = len(v1)
        _ = fi16
    }
    var v17 uint32 = 843
    fmt.Println(v17, v17, v17)
    v17 += 545
    v4 += (^(v4))
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: uint32(v4), F1: 4294967295, F2: (v5 / (((v5 % ((v5) | 1))) | 1))}
        if (!false) {
            v20 := GenS3{F0: (^(v5)), F1: v5}
            for fi21 := 0; fi21 < 15; fi21++ {
                switch "\n\t" {
                case "hello":
                    var v22 float32 = -2.25
                    _ = v22
                }
                _ = fi21
            }
            v23 := []int{fi18, -9223372036854775808}
            _ = v20
            _ = v23
        }
        v4 -= (v4 | (v4 - (v4 << ((uint(v17)) & 63))))
        _ = fi18
        _ = v19
    }
    v24 := GenS2{F0: 4294967295, F1: 471, F2: 1}
    var v25 float32 = float32(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
