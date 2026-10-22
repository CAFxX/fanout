package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() int32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch "a" {
        case "a":
            switch "" {
            case "hello":
                v7 := GenS1{F0: 191, F1: 2, F2: (^(fi6))}
                _ = v7
            case "\n\t":
                var v8 uint64 = 1
                _ = v8
            }
        default:
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS2{F0: 100, F1: -9223372036854775808}
                _ = fi9
                _ = v10
            }
        }
        var v11 int16 = -1
        var v12 uint64 = 1
        _ = fi6
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS3{F0: 65535, F1: 138, F2: 210}
            switch "foo bar" {
            case "foo bar":
                var v16 uint64 = 2
                _ = v16
            case "a":
                v17 := GenS2{F0: 127, F1: fi14}
                _ = v17
            case "\n\t":
                v18 := GenS2{F0: -1, F1: -1}
                _ = v18
            default:
                fmt.Println(fi14, fi14)
            }
            _ = fi14
            _ = v15
        }
        var v19 float64 = 1.5
        _ = fi13
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        switch (fi20 & fi20) {
        case -2147483648:
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint32 = 4294967295
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS3{F0: 65535, F1: 18446744073709551615, F2: 9223372036854775807}
                v25 := GenS2{F0: int8(fi23), F1: (fi23 / ((-2147483648) | 1))}
                v26 := []int{(fi23 ^ (-(fi23))), 4294967295, (^(fi23)), 1}
                switch 4294967295 {
                case -1:
                    switch 127 {
                    case -141:
                        var v27 int16 = 655
                        _ = v27
                    }
                case 128:
                    var v28 bool = (false || true)
                    _ = v28
                }
                fmt.Println(fi23, fi23)
                if (true && true) {
                    switch fi23 {
                    case 150:
                        var v29 float64 = 0.0
                        _ = v29
                    case 9223372036854775807:
                        var v30 uint64 = 581
                        v30 -= 18446744073709551615
                        v30 += v30
                        _ = v30
                    }
                    fmt.Println(fi23, fi23)
                    v31 := GenS1{F0: 225, F1: 255, F2: (fi23 / ((fi23) | 1))}
                    var v32 int32 = 897
                    var v33 uint = 65535
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    var v34 int32 = 234
                    v26 = append(v26, 769)
                    _ = v34
                }
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 bool = true
            _ = fi35
            _ = v36
        }
        _ = fi20
    }
    v37 := GenS2{F0: -1, F1: -39}
    _ = v37
    return 0
}
func genFn1() int64 {
    var v38 bool = (!((false || false) && (0 >= 100)))
    v38 = ((-1 < -282) && (1 == -315))
    switch 42 {
    case -9223372036854775808:
        v39 := GenS2{F0: 51, F1: -9223372036854775808}
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        switch (fi40 ^ fi40) {
        case 1:
            v41 := []int{}
            _ = v41
        case 9223372036854775807:
            v42 := GenS4{F0: 42, F1: 42}
            _ = v42
        default:
            v43 := GenS3{F0: 42, F1: 0, F2: -439}
            _ = v43
        }
        _ = fi40
    }
    _ = v38
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 uint32 = 2147483647
    var v3 float64 = (((0.1 - 1.5) / 1e10) - -2.25)
    var v4 int64 = 378
    var v5 float64 = float64(0.0)
    var v44 uint = 35
    for fi45 := 0; fi45 < 2; fi45++ {
        var v46 uint64 = 1
        var v47 uint = 584
        _ = fi45
        _ = v46
        _ = v47
    }
    if true {
        var v48 int32 = 0
        v3 = ((v3 * 1.5) + (v3 / v3))
        _ = v48
    } else {
        for fi49 := 0; fi49 < 1; fi49++ {
            var v50 uint32 = v2
            _ = fi49
            _ = v50
        }
    }
    v51 := GenS5{F0: (v2 + (v2 | (v2 ^ (v2 - 4294967295)))), F1: v44}
    v4 -= 1
    if (false && (v1 < (^(v1)))) {
        fmt.Println(v5)
        var v52 uint = (v44 / ((65535) | 1))
        var v53 uint = 0
        _ = v52
        _ = v53
    }
    var v54 int16 = -1
    var v55 uint16 = uint16(v2)
    v56 := GenS5{F0: (v2 + (v2 | (v2 | v2))), F1: (v44 ^ v44)}
    var v57 uint16 = v55
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
