package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0(p0 int32, p1 int8) int8 {
    p0 += (p0 << ((65535) & 63))
    if true {
        var v6 uint32 = 2147483648
        _ = v6
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 bool) string {
    v7 := []int{446, 0}
    switch 128 {
    case -1:
        if len(v7) > 0 {
            v7[0] = 1
        }
    }
    fmt.Println(p0)
    var v8 int8 = -128
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 3.14159
    var v3 bool = (-2.25 == 0.1)
    var v4 int8 = 127
    v5 := []uint32{2147483647, 415, 0, 556}
    var v9 int32 = -273
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint64 = 18446744073709551615
        v3 = true
        _ = fi10
        _ = v11
    }
    if (!((((v2 + v2) * 1.5) * v2) == v2)) {
        v12 := GenS1{F0: 42, F1: v9}
        v13 := GenS2{F0: 18446744073709551615, F1: 32767, F2: 989}
        var v14 uint32 = 0
        _ = v12
        _ = v13
        _ = v14
    }
    switch "a" {
    case "":
        if ((0.0 != 0.0) && ((32767 <= 127) || v3)) {
            v15 := GenS1{F0: -1, F1: (v9 | v9)}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v15
        } else {
            if false {
                if v3 {
                    var v16 uint16 = 65535
                    v3 = (-977 < 2147483647)
                    v4 += int8(v16)
                    _ = v16
                } else {
                    if ((v9 * (v9 - (v9 | (^(v9))))) != (v9 >> ((593) & 63))) {
                        switch 426 {
                        case -2147483648:
                            var v17 uint64 = 122
                            _ = v17
                        }
                    }
                }
                v4 += -52
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    if (v3 && (v3 && false)) {
                        switch (fi18 % (((fi18 * (fi18 << ((65535) & 63)))) | 1)) {
                        case -9223372036854775808:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 32768:
                            fmt.Println(v4, v4)
                        case -965:
                            var v19 int32 = (v9 | (v9 | v9))
                            _ = v19
                        }
                    } else {
                        var v20 uint32 = 2
                        _ = v20
                    }
                    fmt.Println(v3, fi18)
                    var v21 uint16 = 2
                    _ = fi18
                    _ = v21
                }
                for fi22 := 0; fi22 < 10; fi22++ {
                    if (0 >= genFn0(v9, (v4 | genFn0((v9 - v9), -1)))) {
                        delete(v1, "\n\t")
                        var v23 float32 = 0.1
                        _ = v23
                    }
                    _ = fi22
                }
                if (v3 && v3) {
                    v9 = v9
                    var v24 float64 = (0.0 - float64((1e10 - 1.5)))
                    _ = v24
                }
            }
        }
    case "\n\t":
        for fi25 := 0; fi25 < 12; fi25++ {
            _ = len(v5)
            _ = fi25
        }
    case "hello":
        v3 = (true || true)
    default:
        v26 := GenS3{F0: v9, F1: (v3 || ("hello" <= ("foo bar" + ("a" + ""))))}
        _ = v26
    }
    var v27 bool = (true && v3)
    _ = len(v1)
    v5 = append(v5, 953)
    fmt.Println(v27)
    switch -2147483648 {
    case 0:
        v28 := GenS1{F0: 32768, F1: (v9 & v9)}
        _ = v28
    case -1:
        switch 1 {
        case 9223372036854775807:
            var v29 int16 = -319
            _ = v29
        default:
            var v30 int16 = 1000
            _ = v30
        }
    }
    switch -2147483648 {
    case 128:
        v31 := GenS4{F0: 100, F1: 128, F2: false}
        _ = v31
    case 32768:
        v1["\n\t"] = 42
    case 1:
        var v32 uint16 = 32767
        _ = v32
    default:
        if v27 {
            v33 := GenS2{F0: 18446744073709551615, F1: -597, F2: 1}
            var v34 float64 = 0.0
            _ = v33
            _ = v34
        }
    }
    v27 = (65535 >= 1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint32 = 2147483648
        fmt.Println(v27, v2)
        switch (fi35 ^ 0) {
        case -1:
            var v37 int64 = 1
            _ = v37
        default:
            var v38 int32 = v9
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    v9 += v9
    switch -9223372036854775808 {
    case 1:
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
