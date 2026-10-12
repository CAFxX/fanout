package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t"}
    var v3 uint64 = 493
    var v4 int32 = 318
    v3 = 1
    _ = len(v1)
    var v5 uint16 = 2
    if (!false) {
        v4 -= 1
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 472, F1: 2}
            for fi8 := 0; fi8 < 20; fi8++ {
                if true {
                    switch "a" {
                    case "hello":
                        v5 += 255
                    case "\n\t":
                        var v9 uint = 255
                        _ = v9
                    }
                    var v10 int32 = v4
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v10
                }
                v11 := GenS2{F0: -1}
                _ = fi8
                _ = v11
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                switch fi6 {
                case 457:
                    fmt.Println(v5)
                }
                if ((!((v5 | 42) > 624)) || false) {
                    switch "a" {
                    case "\n\t":
                        v3 = v3
                    case "hello":
                        v5 = 755
                    default:
                        v13 := GenS2{F0: 100}
                        _ = v13
                    }
                    var v14 int64 = 0
                    _ = v14
                }
                _ = fi12
            }
            _ = fi6
            _ = v7
        }
        if (!false) {
            v3 = v3
            v1[((((((("" + "a") + "") + (("foo bar" + "") + "foo bar")) + ("foo bar" + "foo bar")) + ("a" + "hello")) + (((("foo bar" + ("\n\t" + "hello")) + "\n\t") + "a") + "a")) + "hello")] = -342
            var v15 uint32 = 251
            _ = v15
        }
    } else {
        var v16 int = 4294967295
        _ = v16
    }
    v4 += v4
    fmt.Println(v3, v5, v4)
    for fi17 := 0; fi17 < 6; fi17++ {
        if false {
            v18 := GenS3{F0: 1, F1: 65535}
            v4 = 382
            if (!(593 != -15)) {
                fmt.Println(v3, fi17)
                v19 := GenS1{F0: v3, F1: 175}
                switch int(v5) {
                case -9223372036854775808:
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v3 -= v3
                        if false {
                            if false {
                                v21 := GenS4{F0: v4, F1: false, F2: 65535}
                                _ = v21
                            } else {
                                v4 -= v4
                                v3 += v3
                                var v22 uint16 = 42
                                _ = v22
                            }
                        }
                        fmt.Println(v5, fi17, v5)
                        if false {
                            v23 := GenS5{F0: int(v3), F1: -563, F2: -9223372036854775808}
                            var v24 int16 = -32768
                            _ = v23
                            _ = v24
                        } else {
                            v4 = (v4 / ((v4) | 1))
                            var v25 uint64 = v3
                            v4 -= (v4 * 779)
                            switch (fi20 ^ (^(fi20))) {
                            case -205:
                                v1["hello"] = 32767
                            case 1:
                                var v26 int16 = 32767
                                _ = v26
                            default:
                                if false {
                                    if (false || ((937 != 65535) && false)) {
                                        fmt.Println(fi20, v5)
                                        v27 := GenS6{F0: v3}
                                        _ = v27
                                    } else {
                                        var v28 int32 = -466
                                        fmt.Println(v28, v5)
                                        _ = v28
                                    }
                                }
                            }
                            var v29 int = 0
                            var v30 uint = 144
                            var v31 int64 = 1
                            v32 := GenS5{F0: (v29 << (((v30 / (((^(v30))) | 1))) & 63)), F1: v31, F2: v31}
                            fmt.Println(v4, v5, v25)
                            v31 += (-(v31))
                            _ = v25
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = fi20
                    }
                default:
                    v33 := GenS7{F0: ((((("hello" + "foo bar") + ((("" + ("\n\t" + "foo bar")) + ("foo bar" + "a")) + "")) + "foo bar") + "\n\t") + "\n\t"), F1: (("foo bar" + "\n\t") + ""), F2: (fi17 + 558)}
                    _ = v33
                }
                _ = v19
            }
            _ = v18
        }
        if true {
            var v34 int32 = (-(v4))
            for fi35 := 0; fi35 < 10; fi35++ {
                fmt.Println(v4, fi17, v4)
                fmt.Println(v3, fi35)
                fmt.Println(v5, fi35)
                _ = fi35
            }
            _ = v34
        } else {
            v36 := v1[("foo bar" + "a")]
            var v37 int32 = 42
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = v36
            _ = v37
        }
        _ = fi17
    }
    if ((!((1.5 >= 3.14159) && true)) || ((1 > 0) || false)) {
        var v38 string = ""
        switch "hello" {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v5 = v5
        }
        _ = v38
    }
    var v39 uint16 = (^(v5))
    var v40 uint = 18446744073709551615
    v40 -= v40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
