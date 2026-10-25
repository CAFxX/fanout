package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 float32 = 3.14159
    v4 := []string{(("\n\t" + "\n\t") + ("" + "a"))}
    v3 = (v3 * ((0.1 / v3) * 0.0))
    var v5 uint16 = 0
    v1 -= (v1 << ((42) & 63))
    switch "foo bar" {
    case "foo bar":
        var v6 uint16 = v5
        _ = v6
    case "a":
        _ = len(v4)
    default:
        if false {
            var v7 int32 = -1
            _ = v7
        } else {
            delete(v2, "hello")
        }
    }
    fmt.Println(v5, v3)
    v1 += v1
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            if (!true) {
                if false {
                    v10 := GenS1{F0: 168, F1: 42}
                    if false {
                        var v11 uint16 = v5
                        switch ("" + (("a" + "") + ("hello" + (("hello" + "") + "hello")))) {
                        case "hello":
                            var v12 uint64 = 42
                            _ = v12
                        default:
                            var v13 int32 = 2147483647
                            _ = v13
                        }
                        if (fi9 == 4294967295) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v14 int8 = -29
                            _ = v14
                        }
                        _ = v11
                    } else {
                        v3 += v3
                        v5 = 2
                    }
                    _ = v10
                }
                v15 := v2["\n\t"]
                switch "" {
                case "":
                    var v16 int64 = v1
                    _ = v16
                case "\n\t":
                    switch fi9 {
                    case 0:
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v5 += v5
                        case "":
                            v1 -= v1
                        default:
                            v1 = v1
                        }
                    }
                case "a":
                    _ = len(v4)
                }
                var v17 uint16 = (v5 * (v5 & (v5 % ((v5) | 1))))
                v18 := GenS2{F0: 189, F1: (127 >= 0)}
                _ = v15
                _ = v17
                _ = v18
            }
            var v19 uint = 731
            var v20 int = (fi9 % ((fi8) | 1))
            if true {
                v19 += (^(v19))
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v21 string = ("" + "\n\t")
                var v22 float64 = -2.25
                _ = v21
                _ = v22
            }
            fmt.Println(v3, v3, v3)
            _ = fi9
            _ = v19
            _ = v20
        }
        delete(v2, (("hello" + "a") + (("" + ("foo bar" + ("a" + "foo bar"))) + "hello")))
        v5 = (v5 | (v5 >> ((1) & 63)))
        _ = fi8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = -1
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 uint64 = 4294967295
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 uint32 = 338
            var v27 float64 = (((float64(v3) * (float64(1.5) * 1e10)) * 1.5) * (0.1 - -2.25))
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    var v28 uint16 = v5
    var v29 uint16 = 255
    fmt.Println(v3)
    if true {
        var v30 int = -328
        _ = v30
    }
    if false {
        delete(v2, "foo bar")
    } else {
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
