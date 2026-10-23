package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int32 = 141
    var v2 uint16 = 42
    var v3 bool = true
    var v4 uint = 1
    var v5 bool = (0 >= 1)
    var v6 float32 = 1e10
    var v7 int = 32768
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v3)
        fmt.Println(v1)
        if v5 {
            v3 = (1000 < 375)
            if false {
                v3 = false
            }
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int16 = int16(v2)
        v11 := GenS1{F0: v4, F1: 1, F2: (v4 ^ (v4 / (((v4 << (((^(v4))) & 63))) | 1)))}
        var v12 uint = v4
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    switch ((("a" + "hello") + "foo bar") + "\n\t") {
    case "hello":
        v3 = false
    }
    var v13 int8 = -1
    var v14 uint64 = 216
    var v15 uint8 = 0
    if (42 > 65535) {
        v4 = v4
        for fi16 := 0; fi16 < 5; fi16++ {
            v15 = 127
            v2 -= (^(v2))
            _ = fi16
        }
    } else {
        fmt.Println(v3, v5)
    }
    if true {
        v17 := []int{}
        fmt.Println(v13)
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if ((!("a" > ((("\n\t" + "foo bar") + "hello") + "a"))) && (v3 && v5)) {
            v2 += 329
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS2{F0: int16(v14)}
                _ = fi19
                _ = v20
            }
            fmt.Println(v1, v4)
        }
        v14 += (v14 * uint64(v4))
        _ = fi18
    }
    switch v7 {
    case -16:
        v21 := []int{}
        _ = v21
    case -318:
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := []int{v7, 2147483647, (fi22 << ((v4) & 63))}
            for fi24 := 0; fi24 < 6; fi24++ {
                fmt.Println(v2)
                var v25 uint16 = uint16(v4)
                v23 = append(v23, 0)
                _ = fi24
                _ = v25
            }
            var v26 uint8 = (v15 << (((v4 + v4)) & 63))
            _ = fi22
            _ = v23
            _ = v26
        }
    case 9223372036854775807:
        v2 -= v2
    }
    switch int(v15) {
    case 628:
        v4 = (v4 >> ((v4) & 63))
    case 42:
        v27 := []int{(v7 + 42)}
        _ = v27
    default:
        for fi28 := 0; fi28 < 3; fi28++ {
            if (!(0 > v2)) {
                if false {
                    v29 := []int{v7, v7, v7, (v7 >> (((v4 | (v4 & 4294967295))) & 63))}
                    if ((v13 % ((102) | 1)) == v13) {
                        if (110 != v15) {
                            var v30 uint64 = v14
                            switch "" {
                            case "a":
                                if (!(!false)) {
                                    if (((-2.25 / v6) + v6) != ((0.1 - v6) * v6)) {
                                        switch "a" {
                                        case "":
                                            var v31 float32 = (1e10 + (-2.25 + v6))
                                            _ = v31
                                        case "hello":
                                            v5 = (v5 || v3)
                                        case "\n\t":
                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                if len(v29) > 0 {
                                                    v29[0] = 146
                                                }
                                                fmt.Println(v5, fi32, v6)
                                                v2 -= v2
                                                _ = fi32
                                            }
                                        default:
                                            if (v3 && v3) {
                                                v1 = 914
                                                var v33 uint8 = 1
                                                if len(v29) >= 2 {
                                                    v29 = v29[1:len(v29)]
                                                }
                                                _ = v33
                                            } else {
                                                fmt.Println(v30, fi28, v30)
                                            }
                                        }
                                        v4 += (v4 | (v4 - 0))
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v34 uint8 = (v15 >> (((^(v4))) & 63))
                                            _ = v34
                                        case "\n\t":
                                            switch "hello" {
                                            case "":
                                                v35 := GenS2{F0: 32767}
                                                _ = v35
                                            case "\n\t":
                                                var v36 uint64 = (^(v30))
                                                _ = v36
                                            case "foo bar":
                                                var v37 float64 = -2.25
                                                _ = v37
                                            default:
                                                if true {
                                                    _ = len(v29)
                                                    fmt.Println(v3, v30, v3)
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v2, v1)
                                }
                            case "foo bar":
                                if false {
                                    fmt.Println(v2, v13)
                                    v13 -= (v13 | (v13 >> (((v4 & (v4 | (v4 << ((1) & 63))))) & 63)))
                                    if (!v5) {
                                        v15 -= (v15 & v15)
                                        v4 = (v4 / ((v4) | 1))
                                        var v38 string = ("a" + ("hello" + "a"))
                                        _ = v38
                                    } else {
                                        for fi39 := 0; fi39 < 3; fi39++ {
                                            var v40 uint16 = (v2 + v2)
                                            fmt.Println(v5, v5, v2)
                                            switch v7 {
                                            case 32768:
                                                v13 = v13
                                            case -215:
                                                if v5 {
                                                    fmt.Println(v14, fi28)
                                                }
                                            case 1:
                                                v30 = (v30 & 4294967295)
                                            default:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v30, v3, v15)
                                                }
                                                fmt.Println(v40, v14, v7)
                                                v13 -= v13
                                            }
                                            _ = fi39
                                            _ = v40
                                        }
                                    }
                                }
                            }
                            v14 += (^(v30))
                            _ = v30
                        }
                        var v41 int32 = v1
                        if v5 {
                            for fi42 := 0; fi42 < 16; fi42++ {
                                switch ("hello" + "a") {
                                case "\n\t":
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                case "hello":
                                    var v43 uint8 = 127
                                    _ = v43
                                case "a":
                                    fmt.Println(v41, v6)
                                }
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                _ = fi42
                            }
                            for fi44 := 0; fi44 < 19; fi44++ {
                                v13 -= v13
                                v3 = v3
                                if (v2 > (v2 ^ (^(v2)))) {
                                    if ((v2 | (v2 << ((769) & 63))) > (v2 & (v2 / (((^(v2))) | 1)))) {
                                        var v45 uint16 = (v2 / ((v2) | 1))
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        _ = v45
                                    } else {
                                        v46 := GenS1{F0: uint(v14), F1: 0, F2: (v4 << (((v4 - v4)) & 63))}
                                        var v47 int32 = v41
                                        _ = v46
                                        _ = v47
                                    }
                                    var v48 uint = v4
                                    _ = v48
                                }
                                _ = fi44
                            }
                            v5 = ((v6 + 0.0) >= 0.0)
                        }
                        _ = v41
                    } else {
                        fmt.Println(v13, v13, v5)
                    }
                    _ = v29
                }
                v2 = (v2 - (v2 % ((v2) | 1)))
            }
            _ = fi28
        }
    }
    v49 := []int{}
    for fi50 := 0; fi50 < 8; fi50++ {
        v51 := GenS3{F0: (^(v14))}
        _ = fi50
        _ = v51
    }
    v4 = v4
    if (int(v13) <= (v7 ^ int(v1))) {
        v49 = append(v49, (v7 + v7))
        if len(v49) >= 2 {
            v49 = v49[1:len(v49)]
        }
    } else {
        if (v13 == (v13 - v13)) {
            _ = len(v49)
            if len(v49) > 0 {
                v49[0] = v7
            }
        } else {
            v52 := GenS4{F0: (float64(float32(((0.1 / (0.0 + ((0.0 - 3.14159) - 3.14159))) * float64(1e10)))) < (1e10 / 1.5))}
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
