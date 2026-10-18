package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint; F2 int8 }

func main() {
    var v1 uint32 = 2
    var v2 float32 = 0.1
    var v3 uint8 = 128
    v4 := make(map[uint32]uint32)
    var v5 int16 = 756
    v6 := []int{-2147483648, 42, 9223372036854775807, 466}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if (-2.25 <= 1e10) {
                v9 := GenS1{F0: (!((int8(v5) > 1) && true)), F1: (fi8 * fi8), F2: 19}
                _ = v9
            }
            var v10 uint64 = 42
            _ = fi8
            _ = v10
        }
        var v11 uint16 = 42
        _ = fi7
        _ = v11
    }
    fmt.Println(v3, v3, v5)
    fmt.Println(v2, v5, v2)
    var v12 uint32 = (v1 / (((v1 - (v1 % ((v1) | 1)))) | 1))
    var v13 int = 1
    v14 := GenS1{F0: true, F1: v13, F2: 18446744073709551615}
    v15 := GenS1{F0: false, F1: v13, F2: 398}
    delete(v4, 1)
    var v16 int64 = -1
    var v17 uint32 = v1
    fmt.Println(v3, v12, v16)
    var v18 bool = (true && (32767 != uint16(v3)))
    for fi19 := 0; fi19 < 11; fi19++ {
        v2 += 3.14159
        v20 := GenS1{F0: (("a" + "a") < ("foo bar" + "a")), F1: (fi19 - (fi19 >> ((17) & 63))), F2: 2}
        _ = fi19
        _ = v20
    }
    v4[(v12 % ((v1) | 1))] = v1
    v5 -= 291
    fmt.Println(v5)
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            if (982 == 255) {
                var v23 bool = ((v5 - (v5 % ((v5) | 1))) == (v5 | (^(v5))))
                fmt.Println(v18)
                v24 := GenS2{F0: 255, F1: v3, F2: 647}
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        var v25 uint32 = v17
        _ = fi21
        _ = v25
    }
    v26 := GenS3{F0: v3, F1: 1, F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
