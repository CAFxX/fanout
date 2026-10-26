package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 1
    v3 := []uint32{4294967295, 65535}
    v4 := make(map[string]int)
    var v5 int8 = -61
    if (26 != 255) {
        switch v2 {
        case 32767:
            v5 += -128
        }
        _ = len(v3)
        var v6 int64 = -212
        _ = v6
    } else {
        v7 := GenS1{F0: ("hello" != ""), F1: (true && (-9223372036854775808 > -1))}
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int8 = 0
            if false {
                v10 := GenS1{F0: (true || (("hello" + "\n\t") != ("" + "hello"))), F1: (128 != 246)}
                _ = v10
            }
            var v11 int8 = v5
            _ = fi8
            _ = v11
            _ = v9
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v7
    }
    var v12 bool = (1 >= 42)
    var v13 uint32 = 4294967295
    v13 = v13
    if len(v3) > 0 {
        v3[0] = (^(v13))
    }
    v14 := GenS1{F0: (v2 <= (v2 % (((-(v2))) | 1))), F1: true}
    if false {
        _ = len(v1)
    } else {
        fmt.Println(v2, v2, v12)
    }
    v2 = -9223372036854775808
    var v15 int64 = -554
    v13 = v13
    v15 += -1
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
