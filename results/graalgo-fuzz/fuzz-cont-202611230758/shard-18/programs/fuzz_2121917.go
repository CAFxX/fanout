package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint8, p1 bool, p2 uint64) int64 {
    switch 117 {
    case 230:
        v5 := []int{-885, 1}
        _ = v5
    case 4294967295:
        if (127 > 42) {
            for fi6 := 0; fi6 < 10; fi6++ {
                p1 = (true && false)
                for fi7 := 0; fi7 < 8; fi7++ {
                    p2 += (^(p2))
                    v8 := GenS1{F0: 0}
                    p0 -= 128
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint = 4294967295
                    if true {
                        v10 -= v10
                        switch fi6 {
                        case 42:
                            v10 -= 0
                        default:
                            fmt.Println(fi6, fi6)
                        }
                        var v11 int32 = -1
                        switch (fi9 + fi9) {
                        case 42:
                            p1 = (!true)
                        case -221:
                            if p1 {
                                switch fi9 {
                                case -963:
                                    v10 = v10
                                case 2147483647:
                                    var v12 int32 = 2147483647
                                    _ = v12
                                }
                            } else {
                                if ((2147483648 > 0) && p1) {
                                    v13 := GenS2{F0: 65535}
                                    var v14 bool = (v10 >= 65535)
                                    fmt.Println(v11, p0, p2)
                                    _ = v13
                                    _ = v14
                                }
                            }
                        case 32767:
                            v15 := []int{fi9, fi9, fi6}
                            _ = v15
                        }
                        var v16 bool = (58 > 125)
                        _ = v11
                        _ = v16
                    } else {
                        v17 := []int{(fi9 ^ fi9), fi6, (^(fi9))}
                        _ = v17
                    }
                    if true {
                        v10 += (v10 * v10)
                        v18 := GenS1{F0: 0}
                        var v19 int32 = 2147483647
                        _ = v18
                        _ = v19
                    } else {
                        v20 := []int{-2147483648, (fi6 * (fi9 * (fi6 + 42)))}
                        _ = v20
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi6
            }
        }
    case -335:
        p2 = p2
    }
    var v21 float32 = (1.5 + 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 638
}
func genFn1() int8 {
    var v22 uint64 = 0
    var v23 int16 = -1
    fmt.Println(v22)
    switch -688 {
    case 42:
        v23 = v23
    default:
        var v24 int32 = -1
        _ = v24
    }
    _ = v22
    _ = v23
    return 0
}

func main() {
    var v1 uint = 1
    var v2 int8 = 127
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    v2 -= genFn1()
    if true {
        v1 -= v1
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        fmt.Println(v1, fi25, fi25)
        v1 -= v1
        v1 = (v1 & (v1 & (v1 * (v1 << (((^(v1))) & 63)))))
        _ = fi25
    }
    _ = len(v4)
    var v26 int8 = (v2 << ((v1) & 63))
    switch -468 {
    case 567:
        v1 += (v1 | v1)
    }
    var v27 int = 2147483647
    v28 := GenS3{F0: 0, F1: ("hello" + ("" + "foo bar"))}
    var v29 uint = v1
    v27 -= v27
    v29 += v1
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 float64 = 3.14159
        v31 -= v31
        v31 -= v31
        _ = fi30
        _ = v31
    }
    var v32 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
