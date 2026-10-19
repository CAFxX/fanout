package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 bool }

func main() {
    var v1 int16 = -981
    v2 := []string{("" + ("" + "a")), "\n\t", "", "hello"}
    var v3 int = -363
    var v4 uint64 = 42
    v5 := []string{"hello", "", ((("hello" + ("" + (("\n\t" + "a") + "foo bar"))) + "\n\t") + ""), (("\n\t" + "foo bar") + "foo bar")}
    var v6 uint = 255
    switch ("a" + "") {
    case "foo bar":
        v6 -= (v6 / (((^(v6))) | 1))
    default:
        v6 = (v6 ^ v6)
    }
    fmt.Println(v3, v1, v6)
    v4 += (v4 & 579)
    var v7 uint8 = 127
    fmt.Println(v6)
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: int8(fi9), F1: 1}
        var v11 int64 = 156
        v1 = 32767
        _ = fi9
        _ = v10
        _ = v11
    }
    switch -1 {
    case 42:
        for fi12 := 0; fi12 < 6; fi12++ {
            if (float32(0.1) == v8) {
                v7 += v7
                var v13 float32 = (v8 + v8)
                _ = v13
            }
            v8 = (3.14159 / (((float32((float64(v8) * (1.5 * 3.14159))) - float32(-2.25)) * v8) - 1e10))
            for fi14 := 0; fi14 < 13; fi14++ {
                v3 += fi12
                switch fi14 {
                case -249:
                    var v15 int16 = -1000
                    _ = v15
                default:
                    fmt.Println(v3)
                }
                _ = fi14
            }
            _ = fi12
        }
    case 32767:
        fmt.Println(v1, v6, v4)
    case -453:
        var v16 int32 = -393
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v5 = append(v5, "a")
        if ((true || (100 == -128)) && (!(!false))) {
            _ = len(v2)
            var v18 int64 = 0
            switch (v3 ^ -714) {
            case -264:
                var v19 string = "\n\t"
                _ = v19
            case 876:
                var v20 int = fi17
                _ = v20
            }
            _ = v18
        }
        v1 -= v1
        _ = fi17
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v5 = append(v5, "foo bar")
        v6 = v6
        if false {
            fmt.Println(v1)
            fmt.Println(v8, fi21)
        } else {
            var v22 int16 = (v1 * -1)
            for fi23 := 0; fi23 < 15; fi23++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    v3 += (-(v3))
                    v24 := GenS2{F0: (!false), F1: 32767, F2: 127}
                    switch "" {
                    case "foo bar":
                        fmt.Println(fi21)
                    case "\n\t":
                        v4 += v4
                    }
                    _ = v24
                } else {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            var v25 int = (fi23 & (-(v3)))
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            _ = v25
                        } else {
                            switch fi23 {
                            case 1:
                                fmt.Println(v1, v4, v6)
                            case 0:
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            case -1:
                                _ = len(v2)
                            default:
                                fmt.Println(fi23, fi21, v1)
                            }
                            fmt.Println(v3, v7)
                        }
                    case "foo bar":
                        v27 := GenS2{F0: (!(!(!true))), F1: 65535, F2: -1}
                        _ = v27
                    }
                    if true {
                        switch 128 {
                        case -644:
                            v28 := GenS3{F0: 515, F1: v4}
                            _ = v28
                        case -2147483648:
                            if (-2.25 > -2.25) {
                                if len(v5) > 0 {
                                    v5[0] = ""
                                }
                                v22 = v1
                                if len(v2) > 0 {
                                    v2[0] = ((("\n\t" + ("" + ("\n\t" + "foo bar"))) + ("" + ("foo bar" + ("hello" + "a")))) + (((("" + ("hello" + "hello")) + "\n\t") + "hello") + "hello"))
                                }
                                var v29 bool = false
                                _ = v29
                            }
                        case -615:
                            switch -2147483648 {
                            case 575:
                                v6 = (v6 & (v6 ^ (v6 ^ v6)))
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v30 uint8 = 42
                        _ = v30
                    }
                }
                _ = fi23
            }
            v4 -= (v4 % ((2) | 1))
            _ = v22
        }
        _ = fi21
    }
    v4 += (^(v4))
    v1 = (v1 | 32767)
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS4{F0: v3, F1: 1, F2: false}
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
