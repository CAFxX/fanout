package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int8, p1 bool, p2 int64) uint {
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            p2 -= p2
        case "hello":
            if (!true) {
                v4 := GenS1{F0: 2, F1: p2}
                _ = v4
            } else {
                v5 := GenS1{F0: 0, F1: 235}
                _ = v5
            }
        default:
            v6 := []int{9223372036854775807, -9223372036854775808}
            _ = v6
        }
    default:
        if false {
            if (2 < 255) {
                var v7 string = ("hello" + "")
                if ((!p1) && true) {
                    if (float32(1.5) >= 0.0) {
                        v7 = ((("" + v7) + v7) + v7)
                    }
                } else {
                    var v8 int8 = (p0 % ((p0) | 1))
                    v9 := []int{-9223372036854775808}
                    if (358 == 1) {
                        p0 = p0
                        fmt.Println(p1, p1)
                        v10 := GenS2{F0: 0, F1: 4294967295}
                        _ = v10
                    }
                    _ = v8
                    _ = v9
                }
                if p1 {
                    p2 -= -686
                    if (!p1) {
                        p0 -= 127
                        v11 := []int{-2147483648, 128}
                        if len(v11) > 0 {
                            v11[0] = -643
                        }
                        _ = v11
                    } else {
                        var v12 uint16 = 32767
                        v13 := GenS1{F0: 210, F1: 0}
                        _ = v12
                        _ = v13
                    }
                }
                _ = v7
            } else {
                fmt.Println(p0, p2, p1)
                if (!true) {
                    switch 127 {
                    case 856:
                        v14 := []int{-2147483648, 0, -1, 1}
                        _ = v14
                    case 4294967295:
                        v15 := []int{-2147483648, -9223372036854775808, 128}
                        _ = v15
                    }
                    var v16 uint32 = 2147483647
                    _ = v16
                } else {
                    v17 := GenS1{F0: 4294967295, F1: (p2 / (((p2 ^ (p2 >> ((273) & 63)))) | 1))}
                    _ = v17
                }
                var v18 float32 = (((1e10 * float32(1e10)) * -2.25) - -2.25)
                _ = v18
            }
        }
    }
    p2 += p2
    p1 = p1
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 float32 = (float32(0.1) * 1e10)
        var v21 int64 = p2
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint16 = 756
    var v2 uint64 = 0
    var v3 uint = 580
    v22 := GenS2{F0: 42, F1: 0}
    fmt.Println(v2, v1)
    if (372 <= 2147483647) {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := []int{0}
            var v25 uint64 = (v2 >> (((v3 / ((v3) | 1))) & 63))
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        v3 = genFn0(100, ((!false) || ((!((1.5 - (-2.25 - 0.0)) != -2.25)) && (!false))), 567)
    }
    switch 127 {
    case 0:
        v26 := []int{1, 629, -630, -2147483648}
        _ = v26
    case 32767:
        v27 := []int{-2147483648, -537, -725}
        _ = v27
    }
    var v28 uint16 = (v1 + 488)
    v29 := []int{128}
    v30 := GenS2{F0: 32767, F1: v2}
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    if ("foo bar" < "") {
        v3 -= (v3 >> ((v3) & 63))
    } else {
        var v31 uint64 = v2
        _ = v31
    }
    fmt.Println(v3)
    var v32 uint32 = 1
    var v33 float32 = (-2.25 * 0.1)
    if false {
        fmt.Println(v1, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
}
