package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int32) int64 {
    v6 := []int{1}
    switch ("a" + "hello") {
    case "":
        fmt.Println(p0)
    case "a":
        if len(v6) > 0 {
            v6[0] = -691
        }
    case "foo bar":
        _ = len(v6)
    default:
        switch 32768 {
        case 127:
            if ((0.1 / 1e10) >= 0.1) {
                _ = len(v6)
            } else {
                var v7 float32 = 0.0
                _ = v7
            }
        default:
            p0 = p0
        }
    }
    fmt.Println(p0)
    p0 = (p0 / (((p0 << ((671) & 63))) | 1))
    _ = p0
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 string, p1 uint64, p2 uint) uint32 {
    p1 = p1
    fmt.Println(p0)
    v8 := GenS1{F0: 0, F1: true, F2: (p1 << ((p2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 926
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 255
    var v3 uint64 = 42
    var v4 uint8 = (v1 - v1)
    var v5 int = -9223372036854775808
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS2{F0: 4294967295, F1: 605, F2: true}
        _ = fi9
        _ = v10
    }
    switch "foo bar" {
    case "foo bar":
        v2 += 18446744073709551615
    default:
        if ("\n\t" >= ("\n\t" + "foo bar")) {
            for fi11 := 0; fi11 < 20; fi11++ {
                v5 += 128
                _ = fi11
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    if (!((255 != 305) && (100 >= 1))) {
                        v14 := []int{(fi12 / (((fi12 >> ((v2) & 63))) | 1)), (-(fi12))}
                        v4 = (v1 / (((v4 & (v4 | 1))) | 1))
                        _ = v14
                    }
                    var v15 uint16 = 42
                    _ = fi13
                    _ = v15
                }
                v16 := GenS2{F0: (v3 ^ (^(v3))), F1: 980, F2: true}
                _ = fi12
                _ = v16
            }
        }
    }
    v17 := GenS3{F0: (v3 % ((uint64(v2)) | 1)), F1: 65535}
    v18 := GenS4{F0: 2, F1: (v1 / (((v4 | v1)) | 1))}
    v5 = v5
    v3 += v3
    v3 = v3
    v4 += uint8(v3)
    switch ("foo bar" + ("" + "hello")) {
    case "hello":
        v1 = v1
    case "a":
        fmt.Println(v4, v5)
    }
    if true {
        v1 = v4
        var v19 int = -9223372036854775808
        v20 := []int{32767}
        _ = v19
        _ = v20
    } else {
        var v21 float64 = ((3.14159 + 3.14159) + 1e10)
        _ = v21
    }
    v22 := []int{}
    var v23 uint8 = v1
    v24 := GenS1{F0: 42, F1: false, F2: v3}
    if len(v22) > 0 {
        v22[0] = 4294967295
    }
    var v25 bool = (((v23 | v4) > (v1 % (((v1 << ((504) & 63))) | 1))) || ("foo bar" >= "foo bar"))
    if false {
        fmt.Println(v25)
    } else {
        var v26 uint64 = v3
        _ = v26
    }
    v5 -= (v5 / ((v5) | 1))
    var v27 uint8 = 20
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
