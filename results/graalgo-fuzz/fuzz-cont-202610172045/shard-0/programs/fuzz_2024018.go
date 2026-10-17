package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 bool; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
func genFn0(p0 int32) string {
    if ("" <= "") {
        v5 := GenS1{F0: -1, F1: 4294967295}
        _ = v5
    } else {
        p0 = (p0 | -2147483648)
        if (42 <= 1) {
            p0 -= p0
            p0 -= (p0 ^ (p0 + p0))
        }
        p0 -= (p0 << ((255) & 63))
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p0)
        _ = fi6
    }
    _ = p0
    return "a"
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    v2 := make(map[uint32]string)
    v3 := []string{(v1 + v1)}
    v4 := make(map[string]bool)
    v1 = v1
    switch v1 {
    case "a":
        var v7 int16 = 301
        _ = v7
    case "":
        var v8 uint32 = 42
        _ = v8
    case "\n\t":
        v1 = (v1 + v1)
    default:
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v1 = v1
        if (v1 < "hello") {
            var v10 float32 = 0.0
            if false {
                var v11 float32 = v10
                fmt.Println(fi9, v1, fi9)
                v10 += -2.25
                _ = v11
            }
            var v12 uint8 = 255
            _ = v10
            _ = v12
        }
        if (2 != 943) {
            if (32767 <= 685) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    fmt.Println(fi9, fi13)
                    if len(v3) > 0 {
                        v3[0] = v1
                    }
                    _ = fi13
                }
            }
        }
        _ = fi9
    }
    switch "a" {
    case "foo bar":
        for fi14 := 0; fi14 < 6; fi14++ {
            v1 = ""
            if (false && (!false)) {
                var v15 int = (fi14 / (((^(fi14))) | 1))
                switch (fi14 & (fi14 - 127)) {
                case 2147483647:
                    var v16 int16 = -32768
                    _ = v16
                case 32768:
                    if (0.1 < -2.25) {
                        var v17 uint32 = 836
                        var v18 int64 = 1
                        switch (v15 << ((uint(v17)) & 63)) {
                        case -645:
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v20 := GenS1{F0: (fi14 | (v15 << ((0) & 63))), F1: 2}
                                var v21 string = ""
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        case 4294967295:
                            v18 = v18
                        case 128:
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v23 := GenS1{F0: fi22, F1: 3}
                                _ = fi22
                                _ = v23
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        for fi24 := 0; fi24 < 2; fi24++ {
                            if true {
                                v25 := GenS2{F0: 42}
                                var v26 float32 = 3.14159
                                var v27 uint64 = 2
                                v28 := GenS2{F0: 127}
                                _ = v25
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                v1 = (v1 + v1)
                            }
                            if true {
                                var v29 int16 = -166
                                _ = v29
                            }
                            _ = fi24
                        }
                        for fi30 := 0; fi30 < 7; fi30++ {
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            _ = fi30
                        }
                        var v31 int8 = 46
                        _ = v31
                    }
                default:
                    var v32 uint16 = 255
                    _ = v32
                }
                _ = v15
            } else {
                var v33 uint16 = 2
                _ = v33
            }
            fmt.Println(v1, v1)
            _ = fi14
        }
    case "":
        var v34 uint16 = 803
        _ = v34
    case "\n\t":
        v35 := GenS3{F0: ("foo bar" + ("a" + v1)), F1: true, F2: 146}
        _ = v35
    }
    fmt.Println(v1)
    delete(v2, 690)
    v1 = v1
    switch 127 {
    case 32767:
        v1 = (v1 + "\n\t")
    case -831:
        v1 = "\n\t"
    }
    fmt.Println(v1)
    _ = len(v3)
    v36 := GenS4{F0: 24, F1: 2147483648, F2: "hello"}
    var v37 int = 747
    var v38 string = v1
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 int8 = 42
        var v41 int32 = -1
        fmt.Println(v37, v40)
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v1, v38, v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
