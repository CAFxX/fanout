package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint16 = 32767
    v2 := []int{32768, -1, -9223372036854775808}
    v3 := make(map[string]string)
    switch 42 {
    case 523:
        if true {
            v4 := GenS1{F0: -85, F1: 65535, F2: (!true)}
            _ = v4
        }
    case -9223372036854775808:
        v5 := GenS1{F0: -102, F1: 65535, F2: true}
        _ = v5
    default:
        var v6 uint = 0
        _ = v6
    }
    if (((127 <= 128) || ((-376 == 0) || (18446744073709551615 < 4294967295))) || (true || (88 < -128))) {
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                v2 = append(v2, -776)
                if (!(1000 <= -32768)) {
                    v10 := GenS2{F0: -699, F1: -128}
                    if true {
                        v7 = 2147483648
                        if (42 == -32768) {
                            fmt.Println(fi8)
                        } else {
                            fmt.Println(v1, v1)
                            var v11 uint16 = v1
                            _ = v11
                        }
                    }
                    v12 := GenS3{F0: (1 < -1), F1: 0, F2: 32767}
                    _ = v10
                    _ = v12
                }
                var v13 uint16 = v1
                _ = fi9
                _ = v13
            }
            var v14 uint = 65535
            _ = fi8
            _ = v14
        }
        var v15 int32 = -466
        _ = v15
        _ = v7
    } else {
        if false {
            v1 -= v1
            fmt.Println(v1)
            v1 += 65535
        } else {
            var v16 uint64 = 0
            var v17 uint16 = v1
            _ = v16
            _ = v17
        }
        var v18 int8 = -124
        if true {
            var v19 bool = (227 <= 0)
            _ = v19
        } else {
            if true {
                var v20 float64 = (3.14159 - float64(1e10))
                for fi21 := 0; fi21 < 2; fi21++ {
                    v20 -= (float64(float32((v20 + v20))) * v20)
                    v22 := GenS3{F0: (((511 <= -9223372036854775808) && false) && (!(!((("hello" + "hello") + ("hello" + "hello")) > "\n\t")))), F1: -437, F2: -32768}
                    if (-331 <= 42) {
                        if (-32768 <= 1000) {
                            v1 += 362
                            var v23 float32 = (0.1 * ((-2.25 / 0.1) / (0.1 / 1.5)))
                            v18 -= (v18 & v18)
                            _ = v23
                        } else {
                            var v24 int16 = 32767
                            fmt.Println(v20, fi21)
                            if (18446744073709551615 == 2) {
                                var v25 int8 = (v18 & v18)
                                _ = v25
                            } else {
                                var v26 uint32 = 580
                                var v27 int64 = 0
                                _ = v26
                                _ = v27
                            }
                            _ = v24
                        }
                    } else {
                        fmt.Println(v1, v20, v18)
                        v20 = v20
                    }
                    _ = fi21
                    _ = v22
                }
                var v28 int8 = v18
                _ = v20
                _ = v28
            }
            switch "\n\t" {
            case "\n\t":
                v29 := GenS4{F0: ""}
                _ = v29
            }
        }
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v30 uint = 4294967295
    var v31 int64 = -430
    var v32 int = 582
    v33 := GenS4{F0: "a"}
    for fi34 := 0; fi34 < 20; fi34++ {
        switch (-(fi34)) {
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = (^(fi34))
            }
        case 32768:
            for fi35 := 0; fi35 < 3; fi35++ {
                if false {
                    v36 := GenS2{F0: -32768, F1: 0}
                    var v37 bool = (566 == (v31 >> (((^(v30))) & 63)))
                    _ = v36
                    _ = v37
                }
                var v38 uint = (^(v30))
                fmt.Println(fi34, fi35, v30)
                _ = fi35
                _ = v38
            }
        }
        if (65535 > 696) {
            switch "\n\t" {
            case "foo bar":
                if (2 < 65535) {
                    v32 += fi34
                }
            case "\n\t":
                v31 = v31
            case "hello":
                v30 = (v30 << (((v30 ^ 1)) & 63))
            default:
                fmt.Println(v32, v1)
            }
            v3["\n\t"] = ("foo bar" + "\n\t")
        }
        fmt.Println(fi34, fi34, v31)
        _ = fi34
    }
    v30 = (^(v30))
    v30 = (v30 | (^(v30)))
    var v39 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v39
}
