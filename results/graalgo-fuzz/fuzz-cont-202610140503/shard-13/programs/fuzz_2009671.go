package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := []uint32{4294967295}
    var v2 uint64 = 18446744073709551615
    var v3 uint = 42
    v4 := []int{-1, 924, 0}
    var v5 string = "hello"
    var v6 int16 = 1
    fmt.Println(v3, v3)
    v7 := GenS1{F0: 42}
    var v8 bool = ((!true) || false)
    v9 := GenS1{F0: 128}
    v6 = (v6 + 90)
    var v10 int8 = 1
    if ((v10 << (((v3 / ((v3) | 1))) & 63)) <= (v10 | v10)) {
        v11 := GenS1{F0: uint8(v2)}
        _ = v11
    }
    fmt.Println(v3, v10)
    v6 = v6
    v12 := GenS2{F0: -9223372036854775808}
    var v13 int32 = 1
    if (v3 <= (v3 | v3)) {
        if (v10 == 35) {
            v14 := GenS3{F0: (v10 ^ -1), F1: 127}
            for fi15 := 0; fi15 < 20; fi15++ {
                if false {
                    v16 := GenS1{F0: 44}
                    _ = v16
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v5 = v5
                if len(v4) > 0 {
                    v4[0] = fi15
                }
                _ = fi15
            }
            var v17 uint = v3
            _ = v14
            _ = v17
        } else {
            v18 := GenS4{F0: 870}
            _ = v18
        }
        v3 = v3
    }
    if (!(0.0 != -2.25)) {
        v13 = 1
        var v19 uint8 = 1
        for fi20 := 0; fi20 < 1; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi20
        }
        _ = v19
    }
    _ = len(v4)
    if false {
        v21 := GenS5{F0: -9223372036854775808, F1: (127 != 93)}
        v2 -= v2
        var v22 uint8 = 42
        _ = v21
        _ = v22
    } else {
        v4 = append(v4, 0)
        v23 := GenS6{F0: 0}
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v6)
        if (159 < 1) {
            switch (fi24 + fi24) {
            case 42:
                v6 -= v6
            case 9223372036854775807:
                var v25 string = ("\n\t" + "foo bar")
                _ = v25
            case 32768:
                v26 := GenS5{F0: -1, F1: v8}
                _ = v26
            default:
                v1 = append(v1, 827)
            }
            var v27 uint16 = 0
            _ = v27
        }
        _ = fi24
    }
    v28 := GenS4{F0: 1}
    v29 := GenS6{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
