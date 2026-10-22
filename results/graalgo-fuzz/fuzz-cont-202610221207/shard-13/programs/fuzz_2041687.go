package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint32; F2 string }

func main() {
    v1 := []string{}
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    var v5 uint16 = 267
    var v6 int16 = 432
    delete(v3, "a")
    fmt.Println(v2, v5, v4)
    if true {
        var v7 uint8 = 127
        if (!(9223372036854775807 <= 1)) {
            fmt.Println(v4)
        } else {
            v8 := GenS1{F0: v5, F1: (^(v2)), F2: v7}
            _ = v8
        }
        var v9 float64 = 3.14159
        _ = v7
        _ = v9
    }
    v2 = 4294967295
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 float64 = -2.25
        v1 = append(v1, ("\n\t" + ""))
        _ = fi10
        _ = v11
    }
    if false {
        fmt.Println(v2)
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 -= v2
            if (false || (true && true)) {
                var v13 uint = 4294967295
                _ = v13
            } else {
                v14 := v3[("foo bar" + (((("foo bar" + (((((((("hello" + (("\n\t" + ("" + "hello")) + ("" + ("" + ("\n\t" + (("foo bar" + "\n\t") + ("foo bar" + "a"))))))) + "hello") + "\n\t") + "") + (((("\n\t" + ("\n\t" + ("" + "hello"))) + "hello") + "\n\t") + ("\n\t" + "a"))) + ("hello" + ("foo bar" + "a"))) + "\n\t") + "a")) + (((("" + "hello") + "hello") + "") + "foo bar")) + ("a" + "a")) + ("a" + ("\n\t" + "a"))))]
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 float32 = (1.5 / float32(1e10))
                    delete(v3, (("a" + "") + ((("\n\t" + "") + "foo bar") + (("hello" + ("foo bar" + ((((("foo bar" + ((("foo bar" + "") + "hello") + ("foo bar" + "\n\t"))) + (((("hello" + ("hello" + ("hello" + ("\n\t" + "\n\t")))) + "") + "foo bar") + "hello")) + "") + "foo bar") + "\n\t"))) + "foo bar"))))
                    var v17 int8 = 104
                    v6 -= (v6 / (((^(v6))) | 1))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v18 := GenS1{F0: (v5 | (v5 & 0)), F1: (v2 - (v14 * uint32(v5))), F2: 0}
                _ = v14
                _ = v18
            }
            var v19 uint64 = 0
            _ = fi12
            _ = v19
        }
        if (false || false) {
            v4 = v4
        } else {
            v20 := GenS2{F0: (v6 % ((v6) | 1))}
            var v21 uint16 = v5
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int8 = 100
                fmt.Println(v4, v5, v6)
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        }
    }
    v2 = (v2 - v2)
    v6 -= 0
    v6 -= (v6 << ((343) & 63))
    var v24 uint32 = (v2 ^ (^(v2)))
    for fi25 := 0; fi25 < 5; fi25++ {
        switch 631 {
        case -1:
            fmt.Println(fi25, v5, v5)
        case -2147483648:
            if ((false && false) || false) {
                v6 = v6
                switch fi25 {
                case 32767:
                    v24 -= v24
                default:
                    v26 := GenS3{F0: (v6 >> ((220) & 63)), F1: (v24 ^ (v2 / ((v24) | 1))), F2: (("foo bar" + "hello") + "foo bar")}
                    _ = v26
                }
            } else {
                var v27 string = "foo bar"
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(v24, v24, v4)
                    v5 += (v5 >> ((4294967295) & 63))
                    _ = fi28
                }
                _ = v27
            }
        default:
            var v29 int16 = (v6 << ((940) & 63))
            _ = v29
        }
        v30 := GenS1{F0: 893, F1: 0, F2: 0}
        fmt.Println(v24, v2, v5)
        _ = fi25
        _ = v30
    }
    var v31 uint64 = 42
    if (0.1 > 1e10) {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint64 = 0
            var v34 int = fi32
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 int16 = v6
        _ = v35
    } else {
        switch "hello" {
        case "":
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 bool = (false && true)
                _ = fi36
                _ = v37
            }
        case "foo bar":
            var v38 int32 = -2147483648
            _ = v38
        case "\n\t":
            v2 += (v24 / (((^(v24))) | 1))
        default:
            var v39 float32 = 1.5
            _ = v39
        }
        _ = len(v1)
        v24 += 2147483648
    }
    v40 := GenS3{F0: v6, F1: v2, F2: ((("" + "foo bar") + ((("" + "\n\t") + (("a" + "\n\t") + "\n\t")) + "\n\t")) + "foo bar")}
    v24 = (v24 >> ((18446744073709551615) & 63))
    var v41 uint64 = v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
