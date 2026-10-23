package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16; F1 int8; F2 uint64 }

func main() {
    v1 := []string{}
    var v2 int = -1
    var v3 bool = true
    var v4 int16 = -306
    var v5 float64 = (1e10 / 3.14159)
    v2 = (v2 + 2147483647)
    var v6 uint16 = 65535
    fmt.Println(v6, v5, v3)
    if ((!v3) || true) {
        if ((!(!true)) && (329 <= 260)) {
            switch "foo bar" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                fmt.Println(v2, v3)
            case "foo bar":
                var v7 float32 = -2.25
                _ = v7
            default:
                v1 = append(v1, ("" + "foo bar"))
            }
            switch "a" {
            case "foo bar":
                if true {
                    var v8 int64 = 9223372036854775807
                    _ = v8
                } else {
                    v1 = append(v1, "")
                    v9 := GenS1{F0: 1, F1: 0}
                    _ = v9
                }
            case "a":
                switch (v2 / (((-(v2))) | 1)) {
                case 127:
                    switch (-(v2)) {
                    case 32767:
                        switch (v2 + (v2 >> ((885) & 63))) {
                        case 2147483647:
                            var v10 bool = (!true)
                            _ = v10
                        case 9223372036854775807:
                            for fi11 := 0; fi11 < 5; fi11++ {
                                for fi12 := 0; fi12 < 18; fi12++ {
                                    v13 := GenS2{F0: (fi12 & (fi12 ^ fi12)), F1: 142, F2: -9223372036854775808}
                                    var v14 uint8 = 127
                                    var v15 string = ""
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                                _ = fi11
                            }
                        }
                    }
                case 477:
                    fmt.Println(v6, v5)
                case -2147483648:
                    if (!true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v16 := GenS2{F0: (v2 >> ((1) & 63)), F1: 1, F2: 1}
                        _ = v16
                    } else {
                        v17 := GenS3{F0: 42, F1: 0, F2: -589}
                        var v18 string = "foo bar"
                        if (true && (false || (0.1 < 0.0))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v6 += (^(v6))
                        }
                        _ = v17
                        _ = v18
                    }
                }
            }
        } else {
            v19 := GenS4{F0: 65535, F1: 0}
            if (v3 && v3) {
                v20 := GenS5{F0: v4, F1: -1, F2: 18446744073709551615}
                _ = v20
            }
            _ = v19
        }
    }
    switch "a" {
    case "foo bar":
        if (2147483647 == 1) {
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 int64 = 935
                fmt.Println(v5, v22, v4)
                v23 := GenS5{F0: (v4 + 189), F1: -60, F2: 6}
                _ = fi21
                _ = v22
                _ = v23
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
    case "hello":
        var v24 uint32 = 42
        _ = v24
    case "":
        var v25 int = 9223372036854775807
        _ = v25
    }
    if true {
        fmt.Println(v6, v4, v5)
    }
    switch ("a" + ((("\n\t" + ((("a" + ("hello" + "a")) + "foo bar") + ("hello" + ("\n\t" + "")))) + "a") + "\n\t")) {
    case "":
        var v26 int8 = 42
        _ = v26
    default:
        if false {
            if (false || false) {
                v6 = (v6 >> ((0) & 63))
                v1 = append(v1, (("\n\t" + ("hello" + "")) + "foo bar"))
            } else {
                var v27 uint = 960
                _ = v27
            }
            v28 := GenS3{F0: 349, F1: 131, F2: -9223372036854775808}
            var v29 uint32 = 838
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v5 += 0.1
        _ = fi30
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v1 = append(v1, (("hello" + "") + ""))
        _ = fi31
    }
    var v32 string = (("a" + ((("foo bar" + (("\n\t" + "") + "a")) + (("a" + ("foo bar" + "hello")) + "\n\t")) + "\n\t")) + ("a" + "hello"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
