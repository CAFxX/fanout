package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0(p0 uint32, p1 uint8, p2 int32) string {
    fmt.Println(p0, p0, p0)
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "" {
        case "":
            p2 -= p2
        }
        _ = fi6
    }
    v7 := GenS1{F0: 9223372036854775807, F1: (p2 * p2)}
    v8 := GenS1{F0: 176, F1: p2}
    var v9 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ("hello" + (((("\n\t" + "hello") + "") + "") + ""))
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]bool)
    var v3 int8 = 42
    var v4 uint16 = 65535
    v5 := []uint32{343, 478, 2147483648}
    v10 := GenS2{F0: 0}
    v11 := GenS2{F0: 42}
    var v12 int64 = -256
    v12 = (v12 << ((132) & 63))
    v13 := GenS3{F0: 2147483648}
    v14 := GenS4{F0: 53}
    v4 += (^(v4))
    v3 = v3
    switch -404 {
    case -9223372036854775808:
        v15 := GenS5{F0: -133, F1: v12}
        _ = v15
    case 1:
        if len(v5) > 0 {
            v5[0] = 798
        }
    default:
        for fi16 := 0; fi16 < 18; fi16++ {
            delete(v2, (fi16 >> ((65535) & 63)))
            v2[fi16] = false
            _ = fi16
        }
    }
    switch ("" + "foo bar") {
    case "a":
        var v17 uint8 = 207
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 7; fi18++ {
            v12 += v12
            if false {
                fmt.Println(v3, v3, v3)
                for fi19 := 0; fi19 < 6; fi19++ {
                    v20 := GenS3{F0: 65535}
                    switch fi19 {
                    case -2147483648:
                        if true {
                            var v21 int64 = 1
                            v22 := GenS1{F0: (fi18 + fi18), F1: 1}
                            v2[(-(fi19))] = (!(fi19 > (fi19 & fi19)))
                            _ = v21
                            _ = v22
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            v4 = (^(v4))
                            fmt.Println(fi18)
                            if true {
                                _ = len(v2)
                                v23 := GenS1{F0: 32767, F1: 1}
                                _ = v23
                            }
                            fmt.Println(v4)
                        }
                    }
                    var v24 int32 = 0
                    v24 -= v24
                    if ((0 != 1) && (1 <= 18446744073709551615)) {
                        v25 := GenS6{F0: 18446744073709551615, F1: 128, F2: false}
                        var v26 uint16 = 42
                        v24 -= 42
                        _ = v25
                        _ = v26
                    } else {
                        v4 -= 0
                        v24 += (^(v24))
                        v4 -= (v4 + (v4 ^ (v4 + (v4 * v4))))
                    }
                    _ = fi19
                    _ = v20
                    _ = v24
                }
            }
            _ = fi18
        }
    default:
        fmt.Println(v12)
    }
    v12 += -367
    for fi27 := 0; fi27 < 10; fi27++ {
        switch (-(fi27)) {
        case 9223372036854775807:
            var v28 float32 = 0.1
            _ = v28
        default:
            var v29 uint = 374
            _ = v29
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
