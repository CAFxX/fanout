package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    var v1 string = ""
    var v2 int16 = 1
    var v3 int8 = -95
    v4 := []string{"foo bar", v1, (v1 + v1), ("" + "foo bar")}
    v5 := make(map[int]string)
    v6 := v5[633]
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: -389}
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                var v10 int = (fi7 << ((2) & 63))
                _ = v10
            } else {
                if (v6 >= "hello") {
                    fmt.Println(v6, v1)
                    v6 = v1
                    var v11 uint8 = 127
                    _ = v11
                }
                v1 = (("hello" + v1) + v1)
                switch "" {
                case "a":
                    fmt.Println(v2, v2, fi9)
                case "":
                    switch (fi7 >> ((2) & 63)) {
                    case 127:
                        v1 = "a"
                    case -9223372036854775808:
                        v1 = v1
                    }
                }
                var v12 float64 = 3.14159
                _ = v12
            }
            var v13 uint = 255
            fmt.Println(v3, v1)
            if true {
                _ = len(v5)
                var v14 int8 = 3
                _ = v14
            }
            var v15 uint16 = 2
            _ = fi9
            _ = v13
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    switch (v6 + "hello") {
    case "":
        v6 = v1
    case "a":
        if true {
            v2 += 42
            if (!((((v2 != v2) || false) || false) || false)) {
                v3 -= v3
                if (315 > -1) {
                    fmt.Println(v6)
                } else {
                    fmt.Println(v6)
                    v16 := GenS2{F0: (v3 / (((v3 << ((42) & 63))) | 1)), F1: v6, F2: -175}
                    var v17 uint16 = 255
                    _ = v16
                    _ = v17
                }
                var v18 float64 = ((0.1 / 3.14159) * float64(1.5))
                _ = v18
            }
            _ = len(v5)
        } else {
            fmt.Println(v2, v3, v2)
            var v19 int8 = v3
            _ = v19
        }
    }
    fmt.Println(v1, v6, v1)
    fmt.Println(v1)
    v1 = v6
    var v20 uint16 = 2
    switch v6 {
    case "\n\t":
        v21 := GenS2{F0: (-(v3)), F1: "foo bar", F2: 499}
        _ = v21
    default:
        var v22 int16 = 313
        _ = v22
    }
    switch 128 {
    case 335:
        v23 := GenS3{F0: 9223372036854775807}
        _ = v23
    }
    var v24 uint8 = 128
    switch "" {
    case "\n\t":
        v4 = append(v4, "a")
    case "a":
        var v25 uint32 = 2
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := GenS4{F0: -9223372036854775808, F1: -2147483648}
            switch v6 {
            case "foo bar":
                switch fi26 {
                case 709:
                    v20 = v20
                case 821:
                    var v28 uint32 = 2147483648
                    _ = v28
                default:
                    if (!((true || true) && true)) {
                        var v29 uint8 = v24
                        switch 932 {
                        case 650:
                            fmt.Println(fi26, v1, v1)
                        case 47:
                            v2 = 42
                        case 1:
                            var v30 string = ""
                            _ = v30
                        default:
                            switch (fi26 % ((-418) | 1)) {
                            case 32768:
                                if false {
                                    v20 += (^(v20))
                                    var v31 float32 = float32(((0.0 - (1e10 / 3.14159)) * ((1e10 / 1.0) - 0.0)))
                                    _ = v31
                                }
                            case 495:
                                v4 = append(v4, v6)
                            case 2147483647:
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    switch "\n\t" {
                                    case "":
                                        v24 = (v29 * (v24 % ((5) | 1)))
                                    case "foo bar":
                                        v2 += v2
                                    default:
                                        switch -140 {
                                        case 312:
                                            var v33 uint16 = v20
                                            _ = v33
                                        case 4294967295:
                                            v34 := GenS5{F0: (fi32 + 32767), F1: -838}
                                            _ = v34
                                        }
                                    }
                                    _ = fi32
                                }
                            default:
                                v5[32767] = v1
                            }
                        }
                        if (v24 >= 2) {
                            _ = len(v4)
                            v6 = ((v1 + "") + "hello")
                            if (false && false) {
                                var v35 int8 = v3
                                _ = v35
                            }
                        } else {
                            _ = len(v5)
                        }
                        _ = v29
                    }
                }
            }
            _ = fi26
            _ = v27
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            fmt.Println(v6)
            v5[4294967295] = (v6 + (v1 + "foo bar"))
        } else {
            for fi36 := 0; fi36 < 15; fi36++ {
                switch (v6 + "") {
                case "":
                    v4 = append(v4, "foo bar")
                case "a":
                    v1 = v1
                case "\n\t":
                    var v37 uint64 = 4294967295
                    _ = v37
                }
                if false {
                    var v38 int16 = (v2 + (v2 + 32767))
                    _ = len(v4)
                    for fi39 := 0; fi39 < 17; fi39++ {
                        v40 := GenS3{F0: 0}
                        fmt.Println(v20, v24, fi39)
                        v6 = ("hello" + v6)
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
                fmt.Println(fi36, v6)
                _ = fi36
            }
            v20 += v20
            var v41 int8 = (v3 & v3)
            _ = v41
        }
    } else {
        v2 += v2
    }
    v1 = (v1 + v1)
    fmt.Println(v20, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
