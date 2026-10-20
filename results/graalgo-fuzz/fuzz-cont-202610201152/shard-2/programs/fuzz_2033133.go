package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0() int8 {
    if false {
        var v7 string = "foo bar"
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int32 = 1
            v10 := GenS1{F0: 127}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            if (((0.0 * (((float64(0.1) + float64(0.1)) + -2.25) + 0.0)) - (0.0 + (1.5 * 0.0))) < (((0.1 + (3.14159 + 1.5)) - 1e10) * (-2.25 * 1.5))) {
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 int8 = -63
                    if (!false) {
                        v14 := GenS1{F0: 128}
                        switch ("foo bar" + (v7 + v7)) {
                        case "hello":
                            fmt.Println(v13, fi11)
                        case "a":
                            if true {
                                v7 = "hello"
                                v13 = v13
                            }
                        case "\n\t":
                            v7 = v7
                        default:
                            var v15 uint = 1
                            _ = v15
                        }
                        _ = v14
                    }
                    _ = fi12
                    _ = v13
                }
            }
            fmt.Println(v7, v7)
            v7 = v7
            _ = fi11
        }
        _ = v7
    }
    switch 2147483647 {
    case 997:
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint = 210
            fmt.Println(fi16, v17)
            _ = fi16
            _ = v17
        }
    case 42:
        var v18 uint16 = 65535
        _ = v18
    case 32767:
        switch 1 {
        case -2147483648:
            var v19 float32 = (-2.25 - 1.5)
            _ = v19
        }
    default:
        switch 623 {
        case -9223372036854775808:
            v20 := []int{-470, 225, 404, 4294967295}
            _ = v20
        case 32768:
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 int = fi21
                v22 -= (^(v22))
                _ = fi21
                _ = v22
            }
        case -408:
            v23 := GenS1{F0: 0}
            _ = v23
        }
    }
    var v24 bool = (-249 == 32768)
    _ = v24
    return 56
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -100
    var v3 uint16 = 1
    var v4 uint8 = 2
    v5 := make(map[uint32]string)
    v6 := []string{("" + "")}
    if false {
        v6 = append(v6, "a")
        fmt.Println(v2, v2)
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v4 += (^(v4))
        _ = fi25
    }
    if true {
        fmt.Println(v3, v4, v4)
        v5[2] = "hello"
    }
    if len(v6) > 0 {
        v6[0] = ("hello" + ("\n\t" + ("foo bar" + "foo bar")))
    }
    fmt.Println(v2)
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            fmt.Println(v4, fi27)
            v28 := GenS1{F0: v4}
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    switch "\n\t" {
    case "\n\t":
        v29 := GenS2{F0: 4294967295, F1: (v2 / ((v2) | 1))}
        _ = v29
    case "":
        v30 := GenS2{F0: 588, F1: v2}
        _ = v30
    case "foo bar":
        v31 := GenS2{F0: 42, F1: (v2 % ((v2) | 1))}
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        if false {
            for fi33 := 0; fi33 < 14; fi33++ {
                if true {
                    var v34 uint32 = 423
                    _ = v34
                }
                var v35 int64 = -299
                _ = fi33
                _ = v35
            }
            v4 = 37
            var v36 uint64 = 4294967295
            _ = v36
        }
        var v37 uint = 4294967295
        _ = fi32
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
