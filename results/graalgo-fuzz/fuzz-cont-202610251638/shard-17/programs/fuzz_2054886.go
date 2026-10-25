package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 int, p1 uint8) string {
    p0 = (p0 * (-(p0)))
    p0 = p0
    _ = p0
    _ = p1
    return ("foo bar" + "hello")
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483648, 2147483648, 2147483647, 42}
    v3 := make(map[uint32]uint32)
    if (128 == 214) {
        var v4 string = "a"
        if (-1 >= 100) {
            if false {
                v4 = ("\n\t" + v4)
            }
        } else {
            v4 = "hello"
            var v5 int16 = -32768
            v6 := GenS1{F0: 1, F1: (1e10 < 0.1)}
            _ = v5
            _ = v6
        }
        _ = v4
    }
    var v7 uint64 = 624
    v8 := GenS1{F0: 203, F1: true}
    switch -67 {
    case 128:
        switch -1 {
        case 193:
            var v9 int16 = 1
            _ = v9
        case -9223372036854775808:
            fmt.Println(v7)
        case -935:
            fmt.Println(v7)
        default:
            var v10 int64 = 9223372036854775807
            _ = v10
        }
    case 4294967295:
        v11 := GenS1{F0: 255, F1: true}
        _ = v11
    case 345:
        v12 := GenS1{F0: 42, F1: true}
        _ = v12
    default:
        v2 = append(v2, 292)
    }
    if false {
        var v13 int64 = -1
        var v14 uint = 1
        v15 := v3[792]
        _ = v13
        _ = v14
        _ = v15
    }
    v7 = (^(v7))
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(v7, v7)
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(fi16, fi16, fi17)
            v18 := GenS1{F0: 145, F1: (32767 > 0)}
            if false {
                v19 := GenS1{F0: 128, F1: (!false)}
                _ = v19
            }
            var v20 float64 = 1e10
            v20 -= ((v20 / (v20 / v20)) + 0.1)
            v21 := GenS1{F0: 255, F1: (int64(fi16) == 0)}
            _ = fi17
            _ = v18
            _ = v20
            _ = v21
        }
        v7 -= v7
        _ = fi16
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v7 = v7
        v23 := GenS1{F0: 42, F1: true}
        _ = fi22
        _ = v23
    }
    if false {
        var v24 int32 = 2147483647
        _ = v24
    } else {
        _ = len(v1)
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint64 = 521
            _ = fi25
            _ = v26
        }
    }
    v27 := GenS2{F0: -1, F1: -32768, F2: 32767}
    v7 -= v7
    for fi28 := 0; fi28 < 1; fi28++ {
        if true {
            var v29 uint32 = 65535
            _ = v29
        } else {
            var v30 uint64 = (^(v7))
            v7 = v30
            v30 += v7
            _ = v30
        }
        switch fi28 {
        case -837:
            v31 := GenS2{F0: 0, F1: -1, F2: -32768}
            _ = v31
        case 42:
            var v32 int = fi28
            _ = v32
        case -2147483648:
            var v33 bool = false
            _ = v33
        }
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
