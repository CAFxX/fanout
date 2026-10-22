package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 bool }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    var v1 int = 1
    var v2 uint = 1
    var v3 bool = (533 > 1)
    v4 := make(map[uint32]int)
    var v5 uint = 722
    v6 := []string{"", "", "\n\t", "hello"}
    v4[65535] = (v1 | v1)
    v3 = ((0.1 / 1.5) >= 0.0)
    v7 := GenS1{F0: 2, F1: 41}
    switch 9223372036854775807 {
    case 42:
        if false {
            v1 += 32767
            v3 = (1 <= 102)
            for fi8 := 0; fi8 < 19; fi8++ {
                v1 = (fi8 << (((v2 & (v2 << ((v2) & 63)))) & 63))
                _ = fi8
            }
        } else {
            v9 := GenS2{F0: 2, F1: 1, F2: (18446744073709551615 > 54)}
            _ = v9
        }
    case 4294967295:
        var v10 int32 = -1
        _ = v10
    case -9223372036854775808:
        var v11 uint = 533
        _ = v11
    }
    v12 := GenS3{F0: (v2 - 255), F1: v3}
    v2 += v5
    var v13 bool = v3
    switch (v1 << (((v2 / ((v2) | 1))) & 63)) {
    case 1:
        var v14 int8 = -1
        _ = v14
    case 233:
        v4[65535] = v1
    case -2147483648:
        v3 = true
    }
    var v15 int32 = -123
    if ((42 < 42) || v3) {
        v16 := GenS3{F0: (v5 ^ v2), F1: v13}
        _ = len(v4)
        _ = v16
    }
    v2 += v5
    if false {
        switch 127 {
        case 32767:
            var v17 int32 = (-(v15))
            _ = v17
        }
    } else {
        var v18 int = v1
        var v19 uint16 = 2
        _ = v18
        _ = v19
    }
    var v20 int8 = 100
    var v21 int16 = 32767
    v15 -= 42
    for fi22 := 0; fi22 < 16; fi22++ {
        v1 = (fi22 << ((4294967295) & 63))
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
