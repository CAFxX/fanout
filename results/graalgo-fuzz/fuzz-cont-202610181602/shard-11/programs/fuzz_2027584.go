package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int8 = 121
    var v2 uint32 = 723
    var v3 uint16 = 65535
    var v4 float64 = (-2.25 * 1.5)
    var v5 int32 = 2147483647
    fmt.Println(v1, v5)
    v3 = v3
    v6 := GenS1{F0: 2147483647}
    fmt.Println(v4, v2)
    if true {
        v5 = 680
        switch "foo bar" {
        case "foo bar":
            v7 := []int{42, 193, 4294967295, 127}
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: (v5 - v5)}
            _ = v8
        default:
            v5 = v5
        }
    } else {
        var v9 int = 127
        v3 += (v3 ^ (v3 - v3))
        fmt.Println(v2, v9, v2)
        _ = v9
    }
    v10 := []int{32768, -1, 32767}
    if false {
        switch "a" {
        case "\n\t":
            if false {
                for fi11 := 0; fi11 < 8; fi11++ {
                    switch fi11 {
                    case 4294967295:
                        v12 := GenS1{F0: (v5 * 469)}
                        _ = v12
                    case 32767:
                        v4 += v4
                    case 510:
                        v2 += uint32(v3)
                    }
                    v4 += v4
                    _ = fi11
                }
                for fi13 := 0; fi13 < 5; fi13++ {
                    v1 -= (v1 / ((v1) | 1))
                    if true {
                        fmt.Println(v3, fi13)
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        fmt.Println(v2, fi13)
                        var v15 int16 = 0
                        v1 = v1
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
            }
        case "a":
            var v16 float64 = (v4 + 0.1)
            _ = v16
        case "hello":
            if (((("foo bar" + "\n\t") + "\n\t") + "a") != "foo bar") {
                if (true || false) {
                    if (255 == 95) {
                        v3 = 2
                    }
                    v10 = append(v10, -2147483648)
                } else {
                    fmt.Println(v1, v4, v3)
                }
                switch 591 {
                case 417:
                    var v17 uint8 = uint8(v2)
                    _ = v17
                case -93:
                    v18 := GenS1{F0: (v5 << ((257) & 63))}
                    _ = v18
                case 391:
                    if ((1.5 * v4) != v4) {
                        if (672 >= 275) {
                            for fi19 := 0; fi19 < 5; fi19++ {
                                var v20 bool = (65535 > (v2 & (^(v2))))
                                v21 := GenS1{F0: v5}
                                v5 = v5
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(v3)
                            fmt.Println(v3)
                        }
                        if (true || ((3.14159 / 1.0) == float32((1.5 / v4)))) {
                            var v22 float32 = -2.25
                            var v23 uint64 = 42
                            v24 := GenS1{F0: v5}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        v3 = 364
                    }
                }
            } else {
                var v25 int32 = (v5 << ((255) & 63))
                v26 := GenS1{F0: -2147483648}
                _ = v25
                _ = v26
            }
        }
        v27 := GenS2{F0: (v2 | (v2 | 809))}
        for fi28 := 0; fi28 < 13; fi28++ {
            v4 += v4
            _ = fi28
        }
        _ = v27
    }
    var v29 float32 = 0.0
    v1 -= (v1 + (v1 % ((v1) | 1)))
    v2 = 0
    fmt.Println(v3)
    if true {
        v30 := GenS2{F0: (v2 >> ((0) & 63))}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 13; fi31++ {
            switch 4294967295 {
            case -9223372036854775808:
                switch ("a" + "hello") {
                case "foo bar":
                    var v32 float64 = ((1e10 - 1.5) * v4)
                    _ = v32
                }
            case 843:
                v33 := GenS3{F0: fi31, F1: 0}
                _ = v33
            case 9223372036854775807:
                v2 += v2
            default:
                switch "" {
                case "a":
                    var v34 uint32 = 0
                    _ = v34
                case "\n\t":
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
            }
            var v35 float64 = v4
            _ = fi31
            _ = v35
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 int32 = (v5 ^ (v5 << ((42) & 63)))
            v38 := GenS1{F0: v37}
            if true {
                if true {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                switch 400 {
                case 165:
                    _ = len(v10)
                }
                v4 += float64(((v29 / ((float32(v4) - v29) - (v29 - v29))) + (float32((float64(float32(3.14159)) / 0.1)) / float32(((1e10 * ((float64(v29) + (v4 * (((v4 * float64(1.5)) * (v4 + 1e10)) * (3.14159 + float64(v29))))) / v4)) * v4)))))
            }
            _ = fi36
            _ = v37
            _ = v38
        }
        fmt.Println(v3, v2)
    }
    if (((1 >= 2147483647) || (-40 == (v1 | (^(v1))))) && (16 < 239)) {
        switch 32767 {
        case 9223372036854775807:
            if (-2.25 <= (float32((3.14159 / (v4 * (v4 / (((((v4 * float64((float32(v4) * (((v29 * v29) - v29) / 0.1)))) * v4) + float64(float32(((-2.25 - (v4 + float64(v29))) + 3.14159)))) + -2.25) / v4))))) + float32((1e10 + ((0.1 + v4) + 1.5))))) {
                var v39 bool = (1 >= 0)
                if (0 == 128) {
                    v29 += v29
                    var v40 uint64 = 4294967295
                    if ((3.14159 > (v29 - v29)) && (9223372036854775807 > 0)) {
                        _ = len(v10)
                    } else {
                        v41 := GenS4{F0: -1}
                        _ = v41
                    }
                    _ = v40
                }
                _ = v39
            } else {
                if (!true) {
                    var v42 float32 = (v29 * -2.25)
                    if (v3 == v3) {
                        fmt.Println(v4, v5, v3)
                    }
                    var v43 int8 = int8(v3)
                    _ = v42
                    _ = v43
                }
                if (v4 <= v4) {
                    v4 += float64(-2.25)
                    var v44 uint = 1
                    _ = v44
                }
                v5 -= v5
            }
        case 32767:
            fmt.Println(v1, v1)
        case 2147483647:
            if false {
                var v45 uint64 = 42
                v45 = (v45 & (v45 >> ((42) & 63)))
                for fi46 := 0; fi46 < 7; fi46++ {
                    v47 := GenS3{F0: (fi46 >> ((2) & 63)), F1: -1}
                    fmt.Println(fi46)
                    var v48 uint64 = (v45 >> ((uint(v45)) & 63))
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = v45
            }
        }
        for fi49 := 0; fi49 < 18; fi49++ {
            _ = len(v10)
            if (0 >= 2147483648) {
                if (((!false) && false) || true) {
                    var v50 uint8 = 255
                    _ = v50
                } else {
                    if false {
                        var v51 string = ""
                        for fi52 := 0; fi52 < 18; fi52++ {
                            v53 := GenS1{F0: (v5 + int32(v2))}
                            _ = fi52
                            _ = v53
                        }
                        v3 = v3
                        _ = v51
                    }
                    var v54 int64 = -1
                    if len(v10) > 0 {
                        v10[0] = 817
                    }
                    _ = v54
                }
            } else {
                fmt.Println(v5, v29)
            }
            fmt.Println(v3)
            _ = fi49
        }
    }
    for fi55 := 0; fi55 < 15; fi55++ {
        var v56 uint64 = 4294967295
        v5 += v5
        _ = fi55
        _ = v56
    }
    switch ("hello" + "") {
    case "a":
        v1 += v1
    case "foo bar":
        v10 = append(v10, -1)
    case "":
        if true {
            switch -563 {
            case 2147483647:
                v10 = append(v10, -9223372036854775808)
            }
        } else {
            for fi57 := 0; fi57 < 18; fi57++ {
                var v58 int16 = 0
                _ = fi57
                _ = v58
            }
            v2 += uint32(v3)
            var v59 bool = (42 == 197)
            _ = v59
        }
    default:
        fmt.Println(v4, v4)
    }
    v3 = (v3 + (^(v3)))
    v60 := GenS3{F0: 723, F1: (v5 << ((255) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
