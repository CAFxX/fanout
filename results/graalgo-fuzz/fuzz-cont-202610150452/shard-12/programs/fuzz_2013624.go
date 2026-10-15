package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS4 struct { F0 int; F1 int8; F2 uint64 }
func genFn0() int32 {
    if (true || true) {
        var v4 float32 = (float32(1.5) / 0.1)
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 705, F1: (673 <= 1), F2: 127}
            fmt.Println(v4)
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS2{F0: -32768, F1: 128, F2: (^(fi5))}
                var v9 uint64 = 2
                v9 += (v9 / ((484) | 1))
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi5
            _ = v6
        }
        v4 -= 1e10
        _ = v4
    } else {
        var v10 float64 = 0.0
        var v11 uint8 = 132
        _ = v10
        _ = v11
    }
    var v12 uint = 330
    v12 = 58
    var v13 uint16 = 809
    _ = v12
    _ = v13
    return 2147483647
}
func genFn1(p0 uint64) int16 {
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{(fi14 >> ((305) & 63)), fi14}
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 0
    if true {
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(v16, v16)
            _ = fi17
        }
        var v18 int = -558
        if ("foo bar" == "hello") {
            var v19 uint32 = 2
            _ = v19
        } else {
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := []int{fi20, (fi20 << ((65535) & 63))}
                if false {
                    v21 = append(v21, (v18 << ((42) & 63)))
                    fmt.Println(p0, v16)
                } else {
                    v22 := GenS2{F0: -475, F1: v18, F2: (fi20 << ((1) & 63))}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        }
        _ = v18
    }
    _ = p0
    _ = v16
    return 1
}

func main() {
    var v1 uint = 802
    var v2 int8 = 0
    var v3 bool = (true && (true && (!false)))
    v2 -= v2
    var v23 uint = 2
    v23 -= 65535
    if (v3 || v3) {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := []int{}
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 uint8 = 0
                if (2 <= 42) {
                    var v28 uint64 = 0
                    var v29 int = (fi24 + int(v28))
                    if false {
                        fmt.Println(v2, v27, fi26)
                        var v30 int8 = 127
                        _ = v30
                    } else {
                        var v31 float32 = 3.14159
                        v32 := GenS3{F0: (v28 >> (((v23 + uint(v28))) & 63)), F1: 2, F2: -9223372036854775808}
                        _ = v31
                        _ = v32
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    v2 -= 127
    switch 127 {
    case 132:
        switch "a" {
        case "\n\t":
            var v33 int16 = 42
            _ = v33
        case "":
            var v34 int8 = v2
            _ = v34
        default:
            var v35 uint16 = 714
            _ = v35
        }
    case 32767:
        fmt.Println(v23, v2, v1)
    default:
        v2 = -128
    }
    var v36 uint64 = 4294967295
    v2 = (-(v2))
    if false {
        v37 := []int{-2147483648, 9223372036854775807}
        if v3 {
            v38 := GenS3{F0: (v36 << (((v23 - (v23 & (v1 ^ 4294967295)))) & 63)), F1: 0, F2: 1}
            var v39 bool = v3
            if (9223372036854775807 > 128) {
                var v40 int = 2147483647
                v41 := GenS2{F0: 317, F1: (^(v40)), F2: -2147483648}
                _ = v40
                _ = v41
            }
            _ = v38
            _ = v39
        } else {
            var v42 uint32 = 370
            _ = v42
        }
        switch "" {
        case "\n\t":
            switch "a" {
            case "hello":
                if v3 {
                    if (v3 && (!v3)) {
                        var v43 uint32 = 0
                        v36 += 4294967295
                        v43 = v43
                        _ = v43
                    }
                    if (1e10 >= 3.14159) {
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v1 -= v23
                            if false {
                                v2 += -11
                            }
                            _ = fi44
                        }
                    } else {
                        v2 = v2
                        v2 = 96
                    }
                }
            default:
                switch "a" {
                case "foo bar":
                    var v45 float32 = 1e10
                    _ = v45
                }
            }
        case "foo bar":
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
        }
        _ = v37
    }
    if (1 == 42) {
        for fi46 := 0; fi46 < 13; fi46++ {
            for fi47 := 0; fi47 < 1; fi47++ {
                v1 = 4294967295
                switch "hello" {
                case "hello":
                    v23 = 1
                }
                var v48 float64 = 3.14159
                _ = fi47
                _ = v48
            }
            _ = fi46
        }
    }
    v49 := GenS4{F0: -1, F1: 1, F2: v36}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v36, v3, v3)
    case "hello":
        v2 -= v2
    default:
        v50 := []int{}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v49
}
