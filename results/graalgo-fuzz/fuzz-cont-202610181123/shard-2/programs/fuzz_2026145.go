package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool; F1 bool; F2 int32 }
func genFn0(p0 uint32, p1 int16) int16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 -= 42
        v6 := GenS1{F0: 1, F1: 0}
        _ = fi5
        _ = v6
    }
    if false {
        var v7 int = 42
        _ = v7
    }
    switch "\n\t" {
    case "":
        if false {
            switch ("" + "a") {
            case "hello":
                v8 := []int{128, 1, -504, -420}
                _ = v8
            case "":
                fmt.Println(p1)
            default:
                v9 := GenS2{F0: 42}
                _ = v9
            }
            fmt.Println(p0, p1)
        }
    default:
        if (true || (0 >= 2147483647)) {
            p0 += p0
            fmt.Println(p0, p0)
        }
    }
    v10 := GenS1{F0: 1, F1: 255}
    v11 := []int{-31, -2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p1
}

func main() {
    var v1 int64 = -506
    v2 := []int{127}
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v12 uint32 = 4294967295
    var v13 uint8 = 128
    fmt.Println(v12)
    v14 := v4[("hello" + "")]
    var v15 uint16 = 65535
    if (false && (int16(v13) == -492)) {
        var v16 float64 = (((3.14159 - 0.1) / (3.14159 / float64(1e10))) * (0.0 * (((-2.25 * (0.1 * 0.0)) / 3.14159) * (1e10 + 0.1))))
        fmt.Println(v12, v12, v14)
        _ = v16
    } else {
        if true {
            var v17 uint8 = (v13 >> ((950) & 63))
            v15 -= 135
            _ = v17
        }
        v18 := GenS2{F0: 271}
        _ = v18
    }
    var v19 uint = 65535
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: 529, F1: v15}
        var v22 float32 = -2.25
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v13, v14, v15)
        case "a":
            if true {
                v14 = (v14 << (((v19 / ((v19) | 1))) & 63))
                var v23 string = "foo bar"
                _ = v23
            }
        case "hello":
            if (205 != (v15 >> ((65535) & 63))) {
                v1 = v1
                v19 -= v19
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := GenS2{F0: 207}
                    _ = fi24
                    _ = v25
                }
            }
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    if false {
        var v26 uint16 = v15
        _ = v26
    } else {
        switch ("foo bar" + ((("hello" + "\n\t") + "foo bar") + ("foo bar" + ((("\n\t" + "\n\t") + "hello") + "hello")))) {
        case "":
            v3[(v12 / ((0) | 1))] = ("hello" + "")
        case "hello":
            for fi27 := 0; fi27 < 10; fi27++ {
                v13 = 152
                _ = fi27
            }
        case "a":
            switch "foo bar" {
            case "foo bar":
                switch "" {
                case "hello":
                    switch -2147483648 {
                    case 42:
                        if ((v15 | 0) == uint16(v19)) {
                            v15 = 1
                            fmt.Println(v1, v13)
                            switch "a" {
                            case "\n\t":
                                var v28 float64 = (0.1 * 0.1)
                                _ = v28
                            default:
                                v15 = (^(v15))
                            }
                        } else {
                            var v29 int32 = int32(v19)
                            v30 := GenS2{F0: 177}
                            _ = v29
                            _ = v30
                        }
                    case 236:
                        fmt.Println(v1, v14)
                    case 128:
                        fmt.Println(v1, v14, v13)
                    }
                case "foo bar":
                    var v31 uint32 = (^(v12))
                    _ = v31
                case "a":
                    fmt.Println(v12, v13, v19)
                default:
                    var v32 uint = 0
                    _ = v32
                }
            case "a":
                v12 += (v14 | (v14 / (((v14 | v12)) | 1)))
            case "hello":
                var v33 uint = (v19 | 636)
                _ = v33
            }
        }
        if false {
            var v34 uint16 = (v15 + 0)
            v13 = (v13 >> (((^(v19))) & 63))
            _ = v34
        } else {
            v15 = uint16(v13)
            for fi35 := 0; fi35 < 9; fi35++ {
                fmt.Println(v14)
                v13 = uint8(v14)
                _ = fi35
            }
            v36 := GenS1{F0: -2147483648, F1: v15}
            _ = v36
        }
    }
    v14 = v14
    if false {
        if (((!((1e10 <= 3.14159) || (!false))) && ((42 != -1) && (int(v13) != 0))) || (!((v19 % ((700) | 1)) >= v19))) {
            if true {
                v37 := GenS1{F0: 42, F1: v15}
                _ = v37
            }
        }
        switch "a" {
        case "\n\t":
            var v38 int16 = 42
            _ = v38
        }
    }
    switch 32767 {
    case 958:
        switch "" {
        case "hello":
            var v39 int8 = 127
            _ = v39
        case "foo bar":
            if false {
                if len(v2) > 0 {
                    v2[0] = 1
                }
            } else {
                var v40 uint32 = v12
                _ = v40
            }
        }
    case -1:
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 uint8 = (v13 % ((0) | 1))
            if len(v2) > 0 {
                v2[0] = (fi41 * (fi41 / ((fi41) | 1)))
            }
            _ = fi41
            _ = v42
        }
    case 32768:
        v12 -= v14
    }
    v43 := GenS3{F0: (v1 | 1), F1: v19}
    switch 1 {
    case -2147483648:
        v44 := GenS4{F0: v14, F1: 771, F2: 0}
        _ = v44
    default:
        var v45 uint32 = (v14 << (((v19 >> ((v19) & 63))) & 63))
        _ = v45
    }
    v15 += (^(v15))
    v2 = append(v2, 32767)
    v46 := GenS5{F0: true, F1: true, F2: 0}
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v46
}
