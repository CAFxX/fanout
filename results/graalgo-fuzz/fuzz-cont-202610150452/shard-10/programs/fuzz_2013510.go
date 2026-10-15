package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    var v5 int8 = 42
    var v6 string = ("hello" + "foo bar")
    v4[42] = 42
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v5, v6)
    case 141:
        for fi7 := 0; fi7 < 18; fi7++ {
            v6 = v6
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: -40, F1: (v3 + (v3 >> ((1) & 63)))}
                _ = fi8
                _ = v9
            }
            if false {
                v5 = v5
                var v10 int = fi7
                v2[449] = ((v6 + v6) + (v6 + (v6 + "")))
                _ = v10
            }
            _ = fi7
        }
    case 549:
        v11 := v4[0]
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: v3}
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: 0, F1: 1, F2: 1}
        fmt.Println(v5)
        var v15 uint32 = 968
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        switch v6 {
        case "\n\t":
            _ = len(v1)
        case "a":
            var v17 string = (v6 + "")
            _ = v17
        case "foo bar":
            var v18 uint8 = 2
            _ = v18
        }
        _ = len(v4)
        if true {
            var v19 uint16 = 32767
            _ = v19
        }
        _ = fi16
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v3 = v3
        if true {
            for fi21 := 0; fi21 < 10; fi21++ {
                v6 = (v6 + "foo bar")
                _ = len(v2)
                v5 += -128
                _ = fi21
            }
            var v22 uint32 = 1
            _ = v22
        }
        _ = len(v1)
        _ = fi20
    }
    var v23 uint8 = 255
    switch "a" {
    case "\n\t":
        v1 = append(v1, (v6 + v6))
    case "":
        var v24 uint8 = (v23 ^ 98)
        _ = v24
    default:
        v3 += (^(v3))
    }
    if false {
        if (("\n\t" != v6) || (!(!false))) {
            var v25 int = 4294967295
            _ = v25
        } else {
            var v26 float64 = 0.1
            v26 = v26
            _ = v26
        }
        var v27 int64 = -1
        _ = v27
    } else {
        v28 := GenS2{F0: 371, F1: 190, F2: uint16(v23)}
        if false {
            if len(v1) > 0 {
                v1[0] = (v6 + v6)
            }
        } else {
            _ = len(v4)
            fmt.Println(v23, v3)
            _ = len(v1)
        }
        _ = v28
    }
    v23 = (v23 / (((v23 + v23)) | 1))
    var v29 int64 = -1
    v30 := GenS3{F0: 1, F1: 42, F2: v29}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
