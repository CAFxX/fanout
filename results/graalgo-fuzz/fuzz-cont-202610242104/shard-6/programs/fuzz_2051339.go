package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 65535
    var v3 uint32 = 65535
    v4 := []uint32{v3}
    v5 := make(map[string]string)
    var v6 uint64 = 1
    v2 -= v2
    if false {
        switch "a" {
        case "\n\t":
            if true {
                v2 = uint16(v3)
                v3 = v3
                if (false || true) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                var v7 int64 = -219
                v8 := GenS1{F0: (^(v6)), F1: 1}
                _ = len(v4)
                _ = v7
                _ = v8
            }
        case "hello":
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint64 = 42
                for fi11 := 0; fi11 < 11; fi11++ {
                    if (1 == 0) {
                        var v12 uint8 = 127
                        _ = v12
                    } else {
                        if (!(9223372036854775807 >= -495)) {
                            var v13 int8 = 106
                            _ = v13
                        }
                    }
                    v6 -= (v10 % (((v10 ^ 1)) | 1))
                    v10 += v10
                    _ = fi11
                }
                _ = fi9
                _ = v10
            }
        }
        var v14 float32 = 3.14159
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int8 = 42
            var v17 uint32 = (v3 >> ((65535) & 63))
            fmt.Println(v14, v2)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    } else {
        v18 := GenS1{F0: 586, F1: 948}
        v2 -= 1
        var v19 float64 = float64(float32((3.14159 / 0.1)))
        _ = v18
        _ = v19
    }
    if ((0.0 - (1.5 + (3.14159 + (3.14159 - 0.1)))) > 3.14159) {
        v20 := GenS2{F0: 127}
        v6 -= v6
        for fi21 := 0; fi21 < 19; fi21++ {
            switch "foo bar" {
            case "a":
                v22 := GenS3{F0: 42}
                _ = v22
            case "\n\t":
                v3 = (v3 >> ((0) & 63))
            default:
                var v23 uint64 = 2
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    } else {
        var v24 float64 = 1e10
        _ = v24
    }
    if false {
        v3 = (v3 << ((uint(v6)) & 63))
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    var v25 int32 = 379
    var v26 uint8 = 188
    var v27 string = (((("" + "foo bar") + "") + ("" + "hello")) + ("\n\t" + "\n\t"))
    var v28 string = v27
    v6 = (v6 | v6)
    v27 = v27
    if (((1.5 / 0.1) + ((float32(((1.5 * 1e10) + 1e10)) / (1.5 * (-2.25 + -2.25))) * (1e10 * 3.14159))) == -2.25) {
        v6 = (v6 >> ((255) & 63))
    }
    if (float64(1.5) > 0.0) {
        switch "\n\t" {
        case "hello":
            var v29 uint64 = 845
            _ = v29
        case "":
            v6 = (^(v6))
        }
    } else {
        delete(v5, "a")
        if (v6 >= v6) {
            if (((^(v26)) == v26) && (!true)) {
                var v30 string = "\n\t"
                v31 := GenS2{F0: 42}
                v6 += v6
                _ = v30
                _ = v31
            } else {
                v4 = append(v4, 65535)
                for fi32 := 0; fi32 < 5; fi32++ {
                    if true {
                        v6 -= v6
                    } else {
                        v27 = v27
                        fmt.Println(v26, v2, v2)
                        var v33 uint8 = (v26 >> ((18446744073709551615) & 63))
                        _ = v33
                    }
                    _ = fi32
                }
            }
        } else {
            var v34 uint8 = v26
            for fi35 := 0; fi35 < 9; fi35++ {
                var v36 int = 32768
                _ = fi35
                _ = v36
            }
            _ = v34
        }
    }
    var v37 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
