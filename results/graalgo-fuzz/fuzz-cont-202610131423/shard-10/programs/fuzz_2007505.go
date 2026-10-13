package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 32767
    var v3 uint8 = 99
    var v4 uint16 = (v2 / ((255) | 1))
    var v5 int64 = 0
    v4 -= (v2 - (v2 * v2))
    if true {
        if (true && false) {
            switch "" {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
        } else {
            v1 = append(v1, 700)
            for fi6 := 0; fi6 < 3; fi6++ {
                if (false && false) {
                    var v7 int = fi6
                    _ = v7
                }
                var v8 string = (("\n\t" + "foo bar") + "a")
                _ = fi6
                _ = v8
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v9 string = ("" + ((("hello" + "\n\t") + "foo bar") + "\n\t"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v2 -= 554
    var v10 bool = false
    var v11 int = 32768
    switch "hello" {
    case "\n\t":
        v11 -= (^(v11))
    case "hello":
        v3 += uint8(v2)
    }
    if true {
        var v12 int16 = 1
        v13 := GenS1{F0: int64(v3), F1: (v11 % ((32768) | 1))}
        _ = v12
        _ = v13
    }
    v2 = (v4 / ((uint16(v3)) | 1))
    var v14 uint8 = v3
    v10 = v10
    if true {
        var v15 uint16 = v2
        if (2 > 1) {
            var v16 int16 = 1000
            fmt.Println(v3)
            for fi17 := 0; fi17 < 15; fi17++ {
                if (v5 == v5) {
                    if len(v1) > 0 {
                        v1[0] = 110
                    }
                } else {
                    var v18 float32 = float32(3.14159)
                    v18 = (((v18 - v18) * v18) * (v18 / v18))
                    _ = v18
                }
                v4 = v2
                v19 := GenS1{F0: (v5 ^ -325), F1: 32768}
                _ = fi17
                _ = v19
            }
            _ = v16
        }
        _ = v15
    }
    switch 32768 {
    case 42:
        switch v9 {
        case "":
            if (0 > 18446744073709551615) {
                fmt.Println(v9, v11)
                if true {
                    if (73 == 42) {
                        var v20 uint16 = (v4 >> ((42) & 63))
                        _ = len(v1)
                        if v10 {
                            v1 = append(v1, 2147483647)
                        }
                        _ = v20
                    }
                    fmt.Println(v9, v14, v9)
                } else {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 string = v9
                        fmt.Println(v9, v3)
                        _ = fi21
                        _ = v22
                    }
                    switch v11 {
                    case -2147483648:
                        v23 := GenS1{F0: 1, F1: 1}
                        _ = v23
                    case -404:
                        var v24 uint = 65535
                        _ = v24
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 42
                }
            }
        case "hello":
            v4 = (^(v2))
        case "a":
            var v25 uint8 = v3
            _ = v25
        default:
            var v26 uint8 = (v14 % ((42) | 1))
            _ = v26
        }
    case 0:
        v27 := GenS1{F0: (v5 + (v5 - (v5 - int64(v14)))), F1: (v11 | 127)}
        _ = v27
    }
    var v28 uint8 = (^(v3))
    var v29 bool = (42 < uint64(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
