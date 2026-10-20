package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"hello"}
    var v3 int = 1
    var v4 int64 = 9223372036854775807
    var v5 float32 = (1.5 / 1.5)
    v6 := make(map[uint32]int)
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: 82, F1: -57}
        _ = v7
    case "foo bar":
        var v8 uint = 255
        _ = v8
    }
    v4 = 337
    delete(v6, 2)
    v9 := GenS2{F0: -375, F1: (v4 | 107)}
    switch "a" {
    case "\n\t":
        v10 := GenS3{F0: 32768}
        _ = v10
    case "":
        _ = len(v1)
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = ("\n\t" + ("foo bar" + ""))
        }
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        if true {
            var v11 uint = 636
            var v12 uint64 = 0
            var v13 float64 = 0.0
            _ = v11
            _ = v12
            _ = v13
        }
    case "":
        v4 -= 9223372036854775807
    }
    var v14 int64 = v4
    fmt.Println(v14, v3, v14)
    var v15 int32 = 753
    switch "hello" {
    case "":
        var v16 int32 = -891
        _ = v16
    case "foo bar":
        if (0 == 495) {
            v3 += 128
            v4 = v4
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v4, v5, v3)
                v5 = v5
                var v18 uint32 = 371
                _ = fi17
                _ = v18
            }
        }
    default:
        switch v3 {
        case 0:
            if (1 > 255) {
                v15 += 2147483647
            }
        case 128:
            v19 := GenS4{F0: (-(v15)), F1: 960, F2: 4294967295}
            _ = v19
        case 2147483647:
            switch -821 {
            case 2147483647:
                switch v3 {
                case -569:
                    if (1 == 42) {
                        var v20 int32 = -2147483648
                        for fi21 := 0; fi21 < 11; fi21++ {
                            _ = len(v2)
                            v20 = 1
                            v14 = (v14 ^ v4)
                            _ = fi21
                        }
                        _ = v20
                    }
                case -1:
                    if true {
                        _ = len(v1)
                    } else {
                        v14 += (v14 | (v4 | v14))
                        if false {
                            _ = len(v2)
                        } else {
                            if true {
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    v4 -= (v4 | 1)
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        v24 := GenS2{F0: -1, F1: -352}
                                        v15 += v15
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = fi22
                                }
                                v15 = v15
                                if (180 > 42) {
                                    var v25 uint = 255
                                    v5 = v5
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: (^(v15)), F1: v4}
                                    for fi27 := 0; fi27 < 17; fi27++ {
                                        var v28 bool = (false && (true && true))
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = v26
                                }
                            }
                        }
                        switch 42 {
                        case 32767:
                            if (!true) {
                                var v29 int16 = -81
                                v4 = 1
                                _ = v29
                            } else {
                                switch 32768 {
                                case 42:
                                    v5 -= v5
                                case 128:
                                    v15 += v15
                                case 9223372036854775807:
                                    var v30 bool = (42 <= 42)
                                    _ = v30
                                default:
                                    fmt.Println(v3, v4, v15)
                                }
                                v14 = -917
                            }
                        case 127:
                            v5 = v5
                        case 9223372036854775807:
                            if true {
                                if true {
                                    var v31 uint32 = 65535
                                    switch "" {
                                    case "a":
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                    case "hello":
                                        if (!false) {
                                            v1["a"] = (^(v31))
                                            for fi32 := 0; fi32 < 5; fi32++ {
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                _ = fi32
                                            }
                                        }
                                    case "":
                                        var v33 bool = (false || false)
                                        _ = v33
                                    }
                                    v31 = 570
                                    _ = v31
                                } else {
                                    fmt.Println(v4, v4, v4)
                                    v4 += v4
                                    v6[4294967295] = (v3 % ((v3) | 1))
                                }
                                v5 -= v5
                            } else {
                                v15 = v15
                            }
                        default:
                            v2 = append(v2, "foo bar")
                        }
                    }
                case -351:
                    if false {
                        var v34 string = "hello"
                        _ = v34
                    }
                default:
                    v5 = v5
                }
            default:
                for fi35 := 0; fi35 < 18; fi35++ {
                    var v36 int16 = -573
                    _ = fi35
                    _ = v36
                }
            }
        }
    }
    v37 := GenS5{F0: 65535, F1: 165}
    _ = len(v2)
    v4 -= (v4 | -9223372036854775808)
    var v38 int8 = 42
    v15 -= (v15 & int32(v38))
    v5 = v5
    var v39 uint = 4294967295
    var v40 bool = ((v39 + 1) <= v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
