package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = 1
    var v2 float32 = float32((-2.25 * (0.0 - float64((0.1 / 1.0)))))
    v3 := []int{0, 0, int(v1)}
    v4 := make(map[int]int)
    if true {
        if true {
            switch "foo bar" {
            case "hello":
                v5 := GenS1{F0: 2147483647, F1: 97, F2: 0}
                _ = v5
            }
            if true {
                switch -535 {
                case 128:
                    v2 = float32((1e10 - 3.14159))
                case -678:
                    if ("\n\t" != "foo bar") {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v6 int16 = (v1 << ((1) & 63))
                        var v7 float32 = (1e10 + -2.25)
                        _ = v6
                        _ = v7
                    } else {
                        fmt.Println(v1, v2)
                        var v8 uint8 = 211
                        _ = v8
                    }
                default:
                    v1 -= (v1 >> ((18446744073709551615) & 63))
                }
                v9 := GenS2{F0: (163 <= 646)}
                _ = len(v4)
                _ = v9
            }
            v2 += v2
        } else {
            fmt.Println(v2, v1)
            v10 := v4[32768]
            fmt.Println(v10)
            _ = v10
        }
        fmt.Println(v2, v2)
        switch "foo bar" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
        default:
            var v11 uint32 = 2
            _ = v11
        }
    }
    var v12 uint = 119
    var v13 uint = v12
    var v14 int = 127
    fmt.Println(v13)
    v4[(^(v14))] = v14
    var v15 uint = v13
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v15, v12)
        _ = fi16
    }
    v12 += 65535
    for fi17 := 0; fi17 < 10; fi17++ {
        switch v14 {
        case 42:
            v13 = v12
        }
        _ = fi17
    }
    switch "foo bar" {
    case "hello":
        v14 = (v14 << ((v15) & 63))
    case "a":
        var v18 int32 = 42
        _ = v18
    case "foo bar":
        switch (v14 + v14) {
        case 244:
            if false {
                v19 := GenS1{F0: (v14 << ((260) & 63)), F1: 42, F2: 255}
                _ = v19
            } else {
                switch (v14 / ((v14) | 1)) {
                case 42:
                    if false {
                        if true {
                            v1 = 32767
                            v1 += v1
                        }
                        if false {
                            delete(v4, -9223372036854775808)
                        } else {
                            for fi20 := 0; fi20 < 2; fi20++ {
                                fmt.Println(fi20, v12, v2)
                                var v21 uint = v13
                                v12 = v12
                                _ = fi20
                                _ = v21
                            }
                            v2 = v2
                        }
                        v22 := v4[v14]
                        _ = v22
                    }
                case -1:
                    delete(v4, (v14 | 787))
                case -533:
                    v2 -= v2
                }
            }
        case 2147483647:
            if false {
                switch v14 {
                case 90:
                    v3 = append(v3, (v14 ^ -9223372036854775808))
                case 2147483647:
                    var v23 uint16 = uint16(v15)
                    _ = v23
                }
                v1 = (v1 << ((v15) & 63))
                v4[(v14 / (((v14 << ((2) & 63))) | 1))] = (v14 - v14)
            }
        default:
            v12 += v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
