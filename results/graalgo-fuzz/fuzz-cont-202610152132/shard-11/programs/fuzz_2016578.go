package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8, p1 bool) uint {
    if ((!(762 >= 65535)) && p1) {
        v7 := []int{9223372036854775807, -9223372036854775808, 32768}
        _ = v7
    }
    if p1 {
        p1 = (float32(3.14159) < (-2.25 * 1.5))
        var v8 bool = (p1 || p1)
        _ = v8
    }
    var v9 int = 4294967295
    _ = p0
    _ = p1
    _ = v9
    return 478
}
func genFn1(p0 uint8, p1 string, p2 uint) bool {
    var v10 int = 4294967295
    p1 = ""
    switch v10 {
    case -1:
        v11 := GenS1{F0: (p0 | p0), F1: -482}
        _ = v11
    case 32768:
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(p1)
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS2{F0: fi12, F1: 2, F2: (v10 & (v10 - 9223372036854775807))}
                switch (p1 + (p1 + p1)) {
                case "\n\t":
                    v15 := []int{(fi13 - fi13), fi13, -9223372036854775808, fi12}
                    _ = v15
                case "hello":
                    v16 := GenS2{F0: (fi13 + v10), F1: 255, F2: 559}
                    var v17 uint = 42
                    v18 := GenS1{F0: p0, F1: 382}
                    _ = v16
                    _ = v17
                    _ = v18
                case "foo bar":
                    p2 = 18446744073709551615
                default:
                    v19 := GenS1{F0: 0, F1: int64(p0)}
                    _ = v19
                }
                v20 := GenS3{F0: p2, F1: 9223372036854775807}
                _ = fi13
                _ = v14
                _ = v20
            }
            _ = fi12
        }
    case 32767:
        v21 := []int{42, (v10 / ((v10) | 1))}
        _ = v21
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            if (-6 <= -128) {
                switch "a" {
                case "hello":
                    p0 -= (p0 << ((genFn0(42, ((!(1.5 < -2.25)) && false))) & 63))
                case "foo bar":
                    var v23 string = (((("" + (p1 + ("a" + p1))) + (p1 + p1)) + "a") + (p1 + p1))
                    _ = v23
                case "\n\t":
                    if (false || (!false)) {
                        var v24 int32 = -2147483648
                        _ = v24
                    }
                default:
                    v25 := GenS2{F0: v10, F1: 0, F2: v10}
                    _ = v25
                }
            }
            _ = fi22
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 bool = (true && (430 != 2147483648))
        _ = fi26
        _ = v27
    }
    var v28 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v28
    return true
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    var v5 int32 = 644
    var v6 int32 = v5
    var v29 int32 = v6
    v6 -= v5
    if true {
        switch ("" + "a") {
        case "a":
            v6 = (v5 * v5)
        case "hello":
            for fi30 := 0; fi30 < 20; fi30++ {
                _ = len(v1)
                v31 := GenS4{F0: 2, F1: "hello"}
                delete(v3, "a")
                _ = fi30
                _ = v31
            }
        default:
            fmt.Println(v6)
        }
        switch "" {
        case "a":
            if true {
                v32 := GenS5{F0: -9223372036854775808}
                _ = v32
            } else {
                v33 := GenS3{F0: 1, F1: 42}
                _ = v33
            }
        case "\n\t":
            v6 = v6
        case "":
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 uint32 = 340
                _ = fi34
                _ = v35
            }
        }
        v36 := GenS3{F0: 255, F1: 1}
        _ = v36
    }
    v4 = (1e10 * 1.5)
    v1 = append(v1, 2147483647)
    v4 = v4
    v37 := GenS5{F0: -1}
    delete(v3, "a")
    v29 = (v29 / (((v6 << ((255) & 63))) | 1))
    var v38 bool = false
    fmt.Println(v6)
    v38 = ((1.5 * 1.5) == (1e10 + 1.5))
    var v39 uint8 = 128
    delete(v2, 127)
    v39 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
