package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int32, p1 int32, p2 uint16) uint8 {
    var v5 uint8 = 1
    v6 := GenS1{F0: ((((("foo bar" + "foo bar") + "hello") + "") + ("a" + "foo bar")) + "hello"), F1: ((p2 < p2) || (226 <= uint64(v5))), F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(v5))
}
func genFn1() int {
    var v7 uint64 = 0
    v8 := GenS2{F0: 435}
    _ = v7
    _ = v8
    return -560
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 2
    v3 := []uint32{}
    var v4 float32 = (((0.0 / 0.1) - 3.14159) / 1.5)
    var v9 uint = 0
    for fi10 := 0; fi10 < 2; fi10++ {
        if len(v3) > 0 {
            v3[0] = 174
        }
        var v11 uint = v1
        _ = fi10
        _ = v11
    }
    v4 = float32(float64((v4 + 3.14159)))
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v9, v9, fi12)
        var v13 uint = v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
        _ = v13
    }
    switch "foo bar" {
    case "":
        if false {
            switch 127 {
            case 0:
                fmt.Println(v1)
            case -790:
                var v14 uint32 = 2147483648
                _ = v14
            }
            fmt.Println(v9, v4, v9)
        }
    case "\n\t":
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint8 = v2
            v17 := GenS2{F0: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
    case "foo bar":
        var v18 uint8 = v2
        _ = v18
    default:
        v4 = v4
    }
    if ((v9 >> (((v1 - v1)) & 63)) == (v9 & v9)) {
        var v19 int16 = 1000
        v2 += v2
        switch 127 {
        case 9223372036854775807:
            var v20 uint32 = 241
            _ = v20
        default:
            v21 := GenS3{F0: 42, F1: 2}
            _ = v21
        }
        _ = v19
    } else {
        if true {
            var v22 uint32 = 2
            if len(v3) > 0 {
                v3[0] = 1
            }
            v23 := GenS1{F0: "hello", F1: false, F2: 837}
            _ = v22
            _ = v23
        } else {
            var v24 int16 = -32768
            var v25 uint32 = 42
            v26 := GenS2{F0: 1}
            _ = v24
            _ = v25
            _ = v26
        }
    }
    if (true && true) {
        _ = len(v3)
        if (!false) {
            if len(v3) > 0 {
                v3[0] = 161
            }
            var v27 int16 = 0
            _ = v27
        } else {
            switch "foo bar" {
            case "hello":
                for fi28 := 0; fi28 < 9; fi28++ {
                    v9 = 4294967295
                    if false {
                        _ = len(v3)
                        v4 = v4
                    } else {
                        v1 = 65535
                        v4 = v4
                    }
                    _ = fi28
                }
            case "":
                var v29 uint64 = 2
                _ = v29
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v9 += v9
            }
        }
    } else {
        switch "a" {
        case "a":
            switch -9223372036854775808 {
            case 9223372036854775807:
                v30 := GenS2{F0: 18446744073709551615}
                _ = v30
            case -1:
                switch -1 {
                case 127:
                    fmt.Println(v9)
                case 9223372036854775807:
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v2 = v2
                        fmt.Println(v4, v1)
                        _ = fi31
                    }
                case 128:
                    if false {
                        v32 := GenS1{F0: ("" + "foo bar"), F1: true, F2: 42}
                        fmt.Println(v4, v4, v1)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 11; fi33++ {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                fmt.Println(v1, v4)
                                _ = fi34
                            }
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            for fi35 := 0; fi35 < 15; fi35++ {
                                v1 += 311
                                _ = fi35
                            }
                            _ = fi33
                        }
                        v2 -= genFn0(-883, int32(v1), 42)
                    }
                default:
                    for fi36 := 0; fi36 < 5; fi36++ {
                        _ = len(v3)
                        var v37 int8 = 42
                        _ = fi36
                        _ = v37
                    }
                }
            case 2147483647:
                for fi38 := 0; fi38 < 14; fi38++ {
                    v39 := GenS2{F0: 162}
                    v2 += 128
                    _ = fi38
                    _ = v39
                }
            default:
                v4 = (((1e10 + v4) * v4) * ((v4 + float32(3.14159)) - v4))
            }
        case "foo bar":
            v2 += v2
        default:
            var v40 string = "a"
            _ = v40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            for fi42 := 0; fi42 < 18; fi42++ {
                v43 := GenS3{F0: 0, F1: 177}
                var v44 string = "hello"
                var v45 string = v44
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 2; fi46++ {
                fmt.Println(v9, v1)
                if (527 == 47) {
                    var v47 int32 = 1
                    if len(v3) > 0 {
                        v3[0] = 966
                    }
                    _ = v47
                }
                var v48 float64 = 0.1
                fmt.Println(v48, v1)
                switch "" {
                case "foo bar":
                    switch "foo bar" {
                    case "a":
                        if true {
                            _ = len(v3)
                            v49 := GenS4{F0: 255, F1: 309}
                            _ = v49
                        }
                    default:
                        v50 := GenS5{F0: (fi46 % ((32767) | 1)), F1: true}
                        _ = v50
                    }
                }
                if true {
                    fmt.Println(v4, fi41)
                    switch "" {
                    case "a":
                        _ = len(v3)
                    case "\n\t":
                        var v51 string = ("\n\t" + "a")
                        _ = v51
                    case "":
                        v52 := GenS6{F0: 42}
                        _ = v52
                    default:
                        v53 := GenS1{F0: (("" + ("a" + "a")) + "a"), F1: true, F2: 18446744073709551615}
                        _ = v53
                    }
                }
                v4 -= (0.0 * v4)
                _ = fi46
                _ = v48
            }
            _ = fi41
        }
    }
    var v54 bool = false
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    for fi55 := 0; fi55 < 17; fi55++ {
        if len(v3) > 0 {
            v3[0] = 113
        }
        if (!("a" != (("" + "hello") + "foo bar"))) {
            if len(v3) > 0 {
                v3[0] = 35
            }
            var v56 uint8 = v2
            switch "hello" {
            case "hello":
                v54 = (2147483647 != 2147483647)
            default:
                if v54 {
                    switch fi55 {
                    case -1:
                        v9 = v1
                    case -272:
                        for fi57 := 0; fi57 < 16; fi57++ {
                            v56 = v56
                            var v58 float64 = 3.14159
                            _ = fi57
                            _ = v58
                        }
                    }
                    v59 := GenS7{F0: -1}
                    _ = v59
                }
            }
            _ = v56
        }
        _ = fi55
    }
    var v60 uint32 = 782
    v54 = (9 == (v2 >> ((v9) & 63)))
    fmt.Println(v9)
    var v61 int8 = 86
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v60
    _ = v61
    _ = v9
}
