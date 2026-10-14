package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 uint32 = 0
    v2 := []string{"a", "hello"}
    var v3 int8 = 14
    var v4 bool = (1 != 724)
    var v5 uint64 = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 := GenS1{F0: v5}
    _ = len(v2)
    v3 -= (v3 % ((v3) | 1))
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (!(!true)) {
                var v9 string = ("a" + "a")
                _ = v9
            }
            v5 = v5
            fmt.Println(fi8, v5, v1)
            switch "foo bar" {
            case "":
                if ((v5 / ((v5) | 1)) >= v5) {
                    v10 := GenS2{F0: 42, F1: (-(v3)), F2: 65535}
                    switch fi7 {
                    case -438:
                        if (v4 && (2147483647 == 1)) {
                            if (v5 < (^(v5))) {
                                var v11 bool = false
                                v2 = append(v2, "a")
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                _ = v11
                            } else {
                                fmt.Println(v3)
                            }
                            if (-81 > -128) {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v12 int8 = (v3 & 127)
                                fmt.Println(v12)
                                _ = v12
                            } else {
                                v4 = (477 <= 65535)
                                v2 = append(v2, ("\n\t" + "hello"))
                                if v4 {
                                    if ((!(true && (!true))) || v4) {
                                        switch 0 {
                                        case -9223372036854775808:
                                            if len(v2) > 0 {
                                                v2[0] = ("\n\t" + "foo bar")
                                            }
                                        case -499:
                                            switch "" {
                                            case "hello":
                                                v2 = append(v2, ("foo bar" + "a"))
                                            case "a":
                                                v4 = v4
                                                _ = len(v2)
                                                if v4 {
                                                    var v13 float64 = float64(((3.14159 / 0.1) - 3.14159))
                                                    if len(v2) > 0 {
                                                        v2[0] = ("" + ("foo bar" + ""))
                                                    }
                                                    switch -521 {
                                                    case 4294967295:
                                                        if true {
                                                            v14 := GenS1{F0: (v5 & v5)}
                                                            if (fi7 == (fi8 & fi7)) {
                                                                var v15 uint16 = 42
                                                                v4 = true
                                                                switch (("foo bar" + "\n\t") + "\n\t") {
                                                                case "hello":
                                                                    var v16 bool = false
                                                                    _ = v16
                                                                case "\n\t":
                                                                    v4 = (v4 && (55 != v1))
                                                                case "foo bar":
                                                                    fmt.Println(v13, v15)
                                                                    fmt.Println(fi7, fi8, v3)
                                                                    v4 = (!(v13 >= v13))
                                                                }
                                                                var v17 bool = ((v13 == (float64((0.0 * 0.0)) - 0.0)) && v4)
                                                                if false {
                                                                    var v18 uint = 42
                                                                    _ = v18
                                                                }
                                                                _ = v15
                                                                _ = v17
                                                            }
                                                            v2 = append(v2, "hello")
                                                            _ = v14
                                                        }
                                                        var v19 bool = (4294967295 >= 408)
                                                        _ = v19
                                                    case 528:
                                                        v20 := GenS2{F0: 1, F1: v3, F2: 65535}
                                                        _ = v20
                                                    case 42:
                                                        var v21 float32 = -2.25
                                                        _ = v21
                                                    }
                                                    _ = v13
                                                }
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        default:
                                            v2 = append(v2, "foo bar")
                                        }
                                        fmt.Println(v4, fi7, v3)
                                        v2 = append(v2, ("" + "foo bar"))
                                    }
                                    if (127 != 0) {
                                        var v22 uint64 = (v5 - v5)
                                        _ = v22
                                    } else {
                                        var v23 uint8 = 2
                                        fmt.Println(v3)
                                        v24 := GenS1{F0: v5}
                                        _ = v23
                                        _ = v24
                                    }
                                }
                            }
                        }
                    case -127:
                        v25 := GenS3{F0: "hello", F1: (65535 >= 1)}
                        _ = v25
                    default:
                        if (("\n\t" + ("\n\t" + "a")) != (((("a" + (("foo bar" + "") + ("" + "\n\t"))) + ("" + "foo bar")) + "\n\t") + "hello")) {
                            v1 += (v1 % (((v1 >> ((178) & 63))) | 1))
                            var v26 float32 = 0.0
                            _ = v26
                        }
                    }
                    v5 = (^(v5))
                    _ = v10
                }
            case "hello":
                var v27 bool = v4
                _ = v27
            case "\n\t":
                v28 := GenS3{F0: ("foo bar" + ""), F1: v4}
                _ = v28
            default:
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "foo bar")
                }
            }
            _ = fi8
        }
        _ = fi7
    }
    if v4 {
        fmt.Println(v5, v4, v5)
        var v29 int = 371
        _ = v29
    }
    var v30 uint16 = 42
    if v4 {
        var v31 int = -1
        _ = len(v2)
        var v32 uint8 = 255
        _ = v31
        _ = v32
    }
    switch 127 {
    case 42:
        fmt.Println(v5, v4)
    case 128:
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := GenS1{F0: (v5 * 871)}
            fmt.Println(v4, v3)
            _ = fi33
            _ = v34
        }
    case -646:
        if ((float32(3.14159) * ((0.1 + (float32(-2.25) + 0.1)) * 1.5)) < float32((1.5 / 1.0))) {
            v35 := GenS1{F0: v5}
            var v36 int16 = -32768
            _ = v35
            _ = v36
        } else {
            var v37 int16 = int16(v5)
            _ = v37
        }
    default:
        for fi38 := 0; fi38 < 19; fi38++ {
            if v4 {
                v5 = 18446744073709551615
                v30 = 255
                v5 += (v5 + (v5 % (((^(v5))) | 1)))
            } else {
                var v39 string = ("" + "\n\t")
                v4 = ((!v4) && v4)
                _ = v39
            }
            _ = fi38
        }
    }
    var v40 uint16 = v30
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
