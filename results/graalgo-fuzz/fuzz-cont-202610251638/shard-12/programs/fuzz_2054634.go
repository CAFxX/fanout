package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 uint16) int64 {
    v4 := GenS1{F0: 13}
    v5 := GenS1{F0: -128}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1() int64 {
    var v6 int8 = 100
    v6 = v6
    var v7 string = "hello"
    v7 = v7
    if (42 < v6) {
        if true {
            v6 -= v6
            v7 = "\n\t"
        }
        v8 := GenS1{F0: v6}
        v9 := []int{-79}
        _ = v8
        _ = v9
    } else {
        v6 -= (-(v6))
        v10 := GenS2{F0: 0, F1: 1}
        _ = v10
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 0
    var v2 uint8 = 196
    var v3 int16 = int16(v2)
    v11 := GenS3{F0: 885}
    switch "" {
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int32 = 1
            fmt.Println(v3, v13, v13)
            _ = fi12
            _ = v13
        }
    case "foo bar":
        switch 128 {
        case 9223372036854775807:
            fmt.Println(v3)
        default:
            v1 += v3
        }
    case "hello":
        v14 := []int{}
        _ = v14
    }
    v3 -= (v3 << ((65535) & 63))
    var v15 float64 = float64(1e10)
    var v16 int32 = 736
    var v17 string = "a"
    for fi18 := 0; fi18 < 8; fi18++ {
        if (1 > 4294967295) {
            switch "hello" {
            case "a":
                v16 += v16
            case "\n\t":
                v3 = (v1 + (-(v3)))
            case "":
                if (!false) {
                    v19 := []int{fi18, -1, 27}
                    _ = v19
                }
            default:
                var v20 int = -209
                _ = v20
            }
            fmt.Println(v3, v1, v16)
        }
        _ = fi18
    }
    switch -9223372036854775808 {
    case 479:
        for fi21 := 0; fi21 < 2; fi21++ {
            if false {
                v1 -= v3
                var v22 int16 = v1
                switch "a" {
                case "foo bar":
                    switch ((v17 + v17) + (v17 + (("a" + ("foo bar" + v17)) + ("" + (v17 + v17))))) {
                    case "":
                        var v23 int16 = v1
                        _ = v23
                    case "\n\t":
                        switch "" {
                        case "a":
                            for fi24 := 0; fi24 < 17; fi24++ {
                                fmt.Println(fi24, v17)
                                v25 := GenS2{F0: 65535, F1: 255}
                                _ = fi24
                                _ = v25
                            }
                        case "":
                            fmt.Println(v3, v2)
                        default:
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v16 = v16
                                switch (fi26 & 32767) {
                                case 2147483647:
                                    v27 := []int{(fi21 % ((int(v1)) | 1)), fi21}
                                    _ = v27
                                case -2147483648:
                                    var v28 uint16 = 2
                                    _ = v28
                                default:
                                    v29 := GenS3{F0: fi21}
                                    _ = v29
                                }
                                switch (fi26 / ((fi26) | 1)) {
                                case 32768:
                                    v30 := GenS2{F0: 65535, F1: 913}
                                    _ = v30
                                case 9223372036854775807:
                                    var v31 uint = 507
                                    _ = v31
                                case 575:
                                    v2 += 128
                                    v15 -= 0.1
                                }
                                _ = fi26
                            }
                        }
                    }
                case "\n\t":
                    fmt.Println(v22)
                case "a":
                    v3 += v3
                }
                _ = v22
            }
            _ = fi21
        }
    }
    var v32 float32 = (((float32((v15 * (v15 * (v15 + 1e10)))) - 1.5) - 0.1) / (float32(1.5) + 0.0))
    var v33 uint64 = 1
    v3 -= (^(v1))
    var v34 uint32 = 65535
    v35 := []int{128, 2147483647}
    if (0 >= 1) {
        for fi36 := 0; fi36 < 5; fi36++ {
            _ = len(v35)
            fmt.Println(v33, v17, v17)
            v35 = append(v35, int(v34))
            _ = fi36
        }
        if (false || false) {
            v37 := GenS1{F0: 127}
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        switch (^(fi38)) {
        case 42:
            v32 = v32
        case 9223372036854775807:
            switch v17 {
            case "a":
                v39 := GenS4{F0: v17}
                _ = v39
            default:
                v2 -= v2
            }
        default:
            if false {
                for fi40 := 0; fi40 < 9; fi40++ {
                    if (0.0 < v32) {
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        fmt.Println(v32, v34, fi40)
                        v32 -= -2.25
                        v41 := GenS3{F0: fi38}
                        var v42 int32 = 0
                        var v43 int32 = v16
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi40
                }
            } else {
                v15 += v15
                v44 := GenS2{F0: 483, F1: uint16(v2)}
                _ = v44
            }
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            v16 = 0
            _ = fi45
        }
        _ = fi38
    }
    if ("\n\t" != v17) {
        for fi46 := 0; fi46 < 14; fi46++ {
            v34 += 523
            switch "foo bar" {
            case "foo bar":
                if (42 <= uint16(v33)) {
                    switch (v17 + (v17 + v17)) {
                    case "":
                        v1 += (v3 * (v3 ^ (v3 << ((2) & 63))))
                    }
                    v47 := GenS2{F0: 137, F1: 255}
                    _ = v47
                }
            case "a":
                v32 = 1.5
            case "\n\t":
                var v48 string = v17
                _ = v48
            default:
                v49 := GenS1{F0: 100}
                _ = v49
            }
            _ = fi46
        }
    }
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
