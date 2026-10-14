package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int64 }

func main() {
    var v1 string = "hello"
    var v2 float32 = 1e10
    v3 := []uint32{1, 2147483648}
    if (!(true || (!true))) {
        v4 := GenS1{F0: 127}
        _ = v4
    } else {
        var v5 float32 = (v2 * (1.5 / -2.25))
        fmt.Println(v5)
        v3 = append(v3, 42)
        _ = v5
    }
    v3 = append(v3, 401)
    switch -1 {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 16; fi6++ {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi6
        }
    case -9223372036854775808:
        var v7 int8 = -81
        _ = v7
    default:
        var v8 int16 = 479
        _ = v8
    }
    var v9 string = v1
    var v10 string = (v1 + ((v1 + ((v9 + v1) + (v1 + v1))) + (((v1 + v1) + "foo bar") + "a")))
    switch -476 {
    case 127:
        v2 += 3.14159
    case -556:
        v2 += 0.0
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v11 := GenS1{F0: 9223372036854775807}
    v12 := GenS1{F0: 671}
    for fi13 := 0; fi13 < 6; fi13++ {
        v9 = v1
        fmt.Println(v9, v10)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    v10 = ("a" + v1)
    v1 = "a"
    var v14 int32 = -2147483648
    if (418 <= -490) {
        var v15 uint = 1
        v16 := GenS1{F0: -102}
        _ = v15
        _ = v16
    } else {
        if true {
            fmt.Println(v10, v14, v10)
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            switch -39 {
            case 42:
                v10 = v9
            case -2147483648:
                for fi18 := 0; fi18 < 8; fi18++ {
                    if false {
                        v9 = (v10 + ((((v9 + (v10 + "hello")) + v10) + v9) + v10))
                    } else {
                        var v19 uint8 = 255
                        if false {
                            v20 := GenS2{F0: 1000, F1: "a", F2: 646}
                            var v21 float32 = v2
                            v21 -= v2
                            _ = v20
                            _ = v21
                        } else {
                            if true {
                                if false {
                                    v3 = append(v3, 0)
                                } else {
                                    fmt.Println(v19, v19, v19)
                                }
                                if false {
                                    v22 := GenS3{F0: 2147483647, F1: 2}
                                    switch "foo bar" {
                                    case "foo bar":
                                        v23 := GenS4{F0: 4294967295, F1: 18446744073709551615}
                                        _ = v23
                                    case "hello":
                                        if false {
                                            var v24 uint16 = 255
                                            var v25 uint32 = 581
                                            fmt.Println(v9, v19, v2)
                                            _ = v24
                                            _ = v25
                                        }
                                    case "":
                                        var v26 float32 = float32(1e10)
                                        _ = v26
                                    }
                                    _ = v22
                                }
                            }
                            fmt.Println(fi17, fi18)
                        }
                        var v27 bool = false
                        _ = v19
                        _ = v27
                    }
                    _ = fi18
                }
            case 545:
                v28 := GenS5{F0: 0}
                _ = v28
            }
            if false {
                var v29 bool = (1 < 0)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v29
            } else {
                var v30 int32 = -2147483648
                var v31 uint8 = 255
                _ = v30
                _ = v31
            }
            _ = fi17
        }
    }
    var v32 string = "\n\t"
    switch -1 {
    case -799:
        if false {
            var v33 int8 = -1
            v34 := GenS2{F0: -1, F1: v32, F2: 222}
            switch -9223372036854775808 {
            case 2147483647:
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if (!false) {
                        switch "" {
                        case "\n\t":
                            for fi35 := 0; fi35 < 2; fi35++ {
                                var v36 float32 = (float32(0.1) - 1.5)
                                v36 = 3.14159
                                _ = fi35
                                _ = v36
                            }
                        case "":
                            var v37 float64 = (0.1 + ((0.0 * float64((((float32(1.5) / v2) + v2) / v2))) / -2.25))
                            _ = v37
                        default:
                            v1 = "foo bar"
                        }
                    } else {
                        var v38 uint16 = 300
                        v1 = (v9 + v1)
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 12; fi39++ {
                        var v40 int8 = (v33 >> ((2) & 63))
                        _ = fi39
                        _ = v40
                    }
                    v10 = (("\n\t" + (v10 + "foo bar")) + ("hello" + (("" + v32) + "foo bar")))
                }
            case -1:
                v9 = "a"
            default:
                var v41 string = v32
                _ = v41
            }
            _ = v33
            _ = v34
        } else {
            if false {
                for fi42 := 0; fi42 < 6; fi42++ {
                    var v43 float64 = 0.0
                    v44 := GenS1{F0: fi42}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                var v45 uint16 = 0
                if true {
                    v46 := GenS6{F0: 1, F1: 9223372036854775807}
                    _ = v46
                } else {
                    fmt.Println(v14)
                    var v47 float32 = v2
                    _ = v47
                }
                _ = v45
            } else {
                for fi48 := 0; fi48 < 16; fi48++ {
                    v9 = ("" + v10)
                    switch "hello" {
                    case "\n\t":
                        v14 += 2147483647
                    case "a":
                        if (1 <= 4294967295) {
                            var v49 uint8 = 1
                            _ = v49
                        } else {
                            var v50 uint16 = 32767
                            _ = v50
                        }
                    default:
                        var v51 int16 = 0
                        _ = v51
                    }
                    _ = fi48
                }
            }
        }
    case -427:
        fmt.Println(v1)
    case 42:
        var v52 float64 = (-2.25 + -2.25)
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v9
}
