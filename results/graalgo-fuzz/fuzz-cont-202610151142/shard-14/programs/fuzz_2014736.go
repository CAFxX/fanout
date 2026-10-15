package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int32; F2 string }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 uint16 }

func main() {
    v1 := []uint32{685, 787}
    var v2 uint8 = 42
    var v3 uint8 = (v2 + v2)
    v4 := []uint32{4294967295, 2}
    var v5 uint8 = (v2 % ((1) | 1))
    v6 := make(map[string]uint32)
    fmt.Println(v2, v2)
    v2 += v3
    var v7 uint64 = 2
    v6[""] = 2147483647
    v2 += v2
    v7 += 4294967295
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch "\n\t" {
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    if (((((0.0 / (-2.25 / ((1.5 + (((-2.25 / (3.14159 + 3.14159)) + 1e10) - 3.14159)) + 0.0))) / -2.25) - -2.25) - 0.0) >= (1.5 * -2.25)) {
                        var v9 int = -1
                        _ = v9
                    } else {
                        var v10 int8 = -128
                        delete(v6, "hello")
                        var v11 int = 9223372036854775807
                        _ = v10
                        _ = v11
                    }
                case "hello":
                    v5 = 128
                case "\n\t":
                    v3 -= (^(v2))
                }
            case "":
                fmt.Println(v3, fi8, v5)
            default:
                if (2 <= 2) {
                    v3 -= v5
                    var v12 int8 = 100
                    fmt.Println(v5)
                    _ = v12
                }
            }
            var v13 float32 = 3.14159
            _ = fi8
            _ = v13
        }
    } else {
        v1 = append(v1, 70)
        if len(v4) > 0 {
            v4[0] = 0
        }
    }
    var v14 uint = 598
    _ = len(v6)
    v14 = v14
    if false {
        if true {
            if (0.1 >= 3.14159) {
                var v15 string = "hello"
                if true {
                    var v16 int16 = 32767
                    fmt.Println(v14, v7, v3)
                    _ = v16
                } else {
                    fmt.Println(v3, v14)
                    v5 -= (v5 - 2)
                    v3 -= (v2 / ((127) | 1))
                }
                v17 := GenS1{F0: 255, F1: 42}
                _ = v15
                _ = v17
            } else {
                if (1000 != 0) {
                    v18 := GenS1{F0: 600, F1: 32768}
                    _ = v18
                } else {
                    if (2147483647 < 65535) {
                        fmt.Println(v2, v14)
                    }
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    switch 1 {
                    case -2147483648:
                        if false {
                            for fi20 := 0; fi20 < 4; fi20++ {
                                v21 := GenS2{F0: 0}
                                v22 := GenS3{F0: (!true), F1: -550, F2: "\n\t"}
                                v23 := GenS4{F0: 2147483647, F1: fi19, F2: 65535}
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            var v24 float64 = -2.25
                            _ = v24
                        } else {
                            if (!(false || true)) {
                                switch (fi19 * 2147483647) {
                                case 1:
                                    v25 := GenS4{F0: 1, F1: (fi19 / (((^(fi19))) | 1)), F2: 295}
                                    _ = v25
                                default:
                                    v7 -= v7
                                }
                            }
                        }
                    case -51:
                        _ = len(v1)
                    }
                    _ = fi19
                }
            }
        } else {
            if ("a" < ("a" + ("hello" + ""))) {
                v2 += 42
                var v26 int = -2147483648
                switch 0 {
                case -391:
                    v27 := v6[("\n\t" + "foo bar")]
                    _ = v27
                case 2147483647:
                    if false {
                        _ = len(v4)
                        if (-9223372036854775808 <= -955) {
                            v28 := GenS5{F0: (v26 & v26), F1: 255, F2: 134}
                            _ = v28
                        }
                        v1 = append(v1, 91)
                    }
                case 73:
                    v29 := GenS4{F0: 42, F1: v26, F2: 32767}
                    _ = v29
                }
                _ = v26
            }
            v14 -= v14
            v30 := v6[("foo bar" + ((("hello" + "") + "") + ("hello" + ((("\n\t" + ("hello" + "\n\t")) + ("a" + "foo bar")) + ("a" + "a")))))]
            _ = v30
        }
        v31 := GenS3{F0: (-935 < 128), F1: 42, F2: (("hello" + "\n\t") + "")}
        _ = v31
    }
    var v32 int16 = 42
    var v33 int = 9223372036854775807
    fmt.Println(v14, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
