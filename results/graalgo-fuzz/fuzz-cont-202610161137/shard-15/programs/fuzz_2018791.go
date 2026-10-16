package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint16 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint32 = 394
    var v2 float32 = 3.14159
    var v3 int64 = 42
    v4 := GenS1{F0: 596}
    v5 := []int{}
    v2 = (v2 - 3.14159)
    switch -2147483648 {
    case 870:
        fmt.Println(v1, v1, v2)
    default:
        v6 := GenS1{F0: 85}
        _ = v6
    }
    var v7 float64 = 3.14159
    v7 += 0.0
    if ("foo bar" != ("\n\t" + "a")) {
        switch int(v3) {
        case 127:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint32 = v1
                _ = fi8
                _ = v9
            }
        case 1:
            for fi10 := 0; fi10 < 18; fi10++ {
                v3 = -876
                switch 0 {
                case 1:
                    v11 := GenS1{F0: 255}
                    _ = v11
                default:
                    if ((!false) && ((128 != -769) && false)) {
                        fmt.Println(v1)
                    }
                }
                _ = fi10
            }
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: 32767}
                for fi14 := 0; fi14 < 9; fi14++ {
                    v2 = v2
                    v1 = v1
                    var v15 int = (^(fi14))
                    _ = fi14
                    _ = v15
                }
                v2 = v2
                _ = fi12
                _ = v13
            }
        }
        v16 := GenS1{F0: 32767}
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                v7 = v7
                _ = fi18
            }
            _ = fi17
        }
        _ = v16
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                if true {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v2, v7)
                    case "":
                        v21 := GenS2{F0: "", F1: (v3 & 194), F2: (v3 % ((v3) | 1))}
                        _ = v21
                    default:
                        if true {
                            v7 -= (v7 / (v7 * (3.14159 * v7)))
                            if false {
                                v22 := GenS3{F0: v1}
                                _ = v22
                            }
                            v5 = append(v5, fi20)
                        }
                    }
                    switch (^(fi19)) {
                    case 934:
                        v23 := GenS3{F0: 462}
                        _ = v23
                    case -1:
                        v1 = v1
                    case 1:
                        switch 32768 {
                        case -1:
                            v3 = v3
                        case -2147483648:
                            if false {
                                v5 = append(v5, fi19)
                            }
                        default:
                            _ = len(v5)
                        }
                    }
                } else {
                    if false {
                        if true {
                            fmt.Println(v7, v7, v2)
                            if true {
                                switch -1 {
                                case 66:
                                    var v24 uint32 = v1
                                    _ = v24
                                case -2147483648:
                                    v2 -= 0.0
                                }
                                var v25 int64 = (v3 & (^(v3)))
                                fmt.Println(v25, v7)
                                _ = v25
                            }
                            var v26 uint = 42
                            _ = v26
                        } else {
                            var v27 int16 = 420
                            v3 -= v3
                            _ = v27
                        }
                        var v28 uint64 = 2
                        v29 := GenS4{F0: -1, F1: v3, F2: 65535}
                        _ = v28
                        _ = v29
                    }
                    switch fi20 {
                    case -1:
                        v1 += (v1 % (((^(v1))) | 1))
                    case 32768:
                        v3 -= (v3 + int64(v1))
                    default:
                        var v30 uint64 = 4294967295
                        _ = v30
                    }
                    switch ((("hello" + ("a" + "a")) + "a") + "a") {
                    case "\n\t":
                        var v31 uint8 = 255
                        _ = v31
                    default:
                        switch (fi19 >> ((18446744073709551615) & 63)) {
                        case 4294967295:
                            switch (fi20 % ((fi20) | 1)) {
                            case 32768:
                                v5 = append(v5, fi19)
                            case 1:
                                switch (fi20 % ((fi20) | 1)) {
                                case 960:
                                    v32 := GenS4{F0: 32767, F1: (v3 >> ((4294967295) & 63)), F2: 65535}
                                    _ = v32
                                case 32767:
                                    v7 -= v7
                                }
                            case 4294967295:
                                v7 = ((v7 - 3.14159) - (float64(3.14159) / v7))
                            }
                        case 9223372036854775807:
                            fmt.Println(v7, v1, v3)
                        }
                    }
                }
                if (3.14159 > v2) {
                    v5 = append(v5, fi19)
                    v33 := GenS4{F0: 32767, F1: (v3 & (v3 >> ((981) & 63))), F2: 1}
                    var v34 uint = 2
                    _ = v33
                    _ = v34
                } else {
                    fmt.Println(v2, v2)
                }
                _ = fi20
            }
            switch "a" {
            case "hello":
                var v35 int64 = 42
                _ = v35
            }
            if (true || ((1.5 + v2) == 3.14159)) {
                for fi36 := 0; fi36 < 9; fi36++ {
                    fmt.Println(fi19)
                    _ = len(v5)
                    _ = fi36
                }
                fmt.Println(v7)
                _ = len(v5)
            }
            _ = fi19
        }
        var v37 uint8 = 42
        _ = v37
    }
    if false {
        var v38 float64 = -2.25
        var v39 int16 = -322
        _ = v38
        _ = v39
    }
    if (0 == 42) {
        var v40 int32 = 2147483647
        fmt.Println(v1, v2, v1)
        for fi41 := 0; fi41 < 11; fi41++ {
            fmt.Println(v40, v3)
            _ = fi41
        }
        _ = v40
    }
    if (float32(float64(0.0)) <= v2) {
        if true {
            v42 := GenS5{F0: v3}
            v1 += v1
            var v43 float64 = (v7 * 0.1)
            _ = v42
            _ = v43
        } else {
            _ = len(v5)
        }
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 uint8 = 0
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
