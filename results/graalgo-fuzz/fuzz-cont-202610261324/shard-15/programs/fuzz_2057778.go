package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 int16; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16, p1 int32, p2 int16) string {
    v4 := GenS1{F0: 1, F1: ""}
    p0 = (p0 - (p0 & (p0 & (^(p0)))))
    if (false && (42 != 1)) {
        if false {
            if (!(!true)) {
                switch 32768 {
                case -2147483648:
                    var v5 uint64 = 42
                    _ = v5
                case -800:
                    var v6 uint64 = 83
                    _ = v6
                case 886:
                    for fi7 := 0; fi7 < 12; fi7++ {
                        var v8 int64 = 42
                        if ((p0 - p0) < 2) {
                            v9 := GenS2{F0: -9223372036854775808, F1: 100, F2: false}
                            p1 = p1
                            switch (fi7 >> ((65535) & 63)) {
                            case -914:
                                var v10 int64 = (v8 << ((42) & 63))
                                _ = v10
                            case -95:
                                var v11 bool = (!false)
                                _ = v11
                            case -9223372036854775808:
                                var v12 float32 = 1e10
                                _ = v12
                            }
                            _ = v9
                        }
                        _ = fi7
                        _ = v8
                    }
                }
            }
            if (true && false) {
                v13 := GenS3{F0: p1, F1: p2, F2: 255}
                if true {
                    p2 = (p2 % (((^(p2))) | 1))
                    var v14 int = 0
                    _ = v14
                }
                v15 := []int{9223372036854775807, 127, 4294967295, 140}
                _ = v13
                _ = v15
            } else {
                v16 := GenS1{F0: 2, F1: "a"}
                v17 := []int{}
                _ = v16
                _ = v17
            }
            v18 := []int{127}
            _ = v18
        }
    }
    var v19 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v4
    return ("foo bar" + "")
}

func main() {
    var v1 int8 = -128
    var v2 int32 = 66
    var v3 int8 = v1
    fmt.Println(v2, v3)
    if (2 < 255) {
        fmt.Println(v2, v3)
        if (v2 > v2) {
            v3 += (v3 - (v1 ^ -128))
            switch -9223372036854775808 {
            case -437:
                var v20 float32 = 3.14159
                _ = v20
            default:
                v3 -= (v1 & v1)
            }
        } else {
            v21 := GenS1{F0: 0, F1: "hello"}
            var v22 int8 = 127
            _ = v21
            _ = v22
        }
        v23 := []int{-2147483648, 794, 1}
        _ = v23
    } else {
        v2 += 0
        v1 = 1
    }
    v3 -= v3
    v1 -= (v1 / ((v3) | 1))
    v1 = 127
    v3 -= v1
    v24 := GenS4{F0: 0}
    v25 := GenS3{F0: -2147483648, F1: int16(v1), F2: 587}
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint64 = 18446744073709551615
        v28 := []int{(-(fi26)), (fi26 / ((2147483647) | 1))}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint64 = 727
    for fi30 := 0; fi30 < 4; fi30++ {
        v3 = (v1 | -103)
        _ = fi30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        fmt.Println(v3)
        _ = fi31
    }
    fmt.Println(v29, v2, v29)
    v32 := []int{32767, 988}
    v33 := GenS5{F0: 250, F1: 42}
    fmt.Println(v3, v29)
    var v34 int64 = 0
    v35 := GenS4{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
