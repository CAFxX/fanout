package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := []int{-2147483648, 32767}
    v3 := []uint32{790, 2147483647, 1, 2147483648}
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 2147483648}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    v6 := GenS1{F0: 2}
    switch -665 {
    case 32767:
        var v7 uint = 42
        _ = v7
    default:
        v8 := GenS2{F0: (!(-9223372036854775808 < -724))}
        _ = v8
    }
    var v9 int8 = 1
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint16 = 65535
        if true {
            var v12 uint64 = 4294967295
            var v13 float64 = ((1e10 + 1e10) + (1e10 - 3.14159))
            v12 += 609
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                v11 += (v11 % (((v11 + (v11 << ((18446744073709551615) & 63)))) | 1))
                _ = fi14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v11 += 2
                _ = fi15
            }
        }
        _ = fi10
        _ = v11
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        delete(v4, "a")
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint32 = 2
            _ = fi17
            _ = v18
        }
        fmt.Println(v9, fi16)
        _ = fi16
    }
    fmt.Println(v9)
    var v19 int16 = 32767
    for fi20 := 0; fi20 < 19; fi20++ {
        v19 += v19
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
