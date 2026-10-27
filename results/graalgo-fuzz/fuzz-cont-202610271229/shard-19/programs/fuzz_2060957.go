package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int8; F1 int8; F2 int }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
func genFn0() uint16 {
    var v6 float64 = (3.14159 + 0.0)
    v7 := GenS1{F0: 437}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 string = ("" + "\n\t")
    var v2 float32 = 0.0
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    var v5 float32 = v2
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{(fi9 << ((547) & 63))}
            switch (fi9 ^ -112) {
            case 32767:
                switch fi9 {
                case -2147483648:
                    fmt.Println(v1, v2, v4)
                case 968:
                    v11 := GenS1{F0: 862}
                    _ = v11
                }
            }
            _ = fi9
            _ = v10
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 string = v1
            var v14 uint32 = 282
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    v5 += (0.0 / v5)
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint64 = 250
            switch (fi15 | (^(fi15))) {
            case -2147483648:
                switch (fi15 + fi15) {
                case 42:
                    v18 := []int{(fi16 + (fi15 - fi15)), fi16}
                    _ = v18
                case 127:
                    v4 = v4
                default:
                    var v19 int32 = 2147483647
                    _ = v19
                }
            case 127:
                var v20 uint = 4294967295
                _ = v20
            default:
                v21 := []int{fi15, 9223372036854775807, 2147483647}
                _ = v21
            }
            _ = fi16
            _ = v17
        }
        var v22 uint = 65535
        var v23 float32 = float32(1e10)
        _ = fi15
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := GenS1{F0: 2}
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 bool = false
            var v28 uint16 = 0
            _ = fi26
            _ = v27
            _ = v28
        }
        if (!true) {
            if (3.14159 < 1e10) {
                v1 = v1
                v2 += float32(3.14159)
                v4 = v4
            }
            var v29 int16 = 289
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        if ((1 >= 71) && (1 != 1)) {
            v1 = "\n\t"
        }
        _ = fi30
    }
    fmt.Println(v5, v4)
    for fi31 := 0; fi31 < 20; fi31++ {
        v5 -= (v2 + v2)
        switch 9223372036854775807 {
        case 959:
            for fi32 := 0; fi32 < 19; fi32++ {
                if (fi32 >= (fi31 >> ((387) & 63))) {
                    var v33 bool = false
                    var v34 int64 = v4
                    v35 := []int{(fi32 * 127), fi31, (^(fi31)), -339}
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    switch (fi32 ^ -1) {
                    case 4294967295:
                        fmt.Println(fi31, v1, v2)
                    case 0:
                        switch (fi32 >> ((65535) & 63)) {
                        case 519:
                            switch (("a" + v1) + ((((v1 + v1) + "\n\t") + "hello") + (v1 + "foo bar"))) {
                            case "\n\t":
                                switch v1 {
                                case "a":
                                    var v36 uint16 = 290
                                    _ = v36
                                }
                            case "":
                                fmt.Println(v5)
                            case "foo bar":
                                v37 := GenS1{F0: 42}
                                _ = v37
                            }
                        case 32768:
                            fmt.Println(fi31)
                        default:
                            var v38 float32 = v5
                            _ = v38
                        }
                    case -519:
                        v1 = v1
                    default:
                        if (721 == 220) {
                            v2 += 1.5
                        }
                        v39 := GenS1{F0: 255}
                        _ = v39
                    }
                    v40 := GenS1{F0: 32767}
                    _ = v40
                }
                v41 := GenS2{F0: -9223372036854775808, F1: -128}
                _ = fi32
                _ = v41
            }
        case 478:
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 int16 = -32768
                fmt.Println(v2, v1)
                v44 := GenS3{F0: int8(fi31), F1: 42, F2: 128}
                v45 := GenS2{F0: v4, F1: -39}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        }
        _ = fi31
    }
    v46 := GenS4{F0: -32768, F1: 1, F2: -121}
    var v47 int = 9223372036854775807
    var v48 uint8 = 1
    if ((v48 >> ((18446744073709551615) & 63)) < (v48 + v48)) {
        fmt.Println(v47, v1)
        var v49 string = "foo bar"
        v50 := []int{}
        _ = v49
        _ = v50
    }
    v51 := []int{32767, -929, -735, (v47 << ((967) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v51
}
