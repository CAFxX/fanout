package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = -321
    var v2 uint16 = 2
    v3 := []uint32{2, 2147483647, 164, uint32(v2)}
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 0, F1: ("hello" + "\n\t"), F2: ("foo bar" + "\n\t")}
        if false {
            switch -872 {
            case -408:
                var v6 uint16 = 2
                _ = v6
            case 9223372036854775807:
                var v7 int = fi4
                _ = v7
            default:
                if (false || false) {
                    if false {
                        var v8 bool = false
                        if v8 {
                            _ = len(v3)
                            switch "\n\t" {
                            case "foo bar":
                                v8 = (("hello" + (("foo bar" + "") + ("\n\t" + "foo bar"))) >= (((("a" + "a") + ("" + "hello")) + "") + "a"))
                            case "hello":
                                fmt.Println(v1, fi4)
                            case "":
                                v1 = v1
                            }
                        } else {
                            var v9 bool = (2147483647 <= 42)
                            _ = v9
                        }
                        v1 = int64(v2)
                        _ = v8
                    }
                } else {
                    var v10 uint32 = 65535
                    var v11 int16 = -32768
                    switch (("" + "") + ("\n\t" + "foo bar")) {
                    case "hello":
                        switch fi4 {
                        case 32767:
                            for fi12 := 0; fi12 < 14; fi12++ {
                                if (-775 <= -234) {
                                    v2 -= (v2 - 1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    v1 = (^(v1))
                                    fmt.Println(v1, v10, v1)
                                    v1 = int64(v10)
                                } else {
                                    v13 := GenS1{F0: 4, F1: ("hello" + ("" + ("" + ((("hello" + "foo bar") + "a") + "")))), F2: "\n\t"}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v13
                                }
                                _ = fi12
                            }
                        case -1:
                            var v14 float64 = 0.0
                            _ = v14
                        }
                    case "foo bar":
                        if false {
                            for fi15 := 0; fi15 < 11; fi15++ {
                                v2 = 2
                                _ = fi15
                            }
                            for fi16 := 0; fi16 < 20; fi16++ {
                                var v17 uint16 = 869
                                _ = fi16
                                _ = v17
                            }
                            var v18 int8 = -96
                            _ = v18
                        }
                    default:
                        switch (("a" + ("hello" + "foo bar")) + "foo bar") {
                        case "a":
                            _ = len(v3)
                        default:
                            v10 = v10
                        }
                    }
                    _ = v10
                    _ = v11
                }
            }
            var v19 float32 = float32(((((1.5 / 1.0) - 1.5) / 0.1) + ((1e10 * -2.25) + 1.5)))
            v20 := GenS1{F0: 127, F1: "hello", F2: ("a" + ((("\n\t" + ("a" + "\n\t")) + "a") + "a"))}
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v2 = v2
            v3 = append(v3, 783)
            fmt.Println(fi21)
            _ = fi21
        }
        _ = fi4
        _ = v5
    }
    if (false || false) {
        var v22 float32 = -2.25
        _ = v22
    } else {
        v2 = (v2 << ((4294967295) & 63))
        v1 += (-(v1))
        v2 += 65535
    }
    if false {
        fmt.Println(v1)
        v2 = (v2 - (v2 + v2))
        v23 := GenS1{F0: 42, F1: (("\n\t" + "\n\t") + "a"), F2: ("a" + "")}
        _ = v23
    }
    switch 1 {
    case -9223372036854775808:
        _ = len(v3)
    case -1:
        fmt.Println(v1, v2)
    case 32768:
        if ((-276 <= 0) || true) {
            v1 = (v1 - v1)
        } else {
            if (((("foo bar" + "hello") + "foo bar") + ("foo bar" + "foo bar")) > "\n\t") {
                var v24 int8 = int8(v2)
                fmt.Println(v1)
                _ = v24
            }
        }
    }
    v1 = 686
    var v25 uint64 = 18446744073709551615
    v26 := GenS2{F0: 42}
    v3 = append(v3, 4294967295)
    var v27 uint8 = 2
    v27 -= (v27 >> ((123) & 63))
    v1 += (v1 - v1)
    v25 += (v25 ^ (v25 ^ 4294967295))
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int8 = 0
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 int16 = -544
        _ = fi30
        _ = v31
    }
    var v32 int8 = 1
    if false {
        var v33 uint32 = 65535
        v34 := GenS3{F0: v1}
        _ = v33
        _ = v34
    } else {
        fmt.Println(v32, v27, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
}
