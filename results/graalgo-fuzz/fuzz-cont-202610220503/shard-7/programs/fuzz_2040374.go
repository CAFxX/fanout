package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 1
    var v2 int = -125
    v3 := []uint32{2, 1, 2147483647}
    if false {
        v2 = (v2 - 418)
        switch 42 {
        case 127:
            fmt.Println(v1, v2)
        }
    } else {
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: uint16(v1)}
            fmt.Println(v1, v1)
            _ = fi4
            _ = v5
        }
    }
    var v6 bool = false
    if len(v3) > 0 {
        v3[0] = 65535
    }
    switch v2 {
    case 1:
        v3 = append(v3, uint32(v1))
    default:
        v7 := GenS1{F0: 32767}
        _ = v7
    }
    var v8 int32 = 821
    v9 := GenS2{F0: 2, F1: 255, F2: 2}
    for fi10 := 0; fi10 < 18; fi10++ {
        if (68 == 32767) {
            v8 += v8
        } else {
            if (9223372036854775807 > -1) {
                var v11 int8 = -128
                _ = v11
            }
            var v12 int64 = int64(v1)
            _ = v12
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi10
    }
    v8 -= 0
    if len(v3) > 0 {
        v3[0] = 63
    }
    fmt.Println(v6, v2)
    var v13 bool = v6
    var v14 int16 = -32768
    v13 = (v6 && v13)
    v15 := GenS3{F0: 0}
    v8 += (v8 & (v8 >> ((65535) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
