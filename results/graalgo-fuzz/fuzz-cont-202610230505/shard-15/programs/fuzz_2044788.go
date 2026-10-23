package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0(p0 int32, p1 bool, p2 uint) uint16 {
    v5 := GenS1{F0: 401}
    p2 += p2
    var v6 uint32 = 292
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 45
    var v3 bool = true
    v4 := []int{-578, 42}
    v1 = v1
    switch 0 {
    case 1:
        switch 127 {
        case 614:
            switch -192 {
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 int64 = 9223372036854775807
                    fmt.Println(v9, v2, v9)
                    _ = fi8
                    _ = v9
                }
            }
        }
    case 42:
        v1 = v1
    case -2147483648:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int32 = 2147483647
            v3 = v3
            v1 -= (v1 + v1)
            _ = fi10
            _ = v11
        }
    }
    var v12 uint64 = 398
    v13 := GenS1{F0: 1000}
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v3, v1)
        _ = fi14
    }
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 10; fi16++ {
        switch "foo bar" {
        case "foo bar":
            v17 := GenS2{F0: 42, F1: 32767, F2: v1}
            _ = v17
        case "hello":
            if len(v4) > 0 {
                v4[0] = (-(fi16))
            }
        case "a":
            var v18 uint32 = 166
            _ = v18
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi16
    }
    v4 = append(v4, 1)
    if len(v4) > 0 {
        v4[0] = 128
    }
    if v3 {
        var v19 int8 = 1
        _ = v19
    } else {
        if len(v4) > 0 {
            v4[0] = int(v12)
        }
    }
    v20 := GenS3{F0: 65535, F1: 875, F2: 330}
    v12 += v12
    fmt.Println(v12, v2)
    var v21 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
