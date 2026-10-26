package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint64 = 86
    v2 := []uint32{42, 65535, 2147483648, 2147483647}
    var v3 float64 = 1e10
    var v4 uint32 = 2
    v5 := []string{"", ("" + ""), "foo bar"}
    v1 = v1
    fmt.Println(v3)
    _ = len(v2)
    var v6 float64 = float64((((1.5 / 1.0) / float32((0.0 / v3))) + (float32((v3 * v3)) - 0.0)))
    if false {
        switch 0 {
        case -324:
            v4 = v4
        default:
            for fi7 := 0; fi7 < 18; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    if true {
                        if ((fi7 - (fi8 * (fi8 * -759))) != (-(fi7))) {
                            fmt.Println(v6, fi7, v6)
                            var v9 uint8 = 134
                            v10 := GenS1{F0: (((3.14159 / float64(1e10)) * ((v3 + float64(1e10)) - v3)) != v6)}
                            var v11 string = "hello"
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                        var v12 uint16 = 1
                        _ = v12
                    }
                    if true {
                        fmt.Println(fi7, v4)
                        fmt.Println(v6, v4)
                        v3 = float64((((3.14159 / float32(v6)) / -2.25) * 1.5))
                        fmt.Println(fi7, fi7, v1)
                        v3 = 0.0
                    } else {
                        var v13 uint = 18446744073709551615
                        var v14 uint16 = 2
                        _ = v13
                        _ = v14
                    }
                    var v15 uint16 = 42
                    _ = fi8
                    _ = v15
                }
                _ = fi7
            }
        }
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint64 = v1
            v17 -= 705
            var v18 int = (fi16 << ((255) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
        v3 = (v6 * 1e10)
        v19 := GenS1{F0: ((-400 < 127) || ((3.14159 * -2.25) != (1.5 + 0.0)))}
        _ = v19
    }
    switch ("\n\t" + ("foo bar" + "hello")) {
    case "":
        if (-2147483648 <= 81) {
            v1 -= (v1 * v1)
        } else {
            if (!((2147483647 >= 0) || (0 <= 0))) {
                if true {
                    v4 += (v4 << ((397) & 63))
                    v6 = v3
                }
            } else {
                var v20 uint16 = 2
                _ = v20
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                if (fi21 == 32767) {
                    var v22 uint = 4294967295
                    v23 := GenS2{F0: v4, F1: 42}
                    v24 := GenS1{F0: false}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
        }
    case "foo bar":
        for fi25 := 0; fi25 < 9; fi25++ {
            if len(v2) > 0 {
                v2[0] = v4
            }
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            _ = fi25
        }
    case "a":
        v6 -= 0.1
    }
    v1 = 0
    v3 += (0.1 + v3)
    for fi26 := 0; fi26 < 2; fi26++ {
        if ((!((false || true) || ((1e10 / float32(1.5)) <= float32(v6)))) && (v3 == float64((3.14159 + (1.5 * 3.14159))))) {
            if false {
                var v27 int32 = -2147483648
                if false {
                    switch "hello" {
                    case "\n\t":
                        for fi28 := 0; fi28 < 1; fi28++ {
                            switch "a" {
                            case "foo bar":
                                switch "hello" {
                                case "\n\t":
                                    var v29 uint16 = 237
                                    _ = v29
                                case "":
                                    var v30 int64 = -1
                                    _ = v30
                                default:
                                    var v31 int8 = 51
                                    _ = v31
                                }
                            case "hello":
                                var v32 int64 = 0
                                _ = v32
                            }
                            _ = fi28
                        }
                    case "hello":
                        _ = len(v2)
                    default:
                        var v33 float32 = float32(v6)
                        _ = v33
                    }
                }
                _ = v27
            }
        } else {
            v3 += float64(-2.25)
            var v34 string = (("\n\t" + "") + "hello")
            for fi35 := 0; fi35 < 14; fi35++ {
                v36 := GenS3{F0: 255, F1: 279}
                v4 = v4
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        if true {
            var v37 float64 = (float64(0.1) - float64((3.14159 * (float32(v6) / 1e10))))
            v6 = (v37 * v3)
            _ = v37
        }
        _ = fi26
    }
    var v38 uint32 = 875
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "\n\t")
    }
    fmt.Println(v6)
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := GenS3{F0: 0, F1: (v4 * (v4 & 0))}
        _ = fi39
        _ = v40
    }
    switch 32768 {
    case 4294967295:
        v41 := GenS2{F0: (v4 | v4), F1: -128}
        _ = v41
    case 127:
        fmt.Println(v1)
    case -9223372036854775808:
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := GenS2{F0: (^(v4)), F1: 0}
            v6 = (v6 + v3)
            var v44 uint16 = uint16(v4)
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    if (1 >= 0) {
        v1 = (v1 >> ((2) & 63))
        switch -970 {
        case 32768:
            var v45 float32 = -2.25
            _ = v45
        case -2147483648:
            switch (("foo bar" + ("foo bar" + "\n\t")) + ("" + ("foo bar" + ""))) {
            case "a":
                fmt.Println(v4, v4, v3)
            }
        case 1:
            fmt.Println(v3, v6, v4)
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = len(v5)
        v46 := GenS1{F0: (0 > -1)}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
