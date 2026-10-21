package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint8 = 1
    var v2 uint8 = v1
    v3 := []uint32{}
    v4 := GenS1{F0: 0, F1: (v1 ^ v1)}
    v5 := GenS2{F0: 759, F1: 18446744073709551615}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v1)
    v2 = (v2 / (((v2 >> ((uint(v2)) & 63))) | 1))
    _ = len(v3)
    fmt.Println(v1)
    for fi6 := 0; fi6 < 1; fi6++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            if (true && (!(fi6 != fi6))) {
                var v7 string = "a"
                if (true || (4294967295 > 4294967295)) {
                    if ((1e10 * (0.0 - ((-2.25 + 1e10) * 0.0))) >= 0.0) {
                        v1 += 255
                    }
                }
                var v8 uint8 = (v1 + v2)
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 19; fi9++ {
                    _ = len(v3)
                    v2 = (v2 ^ v2)
                    _ = fi9
                }
                if (0.0 < 3.14159) {
                    v1 = (v1 / ((63) | 1))
                }
                if len(v3) > 0 {
                    v3[0] = 64
                }
            }
        case "foo bar":
            v3 = append(v3, 65535)
        default:
            v2 += (v1 + (v1 + (v1 >> ((4294967295) & 63))))
        }
        _ = fi6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
