package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool, p2 int64) int16 {
    p0 += 2
    if (((((1.5 + 1.5) + (3.14159 + 1.5)) - 1.5) / 1e10) != 1e10) {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 float64 = 3.14159
            switch "" {
            case "":
                for fi7 := 0; fi7 < 14; fi7++ {
                    p0 = p0
                    v8 := []int{(fi7 % ((-9223372036854775808) | 1))}
                    _ = fi7
                    _ = v8
                }
            default:
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int32 = 159
                    var v11 uint8 = 69
                    fmt.Println(v6)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            var v12 float64 = 0.0
            _ = fi5
            _ = v12
            _ = v6
        }
        if (1000 < -825) {
            var v13 int64 = (p2 % ((1) | 1))
            _ = v13
        }
    }
    var v14 uint64 = 0
    var v15 uint = 376
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := []string{("\n\t" + "")}
    v2 := []int{-2147483648, 32768}
    var v3 uint16 = 2
    var v4 int64 = 1
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int16 = 42
        fmt.Println(fi16)
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 1}
    v4 += 42
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 float32 = ((3.14159 * (((1.5 - (0.0 / (0.1 + ((1.5 + 0.1) - (3.14159 - (1.5 + -2.25)))))) / 1e10) + ((0.1 / 1.5) / (1e10 * 1e10)))) / 1e10)
            _ = len(v1)
            var v22 uint8 = 2
            _ = fi20
            _ = v21
            _ = v22
        }
        if true {
            v3 -= 65535
            for fi23 := 0; fi23 < 11; fi23++ {
                switch ("a" + ("hello" + ("\n\t" + "foo bar"))) {
                case "\n\t":
                    switch fi19 {
                    case 2147483647:
                        v24 := GenS1{F0: 0}
                        _ = v24
                    case 0:
                        v4 -= (^(v4))
                    case -941:
                        var v25 float64 = 3.14159
                        _ = v25
                    default:
                        v3 = 531
                    }
                }
                switch "a" {
                case "hello":
                    var v26 float32 = 0.1
                    _ = v26
                case "\n\t":
                    fmt.Println(fi19, fi23, fi19)
                case "":
                    switch "hello" {
                    case "hello":
                        if len(v2) > 0 {
                            v2[0] = 32768
                        }
                        v3 += 0
                        var v27 int32 = 2147483647
                        _ = v27
                    case "a":
                        v4 = 171
                    case "\n\t":
                        v3 -= (^(v3))
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                }
                _ = fi23
            }
        } else {
            v28 := GenS2{F0: 144}
            _ = v28
        }
        _ = fi19
    }
    fmt.Println(v3, v4)
    v4 = v4
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS2{F0: -752}
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi29
        _ = v30
    }
    switch 32768 {
    case -1:
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS2{F0: genFn0((v3 * (v3 + 217)), ((2147483647 > 0) && true), v4)}
            v33 := GenS1{F0: 319}
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    switch ("\n\t" + (("\n\t" + "") + "hello")) {
    case "\n\t":
        var v34 int = -1
        _ = v34
    case "":
        _ = len(v1)
    default:
        v4 -= (^(v4))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
