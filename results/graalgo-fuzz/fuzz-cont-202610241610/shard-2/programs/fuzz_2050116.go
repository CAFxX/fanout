package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    v1 := []string{"foo bar", "", (("\n\t" + "hello") + ("a" + ("foo bar" + ""))), ""}
    v2 := []uint32{332, 4294967295, 42, 221}
    v3 := make(map[uint32]uint32)
    var v4 uint = 42
    v5 := []string{"", "foo bar", "a"}
    var v6 uint32 = 2147483648
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v4)
    v7 := GenS1{F0: "\n\t", F1: -32768, F2: 127}
    v6 = 0
    if (false && false) {
        v6 = 4294967295
        v4 -= (v4 >> ((42) & 63))
        if true {
            switch -620 {
            case -1:
                v8 := GenS1{F0: "a", F1: 0, F2: 236}
                _ = v8
            case 1:
                for fi9 := 0; fi9 < 15; fi9++ {
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            fmt.Println(fi9)
                            if false {
                                if (fi9 <= (fi9 | fi9)) {
                                    var v10 int = (-(fi9))
                                    _ = v10
                                } else {
                                    switch fi9 {
                                    case -1:
                                        fmt.Println(v6, fi9)
                                    case 32768:
                                        if true {
                                            if false {
                                                var v11 int32 = 1
                                                for fi12 := 0; fi12 < 15; fi12++ {
                                                    var v13 int16 = -1
                                                    fmt.Println(v4)
                                                    var v14 bool = (false || true)
                                                    _ = fi12
                                                    _ = v13
                                                    _ = v14
                                                }
                                                _ = v11
                                            }
                                            _ = len(v5)
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                var v15 uint32 = (v6 / ((4294967295) | 1))
                                                _ = v15
                                            case "\n\t":
                                                var v16 uint32 = (v6 * v6)
                                                _ = v16
                                            default:
                                                if (0.0 == 0.0) {
                                                    if true {
                                                        switch ("hello" + "\n\t") {
                                                        case "a":
                                                            v4 -= v4
                                                        case "hello":
                                                            v17 := GenS1{F0: "foo bar", F1: 1, F2: 226}
                                                            _ = v17
                                                        }
                                                    } else {
                                                        v18 := GenS2{F0: true, F1: (^(v4))}
                                                        _ = v18
                                                    }
                                                    _ = len(v3)
                                                    v4 = 1
                                                } else {
                                                    _ = len(v3)
                                                    var v19 uint16 = 2
                                                    _ = v19
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            var v20 string = "\n\t"
                            v6 = (v6 - (^(v6)))
                            _ = v20
                        }
                    case "a":
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                    }
                    if true {
                        if false {
                            var v21 int64 = -9223372036854775808
                            _ = v21
                        } else {
                            fmt.Println(fi9, v4)
                            if false {
                                v6 += 65535
                                v6 -= (^(v6))
                                v4 += (v4 << ((255) & 63))
                            } else {
                                v6 -= v6
                                switch "\n\t" {
                                case "foo bar":
                                    v6 = 568
                                case "hello":
                                    v4 = 42
                                case "\n\t":
                                    fmt.Println(v4, fi9)
                                default:
                                    var v22 int8 = 1
                                    _ = v22
                                }
                            }
                        }
                        fmt.Println(fi9, v6, fi9)
                        fmt.Println(v6)
                    } else {
                        v23 := GenS2{F0: (!(!(((("foo bar" + "") + "\n\t") + "foo bar") != "\n\t"))), F1: v4}
                        _ = len(v3)
                        _ = v23
                    }
                    _ = fi9
                }
            case 9223372036854775807:
                var v24 int = 42
                _ = v24
            default:
                for fi25 := 0; fi25 < 11; fi25++ {
                    switch "" {
                    case "foo bar":
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v27 := GenS1{F0: "", F1: -1, F2: 42}
                            switch "hello" {
                            case "a":
                                var v28 uint64 = 424
                                _ = v28
                            default:
                                v3[v6] = v6
                            }
                            var v29 uint = 42
                            _ = fi26
                            _ = v27
                            _ = v29
                        }
                    default:
                        if true {
                            var v30 uint32 = (v6 | (v6 & uint32(v4)))
                            _ = v30
                        }
                    }
                    v31 := GenS1{F0: (("" + "\n\t") + (("\n\t" + "a") + ("hello" + ("a" + ("hello" + ("a" + ("a" + "hello"))))))), F1: -1, F2: 0}
                    _ = fi25
                    _ = v31
                }
            }
            var v32 uint = v4
            v33 := GenS1{F0: "", F1: 1, F2: 141}
            _ = v32
            _ = v33
        } else {
            if true {
                fmt.Println(v4, v4)
                var v34 float64 = 1e10
                _ = v34
            }
            var v35 bool = true
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v4 -= (v4 - (v4 * 2))
        _ = fi36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 float32 = 1.5
        var v39 float32 = v38
        _ = fi37
        _ = v38
        _ = v39
    }
    switch 4294967295 {
    case 4294967295:
        fmt.Println(v4, v4)
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS2{F0: (true && (((-2.25 * 0.1) / 3.14159) != 3.14159)), F1: 42}
        _ = fi40
        _ = v41
    }
    v42 := GenS2{F0: (1 >= 0), F1: v4}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
