package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }

func main() {
    var v1 float64 = (float64((0.0 * 1.5)) - 0.1)
    v2 := []uint32{65535, 600, 15}
    v3 := []string{"\n\t"}
    var v4 int16 = 1000
    v5 := []uint32{505, 59}
    v6 := make(map[int]string)
    if true {
        v3 = append(v3, "\n\t")
    }
    var v7 uint32 = 2147483648
    var v8 int16 = -870
    if true {
        v6[-9223372036854775808] = ""
        v8 += 1000
        if ((!(!true)) && true) {
            var v9 float64 = v1
            _ = v9
        } else {
            var v10 int32 = -2147483648
            _ = v10
        }
    }
    var v11 uint = 900
    var v12 int8 = 51
    var v13 uint64 = 628
    var v14 uint32 = 2147483647
    var v15 bool = ("foo bar" > (("hello" + "a") + ("a" + ("\n\t" + "\n\t"))))
    var v16 int64 = 249
    if v15 {
        switch "\n\t" {
        case "foo bar":
            for fi17 := 0; fi17 < 2; fi17++ {
                v11 -= (^(v11))
                for fi18 := 0; fi18 < 20; fi18++ {
                    _ = len(v6)
                    v6[-1] = "hello"
                    _ = fi18
                }
                v19 := GenS1{F0: "foo bar"}
                _ = fi17
                _ = v19
            }
        case "hello":
            v20 := GenS2{F0: v7, F1: 4294967295}
            _ = v20
        case "a":
            for fi21 := 0; fi21 < 16; fi21++ {
                switch "\n\t" {
                case "a":
                    v6[(fi21 << (((v11 | v11)) & 63))] = (("" + (("hello" + "") + "\n\t")) + ("a" + ""))
                case "\n\t":
                    if (!(-410 >= 438)) {
                        var v22 int16 = (v4 ^ v4)
                        v22 = 32767
                        _ = v22
                    }
                }
                if false {
                    var v23 string = "foo bar"
                    v16 = (-(v16))
                    v4 = 626
                    _ = v23
                } else {
                    v24 := GenS2{F0: 65535, F1: fi21}
                    v1 = float64(0.0)
                    _ = v24
                }
                var v25 int16 = (^(v8))
                _ = fi21
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 4; fi26++ {
                v16 = (v16 | v16)
                delete(v6, (fi26 >> ((v11) & 63)))
                v5 = append(v5, (v14 ^ 747))
                _ = fi26
            }
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            v12 = (v12 + 71)
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 string = (("foo bar" + "a") + ("hello" + "a"))
                if false {
                    v14 += 166
                    if v15 {
                        var v30 int64 = (v16 / ((v16) | 1))
                        fmt.Println(v30)
                        switch "" {
                        case "foo bar":
                            var v31 int16 = v4
                            if (v15 || v15) {
                                var v32 string = ("a" + ((v29 + ((v29 + ((v29 + v29) + "a")) + "a")) + ""))
                                _ = v32
                            } else {
                                v33 := GenS3{F0: (v13 | (v13 & (v13 * v13)))}
                                if (v12 >= (v12 - (v12 & v12))) {
                                    var v34 int32 = 764
                                    var v35 uint64 = (v13 ^ 1)
                                    switch 2147483647 {
                                    case 4294967295:
                                        fmt.Println(v14, fi28, v35)
                                    case 2147483647:
                                        v13 -= (v35 ^ (v13 ^ v35))
                                    case 128:
                                        if true {
                                            var v36 float64 = ((float64((0.0 / 1e10)) / (1e10 / float64(3.14159))) / 1.5)
                                            fmt.Println(v31, v31)
                                            _ = v36
                                        }
                                    default:
                                        v30 -= int64(v11)
                                    }
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch (fi27 ^ fi28) {
                                    case 127:
                                        v12 -= v12
                                    case -144:
                                        v30 += v30
                                    }
                                    v37 := GenS3{F0: (^(v13))}
                                    _ = v37
                                }
                                _ = v33
                            }
                            _ = v31
                        }
                        _ = v30
                    } else {
                        _ = len(v6)
                        fmt.Println(v11)
                        var v38 int = fi27
                        _ = v38
                    }
                }
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS1{F0: "\n\t"}
            switch ("hello" + "") {
            case "foo bar":
                var v41 int64 = 156
                _ = v41
            case "hello":
                v42 := GenS3{F0: v13}
                _ = v42
            default:
                v14 -= 2147483647
            }
            fmt.Println(v11, v1, v4)
            _ = fi39
            _ = v40
        }
    } else {
        switch -1 {
        case 2147483647:
            if ((v1 >= (v1 * (((v1 * (v1 + (((0.1 * v1) - v1) + (((v1 - v1) - 1.5) / v1)))) / v1) * 1e10))) && false) {
                v15 = (v15 || ((2 != 42) && true))
                fmt.Println(v1, v16, v12)
            }
        default:
            var v43 float64 = 3.14159
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        fmt.Println(v14, v13, v16)
        v7 = v7
        _ = fi44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        v46 := GenS1{F0: ("a" + "\n\t")}
        _ = fi45
        _ = v46
    }
    switch 2147483647 {
    case 42:
        var v47 uint8 = 255
        _ = v47
    case -154:
        fmt.Println(v16, v12)
    case -973:
        switch -120 {
        case -1:
            v48 := GenS4{F0: 4294967295, F1: 128, F2: (("" + "\n\t") + (("a" + ((("hello" + "") + (("foo bar" + "foo bar") + "foo bar")) + "")) + ("foo bar" + ("a" + ("a" + "")))))}
            _ = v48
        case 128:
            var v49 int = -2147483648
            _ = v49
        default:
            v50 := GenS2{F0: 813, F1: 128}
            _ = v50
        }
    default:
        v11 = 18446744073709551615
    }
    var v51 int16 = 1
    v6[-9223372036854775808] = ("" + ("hello" + ((((("" + "a") + "hello") + (("\n\t" + "\n\t") + "hello")) + "foo bar") + ("" + "hello"))))
    var v52 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}
