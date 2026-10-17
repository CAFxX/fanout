package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int8 = 1
    v4 := make(map[int]uint32)
    var v5 int64 = -9223372036854775808
    switch -1 {
    case 9223372036854775807:
        v6 := []int{9223372036854775807, -1}
        _ = v6
    case 32767:
        var v7 string = ("\n\t" + ("hello" + ("a" + "\n\t")))
        _ = v7
    }
    var v8 int = 42
    v5 = (v5 * (^(v5)))
    if (-2.25 == 0.0) {
        fmt.Println(v3, v8, v5)
        v9 := GenS1{F0: -991, F1: (-(v8))}
        _ = v9
    } else {
        switch 32767 {
        case 128:
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    fmt.Println(v3, v8)
                    var v11 int16 = 1
                    _ = fi10
                    _ = v11
                }
                v5 += v5
            } else {
                var v12 int = (v8 | 360)
                _ = v12
            }
        default:
            v3 = (v3 << ((42) & 63))
        }
    }
    v3 = 42
    v3 = v3
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint32 = 4294967295
        var v15 int64 = 9223372036854775807
        var v16 uint64 = 1
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint16 = 255
    v18 := []int{v8, 164, v8}
    var v19 uint8 = 2
    switch v8 {
    case 1:
        for fi20 := 0; fi20 < 4; fi20++ {
            switch 32767 {
            case 2147483647:
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch 32768 {
                    case 2147483647:
                        fmt.Println(fi21, v17)
                    case 0:
                        v5 = v5
                    }
                    fmt.Println(v5)
                    v22 := GenS1{F0: -1, F1: -974}
                    _ = fi21
                    _ = v22
                }
            case 1:
                v23 := GenS2{F0: 963, F1: (^(v8))}
                _ = v23
            case 32768:
                for fi24 := 0; fi24 < 15; fi24++ {
                    v5 += 42
                    if (0 <= 1) {
                        if len(v18) > 0 {
                            v18[0] = v8
                        }
                    } else {
                        v25 := GenS1{F0: 1, F1: (^(fi20))}
                        var v26 int16 = 42
                        switch (v8 * fi24) {
                        case 2147483647:
                            delete(v4, (v8 >> ((255) & 63)))
                            if false {
                                v27 := GenS1{F0: v26, F1: (fi24 >> ((42) & 63))}
                                fmt.Println(v3, v8, v8)
                                _ = v27
                            }
                        case 4294967295:
                            v8 -= 9223372036854775807
                        case 438:
                            _ = len(v18)
                        }
                        _ = v25
                        _ = v26
                    }
                    var v28 int8 = v3
                    _ = fi24
                    _ = v28
                }
            default:
                _ = len(v2)
            }
            _ = fi20
        }
    case -9223372036854775808:
        v8 = v8
    default:
        v29 := GenS2{F0: -32768, F1: v8}
        _ = v29
    }
    v30 := GenS3{F0: 317, F1: 15}
    if len(v18) > 0 {
        v18[0] = 127
    }
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
