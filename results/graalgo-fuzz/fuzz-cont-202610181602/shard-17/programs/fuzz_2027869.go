package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 255
    v3 := []int{}
    switch -290 {
    case 128:
        v4 := GenS1{F0: 127}
        _ = v4
    case 127:
        if true {
            v5 := GenS1{F0: 0}
            _ = v5
        }
    }
    v6 := GenS1{F0: 32767}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v7 int16 = 1
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    switch 4294967295 {
    case 42:
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS2{F0: 255}
            if false {
                var v11 int8 = -42
                if ((-2.25 / -2.25) >= (1e10 / 3.14159)) {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v2 = (v2 >> ((2) & 63))
                        v11 = v11
                        if (!(0.1 != (3.14159 + (0.0 - -2.25)))) {
                            v13 := v1[65535]
                            v13 = v13
                            _ = v13
                        }
                        switch ("\n\t" + "\n\t") {
                        case "":
                            v11 += (^(v11))
                        }
                        _ = fi12
                    }
                    v14 := GenS2{F0: 0}
                    _ = v14
                }
                _ = v11
            } else {
                switch ("hello" + "") {
                case "hello":
                    _ = len(v3)
                case "foo bar":
                    if (693 == 18446744073709551615) {
                        if (fi9 == -1) {
                            fmt.Println(v2, fi9)
                            fmt.Println(fi9, fi9)
                        }
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v1[65535] = 42
                            v16 := GenS2{F0: 65535}
                            if len(v3) > 0 {
                                v3[0] = (-(fi15))
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 float32 = float32(float64((0.1 * 1e10)))
                    _ = fi17
                    _ = v18
                }
            }
            _ = fi9
            _ = v10
        }
    }
    var v19 uint8 = 1
    v19 -= v19
    for fi20 := 0; fi20 < 4; fi20++ {
        if false {
            switch "" {
            case "foo bar":
                if (("\n\t" + "a") <= ("a" + "foo bar")) {
                    if ((^(fi20)) > fi20) {
                        var v21 int16 = int16(v2)
                        if true {
                            v22 := GenS2{F0: 148}
                            fmt.Println(v21, v21, fi20)
                            if (643 != 671) {
                                v23 := v1[441]
                                _ = v23
                            }
                            _ = v22
                        } else {
                            if true {
                                var v24 uint16 = v2
                                v25 := GenS3{F0: (v19 & (v19 << ((65535) & 63))), F1: -128}
                                _ = v24
                                _ = v25
                            }
                            v19 += (v19 | (v19 / ((v19) | 1)))
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v19 = v19
                            var v27 int32 = 0
                            _ = fi26
                            _ = v27
                        }
                        _ = v21
                    }
                    v1[2147483648] = 979
                } else {
                    switch (fi20 ^ fi20) {
                    case -9223372036854775808:
                        v19 = v19
                    case 128:
                        fmt.Println(v19)
                    case -331:
                        var v28 uint64 = 0
                        _ = v28
                    default:
                        if (true && true) {
                            v19 = (^(v19))
                            switch fi20 {
                            case 127:
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v2, fi20, v2)
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "a":
                                    var v29 uint32 = 4294967295
                                    _ = v29
                                }
                            case -627:
                                var v30 uint64 = 18446744073709551615
                                _ = v30
                            default:
                                if len(v3) > 0 {
                                    v3[0] = fi20
                                }
                            }
                        }
                    }
                }
            }
            delete(v1, 339)
            fmt.Println(v19, v19)
        }
        _ = fi20
    }
    v2 = v2
    v19 -= v19
    delete(v1, 0)
    fmt.Println(v19, v2)
    switch ("foo bar" + ("foo bar" + ("" + ((("\n\t" + "hello") + ("foo bar" + "a")) + "\n\t")))) {
    case "":
        v31 := GenS3{F0: v19, F1: int8(v2)}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
