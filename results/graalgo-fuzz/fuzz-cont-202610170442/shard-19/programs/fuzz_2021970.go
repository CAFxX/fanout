package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
func genFn0(p0 uint64, p1 string, p2 string) uint16 {
    if (0 > 0) {
        p0 = (p0 - (p0 | 4294967295))
        switch 32767 {
        case 42:
            fmt.Println(p1, p1)
        }
        fmt.Println(p2, p1, p0)
    }
    v6 := GenS1{F0: 185, F1: (^(p0))}
    v7 := []int{}
    var v8 int32 = 653
    for fi9 := 0; fi9 < 15; fi9++ {
        _ = len(v7)
        p0 -= 1
        var v10 uint = uint(p0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    v5 := []uint32{0, 100, 2147483648, 2147483647}
    v1 = v1
    var v11 uint32 = 2147483648
    delete(v3, -1)
    switch 78 {
    case 1:
        fmt.Println(v1)
    }
    var v12 int64 = v4
    v4 = v12
    fmt.Println(v12)
    v13 := GenS2{F0: v1, F1: v1, F2: genFn0(v1, "hello", "\n\t")}
    var v14 uint = 255
    fmt.Println(v11, v14)
    if false {
        switch "\n\t" {
        case "a":
            v5 = append(v5, (v11 | v11))
        }
        if true {
            var v15 bool = false
            _ = v15
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                delete(v2, (v11 << (((^(v14))) & 63)))
                if false {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if true {
                            if len(v5) > 0 {
                                v5[0] = 0
                            }
                        }
                        var v18 int16 = -32768
                        _ = fi17
                        _ = v18
                    }
                    var v19 uint16 = 65535
                    if false {
                        switch fi16 {
                        case 517:
                            switch ("" + "hello") {
                            case "hello":
                                v20 := GenS2{F0: v1, F1: v1, F2: v19}
                                _ = v20
                            case "foo bar":
                                _ = len(v5)
                            default:
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    v22 := GenS2{F0: v1, F1: 0, F2: (v19 >> ((v14) & 63))}
                                    fmt.Println(v14, fi21)
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        }
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 int8 = 112
                            _ = fi23
                            _ = v24
                        }
                        if (int(v19) >= -2147483648) {
                            v4 -= (v4 ^ -21)
                            delete(v2, 608)
                            v1 += v1
                        }
                    } else {
                        if true {
                            v4 = (v12 >> ((232) & 63))
                            delete(v2, (v11 + v11))
                            var v25 int64 = 0
                            _ = v25
                        } else {
                            fmt.Println(v4, v14, v1)
                            var v26 int8 = -47
                            _ = v26
                        }
                    }
                    _ = v19
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 float64 = 3.14159
                    var v29 bool = true
                    var v30 uint32 = v11
                    switch (fi27 << ((42) & 63)) {
                    case 634:
                        switch int(v30) {
                        case 80:
                            v28 += 0.1
                        default:
                            switch "foo bar" {
                            case "a":
                                v31 := GenS3{F0: 255, F1: "\n\t", F2: ("\n\t" + ("hello" + ""))}
                                _ = v31
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi16
            }
        }
    } else {
        if len(v5) > 0 {
            v5[0] = uint32(v14)
        }
        v4 = 566
    }
    if false {
        fmt.Println(v12)
        v1 += uint64(v14)
        delete(v3, -744)
    } else {
        v12 += -9223372036854775808
        v12 -= -91
        var v32 uint64 = v1
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
