package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = -1
    v2 := []int{-348}
    var v3 int64 = -1
    var v4 uint = 2
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v5 float64 = (float64(3.14159) * 1e10)
    if len(v2) > 0 {
        v2[0] = 692
    }
    var v6 uint16 = 317
    if false {
        v6 -= (v6 ^ 774)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        fmt.Println(v5, v4)
        if true {
            _ = len(v2)
            for fi7 := 0; fi7 < 8; fi7++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v4, v3)
                case "foo bar":
                    v3 += (v3 - (v3 ^ 0))
                case "":
                    for fi8 := 0; fi8 < 13; fi8++ {
                        var v9 int8 = 0
                        _ = fi8
                        _ = v9
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi7
            }
            var v10 float64 = v5
            _ = v10
        }
        _ = len(v2)
    }
    switch 201 {
    case 0:
        fmt.Println(v4, v3)
    case -463:
        v3 = 492
    case 127:
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
    }
    var v13 float64 = (float64(1.5) - (v5 * v5))
    if ((v6 * (v6 << ((v4) & 63))) == 0) {
        if ((int64(v6) != v3) && true) {
            switch 9223372036854775807 {
            case 741:
                var v14 uint64 = 674
                _ = v14
            case -1:
                _ = len(v2)
            case 42:
                v4 = (v4 + v4)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v5 = 1.5
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint32 = 42
                    fmt.Println(v1, v13)
                    fmt.Println(v1, v5)
                    _ = fi16
                    _ = v17
                }
                var v18 float32 = 1.5
                if (!((100 != int8(v1)) || true)) {
                    switch (fi15 & fi15) {
                    case -82:
                        if true {
                            v1 += (v1 % (((v1 & -32768)) | 1))
                            v19 := GenS1{F0: 2}
                            for fi20 := 0; fi20 < 5; fi20++ {
                                v13 = 0.1
                                fmt.Println(v5, v18, fi20)
                                v6 -= v6
                                _ = fi20
                            }
                            _ = v19
                        } else {
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v21 uint8 = 48
                                _ = len(v2)
                                _ = v21
                            }
                            var v22 int8 = 1
                            _ = v22
                        }
                    }
                    if (float64(((v18 + v18) - -2.25)) < (1.5 - (v13 - (v13 * ((v13 + v5) / v5))))) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = fi15
                _ = v18
            }
        } else {
            fmt.Println(v3)
            _ = len(v2)
        }
        v23 := GenS1{F0: 42}
        v24 := GenS1{F0: 18446744073709551615}
        _ = v23
        _ = v24
    } else {
        fmt.Println(v3)
    }
    var v25 float32 = float32(1.5)
    var v26 uint = v4
    var v27 int8 = 100
    v27 = -47
    v25 -= (v25 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
