package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 string) uint8 {
    var v5 uint16 = 667
    var v6 bool = ((42 == 42) || ((0.0 <= ((((3.14159 + 0.0) / 1e10) / (1.5 * (3.14159 - 1e10))) - ((3.14159 * float64(((-2.25 * float32(0.1)) * 0.0))) * 0.1))) && true))
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(p0, v5)
        switch ((p0 + (p0 + p0)) + "\n\t") {
        case "":
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint16 = (v5 ^ v5)
                if (!true) {
                    v10 -= (v5 ^ 255)
                    fmt.Println(v5)
                } else {
                    var v11 uint64 = (^(v7))
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        case "foo bar":
            p0 = ("hello" + ((((p0 + p0) + "") + ("\n\t" + (p0 + (p0 + "a")))) + (("" + ((p0 + p0) + (p0 + ""))) + p0)))
        case "\n\t":
            v12 := []int{(fi8 ^ (fi8 + fi8))}
            _ = v12
        default:
            v13 := []int{(fi8 * (fi8 ^ (fi8 / ((fi8) | 1)))), -9223372036854775808}
            _ = v13
        }
        switch fi8 {
        case 42:
            v14 := GenS1{F0: (v7 & (v7 >> ((0) & 63))), F1: 1}
            _ = v14
        case 9223372036854775807:
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 float64 = 1.5
                var v17 uint16 = uint16(v7)
                _ = fi15
                _ = v16
                _ = v17
            }
        case -1:
            v6 = ((v5 != 32767) && (((3.14159 / (float32(1e10) + 0.0)) / 3.14159) != 0.0))
        default:
            v5 -= v5
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 1
    var v2 float32 = ((float32((float64(0.1) * (float64(0.1) - 0.1))) + 0.0) - 1.5)
    var v3 uint16 = 0
    var v4 uint32 = 421
    v18 := []int{42, 0}
    v19 := GenS2{F0: 1, F1: 0, F2: v4}
    v4 -= 0
    v20 := GenS3{F0: 489, F1: 95}
    switch 128 {
    case 127:
        for fi21 := 0; fi21 < 19; fi21++ {
            v2 += ((v2 - v2) * (((((v2 - 3.14159) - float32((float64(v2) / 3.14159))) * v2) - v2) + v2))
            _ = fi21
        }
    default:
        var v22 int32 = 728
        _ = v22
    }
    v3 = (v3 | v3)
    v3 -= v3
    if (0.1 <= (-2.25 / 1.0)) {
        fmt.Println(v1, v2, v2)
        for fi23 := 0; fi23 < 4; fi23++ {
            switch "\n\t" {
            case "":
                v18 = append(v18, fi23)
            case "\n\t":
                if true {
                    _ = len(v18)
                    v2 = v2
                }
            case "hello":
                for fi24 := 0; fi24 < 2; fi24++ {
                    v1 = v1
                    _ = fi24
                }
            }
            if (2 > 486) {
                switch "foo bar" {
                case "\n\t":
                    var v25 bool = (!("\n\t" == "foo bar"))
                    _ = v25
                default:
                    v2 = ((v2 - v2) / v2)
                }
            }
            _ = fi23
        }
        switch "a" {
        case "\n\t":
            if false {
                v4 = (v4 - (v4 >> ((65535) & 63)))
            }
        default:
            for fi26 := 0; fi26 < 13; fi26++ {
                v1 = (v1 ^ 209)
                _ = fi26
            }
        }
    } else {
        v27 := GenS1{F0: 4294967295, F1: 463}
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if false {
            switch (((("a" + "foo bar") + (("\n\t" + "foo bar") + ("" + ""))) + "hello") + (("\n\t" + "hello") + "\n\t")) {
            case "a":
                v29 := GenS2{F0: 2, F1: 4294967295, F2: v4}
                _ = v29
            }
            fmt.Println(v4)
            v30 := GenS4{F0: (v4 / (((^(v4))) | 1)), F1: 9223372036854775807, F2: 127}
            _ = v30
        } else {
            v2 = v2
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
