package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 uint64) int32 {
    p1 -= (p1 - 0)
    var v4 uint64 = 42
    v5 := []int{127, -2147483648}
    fmt.Println(p1)
    p0 -= p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1(p0 bool, p1 int16) uint32 {
    var v6 bool = true
    v7 := GenS1{F0: 536, F1: 480, F2: 1}
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS2{F0: 255, F1: 0, F2: false}
        v10 := GenS2{F0: 128, F1: 0, F2: (v6 && true)}
        fmt.Println(p1, p0)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    v3 := []string{"hello"}
    v2 = v2
    var v11 uint64 = 18446744073709551615
    v12 := GenS3{F0: 0, F1: -545}
    v11 += v11
    v11 = v11
    var v13 uint16 = 2
    var v14 uint8 = 47
    v15 := GenS4{F0: (!(-588 <= 1))}
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            if false {
                v18 := GenS1{F0: 18446744073709551615, F1: 255, F2: -1}
                var v19 uint32 = 1
                var v20 int64 = 9223372036854775807
                v19 = 4294967295
                _ = v18
                _ = v19
                _ = v20
            }
            v13 = (^(v13))
            var v21 uint32 = 4294967295
            var v22 uint8 = (v14 | 255)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi17
            _ = v21
            _ = v22
        }
        v13 = (v13 | (v13 * 255))
        _ = fi16
    }
    v23 := GenS3{F0: 1, F1: 0}
    fmt.Println(v2, v2, v14)
    v24 := GenS2{F0: 42, F1: 9223372036854775807, F2: false}
    switch "" {
    case "hello":
        fmt.Println(v2, v11)
    default:
        for fi25 := 0; fi25 < 10; fi25++ {
            switch (v2 + (v2 + v2)) {
            case "foo bar":
                v14 = (v14 >> ((18446744073709551615) & 63))
            case "\n\t":
                v13 = (v13 - 2)
            }
            _ = fi25
        }
    }
    fmt.Println(v14, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
