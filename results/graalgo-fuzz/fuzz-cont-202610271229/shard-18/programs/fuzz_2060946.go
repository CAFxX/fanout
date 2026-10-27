package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint32 = 1
    v3 := []uint32{(v2 * v2)}
    var v4 int32 = -2147483648
    v5 := []uint32{(v2 ^ v2), v2, 65535, 4294967295}
    v6 := make(map[int]int)
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "" {
        case "\n\t":
            v2 -= 0
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint64 = 495
            fmt.Println(fi9, v11)
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: false, F1: 128}
        _ = fi9
        _ = v12
    }
    var v13 int16 = 1
    v14 := GenS1{F0: ((false || ((0.1 != 1.5) && false)) || true), F1: 180}
    v1 -= (-(v1))
    for fi15 := 0; fi15 < 8; fi15++ {
        switch "\n\t" {
        case "":
            v4 = -2147483648
        case "a":
            var v16 int16 = (v13 >> ((uint(v2)) & 63))
            _ = v16
        case "hello":
            if true {
                _ = len(v6)
                var v17 uint32 = v2
                _ = v17
            }
        }
        _ = fi15
    }
    var v18 int8 = 1
    var v19 int32 = -2147483648
    switch (v1 << ((633) & 63)) {
    case -2147483648:
        var v20 int16 = v13
        _ = v20
    case 9223372036854775807:
        if true {
            var v21 int16 = (v13 + 1)
            _ = v21
        } else {
            if false {
                if len(v5) > 0 {
                    v5[0] = v2
                }
            } else {
                v22 := GenS1{F0: false, F1: 5}
                for fi23 := 0; fi23 < 13; fi23++ {
                    v6[fi23] = v1
                    _ = len(v5)
                    _ = fi23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    v18 = (^(v18))
                    _ = fi24
                }
                _ = v22
            }
        }
    case 4294967295:
        for fi25 := 0; fi25 < 7; fi25++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
