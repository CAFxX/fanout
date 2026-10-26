package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int64; F1 int64; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int32; F1 uint; F2 string }

func main() {
    var v1 int = 690
    v2 := []string{}
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 bool = v3
    if v3 {
        fmt.Println(v5, v1)
        if false {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 int64 = 42
                fmt.Println(v5)
                _ = fi6
                _ = v7
            }
        } else {
            v8 := GenS1{F0: (!v5), F1: 1, F2: (("hello" < "\n\t") && (true || v5))}
            fmt.Println(v5)
            v5 = (!v3)
            _ = v8
        }
        if v3 {
            for fi9 := 0; fi9 < 2; fi9++ {
                v1 = (fi9 - fi9)
                for fi10 := 0; fi10 < 13; fi10++ {
                    v1 += v1
                    var v11 uint = 735
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            switch (v1 / ((v1) | 1)) {
            case -9223372036854775808:
                fmt.Println(v5)
            case 32767:
                var v12 uint16 = 42
                _ = v12
            }
            var v13 bool = (!true)
            _ = v13
        } else {
            v14 := GenS1{F0: (!(325 >= 888)), F1: 576, F2: v5}
            v5 = v3
            v15 := GenS2{F0: 42, F1: ""}
            _ = v14
            _ = v15
        }
    } else {
        v3 = ((3.14159 * (0.1 / ((3.14159 - ((3.14159 * (1.5 * float32((1.5 + 0.1)))) + ((((((1e10 - (3.14159 - (0.1 * 3.14159))) / float32(1.5)) - -2.25) * (-2.25 / 1e10)) / ((-2.25 - 1e10) / float32(1e10))) * (0.1 * 3.14159)))) - 3.14159))) <= 0.1)
        var v16 int64 = 1
        _ = v16
    }
    v5 = v5
    v3 = (v3 && v3)
    v5 = (131 >= 255)
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS3{F0: 42, F1: -1, F2: 4294967295}
        for fi19 := 0; fi19 < 19; fi19++ {
            v5 = (v5 && v5)
            switch (fi19 - v1) {
            case -628:
                switch (fi17 ^ 2147483647) {
                case -2147483648:
                    switch (("hello" + ("hello" + ("foo bar" + "\n\t"))) + "a") {
                    case "\n\t":
                        if (("foo bar" + "hello") > "") {
                            var v20 float32 = 1e10
                            if (false || v3) {
                                v20 -= (float32(((3.14159 + ((-2.25 + 0.0) + -2.25)) - 1.5)) / 0.1)
                                v20 += (v20 * v20)
                                if (0 <= 1) {
                                    var v21 bool = false
                                    if (607 <= 4294967295) {
                                        switch v1 {
                                        case -1:
                                            switch 2147483647 {
                                            case -238:
                                                if (("\n\t" + ("hello" + ("foo bar" + "\n\t"))) >= ("\n\t" + "hello")) {
                                                    v5 = (-183 < -342)
                                                }
                                            case -196:
                                                fmt.Println(v5, v1)
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (((1 == 42) || v3) || ((!true) || v3)) {
                                                        var v22 int32 = 644
                                                        switch ((("" + ("hello" + ("a" + "\n\t"))) + "a") + (("a" + "foo bar") + (("" + "") + ("a" + "foo bar")))) {
                                                        case "":
                                                            var v23 uint64 = 2
                                                            _ = v23
                                                        case "a":
                                                            var v24 uint8 = 0
                                                            _ = v24
                                                        default:
                                                            delete(v4, 65535)
                                                        }
                                                        if (!(false || (v3 && ((!(42 < -32768)) && true)))) {
                                                            _ = len(v2)
                                                            var v25 uint16 = 255
                                                            _ = v25
                                                        }
                                                        _ = v22
                                                    } else {
                                                        if true {
                                                            if false {
                                                                v20 -= v20
                                                                v26 := GenS1{F0: ((("foo bar" + (("hello" + "\n\t") + "")) + "foo bar") == (("" + (("foo bar" + "foo bar") + ((("" + (("hello" + "foo bar") + (("a" + ("hello" + "")) + ("a" + "\n\t")))) + "\n\t") + "\n\t"))) + ("a" + "hello"))), F1: 0, F2: v21}
                                                                _ = v26
                                                            }
                                                        }
                                                        v27 := GenS4{F0: fi17}
                                                        _ = v27
                                                    }
                                                case "\n\t":
                                                    var v28 uint32 = 2147483648
                                                    _ = v28
                                                case "":
                                                    fmt.Println(v5)
                                                default:
                                                    var v29 uint8 = 42
                                                    _ = v29
                                                }
                                            }
                                        case 32768:
                                            var v30 int = fi17
                                            _ = v30
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                        }
                                        switch (fi17 ^ fi17) {
                                        case 2147483647:
                                            var v31 string = "a"
                                            _ = v31
                                        }
                                        fmt.Println(v20)
                                    } else {
                                        v32 := GenS3{F0: -1, F1: -9223372036854775808, F2: 1}
                                        _ = v32
                                    }
                                    _ = v21
                                } else {
                                    var v33 uint = 4294967295
                                    switch "" {
                                    case "":
                                        v20 -= v20
                                    }
                                    v20 = v20
                                    _ = v33
                                }
                                var v34 float32 = float32(0.0)
                                delete(v4, 0)
                                _ = v34
                            }
                            var v35 float64 = (((1.5 / (float64(v20) + float64(v20))) * (-2.25 + 1.5)) - (0.1 + -2.25))
                            switch "\n\t" {
                            case "a":
                                if v5 {
                                    var v36 int16 = 1000
                                    _ = v36
                                } else {
                                    v4[2] = 1
                                }
                            case "\n\t":
                                var v37 int64 = 42
                                _ = v37
                            case "foo bar":
                                switch "\n\t" {
                                case "":
                                    var v38 uint32 = 65535
                                    _ = v38
                                case "foo bar":
                                    if (!true) {
                                        v35 = v35
                                        switch ("foo bar" + "hello") {
                                        case "":
                                            v39 := GenS2{F0: 32767, F1: ""}
                                            _ = v39
                                        case "a":
                                            var v40 int64 = 9223372036854775807
                                            _ = v40
                                        case "hello":
                                            v1 = v1
                                        }
                                    }
                                case "hello":
                                    if v3 {
                                        var v41 int = -944
                                        _ = v41
                                    } else {
                                        if (v3 && (0 <= 18446744073709551615)) {
                                            var v42 uint32 = 1
                                            if v5 {
                                                var v43 uint64 = 403
                                                _ = v43
                                            }
                                            _ = v42
                                        }
                                        _ = len(v2)
                                    }
                                default:
                                    v44 := GenS5{F0: 0}
                                    _ = v44
                                }
                            }
                            _ = v20
                            _ = v35
                        }
                    case "a":
                        var v45 float32 = float32(float64(0.1))
                        _ = v45
                    default:
                        switch v1 {
                        case 9223372036854775807:
                            if (v3 || (v5 || v3)) {
                                v5 = (("" + ("a" + "\n\t")) != "hello")
                                if v5 {
                                    v2 = append(v2, "\n\t")
                                    fmt.Println(fi19)
                                }
                            } else {
                                fmt.Println(v5, fi17, fi17)
                                if (!(true || (!v5))) {
                                    if (!(!false)) {
                                        var v46 bool = false
                                        if (18446744073709551615 <= 2) {
                                            fmt.Println(fi17, v3, fi17)
                                            v1 = v1
                                        }
                                        var v47 uint16 = 511
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := GenS4{F0: fi19}
                                    _ = v48
                                } else {
                                    var v49 int8 = -126
                                    _ = v49
                                }
                                v50 := GenS6{F0: 0, F1: 318, F2: 0}
                                _ = v50
                            }
                        case -9223372036854775808:
                            v1 += 0
                        }
                    }
                case 1:
                    v3 = (false && (("hello" + ("a" + ((("hello" + "a") + "") + "hello"))) >= ("\n\t" + "\n\t")))
                case 128:
                    v5 = v5
                default:
                    v5 = (!v5)
                }
            }
            _ = fi19
        }
        v1 -= fi17
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v5)
    if false {
        for fi51 := 0; fi51 < 11; fi51++ {
            var v52 bool = v5
            _ = fi51
            _ = v52
        }
    } else {
        v4[4294967295] = 2
    }
    var v53 uint16 = 42
    v54 := GenS7{F0: 42, F1: 1}
    switch ("a" + "") {
    case "foo bar":
        v55 := GenS8{F0: 2147483647, F1: 676, F2: "\n\t"}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
