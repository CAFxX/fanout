package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0(p0 uint16) uint32 {
    p0 -= (p0 << ((65535) & 63))
    switch "hello" {
    case "hello":
        fmt.Println(p0)
    }
    _ = p0
    return 314
}
func genFn1(p0 uint64, p1 int8, p2 bool) int8 {
    v5 := GenS1{F0: 1}
    p0 = (p0 & p0)
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = 18446744073709551615
    var v3 uint16 = 65535
    var v4 uint32 = 65535
    var v6 int = 32767
    if (genFn1(918, genFn1(42, 0, (("hello" >= "") || (!true))), false) <= -1) {
        var v7 uint32 = 4294967295
        var v8 uint = (v1 - v2)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint32 = v4
        var v11 int16 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: v3, F2: 127}
    var v13 uint8 = 2
    var v14 int64 = 42
    if (true && true) {
        if false {
            var v15 int64 = 802
            v16 := GenS1{F0: -9223372036854775808}
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: (v14 >> ((v2) & 63))}
                v19 := []int{(fi17 >> (((v1 % (((v2 % (((v1 >> ((1) & 63))) | 1))) | 1))) & 63)), -91}
                for fi20 := 0; fi20 < 1; fi20++ {
                    v3 = v3
                    var v21 bool = true
                    v22 := GenS1{F0: v14}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := []int{128, (v6 + (v6 - 2147483647))}
                _ = fi23
                _ = v24
            }
        }
    } else {
        v25 := GenS3{F0: 228, F1: (^(v3)), F2: 0}
        _ = v25
    }
    v26 := GenS4{F0: false, F1: -32}
    var v27 float32 = 1.5
    v28 := GenS5{F0: genFn0((v3 & (v3 & 855))), F1: 1, F2: (v3 >> ((v2) & 63))}
    switch (v6 % ((v6) | 1)) {
    case 0:
        v29 := []int{}
        _ = v29
    case 438:
        if (false || (uint(v13) == 18446744073709551615)) {
            v30 := []int{int(v14), v6, int(v2)}
            v13 = v13
            v31 := GenS3{F0: (v13 >> (((v2 << (((v1 + v2)) & 63))) & 63)), F1: uint16(v13), F2: -495}
            _ = v30
            _ = v31
        } else {
            var v32 bool = (39 <= 1)
            _ = v32
        }
    default:
        v4 = v4
    }
    v33 := GenS1{F0: v14}
    for fi34 := 0; fi34 < 4; fi34++ {
        switch "" {
        case "hello":
            fmt.Println(fi34)
        case "":
            fmt.Println(v14, v6, fi34)
        case "\n\t":
            switch 2147483647 {
            case 127:
                v35 := GenS1{F0: v14}
                _ = v35
            case 32768:
                var v36 uint16 = 1
                _ = v36
            case 0:
                fmt.Println(v2, v6)
            default:
                v2 += (^(v2))
            }
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v6
}
