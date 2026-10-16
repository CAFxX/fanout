package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 int8 }
type GenS4 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 0
    v3 := []string{("hello" + ("foo bar" + "\n\t")), "\n\t", "foo bar"}
    var v4 uint16 = 42
    v4 = uint16(v2)
    v5 := GenS1{F0: -29, F1: 0, F2: "foo bar"}
    if (0 != 0) {
        var v6 bool = (false && true)
        fmt.Println(v4)
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float64 = (1.5 + 1.5)
            v2 = (v2 + v2)
            var v9 int = fi7
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        var v10 uint64 = 2
        fmt.Println(v4)
        _ = v10
    }
    switch "" {
    case "":
        v4 = v4
    }
    fmt.Println(v2, v4)
    fmt.Println(v2, v4)
    v3 = append(v3, "hello")
    var v11 bool = (0 != 4294967295)
    if false {
        for fi12 := 0; fi12 < 14; fi12++ {
            v4 = v4
            v1[0] = 2147483647
            var v13 uint32 = 2147483647
            _ = fi12
            _ = v13
        }
        switch "a" {
        case "foo bar":
            var v14 uint = v2
            _ = v14
        default:
            if (1 < 0) {
                v4 -= v4
                v4 = v4
                var v15 string = "foo bar"
                _ = v15
            } else {
                v2 = (v2 << ((548) & 63))
            }
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            if (!true) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v11 = (v11 && v11)
                    _ = fi17
                }
            } else {
                switch "hello" {
                case "hello":
                    v18 := v1[fi16]
                    _ = v18
                default:
                    v4 -= 2
                }
                fmt.Println(v4, v4)
            }
            v19 := GenS2{F0: 0, F1: -109}
            if (!(-1 <= 9223372036854775807)) {
                var v20 uint16 = (v4 / (((^(v4))) | 1))
                _ = v20
            }
            _ = fi16
            _ = v19
        }
    } else {
        var v21 int64 = 42
        v21 = v21
        fmt.Println(v21, v4)
        _ = v21
    }
    if v11 {
        var v22 int = int(v4)
        for fi23 := 0; fi23 < 19; fi23++ {
            v4 = (^(v4))
            fmt.Println(fi23)
            var v24 string = "\n\t"
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        var v25 int8 = 0
        if v11 {
            switch "foo bar" {
            case "":
                v26 := GenS1{F0: v25, F1: 587, F2: "\n\t"}
                _ = v26
            }
            var v27 uint8 = 97
            if v11 {
                var v28 int64 = 42
                var v29 uint8 = (v27 ^ v27)
                v30 := GenS3{F0: v28, F1: -2147483648, F2: v25}
                _ = v28
                _ = v29
                _ = v30
            } else {
                _ = len(v1)
            }
            _ = v27
        }
        var v31 string = ""
        _ = v25
        _ = v31
    }
    var v32 uint = uint(v4)
    v33 := GenS2{F0: 0, F1: -63}
    v34 := GenS4{F0: 0, F1: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
