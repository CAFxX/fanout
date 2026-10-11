package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int }

func main() {
    var v1 int16 = -386
    var v2 float32 = float32(((0.0 - (1.5 / 0.1)) * (float64(-2.25) * ((1e10 - 3.14159) - (0.1 * (1.5 + 1.5))))))
    var v3 int8 = 43
    v4 := []uint32{4294967295, 1, 63}
    v5 := make(map[string]string)
    _ = len(v4)
    if (!false) {
        v5["hello"] = ""
    }
    v3 += (v3 / (((v3 % (((v3 * (v3 >> ((255) & 63)))) | 1))) | 1))
    v6 := GenS1{F0: "hello", F1: v3, F2: (v3 << ((0) & 63))}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (2 <= 1) {
            v4 = append(v4, 2147483647)
        }
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if true {
                var v11 bool = (((3.14159 / 1.0) / (0.1 + 3.14159)) < 1.5)
                _ = len(v4)
                if true {
                    if (2 == 1) {
                        var v12 int32 = 600
                        var v13 uint32 = 1
                        _ = v12
                        _ = v13
                    }
                } else {
                    var v14 int8 = -67
                    var v15 uint8 = 2
                    v14 = (v14 - 100)
                    _ = v14
                    _ = v15
                }
                v5["\n\t"] = "a"
                v16 := GenS1{F0: "\n\t", F1: (-(v3)), F2: -1}
                if v11 {
                    v3 = (-(v3))
                    switch ("a" + "hello") {
                    case "":
                        fmt.Println(v2, v11)
                    case "\n\t":
                        v17 := GenS1{F0: "a", F1: v3, F2: (-(v3))}
                        _ = v17
                    case "foo bar":
                        var v18 int = fi10
                        _ = v18
                    }
                    v3 = (v3 * 9)
                    var v19 float64 = (-2.25 / float64(v2))
                    var v20 float32 = (v2 + (((0.1 + v2) - float32(v19)) + (v2 - (v2 / 1e10))))
                    _ = v19
                    _ = v20
                } else {
                    var v21 string = "hello"
                    v5[(("" + v21) + "")] = ("\n\t" + "foo bar")
                    _ = v21
                }
                fmt.Println(v2, v1, v3)
                v22 := GenS2{F0: 0, F1: (v1 << ((255) & 63))}
                _ = v11
                _ = v16
                _ = v22
            }
            _ = len(v5)
            _ = fi10
        }
        _ = fi9
    }
    var v23 uint8 = 49
    if false {
        var v24 int8 = (-(v3))
        var v25 int16 = v1
        if false {
            switch 9223372036854775807 {
            case 32767:
                switch ((("" + ((("a" + ("hello" + (("hello" + "") + ""))) + "hello") + ("" + "foo bar"))) + ("\n\t" + "a")) + "\n\t") {
                case "\n\t":
                    var v26 uint16 = 245
                    _ = v26
                case "":
                    v27 := GenS2{F0: 1, F1: v1}
                    _ = v27
                case "a":
                    fmt.Println(v2, v1)
                default:
                    for fi28 := 0; fi28 < 6; fi28++ {
                        v23 += (v23 & (^(v23)))
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v30 := GenS3{F0: 42, F1: 1, F2: 127}
                            v2 -= (v2 - (v2 * v2))
                            _ = fi29
                            _ = v30
                        }
                        v3 = v3
                        _ = fi28
                    }
                }
            }
            var v31 float64 = float64((v2 * v2))
            fmt.Println(v2, v2)
            _ = v31
        }
        _ = v24
        _ = v25
    } else {
        var v32 int8 = (v3 / (((v3 / ((1) | 1))) | 1))
        _ = v32
    }
    var v33 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
