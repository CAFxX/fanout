package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int32) uint32 {
    v6 := []int{2147483647}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 631
}
func genFn1(p0 int64, p1 uint32, p2 uint64) int32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{231, fi7, 42, (fi7 - (fi7 - fi7))}
        var v9 uint16 = 255
        for fi10 := 0; fi10 < 11; fi10++ {
            switch "a" {
            case "\n\t":
                v11 := GenS1{F0: (!(((!false) || (uint32(v9) != 348)) || (479 >= 1))), F1: p0, F2: 127}
                _ = v11
            case "a":
                v12 := GenS1{F0: true, F1: (p0 + p0), F2: int8(p2)}
                _ = v12
            case "":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            p0 = (p0 + (p0 >> ((2) & 63)))
            switch ((("\n\t" + "hello") + "hello") + (("" + ("\n\t" + "a")) + "\n\t")) {
            case "\n\t":
                v8 = append(v8, 2147483647)
            case "foo bar":
                v13 := GenS2{F0: -178, F1: 2, F2: (p2 + (^(p2)))}
                _ = v13
            }
            v9 = (v9 & 874)
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v14 := GenS2{F0: -2147483648, F1: 963, F2: 4294967295}
    v15 := GenS3{F0: p0, F1: 1, F2: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return -1
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 2
    v3 := make(map[uint32]uint32)
    v4 := []int{9223372036854775807, 127, 127, -2147483648}
    var v5 int = 128
    v16 := GenS4{F0: genFn0(-2147483648)}
    fmt.Println(v2, v5)
    v17 := GenS2{F0: 2147483647, F1: 4294967295, F2: 551}
    fmt.Println(v2, v2)
    if (false || false) {
        for fi18 := 0; fi18 < 4; fi18++ {
            v5 += -2147483648
            fmt.Println(fi18, v1, fi18)
            var v19 int16 = -1
            _ = fi18
            _ = v19
        }
    }
    _ = len(v4)
    for fi20 := 0; fi20 < 11; fi20++ {
        switch v5 {
        case -9223372036854775808:
            v5 = 155
        case 780:
            var v21 string = ("foo bar" + "\n\t")
            _ = v21
        case -2147483648:
            v2 = 424
        }
        _ = len(v4)
        fmt.Println(v5, v5)
        _ = fi20
    }
    fmt.Println(v1, v5)
    for fi22 := 0; fi22 < 6; fi22++ {
        switch "\n\t" {
        case "foo bar":
            v23 := GenS4{F0: uint32(v1)}
            _ = v23
        }
        _ = fi22
    }
    switch v5 {
    case 197:
        if false {
            var v24 int16 = 573
            _ = v24
        }
    case -9223372036854775808:
        var v25 uint8 = v1
        _ = v25
    case 128:
        if false {
            v4 = append(v4, -246)
            switch "foo bar" {
            case "a":
                var v26 int32 = 1
                _ = v26
            case "foo bar":
                switch -629 {
                case -1:
                    v1 = (^(v1))
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
