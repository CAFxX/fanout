package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{42, 0, 0}
    var v3 uint8 = 2
    var v4 uint8 = v3
    if len(v2) > 0 {
        v2[0] = 581
    }
    switch 9223372036854775807 {
    case -993:
        v4 = (v4 % (((v4 / ((v3) | 1))) | 1))
    default:
        _ = len(v2)
    }
    v4 = (v3 - v4)
    switch -2147483648 {
    case 4294967295:
        if (946 == 18446744073709551615) {
            fmt.Println(v3, v1, v3)
        }
    case -321:
        if true {
            if false {
                v3 = (v4 ^ v4)
                for fi5 := 0; fi5 < 3; fi5++ {
                    if false {
                        v6 := GenS1{F0: 65535}
                        _ = v6
                    }
                    v1 = v1
                    v1 = v1
                    _ = fi5
                }
            }
            if (true || (0 < uint(v3))) {
                v3 -= 127
                v7 := GenS1{F0: 558}
                _ = v7
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case -2147483648:
        v8 := GenS1{F0: 2}
        _ = v8
    default:
        v9 := GenS1{F0: 0}
        _ = v9
    }
    v10 := GenS2{F0: 0}
    var v11 uint16 = 255
    v12 := GenS2{F0: -7}
    v13 := GenS1{F0: 0}
    v14 := GenS1{F0: 270}
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
