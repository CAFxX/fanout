package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint; F2 uint8 }

func main() {
    var v1 uint8 = 255
    var v2 int16 = 536
    var v3 uint = 785
    v4 := make(map[string]uint32)
    var v5 int8 = -117
    v6 := []string{(((("hello" + "\n\t") + ("hello" + ("a" + "\n\t"))) + "") + ("\n\t" + "\n\t")), "foo bar", "\n\t", "hello"}
    v7 := GenS1{F0: v1}
    v8 := GenS2{F0: -9223372036854775808}
    var v9 uint = (v3 << ((v3) & 63))
    switch "hello" {
    case "hello":
        if ((((1e10 - (3.14159 * (((0.0 / 1.5) / 1.0) - (0.1 / 0.1)))) * -2.25) / ((3.14159 + 0.0) / 1.5)) > float32(-2.25)) {
            v1 -= 2
        }
    case "\n\t":
        v1 += 127
    case "a":
        var v10 float64 = -2.25
        _ = v10
    default:
        v5 -= 42
    }
    var v11 int32 = 1
    if len(v6) > 0 {
        v6[0] = ((("foo bar" + "hello") + ("" + ((("a" + ((("hello" + "\n\t") + "a") + "")) + "") + "a"))) + "hello")
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        _ = len(v6)
        if (!true) {
            v5 -= int8(v3)
            v13 := GenS1{F0: (v1 << ((v3) & 63))}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
                _ = fi14
            }
            v11 = 1
        }
        _ = fi12
    }
    var v15 uint16 = 32767
    for fi16 := 0; fi16 < 10; fi16++ {
        v2 = v2
        _ = fi16
    }
    v17 := GenS3{F0: v9, F1: v3, F2: v15}
    v1 += (^(v1))
    if false {
        var v18 string = "hello"
        v1 += v1
        switch 448 {
        case 512:
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint16 = (v15 - 62)
                var v21 uint64 = 2
                _ = fi19
                _ = v20
                _ = v21
            }
        case 643:
            switch 458 {
            case 2147483647:
                if ((v9 | (v3 >> ((v9) & 63))) >= (v9 | (v9 ^ 277))) {
                    fmt.Println(v15, v5)
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 uint8 = 31
                        if (1e10 == float64(0.0)) {
                            if (true && (float64(1e10) == (3.14159 - 0.1))) {
                                var v24 uint8 = v23
                                v23 = 34
                                if true {
                                    switch (fi22 / ((127) | 1)) {
                                    case 1:
                                        v25 := GenS4{F0: fi22}
                                        _ = v25
                                    case 442:
                                        var v26 uint16 = v15
                                        _ = v26
                                    }
                                    var v27 int64 = -151
                                    _ = v27
                                } else {
                                    fmt.Println(v15)
                                    v4["\n\t"] = 690
                                }
                                _ = v24
                            } else {
                                var v28 int64 = 130
                                var v29 uint32 = uint32(v23)
                                _ = v28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 20; fi30++ {
                                if false {
                                    v31 := GenS5{F0: (v11 / (((^(v11))) | 1))}
                                    v2 = 939
                                    v32 := GenS6{F0: 0, F1: 18446744073709551615, F2: 128}
                                    _ = v31
                                    _ = v32
                                }
                                _ = fi30
                            }
                        } else {
                            v18 = (("a" + "a") + v18)
                            var v33 uint8 = 127
                            _ = v33
                        }
                        _ = fi22
                        _ = v23
                    }
                    if (float32((3.14159 * 0.0)) != (((-2.25 + 0.0) / 1e10) + ((1e10 - 0.0) - ((1e10 + float32(0.0)) / 3.14159)))) {
                        v9 = (v3 | 18446744073709551615)
                    } else {
                        var v34 string = "\n\t"
                        _ = v34
                    }
                } else {
                    var v35 uint32 = 2147483647
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v5 = -128
                        v37 := GenS4{F0: (fi36 * (fi36 * fi36))}
                        for fi38 := 0; fi38 < 18; fi38++ {
                            switch fi38 {
                            case 1:
                                v35 -= v35
                            case -9223372036854775808:
                                var v39 int32 = int32(v5)
                                _ = v39
                            default:
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    if (!(!false)) {
                                        if true {
                                            var v40 uint8 = (v1 + 1)
                                            var v41 bool = false
                                            _ = v40
                                            _ = v41
                                        } else {
                                            var v42 float64 = (3.14159 + (3.14159 + 1.5))
                                            v9 = 4294967295
                                            _ = v42
                                        }
                                        v43 := v4["hello"]
                                        var v44 int64 = -350
                                        v45 := GenS4{F0: (fi36 | (fi36 - fi38))}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v11 -= (v11 & (v11 & (^(v11))))
                                        switch v18 {
                                        case "foo bar":
                                            fmt.Println(v9, v9)
                                        case "\n\t":
                                            var v46 int32 = (v11 * (v11 / ((v11) | 1)))
                                            _ = v46
                                        case "a":
                                            switch int(v2) {
                                            case 32767:
                                                v1 -= 1
                                            case -238:
                                                var v47 int32 = v11
                                                v9 = 2
                                                switch -613 {
                                                case 2147483647:
                                                    fmt.Println(v3)
                                                case 127:
                                                    var v48 uint16 = v15
                                                    _ = v48
                                                }
                                                _ = v47
                                            }
                                        }
                                        fmt.Println(v18)
                                    }
                                    v49 := GenS5{F0: (v11 ^ (^(v11)))}
                                    _ = v49
                                case "hello":
                                    _ = len(v4)
                                case "foo bar":
                                    var v50 int = (fi36 / ((128) | 1))
                                    _ = v50
                                default:
                                    var v51 int64 = int64(v5)
                                    _ = v51
                                }
                            }
                            _ = fi38
                        }
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            case 127:
                var v52 int8 = 29
                _ = v52
            }
        case 368:
            fmt.Println(v2, v9, v5)
        default:
            var v53 uint8 = v1
            _ = v53
        }
        _ = v18
    }
    var v54 int = 54
    var v55 int32 = v11
    for fi56 := 0; fi56 < 18; fi56++ {
        for fi57 := 0; fi57 < 3; fi57++ {
            v6 = append(v6, (("a" + ("\n\t" + "hello")) + "hello"))
            _ = fi57
        }
        _ = fi56
    }
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    var v58 uint = v3
    v5 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v58
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
