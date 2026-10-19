package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 uint32; F2 int }

func main() {
    v1 := []string{"foo bar", "foo bar"}
    var v2 uint64 = 853
    var v3 bool = (255 != 618)
    var v4 float32 = ((1.5 / -2.25) + 3.14159)
    v5 := make(map[string]string)
    fmt.Println(v3)
    v2 = v2
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v3, fi6, v2)
        var v7 float64 = -2.25
        if ("a" >= "a") {
            fmt.Println(v7, v2)
        }
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        if v3 {
            switch (("a" + ("" + (("foo bar" + "foo bar") + "foo bar"))) + ("foo bar" + (("hello" + ("\n\t" + ("" + "hello"))) + ("hello" + "\n\t")))) {
            case "hello":
                v9 := GenS1{F0: int16(fi8)}
                _ = v9
            case "\n\t":
                switch (-(fi8)) {
                case 128:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        fmt.Println(fi10)
                        v5["foo bar"] = "foo bar"
                        switch "a" {
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v11 int64 = 9223372036854775807
                            _ = v11
                        }
                        if v3 {
                            var v12 uint64 = 1
                            v2 -= v2
                            switch (fi10 % (((^(fi8))) | 1)) {
                            case 42:
                                v3 = (!(42 != 9223372036854775807))
                            case -659:
                                var v13 uint8 = 2
                                _ = v13
                            }
                            _ = v12
                        } else {
                            if (!(18446744073709551615 >= 2)) {
                                if v3 {
                                    v5[("hello" + "")] = ((("a" + ("\n\t" + (("\n\t" + "hello") + "a"))) + "\n\t") + "foo bar")
                                    var v14 int16 = 42
                                    _ = v14
                                } else {
                                    var v15 string = ("\n\t" + ("hello" + "a"))
                                    switch (fi8 + 0) {
                                    case 4294967295:
                                        v4 = v4
                                    default:
                                        if v3 {
                                            var v16 int8 = -128
                                            v4 += float32((1e10 - 3.14159))
                                            v17 := GenS1{F0: 32767}
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                    v18 := GenS1{F0: 43}
                                    _ = v15
                                    _ = v18
                                }
                                v2 = (v2 ^ (v2 << ((57) & 63)))
                            } else {
                                switch "" {
                                case "foo bar":
                                    v3 = (fi10 < (fi10 + fi10))
                                default:
                                    switch 32767 {
                                    case -852:
                                        v19 := GenS1{F0: 1000}
                                        _ = v19
                                    case 9223372036854775807:
                                        v4 -= v4
                                    case -1:
                                        var v20 int16 = 42
                                        _ = v20
                                    default:
                                        switch (fi8 << ((18446744073709551615) & 63)) {
                                        case 42:
                                            var v21 uint8 = 1
                                            _ = v21
                                        case 921:
                                            v22 := GenS2{F0: v2, F1: v2}
                                            _ = v22
                                        case 4294967295:
                                            v23 := GenS1{F0: 0}
                                            _ = v23
                                        }
                                    }
                                }
                                v24 := GenS1{F0: -1}
                                if false {
                                    switch ("hello" + "foo bar") {
                                    case "a":
                                        v2 += 926
                                        switch "\n\t" {
                                        case "a":
                                            v2 = v2
                                        }
                                        var v25 bool = v3
                                        if (((float64(0.0) + 0.0) + 0.0) > -2.25) {
                                            var v26 int16 = 1000
                                            v26 = v26
                                            _ = v26
                                        } else {
                                            var v27 uint64 = v2
                                            _ = v27
                                        }
                                        switch "hello" {
                                        case "hello":
                                            v3 = v25
                                        default:
                                            var v28 int8 = -93
                                            _ = v28
                                        }
                                        var v29 int16 = 23
                                        fmt.Println(fi10, v3)
                                        var v30 int = 42
                                        var v31 int8 = 1
                                        _ = v25
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    case "\n\t":
                                        if (!true) {
                                            fmt.Println(v3)
                                        }
                                    }
                                    fmt.Println(fi8, fi10, fi8)
                                    if false {
                                        switch "hello" {
                                        case "":
                                            v32 := GenS1{F0: -1}
                                            _ = v32
                                        case "a":
                                            v33 := GenS3{F0: v2, F1: "a"}
                                            _ = v33
                                        case "foo bar":
                                            v5["\n\t"] = "a"
                                        }
                                    }
                                } else {
                                    v34 := GenS1{F0: 1000}
                                    delete(v5, ("hello" + "foo bar"))
                                    if v3 {
                                        switch (fi8 & 2147483647) {
                                        case 1:
                                            v5[""] = "foo bar"
                                        default:
                                            v3 = false
                                        }
                                        fmt.Println(v3, fi8)
                                        switch (fi8 / ((fi10) | 1)) {
                                        case 175:
                                            fmt.Println(fi10, v3)
                                        case -9223372036854775808:
                                            switch (^(fi8)) {
                                            case 32767:
                                                var v35 int8 = 100
                                                _ = v35
                                            case 154:
                                                if (0.1 == (float64(v4) - (1.5 * float64((v4 * v4))))) {
                                                    var v36 uint8 = 91
                                                    _ = v36
                                                }
                                            default:
                                                var v37 uint8 = 1
                                                _ = v37
                                            }
                                        default:
                                            var v38 uint32 = 42
                                            _ = v38
                                        }
                                    }
                                    _ = v34
                                }
                                _ = v24
                            }
                            v2 += v2
                        }
                        var v39 int16 = 32767
                        _ = fi10
                        _ = v39
                    }
                case 70:
                    var v40 uint = 2
                    _ = v40
                }
            }
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 uint = 4294967295
                fmt.Println(fi8, v2, v2)
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = fi41
                _ = v42
            }
        }
        var v43 uint8 = 42
        _ = fi8
        _ = v43
    }
    var v44 float32 = 0.0
    for fi45 := 0; fi45 < 16; fi45++ {
        v46 := GenS4{F0: 42, F1: 416}
        v3 = (!(v44 == v4))
        _ = fi45
        _ = v46
    }
    v47 := GenS5{F0: 42, F1: 0, F2: -1}
    var v48 uint16 = 255
    var v49 float64 = (0.0 + 0.1)
    if ((^(v48)) <= (v48 / ((2) | 1))) {
        v49 -= v49
    } else {
        fmt.Println(v2, v3)
    }
    v50 := GenS6{F0: 1, F1: 536, F2: 871}
    v51 := v5["\n\t"]
    fmt.Println(v3, v49)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
