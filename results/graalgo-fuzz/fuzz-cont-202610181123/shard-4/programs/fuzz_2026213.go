package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{65535, 0}
    var v3 int32 = 749
    v4 := make(map[string]bool)
    var v5 uint64 = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v2 = append(v2, 170)
    } else {
        var v6 uint8 = 127
        _ = v6
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v1 -= v1
            var v10 string = "foo bar"
            var v11 uint16 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v3, v3)
        v12 := GenS2{F0: -1, F1: ((!true) || false), F2: int64(v3)}
        _ = fi8
        _ = v12
    }
    v1 = (-(v1))
    if true {
        var v13 bool = true
        var v14 uint64 = v5
        switch 2147483647 {
        case 32767:
            fmt.Println(v5)
        case 42:
            var v15 float64 = 1e10
            _ = v15
        case 9223372036854775807:
            if false {
                v2 = append(v2, 65535)
                var v16 uint8 = 1
                _ = v16
            } else {
                v3 = 42
                var v17 int16 = v1
                switch "hello" {
                case "":
                    if (1e10 > ((float64(float32(((1e10 - float64(float32(float64(3.14159)))) * 0.0))) + (-2.25 + -2.25)) * 3.14159)) {
                        if true {
                            var v18 int32 = v3
                            if (false || true) {
                                v19 := GenS3{F0: -1, F1: 42}
                                switch "\n\t" {
                                case "hello":
                                    v5 -= (^(v5))
                                case "\n\t":
                                    for fi20 := 0; fi20 < 4; fi20++ {
                                        fmt.Println(v14)
                                        var v21 int = int(v3)
                                        switch ("hello" + (("\n\t" + "\n\t") + "foo bar")) {
                                        case "hello":
                                            var v22 uint32 = 406
                                            _ = v22
                                        case "foo bar":
                                            v23 := GenS2{F0: -44, F1: (320 > 42), F2: -1}
                                            _ = v23
                                        case "a":
                                            var v24 int = (fi20 | (v21 & fi20))
                                            _ = v24
                                        default:
                                            v14 = (v14 / (((v14 + v14)) | 1))
                                        }
                                        _ = fi20
                                        _ = v21
                                    }
                                case "a":
                                    if (v13 && (2 != 2)) {
                                        var v25 uint = 0
                                        if (v5 > 4294967295) {
                                            v1 = v1
                                            for fi26 := 0; fi26 < 13; fi26++ {
                                                var v27 uint32 = 912
                                                switch ("a" + "\n\t") {
                                                case "\n\t":
                                                    if (1e10 != (float64((((1e10 - ((1e10 / 1.5) * 3.14159)) + 0.1) / float32(-2.25))) + 0.0)) {
                                                        var v28 bool = false
                                                        _ = v28
                                                    }
                                                }
                                                _ = fi26
                                                _ = v27
                                            }
                                            switch 127 {
                                            case 1:
                                                fmt.Println(v5)
                                            case 0:
                                                var v29 uint32 = 4294967295
                                                _ = v29
                                            }
                                        }
                                        var v30 string = (("hello" + ("" + "a")) + "\n\t")
                                        _ = v25
                                        _ = v30
                                    }
                                default:
                                    var v31 bool = true
                                    _ = v31
                                }
                                fmt.Println(v1, v18, v18)
                                _ = v19
                            } else {
                                var v32 uint64 = 798
                                _ = v32
                            }
                            _ = v18
                        } else {
                            v3 = 0
                            for fi33 := 0; fi33 < 9; fi33++ {
                                v34 := GenS1{F0: 255}
                                _ = fi33
                                _ = v34
                            }
                            var v35 string = "\n\t"
                            _ = v35
                        }
                        var v36 uint = 65535
                        var v37 string = "a"
                        _ = v36
                        _ = v37
                    }
                case "hello":
                    var v38 int = -2147483648
                    _ = v38
                default:
                    var v39 uint = 699
                    _ = v39
                }
                _ = v17
            }
        default:
            fmt.Println(v5, v14)
        }
        _ = v13
        _ = v14
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi40
        }
        delete(v4, "")
        var v41 uint32 = uint32(v5)
        _ = v41
    }
    switch 4294967295 {
    case -2147483648:
        var v42 uint16 = 1
        _ = v42
    case 1:
        v4["\n\t"] = true
    default:
        v43 := GenS2{F0: 1, F1: ((-2.25 * 0.0) != ((float32(-2.25) * 0.0) / -2.25)), F2: 62}
        _ = v43
    }
    var v44 float64 = float64(0.0)
    if false {
        v3 = (v3 / ((v3) | 1))
    }
    v44 = (-2.25 * v44)
    var v45 float64 = float64(-2.25)
    switch 32768 {
    case -139:
        for fi46 := 0; fi46 < 6; fi46++ {
            switch (^(fi46)) {
            case 1:
                v45 = (v45 + (v45 - 0.0))
            }
            _ = fi46
        }
    case -1:
        v44 = 0.0
    case 42:
        if (!false) {
            for fi47 := 0; fi47 < 3; fi47++ {
                v44 = ((v44 - (0.1 / 0.1)) + v45)
                var v48 float64 = (0.1 * v44)
                _ = fi47
                _ = v48
            }
            if (662 != v1) {
                v3 -= (v3 * v3)
            }
            if (false && false) {
                for fi49 := 0; fi49 < 17; fi49++ {
                    delete(v4, "a")
                    v5 -= v5
                    var v50 uint = 1
                    _ = fi49
                    _ = v50
                }
                var v51 int32 = int32(v1)
                switch 42 {
                case 743:
                    v45 = v45
                case 2147483647:
                    fmt.Println(v3, v1)
                }
                _ = v51
            } else {
                if (!(false && true)) {
                    var v52 int32 = 0
                    _ = v52
                }
                v45 += v45
                switch 0 {
                case -867:
                    var v53 float64 = v44
                    _ = v53
                case 32768:
                    _ = len(v2)
                }
            }
        } else {
            if ("a" > "hello") {
                if true {
                    v3 = 42
                    fmt.Println(v45, v1)
                    if false {
                        v44 = ((v44 * v45) - float64((-2.25 / 1e10)))
                        v3 = (v3 + v3)
                        var v54 float32 = (((1.5 / -2.25) - ((-2.25 / (1e10 * ((0.1 / -2.25) / (((((((3.14159 / 1.5) * 1e10) - 0.0) / 1.5) - (1e10 + float32((((0.0 * -2.25) + float64(((0.0 + (((0.0 / (float32(-2.25) + 3.14159)) / 1.5) * (0.1 * -2.25))) / 1.5))) * v44)))) * 1e10) - ((0.0 * 0.1) + (1.5 - 0.1)))))) - 1e10)) * 3.14159)
                        _ = v54
                    } else {
                        if (127 >= 42) {
                            var v55 int8 = 30
                            v3 += v3
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 10; fi56++ {
                            switch int(v3) {
                            case 388:
                                for fi57 := 0; fi57 < 7; fi57++ {
                                    delete(v4, "hello")
                                    v58 := GenS4{F0: 2147483647}
                                    v59 := GenS3{F0: 263, F1: 255}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            }
                            var v60 int64 = 616
                            var v61 uint8 = 127
                            _ = fi56
                            _ = v60
                            _ = v61
                        }
                    }
                } else {
                    if (32767 <= -2147483648) {
                        v62 := GenS5{F0: ""}
                        fmt.Println(v44, v44, v3)
                        fmt.Println(v45, v1)
                        _ = v62
                    }
                    var v63 uint32 = 2
                    var v64 float32 = 0.1
                    _ = v63
                    _ = v64
                }
                v65 := GenS6{F0: (18446744073709551615 != 486)}
                var v66 uint64 = (v5 / ((4294967295) | 1))
                _ = v65
                _ = v66
            }
        }
    }
    v45 = ((3.14159 / v45) * (3.14159 * ((float64(0.1) * 0.0) + v44)))
    if (0.0 <= (0.1 - (3.14159 + 0.0))) {
        var v67 uint8 = 128
        fmt.Println(v3)
        var v68 int = 0
        _ = v67
        _ = v68
    } else {
        var v69 int8 = 127
        v70 := GenS2{F0: -9223372036854775808, F1: (!true), F2: 850}
        _ = v69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v7
}
