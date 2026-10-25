package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint16; F2 string }

func main() {
    var v1 int64 = -510
    var v2 int64 = (^(v1))
    v3 := make(map[string]int)
    v4 := []string{("a" + (("hello" + "foo bar") + "\n\t")), ((((("hello" + "foo bar") + ("" + "hello")) + "") + "") + "")}
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 255, F1: 4294967295}
        _ = fi6
        _ = v7
    }
    if true {
        v1 = v1
        var v8 float32 = ((float32(v5) / 1.0) + (0.1 - float32(v5)))
        switch 2147483647 {
        case 128:
            for fi9 := 0; fi9 < 16; fi9++ {
                v5 -= 0.1
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 float32 = (0.1 + v8)
                    v11 += 1.5
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        case -9223372036854775808:
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        _ = v8
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint = 18446744073709551615
            for fi14 := 0; fi14 < 16; fi14++ {
                v13 = (v13 / ((v13) | 1))
                _ = fi14
            }
            v1 = (v2 + v1)
            _ = fi12
            _ = v13
        }
        if (!((0 <= 32767) && (42 < 1000))) {
            switch "hello" {
            case "":
                v15 := v3["a"]
                _ = v15
            default:
                var v16 uint64 = 2
                _ = v16
            }
        } else {
            if true {
                var v17 float64 = v5
                _ = v17
            }
            v2 = -28
            if (119 < 2) {
                var v18 uint = 2
                if true {
                    v19 := GenS2{F0: v18, F1: 18446744073709551615}
                    v5 = v5
                    var v20 float32 = -2.25
                    _ = v19
                    _ = v20
                } else {
                    switch ("a" + "") {
                    case "":
                        switch 127 {
                        case 330:
                            v21 := GenS1{F0: v18, F1: (v18 | (v18 - (v18 / ((42) | 1))))}
                            _ = v21
                        default:
                            var v22 uint8 = uint8(v18)
                            _ = v22
                        }
                    case "hello":
                        v23 := GenS3{F0: 0}
                        _ = v23
                    }
                }
                v5 += (float64(0.0) * ((v5 / 3.14159) / (v5 * v5)))
                _ = v18
            }
        }
        if true {
            fmt.Println(v1, v5)
            var v24 int32 = 828
            _ = v24
        } else {
            if true {
                switch ("hello" + ("" + "hello")) {
                case "":
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v2 += (^(v2))
                        switch "hello" {
                        case "a":
                            v2 = 1
                        case "hello":
                            fmt.Println(v1)
                        case "":
                            _ = len(v3)
                        default:
                            for fi26 := 0; fi26 < 1; fi26++ {
                                v2 = -331
                                v27 := GenS4{F0: 476}
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = fi25
                    }
                case "\n\t":
                    v2 = v2
                case "hello":
                    v5 = v5
                }
                v2 += 51
            } else {
                if true {
                    v1 = (v2 | 42)
                }
            }
        }
    }
    var v28 int8 = 1
    var v29 float64 = v5
    if (0.0 > -2.25) {
        v30 := v3["foo bar"]
        v2 = (v2 / ((-662) | 1))
        fmt.Println(v28, v2, v29)
        _ = v30
    } else {
        _ = len(v3)
        v1 += v1
        _ = len(v4)
    }
    v31 := GenS1{F0: 4294967295, F1: 2}
    _ = len(v3)
    var v32 uint64 = 4294967295
    var v33 string = ("foo bar" + "hello")
    for fi34 := 0; fi34 < 20; fi34++ {
        v35 := GenS4{F0: 42}
        v36 := GenS5{F0: (v2 % ((v1) | 1)), F1: 668, F2: (("" + (v33 + ("\n\t" + v33))) + v33)}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 bool = ((false && false) && false)
    if v37 {
        _ = len(v4)
        var v38 int16 = 1000
        _ = v38
    } else {
        fmt.Println(v32)
    }
    var v39 bool = true
    v40 := GenS4{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
