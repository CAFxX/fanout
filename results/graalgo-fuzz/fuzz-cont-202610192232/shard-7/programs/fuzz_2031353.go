package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 uint, p1 uint, p2 uint) uint64 {
    p1 = (p0 + 65535)
    var v4 int = -634
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 uint64 = 0
    var v2 int64 = -9223372036854775808
    var v3 int64 = v2
    switch 4294967295 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 14; fi5++ {
            v3 = v2
            v2 += -20
            _ = fi5
        }
    }
    var v6 uint32 = 2147483647
    var v7 uint16 = 585
    v8 := GenS1{F0: (v2 - 9223372036854775807), F1: v1}
    v9 := GenS1{F0: (^(v2)), F1: (v1 / (((v1 - v1)) | 1))}
    v10 := []int{122}
    var v11 uint16 = (v7 / ((v7) | 1))
    var v12 uint64 = (v1 * (^(v1)))
    var v13 uint = 2
    fmt.Println(v13, v12, v6)
    switch "a" {
    case "hello":
        switch 32767 {
        case 2147483647:
            var v14 int8 = int8(v13)
            _ = v14
        }
    case "foo bar":
        var v15 uint16 = (^(v7))
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 bool = false
            var v18 uint16 = v11
            v1 = v12
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 int64 = v2
    v20 := GenS1{F0: (v19 * (v2 - 927)), F1: genFn0(v13, 0, v13)}
    var v21 int64 = 1
    if (-2.25 < (-2.25 - -2.25)) {
        switch "foo bar" {
        case "":
            if true {
                if (!true) {
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
                var v23 int = 2147483647
                var v24 uint64 = v1
                _ = v23
                _ = v24
            } else {
                v11 += v7
                if (uint32(v12) == 0) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                } else {
                    v25 := GenS1{F0: (v2 << ((65535) & 63)), F1: 913}
                    _ = v25
                }
                if (49 < 42) {
                    if len(v10) > 0 {
                        v10[0] = 127
                    }
                } else {
                    var v26 string = ("foo bar" + "a")
                    _ = v26
                }
            }
        case "a":
            var v27 float32 = 3.14159
            _ = v27
        case "\n\t":
            _ = len(v10)
        default:
            var v28 int16 = -958
            _ = v28
        }
        fmt.Println(v12)
    } else {
        v3 += v19
        v10 = append(v10, 42)
        v29 := GenS2{F0: false, F1: v21, F2: (v13 >> ((uint(v1)) & 63))}
        _ = v29
    }
    v30 := GenS3{F0: 2147483648, F1: "foo bar"}
    for fi31 := 0; fi31 < 1; fi31++ {
        v2 = (-(v21))
        for fi32 := 0; fi32 < 19; fi32++ {
            v2 = v19
            _ = fi32
        }
        _ = fi31
    }
    v33 := GenS4{F0: 1, F1: -109, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v33
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
