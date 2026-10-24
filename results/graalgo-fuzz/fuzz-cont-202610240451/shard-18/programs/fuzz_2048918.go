package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int }

func main() {
    v1 := []int{42}
    var v2 float32 = -2.25
    var v3 string = (("foo bar" + "foo bar") + "foo bar")
    switch 0 {
    case 2147483647:
        fmt.Println(v2, v2)
    case 32768:
        if (2 != 0) {
            var v4 uint16 = 643
            var v5 uint = 42
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v4
            _ = v5
        } else {
            fmt.Println(v3)
        }
    }
    v6 := GenS1{F0: (42 != -2147483648), F1: 65535}
    v7 := GenS1{F0: true, F1: 2147483647}
    var v8 uint32 = 65535
    v9 := GenS1{F0: true, F1: v8}
    switch -449 {
    case 2147483647:
        var v10 uint16 = 2
        _ = v10
    }
    v11 := GenS2{F0: (int32(v8) == 42), F1: v3}
    for fi12 := 0; fi12 < 19; fi12++ {
        v1 = append(v1, fi12)
        _ = fi12
    }
    if true {
        if ((v8 | 0) != 978) {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint16 = 555
                var v15 string = v3
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            if true {
                switch -440 {
                case -232:
                    fmt.Println(v3, v2)
                case 736:
                    v1 = append(v1, 705)
                default:
                    if false {
                        switch (("foo bar" + v3) + v3) {
                        case "\n\t":
                            v8 = v8
                        default:
                            if true {
                                for fi16 := 0; fi16 < 11; fi16++ {
                                    var v17 float32 = 3.14159
                                    var v18 int8 = -41
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                                v2 = v2
                            } else {
                                v2 = v2
                                v19 := GenS2{F0: false, F1: (v3 + v3)}
                                if false {
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        v3 = v3
                                        if true {
                                            v21 := GenS3{F0: 4294967295, F1: (fi20 % (((fi20 << ((814) & 63))) | 1)), F2: v8}
                                            var v22 int32 = 634
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 int = (fi20 >> ((2) & 63))
                                            for fi24 := 0; fi24 < 4; fi24++ {
                                                var v25 int = fi20
                                                v23 -= fi20
                                                _ = fi24
                                                _ = v25
                                            }
                                            v26 := GenS4{F0: v3, F1: 42}
                                            _ = v23
                                            _ = v26
                                        }
                                        _ = fi20
                                    }
                                }
                                _ = v19
                            }
                        }
                        var v27 int64 = -1
                        if (32768 == 127) {
                            v28 := GenS4{F0: (v3 + "a"), F1: 4294967295}
                            switch "foo bar" {
                            case "hello":
                                v2 = float32(0.1)
                            }
                            _ = v28
                        }
                        _ = v27
                    } else {
                        var v29 int16 = -589
                        v29 = (^(v29))
                        _ = len(v1)
                        _ = v29
                    }
                }
                v8 += (^(v8))
                _ = len(v1)
            }
            var v30 int32 = -364
            if true {
                v31 := GenS5{F0: -9223372036854775808}
                if (!false) {
                    var v32 float32 = 1.5
                    v30 += v30
                    _ = v32
                } else {
                    switch 32767 {
                    case 32768:
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v2 = v2
                            v3 = v3
                        } else {
                            if (true && true) {
                                switch 531 {
                                case -204:
                                    v1 = append(v1, 4294967295)
                                case 4294967295:
                                    for fi33 := 0; fi33 < 11; fi33++ {
                                        for fi34 := 0; fi34 < 12; fi34++ {
                                            switch v3 {
                                            case "hello":
                                                fmt.Println(v30, v30)
                                            case "a":
                                                if (!true) {
                                                    v35 := GenS1{F0: (v8 == v8), F1: v8}
                                                    if (!(true || true)) {
                                                        v36 := GenS5{F0: fi33}
                                                        var v37 int8 = 100
                                                        switch 128 {
                                                        case 127:
                                                            if true {
                                                                var v38 int = fi33
                                                                _ = v38
                                                            } else {
                                                                fmt.Println(fi34, v8)
                                                                v3 = v3
                                                            }
                                                        case -355:
                                                            v37 += (-(v37))
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    v30 = (^(v30))
                                                    var v39 uint8 = 105
                                                    _ = v35
                                                    _ = v39
                                                }
                                            default:
                                                v30 = v30
                                            }
                                            var v40 float64 = 1e10
                                            var v41 int8 = 100
                                            _ = fi34
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = fi33
                                    }
                                default:
                                    if false {
                                        var v42 float64 = float64(v2)
                                        if (18446744073709551615 <= 4294967295) {
                                            v3 = "\n\t"
                                            for fi43 := 0; fi43 < 3; fi43++ {
                                                v44 := GenS5{F0: fi43}
                                                for fi45 := 0; fi45 < 12; fi45++ {
                                                    v8 = v8
                                                    _ = fi45
                                                }
                                                v42 += ((v42 / (float64(v2) + (0.0 + (float64(v2) + v42)))) / v42)
                                                _ = fi43
                                                _ = v44
                                            }
                                        }
                                        _ = v42
                                    } else {
                                        v2 = v2
                                        for fi46 := 0; fi46 < 20; fi46++ {
                                            v30 = 42
                                            _ = fi46
                                        }
                                        switch (v3 + v3) {
                                        case "hello":
                                            var v47 int = 127
                                            _ = v47
                                        case "foo bar":
                                            fmt.Println(v2, v2, v30)
                                        default:
                                            switch "\n\t" {
                                            case "a":
                                                if true {
                                                    for fi48 := 0; fi48 < 6; fi48++ {
                                                        var v49 int64 = 9223372036854775807
                                                        v50 := GenS1{F0: (!true), F1: (v8 >> ((42) & 63))}
                                                        fmt.Println(v30)
                                                        _ = fi48
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                }
                                            case "foo bar":
                                                v1 = append(v1, 693)
                                            default:
                                                var v51 uint64 = 18446744073709551615
                                                _ = v51
                                            }
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v8, v3, v30)
                                case "a":
                                    var v52 int8 = 42
                                    _ = v52
                                case "foo bar":
                                    switch 42 {
                                    case 9223372036854775807:
                                        var v53 uint32 = (v8 * (v8 >> ((896) & 63)))
                                        _ = v53
                                    default:
                                        v8 += v8
                                    }
                                default:
                                    var v54 float32 = float32(0.0)
                                    _ = v54
                                }
                            } else {
                                v3 = v3
                                fmt.Println(v3, v2, v2)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                    case 127:
                        v2 += v2
                    case 2147483647:
                        v30 -= v30
                    }
                    var v55 string = ""
                    _ = v55
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v31
            } else {
                if (128 != 0) {
                    v56 := GenS3{F0: (v8 >> ((520) & 63)), F1: -2147483648, F2: v8}
                    fmt.Println(v3)
                    v8 += v8
                    _ = v56
                }
                v57 := GenS4{F0: (("\n\t" + "hello") + v3), F1: 255}
                _ = v57
            }
            _ = v30
        }
    }
    var v58 uint16 = 2
    if len(v1) > 0 {
        v1[0] = 62
    }
    v3 = v3
    for fi59 := 0; fi59 < 20; fi59++ {
        v2 = 1e10
        if len(v1) > 0 {
            v1[0] = fi59
        }
        _ = fi59
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        fmt.Println(fi60, v3, fi60)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v3, v2, v2)
        _ = fi60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v58
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
