package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 uint, p1 uint64, p2 int) int8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = (p0 % ((255) | 1))
        _ = fi6
        _ = v7
    }
    v8 := []int{9223372036854775807, p2, 2147483647}
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(fi9, p2)
        var v10 int16 = 1000
        _ = fi9
        _ = v10
    }
    switch "hello" {
    case "\n\t":
        if ((0.0 <= 3.14159) || (!(false && (4294967295 >= 269)))) {
            switch "foo bar" {
            case "\n\t":
                v11 := GenS1{F0: -367, F1: 147, F2: 42}
                _ = v11
            case "hello":
                switch (p2 >> ((p0) & 63)) {
                case 4294967295:
                    v8 = append(v8, (p2 * 127))
                }
            }
        }
    }
    var v12 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return 1
}

func main() {
    var v1 uint = 360
    v2 := []int{-663, 128, 920}
    var v3 uint64 = 702
    var v4 int16 = -873
    var v5 uint64 = (v3 + 0)
    v13 := GenS1{F0: (v4 % (((v4 + 1)) | 1)), F1: 1, F2: 2}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        switch (fi14 - fi14) {
        case -466:
            for fi15 := 0; fi15 < 17; fi15++ {
                fmt.Println(v3, fi14)
                if len(v2) > 0 {
                    v2[0] = fi14
                }
                _ = fi15
            }
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = ((v3 ^ 1) >= (^(v3)))
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "foo bar":
        var v18 int16 = (v4 / ((0) | 1))
        _ = v18
    }
    var v19 string = ("" + ("\n\t" + "\n\t"))
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint16 = uint16(v1)
        _ = len(v2)
        v5 -= uint64(v1)
        _ = fi20
        _ = v21
    }
    _ = len(v2)
    v3 = v3
    v3 = 180
    fmt.Println(v3)
    fmt.Println(v5)
    switch v19 {
    case "hello":
        v4 -= (v4 * 32767)
    case "":
        if (!(!false)) {
            v2 = append(v2, 2147483647)
        }
    default:
        for fi22 := 0; fi22 < 19; fi22++ {
            switch fi22 {
            case 42:
                v23 := GenS1{F0: (v4 | (v4 * (-(v4)))), F1: 0, F2: 42}
                _ = v23
            case -2147483648:
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := GenS1{F0: v4, F1: -2147483648, F2: 212}
                    _ = fi24
                    _ = v25
                }
            case 32767:
                fmt.Println(v4, fi22)
            }
            _ = fi22
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v2 = append(v2, fi26)
        switch "a" {
        case "a":
            v5 = v3
        case "foo bar":
            v3 -= (^(v3))
        case "\n\t":
            v5 = v3
        }
        v19 = ((v19 + v19) + "a")
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
