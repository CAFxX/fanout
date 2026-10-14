package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 int) int8 {
    var v4 int32 = 2147483647
    var v5 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 uint = 181
    var v3 uint = (v2 * v2)
    for fi6 := 0; fi6 < 17; fi6++ {
        _ = len(v1)
        _ = fi6
    }
    v3 = 255
    v2 = (v2 | 2)
    v2 -= 1
    v3 = (^(v2))
    var v7 int32 = 124
    for fi8 := 0; fi8 < 19; fi8++ {
        if (false && true) {
            fmt.Println(v2)
            switch (fi8 >> ((v3) & 63)) {
            case 951:
                var v9 int8 = 45
                _ = v9
            case 665:
                switch "a" {
                case "a":
                    v2 = (v3 / (((v2 ^ v3)) | 1))
                }
            case 32767:
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS1{F0: 349, F1: "foo bar"}
                    v12 := GenS1{F0: 2147483648, F1: ("\n\t" + ("hello" + (("foo bar" + ("hello" + "\n\t")) + "\n\t")))}
                    v3 = v3
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            switch (^(fi8)) {
            case 2147483647:
                var v13 int = fi8
                _ = v13
            default:
                switch "hello" {
                case "foo bar":
                    var v14 string = "\n\t"
                    _ = v14
                default:
                    if (2 == 0) {
                        v3 = 284
                    }
                }
            }
        } else {
            v2 = (v2 % (((v3 / ((v3) | 1))) | 1))
            v3 -= v3
            v15 := GenS2{F0: 1}
            _ = v15
        }
        v16 := GenS2{F0: 2}
        _ = fi8
        _ = v16
    }
    if (42 == 4294967295) {
        var v17 uint32 = 24
        switch "\n\t" {
        case "":
            v18 := GenS1{F0: 668, F1: (("a" + ("a" + "\n\t")) + "")}
            _ = v18
        case "a":
            if false {
                if true {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch "hello" {
                        case "foo bar":
                            if (!((("\n\t" + ("\n\t" + "hello")) + ((("\n\t" + ("\n\t" + ((("foo bar" + "") + ("a" + "")) + "hello"))) + "") + "")) == (("a" + "") + "foo bar"))) {
                                if true {
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        v1 = append(v1, "foo bar")
                                        var v21 uint = (v2 / ((v2) | 1))
                                        v21 = v3
                                        switch "hello" {
                                        case "":
                                            if true {
                                                v3 = 4294967295
                                            }
                                            v2 += (^(v21))
                                            var v22 int32 = (v7 >> ((v3) & 63))
                                            _ = v22
                                        default:
                                            fmt.Println(v7, v21, v2)
                                        }
                                        _ = fi20
                                        _ = v21
                                    }
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        var v23 uint8 = 127
                                        _ = v23
                                    default:
                                        var v24 int8 = 0
                                        _ = v24
                                    }
                                    var v25 int32 = int32(v3)
                                    _ = v25
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = ("\n\t" + ("hello" + "foo bar"))
                                    }
                                }
                            }
                        case "":
                            v7 += (v7 >> (((^(v3))) & 63))
                        case "a":
                            var v26 int16 = 42
                            _ = v26
                        }
                        _ = fi19
                    }
                    var v27 uint64 = 4294967295
                    v28 := GenS3{F0: (^(v3))}
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v2, v2)
                }
            }
        }
        var v29 uint32 = v17
        _ = v17
        _ = v29
    }
    if (v2 < (v2 & 0)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v30 int32 = v7
    if false {
        v2 -= (v3 % (((v3 - (v3 >> ((v2) & 63)))) | 1))
        v30 = (v7 & (v7 % (((^(v7))) | 1)))
        var v31 uint8 = 1
        _ = v31
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v7 = (v30 >> (((v3 ^ (^(v2)))) & 63))
    for fi32 := 0; fi32 < 3; fi32++ {
        v33 := GenS4{F0: 127, F1: 215}
        var v34 string = "a"
        v30 += 2147483647
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 uint8 = 255
    switch 9223372036854775807 {
    case -2147483648:
        fmt.Println(v3, v35)
    }
    fmt.Println(v30)
    switch ("hello" + "a") {
    case "foo bar":
        if (("\n\t" + (("a" + (((("" + ("hello" + "hello")) + ("foo bar" + "a")) + "") + ("a" + ("\n\t" + "\n\t")))) + "foo bar")) >= "") {
            if (!(!(float32((((3.14159 / 3.14159) - (0.0 - 1.5)) + 3.14159)) <= 3.14159))) {
                if len(v1) > 0 {
                    v1[0] = (("\n\t" + "\n\t") + "")
                }
            } else {
                var v36 int32 = v30
                _ = v36
            }
            var v37 uint32 = 2147483648
            _ = v37
        }
    case "a":
        var v38 int8 = -128
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v7
}
