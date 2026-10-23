package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 103
    v3 := []int{103, -419, 1, 1}
    v4 := make(map[string]uint32)
    var v5 int8 = -11
    var v6 uint = 470
    var v7 uint16 = 0
    v2 += 1
    for fi8 := 0; fi8 < 18; fi8++ {
        delete(v4, "foo bar")
        if (3.14159 > 1e10) {
            if (3.14159 == float32(float64((float32((0.1 + (0.1 + (1e10 - 0.1)))) + 0.1)))) {
                if true {
                    var v9 int8 = v5
                    v6 = 255
                    fmt.Println(fi8, v9, fi8)
                    _ = v9
                } else {
                    v5 = v5
                }
            } else {
                v10 := GenS1{F0: 0, F1: v6}
                _ = v10
            }
            switch "hello" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                _ = len(v4)
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = (fi8 / ((fi8) | 1))
                }
            default:
                switch (-(fi8)) {
                case -475:
                    v11 := GenS1{F0: 32767, F1: (^(v6))}
                    _ = v11
                case 42:
                    v3 = append(v3, int(v1))
                default:
                    if (1 < -1) {
                        var v12 uint32 = uint32(v7)
                        v13 := GenS2{F0: v5, F1: "a", F2: 1}
                        if (!true) {
                            fmt.Println(v6)
                        }
                        _ = v12
                        _ = v13
                    } else {
                        switch int(v6) {
                        case 1:
                            fmt.Println(fi8, v1, fi8)
                        }
                        fmt.Println(fi8, v2, v5)
                    }
                }
            }
        } else {
            var v14 int32 = -1
            v3 = append(v3, int(v5))
            _ = v14
        }
        if (9223372036854775807 > 710) {
            v2 -= 4294967295
        }
        _ = fi8
    }
    var v15 uint16 = uint16(v1)
    if true {
        if ((false || false) || false) {
            var v16 int = 32767
            var v17 int8 = v5
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint64 = v2
                v17 = int8(v7)
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            v6 += 255
            switch "\n\t" {
            case "":
                fmt.Println(v15)
            case "\n\t":
                var v20 int = 32767
                _ = v20
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 700
                }
            default:
                v15 = v7
            }
            v6 = 18446744073709551615
        }
        var v21 int32 = 840
        v22 := GenS2{F0: 42, F1: "\n\t", F2: (^(v21))}
        _ = v21
        _ = v22
    }
    fmt.Println(v7)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int8 = int8(v1)
        delete(v4, "hello")
        _ = fi23
        _ = v24
    }
    v25 := GenS3{F0: uint(v7)}
    for fi26 := 0; fi26 < 14; fi26++ {
        v6 = (v6 & v6)
        for fi27 := 0; fi27 < 19; fi27++ {
            if false {
                var v28 uint = (v6 - (v6 % (((v6 * 232)) | 1)))
                _ = v28
            } else {
                v15 -= 32767
                switch fi27 {
                case -99:
                    var v29 uint64 = (^(v2))
                    _ = v29
                case 2147483647:
                    if (true && ((v15 & 42) >= v7)) {
                        fmt.Println(v1)
                        v30 := GenS2{F0: -66, F1: ("hello" + "foo bar"), F2: 42}
                        _ = v30
                    } else {
                        var v31 uint64 = 856
                        if true {
                            v5 -= -128
                            var v32 int16 = -475
                            _ = v32
                        }
                        v33 := v4[("\n\t" + "foo bar")]
                        _ = len(v4)
                        _ = v31
                        _ = v33
                    }
                }
            }
            if true {
                fmt.Println(fi26, v1)
            }
            _ = fi27
        }
        _ = fi26
    }
    v34 := GenS3{F0: v6}
    fmt.Println(v6)
    var v35 int = 9223372036854775807
    var v36 int32 = -2147483648
    var v37 int32 = v36
    if ((("hello" + (("" + (("foo bar" + "") + ("hello" + ("\n\t" + "a")))) + "a")) > ((("\n\t" + "") + "a") + "a")) || true) {
        var v38 uint = (v6 % ((v6) | 1))
        v37 = 624
        v35 = 292
        _ = v38
    } else {
        v39 := GenS4{F0: "\n\t", F1: (v15 | v7), F2: v1}
        v5 += v5
        v40 := GenS1{F0: 1, F1: (v6 / ((v6) | 1))}
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
