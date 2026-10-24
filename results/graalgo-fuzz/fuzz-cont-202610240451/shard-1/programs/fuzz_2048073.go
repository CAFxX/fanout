package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    var v3 uint8 = 182
    var v4 bool = (false || false)
    var v5 uint64 = 0
    fmt.Println(v1)
    switch 127 {
    case 32768:
        switch -253 {
        case 1:
            v5 = (v5 - v5)
        case -1:
            var v6 bool = v4
            _ = v6
        }
    case 9223372036854775807:
        v3 += 42
    }
    switch 98 {
    case 2147483647:
        fmt.Println(v5)
    }
    var v7 uint16 = 0
    v8 := []int{}
    var v9 int32 = -2147483648
    if ("hello" == "a") {
        v10 := GenS1{F0: v5, F1: 0}
        _ = v10
    } else {
        v9 = 557
    }
    v1 += v1
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float64 = -2.25
        if v4 {
            var v13 int8 = -85
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint64 = (v5 ^ 356)
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            var v16 uint = 255
            var v17 string = ""
            var v18 float32 = 3.14159
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    var v19 int64 = -9223372036854775808
    fmt.Println(v1, v4, v4)
    var v20 bool = (!(267 > v7))
    v4 = ((float64(1e10) <= (1e10 / 1.0)) || false)
    switch 32768 {
    case 468:
        v21 := GenS1{F0: (v5 | v5), F1: 2}
        _ = v21
    }
    switch ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + ("a" + "a"))) {
    case "\n\t":
        var v22 string = ("\n\t" + "hello")
        _ = v22
    case "a":
        switch 127 {
        case 42:
            var v23 int32 = 85
            _ = v23
        case 128:
            switch "hello" {
            case "\n\t":
                if len(v8) > 0 {
                    v8[0] = -9223372036854775808
                }
            case "a":
                var v24 uint8 = (v3 / (((v3 & (^(v3)))) | 1))
                _ = v24
            case "foo bar":
                v25 := GenS2{F0: 928, F1: v19}
                _ = v25
            default:
                if ((!v4) && (v7 != v7)) {
                    fmt.Println(v1, v1, v19)
                    fmt.Println(v19)
                    if v20 {
                        switch 2147483647 {
                        case -939:
                            if false {
                                var v26 float32 = 0.0
                                v3 += (v3 * uint8(v7))
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v28 := GenS2{F0: 32767, F1: (v19 % (((v19 - v19)) | 1))}
                                        _ = v28
                                    case "foo bar":
                                        v9 += v9
                                    default:
                                        fmt.Println(fi27, v19, v5)
                                    }
                                    _ = fi27
                                }
                                _ = v26
                            }
                        case 2147483647:
                            v1 -= (1e10 / (v1 - 0.1))
                        case -9223372036854775808:
                            switch (("" + "") + ("foo bar" + "hello")) {
                            case "\n\t":
                                if v4 {
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        fmt.Println(fi29)
                                        _ = len(v2)
                                        _ = fi29
                                    }
                                    v2[(("\n\t" + "hello") + "")] = -878
                                    fmt.Println(v20, v20)
                                } else {
                                    v7 -= 823
                                    var v30 int32 = (v9 % ((v9) | 1))
                                    _ = v30
                                }
                            default:
                                fmt.Println(v20, v1)
                            }
                        default:
                            v31 := GenS3{F0: -913, F1: -1}
                            _ = v31
                        }
                        var v32 uint32 = 65535
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            for fi34 := 0; fi34 < 20; fi34++ {
                                _ = len(v8)
                                var v35 uint16 = (v7 + (v7 >> ((1) & 63)))
                                _ = fi34
                                _ = v35
                            }
                            _ = fi33
                        }
                        v36 := GenS1{F0: 591, F1: 65535}
                        v9 += 433
                        _ = v36
                    }
                } else {
                    if v4 {
                        v37 := GenS4{F0: (v9 >> ((0) & 63))}
                        var v38 uint32 = 4294967295
                        _ = len(v8)
                        _ = v37
                        _ = v38
                    }
                }
            }
        case 32768:
            fmt.Println(v3)
        default:
            if false {
                var v39 string = "foo bar"
                var v40 float64 = (v1 / v1)
                _ = v39
                _ = v40
            }
        }
    case "foo bar":
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 uint64 = uint64(v3)
            _ = fi41
            _ = v42
        }
    }
    v8 = append(v8, 882)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
