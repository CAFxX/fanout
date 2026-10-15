package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int }
func genFn0(p0 bool, p1 bool) int16 {
    var v5 uint16 = 65535
    p1 = (-1 < -20)
    v6 := []int{128, 127}
    for fi7 := 0; fi7 < 12; fi7++ {
        if p0 {
            v6 = append(v6, fi7)
            v5 += (v5 ^ 2)
            var v8 uint32 = 2147483647
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -876
}

func main() {
    var v1 int8 = -1
    var v2 float64 = 1e10
    v3 := []uint32{44, 0}
    var v4 int = 323
    v9 := GenS1{F0: 18446744073709551615}
    var v10 int = v4
    switch "foo bar" {
    case "foo bar":
        if (0.1 >= (float32(v2) + (1e10 / 1e10))) {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 uint16 = 255
                switch ((("a" + "\n\t") + "foo bar") + "a") {
                case "a":
                    fmt.Println(v12, v4)
                }
                var v13 float64 = -2.25
                _ = fi11
                _ = v12
                _ = v13
            }
            v4 = v10
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch (v10 / ((-107) | 1)) {
        case -2147483648:
            v15 := GenS1{F0: 0}
            _ = v15
        case 32768:
            v16 := GenS1{F0: 2}
            _ = v16
        case 4294967295:
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint16 = 65535
                _ = fi17
                _ = v18
            }
        default:
            v2 = (v2 / v2)
        }
        v2 += v2
        if false {
            for fi19 := 0; fi19 < 10; fi19++ {
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v2, v10, v10)
                    case "hello":
                        if true {
                            switch "" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                            case "hello":
                                v2 = (v2 + v2)
                            default:
                                switch v10 {
                                case -327:
                                    _ = len(v3)
                                }
                            }
                            switch (("\n\t" + "hello") + "a") {
                            case "a":
                                v4 -= (v4 % (((^(v10))) | 1))
                            case "\n\t":
                                var v20 bool = true
                                _ = v20
                            default:
                                if false {
                                    v1 = 0
                                    if false {
                                        v21 := GenS1{F0: 0}
                                        _ = v21
                                    }
                                    v22 := GenS1{F0: 2}
                                    v4 = v10
                                    v1 = (v1 * (v1 / (((v1 % ((-1) | 1))) | 1)))
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: 2}
                                    fmt.Println(fi19, v1)
                                    _ = v23
                                }
                            }
                        } else {
                            var v24 float64 = 1.5
                            _ = v24
                        }
                    case "":
                        v25 := GenS2{F0: 158}
                        _ = v25
                    }
                    switch (v10 ^ (fi14 * (fi14 & fi19))) {
                    case 508:
                        fmt.Println(fi19, v4)
                    default:
                        v3 = append(v3, 2147483647)
                    }
                }
                _ = fi19
            }
        } else {
            var v26 uint = 871
            for fi27 := 0; fi27 < 14; fi27++ {
                v4 -= 247
                _ = fi27
            }
            v4 = (^(fi14))
            _ = v26
        }
        _ = fi14
    }
    if true {
        for fi28 := 0; fi28 < 17; fi28++ {
            v10 = fi28
            _ = fi28
        }
        v3 = append(v3, 0)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 := GenS1{F0: 81}
    if len(v3) > 0 {
        v3[0] = 1
    }
    v30 := GenS2{F0: 60}
    if (2147483647 > 42) {
        var v31 uint32 = 65535
        _ = len(v3)
        fmt.Println(v1, v31, v4)
        _ = v31
    } else {
        if (1 <= 42) {
            for fi32 := 0; fi32 < 11; fi32++ {
                switch v4 {
                case 244:
                    v33 := GenS1{F0: 0}
                    _ = v33
                default:
                    v34 := GenS3{F0: 1, F1: (^(v1)), F2: 18446744073709551615}
                    _ = v34
                }
                v35 := GenS4{F0: v10}
                _ = fi32
                _ = v35
            }
            if (0 >= 475) {
                var v36 uint = 18446744073709551615
                for fi37 := 0; fi37 < 1; fi37++ {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        var v39 int8 = (v1 * -1)
                        var v40 string = "hello"
                        v1 = (v39 << ((v36) & 63))
                        var v41 int8 = (-(v39))
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = fi37
                }
                v42 := GenS3{F0: 1, F1: v1, F2: 1}
                _ = v36
                _ = v42
            }
            var v43 int8 = (-(v1))
            _ = v43
        }
        v2 = 0.0
        var v44 uint8 = 127
        _ = v44
    }
    var v45 float32 = (((3.14159 + 0.0) + float32((3.14159 / v2))) - (1.5 - -2.25))
    fmt.Println(v1, v2)
    var v46 uint8 = 136
    var v47 int16 = -1
    v45 = ((v45 - 0.1) / v45)
    switch ("a" + "foo bar") {
    case "\n\t":
        v1 -= v1
    case "a":
        var v48 uint8 = (v46 << ((0) & 63))
        _ = v48
    case "":
        for fi49 := 0; fi49 < 2; fi49++ {
            v50 := GenS2{F0: (v46 - (v46 ^ 0))}
            v51 := GenS3{F0: 65535, F1: 1, F2: 797}
            v2 -= v2
            _ = fi49
            _ = v50
            _ = v51
        }
    default:
        for fi52 := 0; fi52 < 3; fi52++ {
            v53 := GenS4{F0: 9223372036854775807}
            _ = fi52
            _ = v53
        }
    }
    switch v4 {
    case -230:
        v47 = genFn0(true, true)
    case -2147483648:
        _ = len(v3)
    default:
        switch 600 {
        case 42:
            fmt.Println(v47, v47)
        case 32767:
            switch v4 {
            case 9223372036854775807:
                var v54 int32 = 42
                _ = v54
            case -2147483648:
                var v55 uint8 = 255
                _ = v55
            }
        case -9223372036854775808:
            fmt.Println(v45)
        }
    }
    if true {
        v56 := GenS2{F0: v46}
        if true {
            v10 += (^(v10))
            var v57 int16 = 1
            switch (("" + ("foo bar" + "hello")) + "\n\t") {
            case "":
                var v58 int64 = 42
                _ = v58
            default:
                var v59 float32 = ((v45 / (v45 * v45)) + -2.25)
                _ = v59
            }
            _ = v57
        } else {
            var v60 uint8 = v46
            _ = v60
        }
        _ = v56
    } else {
        switch "\n\t" {
        case "hello":
            var v61 int = 761
            _ = v61
        case "a":
            switch "foo bar" {
            case "a":
                switch "\n\t" {
                case "":
                    var v62 uint16 = 321
                    _ = v62
                default:
                    v4 += v10
                }
            default:
                if (65535 < 42) {
                    fmt.Println(v2, v2, v47)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        case "\n\t":
            if (!((^(v46)) >= 0)) {
                var v63 int32 = -1
                var v64 int32 = (^(v63))
                _ = v63
                _ = v64
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v9
}
