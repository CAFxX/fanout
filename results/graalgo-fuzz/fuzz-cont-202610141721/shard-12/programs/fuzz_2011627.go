package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 bool }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    switch ((("hello" + "\n\t") + ("a" + "a")) + ("hello" + ((("hello" + "foo bar") + "a") + ""))) {
    case "\n\t":
        v4 := []int{-270, 2147483647, 32768}
        _ = v4
    case "a":
        fmt.Println(p0)
    case "":
        fmt.Println(p0)
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                switch "a" {
                case "\n\t":
                    var v7 uint64 = (p0 >> ((4294967295) & 63))
                    _ = v7
                default:
                    switch fi6 {
                    case -9223372036854775808:
                        var v8 uint64 = 18446744073709551615
                        _ = v8
                    default:
                        var v9 int8 = 104
                        _ = v9
                    }
                }
                var v10 int64 = -416
                var v11 uint8 = 127
                _ = fi6
                _ = v10
                _ = v11
            }
            v12 := []int{(fi5 >> ((2) & 63))}
            switch "foo bar" {
            case "foo bar":
                p0 -= p0
            default:
                p0 += (p0 % ((p0) | 1))
            }
            _ = fi5
            _ = v12
        }
    }
    v13 := []int{0, -50}
    p0 = (p0 / ((p0) | 1))
    var v14 uint = 217
    _ = p0
    _ = v13
    _ = v14
    return 127
}

func main() {
    v1 := []int{-146, 304, -851}
    var v2 int = -835
    var v3 int8 = 0
    switch "a" {
    case "a":
        switch 1 {
        case -1:
            if true {
                var v15 uint32 = 313
                var v16 uint64 = 278
                _ = v15
                _ = v16
            }
        case -9223372036854775808:
            if ((-117 > 787) || true) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    switch v2 {
                    case -80:
                        fmt.Println(v3)
                    case 456:
                        _ = len(v1)
                    default:
                        var v17 float32 = 3.14159
                        _ = v17
                    }
                default:
                    fmt.Println(v2, v2)
                }
                v18 := GenS1{F0: 18446744073709551615}
                var v19 float32 = -2.25
                _ = v18
                _ = v19
            }
        }
    case "hello":
        v3 -= (-(v3))
    case "foo bar":
        fmt.Println(v2)
    }
    if (0 > 65535) {
        v2 = v2
        fmt.Println(v2, v3)
    }
    fmt.Println(v2)
    v3 -= -75
    v2 -= v2
    switch v2 {
    case 32767:
        var v20 uint = 0
        _ = v20
    case 0:
        switch "" {
        case "\n\t":
            v3 -= 100
        default:
            v2 = (v2 - 32767)
        }
    case 145:
        fmt.Println(v3, v3)
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        if (true && ((-1 <= 1) && (float64(1e10) <= 0.1))) {
            switch (fi21 >> ((1) & 63)) {
            case -2147483648:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 128:
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(v3)
                    var v23 int8 = v3
                    v24 := GenS1{F0: 194}
                    v3 += -128
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            if true {
                var v25 uint16 = 2
                _ = v25
            }
            v26 := GenS2{F0: 633, F1: 1, F2: true}
            _ = v26
        }
        fmt.Println(v3)
        var v27 bool = (0 <= 0)
        _ = fi21
        _ = v27
    }
    if true {
        v1 = append(v1, (v2 & 32767))
        var v28 uint16 = 255
        switch (v2 - v2) {
        case 42:
            v1 = append(v1, v2)
        case 1:
            fmt.Println(v28, v28, v3)
        }
        _ = v28
    } else {
        var v29 uint = 1
        var v30 float32 = ((float32(0.0) * ((0.1 * (1.5 - -2.25)) / 1.0)) - 1e10)
        _ = v29
        _ = v30
    }
    v31 := GenS3{F0: 0, F1: 243}
    fmt.Println(v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
