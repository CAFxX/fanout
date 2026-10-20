package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int32 = 181
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := []int{4294967295, 9223372036854775807, 955}
    var v5 int32 = v1
    switch ("foo bar" + (("foo bar" + "") + "hello")) {
    case "a":
        v1 += (v5 << ((255) & 63))
    default:
        fmt.Println(v1)
    }
    v6 := GenS1{F0: 255, F1: (65535 > 670), F2: 42}
    v3[-181] = ""
    var v7 int8 = -1
    if false {
        var v8 uint = 506
        v5 = v1
        _ = v8
    } else {
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            case "hello":
                var v9 int = 127
                _ = v9
            case "a":
                switch "a" {
                case "a":
                    if false {
                        switch -1 {
                        case 4294967295:
                            for fi10 := 0; fi10 < 4; fi10++ {
                                if len(v4) > 0 {
                                    v4[0] = (fi10 | (^(fi10)))
                                }
                                fmt.Println(v1, v5)
                                _ = fi10
                            }
                        default:
                            if (2 >= 0) {
                                fmt.Println(v5)
                            } else {
                                fmt.Println(v1, v7, v7)
                                _ = len(v3)
                                v11 := GenS1{F0: 32767, F1: false, F2: 42}
                                _ = v11
                            }
                        }
                        switch "" {
                        case "":
                            fmt.Println(v7, v5, v5)
                        default:
                            v12 := v3[2147483647]
                            _ = v12
                        }
                    }
                case "\n\t":
                    var v13 int32 = -1
                    _ = v13
                }
            }
        case "foo bar":
            switch 42 {
            case -552:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v5 -= (v5 | (v1 % ((v1) | 1)))
                    var v15 int64 = 996
                    _ = fi14
                    _ = v15
                }
            case -2147483648:
                v5 -= (v1 * (v5 | (v5 - v1)))
            }
        }
        v16 := GenS1{F0: 2, F1: true, F2: 489}
        var v17 uint64 = 1
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        switch int(v7) {
        case 558:
            var v19 float32 = ((1e10 - -2.25) * 3.14159)
            _ = v19
        }
        fmt.Println(v5, v1, v1)
        _ = fi18
    }
    v5 = v1
    v7 = v7
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "\n\t" {
        case "\n\t":
            v5 = int32(fi20)
        }
        v21 := GenS2{F0: 4294967295, F1: 950}
        _ = fi20
        _ = v21
    }
    switch -69 {
    case 127:
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: 1, F1: (1.5 < 3.14159), F2: 65535}
                _ = fi23
                _ = v24
            }
            v1 = 2147483647
            _ = fi22
        }
    case 4294967295:
        if false {
            var v25 int = -9223372036854775808
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int64 = 9223372036854775807
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
