package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{32767, 1, -1, 851}
    var v3 int64 = 9223372036854775807
    var v4 float64 = 1.5
    _ = len(v2)
    v3 = v3
    var v5 int16 = 703
    var v6 int16 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 132
        delete(v1, "foo bar")
        for fi9 := 0; fi9 < 7; fi9++ {
            switch ("hello" + "") {
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                if ((v5 + (v6 ^ -32768)) < v6) {
                    v4 = (1.5 + (0.1 - ((3.14159 / v4) * v4)))
                    var v10 float64 = 1e10
                    v6 -= -32768
                    _ = v10
                } else {
                    fmt.Println(v8)
                    if len(v2) > 0 {
                        v2[0] = (fi9 % ((fi7) | 1))
                    }
                    if len(v2) > 0 {
                        v2[0] = int(v6)
                    }
                }
            case "hello":
                switch (("\n\t" + ("hello" + "hello")) + "hello") {
                case "foo bar":
                    switch "" {
                    case "":
                        v11 := GenS1{F0: fi7, F1: 74}
                        _ = v11
                    }
                }
            default:
                _ = len(v2)
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if (0 != 65535) {
        if false {
            v3 += v3
            if (v4 == -2.25) {
                for fi12 := 0; fi12 < 13; fi12++ {
                    fmt.Println(v5, v3)
                    v3 = (v3 + (v3 / ((-9223372036854775808) | 1)))
                    v13 := GenS1{F0: (fi12 % ((fi12) | 1)), F1: -9}
                    _ = fi12
                    _ = v13
                }
                v14 := GenS2{F0: 9223372036854775807, F1: 0}
                _ = v14
            } else {
                var v15 float64 = ((1e10 - v4) - (v4 + -2.25))
                for fi16 := 0; fi16 < 14; fi16++ {
                    if false {
                        v15 = (((0.0 - (v15 * float64((3.14159 / 3.14159)))) + (-2.25 + 1.5)) - float64(0.0))
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if true {
                                v18 := GenS2{F0: v3, F1: 247}
                                v19 := GenS3{F0: 42}
                                _ = v18
                                _ = v19
                            } else {
                                var v20 uint = 2
                                v6 = 42
                                _ = v20
                            }
                            var v21 float64 = (v4 - (v4 + v15))
                            _ = fi17
                            _ = v21
                        }
                    }
                    var v22 uint = 1
                    _ = fi16
                    _ = v22
                }
                _ = v15
            }
            switch "hello" {
            case "":
                var v23 int64 = 9223372036854775807
                _ = v23
            case "\n\t":
                fmt.Println(v5, v3, v6)
            default:
                var v24 int16 = 42
                _ = v24
            }
        }
        var v25 int16 = -1
        if len(v2) > 0 {
            v2[0] = 0
        }
        _ = v25
    }
    fmt.Println(v6)
    switch -49 {
    case -998:
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := GenS1{F0: (fi26 >> ((289) & 63)), F1: 17}
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
