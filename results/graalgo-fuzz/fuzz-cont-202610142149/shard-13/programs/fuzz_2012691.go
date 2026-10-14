package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int8 = -128
    v5 := []uint32{}
    if false {
        v1 += 0
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        _ = len(v2)
        _ = fi6
    }
    if true {
        if (1 < 42) {
            var v7 string = "hello"
            delete(v3, 2147483648)
            _ = v7
        }
    }
    var v8 string = "a"
    v9 := GenS1{F0: 651}
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint64 = 696
            switch "\n\t" {
            case "hello":
                v11 -= v11
            }
            v12 := GenS2{F0: 42, F1: (v1 << ((781) & 63)), F2: 18446744073709551615}
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
    }
    v13 := GenS2{F0: 0, F1: (v1 + (v1 << ((4294967295) & 63))), F2: 65535}
    for fi14 := 0; fi14 < 14; fi14++ {
        switch v8 {
        case "hello":
            v4 += (^(v4))
        case "":
            if len(v5) > 0 {
                v5[0] = 0
            }
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = 2
            }
        }
        _ = fi14
    }
    v8 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
