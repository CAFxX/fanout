package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
func genFn0() uint32 {
    var v5 float64 = 1e10
    v5 = -2.25
    _ = v5
    return 65535
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 127
    var v3 int16 = int16(v1)
    var v4 int32 = -97
    v1 -= (v1 << ((270) & 63))
    v6 := []int{-2147483648}
    if true {
        var v7 uint64 = 2
        _ = v7
    }
    v8 := GenS1{F0: 32767, F1: 128}
    v9 := GenS2{F0: (v2 - v2)}
    for fi10 := 0; fi10 < 8; fi10++ {
        if len(v6) > 0 {
            v6[0] = 763
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: 0, F1: (^(v1))}
        _ = fi11
        _ = v12
    }
    var v13 int16 = (v3 >> ((403) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
