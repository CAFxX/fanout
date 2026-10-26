package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint8 = 128
    var v2 float64 = 0.1
    var v3 uint8 = v1
    var v4 int = -9223372036854775808
    var v5 float64 = 1e10
    if false {
        var v6 uint16 = uint16(v3)
        v3 -= (v3 & v3)
        if false {
            var v7 int64 = 1
            v8 := []int{42}
            _ = v7
            _ = v8
        }
        _ = v6
    } else {
        v5 -= (v5 / v5)
        v9 := []int{v4, 32768, (-(v4)), (v4 & v4)}
        switch v4 {
        case 32767:
            fmt.Println(v5)
        }
        _ = v9
    }
    v10 := []int{v4, 805, (v4 - (v4 + 4294967295))}
    fmt.Println(v5, v2)
    v11 := GenS1{F0: (true || false), F1: 434, F2: 2}
    if false {
        v5 = (1.5 + 1.5)
        if (782 > 32767) {
            for fi12 := 0; fi12 < 1; fi12++ {
                switch "hello" {
                case "hello":
                    fmt.Println(v5)
                default:
                    switch "a" {
                    case "a":
                        var v13 uint8 = 1
                        _ = v13
                    default:
                        if false {
                            fmt.Println(v1, v3)
                            if true {
                                if len(v10) > 0 {
                                    v10[0] = 829
                                }
                                for fi14 := 0; fi14 < 7; fi14++ {
                                    if ((-27 != 1) || (true || true)) {
                                        var v15 string = "foo bar"
                                        v1 += 128
                                        _ = v15
                                    } else {
                                        v3 = v1
                                    }
                                    _ = fi14
                                }
                            }
                            v16 := GenS1{F0: true, F1: 255, F2: 1}
                            _ = v16
                        }
                    }
                }
                if false {
                    v17 := GenS2{F0: -801}
                    _ = v17
                } else {
                    fmt.Println(v4)
                    v1 -= v3
                    v18 := GenS2{F0: -9223372036854775808}
                    _ = v18
                }
                var v19 int32 = -390
                _ = fi12
                _ = v19
            }
            switch "foo bar" {
            case "\n\t":
                var v20 uint64 = 810
                _ = v20
            }
        } else {
            v1 = 42
            var v21 int = v4
            switch (v21 + v4) {
            case 42:
                v22 := GenS2{F0: 9223372036854775807}
                _ = v22
            case -629:
                fmt.Println(v21, v4, v2)
            }
            _ = v21
        }
    } else {
        var v23 int8 = 0
        _ = v23
    }
    switch "\n\t" {
    case "a":
        switch ("\n\t" + "a") {
        case "":
            fmt.Println(v3)
        default:
            fmt.Println(v2, v5, v2)
        }
    case "hello":
        v2 += v5
    default:
        var v24 string = "a"
        _ = v24
    }
    var v25 int64 = 0
    var v26 uint32 = 701
    fmt.Println(v1, v3)
    var v27 uint16 = 1
    for fi28 := 0; fi28 < 2; fi28++ {
        switch v4 {
        case -803:
            if false {
                if true {
                    if ((v2 / (v5 + (((v5 / v2) + v5) / 1.0))) == v2) {
                        switch fi28 {
                        case -171:
                            v29 := GenS2{F0: (v25 % ((v25) | 1))}
                            _ = v29
                        case 128:
                            var v30 int = fi28
                            _ = v30
                        }
                        if false {
                            fmt.Println(v2, v3, v5)
                            for fi31 := 0; fi31 < 9; fi31++ {
                                var v32 float64 = (v2 - (-2.25 / (float64((1e10 * 0.0)) * (((0.0 + v5) - 0.0) - 3.14159))))
                                _ = fi31
                                _ = v32
                            }
                        }
                        v3 += v3
                    } else {
                        var v33 uint8 = (^(v3))
                        if (1 > int8(v26)) {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            var v34 float64 = float64(-2.25)
                            for fi35 := 0; fi35 < 1; fi35++ {
                                switch (v4 % (((fi28 & v4)) | 1)) {
                                case 4294967295:
                                    v5 = v34
                                case 127:
                                    fmt.Println(fi28, v3)
                                case 35:
                                    switch ("" + "foo bar") {
                                    case "a":
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                }
                                var v36 float32 = 0.0
                                _ = fi35
                                _ = v36
                            }
                            _ = v34
                        } else {
                            switch "hello" {
                            case "a":
                                for fi37 := 0; fi37 < 17; fi37++ {
                                    switch "" {
                                    case "foo bar":
                                        var v38 bool = false
                                        _ = v38
                                    case "\n\t":
                                        var v39 uint64 = 18446744073709551615
                                        _ = v39
                                    default:
                                        if false {
                                            var v40 uint = 42
                                            _ = v40
                                        } else {
                                            v33 -= v33
                                            v5 = 3.14159
                                            v1 = (^(v33))
                                        }
                                    }
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi37, v25)
                                    case "":
                                        v41 := GenS2{F0: v25}
                                        _ = v41
                                    }
                                    v42 := GenS1{F0: ((fi37 > v4) && (2 == 290)), F1: (v27 | v27), F2: 559}
                                    _ = fi37
                                    _ = v42
                                }
                            default:
                                if (true && false) {
                                    var v43 int32 = 0
                                    _ = v43
                                }
                            }
                        }
                        fmt.Println(v2, v25)
                        _ = v33
                    }
                    if (true || (!true)) {
                        var v44 uint16 = 32767
                        var v45 uint16 = 882
                        _ = v44
                        _ = v45
                    }
                }
                v3 = 1
            }
        case 9223372036854775807:
            v46 := GenS1{F0: (((((float64(1.5) - (-2.25 / v5)) - v2) - (1.5 / (((v2 * ((v2 / v5) * 1.5)) / v2) - v5))) * float64((-2.25 + 0.0))) > v5), F1: (v27 ^ (v27 + (v27 & 32767))), F2: 18446744073709551615}
            _ = v46
        case 4294967295:
            v2 -= v2
        default:
            v26 += v26
        }
        fmt.Println(v2, v1, v4)
        v4 = (fi28 - (-(fi28)))
        _ = fi28
    }
    var v47 uint64 = 42
    if false {
        v10 = append(v10, (v4 << ((72) & 63)))
        v48 := GenS3{F0: (0 > 1000)}
        _ = v48
    }
    v49 := GenS2{F0: v25}
    for fi50 := 0; fi50 < 13; fi50++ {
        switch "\n\t" {
        case "foo bar":
            for fi51 := 0; fi51 < 2; fi51++ {
                if false {
                    var v52 string = "foo bar"
                    v25 -= -438
                    _ = v52
                }
                switch ("\n\t" + "a") {
                case "hello":
                    if (false && false) {
                        v26 += v26
                    } else {
                        v53 := GenS2{F0: 1}
                        var v54 uint = 0
                        _ = v53
                        _ = v54
                    }
                case "a":
                    v55 := GenS2{F0: (-(v25))}
                    _ = v55
                case "foo bar":
                    switch "a" {
                    case "":
                        var v56 int16 = -32768
                        var v57 uint = 65535
                        _ = v56
                        _ = v57
                    case "\n\t":
                        if true {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            switch "hello" {
                            case "\n\t":
                                var v58 int = (-(fi50))
                                v1 = v3
                                _ = v58
                            case "":
                                v59 := GenS2{F0: (v25 - (v25 % ((int64(v4)) | 1)))}
                                _ = v59
                            case "hello":
                                v47 -= v47
                            }
                            v60 := GenS2{F0: (v25 | (v25 - v25))}
                            v26 = (v26 & v26)
                            _ = v60
                        } else {
                            v61 := GenS3{F0: true}
                            var v62 uint = 255
                            _ = v61
                            _ = v62
                        }
                    }
                default:
                    var v63 uint8 = 127
                    _ = v63
                }
                _ = fi51
            }
        case "\n\t":
            for fi64 := 0; fi64 < 1; fi64++ {
                v2 += (float64((-2.25 * float32(float64(((((float32(1.5) + -2.25) + (float32(float64(0.1)) + 1e10)) / 1.0) / 0.1))))) * v2)
                v65 := GenS4{F0: -1}
                fmt.Println(fi64, v2)
                _ = fi64
                _ = v65
            }
        default:
            v10 = append(v10, (^(fi50)))
        }
        _ = fi50
    }
    switch v4 {
    case 127:
        var v66 string = ("hello" + (("hello" + "hello") + ("foo bar" + "")))
        _ = v66
    }
    if (2147483647 > -1) {
        if ((false || ((!((^(v4)) >= (v4 ^ v4))) || ((!false) && true))) && (v47 >= v47)) {
            for fi67 := 0; fi67 < 14; fi67++ {
                if len(v10) > 0 {
                    v10[0] = (fi67 ^ (v4 | (^(v4))))
                }
                var v68 uint = 18446744073709551615
                var v69 int = 606
                _ = fi67
                _ = v68
                _ = v69
            }
        }
    }
    v25 += 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
}
