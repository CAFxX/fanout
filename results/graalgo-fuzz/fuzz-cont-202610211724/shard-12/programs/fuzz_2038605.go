package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    var v1 int64 = -86
    v2 := []string{("" + "hello"), ((("hello" + "hello") + "") + "foo bar"), ("foo bar" + "\n\t")}
    var v3 float64 = 0.0
    v4 := []string{(((("" + "foo bar") + ((("a" + "foo bar") + "\n\t") + ("a" + "a"))) + "foo bar") + ("foo bar" + (("foo bar" + "foo bar") + "hello")))}
    var v5 uint32 = 365
    v6 := []uint32{(v5 * v5), v5, v5, 2147483648}
    var v7 float64 = v3
    if ((false && true) && false) {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint32 = (^(v5))
            var v10 uint64 = 18446744073709551615
            v5 = (^(v5))
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: -514}
        _ = v11
    }
    fmt.Println(v7)
    v7 += 3.14159
    if true {
        v3 = v3
        v3 += v3
    } else {
        v12 := GenS2{F0: 31}
        if (42 > 42) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v5 -= v5
                _ = fi13
            }
            v14 := GenS3{F0: -128}
            _ = v14
        }
        switch 32767 {
        case 127:
            fmt.Println(v1, v1)
        case 9223372036854775807:
            v15 := GenS2{F0: (v5 >> ((42) & 63))}
            _ = v15
        case 4294967295:
            for fi16 := 0; fi16 < 16; fi16++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v6 = append(v6, (^(v5)))
                _ = fi16
            }
        default:
            var v17 float32 = (1.5 / 0.1)
            _ = v17
        }
        _ = v12
    }
    var v18 uint8 = 82
    var v19 uint8 = (v18 & v18)
    switch "a" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if ("a" > "hello") {
            v3 = v7
        } else {
            if (!false) {
                v20 := GenS3{F0: -1}
                _ = v20
            } else {
                var v21 uint16 = 2
                if false {
                    var v22 float32 = (1.5 - 3.14159)
                    _ = v22
                } else {
                    var v23 int16 = 0
                    v24 := GenS3{F0: 127}
                    v25 := GenS4{F0: 147, F1: 42, F2: (^(v5))}
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: 42}
                _ = v21
                _ = v26
            }
        }
    }
    var v27 uint64 = uint64(v19)
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
