package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int8 }

func main() {
    var v1 int16 = 1
    v2 := []string{""}
    var v3 bool = (((-2147483648 < 2147483647) && (0.0 > 0.1)) && true)
    var v4 float64 = (0.0 / 1.5)
    if v3 {
        v1 += (v1 * (v1 + v1))
    } else {
        var v5 uint64 = 179
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(v5, fi6)
                if v3 {
                    v3 = true
                    v8 := GenS1{F0: 42}
                    switch (fi6 & (-(fi7))) {
                    case 42:
                        if v3 {
                            if ("hello" > ("a" + "hello")) {
                                v1 = -1
                            }
                            v5 += (^(v5))
                            v9 := GenS2{F0: 0}
                            _ = v9
                        } else {
                            if v3 {
                                if v3 {
                                    if (true || v3) {
                                        fmt.Println(fi7, v1)
                                        v3 = (((("foo bar" <= ("\n\t" + "\n\t")) || (v5 < v5)) || true) && v3)
                                    }
                                    _ = len(v2)
                                    v3 = (127 != -1)
                                } else {
                                    var v10 int64 = 9223372036854775807
                                    _ = v10
                                }
                                v4 = v4
                                v11 := GenS2{F0: 0}
                                _ = len(v2)
                                _ = v11
                            } else {
                                var v12 int8 = 42
                                v13 := GenS1{F0: 0}
                                v5 += 4294967295
                                _ = v12
                                _ = v13
                            }
                        }
                    case 32768:
                        switch int(v1) {
                        case 32768:
                            switch "" {
                            case "a":
                                switch fi6 {
                                case -29:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    if false {
                                        var v14 int8 = 17
                                        _ = v14
                                    }
                                }
                                v2 = append(v2, (("hello" + "\n\t") + "foo bar"))
                                _ = len(v2)
                                v2 = append(v2, "")
                                if (v4 < (float64((1.5 + -2.25)) * v4)) {
                                    v5 -= v5
                                    fmt.Println(v3, v1)
                                    var v15 uint16 = 42
                                    switch fi6 {
                                    case 0:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 535:
                                        if (v3 || v3) {
                                            switch "hello" {
                                            case "hello":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if ((v5 >> ((794) & 63)) > 255) {
                                                        fmt.Println(fi7)
                                                    }
                                                case "\n\t":
                                                    v4 = (0.0 + v4)
                                                case "":
                                                    v16 := GenS1{F0: 127}
                                                    _ = v16
                                                }
                                            }
                                            v3 = v3
                                            var v17 int64 = -9223372036854775808
                                            _ = v17
                                        } else {
                                            var v18 uint = 65535
                                            _ = v18
                                        }
                                    }
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        var v19 float64 = float64(float32(v4))
                                        _ = v19
                                    case "a":
                                        v15 = 65535
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(v4)
                                    }
                                    if (-1 == 1) {
                                        fmt.Println(v3)
                                        v15 -= v15
                                    } else {
                                        var v20 string = "foo bar"
                                        v21 := GenS3{F0: 2}
                                        _ = v20
                                        _ = v21
                                    }
                                    v22 := GenS4{F0: (v15 & (v15 * (v15 ^ v15)))}
                                    var v23 uint16 = (v15 / (((v15 << ((1) & 63))) | 1))
                                    _ = v15
                                    _ = v22
                                    _ = v23
                                }
                                if len(v2) > 0 {
                                    v2[0] = ("foo bar" + "")
                                }
                                var v24 int = (^(fi7))
                                _ = v24
                            case "\n\t":
                                fmt.Println(v5, v3, v3)
                            case "hello":
                                v5 -= v5
                            default:
                                v3 = (-2.25 <= (((1.5 + (-2.25 + 1e10)) + ((3.14159 / ((1e10 * ((0.1 * float32(v4)) * (float32(1e10) * 1.5))) * (1e10 + 1.5))) * 1e10)) + 1e10))
                            }
                        }
                    default:
                        v25 := GenS3{F0: 2}
                        _ = v25
                    }
                    _ = v8
                }
                _ = fi7
            }
            fmt.Println(v3, v5, v4)
            v3 = v3
            _ = fi6
        }
        if (154 > 2) {
            v5 += (^(v5))
        } else {
            if (v3 && v3) {
                v4 = v4
            }
            var v26 uint = 2
            _ = v26
        }
        _ = v5
    }
    var v27 float32 = -2.25
    if (((127 == 42) || v3) && (3.14159 != (0.1 * float64(1.5)))) {
        switch "" {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "hello":
            if (!(!(363 != 2))) {
                for fi28 := 0; fi28 < 16; fi28++ {
                    v3 = v3
                    if (50 >= 976) {
                        switch "\n\t" {
                        case "a":
                            v4 = v4
                        case "\n\t":
                            fmt.Println(v27, v4, fi28)
                        default:
                            var v29 float64 = v4
                            _ = v29
                        }
                        if v3 {
                            if true {
                                if ("a" >= "a") {
                                    if false {
                                        fmt.Println(v27, v3)
                                        fmt.Println(v27, v27)
                                        var v30 int16 = v1
                                        _ = v30
                                    } else {
                                        if false {
                                            var v31 float64 = v4
                                            v32 := GenS1{F0: 100}
                                            _ = len(v2)
                                            _ = v31
                                            _ = v32
                                        }
                                        switch ("a" + "") {
                                        case "":
                                            var v33 uint64 = 7
                                            _ = v33
                                        }
                                        fmt.Println(v3)
                                    }
                                } else {
                                    var v34 int64 = -511
                                    if ((0 <= 128) || false) {
                                        fmt.Println(v27, v34, v34)
                                        if v3 {
                                            var v35 bool = (!v3)
                                            for fi36 := 0; fi36 < 11; fi36++ {
                                                if (v35 && ((fi36 - 2147483647) <= (-(fi28)))) {
                                                    v37 := GenS5{F0: (v34 % (((-(v34))) | 1))}
                                                    var v38 int64 = -9223372036854775808
                                                    _ = v37
                                                    _ = v38
                                                }
                                                _ = fi36
                                            }
                                            for fi39 := 0; fi39 < 20; fi39++ {
                                                var v40 uint32 = 4294967295
                                                v41 := GenS4{F0: 1}
                                                _ = fi39
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v35
                                        } else {
                                            for fi42 := 0; fi42 < 8; fi42++ {
                                                var v43 int16 = (v1 ^ (-(v1)))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = fi42
                                                _ = v43
                                            }
                                        }
                                    }
                                    _ = v34
                                }
                            }
                            var v44 int = (-(fi28))
                            var v45 int16 = v1
                            _ = v44
                            _ = v45
                        } else {
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                if v3 {
                                    fmt.Println(v3)
                                }
                            }
                            fmt.Println(fi28)
                            for fi46 := 0; fi46 < 12; fi46++ {
                                var v47 int = (fi46 << ((4294967295) & 63))
                                switch fi46 {
                                case 0:
                                    var v48 int32 = 1
                                    _ = v48
                                case 9223372036854775807:
                                    v49 := GenS6{F0: -8, F1: 127}
                                    _ = v49
                                default:
                                    v47 = 2147483647
                                }
                                v1 -= (v1 % ((1000) | 1))
                                _ = fi46
                                _ = v47
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = ("hello" + (("\n\t" + ("" + "a")) + "\n\t"))
                        }
                    } else {
                        v3 = v3
                    }
                    _ = fi28
                }
            }
        default:
            _ = len(v2)
        }
        var v50 float32 = float32(0.1)
        _ = v50
    } else {
        v1 += v1
    }
    v4 += -2.25
    var v51 uint8 = 1
    v52 := GenS2{F0: 111}
    v4 = (v4 / (float64(v27) * v4))
    fmt.Println(v3, v1)
    var v53 uint16 = 1
    v27 = float32(1.5)
    v1 = (v1 + v1)
    switch "hello" {
    case "foo bar":
        switch (("a" + (("a" + (("a" + "\n\t") + ("" + "\n\t"))) + "a")) + "hello") {
        case "a":
            if v3 {
                fmt.Println(v27)
                v53 = 65535
                var v54 float64 = -2.25
                _ = v54
            } else {
                fmt.Println(v27, v27)
                fmt.Println(v1, v4)
                var v55 uint = 386
                _ = v55
            }
        case "\n\t":
            v3 = (!(87 <= -31))
        }
    case "":
        var v56 int8 = 100
        _ = v56
    default:
        var v57 uint64 = 131
        _ = v57
    }
    for fi58 := 0; fi58 < 13; fi58++ {
        fmt.Println(fi58, v53)
        var v59 int8 = 127
        _ = fi58
        _ = v59
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
}
