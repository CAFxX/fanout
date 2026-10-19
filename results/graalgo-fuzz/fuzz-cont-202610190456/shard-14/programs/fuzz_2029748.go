package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8 }
func genFn0() int64 {
    var v6 int = -360
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = ("foo bar" + "\n\t")
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi9)
            var v10 uint32 = 2147483647
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 18446744073709551615, F1: 2, F2: -1}
    v6 = 128
    v6 += (v6 << ((818) & 63))
    _ = v11
    _ = v6
    return -13
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = v1
    var v3 int64 = 0
    v4 := []string{"\n\t", ("foo bar" + "hello"), "a", ("" + "\n\t")}
    var v5 uint32 = 1
    var v12 uint8 = v1
    switch "a" {
    case "":
        if (2 > 547) {
            var v13 int64 = v3
            var v14 string = ("foo bar" + (("" + "hello") + ""))
            _ = v13
            _ = v14
        }
    case "foo bar":
        var v15 int8 = 0
        _ = v15
    default:
        v12 += (v2 * (v1 / (((v2 << ((0) & 63))) | 1)))
    }
    v5 = 0
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int8 = 127
        _ = fi16
        _ = v17
    }
    switch 127 {
    case 671:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            switch "\n\t" {
            case "":
                switch 127 {
                case 1:
                    fmt.Println(v2, v12)
                default:
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                }
                switch fi18 {
                case -1:
                    if len(v4) > 0 {
                        v4[0] = (("foo bar" + "a") + "a")
                    }
                    if false {
                        v4 = append(v4, "foo bar")
                        if false {
                            v12 += v2
                            v5 = v5
                            v3 = (^(v3))
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v20 float32 = 1.5
                        var v21 bool = false
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int16 = 32767
                        if false {
                            v3 = v3
                            v12 = (v1 << ((42) & 63))
                        }
                        _ = v22
                    }
                    _ = len(v4)
                    if (253 != 1) {
                        fmt.Println(v2)
                    } else {
                        var v23 int16 = 1000
                        switch "hello" {
                        case "\n\t":
                            var v24 int32 = 0
                            _ = v24
                        case "a":
                            v5 = v5
                        }
                        _ = v23
                    }
                    v25 := GenS1{F0: 42, F1: 42, F2: 100}
                    _ = v25
                case 0:
                    _ = len(v4)
                default:
                    var v26 int = fi18
                    _ = v26
                }
                if (genFn0() < v3) {
                    v3 += int64(v5)
                } else {
                    var v27 int32 = 42
                    var v28 uint32 = v5
                    v28 += 2
                    _ = v27
                    _ = v28
                }
                fmt.Println(v3)
                if (79 != 0) {
                    var v29 int32 = -514
                    _ = v29
                }
            case "hello":
                var v30 uint16 = 42
                _ = v30
            }
            v3 += (v3 ^ 839)
            switch "" {
            case "hello":
                var v31 string = ""
                _ = v31
            default:
                if (65535 < 94) {
                    fmt.Println(v3, v1)
                    v3 = v3
                }
            }
            _ = fi19
        }
        _ = fi18
    }
    var v32 string = ("foo bar" + ("foo bar" + "hello"))
    var v33 int16 = -940
    if false {
        v34 := GenS2{F0: 42}
        fmt.Println(v5, v2, v5)
        _ = v34
    } else {
        for fi35 := 0; fi35 < 11; fi35++ {
            switch v32 {
            case "hello":
                v36 := GenS1{F0: 1, F1: 1, F2: 42}
                _ = v36
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi35
        }
        var v37 int = -2147483648
        var v38 uint16 = 32767
        _ = v37
        _ = v38
    }
    v39 := GenS1{F0: 2, F1: 0, F2: 42}
    if true {
        v4 = append(v4, (v32 + (("foo bar" + v32) + "")))
        v33 = (v33 % ((v33) | 1))
        fmt.Println(v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
