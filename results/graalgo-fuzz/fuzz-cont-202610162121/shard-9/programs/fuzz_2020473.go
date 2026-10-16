package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }

func main() {
    v1 := []string{"a"}
    var v2 float32 = float32(1e10)
    var v3 uint16 = 32767
    var v4 float32 = (v2 + 0.0)
    if ((32767 == 598) && true) {
        var v5 int16 = 0
        fmt.Println(v3, v5, v5)
        fmt.Println(v4, v3, v2)
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint64 = 304
            for fi8 := 0; fi8 < 14; fi8++ {
                v7 += 0
                var v9 int32 = 1
                var v10 int = fi6
                _ = fi8
                _ = v10
                _ = v9
            }
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi6
            _ = v7
        }
    }
    v11 := GenS1{F0: 4294967295, F1: int32(v3), F2: 0}
    if (42 <= v3) {
        for fi12 := 0; fi12 < 11; fi12++ {
            if (0 <= 100) {
                v13 := GenS1{F0: 0, F1: 42, F2: 9223372036854775807}
                _ = v13
            }
            switch ("a" + ("" + "foo bar")) {
            case "foo bar":
                fmt.Println(v4, v2)
            case "a":
                switch "" {
                case "a":
                    if false {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v3 = (v3 * (^(v3)))
                            _ = fi14
                        }
                        if true {
                            if (true || true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    v3 += v3
                                    switch "hello" {
                                    case "a":
                                        v1 = append(v1, ("\n\t" + ("hello" + ("\n\t" + "a"))))
                                    }
                                } else {
                                    v4 = (v4 * v2)
                                }
                                v15 := GenS1{F0: 2147483648, F1: -227, F2: -1}
                                _ = v15
                            }
                            var v16 uint = 65535
                            _ = v16
                        }
                    } else {
                        var v17 float64 = 1.5
                        _ = v17
                    }
                case "\n\t":
                    _ = len(v1)
                default:
                    var v18 uint = 658
                    _ = v18
                }
            }
            _ = len(v1)
            _ = fi12
        }
        v1 = append(v1, "\n\t")
        switch -578 {
        case -818:
            v2 += (((((1.5 / (v4 * -2.25)) / (v2 * v4)) / float32(1.5)) + 1e10) * 1.5)
        default:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            v4 = v4
            _ = fi19
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: 45, F1: -2147483648, F2: 0}
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v2 = float32(float64(v4))
            _ = fi20
            _ = v21
        }
    }
    v4 += (1e10 + 1e10)
    var v22 uint8 = 127
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int32 = -1
        _ = fi23
        _ = v24
    }
    _ = len(v1)
    var v25 bool = (776 >= int32(v22))
    if (!((!(0.0 < 1.5)) && (-1 > 32768))) {
        var v26 uint8 = (v22 % ((v22) | 1))
        v27 := GenS1{F0: 4294967295, F1: -70, F2: 9223372036854775807}
        _ = v26
        _ = v27
    }
    var v28 uint8 = (^(v22))
    var v29 uint8 = (v28 / ((v28) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
