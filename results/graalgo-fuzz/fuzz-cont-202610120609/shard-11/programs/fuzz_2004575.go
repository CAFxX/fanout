package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint16 }

func main() {
    var v1 int32 = 2147483647
    v2 := []int{32767}
    var v3 uint16 = 1
    if (1e10 > (1.5 / float32(1e10))) {
        v4 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + "a"), F2: (v1 << ((4294967295) & 63))}
        v2 = append(v2, -1)
        if true {
            fmt.Println(v3)
            v3 = (v3 - (v3 + v3))
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
        } else {
            v3 += 125
            fmt.Println(v1, v1)
            v1 += (v1 / ((-1) | 1))
        }
        _ = v4
    }
    switch "hello" {
    case "a":
        var v5 int = 32767
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 85, F1: "foo bar", F2: (v1 % (((-(v1))) | 1))}
        v8 := GenS2{F0: 2, F1: false, F2: (^(v3))}
        var v9 int16 = 42
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int16 = 42
    var v11 bool = true
    v1 -= (v1 / (((v1 | -1)) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= (v1 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
