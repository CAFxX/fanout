package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{32767}
    var v3 uint8 = 255
    v4 := []string{"foo bar", ""}
    var v5 uint64 = 1
    v3 += v3
    if (-626 >= 1) {
        var v6 int = -354
        v5 += (v5 % (((v5 & v5)) | 1))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                switch "a" {
                case "\n\t":
                    v5 = v5
                case "a":
                    switch int(v3) {
                    case 9223372036854775807:
                        v9 := GenS1{F0: uint64(v3), F1: -1, F2: 2147483648}
                        v3 = uint8(v5)
                        switch "hello" {
                        case "a":
                            var v10 uint32 = 2
                            _ = v10
                        }
                        _ = v9
                    case 345:
                        v5 += v5
                    default:
                        if false {
                            fmt.Println(fi7, fi7, fi8)
                        } else {
                            v3 -= 255
                            if true {
                                switch fi8 {
                                case 127:
                                    var v11 uint16 = 2
                                    _ = v11
                                }
                            } else {
                                fmt.Println(fi7, v3, fi8)
                                v3 += uint8(v5)
                                v3 = (v3 & 128)
                            }
                        }
                        delete(v1, ((((("foo bar" + "hello") + "hello") + "") + ("foo bar" + "a")) + "foo bar"))
                    }
                case "":
                    if len(v2) > 0 {
                        v2[0] = (-(fi7))
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v3, v5)
                    case "hello":
                        v3 -= (v3 + v3)
                    case "":
                        _ = len(v4)
                    default:
                        v5 += v5
                    }
                }
                v3 = (^(v3))
                if true {
                    switch 1 {
                    case -353:
                        if (int32(fi8) < 42) {
                            v3 -= (^(v3))
                        }
                    case 4294967295:
                        fmt.Println(fi8, fi8)
                    case 851:
                        switch fi8 {
                        case 4294967295:
                            if false {
                                var v12 bool = false
                                _ = v12
                            } else {
                                v2 = append(v2, fi8)
                                v13 := GenS2{F0: ("\n\t" + "foo bar")}
                                v3 += v3
                                switch (fi7 + (fi7 | (fi7 - fi7))) {
                                case 516:
                                    fmt.Println(fi8, v3, fi8)
                                default:
                                    var v14 int16 = -1
                                    _ = v14
                                }
                                _ = v13
                            }
                        case 9223372036854775807:
                            v15 := v1[(("foo bar" + "foo bar") + "\n\t")]
                            _ = v15
                        case -968:
                            switch fi8 {
                            case 127:
                                v16 := GenS3{F0: 1, F1: 4294967295}
                                _ = v16
                            }
                        }
                    }
                    v5 += v5
                } else {
                    var v17 uint32 = 2
                    _ = v17
                }
                _ = fi8
            }
            var v18 int8 = int8(v5)
            _ = fi7
            _ = v18
        }
        var v19 string = (("a" + ((("foo bar" + (("a" + "a") + ("" + "a"))) + "\n\t") + "")) + ("a" + (("" + "\n\t") + "\n\t")))
        v20 := GenS4{F0: -2147483648, F1: 42}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v3 -= (v3 >> ((845) & 63))
        for fi22 := 0; fi22 < 8; fi22++ {
            v3 = (v3 & 36)
            var v23 int8 = 0
            _ = fi22
            _ = v23
        }
        var v24 int32 = 0
        _ = fi21
        _ = v24
    }
    var v25 int = 83
    if (2147483647 > 0) {
        var v26 int64 = -500
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 bool = false
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    var v29 int16 = 32767
    var v30 int16 = 42
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint64 = v5
        _ = fi31
        _ = v32
    }
    var v33 int64 = -9223372036854775808
    if false {
        v33 = v33
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS2{F0: "a"}
            _ = fi34
            _ = v35
        }
    } else {
        switch "" {
        case "foo bar":
            v33 = -1
        case "":
            for fi36 := 0; fi36 < 3; fi36++ {
                if (0 == uint16(v5)) {
                    v30 += (v29 % (((v29 * (v29 * -32768))) | 1))
                    var v37 string = "\n\t"
                    _ = v37
                }
                _ = fi36
            }
        case "a":
            var v38 int8 = -128
            _ = v38
        }
        v5 -= v5
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
