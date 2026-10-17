package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{459, 2147483647, 4294967295}
    var v3 int = 95
    var v4 float64 = float64(0.0)
    if (127 < -128) {
        for fi5 := 0; fi5 < 19; fi5++ {
            _ = len(v1)
            _ = fi5
        }
        var v6 uint16 = 490
        _ = v6
    } else {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (-112 >= -1) {
                fmt.Println(v4, v4)
                fmt.Println(v3, v3)
                if false {
                    fmt.Println(v3, v3)
                } else {
                    v7 := v1[455]
                    switch "foo bar" {
                    case "":
                        for fi8 := 0; fi8 < 12; fi8++ {
                            v4 += v4
                            if (!true) {
                                var v9 uint64 = 300
                                _ = v9
                            } else {
                                var v10 uint16 = 42
                                _ = v10
                            }
                            var v11 int8 = 1
                            _ = fi8
                            _ = v11
                        }
                    case "foo bar":
                        var v12 uint8 = 42
                        _ = v12
                    case "\n\t":
                        for fi13 := 0; fi13 < 15; fi13++ {
                            if ("" >= "a") {
                                v1[(v7 ^ 884)] = (v7 >> ((42) & 63))
                                delete(v1, 1)
                            } else {
                                v14 := GenS1{F0: "foo bar"}
                                switch "" {
                                case "":
                                    v15 := GenS1{F0: ""}
                                    _ = v15
                                case "hello":
                                    v7 = v7
                                case "foo bar":
                                    var v16 string = (("hello" + "") + "")
                                    _ = v16
                                default:
                                    fmt.Println(fi13)
                                }
                                _ = v14
                            }
                            switch "\n\t" {
                            case "":
                                var v17 int16 = 441
                                _ = v17
                            case "foo bar":
                                v18 := v1[v7]
                                _ = v18
                            default:
                                v19 := GenS1{F0: "a"}
                                _ = v19
                            }
                            var v20 string = ("\n\t" + ("\n\t" + (("foo bar" + ("\n\t" + "a")) + "foo bar")))
                            _ = fi13
                            _ = v20
                        }
                    default:
                        if (18446744073709551615 < 282) {
                            var v21 int32 = -2147483648
                            switch ((("foo bar" + (((("a" + "\n\t") + "foo bar") + (((("hello" + "a") + "foo bar") + ("hello" + ((("\n\t" + ("foo bar" + ("a" + ""))) + "a") + (("a" + ("hello" + "hello")) + (("foo bar" + "") + ("foo bar" + "hello")))))) + ("a" + "hello"))) + "a")) + "hello") + "a") {
                            case "foo bar":
                                fmt.Println(v3)
                            default:
                                v22 := GenS1{F0: ("\n\t" + "foo bar")}
                                _ = v22
                            }
                            _ = v21
                        }
                    }
                    _ = v7
                }
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 uint16 = 42
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
            }
        } else {
            fmt.Println(v3, v4)
        }
    }
    v1[42] = 2
    var v26 int32 = -1
    var v27 int16 = 1
    v26 -= (^(v26))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int64 = -9223372036854775808
    delete(v1, 154)
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 uint32 = 866
        var v31 uint64 = 4294967295
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint16 = 2
            v28 -= (v28 / ((v28) | 1))
            _ = fi32
            _ = v33
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    var v34 int32 = 1
    switch "a" {
    case "":
        var v35 float64 = v4
        _ = v35
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi36 := 0; fi36 < 1; fi36++ {
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 uint16 = 32767
                v34 += v26
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    if (!true) {
        var v39 int = (v3 | 358)
        v27 = (v27 % (((v27 / ((v27) | 1))) | 1))
        v2 = append(v2, 4294967295)
        _ = v39
    } else {
        v3 -= (v3 + (v3 >> ((2) & 63)))
        v1[2] = 876
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}
