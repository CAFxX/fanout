package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 int = 665
    var v2 uint16 = 42
    var v3 int8 = -128
    var v4 uint32 = 42
    switch (v1 % (((v1 >> ((65535) & 63))) | 1)) {
    case 32767:
        v2 += 42
    case -1:
        var v5 uint = 621
        _ = v5
    case 128:
        v1 = v1
    }
    var v6 int8 = 127
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint32 = v4
            _ = len(v9)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{}
    for fi13 := 0; fi13 < 3; fi13++ {
        v2 -= (v2 >> ((191) & 63))
        _ = fi13
    }
    var v14 uint = 255
    if (-1 == 0) {
        fmt.Println(v2)
        v1 -= (v1 * (-(v1)))
    } else {
        v2 = 629
        v12 = append(v12, 4294967295)
        if (v14 <= (v14 >> (((v14 & (v14 | (^(v14))))) & 63))) {
            switch v1 {
            case -360:
                if false {
                    v15 := GenS1{F0: 595, F1: -615}
                    v1 = (v1 + v1)
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v15
                }
            case -618:
                v3 += int8(v2)
            case -1:
                v16 := GenS1{F0: -32768, F1: 587}
                _ = v16
            }
        }
    }
    var v17 int = (v1 & -993)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
