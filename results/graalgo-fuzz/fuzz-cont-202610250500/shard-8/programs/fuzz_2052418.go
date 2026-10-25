package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -29
    var v3 int16 = 1
    v4 := make(map[int]bool)
    v5 := []uint32{455}
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                v2 -= v2
                v2 -= v2
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v7 int32 = 241
                if ((true && (!((v3 ^ 828) == v3))) || ((!false) && false)) {
                    switch (fi6 ^ (fi6 | fi6)) {
                    case -839:
                        switch "a" {
                        case "\n\t":
                            v8 := GenS1{F0: 42, F1: (("a" + "hello") + "a")}
                            _ = v8
                        default:
                            v9 := GenS2{F0: fi6, F1: 42}
                            _ = v9
                        }
                    case 4294967295:
                        switch fi6 {
                        case 1:
                            var v10 int64 = -1
                            _ = v10
                        case 584:
                            var v11 int32 = v7
                            _ = v11
                        }
                    case 42:
                        v12 := GenS3{F0: 1, F1: int64(v3)}
                        _ = v12
                    }
                    var v13 int16 = -1
                    _ = v13
                } else {
                    var v14 uint8 = 128
                    _ = v14
                }
                _ = v7
            }
            if (v2 < 100) {
                switch fi6 {
                case 4294967295:
                    v3 = (v3 + int16(fi6))
                case -1:
                    var v15 int32 = -1
                    _ = v15
                }
                var v16 int16 = v3
                _ = v16
            }
            switch (fi6 / ((32768) | 1)) {
            case 1:
                var v17 int16 = v3
                _ = v17
            case 9223372036854775807:
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS3{F0: 9223372036854775807, F1: -207}
                    _ = fi18
                    _ = v19
                }
            default:
                v2 = v2
            }
            _ = fi6
        }
    }
    v20 := v1[2147483647]
    var v21 int16 = v3
    v22 := GenS1{F0: 794, F1: "\n\t"}
    switch v20 {
    case 0:
        v23 := GenS4{F0: 65535, F1: v21, F2: 255}
        _ = v23
    case 108:
        fmt.Println(v3, v2)
    }
    var v24 uint64 = 42
    v25 := v1[704]
    v24 -= 1
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint16 = 42
        _ = fi26
        _ = v27
    }
    v28 := GenS5{F0: 65535, F1: (^(v24)), F2: v21}
    if len(v5) > 0 {
        v5[0] = 0
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v29 float64 = (1e10 * (1.5 - 1e10))
    if false {
        var v30 int32 = 2147483647
        fmt.Println(v21, v21)
        _ = v30
    }
    var v31 int32 = 0
    switch v25 {
    case -237:
        v31 += v31
    case -558:
        fmt.Println(v24, v20)
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 uint64 = (v24 >> ((4294967295) & 63))
            fmt.Println(v3, v34)
            v21 = (^(v3))
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 string = "hello"
            var v37 uint16 = 32767
            if (!(!(true || (true && false)))) {
                v36 = v36
            }
            delete(v4, 134)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi32
    }
    if false {
        v38 := v4[(v25 | (v25 + 127))]
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
