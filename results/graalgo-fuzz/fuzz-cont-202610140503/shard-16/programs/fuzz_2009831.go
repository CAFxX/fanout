package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int16; F2 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32(-2.25)
    var v3 uint8 = 67
    v4 := []string{}
    var v5 int64 = 9223372036854775807
    var v6 uint = uint(v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 -= (^(v6))
        var v8 int = fi7
        var v9 uint16 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 557, F1: uint64(v6)}
    for fi11 := 0; fi11 < 14; fi11++ {
        switch "\n\t" {
        case "hello":
            var v12 string = "foo bar"
            _ = v12
        case "":
            if true {
                v3 -= v3
            }
        default:
            if false {
                v13 := GenS2{F0: -90}
                v2 = -2.25
                _ = v13
            } else {
                v2 += 3.14159
            }
        }
        var v14 float32 = (v2 + 1e10)
        _ = fi11
        _ = v14
    }
    v6 = v6
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "hello"
        switch fi15 {
        case 32768:
            fmt.Println(fi15)
        case -1:
            v4 = append(v4, v16)
        case 127:
            var v17 uint = (v6 >> ((180) & 63))
            _ = v17
        default:
            var v18 uint = (v6 << ((v6) & 63))
            _ = v18
        }
        var v19 uint8 = 2
        _ = fi15
        _ = v16
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if (((false || (!false)) && false) && false) {
            v21 := GenS3{F0: (fi20 << ((18446744073709551615) & 63))}
            var v22 uint8 = 104
            v2 -= v2
            _ = v21
            _ = v22
        } else {
            var v23 uint32 = 313
            switch (fi20 << ((uint(v3)) & 63)) {
            case 0:
                switch "\n\t" {
                case "foo bar":
                    if true {
                        switch ("hello" + "foo bar") {
                        case "":
                            for fi24 := 0; fi24 < 5; fi24++ {
                                if true {
                                    if (!(((fi24 | fi20) <= 32767) && (!(((42 != int32(fi20)) || true) || false)))) {
                                        v6 = 65535
                                        delete(v1, "\n\t")
                                        if true {
                                            v1["a"] = "foo bar"
                                        }
                                        v25 := GenS3{F0: 1}
                                        switch "foo bar" {
                                        case "hello":
                                            if (v5 == (-(v5))) {
                                                v5 = v5
                                                switch (fi24 / ((fi20) | 1)) {
                                                case 9223372036854775807:
                                                    fmt.Println(v3, v2, v2)
                                                case 32767:
                                                    v6 -= (v6 & (v6 << ((222) & 63)))
                                                default:
                                                    v26 := GenS4{F0: (-128 < 22), F1: 127}
                                                    _ = v26
                                                }
                                            } else {
                                                var v27 uint32 = 2147483647
                                                v27 = (v23 / ((4294967295) | 1))
                                                fmt.Println(v3, v23)
                                                v23 = v27
                                                _ = v27
                                            }
                                        case "\n\t":
                                            v5 = v5
                                            var v28 uint64 = 1
                                            if true {
                                                v4 = append(v4, ((((((("hello" + ("a" + "\n\t")) + "a") + "a") + "\n\t") + "foo bar") + "") + ("" + ("hello" + "hello"))))
                                                _ = len(v4)
                                            }
                                            var v29 int = fi20
                                            _ = v28
                                            _ = v29
                                        default:
                                            var v30 uint64 = 1
                                            _ = v30
                                        }
                                        _ = v25
                                    }
                                } else {
                                    v3 -= v3
                                    v2 = float32(((0.1 / 1.5) + (0.1 - float64(float32(1e10)))))
                                    switch (("foo bar" + "foo bar") + "hello") {
                                    case "hello":
                                        switch "hello" {
                                        case "a":
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = ("\n\t" + ("" + ("foo bar" + ("hello" + "foo bar"))))
                                                }
                                            }
                                        case "hello":
                                            fmt.Println(fi20)
                                        default:
                                            if true {
                                                v3 -= (v3 * (v3 * v3))
                                                var v31 uint8 = 42
                                                v2 = ((v2 * v2) / ((v2 - (v2 * (v2 - ((v2 * (float32(float64(((3.14159 * v2) * 0.1))) + v2)) * 3.14159)))) - v2))
                                                v6 -= v6
                                                _ = v31
                                            }
                                        }
                                    }
                                    var v32 bool = (("a" + "a") == "\n\t")
                                    _ = v32
                                }
                                _ = fi24
                            }
                        case "a":
                            var v33 int32 = 0
                            _ = v33
                        case "\n\t":
                            fmt.Println(v3, v2, v5)
                        }
                        if (!true) {
                            for fi34 := 0; fi34 < 18; fi34++ {
                                v35 := GenS5{F0: "\n\t"}
                                _ = fi34
                                _ = v35
                            }
                        } else {
                            var v36 int = (^(fi20))
                            v37 := GenS6{F0: 42, F1: 586, F2: 1}
                            for fi38 := 0; fi38 < 9; fi38++ {
                                _ = len(v4)
                                _ = fi38
                            }
                            _ = v36
                            _ = v37
                        }
                        fmt.Println(v2, v5, v2)
                    } else {
                        v3 += v3
                        if false {
                            var v39 int16 = int16(v23)
                            var v40 string = "a"
                            _ = v39
                            _ = v40
                        }
                        var v41 int = fi20
                        _ = v41
                    }
                }
            case 4294967295:
                for fi42 := 0; fi42 < 2; fi42++ {
                    if true {
                        v1["foo bar"] = "hello"
                    }
                    var v43 uint8 = uint8(v6)
                    _ = fi42
                    _ = v43
                }
            case -337:
                fmt.Println(v5, v5, v3)
            default:
                v3 = 128
            }
            v23 -= v23
            _ = v23
        }
        v44 := GenS5{F0: ("foo bar" + "\n\t")}
        if (907 < 0) {
            var v45 int64 = v5
            v6 = v6
            _ = v45
        }
        _ = fi20
        _ = v44
    }
    fmt.Println(v3, v2, v3)
    if true {
        v2 += 0.0
        var v46 int32 = 667
        v47 := GenS2{F0: -1}
        _ = v46
        _ = v47
    } else {
        var v48 int32 = 0
        v4 = append(v4, "foo bar")
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
