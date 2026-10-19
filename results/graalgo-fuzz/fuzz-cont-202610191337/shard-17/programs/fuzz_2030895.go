package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 42
    var v3 int64 = 748
    var v4 int = 2147483647
    var v5 uint16 = 0
    var v6 int8 = 100
    var v7 float32 = 3.14159
    v8 := []int{v4, v4, 127, (v4 % (((v4 * 784)) | 1))}
    if (!(864 != 4294967295)) {
        v9 := GenS1{F0: (^(v2))}
        var v10 float32 = v7
        v4 -= v4
        _ = v10
        _ = v9
    }
    switch v4 {
    case 612:
        v2 -= 78
    case -1:
        var v11 int16 = 1
        _ = v11
    case -180:
        if false {
            if ((float64(3.14159) + float64((float32(0.1) * v7))) > 1e10) {
                v12 := GenS1{F0: v2}
                if false {
                    if true {
                        var v13 uint16 = v5
                        if true {
                            v13 += v5
                            if true {
                                v14 := GenS1{F0: v2}
                                var v15 uint64 = 670
                                _ = v14
                                _ = v15
                            }
                        }
                        var v16 uint32 = 2
                        _ = v13
                        _ = v16
                    }
                    if false {
                        var v17 int16 = 873
                        var v18 bool = (-1 > v6)
                        _ = v17
                        _ = v18
                    } else {
                        v2 -= (v2 * (v2 >> ((1) & 63)))
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        if true {
                            fmt.Println(v5, v7)
                            v3 -= 9223372036854775807
                        }
                    }
                }
                _ = v12
            } else {
                v19 := GenS1{F0: v2}
                _ = v19
            }
        }
    default:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v20 uint16 = 118
    for fi21 := 0; fi21 < 20; fi21++ {
        switch int(v3) {
        case 9223372036854775807:
            var v22 float64 = -2.25
            _ = v22
        default:
            if false {
                var v23 bool = true
                if (42 < -111) {
                    var v24 uint = 1
                    _ = v24
                } else {
                    var v25 uint16 = 891
                    fmt.Println(v25, v25, v1)
                    _ = v25
                }
                _ = v23
            } else {
                var v26 bool = (!true)
                v8 = append(v8, (v4 | v4))
                v3 -= 9223372036854775807
                _ = v26
            }
        }
        _ = fi21
    }
    v27 := GenS2{F0: (v2 | 244)}
    switch v4 {
    case 1:
        v3 += v3
    case 9223372036854775807:
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 int = -704
            switch ("a" + "\n\t") {
            case "\n\t":
                if false {
                    switch (v4 * (v29 % (((fi28 - (-(v4)))) | 1))) {
                    case 324:
                        v8 = append(v8, (fi28 | fi28))
                    }
                    v30 := GenS3{F0: fi28, F1: (!((v3 ^ v3) <= (v3 ^ v3)))}
                    _ = v30
                }
            default:
                if ((18446744073709551615 <= 255) && false) {
                    _ = len(v8)
                    if len(v8) > 0 {
                        v8[0] = -875
                    }
                }
            }
            _ = fi28
            _ = v29
        }
    case 127:
        switch "foo bar" {
        case "\n\t":
            v31 := GenS1{F0: (v2 << ((982) & 63))}
            _ = v31
        case "a":
            var v32 int8 = (v6 | 42)
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 7; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                v8 = append(v8, fi34)
                fmt.Println(v2, v2)
                switch "\n\t" {
                case "hello":
                    v20 = (v20 | v5)
                case "":
                    _ = len(v8)
                case "\n\t":
                    v1 += v1
                default:
                    switch 32768 {
                    case 9223372036854775807:
                        v2 = (v2 % ((v2) | 1))
                    }
                }
                _ = fi34
            }
            v4 = fi33
            _ = fi33
        }
    }
    v4 += (v4 | int(v2))
    var v35 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
