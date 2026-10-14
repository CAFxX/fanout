package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int32 = -849
    var v2 uint16 = 0
    var v3 float32 = -2.25
    var v4 int = -9223372036854775808
    v5 := []int{(v4 << ((534) & 63)), 9223372036854775807, 9223372036854775807, 9223372036854775807}
    if (918 >= v2) {
        v3 -= ((v3 / float32(1e10)) - v3)
        v6 := GenS1{F0: (v2 % (((v2 << ((1) & 63))) | 1)), F1: 2}
        _ = v6
    }
    if ((((346 > 2) || false) || true) && false) {
        if true {
            if true {
                v7 := GenS1{F0: v2, F1: 0}
                if (161 == 81) {
                    v4 += v4
                    for fi8 := 0; fi8 < 3; fi8++ {
                        fmt.Println(fi8, v4)
                        _ = fi8
                    }
                    for fi9 := 0; fi9 < 13; fi9++ {
                        var v10 uint = 0
                        var v11 int32 = v1
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                } else {
                    v12 := GenS1{F0: v2, F1: 0}
                    fmt.Println(v2, v2, v3)
                    _ = v12
                }
                v13 := GenS2{F0: 1, F1: 210}
                _ = v13
                _ = v7
            }
            switch "" {
            case "hello":
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: 4294967295, F1: 4294967295}
                    v1 += v1
                    _ = fi14
                    _ = v15
                }
            default:
                v16 := GenS2{F0: 2147483647, F1: 18446744073709551615}
                _ = v16
            }
        }
        if ((false || true) && false) {
            var v17 uint8 = 2
            v1 = (-(v1))
            for fi18 := 0; fi18 < 17; fi18++ {
                v2 += v2
                v4 = v4
                _ = fi18
            }
            _ = v17
        } else {
            var v19 uint = 42
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint8 = 42
            var v22 float32 = v3
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v4 = (v4 * v4)
    switch "\n\t" {
    case "foo bar":
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 uint32 = 0
            v24 += (v24 | v24)
            _ = fi23
            _ = v24
        }
    case "\n\t":
        fmt.Println(v2, v4, v4)
    default:
        v25 := GenS3{F0: 127}
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            _ = len(v5)
            v28 := GenS1{F0: v2, F1: 0}
            _ = fi27
            _ = v28
        }
        if (((101 <= 84) || true) && (("hello" + "") < (("a" + ("hello" + ("\n\t" + ("a" + "foo bar")))) + "a"))) {
            fmt.Println(v1)
        } else {
            fmt.Println(v2, fi26, v2)
            switch "\n\t" {
            case "\n\t":
                var v29 float64 = (1e10 + (0.0 * -2.25))
                _ = v29
            case "a":
                if (2147483647 <= (v4 | fi26)) {
                    switch v4 {
                    case 32767:
                        var v30 uint = 152
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 15; fi31++ {
                            v5 = append(v5, v4)
                            v2 -= (v2 / ((255) | 1))
                            _ = fi31
                        }
                    }
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 int = v4
                        var v34 bool = true
                        switch "a" {
                        case "\n\t":
                            var v35 int16 = -1
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
            }
        }
        _ = fi26
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := GenS3{F0: 124}
        switch fi36 {
        case -718:
            v3 = 1e10
        }
        _ = fi36
        _ = v37
    }
    v2 -= v2
    v2 = v2
    if (1e10 == 0.1) {
        for fi38 := 0; fi38 < 20; fi38++ {
            switch (fi38 >> ((366) & 63)) {
            case 9223372036854775807:
                var v39 uint = 1
                _ = v39
            case 1:
                v4 += fi38
            case 2147483647:
                var v40 int8 = 42
                _ = v40
            }
            _ = fi38
        }
    } else {
        var v41 uint64 = 945
        _ = v41
    }
    v3 += float32(3.14159)
    var v42 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
