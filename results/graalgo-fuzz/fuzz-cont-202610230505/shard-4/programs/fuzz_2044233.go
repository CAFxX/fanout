package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 string; F1 uint8 }

func main() {
    v1 := []string{("foo bar" + ("\n\t" + ("foo bar" + "\n\t")))}
    var v2 float64 = float64(-2.25)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 int = 128
    v5 = v5
    switch v5 {
    case 129:
        if true {
            _ = len(v3)
            if true {
                if false {
                    fmt.Println(v2)
                    delete(v3, 127)
                }
                var v6 string = (("a" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello")
                fmt.Println(v6)
                _ = v6
            } else {
                v7 := GenS1{F0: 1, F1: 9223372036854775807}
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 127, F1: 607}
        var v10 uint64 = 965
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int64 = 42
            var v14 int16 = -32768
            fmt.Println(v2, v5)
            _ = fi12
            _ = v13
            _ = v14
        }
        if (260 > 111) {
            v15 := GenS2{F0: 2}
            _ = v15
        } else {
            v1 = append(v1, "\n\t")
        }
        _ = fi11
    }
    v2 -= v2
    v16 := GenS3{F0: 0}
    v2 -= float64(float32(0.0))
    v5 -= (v5 / ((v5) | 1))
    v17 := GenS4{F0: 2147483648, F1: -2147483648}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 bool = false
        switch "foo bar" {
        case "\n\t":
            v20 := v4[("" + ("a" + ("foo bar" + "a")))]
            _ = v20
        case "a":
            switch "a" {
            case "a":
                switch v5 {
                case -2147483648:
                    if (66 >= -597) {
                        var v21 float32 = 0.0
                        var v22 uint32 = 65535
                        _ = v21
                        _ = v22
                    }
                case 127:
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                case 1:
                    if ((false && (("\n\t" + "foo bar") < "\n\t")) || (3.14159 > 1.5)) {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        if (1 == -101) {
                            v19 = v19
                        } else {
                            var v23 uint32 = 943
                            v5 = fi18
                            _ = v23
                        }
                    }
                default:
                    var v24 string = "a"
                    _ = v24
                }
            default:
                v25 := GenS5{F0: "hello", F1: 138}
                _ = v25
            }
        default:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
