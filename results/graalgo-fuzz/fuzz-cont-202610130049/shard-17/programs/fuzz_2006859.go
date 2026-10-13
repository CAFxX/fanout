package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int; F2 int8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 bool }

func main() {
    v1 := []string{"a", "", "hello", ("foo bar" + ("foo bar" + "foo bar"))}
    var v2 uint8 = 221
    v3 := make(map[string]bool)
    if false {
        v4 := GenS1{F0: 76, F1: ((((("" + "a") + ("foo bar" + ("a" + "foo bar"))) + (("\n\t" + "hello") + "a")) + ("" + "a")) + ((("hello" + (("" + "foo bar") + "")) + ("hello" + ("foo bar" + "hello"))) + "a"))}
        var v5 int32 = 2147483647
        _ = v4
        _ = v5
    } else {
        delete(v3, "foo bar")
        v6 := GenS1{F0: 127, F1: ("foo bar" + (("hello" + "a") + ("" + "a")))}
        _ = v6
    }
    var v7 uint32 = 527
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint = 250
        v9 += 4294967295
        var v10 uint8 = v2
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 bool = ((("a" + "\n\t") != "") || true)
        delete(v3, "foo bar")
        if v12 {
            var v13 int64 = 1
            v14 := GenS1{F0: int8(v7), F1: "foo bar"}
            for fi15 := 0; fi15 < 13; fi15++ {
                if (((!(v12 && true)) || (!(3.14159 >= 1e10))) || (v7 >= (^(v7)))) {
                    if (!(3.14159 > (-2.25 / 1.0))) {
                        v16 := GenS2{F0: -1, F1: 1}
                        if ((fi15 * 42) <= (fi15 << ((0) & 63))) {
                            delete(v3, "a")
                            if v12 {
                                v17 := GenS3{F0: -128}
                                _ = v17
                            }
                            v18 := v3["hello"]
                            _ = v18
                        }
                        _ = v16
                    } else {
                        if v12 {
                            switch fi11 {
                            case -9223372036854775808:
                                v19 := GenS4{F0: 1}
                                _ = v19
                            case 127:
                                v20 := GenS5{F0: -58, F1: fi11, F2: 24}
                                _ = v20
                            case 9223372036854775807:
                                if (0 == 2) {
                                    v21 := GenS2{F0: 42, F1: 1}
                                    _ = v21
                                } else {
                                    _ = len(v3)
                                    switch (fi15 & fi11) {
                                    case -194:
                                        var v22 uint8 = v2
                                        _ = v22
                                    default:
                                        v7 += (v7 / ((v7) | 1))
                                    }
                                    v12 = (fi11 >= fi15)
                                }
                            }
                        }
                        var v23 int16 = 716
                        var v24 string = "foo bar"
                        _ = v23
                        _ = v24
                    }
                    v7 += 84
                } else {
                    if v12 {
                        var v25 uint16 = 255
                        switch fi11 {
                        case 127:
                            switch (fi15 & -1) {
                            case 128:
                                if (!(!(false || (!(!((v2 >= 128) || v12)))))) {
                                    var v26 uint32 = (v7 | v7)
                                    v27 := GenS6{F0: (v13 + (v13 / ((v13) | 1))), F1: (664 >= -531)}
                                    var v28 int8 = 1
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        case 2147483647:
                            _ = len(v3)
                        }
                        _ = v25
                    }
                }
                v1 = append(v1, "\n\t")
                v2 = (v2 & v2)
                _ = fi15
            }
            _ = v13
            _ = v14
        } else {
            var v29 bool = false
            _ = v29
        }
        _ = fi11
        _ = v12
    }
    if (false || false) {
        var v30 int8 = 42
        _ = v30
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v7 = v7
    case "\n\t":
        var v31 float32 = 0.1
        _ = v31
    default:
        v7 = 4294967295
    }
    v32 := v3[""]
    switch 722 {
    case -2147483648:
        var v33 int64 = 9223372036854775807
        _ = v33
    case -1:
        if ((float32((1.5 / ((((float64(1e10) + 1e10) + -2.25) + float64((1e10 - 3.14159))) - float64((float32(0.0) * 0.0))))) / (float32(1e10) / 0.1)) < (0.0 * (1.5 / 1.0))) {
            fmt.Println(v32)
            switch "" {
            case "\n\t":
                v1 = append(v1, ("" + "hello"))
            default:
                v2 -= (v2 | v2)
            }
        } else {
            if (0 > 65535) {
                v3[("foo bar" + "\n\t")] = v32
            }
            var v34 int = -721
            var v35 uint32 = (v7 % ((v7) | 1))
            _ = v34
            _ = v35
        }
    case 42:
        v32 = v32
    default:
        var v36 uint8 = v2
        _ = v36
    }
    fmt.Println(v32, v7)
    var v37 int8 = 1
    var v38 int8 = (v37 * 100)
    switch 920 {
    case 4294967295:
        switch "hello" {
        case "a":
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := GenS7{F0: 0}
                var v41 float32 = float32(1.5)
                var v42 uint64 = uint64(v2)
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        case "\n\t":
            if true {
                fmt.Println(v38)
            }
        }
    case 731:
        for fi43 := 0; fi43 < 9; fi43++ {
            fmt.Println(v7, v32)
            v44 := GenS2{F0: 496, F1: 565}
            _ = fi43
            _ = v44
        }
    case -2147483648:
        v45 := GenS8{F0: 705, F1: false}
        _ = v45
    default:
        for fi46 := 0; fi46 < 4; fi46++ {
            fmt.Println(v38, v32)
            _ = fi46
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v7
}
