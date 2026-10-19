package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 uint }
func genFn0() uint64 {
    var v7 uint = 535
    for fi8 := 0; fi8 < 10; fi8++ {
        v7 = v7
        for fi9 := 0; fi9 < 4; fi9++ {
            v7 = 42
            v10 := GenS1{F0: 425, F1: int(v7), F2: 65535}
            _ = fi9
            _ = v10
        }
        var v11 int = (-(fi8))
        _ = fi8
        _ = v11
    }
    v12 := []int{1, 42}
    v7 += v7
    _ = v12
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 65535
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 4294967295
    var v5 int8 = -1
    v6 := make(map[int]uint32)
    v4 = (v4 - genFn0())
    if (0 > 0) {
        delete(v6, 2147483647)
    }
    v13 := []int{1, 127}
    if ((^(v2)) > v2) {
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                v5 = 127
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 uint64 = v4
                    switch "a" {
                    case "a":
                        fmt.Println(v2, fi14)
                    case "hello":
                        var v17 uint16 = 255
                        _ = v17
                    case "foo bar":
                        if len(v13) > 0 {
                            v13[0] = fi15
                        }
                    default:
                        if ((^(v16)) == (v4 >> ((v2) & 63))) {
                            var v18 string = ("" + ("hello" + ("hello" + "hello")))
                            if len(v13) > 0 {
                                v13[0] = -2147483648
                            }
                            fmt.Println(fi14)
                            _ = v18
                        } else {
                            v19 := GenS1{F0: 65535, F1: fi15, F2: 449}
                            var v20 int8 = int8(fi14)
                            var v21 string = "hello"
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                    switch "\n\t" {
                    case "":
                        var v22 float64 = -2.25
                        _ = v22
                    case "\n\t":
                        v23 := v6[fi15]
                        fmt.Println(fi14, v5, v16)
                        _ = v23
                    case "a":
                        v24 := GenS2{F0: v2, F1: (v5 - -114), F2: 9223372036854775807}
                        _ = v24
                    default:
                        switch 128 {
                        case 32767:
                            v2 += v2
                        default:
                            v16 = v4
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                var v25 int16 = -32768
                _ = v25
            }
            v26 := GenS1{F0: uint16(v2), F1: (fi14 * (fi14 + (^(fi14)))), F2: 0}
            _ = fi14
            _ = v26
        }
        var v27 float32 = 1.5
        _ = v27
    } else {
        v28 := v6[-585]
        var v29 bool = false
        _ = v28
        _ = v29
    }
    v4 -= (v4 ^ v4)
    var v30 float32 = 0.0
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS2{F0: v2, F1: 90, F2: (^(v3))}
        _ = fi31
        _ = v32
    }
    v3 += (v3 - 0)
    fmt.Println(v30, v4)
    if ((v5 ^ (v5 ^ 0)) == (v5 * v5)) {
        if false {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            var v33 int32 = 0
            _ = v33
        }
        var v34 string = ""
        _ = v34
    } else {
        switch 940 {
        case 128:
            var v35 uint32 = 0
            _ = v35
        default:
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
    }
    v2 = 0
    switch ((("\n\t" + ("" + ("hello" + "a"))) + "hello") + "") {
    case "foo bar":
        fmt.Println(v3, v5)
    default:
        switch "a" {
        case "foo bar":
            if true {
                fmt.Println(v2)
                switch "hello" {
                case "a":
                    if ((("" + "a") + ("" + ("" + "foo bar"))) > "\n\t") {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 float64 = float64(v30)
                            v13 = append(v13, fi36)
                            v38 := GenS3{F0: v2, F1: 372, F2: "hello"}
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        v30 = (((float32(0.0) * v30) - (v30 * v30)) + (0.0 - (((((0.0 - 0.0) / (v30 / (v30 * 1.5))) * ((v30 / (v30 + (v30 * ((1.5 / (1.5 + v30)) - float32(-2.25))))) - 3.14159)) / v30) + (v30 * ((0.1 * v30) - ((1e10 - v30) + (v30 / (3.14159 + v30))))))))
                    } else {
                        var v39 int32 = -2147483648
                        var v40 uint16 = 65535
                        _ = v39
                        _ = v40
                    }
                case "foo bar":
                    if (42 <= -101) {
                        v2 += (v2 & v2)
                        v41 := GenS4{F0: 938}
                        _ = v41
                    }
                }
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                var v42 int = 1
                for fi43 := 0; fi43 < 14; fi43++ {
                    v5 = (v5 >> (((^(v2))) & 63))
                    v4 = 704
                    _ = fi43
                }
                var v44 uint16 = 0
                _ = v42
                _ = v44
            }
        case "hello":
            var v45 float64 = (1e10 - 0.1)
            _ = v45
        case "a":
            v5 -= (v5 ^ (v5 | (v5 / (((^(v5))) | 1))))
        default:
            switch 32768 {
            case -9223372036854775808:
                v46 := GenS1{F0: 126, F1: 42, F2: 460}
                _ = v46
            case 32768:
                v2 += v2
            default:
                var v47 float64 = 3.14159
                _ = v47
            }
        }
    }
    _ = len(v1)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
