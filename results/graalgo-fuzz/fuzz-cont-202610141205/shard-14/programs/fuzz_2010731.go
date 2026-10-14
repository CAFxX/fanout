package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 string, p1 uint64, p2 int16) int32 {
    p1 -= p1
    p2 += (-(p2))
    for fi4 := 0; fi4 < 11; fi4++ {
        fmt.Println(p0, p2)
        v5 := GenS1{F0: -2147483648, F1: fi4, F2: 1}
        _ = fi4
        _ = v5
    }
    switch (p0 + p0) {
    case "\n\t":
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(fi6)
            switch p0 {
            case "":
                switch p0 {
                case "\n\t":
                    for fi7 := 0; fi7 < 6; fi7++ {
                        if (!false) {
                            fmt.Println(p1)
                        }
                        if (166 >= 705) {
                            v8 := GenS2{F0: 0}
                            p0 = "hello"
                            if (fi6 >= 1) {
                                var v9 uint64 = (p1 << ((18446744073709551615) & 63))
                                _ = v9
                            }
                            _ = v8
                        }
                        _ = fi7
                    }
                case "":
                    var v10 bool = (2147483647 != -820)
                    _ = v10
                case "foo bar":
                    if (4294967295 != 350) {
                        var v11 int32 = 42
                        _ = v11
                    } else {
                        p0 = p0
                    }
                default:
                    p0 = (p0 + "\n\t")
                }
            }
            _ = fi6
        }
    default:
        if false {
            var v12 float32 = 0.1
            p0 = p0
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint32 = 65535
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                switch -1 {
                case -1:
                    v16 := GenS3{F0: -9223372036854775808}
                    _ = v16
                case -2147483648:
                    var v17 uint32 = 2
                    _ = v17
                default:
                    if false {
                        fmt.Println(fi15, fi15, p0)
                        var v18 uint64 = 0
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 42, F1: fi15, F2: 94}
                        var v20 float64 = 0.0
                        p0 = p0
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi15
            }
            switch 32767 {
            case -955:
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 uint = 42
                    var v23 string = "foo bar"
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case -9223372036854775808:
                var v24 float64 = 3.14159
                _ = v24
            default:
                v25 := GenS2{F0: 0}
                _ = v25
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -766
}

func main() {
    v1 := []string{("" + "a"), "foo bar", "hello"}
    v2 := []string{"a", ("foo bar" + "\n\t"), ("a" + "hello")}
    var v3 uint = 93
    var v26 uint32 = 193
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v27 int = 2147483647
    v26 = 0
    var v28 uint32 = 1
    var v29 int32 = -2147483648
    fmt.Println(v29, v29)
    v28 += v28
    v30 := GenS2{F0: 73}
    for fi31 := 0; fi31 < 8; fi31++ {
        v3 += (v3 ^ (v3 % (((v3 % (((v3 / ((1) | 1))) | 1))) | 1)))
        v32 := GenS4{F0: 255}
        switch ("" + "hello") {
        case "\n\t":
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(v3, v28, v26)
                v3 += uint(v28)
                _ = fi33
            }
        case "hello":
            switch (fi31 % ((fi31) | 1)) {
            case 32768:
                v26 = v28
            case 127:
                fmt.Println(v29)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "a":
            if (!true) {
                _ = len(v1)
            }
        default:
            switch "" {
            case "\n\t":
                for fi34 := 0; fi34 < 5; fi34++ {
                    v3 += 4294967295
                    _ = fi34
                }
            }
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
