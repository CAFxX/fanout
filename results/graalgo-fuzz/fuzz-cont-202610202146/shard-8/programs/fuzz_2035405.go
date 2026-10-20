package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 479
    var v3 string = "a"
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(v2, fi4, v3)
        _ = fi4
    }
    v3 = v3
    v5 := GenS1{F0: (v2 << ((328) & 63))}
    v3 = ""
    v6 := []int{-918, 32768, 9223372036854775807, 32767}
    if ((false || true) || (0 != 42)) {
        var v7 uint16 = (v2 % (((v2 >> ((2) & 63))) | 1))
        v1 -= v1
        _ = v7
    }
    v1 = (v1 - v1)
    if (v3 <= (v3 + v3)) {
        var v8 uint16 = (v2 & (v2 << ((0) & 63)))
        v9 := GenS2{F0: 4294967295}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: v2}
            v12 := GenS1{F0: (v2 >> ((1) & 63))}
            var v13 int64 = -1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := GenS3{F0: 2, F1: 112}
        for fi15 := 0; fi15 < 15; fi15++ {
            v3 = v3
            v1 = (v1 << ((109) & 63))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi15
        }
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
