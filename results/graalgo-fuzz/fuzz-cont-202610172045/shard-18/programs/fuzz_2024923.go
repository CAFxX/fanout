package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8) uint {
    p0 += p0
    v7 := GenS1{F0: -626, F1: 517}
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 bool) int8 {
    switch "hello" {
    case "hello":
        v8 := GenS1{F0: 0, F1: 42}
        _ = v8
    case "":
        v9 := GenS2{F0: 1, F1: 931}
        _ = v9
    default:
        switch ("\n\t" + "foo bar") {
        case "a":
            if (3.14159 <= float64(3.14159)) {
                var v10 int = 32767
                _ = v10
            } else {
                switch -305 {
                case 127:
                    var v11 uint8 = 177
                    _ = v11
                case 613:
                    if (!p1) {
                        v12 := []int{-249, 233}
                        _ = v12
                    }
                }
                p1 = ((p0 | (p0 >> ((330) & 63))) == (p0 / ((-9223372036854775808) | 1)))
                for fi13 := 0; fi13 < 7; fi13++ {
                    p1 = p1
                    _ = fi13
                }
            }
        case "\n\t":
            var v14 int8 = 0
            _ = v14
        case "hello":
            if (32768 < 4294967295) {
                v15 := []int{32768, 14, 2147483647, 1}
                _ = v15
            } else {
                var v16 int8 = 42
                fmt.Println(p0, p0)
                if (!false) {
                    switch 127 {
                    case 1:
                        var v17 string = "\n\t"
                        _ = v17
                    case 42:
                        v16 += (v16 | -128)
                    case 79:
                        var v18 uint8 = 57
                        _ = v18
                    default:
                        var v19 int64 = (p0 + p0)
                        _ = v19
                    }
                } else {
                    if (false && true) {
                        fmt.Println(v16, p1)
                    } else {
                        v20 := []int{-9223372036854775808, -9223372036854775808}
                        _ = v20
                    }
                    v16 -= 127
                    switch ("" + ("a" + "hello")) {
                    case "":
                        var v21 uint64 = 42
                        _ = v21
                    case "foo bar":
                        if ("\n\t" <= ((("foo bar" + "") + "hello") + "hello")) {
                            var v22 uint = 135
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 6; fi23++ {
                                p1 = ((fi23 >> ((genFn0(v16)) & 63)) > fi23)
                                _ = fi23
                            }
                            v24 := GenS1{F0: 1, F1: 1}
                            var v25 uint = genFn0((v16 << ((907) & 63)))
                            _ = v24
                            _ = v25
                        }
                    }
                }
                _ = v16
            }
        }
    }
    p1 = p1
    v26 := []int{4294967295, -109, -759}
    v27 := GenS3{F0: 18446744073709551615}
    v26 = append(v26, 667)
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 127
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 int64 = 42
    var v5 int32 = 1
    var v6 int = -927
    switch "hello" {
    case "foo bar":
        fmt.Println(v4, v4)
    case "\n\t":
        var v28 uint8 = 42
        _ = v28
    default:
        var v29 uint8 = 127
        _ = v29
    }
    switch 2147483647 {
    case 4294967295:
        v30 := GenS4{F0: 981, F1: 1}
        _ = v30
    case 2147483647:
        var v31 string = ("\n\t" + ("\n\t" + "hello"))
        _ = v31
    case 0:
        fmt.Println(v6, v1)
    default:
        var v32 int16 = -479
        _ = v32
    }
    fmt.Println(v1, v1)
    v4 += (v4 - v4)
    _ = len(v3)
    var v33 int8 = int8(v1)
    v2["\n\t"] = 295
    v3 = append(v3, "a")
    fmt.Println(v4, v33, v6)
    v1 -= v1
    if (v5 < v5) {
        v6 += v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
