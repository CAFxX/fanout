package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]string)
    v3 := []int{-1, -731, 32767}
    var v4 string = "hello"
    var v5 uint8 = 42
    v6 := GenS1{F0: 820, F1: 1, F2: 0}
    v7 := v2[v4]
    var v8 uint = 4294967295
    var v9 int16 = -57
    v7 = v7
    _ = len(v3)
    for fi10 := 0; fi10 < 14; fi10++ {
        v4 = "\n\t"
        v2["\n\t"] = "a"
        if (v5 != v5) {
            var v11 float64 = 1e10
            switch int(v8) {
            case 104:
                v9 = (-(v9))
            default:
                v12 := GenS1{F0: 1, F1: v9, F2: 65535}
                _ = v12
            }
            _ = v11
        }
        _ = fi10
    }
    v13 := GenS1{F0: 2147483648, F1: (v9 - v9), F2: 419}
    v8 -= 1
    v1 = "\n\t"
    switch v1 {
    case "hello":
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS2{F0: 65535, F1: 42, F2: (fi14 / ((fi14) | 1))}
            _ = fi14
            _ = v15
        }
    case "\n\t":
        var v16 uint = v8
        _ = v16
    default:
        v17 := v2["hello"]
        _ = v17
    }
    v18 := GenS3{F0: (^(v9)), F1: false, F2: 17}
    v5 -= v5
    if (uint64(v8) > 18446744073709551615) {
        var v19 uint32 = 65535
        _ = v19
    }
    v20 := GenS3{F0: (v9 << ((v8) & 63)), F1: (-1 <= 127), F2: 2}
    for fi21 := 0; fi21 < 6; fi21++ {
        switch fi21 {
        case 9223372036854775807:
            if (v4 > (v7 + v4)) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    v4 = "\n\t"
                    fmt.Println(fi21, v9, v1)
                    _ = fi22
                }
                v23 := GenS4{F0: (("a" + ("\n\t" + v1)) + ((v1 + "\n\t") + v1)), F1: 42}
                var v24 uint32 = 0
                _ = v23
                _ = v24
            }
        case -727:
            var v25 int32 = int32(fi21)
            _ = v25
        default:
            _ = len(v3)
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
