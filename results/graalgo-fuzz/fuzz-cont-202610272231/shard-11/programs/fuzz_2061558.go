package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    var v1 int32 = 1
    var v2 int16 = -750
    var v3 int16 = v2
    var v4 int8 = 106
    switch -2147483648 {
    case 127:
        for fi5 := 0; fi5 < 1; fi5++ {
            switch "foo bar" {
            case "a":
                var v6 uint = 4294967295
                _ = v6
            }
            _ = fi5
        }
    default:
        fmt.Println(v4)
    }
    v7 := []int{127}
    v3 += (^(v2))
    for fi8 := 0; fi8 < 15; fi8++ {
        v1 -= v1
        switch fi8 {
        case 539:
            v1 = 1
        case 4294967295:
            var v9 int64 = 1
            _ = v9
        default:
            for fi10 := 0; fi10 < 12; fi10++ {
                if len(v7) > 0 {
                    v7[0] = 32767
                }
                v11 := GenS1{F0: (v1 - -303), F1: 1}
                v1 -= v1
                if true {
                    v2 -= (v2 << ((0) & 63))
                    if (4294967295 < 42) {
                        var v12 float32 = 3.14159
                        var v13 int32 = -71
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: 53, F1: 256, F2: 2}
                    _ = v14
                }
                v15 := GenS3{F0: (false || (42 != 42)), F1: 2147483647}
                _ = fi10
                _ = v11
                _ = v15
            }
        }
        _ = fi8
    }
    v16 := GenS4{F0: 464}
    _ = len(v7)
    switch "hello" {
    case "\n\t":
        v4 -= 100
    case "hello":
        _ = len(v7)
    case "a":
        if true {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "a" {
                case "":
                    fmt.Println(v1, v1)
                case "hello":
                    v7 = append(v7, (fi17 | (fi17 ^ fi17)))
                default:
                    switch "a" {
                    case "a":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v3 = (v3 & v3)
                            fmt.Println(fi18, fi17)
                            _ = fi18
                        }
                    case "hello":
                        fmt.Println(v3, v2)
                    default:
                        var v19 int16 = (^(v3))
                        _ = v19
                    }
                }
                _ = fi17
            }
        }
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS5{F0: (fi20 / (((fi20 ^ fi20)) | 1))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v22 int16 = v2
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v23 := GenS6{F0: 0}
    var v24 float32 = (3.14159 - ((1e10 / (0.1 + (((((((0.0 + ((0.1 - (1e10 / 0.1)) / (3.14159 - (-2.25 + (3.14159 * -2.25))))) + 1e10) + -2.25) - 0.1) * -2.25) + 0.1) + ((((((1.5 + float32((-2.25 / 1e10))) - float32(1e10)) / 3.14159) + 1.5) * 0.1) - (0.0 / 1.0))))) + 1.5))
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if false {
                v27 := GenS3{F0: true, F1: 2147483648}
                v28 := GenS6{F0: 65535}
                _ = len(v7)
                _ = v27
                _ = v28
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v29 int16 = v2
                _ = len(v7)
                if (!(true && (true || ((^(fi26)) >= (fi25 | fi25))))) {
                    v30 := GenS4{F0: 920}
                    _ = v30
                } else {
                    switch -610 {
                    case 4294967295:
                        fmt.Println(v3, v29)
                        v24 = v24
                    }
                    v31 := GenS7{F0: 1, F1: 1}
                    _ = v31
                }
                if false {
                    if true {
                        var v32 int64 = -9223372036854775808
                        _ = v32
                    }
                    fmt.Println(v1, v29)
                }
                _ = v29
            }
            _ = fi26
        }
        fmt.Println(v24, v4)
        _ = fi25
    }
    fmt.Println(v3)
    for fi33 := 0; fi33 < 2; fi33++ {
        _ = len(v7)
        if false {
            var v34 int16 = (v3 | v2)
            _ = v34
        } else {
            if (true && true) {
                for fi35 := 0; fi35 < 9; fi35++ {
                    if true {
                        v1 += v1
                        var v36 int8 = v4
                        _ = v36
                    }
                    if true {
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                var v37 float32 = ((0.0 + v24) / 3.14159)
                                _ = v37
                            case "a":
                                fmt.Println(fi35, v3)
                            }
                            v7 = append(v7, 918)
                            v4 = v4
                        }
                    } else {
                        var v38 int16 = v3
                        var v39 float32 = (v24 + (3.14159 + v24))
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                }
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS5{F0: (fi33 / ((4294967295) | 1))}
                    v7 = append(v7, fi40)
                    var v42 uint64 = 18446744073709551615
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        switch ("foo bar" + "a") {
        case "\n\t":
            v43 := GenS8{F0: (v3 / ((v3) | 1)), F1: v1, F2: 255}
            _ = v43
        case "hello":
            v3 -= (v3 | v2)
        case "":
            switch (fi33 | fi33) {
            case -103:
                var v44 uint8 = 42
                _ = v44
            case -695:
                var v45 uint64 = 1
                _ = v45
            }
        }
        _ = fi33
    }
    if true {
        switch 42 {
        case 42:
            var v46 int8 = v4
            _ = v46
        }
    } else {
        var v47 uint32 = 49
        v2 = 0
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
