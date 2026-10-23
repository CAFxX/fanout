package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 int8 }

func main() {
    var v1 float32 = 0.0
    var v2 float32 = v1
    var v3 bool = false
    v4 := []string{(("hello" + "\n\t") + "hello"), ("" + ("foo bar" + "a"))}
    v5 := make(map[uint32]string)
    var v6 int = 128
    if v3 {
        fmt.Println(v1, v3)
        if (!(0 < 30)) {
            var v7 int32 = -653
            var v8 string = "hello"
            if v3 {
                if ("hello" != v8) {
                    v8 = (("\n\t" + ((v8 + "") + ((v8 + (v8 + v8)) + ((v8 + "hello") + (v8 + v8))))) + v8)
                }
                v7 = v7
            }
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                v1 = v2
                v10 := GenS1{F0: -1}
                _ = fi9
                _ = v10
            }
            v2 = ((v2 / ((-2.25 - v1) + v2)) + v1)
            if (916 == 32767) {
                fmt.Println(v1, v3, v6)
                if v3 {
                    var v11 float32 = v1
                    _ = v11
                } else {
                    v6 = v6
                    switch (v6 - (v6 / (((v6 >> ((4294967295) & 63))) | 1))) {
                    case 2147483647:
                        var v12 float64 = 1e10
                        _ = v12
                    default:
                        _ = len(v5)
                    }
                    _ = len(v4)
                }
            }
        }
    }
    var v13 uint16 = 1
    for fi14 := 0; fi14 < 16; fi14++ {
        if v3 {
            var v15 string = ((("a" + (("\n\t" + "\n\t") + (("\n\t" + "\n\t") + ""))) + ("\n\t" + "a")) + ("" + "\n\t"))
            if false {
                if v3 {
                    v3 = false
                    v1 -= 0.1
                } else {
                    v16 := GenS2{F0: 1, F1: 0}
                    _ = v16
                }
                var v17 float64 = 3.14159
                v15 = (v15 + v15)
                _ = v17
            }
            v3 = (!false)
            _ = v15
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                _ = len(v5)
                _ = fi18
            }
            fmt.Println(v13, v1, v1)
        }
        switch ("foo bar" + "a") {
        case "a":
            fmt.Println(fi14, v2, v2)
        case "":
            var v19 uint = 0
            _ = v19
        case "\n\t":
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v2, v1)
                v6 -= (^(fi20))
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = fi20
            }
        }
        _ = fi14
    }
    var v21 int64 = -458
    switch v6 {
    case -2147483648:
        fmt.Println(v1, v2, v6)
    case 127:
        switch "" {
        case "\n\t":
            v22 := GenS3{F0: 1, F1: 2147483648, F2: 72}
            _ = v22
        case "":
            fmt.Println(v13, v3, v3)
        case "foo bar":
            _ = len(v4)
        }
    case 178:
        if (!(v3 || (-32768 >= -32768))) {
            var v23 bool = false
            _ = v23
        } else {
            v3 = (v13 >= v13)
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(v2, v21, v2)
                _ = fi24
            }
        }
    }
    fmt.Println(v3)
    if (!v3) {
        v25 := GenS1{F0: v21}
        _ = v25
    } else {
        v21 = (v21 >> ((198) & 63))
        if (-1 == 468) {
            v26 := GenS2{F0: -1, F1: -903}
            v21 -= (v21 / ((v21) | 1))
            _ = v26
        } else {
            v4 = append(v4, ("\n\t" + ("foo bar" + "foo bar")))
            _ = len(v5)
        }
        v3 = v3
    }
    v2 = 0.0
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint8 = 2
        _ = fi27
        _ = v28
    }
    switch -9223372036854775808 {
    case 4294967295:
        if len(v4) > 0 {
            v4[0] = ""
        }
    case 1:
        if (false && ((3.14159 - 1e10) != (v2 * v1))) {
            v29 := GenS1{F0: v21}
            for fi30 := 0; fi30 < 8; fi30++ {
                if (4294967295 == 2) {
                    var v31 string = "\n\t"
                    fmt.Println(v2, v31, v21)
                    _ = v31
                }
                _ = fi30
            }
            _ = v29
        }
    case 128:
        v32 := GenS4{F0: 100, F1: 703, F2: -128}
        _ = v32
    default:
        if (v3 || false) {
            v33 := GenS2{F0: -1, F1: 397}
            _ = v33
        }
    }
    if (42 < 4294967295) {
        var v34 uint = 255
        switch (v6 % (((v6 & 4294967295)) | 1)) {
        case 32768:
            var v35 int16 = -905
            _ = v35
        case -9223372036854775808:
            for fi36 := 0; fi36 < 5; fi36++ {
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                v34 -= v34
                _ = fi36
            }
        default:
            var v37 uint = (v34 - (^(v34)))
            _ = v37
        }
        _ = v34
    }
    var v38 int = 42
    for fi39 := 0; fi39 < 1; fi39++ {
        fmt.Println(v1, v3, v6)
        var v40 float32 = 3.14159
        v41 := GenS4{F0: 127, F1: 4294967295, F2: -1}
        _ = fi39
        _ = v40
        _ = v41
    }
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
