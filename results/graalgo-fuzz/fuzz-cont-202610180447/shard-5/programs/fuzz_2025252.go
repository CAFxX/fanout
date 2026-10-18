package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int, p2 uint) uint {
    if true {
        for fi5 := 0; fi5 < 5; fi5++ {
            p0 = 127
            _ = fi5
        }
        v6 := GenS1{F0: 152, F1: p0}
        _ = v6
    }
    v7 := GenS2{F0: 42, F1: 415}
    p1 += int(p0)
    var v8 int = (p1 | (p1 & (p1 + 128)))
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 = (p0 ^ p0)
        p0 -= (p0 ^ p0)
        switch (fi9 % (((fi9 * p1)) | 1)) {
        case -71:
            fmt.Println(p0, p2)
        default:
            p1 += 32768
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 << ((p2) & 63))
}

func main() {
    var v1 int16 = 32767
    var v2 bool = (!(-595 < -374))
    var v3 uint16 = 2
    v4 := make(map[string]uint32)
    v10 := GenS3{F0: 127, F1: 2147483647}
    v3 = v3
    v2 = (((!v2) || ("foo bar" == "\n\t")) || v2)
    if ((18446744073709551615 < 2) || v2) {
        v11 := GenS4{F0: 4294967295, F1: 1, F2: 0}
        var v12 uint64 = 1
        _ = v11
        _ = v12
    } else {
        switch 383 {
        case 42:
            var v13 int = 425
            _ = v13
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            v1 = (v1 ^ 423)
            _ = fi14
        }
    }
    v15 := []int{-1, 32767}
    _ = len(v15)
    if (112 == 2) {
        var v16 bool = (0 > 255)
        var v17 int32 = 426
        _ = v16
        _ = v17
    } else {
        v18 := GenS3{F0: 0, F1: 1}
        _ = v18
    }
    v19 := GenS5{F0: 14}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
