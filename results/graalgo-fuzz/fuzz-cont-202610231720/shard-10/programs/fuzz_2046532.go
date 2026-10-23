package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 42
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    fmt.Println(v2)
    v2 += (^(v2))
    var v5 float64 = ((0.1 - (1.5 / 1.5)) * (-2.25 * ((-2.25 / 3.14159) / float64(1.5))))
    if false {
        v5 -= v5
        var v6 uint8 = 1
        _ = v6
    } else {
        v5 += 0.1
        if false {
            fmt.Println(v5, v2)
            switch ("" + "a") {
            case "":
                if true {
                    fmt.Println(v5, v2)
                    for fi7 := 0; fi7 < 17; fi7++ {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            switch (fi8 >> ((42) & 63)) {
                            case 128:
                                v2 -= (^(v2))
                            case -762:
                                var v9 uint16 = 42
                                _ = v9
                            case -980:
                                var v10 int8 = -72
                                _ = v10
                            }
                            v11 := v4["hello"]
                            v2 -= 0
                            _ = fi8
                            _ = v11
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v12 uint64 = 42
                            _ = v12
                        }
                        _ = fi7
                    }
                    for fi13 := 0; fi13 < 2; fi13++ {
                        if false {
                            if false {
                                fmt.Println(v2, v5, v5)
                                delete(v3, "a")
                                v14 := []int{(fi13 >> ((726) & 63)), fi13, fi13}
                                _ = v14
                            }
                        }
                        _ = fi13
                    }
                } else {
                    fmt.Println(v5, v5)
                    switch -9223372036854775808 {
                    case 32767:
                        if (926 > 906) {
                            for fi15 := 0; fi15 < 7; fi15++ {
                                var v16 uint = 65535
                                _ = fi15
                                _ = v16
                            }
                        }
                    case -764:
                        if (2147483647 < 1) {
                            switch 42 {
                            case 1:
                                v2 -= v2
                            case 9223372036854775807:
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 int8 = 42
                                    _ = fi17
                                    _ = v18
                                }
                            default:
                                delete(v4, "hello")
                            }
                            if true {
                                v19 := GenS1{F0: 544, F1: 963}
                                switch 2147483647 {
                                case 32767:
                                    var v20 uint = 255
                                    _ = v20
                                case 4294967295:
                                    fmt.Println(v5, v5)
                                default:
                                    var v21 int64 = 9223372036854775807
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                if false {
                                    var v22 float64 = v5
                                    var v23 int8 = -1
                                    var v24 int64 = (v2 - (v2 ^ v2))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                v5 = (1e10 - (0.0 + v5))
                            }
                        } else {
                            var v25 float32 = 0.1
                            _ = v25
                        }
                    default:
                        fmt.Println(v5, v5)
                    }
                }
            default:
                v2 = (v2 + (v2 ^ 601))
            }
            v26 := []int{188}
            _ = v26
        } else {
            v27 := GenS2{F0: 0}
            _ = v27
        }
    }
    delete(v4, "foo bar")
    var v28 string = ("foo bar" + "foo bar")
    v5 += v5
    for fi29 := 0; fi29 < 8; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            fmt.Println(v5, fi29, fi29)
            _ = fi30
        }
        var v31 float64 = (((1.5 / 0.1) / (v5 / v5)) + v5)
        var v32 int64 = (v2 * (-(v2)))
        _ = fi29
        _ = v31
        _ = v32
    }
    delete(v4, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
