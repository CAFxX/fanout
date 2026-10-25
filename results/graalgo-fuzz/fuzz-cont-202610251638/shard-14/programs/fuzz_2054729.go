package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 1
    var v3 uint = 255
    v4 := []int{v1, (v1 << ((68) & 63)), v1, (v1 / ((v1) | 1))}
    v5 := GenS1{F0: 1}
    v2 += (^(v2))
    var v6 int16 = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = len(v4)
            switch "" {
            case "foo bar":
                if (true || true) {
                    switch "" {
                    case "hello":
                        v9 := GenS1{F0: 127}
                        _ = v9
                    case "\n\t":
                        v10 := GenS1{F0: -12}
                        _ = v10
                    case "a":
                        var v11 int = v1
                        _ = v11
                    default:
                        var v12 float32 = (3.14159 / 3.14159)
                        _ = v12
                    }
                    v3 -= v3
                    fmt.Println(fi8, fi8, v6)
                }
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                v2 = (v2 + v2)
            }
            _ = fi8
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (false || false) {
            if (((((32767 > -596) && false) || (false || true)) || (!(2 == 103))) || (((1e10 / 1.0) - 0.0) > 0.0)) {
                var v13 uint32 = 65535
                v6 += v6
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v2)
                _ = fi14
            }
        } else {
            var v15 float64 = 0.0
            fmt.Println(v15, v2, v3)
            _ = v15
        }
        _ = fi7
    }
    var v16 float64 = 3.14159
    v3 -= v3
    switch -960 {
    case 127:
        var v17 int32 = 438
        _ = v17
    default:
        v1 -= v1
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        switch "foo bar" {
        case "\n\t":
            v1 = 42
        case "foo bar":
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(v2)
                _ = fi19
            }
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            switch fi18 {
            case 32768:
                fmt.Println(v3, v1)
            default:
                if (false && ((v3 | (^(v3))) == (v3 << ((v3) & 63)))) {
                    v21 := GenS1{F0: -128}
                    if false {
                        v6 += v6
                        _ = len(v4)
                        switch "a" {
                        case "":
                            var v22 int8 = -103
                            _ = v22
                        default:
                            var v23 float32 = 0.1
                            _ = v23
                        }
                        v24 := GenS2{F0: 651, F1: 42, F2: true}
                        _ = v24
                    } else {
                        if false {
                            v4 = append(v4, v1)
                        }
                        var v25 string = "foo bar"
                        _ = v25
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v21
                }
            }
            if (v16 <= (float64(1.5) - v16)) {
                var v26 int8 = int8(v2)
                if ((v2 & 42) < (^(v2))) {
                    var v27 float64 = v16
                    _ = v27
                }
                _ = v26
            }
            _ = fi20
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi18
    }
    var v28 string = "\n\t"
    if len(v4) > 0 {
        v4[0] = (v1 & v1)
    }
    fmt.Println(v28, v6)
    if false {
        v2 = (v2 << ((v3) & 63))
        fmt.Println(v6, v16)
    }
    var v29 int = 13
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
