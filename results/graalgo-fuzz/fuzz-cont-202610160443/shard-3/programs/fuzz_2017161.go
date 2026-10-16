package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint; F1 int16; F2 int }

func main() {
    var v1 bool = false
    var v2 int = 0
    v3 := []uint32{541, 4294967295, 4294967295, 2147483648}
    if true {
        v2 = 4294967295
    }
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint = 0
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        v2 = -656
        v1 = (-2147483648 == 2147483647)
        var v7 uint8 = 5
        _ = fi6
        _ = v7
    }
    if (100 >= -3) {
        if false {
            var v8 uint8 = 86
            fmt.Println(v1)
            _ = v8
        }
        v1 = (!false)
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 float32 = 3.14159
        v2 -= (^(v2))
        v10 = float32(float64(1.5))
        _ = fi9
        _ = v10
    }
    var v11 uint = 18446744073709551615
    v12 := GenS1{F0: -1}
    fmt.Println(v1, v2, v2)
    fmt.Println(v1, v2, v11)
    if v1 {
        if (32768 < -965) {
            for fi13 := 0; fi13 < 20; fi13++ {
                for fi14 := 0; fi14 < 12; fi14++ {
                    fmt.Println(v1, v2, v11)
                    if v1 {
                        switch (fi13 ^ fi14) {
                        case 32767:
                            v1 = false
                        case -9223372036854775808:
                            var v15 int64 = -1
                            v16 := GenS1{F0: fi14}
                            var v17 int64 = (v15 / (((v15 & (v15 ^ (v15 / ((961) | 1))))) | 1))
                            _ = v15
                            _ = v16
                            _ = v17
                        case 9223372036854775807:
                            if len(v3) > 0 {
                                v3[0] = 2147483648
                            }
                        }
                        v18 := GenS1{F0: (v2 * 4294967295)}
                        switch (("hello" + "") + ("foo bar" + "")) {
                        case "hello":
                            v1 = v1
                        case "foo bar":
                            var v19 uint64 = 129
                            _ = v19
                        case "a":
                            var v20 uint16 = 42
                            _ = v20
                        }
                        _ = v18
                    } else {
                        switch (-(fi13)) {
                        case 127:
                            var v21 int8 = -3
                            switch (fi13 / ((v2) | 1)) {
                            case 9223372036854775807:
                                switch 42 {
                                case 2147483647:
                                    var v22 uint64 = 18446744073709551615
                                    _ = v22
                                default:
                                    fmt.Println(v1)
                                }
                                switch fi13 {
                                case 128:
                                    v1 = v1
                                case -2147483648:
                                    var v23 bool = (!v1)
                                    _ = v23
                                default:
                                    switch (fi14 + (v2 - (^(fi13)))) {
                                    case 42:
                                        var v24 uint32 = 979
                                        if len(v3) > 0 {
                                            v3[0] = (v24 << (((v11 ^ v11)) & 63))
                                        }
                                        var v25 uint8 = 2
                                        if ("foo bar" >= "") {
                                            v24 += 24
                                            if false {
                                                switch (fi14 - (fi14 / ((127) | 1))) {
                                                case 32768:
                                                    v25 -= 36
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483647
                                                }
                                                v26 := GenS1{F0: fi14}
                                                _ = v26
                                            }
                                            switch ("hello" + "foo bar") {
                                            case "a":
                                                v27 := GenS2{F0: v21}
                                                _ = v27
                                            case "foo bar":
                                                var v28 bool = false
                                                _ = len(v3)
                                                _ = v28
                                            default:
                                                v29 := GenS1{F0: fi14}
                                                var v30 int8 = v21
                                                _ = v29
                                                _ = v30
                                            }
                                        } else {
                                            fmt.Println(v1)
                                            v2 += v2
                                            v21 = v21
                                            var v31 float32 = 0.1
                                            var v32 int64 = 1
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v24
                                        _ = v25
                                    case 909:
                                        fmt.Println(v2)
                                    default:
                                        var v33 int64 = -9223372036854775808
                                        _ = v33
                                    }
                                }
                                v34 := GenS3{F0: v11}
                                _ = v34
                            default:
                                if false {
                                    v3 = append(v3, 42)
                                    switch fi14 {
                                    case 128:
                                        v21 -= v21
                                        fmt.Println(v1, v11)
                                        var v35 uint8 = 128
                                        _ = v35
                                    case 9223372036854775807:
                                        switch v2 {
                                        case -1:
                                            fmt.Println(fi13)
                                        case 42:
                                            var v36 int64 = 0
                                            _ = v36
                                        case 32768:
                                            var v37 uint64 = 659
                                            _ = v37
                                        }
                                    case -2147483648:
                                        v11 -= 147
                                        v21 -= (v21 % ((v21) | 1))
                                        v38 := GenS4{F0: (^(v11)), F1: "a"}
                                        _ = v38
                                    default:
                                        v2 -= (v2 >> ((v11) & 63))
                                    }
                                } else {
                                    fmt.Println(v2)
                                    switch fi13 {
                                    case 2147483647:
                                        v21 += v21
                                        v39 := GenS5{F0: ("" + "foo bar"), F1: (((("a" + "\n\t") + ((("\n\t" + "hello") + "\n\t") + (("\n\t" + "a") + ("foo bar" + "\n\t")))) + "foo bar") + (("\n\t" + (("hello" + "hello") + ("a" + "hello"))) + ("hello" + "hello")))}
                                        switch 926 {
                                        case 0:
                                            fmt.Println(v2)
                                        case 4294967295:
                                            _ = len(v3)
                                        case 2147483647:
                                            var v40 int16 = -1
                                            _ = v40
                                        default:
                                            var v41 int32 = -2147483648
                                            if len(v3) > 0 {
                                                v3[0] = 33
                                            }
                                            _ = v41
                                        }
                                        _ = v39
                                    case -2147483648:
                                        v21 = (v21 * (v21 ^ 1))
                                    }
                                    fmt.Println(v2, fi13)
                                }
                            }
                            v1 = true
                            var v42 int8 = -63
                            _ = v21
                            _ = v42
                        case 32768:
                            if (v1 && (v11 < (v11 - v11))) {
                                v1 = (!(!v1))
                            }
                        }
                        var v43 int32 = 0
                        v43 = v43
                        v1 = false
                        _ = v43
                    }
                    _ = fi14
                }
                _ = fi13
            }
            switch "hello" {
            case "foo bar":
                v1 = v1
            case "hello":
                v11 = (v11 | v11)
            case "\n\t":
                for fi44 := 0; fi44 < 19; fi44++ {
                    if true {
                        var v45 int16 = -1
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 20; fi46++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (v1 && (!(3.14159 > (3.14159 - 1e10)))) {
                                switch "hello" {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "":
                                    fmt.Println(v2, fi46, v2)
                                case "foo bar":
                                    var v47 uint8 = 220
                                    _ = v47
                                default:
                                    if v1 {
                                        fmt.Println(v11, fi44, fi46)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            }
                            fmt.Println(fi44, v11, fi44)
                            _ = fi46
                        }
                        fmt.Println(fi44, v11, v11)
                    }
                    _ = fi44
                }
            }
            v1 = true
        }
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS6{F0: 4294967295, F1: 1, F2: v2}
            fmt.Println(fi48)
            _ = fi48
            _ = v49
        }
    }
    if true {
        if v1 {
            for fi50 := 0; fi50 < 14; fi50++ {
                v11 -= 18446744073709551615
                for fi51 := 0; fi51 < 19; fi51++ {
                    var v52 int64 = -1
                    _ = fi51
                    _ = v52
                }
                var v53 string = "foo bar"
                _ = fi50
                _ = v53
            }
        }
    } else {
        var v54 bool = false
        v1 = v1
        _ = v54
    }
    if v1 {
        for fi55 := 0; fi55 < 20; fi55++ {
            fmt.Println(v2, fi55)
            var v56 string = "a"
            _ = fi55
            _ = v56
        }
        v11 = (v11 >> ((v11) & 63))
    }
    for fi57 := 0; fi57 < 1; fi57++ {
        fmt.Println(v1)
        var v58 uint64 = 154
        _ = fi57
        _ = v58
    }
    var v59 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v59
}
