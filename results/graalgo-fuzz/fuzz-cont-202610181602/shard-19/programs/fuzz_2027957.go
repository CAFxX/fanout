package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16) bool {
    p0 += (p0 * (p0 - (p0 >> ((738) & 63))))
    p0 = 255
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int8 = -27
        var v8 uint = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return (!false)
}

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]bool)
    var v3 float64 = ((0.1 + 0.0) / (1e10 + 1.5))
    v4 := make(map[int]bool)
    var v5 uint = 18446744073709551615
    v9 := GenS1{F0: 42, F1: 9223372036854775807}
    var v10 uint = (v5 - (v5 + (^(v5))))
    v11 := GenS1{F0: 2147483647, F1: -2147483648}
    var v12 float64 = ((v3 - v3) - -2.25)
    switch -547 {
    case -618:
        var v13 int16 = 42
        _ = v13
    case -9223372036854775808:
        if false {
            fmt.Println(v3, v3)
            var v14 float64 = ((0.1 + 0.0) - (v12 - v3))
            _ = v14
        } else {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        delete(v2, ("" + ("a" + "foo bar")))
                        var v15 int8 = 1
                        v3 = v3
                        _ = v15
                    }
                default:
                    _ = len(v1)
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: (fi16 | (fi16 * (fi16 ^ -615))), F1: (fi16 + 9223372036854775807)}
                    var v18 int = 4294967295
                    fmt.Println(fi16, v18)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                switch 357 {
                case -1:
                    if true {
                        v5 -= 255
                        v19 := GenS2{F0: 627}
                        _ = v19
                    }
                case 32767:
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := GenS3{F0: (v10 / ((v10) | 1)), F1: 127}
                        v10 = v10
                        if true {
                            v22 := GenS3{F0: (v5 * (v5 / (((v10 >> ((v10) & 63))) | 1))), F1: 60}
                            switch (("\n\t" + "hello") + "\n\t") {
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                var v23 uint = 0
                                _ = v23
                            }
                            v2[""] = false
                            _ = v22
                        } else {
                            if true {
                                if (true || true) {
                                    var v24 float32 = 0.0
                                    v25 := GenS1{F0: fi20, F1: fi20}
                                    _ = v24
                                    _ = v25
                                } else {
                                    v5 = v10
                                    switch fi20 {
                                    case 32768:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 78:
                                        v12 = 1.5
                                    }
                                }
                            }
                            v5 = v5
                        }
                        _ = fi20
                        _ = v21
                    }
                case 1:
                    fmt.Println(v12)
                }
            }
            switch "" {
            case "a":
                var v26 string = ""
                _ = v26
            case "\n\t":
                var v27 string = ((("a" + "foo bar") + "foo bar") + ("hello" + ("foo bar" + "\n\t")))
                _ = v27
            default:
                v28 := GenS2{F0: 477}
                _ = v28
            }
            var v29 int8 = -128
            _ = v29
        }
    case -967:
        fmt.Println(v5, v12)
    default:
        var v30 int32 = 1
        _ = v30
    }
    if false {
        v31 := GenS3{F0: 1, F1: 42}
        v32 := GenS4{F0: 42}
        _ = v31
        _ = v32
    }
    var v33 uint32 = 1
    v5 += v10
    var v34 uint8 = 2
    if len(v1) > 0 {
        v1[0] = "a"
    }
    _ = len(v4)
    v35 := GenS4{F0: -873}
    v5 += v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
