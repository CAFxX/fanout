package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int8 = 127
    var v2 int32 = 42
    var v3 int = 32768
    var v4 uint16 = 42
    var v5 bool = true
    var v6 float64 = ((0.1 + 1e10) + (1e10 / 1e10))
    var v7 int32 = -289
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int8 = v1
        _ = fi9
        _ = v10
    }
    v3 = v3
    fmt.Println(v1)
    var v11 bool = v5
    v3 = -20
    if (-97 > -1) {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v1, v3, v1)
            var v13 uint16 = (v4 / ((65535) | 1))
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            switch 42 {
            case 127:
                var v15 bool = true
                _ = v15
            default:
                fmt.Println(v2, v11)
            }
            _ = fi14
        }
    } else {
        v16 := []int{v3, (v3 ^ v3), v3, (^(v3))}
        fmt.Println(v6, v6)
        v17 := GenS1{F0: v7}
        _ = v16
        _ = v17
    }
    var v18 float32 = (-2.25 * 3.14159)
    v19 := GenS1{F0: (v2 ^ 2147483647)}
    v1 = 127
    v20 := []int{-2147483648, (v3 ^ v3), (v3 % (((^(v3))) | 1))}
    if true {
        var v21 uint32 = 1
        for fi22 := 0; fi22 < 14; fi22++ {
            if v11 {
                if false {
                    if (true || (776 != 18446744073709551615)) {
                        var v23 uint = 255
                        _ = v23
                    }
                    if (-9223372036854775808 < 42) {
                        switch "" {
                        case "hello":
                            switch "hello" {
                            case "hello":
                                _ = len(v20)
                            default:
                                fmt.Println(v11)
                            }
                        case "\n\t":
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        case "":
                            fmt.Println(fi22, v18)
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v20 = append(v20, (v3 << ((1) & 63)))
                            _ = fi24
                        }
                    } else {
                        var v25 bool = (0 >= (v4 - v4))
                        switch ("foo bar" + ("\n\t" + "foo bar")) {
                        case "foo bar":
                            if (float64((-2.25 * (v18 * v18))) > float64(0.0)) {
                                v11 = v11
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    v2 -= (v2 - 2147483647)
                                    _ = fi26
                                }
                            } else {
                                var v27 int8 = (v1 - -44)
                                _ = v27
                            }
                        case "\n\t":
                            v20 = append(v20, 42)
                        case "":
                            v20 = append(v20, v3)
                        }
                        v20 = append(v20, (v3 - (fi22 * v3)))
                        _ = v25
                    }
                    v11 = v11
                } else {
                    fmt.Println(v7)
                    var v28 int8 = (^(v1))
                    for fi29 := 0; fi29 < 20; fi29++ {
                        var v30 uint = 2
                        if v5 {
                            v11 = true
                        }
                        var v31 bool = (!true)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
                fmt.Println(v21)
            }
            var v32 uint = 255
            _ = fi22
            _ = v32
        }
        v3 += 128
        _ = v21
    }
    if len(v20) > 0 {
        v20[0] = (v3 >> ((211) & 63))
    }
    fmt.Println(v5, v18, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
