package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 128
    var v3 uint64 = 4294967295
    v4 := []int{}
    if true {
        v5 := GenS1{F0: ("foo bar" + ("\n\t" + (("hello" + ("hello" + "")) + "hello"))), F1: (v3 / ((v3) | 1)), F2: 25}
        _ = v5
    }
    v6 := GenS1{F0: "\n\t", F1: v3, F2: -1}
    fmt.Println(v3, v2)
    v7 := GenS1{F0: "", F1: (^(v3)), F2: 0}
    if ((v3 < v3) && true) {
        v3 = v3
    }
    if true {
        if ("hello" >= "hello") {
            var v8 int64 = 42
            v3 = 4294967295
            if ((-2.25 * 0.0) == (1e10 + 1.5)) {
                v8 = 9223372036854775807
                for fi9 := 0; fi9 < 18; fi9++ {
                    if false {
                        for fi10 := 0; fi10 < 1; fi10++ {
                            var v11 string = ("foo bar" + "a")
                            _ = fi10
                            _ = v11
                        }
                        var v12 uint16 = 32767
                        _ = v12
                    } else {
                        v2 += (v2 + v2)
                        v13 := GenS2{F0: 42, F1: (v8 + 0), F2: (v8 >> ((18446744073709551615) & 63))}
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v15 := GenS2{F0: 0, F1: v8, F2: v8}
                        v8 = v8
                        _ = fi14
                        _ = v15
                    }
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        for fi16 := 0; fi16 < 17; fi16++ {
                            var v17 float32 = 1.5
                            v8 -= v8
                            delete(v1, (((((("foo bar" + ("hello" + "a")) + "a") + "\n\t") + "") + "a") + ("foo bar" + (("hello" + "foo bar") + "a"))))
                            _ = fi16
                            _ = v17
                        }
                    case "hello":
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi18
                        }
                    default:
                        _ = len(v1)
                    }
                    _ = fi9
                }
                var v19 float32 = -2.25
                _ = v19
            }
            _ = v8
        }
        fmt.Println(v2, v3)
    }
    v3 -= v3
    var v20 int16 = 1000
    switch -1 {
    case 32768:
        var v21 uint8 = v2
        _ = v21
    }
    if (-659 != 852) {
        fmt.Println(v2, v3, v20)
        var v22 int16 = (v20 - -1)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 3; fi23++ {
            _ = len(v1)
            _ = fi23
        }
        v20 = -296
        for fi24 := 0; fi24 < 5; fi24++ {
            v2 = (^(v2))
            switch "\n\t" {
            case "foo bar":
                v25 := v1[(("hello" + "foo bar") + ("\n\t" + "a"))]
                _ = v25
            default:
                switch "foo bar" {
                case "foo bar":
                    v3 += (v3 & v3)
                default:
                    v4 = append(v4, fi24)
                }
            }
            _ = len(v4)
            _ = fi24
        }
    }
    v26 := GenS2{F0: 0, F1: -1, F2: -605}
    var v27 float64 = 1.5
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int64 = 42
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            v27 = (3.14159 - (3.14159 - (v27 * v27)))
            var v31 string = "foo bar"
            v2 -= (v2 - (v2 | v2))
            _ = v31
        }
        _ = fi30
    }
    v3 = v3
    for fi32 := 0; fi32 < 4; fi32++ {
        if len(v4) > 0 {
            v4[0] = fi32
        }
        fmt.Println(v3)
        _ = fi32
    }
    switch "hello" {
    case "foo bar":
        switch -413 {
        case -313:
            for fi33 := 0; fi33 < 15; fi33++ {
                v27 -= v27
                var v34 uint64 = (v3 ^ v3)
                _ = fi33
                _ = v34
            }
        case 42:
            v27 -= v27
        case 9223372036854775807:
            for fi35 := 0; fi35 < 17; fi35++ {
                _ = len(v1)
                _ = fi35
            }
        }
    default:
        var v36 uint = 0
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
