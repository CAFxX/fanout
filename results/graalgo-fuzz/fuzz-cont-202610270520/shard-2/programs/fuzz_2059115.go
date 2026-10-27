package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8; F1 bool; F2 int16 }
type GenS6 struct { F0 int; F1 uint32; F2 bool }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 int64 = -635
    if false {
        var v4 uint8 = 1
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 int8 = -1
            v6 += 36
            _ = fi5
            _ = v6
        }
        _ = v4
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            if (!(((0.1 / 3.14159) + (1.5 - 1.5)) < ((-2.25 - (1e10 / 1.5)) / 1.0))) {
                var v8 int = (fi7 | (fi7 + (fi7 >> ((2) & 63))))
                if (4294967295 != 0) {
                    var v9 uint32 = 65535
                    if false {
                        var v10 float64 = 0.0
                        v10 = (0.1 - v10)
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                switch ("hello" + ("" + ("foo bar" + ("\n\t" + ("hello" + ("\n\t" + "")))))) {
                                case "a":
                                    if (2147483647 <= -2147483648) {
                                        v9 += v9
                                        v3 -= 382
                                        var v11 int16 = 106
                                        _ = v11
                                    }
                                case "\n\t":
                                    v8 = 2147483647
                                }
                            case "a":
                                var v12 uint16 = 0
                                _ = v12
                            case "":
                                for fi13 := 0; fi13 < 7; fi13++ {
                                    fmt.Println(v10, fi13)
                                    v14 := v1["hello"]
                                    var v15 int64 = v3
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                            v16 := GenS1{F0: 1}
                            v10 = v10
                            _ = v16
                        }
                        _ = v10
                    }
                    v2 = (v2 * v2)
                    _ = v9
                } else {
                    v1["a"] = 1
                    if ((fi7 >> ((4294967295) & 63)) > (-(fi7))) {
                        if (-60 > 127) {
                            var v17 string = ("foo bar" + (("foo bar" + "hello") + ""))
                            _ = v17
                        }
                        v18 := GenS2{F0: 1000, F1: (!true), F2: 1}
                        _ = v18
                    } else {
                        v19 := GenS3{F0: ((v2 * v2) < 719), F1: int8(v3), F2: ("" + ("hello" + "hello"))}
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    switch 32768 {
                    case 9223372036854775807:
                        var v21 uint16 = 65535
                        _ = v21
                    case 4294967295:
                        delete(v1, "a")
                    }
                    _ = fi20
                }
                _ = v8
            } else {
                v2 -= (v2 << ((2) & 63))
            }
            switch -1 {
            case -676:
                switch "hello" {
                case "foo bar":
                    var v22 int64 = 1
                    _ = v22
                case "a":
                    v3 = -9223372036854775808
                default:
                    v3 += (v3 & v3)
                }
            case 0:
                var v23 uint32 = 2147483648
                _ = v23
            case 42:
                v3 = v3
            default:
                if ((((0.0 / 1.5) / 1e10) * 0.0) > (1.5 + float32(0.0))) {
                    switch ("" + "hello") {
                    case "a":
                        v3 += (v3 << ((uint(v2)) & 63))
                    }
                    switch (fi7 - fi7) {
                    case 1:
                        v2 = (v2 ^ (v2 + (v2 % (((v2 & (v2 * (^(v2))))) | 1))))
                    }
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 uint = 636
                        v26 := GenS4{F0: 991, F1: (^(v3))}
                        if (32767 > 0) {
                            var v27 float64 = 3.14159
                            v2 = (v2 + 42)
                            fmt.Println(v2, v3)
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                } else {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        v29 := v1[""]
                        var v30 string = ("\n\t" + "\n\t")
                        v31 := GenS1{F0: int32(v3)}
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
            }
            fmt.Println(v2, v3)
            _ = fi7
        }
        fmt.Println(v2, v3)
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := GenS5{F0: 127, F1: false, F2: 32767}
        fmt.Println(v3)
        _ = fi32
        _ = v33
    }
    switch "hello" {
    case "\n\t":
        var v34 uint8 = 0
        _ = v34
    case "foo bar":
        var v35 uint16 = 32767
        _ = v35
    case "":
        var v36 int16 = -588
        _ = v36
    }
    delete(v1, "foo bar")
    switch "\n\t" {
    case "hello":
        var v37 bool = (494 > 1)
        _ = v37
    case "":
        switch -421 {
        case -9223372036854775808:
            var v38 int8 = 0
            _ = v38
        }
    }
    v39 := GenS6{F0: 128, F1: 1, F2: true}
    var v40 bool = (v3 < 42)
    v41 := GenS7{F0: ""}
    var v42 string = "foo bar"
    var v43 int8 = 127
    var v44 int = -2147483648
    v43 = (-(v43))
    fmt.Println(v43)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
