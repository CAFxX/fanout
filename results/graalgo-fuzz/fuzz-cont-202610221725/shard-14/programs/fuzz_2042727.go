package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 111
    v2 := []int{-2147483648, 127, -2147483648, 32767}
    var v3 float32 = (1e10 - 1e10)
    var v4 float32 = v3
    var v5 uint32 = 559
    switch 0 {
    case -289:
        v5 += v5
    case 2147483647:
        switch 32768 {
        case 32768:
            if false {
                if false {
                    fmt.Println(v1, v4)
                    v4 -= (1e10 + v3)
                }
                v1 = 42
                switch "foo bar" {
                case "hello":
                    v5 = v5
                case "\n\t":
                    var v6 float64 = (((0.1 * 3.14159) + 0.1) + (0.1 * ((0.1 * 0.0) - ((0.0 * float64((1e10 - (v4 + -2.25)))) + 3.14159))))
                    _ = v6
                }
            } else {
                v2 = append(v2, 836)
                v2 = append(v2, -1)
                fmt.Println(v4)
            }
        case 0:
            if (231 >= 1) {
                var v7 int32 = 1
                _ = v7
            }
        case -2147483648:
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: "\n\t"}
                _ = fi8
                _ = v9
            }
        }
    case -1:
        v3 -= float32(0.0)
    }
    var v10 int = 42
    if (255 != v1) {
        v3 = v3
        fmt.Println(v3, v10)
    } else {
        v1 -= (v1 << ((v1) & 63))
        var v11 float64 = (3.14159 + float64(((v4 - -2.25) + 3.14159)))
        var v12 uint = v1
        _ = v11
        _ = v12
    }
    var v13 uint64 = 998
    for fi14 := 0; fi14 < 2; fi14++ {
        v1 = (v1 % (((v1 ^ v1)) | 1))
        _ = fi14
    }
    _ = len(v2)
    for fi15 := 0; fi15 < 10; fi15++ {
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi15
    }
    var v16 float32 = 1.5
    v4 += v3
    v16 = float32(1e10)
    if false {
        v13 = 971
        if true {
            if (!true) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v16 = 0.0
                    _ = fi17
                }
            }
            fmt.Println(v10)
            v18 := GenS1{F0: "foo bar"}
            _ = v18
        } else {
            var v19 int = (v10 | v10)
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v3, v10, v3)
                if (2147483647 == (v5 % ((v5) | 1))) {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        switch "" {
                        case "\n\t":
                            var v22 int32 = -1
                            _ = v22
                        case "":
                            v16 = (1e10 + (v16 * v3))
                        }
                        var v23 int = (v19 % (((fi20 & fi21)) | 1))
                        var v24 float32 = ((3.14159 / v16) - float32(((3.14159 - 0.0) / 1.0)))
                        _ = fi21
                        _ = v23
                        _ = v24
                    }
                    switch "foo bar" {
                    case "":
                        v25 := GenS1{F0: "foo bar"}
                        _ = v25
                    case "\n\t":
                        var v26 uint16 = 65535
                        _ = v26
                    default:
                        _ = len(v2)
                    }
                } else {
                    switch v19 {
                    case 2147483647:
                        v27 := GenS2{F0: v5}
                        _ = v27
                    }
                }
                _ = fi20
            }
            v16 = v3
            _ = v19
        }
    } else {
        v2 = append(v2, (-(v10)))
    }
    var v28 int32 = 0
    if true {
        if (446 >= 42) {
            v29 := GenS1{F0: "foo bar"}
            v28 = (v28 ^ v28)
            var v30 int16 = 956
            _ = v29
            _ = v30
        }
    } else {
        fmt.Println(v5, v28, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
