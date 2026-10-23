package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 int16, p1 uint64) uint {
    for fi6 := 0; fi6 < 16; fi6++ {
        if true {
            p0 = p0
            var v7 int16 = (p0 + (p0 >> ((294) & 63)))
            for fi8 := 0; fi8 < 15; fi8++ {
                switch "a" {
                case "":
                    var v9 int8 = 0
                    _ = v9
                case "hello":
                    var v10 int32 = 1
                    _ = v10
                }
                _ = fi8
            }
            _ = v7
        }
        var v11 float64 = 0.1
        p0 -= 32767
        _ = fi6
        _ = v11
    }
    var v12 float64 = 1.5
    _ = p0
    _ = p1
    _ = v12
    return 255
}

func main() {
    v1 := []int{-501}
    v2 := []uint32{608, 393}
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v13 bool = true
    v13 = v13
    v13 = false
    var v14 int32 = 1
    switch "hello" {
    case "a":
        for fi15 := 0; fi15 < 17; fi15++ {
            switch fi15 {
            case 128:
                v13 = false
            default:
                v13 = false
            }
            v13 = v13
            v16 := GenS1{F0: 42, F1: 0}
            _ = fi15
            _ = v16
        }
    case "":
        if ((0 == 127) && v13) {
            var v17 uint32 = 42
            fmt.Println(v14, v14)
            _ = v17
        } else {
            fmt.Println(v13)
        }
    case "hello":
        v13 = v13
    }
    v18 := GenS1{F0: 1, F1: 175}
    v14 = (^(v14))
    switch int(v14) {
    case 4294967295:
        v13 = v13
    default:
        switch "foo bar" {
        case "a":
            var v19 float64 = 0.1
            _ = v19
        case "foo bar":
            v20 := GenS1{F0: 2, F1: 1}
            _ = v20
        default:
            switch -1 {
            case 4294967295:
                for fi21 := 0; fi21 < 18; fi21++ {
                    fmt.Println(v14, fi21, v14)
                    v22 := GenS2{F0: 4294967295, F1: v14, F2: 42}
                    switch "a" {
                    case "":
                        for fi23 := 0; fi23 < 20; fi23++ {
                            if (false && v13) {
                                fmt.Println(v14, fi23, fi21)
                                v14 = v14
                            } else {
                                switch "\n\t" {
                                case "a":
                                    var v24 int = 332
                                    _ = v24
                                }
                                var v25 int64 = -9223372036854775808
                                var v26 uint16 = 255
                                var v27 uint64 = 329
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = fi23
                        }
                    case "foo bar":
                        v28 := GenS2{F0: 65535, F1: (v14 % (((v14 / ((v14) | 1))) | 1)), F2: -861}
                        _ = v28
                    }
                    _ = fi21
                    _ = v22
                }
            case 568:
                v14 -= 1
            }
        }
    }
    if (((0 > 4294967295) && v13) && (!(v13 && v13))) {
        v29 := GenS1{F0: 2, F1: 240}
        var v30 int32 = -1
        v31 := GenS1{F0: 793, F1: 2}
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
