package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16) int8 {
    v6 := GenS1{F0: -9, F1: 42}
    switch "a" {
    case "":
        switch 128 {
        case 32767:
            var v7 int8 = -1
            _ = v7
        case -125:
            fmt.Println(p0)
        default:
            v8 := GenS2{F0: (-32768 == 337), F1: 0, F2: 1}
            _ = v8
        }
    case "hello":
        var v9 bool = true
        _ = v9
    case "foo bar":
        if true {
            if false {
                switch "a" {
                case "":
                    var v10 uint16 = p0
                    _ = v10
                default:
                    fmt.Println(p0)
                }
            }
        }
    default:
        v11 := []int{32768}
        _ = v11
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1() bool {
    var v12 int64 = 639
    v12 += -9223372036854775808
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v12)
            var v15 uint = 962
            v16 := GenS1{F0: 100, F1: 619}
            v17 := GenS1{F0: 0, F1: v15}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    _ = v12
    return true
}

func main() {
    var v1 float32 = 1.5
    var v2 int16 = 0
    var v3 string = "\n\t"
    var v4 float32 = float32((-2.25 - (((-2.25 + 1e10) - -2.25) + 0.0)))
    v5 := []int{-335}
    v18 := GenS3{F0: 797, F1: 255}
    var v19 int = 4294967295
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v20 float64 = ((0.1 - float64((v4 + ((0.1 * float32(1e10)) * v4)))) - (-2.25 * (0.0 / 1e10)))
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float32 = ((-2.25 / 1.5) * float32(v20))
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: 86, F1: 1}
    for fi24 := 0; fi24 < 19; fi24++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 += 3.14159
        var v25 float64 = v20
        _ = fi24
        _ = v25
    }
    var v26 float64 = float64(0.1)
    var v27 int = 42
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 uint32 = 171
        v1 -= (3.14159 + (1.5 / (v4 * 0.1)))
        var v30 uint32 = v29
        _ = fi28
        _ = v29
        _ = v30
    }
    if true {
        var v31 uint64 = 0
        v31 = (v31 / ((359) | 1))
        _ = v31
    }
    fmt.Println(v26, v20)
    if false {
        for fi32 := 0; fi32 < 17; fi32++ {
            if false {
                v19 = int(v2)
            }
            if len(v5) > 0 {
                v5[0] = (-(fi32))
            }
            _ = fi32
        }
        var v33 uint32 = 50
        _ = v33
    }
    var v34 int = (v27 >> ((2) & 63))
    v34 = (v34 + 9223372036854775807)
    switch v3 {
    case "":
        if (!(127 > 255)) {
            _ = len(v5)
        } else {
            switch v19 {
            case 32767:
                v35 := GenS4{F0: genFn0(281)}
                _ = v35
            }
            v20 = v26
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 uint = 18446744073709551615
                _ = fi36
                _ = v37
            }
        }
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = (v34 | (^(v27)))
        }
    case "hello":
        var v38 uint = 2
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        fmt.Println(v27, v26, v26)
        v40 := GenS1{F0: 127, F1: 42}
        _ = fi39
        _ = v40
    }
    switch (v19 << ((65535) & 63)) {
    case 42:
        var v41 int16 = (v2 - (v2 % (((v2 / ((1) | 1))) | 1)))
        _ = v41
    case 9223372036854775807:
        switch v27 {
        case -967:
            var v42 string = ""
            _ = v42
        case 4294967295:
            switch -976 {
            case -196:
                var v43 string = (v3 + v3)
                _ = v43
            default:
                fmt.Println(v3)
            }
        case 32768:
            for fi44 := 0; fi44 < 11; fi44++ {
                v45 := GenS2{F0: (0 == 9223372036854775807), F1: 255, F2: (v2 ^ (-(v2)))}
                var v46 uint64 = 1
                for fi47 := 0; fi47 < 11; fi47++ {
                    if false {
                        switch fi44 {
                        case -2147483648:
                            if len(v5) > 0 {
                                v5[0] = v27
                            }
                        case 1:
                            v48 := GenS4{F0: 44}
                            _ = v48
                        }
                    } else {
                        if true {
                            var v49 string = v3
                            var v50 uint8 = 83
                            _ = v49
                            _ = v50
                        }
                    }
                    v34 -= (fi47 - -1)
                    _ = fi47
                }
                _ = fi44
                _ = v45
                _ = v46
            }
        }
    default:
        var v51 uint = 0
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
