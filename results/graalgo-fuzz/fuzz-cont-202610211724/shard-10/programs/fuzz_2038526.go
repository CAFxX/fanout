package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int; F2 bool }
type GenS5 struct { F0 int16; F1 int; F2 bool }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 string; F1 int; F2 int8 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int; F1 uint64; F2 uint32 }

func main() {
    var v1 int64 = -68
    var v2 int8 = -98
    var v3 int16 = -1
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    var v6 float32 = v4
    v7 := []int{32767, 128, -9223372036854775808}
    var v8 uint8 = 250
    v3 -= -32768
    v9 := v5["a"]
    switch (v9 + ((v9 + v9) + (v9 + ""))) {
    case "":
        fmt.Println(v8, v1)
    }
    if (false && false) {
        if true {
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: v9}
                switch v9 {
                case "foo bar":
                    var v12 float64 = float64(1e10)
                    _ = v12
                case "":
                    v13 := GenS1{F0: v9}
                    _ = v13
                case "\n\t":
                    if true {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            v15 := GenS2{F0: 65535}
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 4; fi16++ {
                            var v17 uint64 = 4294967295
                            _ = len(v5)
                            _ = fi16
                            _ = v17
                        }
                        v6 = v4
                    }
                default:
                    switch v9 {
                    case "":
                        v9 = v9
                    case "a":
                        fmt.Println(fi10, fi10)
                    default:
                        v18 := GenS3{F0: v9}
                        _ = v18
                    }
                }
                _ = fi10
                _ = v11
            }
            if (!false) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 bool = false
                    _ = fi19
                    _ = v20
                }
            }
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
        } else {
            v4 = (v6 - float32((float64((float32(0.0) + 0.0)) + (-2.25 - 3.14159))))
            if false {
                switch -700 {
                case -9223372036854775808:
                    v21 := GenS4{F0: (false && true), F1: 32767, F2: ((true || true) || false)}
                    _ = v21
                case 848:
                    var v22 string = v9
                    _ = v22
                case 1:
                    delete(v5, (v9 + ("foo bar" + ("\n\t" + v9))))
                default:
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if false {
                            if (!((!false) && (v6 != v4))) {
                                v24 := GenS1{F0: "\n\t"}
                                _ = v24
                            }
                            v4 -= v6
                            v6 -= (float32(1.5) + ((v6 - 0.0) + v4))
                        }
                        _ = len(v5)
                        _ = fi23
                    }
                }
            } else {
                var v25 bool = (v4 == v6)
                fmt.Println(v4, v4)
                v8 -= (^(v8))
                _ = v25
            }
        }
    } else {
        v9 = v9
        switch -365 {
        case 0:
            if false {
                var v26 int16 = 32767
                var v27 float32 = v4
                for fi28 := 0; fi28 < 13; fi28++ {
                    fmt.Println(v9, fi28, v26)
                    v29 := GenS5{F0: (v26 ^ (v26 >> ((0) & 63))), F1: fi28, F2: (v1 < (^(v1)))}
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            } else {
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 uint = 2
                    v31 = v31
                    _ = fi30
                    _ = v31
                }
                var v32 int64 = (v1 | v1)
                v33 := GenS3{F0: v9}
                _ = v32
                _ = v33
            }
        }
    }
    v5["a"] = v9
    v3 = v3
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 uint16 = 1
            v4 = v6
            fmt.Println(v8, v2, v36)
            _ = fi35
            _ = v36
        }
        switch "hello" {
        case "":
            var v37 int64 = v1
            _ = v37
        case "foo bar":
            if true {
                for fi38 := 0; fi38 < 6; fi38++ {
                    if false {
                        if (255 == v8) {
                            var v39 bool = true
                            var v40 uint = 1
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(fi38)
                        v41 := GenS1{F0: v9}
                        _ = v41
                    } else {
                        _ = len(v5)
                    }
                    var v42 int = fi38
                    v43 := GenS6{F0: (v2 & v2), F1: false}
                    _ = fi38
                    _ = v42
                    _ = v43
                }
                v44 := GenS7{F0: (v1 ^ (v1 ^ (v1 * -1))), F1: 255}
                _ = v44
            }
        case "hello":
            _ = len(v5)
        default:
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 uint64 = 1
                _ = fi45
                _ = v46
            }
        }
        var v47 uint32 = 2147483647
        _ = fi34
        _ = v47
    }
    var v48 uint8 = (v8 << ((4294967295) & 63))
    if ((((((-2.25 / (1.5 * 1e10)) / 1.0) - 3.14159) * 1.5) - ((3.14159 / (1e10 / (float64((0.0 / 1.0)) * (0.1 / (float64(v6) / 1e10))))) * 0.1)) <= (0.1 * 1e10)) {
        var v49 int32 = -1
        if true {
            switch 32768 {
            case 1:
                for fi50 := 0; fi50 < 3; fi50++ {
                    var v51 float32 = float32(3.14159)
                    if len(v7) > 0 {
                        v7[0] = 2147483647
                    }
                    v1 = (v1 << ((305) & 63))
                    _ = fi50
                    _ = v51
                }
            case 204:
                switch (v9 + "a") {
                case "":
                    v52 := GenS6{F0: (v2 << ((4294967295) & 63)), F1: (v3 != (v3 + v3))}
                    _ = v52
                case "foo bar":
                    switch 769 {
                    case -2147483648:
                        v53 := GenS8{F0: (("hello" + "foo bar") + ("" + v9)), F1: -9223372036854775808, F2: 0}
                        _ = v53
                    default:
                        switch -88 {
                        case 4294967295:
                            switch 32767 {
                            case 197:
                                v5[v9] = v9
                            }
                        case 128:
                            switch -464 {
                            case 32768:
                                switch v9 {
                                case "\n\t":
                                    v6 -= (float32(-2.25) * 1e10)
                                case "foo bar":
                                    v54 := GenS9{F0: ("a" + v9), F1: v49}
                                    _ = v54
                                case "hello":
                                    if (((3.14159 * 1e10) / -2.25) < 0.0) {
                                        v5[""] = "hello"
                                    } else {
                                        switch 32767 {
                                        case 4294967295:
                                            if false {
                                                v55 := GenS5{F0: 769, F1: 2147483647, F2: true}
                                                _ = v55
                                            } else {
                                                v7 = append(v7, -1)
                                                var v56 int8 = -128
                                                if false {
                                                    v6 = (v4 * 1.5)
                                                } else {
                                                    v48 = v48
                                                    for fi57 := 0; fi57 < 7; fi57++ {
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        if (false || false) {
                                                            _ = len(v7)
                                                            for fi58 := 0; fi58 < 1; fi58++ {
                                                                v6 = 0.0
                                                                var v59 uint16 = 49
                                                                _ = fi58
                                                                _ = v59
                                                            }
                                                        }
                                                        var v60 int32 = (v49 ^ (v49 - v49))
                                                        _ = fi57
                                                        _ = v60
                                                    }
                                                }
                                                _ = v56
                                            }
                                        }
                                    }
                                default:
                                    if true {
                                        v3 += (v3 - 1)
                                        if (384 <= 935) {
                                            switch 938 {
                                            case 4294967295:
                                                var v61 float64 = ((0.1 * 1.5) - -2.25)
                                                _ = v61
                                            case 2147483647:
                                                for fi62 := 0; fi62 < 15; fi62++ {
                                                    v9 = (v9 + v9)
                                                    _ = fi62
                                                }
                                            }
                                            fmt.Println(v9, v4, v9)
                                        }
                                    }
                                }
                            }
                        case 9223372036854775807:
                            switch 4294967295 {
                            case -2147483648:
                                var v63 uint = 1
                                _ = v63
                            case 127:
                                v64 := GenS10{F0: v2}
                                _ = v64
                            case 822:
                                fmt.Println(v4, v4)
                            }
                        }
                    }
                }
            case 4294967295:
                for fi65 := 0; fi65 < 12; fi65++ {
                    for fi66 := 0; fi66 < 2; fi66++ {
                        v67 := GenS11{F0: (-(fi65)), F1: 539, F2: 2147483648}
                        _ = fi66
                        _ = v67
                    }
                    _ = fi65
                }
            }
        }
        v8 = 50
        _ = v49
    }
    fmt.Println(v4, v4)
    if (!(false && (true || true))) {
        var v68 float32 = (v6 / -2.25)
        fmt.Println(v48)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v68
    }
    if (4294967295 <= 811) {
        v48 = (v48 & 0)
        fmt.Println(v4)
        v69 := GenS7{F0: (v1 * v1), F1: 0}
        _ = v69
    }
    switch 999 {
    case -267:
        switch 42 {
        case 27:
            fmt.Println(v8)
        case 32767:
            var v70 int = 127
            _ = v70
        case 32768:
            fmt.Println(v1, v8, v1)
        }
    case 289:
        switch ("hello" + "hello") {
        case "":
            if (false && (v4 >= 0.1)) {
                v71 := GenS5{F0: (v3 | (v3 >> ((255) & 63))), F1: 0, F2: false}
                var v72 string = "\n\t"
                _ = v71
                _ = v72
            } else {
                var v73 uint16 = 903
                _ = v73
            }
        case "hello":
            if ((1.5 * 0.1) < 1.5) {
                v9 = ((v9 + v9) + (v9 + v9))
            } else {
                var v74 float64 = (1e10 / (float64(v6) / 1.0))
                _ = v74
            }
        }
    default:
        delete(v5, (v9 + v9))
    }
    var v75 uint8 = 195
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v75:", v75)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v75
    _ = v8
    _ = v9
}
