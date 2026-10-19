package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 string }
func genFn0(p0 int, p1 int32, p2 uint) int64 {
    p1 -= 42
    p2 = p2
    var v6 int8 = int8(p1)
    var v7 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v7 ^ -9223372036854775808)
}
func genFn1(p0 uint, p1 uint16) int {
    if (!true) {
        if (("" + "foo bar") < "a") {
            if (false || false) {
                p0 = (p0 | (^(p0)))
            }
            p0 = p0
            var v8 string = "foo bar"
            _ = v8
        } else {
            fmt.Println(p1, p0)
            for fi9 := 0; fi9 < 14; fi9++ {
                p1 = (p1 & (p1 ^ 0))
                _ = fi9
            }
            var v10 uint32 = 2147483648
            _ = v10
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: 756, F1: p0}
            v13 := GenS1{F0: 2147483647, F1: (^(p0))}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(p0)
        switch "hello" {
        case "foo bar":
            v15 := GenS1{F0: 65535, F1: 4294967295}
            _ = v15
        default:
            p1 = p1
        }
        _ = fi14
    }
    var v16 float64 = (-2.25 / 1.0)
    _ = p0
    _ = p1
    _ = v16
    return 836
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = -9223372036854775808
    var v3 string = ""
    v4 := []string{"foo bar", (v3 + (v3 + v3)), v3, "foo bar"}
    v5 := make(map[string]bool)
    v1 -= v1
    switch "" {
    case "\n\t":
        switch 2147483647 {
        case 127:
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint32 = 0
                v3 = ""
                _ = fi17
                _ = v18
            }
        case 4294967295:
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint8 = 154
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 int = fi21
                    v20 -= (v20 % (((v20 / (((v20 + v20)) | 1))) | 1))
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
        case 32768:
            fmt.Println(v1, v3, v3)
        }
    case "a":
        switch -523 {
        case 32767:
            var v23 float32 = -2.25
            _ = v23
        case 829:
            _ = len(v4)
        case 605:
            v2 -= genFn0(genFn1(2, (v1 >> ((42) & 63))), 2147483647, 4294967295)
        default:
            for fi24 := 0; fi24 < 8; fi24++ {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 float32 = -2.25
                    var v27 int16 = int16(v1)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = len(v4)
                _ = fi24
            }
        }
    case "foo bar":
        _ = len(v4)
    default:
        v28 := GenS2{F0: 18446744073709551615, F1: true}
        _ = v28
    }
    if (887 == -32768) {
        var v29 uint64 = 272
        var v30 uint16 = (v1 + (^(v1)))
        _ = v29
        _ = v30
    } else {
        v31 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: "a"}
        _ = v31
    }
    if false {
        var v32 int16 = -32768
        v32 = -18
        v1 = 42
        _ = v32
    }
    fmt.Println(v1, v3, v3)
    fmt.Println(v3, v1)
    if ((!(1 == 119)) && false) {
        var v33 uint32 = 2
        v3 = v3
        for fi34 := 0; fi34 < 9; fi34++ {
            if false {
                v1 = v1
                var v35 int32 = int32(fi34)
                _ = v35
            } else {
                for fi36 := 0; fi36 < 9; fi36++ {
                    switch "a" {
                    case "hello":
                        switch (("foo bar" + ("" + (v3 + v3))) + ((v3 + ("" + v3)) + ((v3 + v3) + (v3 + v3)))) {
                        case "foo bar":
                            var v37 float64 = 3.14159
                            _ = v37
                        default:
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                        }
                    }
                    var v38 int64 = (v2 + 153)
                    if false {
                        v2 += v38
                        fmt.Println(v1, fi36)
                    }
                    _ = fi36
                    _ = v38
                }
            }
            _ = fi34
        }
        _ = v33
    } else {
        if (!(242 != 1)) {
            fmt.Println(v3)
        }
        v3 = ((v3 + v3) + v3)
    }
    switch -2147483648 {
    case 196:
        v2 = v2
    }
    var v39 int8 = 127
    fmt.Println(v2, v1, v2)
    for fi40 := 0; fi40 < 1; fi40++ {
        var v41 int = (fi40 >> ((917) & 63))
        _ = fi40
        _ = v41
    }
    var v42 float32 = 1.5
    v2 = (v2 + v2)
    for fi43 := 0; fi43 < 13; fi43++ {
        v1 = (v1 - v1)
        var v44 int64 = genFn0(-526, 0, 65535)
        _ = fi43
        _ = v44
    }
    v3 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
