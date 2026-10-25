package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint = 2
    var v2 int8 = 0
    v3 := make(map[int]string)
    var v4 int16 = -833
    v5 := make(map[string]bool)
    fmt.Println(v1)
    if (1 >= 269) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint8 = 127
            v2 += (v2 + v2)
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint = 0
                v7 += (v7 & 32)
                fmt.Println(v2, fi8)
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        if true {
            switch 32767 {
            case 127:
                var v10 int8 = 127
                _ = v10
            default:
                v11 := GenS1{F0: 986, F1: (v2 * 0)}
                _ = v11
            }
            v2 = v2
            v1 = 256
        }
        if (true || true) {
            var v12 int = -9223372036854775808
            _ = v12
        }
    }
    var v13 int = 127
    if (!false) {
        fmt.Println(v2, v2)
        switch "\n\t" {
        case "\n\t":
            var v14 string = "hello"
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 8; fi15++ {
                v13 -= (fi15 - -9223372036854775808)
                v13 = 4294967295
                _ = fi15
            }
        case "foo bar":
            if true {
                v2 = v2
                if false {
                    if (!(0.1 == (1e10 - (((1.5 / 1.0) * 1e10) - (0.1 * 0.0))))) {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            for fi17 := 0; fi17 < 8; fi17++ {
                                v18 := []int{}
                                _ = fi17
                                _ = v18
                            }
                            v13 = (^(fi16))
                            var v19 uint32 = 469
                            _ = fi16
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v4 -= -32768
                            _ = fi20
                        }
                        v21 := []int{}
                        _ = v21
                    }
                } else {
                    delete(v3, 211)
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if (!false) {
                            v23 := GenS2{F0: v2, F1: 9223372036854775807}
                            _ = v23
                        }
                        _ = fi22
                    }
                }
            }
        }
        v4 = (-(v4))
    }
    v1 -= v1
    var v24 string = "foo bar"
    switch "foo bar" {
    case "":
        var v25 uint64 = 4294967295
        _ = v25
    case "\n\t":
        v4 += v4
    default:
        v5[v24] = (!("hello" != ""))
    }
    v26 := GenS1{F0: 0, F1: (v2 | (v2 * 100))}
    v27 := v3[(-(v13))]
    var v28 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
