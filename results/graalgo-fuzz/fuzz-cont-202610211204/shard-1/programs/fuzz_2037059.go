package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 1
    var v3 int = 32767
    v4 := make(map[uint32]bool)
    v5 := []int{(-(v3))}
    v3 += (v3 * (v3 * 9223372036854775807))
    var v6 int16 = v2
    fmt.Println(v2, v6, v6)
    var v7 float32 = (0.0 + float32(1.5))
    switch 32767 {
    case -2147483648:
        for fi8 := 0; fi8 < 6; fi8++ {
            switch (v3 + fi8) {
            case 42:
                var v9 int8 = -128
                _ = v9
            case 4294967295:
                var v10 uint32 = 2
                _ = v10
            }
            v3 = int(v2)
            for fi11 := 0; fi11 < 14; fi11++ {
                v1 = 0.0
                v1 = (v1 - 3.14159)
                _ = fi11
            }
            _ = fi8
        }
    case 42:
        v7 = (v7 * v7)
    default:
        var v12 uint16 = 32767
        _ = v12
    }
    if ((14 < 444) || false) {
        var v13 bool = (0 >= 518)
        v14 := GenS1{F0: "foo bar", F1: 128}
        v15 := v4[0]
        _ = v13
        _ = v14
        _ = v15
    } else {
        switch v3 {
        case 32767:
            fmt.Println(v3)
        case 42:
            for fi16 := 0; fi16 < 5; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    if false {
                        fmt.Println(v6)
                        v6 = v6
                    }
                    _ = fi17
                }
                _ = fi16
            }
        case -940:
            fmt.Println(v3, v7)
        default:
            switch "\n\t" {
            case "hello":
                v6 -= (^(v2))
            case "\n\t":
                fmt.Println(v1, v6, v6)
            case "foo bar":
                v1 = v1
            default:
                switch "foo bar" {
                case "foo bar":
                    v18 := GenS2{F0: 2147483647, F1: ("" + ("a" + "hello"))}
                    _ = v18
                case "\n\t":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 float32 = (v7 - v7)
                        _ = fi19
                        _ = v20
                    }
                case "hello":
                    for fi21 := 0; fi21 < 6; fi21++ {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v6 += 42
                            delete(v4, 65535)
                            _ = len(v4)
                            _ = fi22
                        }
                        v1 += (v1 / (1.5 - 0.0))
                        _ = fi21
                    }
                default:
                    if (-9223372036854775808 <= 42) {
                        for fi23 := 0; fi23 < 2; fi23++ {
                            v7 = ((-2.25 * (0.0 + v7)) / (v7 + v7))
                            _ = fi23
                        }
                    }
                }
            }
        }
        switch -9223372036854775808 {
        case 32768:
            v24 := GenS2{F0: 4294967295, F1: "hello"}
            _ = v24
        default:
            if (127 < 2) {
                var v25 int64 = 9223372036854775807
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := GenS3{F0: 4294967295, F1: 972, F2: v25}
                    v28 := GenS2{F0: 795, F1: "a"}
                    switch (fi26 & 32768) {
                    case 597:
                        if len(v5) > 0 {
                            v5[0] = (v3 | (v3 / (((v3 & (v3 & 2147483647))) | 1)))
                        }
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            } else {
                var v29 string = "hello"
                if (!(20 != 42)) {
                    v6 -= v2
                    var v30 uint16 = 65535
                    for fi31 := 0; fi31 < 17; fi31++ {
                        if ((v6 / (((v6 + (^(v2)))) | 1)) < 42) {
                            v5 = append(v5, v3)
                            var v32 int64 = 9223372036854775807
                            var v33 float32 = v7
                            _ = v32
                            _ = v33
                        }
                        switch (fi31 % (((-(fi31))) | 1)) {
                        case 9223372036854775807:
                            v4[65535] = (0 > -1)
                        }
                        var v34 float32 = (((v7 / (float32(-2.25) + 1e10)) + 1e10) * 3.14159)
                        _ = fi31
                        _ = v34
                    }
                    _ = v30
                }
                var v35 int64 = 42
                _ = v29
                _ = v35
            }
        }
        fmt.Println(v6, v7)
    }
    var v36 uint16 = 32767
    var v37 int8 = 1
    if true {
        v7 = ((0.0 / v7) / (v7 / ((0.1 - 3.14159) - 0.0)))
        var v38 int = (v3 ^ (v3 << ((612) & 63)))
        for fi39 := 0; fi39 < 7; fi39++ {
            v40 := GenS2{F0: 2, F1: (("a" + "") + "foo bar")}
            _ = fi39
            _ = v40
        }
        _ = v38
    } else {
        v37 = (v37 - 127)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
