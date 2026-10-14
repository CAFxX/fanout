package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0() uint64 {
    var v6 bool = (65535 <= 433)
    if (0 != 2147483647) {
        switch 530 {
        case 127:
            fmt.Println(v6)
        default:
            v6 = v6
        }
    }
    v6 = (!((0.1 != (0.1 - 1e10)) && true))
    _ = v6
    return 2
}

func main() {
    var v1 int = 128
    var v2 uint16 = 255
    v3 := []uint32{65535, 2}
    var v4 int16 = 1
    v5 := []int{v1, v1}
    if false {
        fmt.Println(v4)
        if len(v3) > 0 {
            v3[0] = 65535
        }
    } else {
        var v7 string = (("hello" + "") + "\n\t")
        for fi8 := 0; fi8 < 11; fi8++ {
            v4 -= int16(v2)
            v7 = "\n\t"
            _ = fi8
        }
        var v9 int16 = int16(v1)
        _ = v7
        _ = v9
    }
    v2 = (^(v2))
    fmt.Println(v4)
    switch "a" {
    case "foo bar":
        v10 := GenS1{F0: 1, F1: 127, F2: 1}
        _ = v10
    case "\n\t":
        fmt.Println(v1, v4, v2)
    case "":
        v11 := GenS1{F0: 1, F1: 42, F2: -2147483648}
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        _ = len(v3)
        _ = fi12
    }
    if (0 != 230) {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint32 = 2
            if (true && false) {
                fmt.Println(v14, fi13)
                v15 := GenS2{F0: 42, F1: ((true && true) || (true && true))}
                if false {
                    v1 = (^(v1))
                    var v16 uint64 = 84
                    _ = v16
                }
                _ = v15
            }
            switch (v1 << ((255) & 63)) {
            case -2147483648:
                v2 = 261
            default:
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint16 = 2
                    v19 := GenS3{F0: 0, F1: 42, F2: (v18 >> ((434) & 63))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi13
            _ = v14
        }
        v4 -= (v4 % ((v4) | 1))
    } else {
        var v20 bool = true
        _ = v20
    }
    v1 = (v1 & v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
