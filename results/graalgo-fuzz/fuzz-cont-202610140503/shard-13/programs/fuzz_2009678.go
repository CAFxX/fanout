package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483648, 684, 1}
    v3 := []string{}
    v4 := GenS1{F0: 255, F1: -1, F2: 1}
    switch "foo bar" {
    case "\n\t":
        var v5 int32 = 1
        _ = v5
    }
    var v6 uint64 = 18446744073709551615
    var v7 bool = false
    if false {
        fmt.Println(v6, v7)
    }
    v8 := GenS2{F0: 9223372036854775807, F1: 935}
    v6 += 1
    v9 := GenS1{F0: 322, F1: -2147483648, F2: v6}
    v10 := GenS2{F0: 32768, F1: 2147483648}
    if ((255 >= 55) && (v7 || false)) {
        for fi11 := 0; fi11 < 8; fi11++ {
            _ = len(v2)
            v12 := GenS1{F0: 2, F1: 42, F2: (v6 / (((v6 >> ((17) & 63))) | 1))}
            _ = fi11
            _ = v12
        }
        var v13 uint64 = (^(v6))
        v6 = v13
        _ = v13
    }
    if (v7 || true) {
        v1[706] = ("\n\t" + "hello")
    }
    v14 := GenS2{F0: 4294967295, F1: 42}
    for fi15 := 0; fi15 < 7; fi15++ {
        v1[0] = ""
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS2{F0: (fi16 >> ((4294967295) & 63)), F1: 2147483648}
            if v7 {
                v7 = v7
                fmt.Println(fi15, fi15, fi15)
            } else {
                v18 := GenS1{F0: 672, F1: 815, F2: v6}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v19 int16 = 0
                _ = v18
                _ = v19
            }
            v6 -= (^(v6))
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    _ = len(v2)
    v20 := GenS3{F0: 809}
    v21 := GenS1{F0: 0, F1: 0, F2: 1}
    var v22 float64 = 0.0
    v2 = append(v2, 363)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
