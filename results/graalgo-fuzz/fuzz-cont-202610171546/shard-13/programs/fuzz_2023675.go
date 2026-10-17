package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v5 float64 = (((3.14159 - 1.5) - 1.5) - (0.1 - (-2.25 / 1.0)))
    var v6 float64 = -2.25
    var v7 int16 = -144
    fmt.Println(v6, v5)
    var v8 int16 = v7
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 string, p1 uint16, p2 int8) int {
    var v9 uint = 18446744073709551615
    var v10 uint16 = 65535
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{65535, 2147483647}
    switch "hello" {
    case "a":
        if ((v1 * (v1 << ((1) & 63))) < (v1 / ((v1) | 1))) {
            v11 := v2[32767]
            var v12 int8 = 42
            _ = v11
            _ = v12
        }
    case "foo bar":
        if false {
            v13 := GenS1{F0: -896}
            switch "hello" {
            case "a":
                var v14 uint8 = (v1 & (v1 + v1))
                _ = v14
            case "":
                v1 = v1
            default:
                v3[573] = 2147483647
            }
            _ = v13
        }
    default:
        fmt.Println(v1)
    }
    switch "" {
    case "hello":
        var v15 bool = true
        _ = v15
    default:
        v1 = (^(v1))
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        if (42 <= 1) {
            var v17 int8 = 42
            _ = v17
        } else {
            switch (("\n\t" + "hello") + "hello") {
            case "foo bar":
                for fi18 := 0; fi18 < 20; fi18++ {
                    v1 -= 128
                    _ = fi18
                }
            case "hello":
                v19 := v2[(fi16 + fi16)]
                _ = v19
            default:
                v1 -= v1
            }
            v20 := v3[2]
            _ = v20
        }
        _ = fi16
    }
    v21 := GenS1{F0: 1}
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS1{F0: -2147483648}
        _ = fi22
        _ = v23
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (127 <= 1) {
        v24 := GenS1{F0: 0}
        switch (("" + "a") + "hello") {
        case "foo bar":
            var v25 uint16 = 32767
            _ = v25
        case "a":
            if len(v4) > 0 {
                v4[0] = 340
            }
        case "":
            var v26 int = 9223372036854775807
            _ = v26
        default:
            v1 = v1
        }
        _ = v24
    }
    var v27 int32 = -1
    v27 -= 214
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
