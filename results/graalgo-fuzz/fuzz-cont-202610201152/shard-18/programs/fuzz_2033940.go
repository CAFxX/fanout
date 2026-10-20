package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 42
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    v6 := make(map[string]uint32)
    v7 := []int{128}
    var v8 uint8 = 1
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
    } else {
        fmt.Println(v2, v8)
    }
    switch "a" {
    case "hello":
        if true {
            var v11 uint8 = (v8 % ((70) | 1))
            _ = v11
        }
    case "foo bar":
        if (false && true) {
            v12 := GenS1{F0: 42}
            fmt.Println(v2)
            if len(v7) > 0 {
                v7[0] = 32767
            }
            _ = v12
        }
    }
    if (-32768 <= 42) {
        var v13 uint64 = 0
        _ = len(v4)
        _ = v13
    }
    v14 := v5[32767]
    v15 := GenS2{F0: 1, F1: 255}
    if false {
        fmt.Println(v8, v8)
        delete(v3, 4294967295)
        switch "\n\t" {
        case "foo bar":
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int8 = -71
                v14 = "a"
                if false {
                    if (2 <= 309) {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            if (false && (-32768 < -1)) {
                                fmt.Println(fi16)
                                var v19 uint64 = 986
                                v14 = (((v14 + ("a" + "a")) + "\n\t") + v14)
                                _ = v19
                            }
                            _ = fi18
                        }
                        v2 = 21
                    }
                }
                _ = fi16
                _ = v17
            }
        case "\n\t":
            fmt.Println(v2)
        default:
            v14 = v14
        }
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        if ((uint8(v2) >= v8) && (!((879 != 255) || false))) {
            switch fi20 {
            case 32768:
                for fi21 := 0; fi21 < 16; fi21++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v22 := GenS2{F0: int16(fi20), F1: 42}
                    _ = fi21
                    _ = v22
                }
            case 9223372036854775807:
                v23 := GenS3{F0: 0, F1: uint64(v2), F2: (v2 << ((261) & 63))}
                _ = v23
            case 42:
                v6[v14] = 4294967295
            default:
                var v24 int16 = 32767
                _ = v24
            }
            v2 -= v2
        } else {
            v8 = v8
        }
        v25 := GenS4{F0: uint64(v2), F1: 0}
        var v26 float32 = -2.25
        _ = fi20
        _ = v25
        _ = v26
    }
    if (v2 < (v2 | v2)) {
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v2 += v2
                v29 := GenS2{F0: int16(v2), F1: 941}
                _ = fi28
                _ = v29
            }
            v14 = (v14 + ((("foo bar" + v14) + "hello") + v14))
            _ = fi27
        }
        switch "" {
        case "":
            var v30 int16 = 42
            _ = v30
        case "foo bar":
            v8 = 128
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
