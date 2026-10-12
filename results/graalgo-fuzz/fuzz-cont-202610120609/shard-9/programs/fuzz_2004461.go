package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := []int{32767, 755, -2147483648}
    var v2 uint16 = 1
    var v3 int16 = -487
    v1 = append(v1, -2147483648)
    if false {
        v3 -= (v3 & v3)
    } else {
        var v4 int64 = 42
        _ = v4
    }
    var v5 uint64 = 18446744073709551615
    v3 = v3
    switch 128 {
    case 1:
        if len(v1) > 0 {
            v1[0] = 0
        }
    case 32767:
        var v6 uint16 = 42
        _ = v6
    case -169:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v7 string = "hello"
    var v8 int = 521
    var v9 int8 = 1
    if true {
        var v10 int8 = (v9 ^ (v9 * (v9 << ((311) & 63))))
        switch v8 {
        case -765:
            v11 := GenS1{F0: 1}
            _ = v11
        }
        v10 = (v9 >> ((446) & 63))
        _ = v10
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (0 != -1) {
            var v12 uint8 = 42
            fmt.Println(v12, v5)
            if false {
                v13 := GenS1{F0: 177}
                _ = v13
            }
            _ = v12
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v7, v2)
                var v15 uint32 = 891
                for fi16 := 0; fi16 < 7; fi16++ {
                    fmt.Println(fi16, v3)
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
            v17 := GenS2{F0: uint8(v2), F1: -825}
            _ = v17
        }
    }
    var v18 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
