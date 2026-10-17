package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int64 = 0
    var v2 float32 = (0.0 * (0.0 - (0.0 + float32(0.1))))
    var v3 float64 = 1.5
    var v4 uint16 = 2
    var v5 int32 = 2147483647
    switch "foo bar" {
    case "hello":
        v4 = v4
    default:
        if true {
            if true {
                v6 := GenS1{F0: (!true)}
                fmt.Println(v4, v1)
                _ = v6
            }
            fmt.Println(v2, v1, v1)
            v7 := []int{}
            _ = v7
        } else {
            switch -988 {
            case 4294967295:
                var v8 bool = false
                _ = v8
            }
            v4 = v4
        }
    }
    v9 := []int{-9223372036854775808, -72, 2147483647}
    v10 := GenS1{F0: (true && (32767 != 42))}
    switch "foo bar" {
    case "":
        var v11 bool = false
        _ = v11
    case "\n\t":
        if false {
            v2 -= v2
            v1 = -9223372036854775808
        }
    }
    fmt.Println(v4, v3, v2)
    var v12 uint16 = (v4 | v4)
    var v13 float32 = 0.1
    if (float32(v3) > float32(float64(float32((((v3 / v3) / (v3 + 0.1)) - (v3 / v3)))))) {
        v2 = v2
        v13 -= v13
    } else {
        switch "" {
        case "foo bar":
            var v14 string = ("foo bar" + "\n\t")
            _ = v14
        default:
            var v15 uint = 405
            _ = v15
        }
        if true {
            if (-1 > -53) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v1 += -909
                    _ = fi16
                }
            } else {
                v17 := GenS1{F0: (true || (false || false))}
                _ = v17
            }
            if ((2147483647 > 249) && ((!(2147483648 < 2147483647)) && false)) {
                switch -159 {
                case 2147483647:
                    v13 -= v13
                }
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        fmt.Println(v2, fi18)
                        if ((false && true) || true) {
                            if ((v2 * (v13 + ((v2 + 1.5) + v13))) > 1.5) {
                                fmt.Println(v5, fi18)
                                var v20 int16 = 777
                                if false {
                                    var v21 uint16 = 65535
                                    v3 += (0.1 + -2.25)
                                    _ = v21
                                } else {
                                    if true {
                                        v9 = append(v9, (fi19 % ((fi18) | 1)))
                                        var v22 bool = true
                                        v23 := GenS1{F0: (!true)}
                                        _ = len(v9)
                                        v24 := GenS1{F0: v22}
                                        var v25 int = (fi18 & (^(fi18)))
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    v1 += 9223372036854775807
                                }
                                _ = v20
                            }
                            v1 -= (v1 % ((v1) | 1))
                            var v26 int8 = -57
                            _ = v26
                        } else {
                            fmt.Println(v1, v1, v2)
                            var v27 float32 = v2
                            _ = v27
                        }
                        if false {
                            fmt.Println(v5, v13, v3)
                            fmt.Println(v3, v12)
                        }
                        _ = fi19
                    }
                    _ = fi18
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 int = fi28
                    if len(v9) > 0 {
                        v9[0] = (fi28 | 2147483647)
                    }
                    _ = fi28
                    _ = v29
                }
            }
            v30 := GenS1{F0: (9223372036854775807 <= 0)}
            _ = v30
        }
    }
    if (!true) {
        fmt.Println(v3, v5)
        v31 := GenS1{F0: true}
        _ = v31
    }
    var v32 float32 = float32((0.1 - v3))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
