package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 int; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint; F2 int }
type GenS4 struct { F0 int8; F1 string }
func genFn0() uint32 {
    switch 32768 {
    case 128:
        var v4 uint8 = 127
        _ = v4
    case 9223372036854775807:
        v5 := GenS1{F0: 2, F1: 1, F2: 32767}
        _ = v5
    default:
        var v6 int = 42
        _ = v6
    }
    v7 := []int{-9223372036854775808, 581}
    switch ("hello" + ("foo bar" + ("" + "foo bar"))) {
    case "":
        if true {
            var v8 uint = 676
            _ = v8
        }
    default:
        if (-1 == 0) {
            v9 := GenS1{F0: 864, F1: 586, F2: 500}
            _ = v9
        } else {
            v10 := GenS2{F0: "", F1: 2147483647, F2: 1}
            _ = v10
        }
    }
    var v11 uint16 = 0
    switch 2147483647 {
    case 108:
        v7 = append(v7, 42)
    }
    _ = v11
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483648, 42}
    v3 := make(map[string]string)
    var v12 uint16 = 2
    switch ("" + "foo bar") {
    case "\n\t":
        switch 32767 {
        case -9223372036854775808:
            var v13 int8 = 127
            _ = v13
        }
    case "foo bar":
        if false {
            _ = len(v2)
            fmt.Println(v12)
            v14 := v3[("" + "")]
            _ = v14
        }
    case "":
        if false {
            v12 -= (^(v12))
            delete(v3, (((("hello" + "hello") + "a") + "foo bar") + "hello"))
            for fi15 := 0; fi15 < 15; fi15++ {
                if false {
                    _ = len(v1)
                } else {
                    v16 := GenS3{F0: 255, F1: 322, F2: fi15}
                    var v17 int16 = 42
                    _ = v16
                    _ = v17
                }
                v18 := GenS4{F0: -125, F1: "foo bar"}
                _ = fi15
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 string = "foo bar"
            if false {
                var v22 float32 = 0.0
                delete(v1, 679)
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(v12)
        _ = len(v2)
        _ = fi19
    }
    if false {
        v12 -= (v12 % ((v12) | 1))
        fmt.Println(v12)
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint = 776
        var v25 uint = v24
        switch fi23 {
        case 32767:
            if true {
                for fi26 := 0; fi26 < 18; fi26++ {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    if (uint(v12) != v25) {
                        v25 -= 255
                        if (!(255 < 2)) {
                            fmt.Println(v12, fi26)
                        }
                    } else {
                        switch (fi26 >> ((uint(v12)) & 63)) {
                        case 128:
                            v12 -= 0
                            if false {
                                var v27 uint8 = 188
                                _ = v27
                            } else {
                                switch "" {
                                case "\n\t":
                                    v3["\n\t"] = ("a" + "hello")
                                default:
                                    if (1 > int64(v24)) {
                                        var v28 uint = 18446744073709551615
                                        var v29 bool = (((v12 & (v12 << ((v28) & 63))) < 666) || (int64(v12) > -81))
                                        v25 = 255
                                        _ = v28
                                        _ = v29
                                    } else {
                                        v12 += 0
                                        v12 = (v12 % (((v12 >> ((v24) & 63))) | 1))
                                    }
                                }
                            }
                            v30 := GenS3{F0: 127, F1: 1, F2: fi26}
                            v12 -= 21
                            switch 42 {
                            case 100:
                                var v31 uint32 = genFn0()
                                var v32 uint64 = uint64(v25)
                                var v33 int8 = 127
                                v12 -= 593
                                v32 += 42
                                _ = v31
                                _ = v32
                                _ = v33
                            case -2147483648:
                                v34 := GenS1{F0: 811, F1: -2147483648, F2: 2147483647}
                                _ = v34
                            case 635:
                                fmt.Println(v25, fi23, v25)
                            }
                            _ = v30
                        case 127:
                            fmt.Println(fi26, v24, v25)
                        case 32768:
                            if false {
                                v12 -= 218
                                if len(v2) > 0 {
                                    v2[0] = 2147483648
                                }
                            } else {
                                fmt.Println(v12)
                                v25 = (^(v24))
                            }
                        default:
                            var v35 float32 = -2.25
                            _ = v35
                        }
                        v25 -= v25
                        var v36 uint32 = 713
                        _ = v36
                    }
                    v25 = v25
                    _ = fi26
                }
                for fi37 := 0; fi37 < 15; fi37++ {
                    if false {
                        v12 = (v12 - (v12 / (((v12 >> ((v24) & 63))) | 1)))
                        var v38 bool = (255 != v24)
                        var v39 uint8 = 1
                        _ = v38
                        _ = v39
                    } else {
                        fmt.Println(v24)
                        var v40 string = ("\n\t" + "hello")
                        fmt.Println(v40, v24, v24)
                        _ = v40
                    }
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        var v41 uint8 = 2
                        _ = v41
                    case "a":
                        v42 := GenS4{F0: 0, F1: ("hello" + ("\n\t" + "hello"))}
                        var v43 uint32 = genFn0()
                        _ = v42
                        _ = v43
                    }
                    _ = fi37
                }
                fmt.Println(v25)
            }
        case 128:
            v1[646] = 2
        case 0:
            var v44 uint64 = 2
            _ = v44
        default:
            if ((^(fi23)) < (fi23 ^ 507)) {
                var v45 int32 = 42
                switch "foo bar" {
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
                _ = v45
            }
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v12)
    var v46 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v46
}
