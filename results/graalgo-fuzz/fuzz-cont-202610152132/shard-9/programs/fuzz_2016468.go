package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 uint; F1 bool; F2 int }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 957
    v2 := []string{"foo bar", ((("\n\t" + ("a" + "foo bar")) + ("\n\t" + ("hello" + (((("hello" + ("a" + "")) + ("hello" + "hello")) + ((("a" + (("\n\t" + "a") + (("\n\t" + "") + "foo bar"))) + "a") + ("a" + (("foo bar" + "\n\t") + "")))) + "")))) + "a"), "", "a"}
    var v3 int16 = 1
    var v4 int32 = -49
    var v5 int16 = (-(v3))
    switch (("hello" + ("a" + "foo bar")) + "\n\t") {
    case "foo bar":
        v3 = v3
    case "hello":
        var v6 float32 = 3.14159
        _ = v6
    default:
        switch ("foo bar" + ("a" + "a")) {
        case "hello":
            switch "a" {
            case "hello":
                switch ("\n\t" + "a") {
                case "\n\t":
                    v1 = v1
                case "hello":
                    v5 = (v3 | v5)
                default:
                    v7 := GenS1{F0: v1, F1: 1}
                    _ = v7
                }
            default:
                switch 2147483647 {
                case 4294967295:
                    for fi8 := 0; fi8 < 11; fi8++ {
                        if (!true) {
                            v9 := GenS1{F0: v1, F1: 65535}
                            for fi10 := 0; fi10 < 8; fi10++ {
                                v11 := GenS2{F0: (v3 & 164), F1: -9223372036854775808, F2: (fi8 * (-(fi8)))}
                                switch fi10 {
                                case 352:
                                    v1 = v1
                                }
                                if false {
                                    v1 -= (v1 * (^(v1)))
                                    var v12 float64 = 1e10
                                    _ = v12
                                } else {
                                    v1 = 365
                                    switch "\n\t" {
                                    case "hello":
                                        v2 = append(v2, "hello")
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            switch "a" {
                                            case "foo bar":
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            case "a":
                                                v5 -= (v5 ^ (v3 % ((-1) | 1)))
                                            case "":
                                                v5 = -1
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(v1)
                                                case "":
                                                    v13 := GenS3{F0: 2, F1: (330 == v1), F2: 42}
                                                    _ = v13
                                                case "\n\t":
                                                    v2 = append(v2, (((("foo bar" + "foo bar") + "foo bar") + "foo bar") + ("hello" + "a")))
                                                }
                                            }
                                        }
                                    case "":
                                        if false {
                                            v3 -= v3
                                        }
                                    }
                                    switch (fi10 >> ((18446744073709551615) & 63)) {
                                    case -804:
                                        v4 = (v4 % ((v4) | 1))
                                    default:
                                        if ((-104 == 127) || true) {
                                            v5 -= 67
                                            var v14 int8 = -128
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v14
                                        } else {
                                            fmt.Println(v4)
                                            v3 = 42
                                            var v15 uint64 = 18446744073709551615
                                            v1 = (v1 >> ((822) & 63))
                                            _ = v15
                                        }
                                    }
                                    var v16 uint8 = 42
                                    switch "hello" {
                                    case "foo bar":
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                    }
                                    _ = v16
                                }
                                _ = fi10
                                _ = v11
                            }
                            for fi17 := 0; fi17 < 12; fi17++ {
                                switch fi17 {
                                case -215:
                                    v2 = append(v2, (("a" + "a") + "a"))
                                case 4294967295:
                                    var v18 float32 = 0.0
                                    _ = v18
                                }
                                var v19 uint32 = 0
                                fmt.Println(v19, fi8)
                                _ = fi17
                                _ = v19
                            }
                            _ = v9
                        } else {
                            v4 = v4
                        }
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 uint16 = 841
                            v5 += v3
                            _ = fi20
                            _ = v21
                        }
                        fmt.Println(fi8, v1, v4)
                        _ = fi8
                    }
                case 32767:
                    var v22 uint64 = 719
                    _ = v22
                default:
                    switch ((((("hello" + (("hello" + "hello") + (("foo bar" + "\n\t") + "hello"))) + "foo bar") + "hello") + "a") + "\n\t") {
                    case "\n\t":
                        v3 = -32768
                    case "foo bar":
                        var v23 int64 = -810
                        _ = v23
                    default:
                        if false {
                            fmt.Println(v3, v4, v4)
                        } else {
                            v1 = 692
                            v24 := GenS4{F0: 367}
                            _ = v24
                        }
                    }
                }
            }
        case "\n\t":
            var v25 int8 = 1
            _ = v25
        case "foo bar":
            v3 = v5
        default:
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 float32 = (3.14159 + 1.5)
                v3 += v3
                v27 -= v27
                _ = fi26
                _ = v27
            }
        }
    }
    v3 = v3
    v28 := GenS3{F0: 2, F1: true, F2: 611}
    v5 = v5
    fmt.Println(v1, v3, v1)
    v4 = -1
    v4 = 0
    var v29 int16 = (v5 - v3)
    if false {
        v4 -= (v4 % ((v4) | 1))
        v2 = append(v2, "")
    } else {
        fmt.Println(v5, v1, v29)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
