package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    var v1 int32 = 2147483647
    var v2 bool = false
    var v3 uint64 = 0
    var v4 uint = 1
    var v5 int32 = v1
    var v6 uint8 = 128
    v1 = (^(v1))
    var v7 uint8 = uint8(v4)
    if (0.1 > 1e10) {
        var v8 int16 = -99
        fmt.Println(v3, v8, v2)
        var v9 string = ""
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint64 = 2
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        if v2 {
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 float64 = 1e10
                    _ = fi14
                    _ = v15
                }
                v6 += (^(v6))
                _ = fi13
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                v6 += uint8(v3)
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := []int{}
                    var v19 uint32 = 2147483647
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if v2 {
                    fmt.Println(v1, v6)
                }
                _ = fi16
            }
        } else {
            if v2 {
                v6 = (v7 - v7)
                v5 += v1
                var v20 float32 = 0.1
                _ = v20
            }
        }
        var v21 int16 = 551
        _ = v21
    }
    var v22 int32 = v1
    v23 := GenS1{F0: -1}
    v3 += v3
    v24 := GenS2{F0: (v6 >> (((v4 & v4)) & 63)), F1: v4}
    v1 = v22
    v25 := GenS1{F0: int8(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
