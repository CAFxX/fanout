package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float32 = (0.1 / float32((((-2.25 / 3.14159) * 1.5) / -2.25)))
    var v2 uint64 = 941
    v3 := []int{-62}
    var v4 int8 = 0
    var v5 int8 = v4
    var v6 int16 = 0
    v1 = (-2.25 - ((v1 * float32(0.0)) / 1e10))
    if false {
        var v7 uint32 = 4294967295
        _ = v7
    }
    v3 = append(v3, 1)
    switch 4294967295 {
    case -415:
        var v8 float32 = 1e10
        _ = v8
    case 4294967295:
        v9 := GenS1{F0: true, F1: 524, F2: v2}
        _ = v9
    case 1:
        if (65535 == 0) {
            if (false && true) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    _ = len(v3)
                    switch "" {
                    case "foo bar":
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v12 := GenS2{F0: 1}
                            _ = fi11
                            _ = v12
                        }
                    case "\n\t":
                        fmt.Println(v5, v6, v5)
                    case "":
                        switch "foo bar" {
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v3)
                            }
                        }
                    default:
                        var v13 float64 = ((1.5 * (0.0 * -2.25)) / 1.0)
                        _ = v13
                    }
                    switch "\n\t" {
                    case "hello":
                        var v14 uint = 255
                        _ = v14
                    }
                    _ = fi10
                }
            } else {
                if true {
                    v4 += 77
                    if false {
                        switch ("hello" + ("foo bar" + ("hello" + "\n\t"))) {
                        case "":
                            v2 = v2
                        case "a":
                            v5 -= 100
                        case "foo bar":
                            _ = len(v3)
                        default:
                            for fi15 := 0; fi15 < 8; fi15++ {
                                var v16 int16 = v6
                                v3 = append(v3, -1)
                                v16 += 0
                                _ = fi15
                                _ = v16
                            }
                        }
                    }
                    switch 32767 {
                    case 32768:
                        _ = len(v3)
                    case 137:
                        if false {
                            var v17 int8 = (v5 >> ((18446744073709551615) & 63))
                            _ = v17
                        }
                    case 128:
                        for fi18 := 0; fi18 < 14; fi18++ {
                            _ = len(v3)
                            _ = fi18
                        }
                    default:
                        var v19 float32 = ((3.14159 / (3.14159 * float32((-2.25 / 1e10)))) - v1)
                        _ = v19
                    }
                }
            }
            v3 = append(v3, -9223372036854775808)
        }
    default:
        v4 = (-(v5))
    }
    _ = len(v3)
    v20 := GenS1{F0: (((v1 * v1) + v1) >= v1), F1: 49, F2: 2}
    fmt.Println(v1, v2, v5)
    v5 += (v5 >> ((85) & 63))
    if ((-(v6)) == 42) {
        var v21 string = "foo bar"
        _ = len(v3)
        switch 42 {
        case 4294967295:
            v3 = append(v3, 32767)
        case 796:
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 bool = (0 > 1)
                v24 := GenS3{F0: 1}
                for fi25 := 0; fi25 < 17; fi25++ {
                    v2 += (v2 | v2)
                    v4 -= v4
                    _ = fi25
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        case -607:
            fmt.Println(v4)
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
