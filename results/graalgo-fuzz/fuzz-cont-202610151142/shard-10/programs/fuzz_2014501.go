package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("" + "foo bar")
    v3 := make(map[int]bool)
    v4 := []string{v2, "hello", (v2 + v2)}
    v5 := make(map[string]bool)
    _ = len(v3)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    v6 := GenS1{F0: 32767}
    if false {
        switch v2 {
        case "a":
            switch "foo bar" {
            case "\n\t":
                switch 127 {
                case 127:
                    switch -542 {
                    case 32768:
                        v7 := GenS1{F0: -2147483648}
                        _ = v7
                    case -859:
                        if (!(((3.14159 + float64((1.5 - 0.0))) <= 1e10) && false)) {
                            switch "\n\t" {
                            case "foo bar":
                                if false {
                                    if (65535 < 258) {
                                        switch -1 {
                                        case 4294967295:
                                            fmt.Println(v2)
                                        case 793:
                                            fmt.Println(v2)
                                        case -882:
                                            fmt.Println(v2)
                                        }
                                        fmt.Println(v2)
                                        switch v2 {
                                        case "hello":
                                            v2 = (v2 + v2)
                                        default:
                                            switch (v2 + "\n\t") {
                                            case "\n\t":
                                                for fi8 := 0; fi8 < 20; fi8++ {
                                                    v2 = "hello"
                                                    _ = fi8
                                                }
                                            }
                                        }
                                    } else {
                                        v9 := GenS1{F0: 32767}
                                        v10 := GenS1{F0: -508}
                                        _ = v10
                                        _ = v9
                                    }
                                    v2 = ("\n\t" + v2)
                                    if false {
                                        switch 128 {
                                        case 664:
                                            var v11 int = 2147483647
                                            _ = v11
                                        }
                                        var v12 int = 9223372036854775807
                                        _ = v12
                                    }
                                }
                            case "a":
                                v2 = v2
                            case "":
                                for fi13 := 0; fi13 < 11; fi13++ {
                                    if (((!((("a" + "hello") >= "hello") && false)) || (1 <= -1)) && (!(false && (true || (false && false))))) {
                                        var v14 string = ("a" + (v2 + ((((v2 + v2) + v2) + v2) + (v2 + "foo bar"))))
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v14
                                    } else {
                                        v2 = (("\n\t" + (v2 + (("" + "foo bar") + (v2 + (((("\n\t" + v2) + "foo bar") + ("hello" + v2)) + "a"))))) + "")
                                        var v15 uint = 65535
                                        _ = v15
                                    }
                                    for fi16 := 0; fi16 < 17; fi16++ {
                                        v17 := GenS2{F0: 417}
                                        _ = fi16
                                        _ = v17
                                    }
                                    _ = fi13
                                }
                            default:
                                v2 = (("hello" + v2) + "a")
                            }
                            v4 = append(v4, ("\n\t" + "hello"))
                            v2 = "a"
                        } else {
                            var v18 uint32 = 2
                            if true {
                                v18 = v18
                            } else {
                                v2 = ("" + "")
                                var v19 float32 = (0.1 / (3.14159 - 1.5))
                                if (true && (v19 > (v19 - (float32(-2.25) + (v19 / v19))))) {
                                    fmt.Println(v18, v18)
                                    var v20 int = 552
                                    var v21 int = (^(v20))
                                    _ = v20
                                    _ = v21
                                }
                                _ = v19
                            }
                            if false {
                                v18 = (v18 / ((4294967295) | 1))
                                if (42 > 91) {
                                    if false {
                                        var v22 uint8 = 143
                                        v2 = v2
                                        v2 = ((((v2 + (v2 + v2)) + "") + "a") + ((v2 + v2) + ((v2 + v2) + "a")))
                                        _ = v22
                                    }
                                    var v23 string = "a"
                                    fmt.Println(v18, v2)
                                    _ = v23
                                }
                            }
                            _ = v18
                        }
                    case 32767:
                        for fi24 := 0; fi24 < 19; fi24++ {
                            v2 = v2
                            v25 := GenS1{F0: fi24}
                            _ = fi24
                            _ = v25
                        }
                    }
                case 0:
                    v4 = append(v4, v2)
                case 32768:
                    if (true || false) {
                        v2 = ("\n\t" + "\n\t")
                        for fi26 := 0; fi26 < 19; fi26++ {
                            fmt.Println(fi26)
                            var v27 int32 = -2147483648
                            v2 = "a"
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            case "foo bar":
                v2 = v2
            }
        case "hello":
            var v28 bool = (false || (!false))
            _ = v28
        case "":
            v2 = (((v2 + "") + (("foo bar" + ("a" + v2)) + "a")) + v2)
        }
    }
    switch v2 {
    case "hello":
        fmt.Println(v2)
    case "":
        fmt.Println(v2)
    default:
        _ = len(v1)
    }
    var v29 uint8 = 1
    switch 32768 {
    case 32768:
        v2 = ((v2 + v2) + v2)
    case 1:
        if (!false) {
            fmt.Println(v29, v29)
        } else {
            var v30 int32 = -119
            _ = v30
        }
    }
    if false {
        v31 := GenS2{F0: -36}
        switch (v2 + "foo bar") {
        case "a":
            if true {
                v2 = (("" + v2) + v2)
            } else {
                v5[((v2 + v2) + v2)] = false
            }
        case "\n\t":
            if (452 <= 255) {
                v2 = ((v2 + "\n\t") + v2)
                if (100 > 42) {
                    v2 = v2
                } else {
                    v2 = v2
                    v32 := v3[-997]
                    _ = v32
                }
                switch (((v2 + v2) + (v2 + v2)) + "\n\t") {
                case "":
                    v33 := GenS2{F0: 0}
                    _ = v33
                default:
                    v34 := GenS1{F0: 529}
                    _ = v34
                }
            } else {
                var v35 uint = 0
                switch 4294967295 {
                case 9223372036854775807:
                    v29 += (v29 % ((v29) | 1))
                case -9223372036854775808:
                    v36 := GenS1{F0: 1}
                    _ = v36
                }
                var v37 string = "foo bar"
                _ = v35
                _ = v37
            }
        }
        var v38 float32 = 0.1
        _ = v31
        _ = v38
    } else {
        v2 = ""
        for fi39 := 0; fi39 < 13; fi39++ {
            v2 = (v2 + (v2 + v2))
            v2 = "a"
            _ = fi39
        }
    }
    if (0 != 127) {
        for fi40 := 0; fi40 < 6; fi40++ {
            var v41 int32 = -1
            v2 = "a"
            _ = fi40
            _ = v41
        }
        var v42 int8 = 84
        v43 := v1[9223372036854775807]
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 float64 = 0.0
        if (1.5 < (0.1 - 1e10)) {
            v4 = append(v4, (("" + "\n\t") + v2))
        } else {
            if len(v4) > 0 {
                v4[0] = (v2 + v2)
            }
        }
        _ = fi44
        _ = v45
    }
    v2 = v2
    if (!true) {
        delete(v3, 32768)
        v46 := GenS2{F0: -32768}
        _ = v46
    }
    if true {
        var v47 string = "foo bar"
        var v48 float64 = (float64(1e10) / float64(float32((1e10 - -2.25))))
        if true {
            if ((-25 > -9223372036854775808) || true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            delete(v5, "foo bar")
            v49 := GenS3{F0: 42, F1: 53, F2: 2147483648}
            _ = v49
        }
        _ = v47
        _ = v48
    } else {
        var v50 uint = 65535
        fmt.Println(v2, v2, v50)
        var v51 bool = false
        _ = v50
        _ = v51
    }
    var v52 uint64 = 6
    v2 = ""
    var v53 float32 = 0.0
    if (false && (false && (false || true))) {
        v2 = ((v2 + v2) + "")
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (42 >= 2) {
        v54 := v3[-1]
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
