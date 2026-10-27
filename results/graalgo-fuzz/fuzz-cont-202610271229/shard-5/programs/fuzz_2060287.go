package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0() int8 {
    var v5 int8 = 42
    var v6 int = 9223372036854775807
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 float32 = (3.14159 / (0.1 - (3.14159 - 3.14159)))
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 float32 = v1
    if (581 == 2147483648) {
        fmt.Println(v4)
        if true {
            var v7 uint64 = 946
            var v8 uint32 = 0
            for fi9 := 0; fi9 < 3; fi9++ {
                v8 = v8
                _ = fi9
            }
            _ = v7
            _ = v8
        }
    } else {
        var v10 int8 = genFn0()
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int16 = 42
        _ = fi11
        _ = v12
    }
    if (2147483647 == 0) {
        v13 := GenS1{F0: 0, F1: 42}
        v1 = v4
        _ = v13
    }
    fmt.Println(v1, v4)
    v4 = -2.25
    var v14 string = "a"
    v15 := GenS1{F0: 100, F1: 32767}
    switch 32768 {
    case 127:
        if false {
            fmt.Println(v14, v1)
            v14 = v14
        } else {
            var v16 int16 = 32767
            var v17 float32 = 1.5
            _ = v16
            _ = v17
        }
    case 42:
        if true {
            if (0 != 65535) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: 1, F1: 65535}
                    v4 += v1
                    switch (fi18 | (fi18 / ((fi18) | 1))) {
                    case -303:
                        fmt.Println(v4, v1, v14)
                    case 842:
                        _ = len(v2)
                    default:
                        v1 = float32((-2.25 * ((float64((1e10 / 1.0)) * float64(0.0)) / -2.25)))
                    }
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        switch v14 {
                        case "hello":
                            var v22 uint = 377
                            _ = v22
                        case "foo bar":
                            switch ((v14 + (v14 + (("a" + v14) + "a"))) + v14) {
                            case "":
                                if false {
                                    v14 = v14
                                }
                            case "a":
                                var v23 int8 = genFn0()
                                _ = v23
                            default:
                                _ = len(v3)
                            }
                        case "":
                            fmt.Println(v4, v14)
                        default:
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        }
                        _ = fi21
                    }
                    fmt.Println(v14)
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v4 = float32(0.1)
                        _ = fi25
                    }
                    _ = fi20
                }
            }
        } else {
            v3[v14] = -363
        }
    case 0:
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 int32 = 149
                v29 := GenS1{F0: 0, F1: 2}
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    default:
        v4 += (1.5 + v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
