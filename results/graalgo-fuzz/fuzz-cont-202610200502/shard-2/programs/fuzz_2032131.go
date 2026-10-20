package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 int32, p1 bool) uint8 {
    if (803 <= 1000) {
        v4 := GenS1{F0: 4294967295}
        var v5 uint64 = 42
        if p1 {
            v6 := []int{-9223372036854775808}
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
            _ = v6
        }
        _ = v4
        _ = v5
    } else {
        v7 := []int{32768, 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 42}
            _ = fi9
            _ = v10
        }
        var v11 uint16 = 32767
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: 1}
    v13 := GenS1{F0: 1}
    v14 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 int8 = 6
    var v2 int16 = 42
    v3 := make(map[string]int)
    if ((!false) && true) {
        if ((v1 << ((4294967295) & 63)) < v1) {
            switch "" {
            case "a":
                delete(v3, "a")
            case "hello":
                delete(v3, "foo bar")
            case "foo bar":
                switch -2147483648 {
                case 2147483647:
                    var v15 int8 = v1
                    _ = v15
                default:
                    if ((82 > 1) || (!false)) {
                        v2 -= (v2 % ((-32768) | 1))
                        switch "foo bar" {
                        case "a":
                            v1 += (v1 ^ v1)
                        }
                        switch "foo bar" {
                        case "":
                            v16 := GenS2{F0: 18446744073709551615, F1: (-(v2))}
                            _ = v16
                        case "hello":
                            fmt.Println(v1, v2)
                        case "a":
                            if ((18446744073709551615 != 420) && true) {
                                v1 += int8(v2)
                            } else {
                                var v17 int64 = 0
                                var v18 uint16 = 2
                                var v19 int16 = v2
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        }
                    } else {
                        var v20 uint64 = 954
                        _ = v20
                    }
                }
            }
        }
    } else {
        if false {
            v2 = (v2 % (((v2 * (-(v2)))) | 1))
            for fi21 := 0; fi21 < 3; fi21++ {
                switch (fi21 / (((fi21 + (fi21 / (((fi21 >> ((4294967295) & 63))) | 1)))) | 1)) {
                case 128:
                    v2 = v2
                case 72:
                    var v22 bool = true
                    _ = v22
                }
                v23 := GenS3{F0: 2147483647, F1: 662, F2: 32767}
                v24 := GenS1{F0: 337}
                _ = fi21
                _ = v23
                _ = v24
            }
        }
        v25 := GenS4{F0: 9223372036854775807, F1: 42}
        v1 = (v1 - 127)
        _ = v25
    }
    if (1.5 > 1e10) {
        var v26 float64 = ((1.5 - float64(-2.25)) + ((1e10 * ((1.5 * 3.14159) * 3.14159)) / 1.0))
        fmt.Println(v26, v26, v2)
        _ = v26
    }
    v27 := []int{839, 32768}
    switch -9223372036854775808 {
    case 1:
        v28 := v3["foo bar"]
        _ = v28
    case -1:
        var v29 bool = (!false)
        _ = v29
    default:
        fmt.Println(v1, v1)
    }
    v30 := GenS2{F0: 4294967295, F1: 1}
    v2 -= v2
    if (66 < 127) {
        if false {
            var v31 string = "\n\t"
            v1 = (v1 & (v1 + v1))
            _ = v31
        }
        v32 := GenS3{F0: -1, F1: -1, F2: 42}
        _ = v32
    } else {
        var v33 int = 42
        v33 = (v33 / ((v33) | 1))
        _ = v33
    }
    v2 += int16(v1)
    v2 = (v2 * (-(v2)))
    var v34 float32 = 1e10
    if true {
        var v35 float64 = (0.0 * (0.0 / 1.5))
        _ = v35
    }
    var v36 int8 = (v1 | v1)
    v37 := GenS2{F0: 65535, F1: (v2 ^ v2)}
    switch 395 {
    case 357:
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 int64 = 0
            switch (fi38 * fi38) {
            case 634:
                v27 = append(v27, 22)
            case 9223372036854775807:
                if len(v27) > 0 {
                    v27[0] = int(v39)
                }
            }
            _ = fi38
            _ = v39
        }
    default:
        switch 399 {
        case 128:
            v40 := v3["a"]
            _ = v40
        default:
            for fi41 := 0; fi41 < 6; fi41++ {
                if true {
                    v36 = v36
                }
                switch (fi41 / ((fi41) | 1)) {
                case 128:
                    for fi42 := 0; fi42 < 14; fi42++ {
                        v2 = 0
                        _ = fi42
                    }
                default:
                    delete(v3, "")
                }
                _ = fi41
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v36
    _ = v37
}
