package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }

func main() {
    v1 := []uint32{4294967295, 90, 826}
    v2 := make(map[string]int)
    var v3 int64 = 0
    var v4 uint16 = 65535
    v5 := make(map[string]int)
    var v6 int8 = 91
    switch "\n\t" {
    case "hello":
        if (42 <= (v3 / ((v3) | 1))) {
            v7 := GenS1{F0: v6, F1: 42}
            delete(v5, "hello")
            v3 -= (-(v3))
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            v4 = v4
            v9 := GenS2{F0: 0, F1: (^(v6)), F2: -409}
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        switch "foo bar" {
        case "":
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS2{F0: 2, F1: v6, F2: -1}
                v3 -= (v3 | int64(fi11))
                v4 -= 2
                _ = fi11
                _ = v12
            }
        }
        if false {
            var v13 int8 = -128
            _ = v13
        } else {
            var v14 bool = (!false)
            if len(v1) > 0 {
                v1[0] = 76
            }
            if ((3.14159 != ((3.14159 + (-2.25 + 0.0)) * (3.14159 - 0.1))) && ((^(v4)) >= 360)) {
                switch "hello" {
                case "":
                    fmt.Println(fi10, fi10)
                default:
                    v14 = (255 < 99)
                }
            }
            _ = v14
        }
        _ = fi10
    }
    var v15 uint8 = 127
    if true {
        if false {
            if false {
                v6 = (^(v6))
            }
            var v16 uint8 = v15
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v16
        }
        v17 := GenS2{F0: 2, F1: (v6 | v6), F2: -2147483648}
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS2{F0: 2147483647, F1: v6, F2: 42}
            fmt.Println(v3, v4, v3)
            v20 := GenS1{F0: (v6 & v6), F1: (fi18 | 384)}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    } else {
        v6 = (v6 + (v6 * v6))
    }
    var v21 int64 = v3
    v22 := GenS1{F0: v6, F1: 32768}
    var v23 uint32 = 4294967295
    if ((!true) || true) {
        if (0.1 > 1e10) {
            var v24 string = ("hello" + "a")
            _ = v24
        }
    }
    if true {
        if (v6 < (v6 << ((2) & 63))) {
            var v25 int = -1
            _ = v25
        }
        v15 += (v15 << ((86) & 63))
    }
    var v26 bool = (((v4 - (v4 | v4)) == v4) || ((v4 ^ 407) == (v4 | (v4 / ((343) | 1)))))
    v26 = true
    var v27 int = -9223372036854775808
    v1 = append(v1, v23)
    v27 = (v27 / ((v27) | 1))
    if v26 {
        fmt.Println(v6, v6)
        for fi28 := 0; fi28 < 7; fi28++ {
            fmt.Println(v3, v27)
            var v29 int16 = -176
            v30 := GenS3{F0: (v15 & v15), F1: (v4 / (((v4 - 0)) | 1)), F2: v6}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v31 := GenS4{F0: 0, F1: (v4 & (v4 / ((v4) | 1))), F2: v21}
    fmt.Println(v26, v23, v15)
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
