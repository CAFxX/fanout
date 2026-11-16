package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    var v1 uint32 = 1
    v2 := []string{"a", "\n\t", ("a" + ((("\n\t" + ("" + ("a" + "a"))) + "foo bar") + "a")), ""}
    v3 := []uint32{v1, v1, (v1 & (v1 / ((v1) | 1))), (v1 ^ (v1 % (((^(v1))) | 1)))}
    var v4 float64 = 1.5
    v5 := []uint32{v1}
    var v6 uint64 = 116
    var v7 int8 = 86
    switch 42 {
    case 127:
        if len(v3) > 0 {
            v3[0] = (v1 << ((255) & 63))
        }
    case -529:
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: (true && true), F1: 472, F2: 2147483647}
            _ = fi8
            _ = v9
        }
    case 4294967295:
        v10 := GenS2{F0: v6, F1: 402, F2: (v7 * int8(v6))}
        _ = v10
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if true {
        var v11 uint8 = 55
        switch ("a" + "") {
        case "\n\t":
            var v12 int64 = 42
            _ = v12
        }
        v13 := GenS2{F0: 2, F1: (v1 << ((0) & 63)), F2: int8(v11)}
        _ = v11
        _ = v13
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1, v1)
        v5 = append(v5, 2147483647)
    }
    var v14 uint64 = (^(v6))
    fmt.Println(v7)
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS3{F0: 487, F1: (^(fi15)), F2: 225}
        switch (-(fi15)) {
        case 128:
            if len(v5) > 0 {
                v5[0] = (v1 << ((682) & 63))
            }
        case 648:
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS4{F0: 150, F1: 65535, F2: v1}
                _ = fi17
                _ = v18
            }
        default:
            v14 += 42
        }
        v6 = v14
        _ = fi15
        _ = v16
    }
    var v19 int = 42
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS3{F0: 372, F1: (fi20 & 32767), F2: 2147483647}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
