package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]bool)
    v3 := []string{"\n\t"}
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 uint16 = 42
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v7 int16 = v1
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(fi8)
        delete(v5, "foo bar")
        _ = fi8
    }
    v3 = append(v3, ("foo bar" + (("\n\t" + "\n\t") + ("" + ("a" + "")))))
    switch "\n\t" {
    case "\n\t":
        v9 := GenS1{F0: "", F1: (v6 >> ((4294967295) & 63))}
        _ = v9
    case "a":
        v6 -= v6
    case "hello":
        if (true && false) {
            v6 = v6
            v7 = (v7 - (v7 >> ((42) & 63)))
            v10 := GenS1{F0: "", F1: 571}
            _ = v10
        }
    }
    v6 = (v6 & 753)
    switch 0 {
    case 2147483647:
        v1 = 0
    case -1:
        var v11 int32 = 42
        _ = v11
    default:
        if ("a" < ("hello" + "foo bar")) {
            v4["foo bar"] = 32768
        } else {
            if (0 <= 1) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 int = fi12
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                fmt.Println(v6, v1)
                if (0.0 < float32(0.0)) {
                    v1 = v7
                    if (!false) {
                        var v15 uint32 = 177
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v15 += v15
                            switch 32768 {
                            case 128:
                                if (v7 != 0) {
                                    if false {
                                        var v17 int8 = -1
                                        _ = v17
                                    } else {
                                        delete(v5, ("a" + ("hello" + (("foo bar" + "a") + ""))))
                                        var v18 int = -747
                                        fmt.Println(fi16, v6, v18)
                                        _ = v18
                                    }
                                    if ((^(v15)) < (v15 | 42)) {
                                        fmt.Println(v7, v6, fi16)
                                        fmt.Println(v15, fi16)
                                    } else {
                                        var v19 uint16 = v6
                                        _ = v19
                                    }
                                    switch (("hello" + "hello") + "\n\t") {
                                    case "":
                                        fmt.Println(v15, v7)
                                    case "hello":
                                        v2["foo bar"] = (!(true || (float32(0.1) <= 0.1)))
                                    }
                                }
                            case 127:
                                v15 = v15
                            default:
                                v20 := GenS2{F0: -128, F1: 1}
                                _ = v20
                            }
                            _ = fi16
                        }
                        v7 = (v7 >> ((65535) & 63))
                        _ = v15
                    } else {
                        if (4294967295 >= 1) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v5["\n\t"] = (!(!true))
                            v21 := v5["\n\t"]
                            _ = v21
                        }
                    }
                    v6 -= v6
                } else {
                    if true {
                        if (-821 > -9223372036854775808) {
                            switch 9223372036854775807 {
                            case 0:
                                if true {
                                    var v22 string = "\n\t"
                                    fmt.Println(v22, v22, v6)
                                    _ = v22
                                } else {
                                    v6 = (^(v6))
                                    for fi23 := 0; fi23 < 20; fi23++ {
                                        if (-962 != (v7 & v1)) {
                                            for fi24 := 0; fi24 < 13; fi24++ {
                                                switch ("foo bar" + "") {
                                                case "foo bar":
                                                    v7 = -1
                                                case "hello":
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                }
                                                if true {
                                                    v25 := GenS1{F0: "\n\t", F1: v6}
                                                    if true {
                                                        v2["hello"] = ("\n\t" > (("\n\t" + "\n\t") + ((("foo bar" + "a") + "") + "a")))
                                                    }
                                                    v26 := GenS3{F0: 4294967295, F1: 255, F2: 4294967295}
                                                    _ = v25
                                                    _ = v26
                                                }
                                                var v27 int64 = -9223372036854775808
                                                _ = fi24
                                                _ = v27
                                            }
                                            if (130 != 104) {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi23, fi23)
                                                case "a":
                                                    v28 := GenS3{F0: 879, F1: 42, F2: 1}
                                                    _ = v28
                                                case "":
                                                    var v29 uint64 = 42
                                                    _ = v29
                                                default:
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(fi23)
                                                    case "\n\t":
                                                        v7 += (v7 ^ 0)
                                                    case "hello":
                                                        v30 := GenS4{F0: 9223372036854775807}
                                                        _ = v30
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                }
                                                fmt.Println(v6, fi23)
                                                v6 = 255
                                            }
                                        }
                                        _ = fi23
                                    }
                                }
                            default:
                                v31 := GenS2{F0: 127, F1: 168}
                                _ = v31
                            }
                        }
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v7 = v7
                            _ = fi32
                        }
                        switch "foo bar" {
                        case "hello":
                            for fi33 := 0; fi33 < 13; fi33++ {
                                switch "foo bar" {
                                case "foo bar":
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        v4["a"] = fi34
                                        _ = fi34
                                    }
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                }
                                _ = fi33
                            }
                        case "a":
                            v35 := GenS3{F0: 42, F1: 1, F2: 18446744073709551615}
                            _ = v35
                        case "\n\t":
                            v5["foo bar"] = true
                        }
                    } else {
                        v1 = v1
                        v6 += (^(v6))
                    }
                    var v36 uint16 = (v6 | 42)
                    v37 := GenS2{F0: int8(v7), F1: 42}
                    _ = v36
                    _ = v37
                }
            } else {
                var v38 int64 = -285
                _ = v38
            }
            switch "" {
            case "hello":
                v7 += (v7 + 42)
            default:
                v7 = (v1 - v1)
            }
            for fi39 := 0; fi39 < 6; fi39++ {
                for fi40 := 0; fi40 < 7; fi40++ {
                    var v41 int = (^(fi40))
                    _ = fi40
                    _ = v41
                }
                v7 += v1
                v42 := GenS3{F0: 1, F1: 75, F2: 140}
                _ = fi39
                _ = v42
            }
        }
    }
    v1 = (v1 & (v1 >> ((4294967295) & 63)))
    v7 = -32768
    for fi43 := 0; fi43 < 2; fi43++ {
        v1 = -32768
        v44 := GenS2{F0: 127, F1: 42}
        fmt.Println(v6, v6)
        _ = fi43
        _ = v44
    }
    v6 += (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
