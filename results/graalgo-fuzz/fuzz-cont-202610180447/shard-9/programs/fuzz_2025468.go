package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 165
    var v2 int64 = -374
    var v3 uint8 = 128
    switch 127 {
    case 9223372036854775807:
        v2 -= -9223372036854775808
    }
    for fi4 := 0; fi4 < 7; fi4++ {
        switch (fi4 << ((478) & 63)) {
        case 1:
            for fi5 := 0; fi5 < 17; fi5++ {
                v6 := []int{(fi5 / (((fi4 ^ (fi4 | (fi5 - (^(fi4)))))) | 1)), (fi5 - fi4)}
                var v7 uint8 = 246
                _ = fi5
                _ = v6
                _ = v7
            }
        case 2147483647:
            var v8 uint8 = (v3 % ((v3) | 1))
            _ = v8
        default:
            switch "a" {
            case "\n\t":
                v9 := []int{fi4, (fi4 - (fi4 | int(v3))), (fi4 + (fi4 + (-(fi4))))}
                _ = v9
            }
        }
        _ = fi4
    }
    if (false && (260 <= 2)) {
        v3 += 216
        v10 := []int{9223372036854775807, 767, int(v1), -1}
        _ = v10
    }
    v11 := GenS1{F0: 0}
    var v12 int8 = -128
    v2 = v2
    v13 := GenS1{F0: -666}
    if (!(("" <= "hello") || (42 < 32767))) {
        v2 = 9223372036854775807
        v14 := GenS1{F0: -1}
        _ = v14
    }
    v15 := []int{-359}
    v1 += (v3 - v1)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
