package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := []uint32{}
    var v2 int16 = 42
    var v3 uint = 6
    _ = len(v1)
    var v4 int16 = v2
    v4 = v4
    fmt.Println(v4)
    fmt.Println(v3, v4)
    if (594 > 824) {
        if (793 < 18446744073709551615) {
            for fi5 := 0; fi5 < 13; fi5++ {
                for fi6 := 0; fi6 < 1; fi6++ {
                    var v7 int16 = 32767
                    switch (fi6 >> (((v3 << ((255) & 63))) & 63)) {
                    case 9223372036854775807:
                        fmt.Println(fi5)
                    }
                    _ = fi6
                    _ = v7
                }
                _ = fi5
            }
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: -71, F1: fi8}
                _ = fi8
                _ = v9
            }
        }
    } else {
        v2 += (v4 | -749)
        fmt.Println(v4, v4, v2)
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint8 = 1
        v2 += v2
        v1 = append(v1, 65535)
        _ = fi10
        _ = v11
    }
    var v12 int8 = 1
    v2 -= 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
