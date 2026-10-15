package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 uint16, p2 int32) uint8 {
    var v5 int32 = 865
    v6 := GenS1{F0: 475, F1: -1, F2: 42}
    v5 += v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := []int{}
    var v2 uint16 = 2
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 0
        if true {
            fmt.Println(v3)
            v9 := GenS1{F0: 2147483648, F1: -128, F2: int8(v2)}
            _ = v9
        }
        if ("\n\t" >= "\n\t") {
            v10 := v4[fi7]
            _ = v10
        } else {
            fmt.Println(v3, v8)
        }
        _ = fi7
        _ = v8
    }
    var v11 int8 = 0
    v12 := v4[9223372036854775807]
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, v2)
        for fi14 := 0; fi14 < 4; fi14++ {
            if (("hello" >= ("" + "hello")) && true) {
                switch fi14 {
                case 32767:
                    v3 += v3
                case 4294967295:
                    if true {
                        v4[fi14] = v12
                        v2 = v2
                        fmt.Println(fi13, v11)
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            var v15 uint16 = (v2 >> ((616) & 63))
                            _ = v15
                        }
                        v16 := GenS1{F0: uint32(v2), F1: (v11 >> ((65535) & 63)), F2: (v11 ^ v11)}
                        _ = v16
                    }
                case 0:
                    if true {
                        v17 := GenS1{F0: v12, F1: (v11 << ((196) & 63)), F2: v11}
                        fmt.Println(v2, v3)
                        v18 := GenS1{F0: (^(v12)), F1: v11, F2: (v11 % ((v11) | 1))}
                        _ = v17
                        _ = v18
                    }
                default:
                    fmt.Println(v3)
                }
                var v19 uint16 = 42
                _ = v19
            } else {
                switch "foo bar" {
                case "hello":
                    v11 += 0
                default:
                    v20 := GenS1{F0: v12, F1: (v11 + v11), F2: 79}
                    _ = v20
                }
                fmt.Println(v3, v3)
                v3 -= v3
            }
            fmt.Println(v12)
            _ = fi14
        }
        delete(v4, fi13)
        _ = fi13
    }
    v11 += (v11 % ((v11) | 1))
    var v21 uint8 = uint8(v12)
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS2{F0: v3}
        var v24 int32 = 670
        if true {
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := GenS2{F0: (v24 - (v24 - v3))}
                fmt.Println(v11)
                fmt.Println(v12, fi22, v2)
                var v27 int64 = 9223372036854775807
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 uint = uint(v21)
            _ = v28
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    v12 += v12
    var v29 uint64 = 229
    var v30 int32 = (v3 ^ (v3 - 2147483647))
    if true {
        v4[564] = 2147483647
        var v31 int16 = 926
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(v29)
            var v33 uint16 = (v2 % ((uint16(v12)) | 1))
            _ = fi32
            _ = v33
        }
        _ = v31
    } else {
        var v34 uint8 = genFn0(1000, v2, -725)
        v2 += v2
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS2{F0: 1}
        v30 += (v30 + -1)
        v37 := GenS2{F0: v3}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint32 = v12
    var v39 float64 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
}
