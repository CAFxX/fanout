package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int = 32768
    v2 := make(map[string]string)
    var v3 int32 = -560
    v4 := []int{(^(v1)), -717, 4294967295}
    for fi5 := 0; fi5 < 13; fi5++ {
        v3 -= v3
        var v6 uint16 = 42
        v6 -= 2
        _ = fi5
        _ = v6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v2)
    v7 := GenS1{F0: (v3 * (v3 << ((65535) & 63))), F1: 0}
    v4 = append(v4, v1)
    var v8 float64 = 1e10
    v8 += 3.14159
    switch v1 {
    case 926:
        _ = len(v2)
    case 32768:
        switch "a" {
        case "hello":
            v9 := GenS2{F0: 448}
            _ = v9
        case "":
            v3 = (v3 % (((v3 | (v3 % (((v3 % ((v3) | 1))) | 1)))) | 1))
        default:
            if true {
                v8 -= v8
                for fi10 := 0; fi10 < 11; fi10++ {
                    v4 = append(v4, (fi10 ^ (v1 + (fi10 + -878))))
                    _ = fi10
                }
            }
        }
    }
    v1 = v1
    var v11 float32 = (float32(((v8 - v8) - -2.25)) - ((-2.25 * 1e10) / -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
