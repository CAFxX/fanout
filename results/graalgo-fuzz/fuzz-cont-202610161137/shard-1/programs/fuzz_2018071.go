package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    v1 := []int{32767, -9223372036854775808, 128, -352}
    v2 := []int{4294967295, 1}
    var v3 uint16 = 42
    v4 := []int{}
    var v5 uint = 18446744073709551615
    var v6 float64 = 0.1
    if len(v4) > 0 {
        v4[0] = 42
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        switch "a" {
        case "foo bar":
            var v8 uint = 1
            _ = v8
        case "a":
            var v9 int = fi7
            _ = v9
        case "":
            var v10 uint8 = 140
            _ = v10
        default:
            if (v6 > v6) {
                _ = len(v1)
                var v11 uint64 = 0
                _ = v11
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int32 = -2147483648
                    _ = fi12
                    _ = v13
                }
            }
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        switch (("foo bar" + "\n\t") + "\n\t") {
        case "":
            switch "a" {
            case "hello":
                if ((!false) || false) {
                    var v14 uint = v5
                    switch "\n\t" {
                    case "a":
                        v5 -= (v14 + 18446744073709551615)
                    case "":
                        var v15 bool = true
                        _ = v15
                    }
                    v4 = append(v4, 127)
                    _ = v14
                } else {
                    v6 += v6
                    fmt.Println(v3)
                    var v16 int = -686
                    _ = v16
                }
            case "\n\t":
                var v17 uint32 = 1
                _ = v17
            case "":
                for fi18 := 0; fi18 < 16; fi18++ {
                    if (1e10 == v6) {
                        v6 += (v6 / v6)
                        _ = len(v2)
                    }
                    _ = fi18
                }
            }
        case "foo bar":
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 uint64 = 1
                switch (fi19 * (fi19 >> (((v5 | v5)) & 63))) {
                case 32767:
                    switch "foo bar" {
                    case "a":
                        v3 = (^(v3))
                    case "foo bar":
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v22 := GenS1{F0: 42, F1: 42}
                            fmt.Println(v5, v20)
                            var v23 float64 = float64((0.0 + (1.5 + 0.0)))
                            v24 := GenS1{F0: 0, F1: 2147483647}
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    case "hello":
                        var v25 int32 = -1
                        _ = v25
                    default:
                        _ = len(v4)
                    }
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    fmt.Println(v5, fi19)
                    v27 := GenS2{F0: (v5 ^ 42), F1: 1}
                    v20 = v20
                    v2 = append(v2, 9223372036854775807)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi19
                _ = v20
            }
        case "\n\t":
            v6 = (v6 / v6)
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if len(v2) > 0 {
            v2[0] = 99
        }
    } else {
        v5 = (^(v5))
        fmt.Println(v3)
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            switch "hello" {
            case "foo bar":
                switch ("a" + "") {
                case "foo bar":
                    v6 -= (v6 * (v6 * (v6 / (v6 / 1.0))))
                default:
                    var v30 float32 = -2.25
                    _ = v30
                }
            }
            _ = len(v4)
            switch -520 {
            case 2147483647:
                v6 -= ((3.14159 + (float64(0.1) - v6)) / v6)
            case 42:
                switch "foo bar" {
                case "hello":
                    v6 += v6
                case "\n\t":
                    var v31 int16 = -1
                    _ = v31
                case "a":
                    v32 := GenS1{F0: 65535, F1: 594}
                    _ = v32
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            default:
                if false {
                    var v33 int = fi28
                    v3 = (v3 - 288)
                    if false {
                        v33 = fi29
                    } else {
                        var v34 uint16 = v3
                        _ = v34
                    }
                    var v35 int16 = 322
                    var v36 bool = true
                    _ = v33
                    _ = v35
                    _ = v36
                }
            }
            _ = fi29
        }
        v5 = (v5 % ((751) | 1))
        _ = fi28
    }
    var v37 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
