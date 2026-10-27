package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint16 }
func genFn0(p0 uint32) uint32 {
    p0 -= 4294967295
    var v5 uint32 = p0
    _ = p0
    _ = v5
    return (v5 * v5)
}
func genFn1(p0 int16) uint8 {
    v6 := GenS1{F0: -1, F1: 75, F2: 0}
    fmt.Println(p0)
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = -1
        p0 = p0
        _ = fi7
    }
    var v8 string = "\n\t"
    v8 = "\n\t"
    _ = p0
    _ = v6
    _ = v8
    return 245
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{128, 1, 65535}
    var v3 int8 = 1
    var v4 int8 = (v3 ^ (^(v3)))
    v4 += (v4 - (v4 >> ((18446744073709551615) & 63)))
    _ = len(v2)
    fmt.Println(v4, v4, v1)
    v2 = append(v2, 2147483648)
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: -1, F1: 18446744073709551615, F2: 42}
            _ = fi10
            _ = v11
        }
        v1 = (v1 + "")
        _ = fi9
    }
    v4 += (v3 ^ v4)
    switch -2147483648 {
    case -2147483648:
        v2 = append(v2, 4294967295)
    case -1:
        v2 = append(v2, 2147483648)
    default:
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint16 = 0
            v13 += (v13 & v13)
            switch v1 {
            case "":
                if ((!(v1 < v1)) && true) {
                    v4 = v4
                } else {
                    if (!((1 == 0) && true)) {
                        v14 := GenS2{F0: 54, F1: (-(v4))}
                        v15 := GenS2{F0: 42, F1: (-(v3))}
                        v1 = (v1 + "foo bar")
                        _ = v14
                        _ = v15
                    }
                    fmt.Println(v3, fi12)
                    v16 := GenS2{F0: 42, F1: v3}
                    _ = v16
                }
            case "foo bar":
                var v17 uint8 = 255
                _ = v17
            case "\n\t":
                v18 := GenS3{F0: 42, F1: 790, F2: (^(v13))}
                _ = v18
            }
            _ = fi12
            _ = v13
        }
    }
    v3 = 83
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
