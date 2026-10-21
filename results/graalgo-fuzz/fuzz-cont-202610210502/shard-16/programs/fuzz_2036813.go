package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []string{""}
    var v2 float64 = 1.5
    var v3 int = -913
    var v4 uint32 = 1
    var v5 uint32 = (^(v4))
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "\n\t" {
            case "foo bar":
                v9 := GenS1{F0: 0, F1: (v3 << ((0) & 63)), F2: ""}
                _ = v9
            case "\n\t":
                var v10 int16 = 42
                _ = v10
            default:
                fmt.Println(v4, v5)
            }
            _ = fi8
        }
        _ = fi7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS2{F0: true, F1: 42}
    if (95 > 1) {
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v14 := GenS2{F0: true, F1: 1000}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        switch "hello" {
        case "hello":
            v15 := GenS2{F0: (!false), F1: 42}
            _ = v15
        case "":
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS2{F0: true, F1: -32768}
                _ = fi16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS3{F0: 1}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint64 = 2
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 uint8 = uint8(v21)
            v24 := GenS3{F0: 32767}
            var v25 uint8 = 66
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS4{F0: 0, F1: 20, F2: ""}
            v29 := GenS3{F0: 1}
            var v30 uint = 255
            v31 := GenS5{F0: (^(v30)), F1: 42, F2: 42}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
            _ = v31
        }
        switch (-(fi26)) {
        case 318:
            if (!(42 < 0)) {
                switch (v3 << ((0) & 63)) {
                case -2147483648:
                    fmt.Println(v3, v3, fi26)
                case -997:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v5 -= v4
                        v4 += v5
                        if false {
                            v4 += (v6 ^ 0)
                        } else {
                            v5 -= (v6 % ((0) | 1))
                            fmt.Println(v3, v3, fi26)
                        }
                        _ = fi32
                    }
                }
                for fi33 := 0; fi33 < 14; fi33++ {
                    v5 += 1
                    _ = fi33
                }
            } else {
                v5 = 2
                if (false || true) {
                    v34 := GenS6{F0: 2}
                    _ = v34
                }
            }
        case 2147483647:
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 int64 = 9223372036854775807
                var v37 float64 = v2
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        _ = fi26
    }
    var v38 int8 = 100
    if false {
        v5 += 4294967295
        switch (v3 | (v3 + v3)) {
        case 2147483647:
            v1 = append(v1, "\n\t")
        }
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "\n\t")
    }
    var v39 float64 = v2
    v3 = (^(v3))
    var v40 int8 = (-(v38))
    switch "a" {
    case "\n\t":
        for fi41 := 0; fi41 < 6; fi41++ {
            if false {
                v2 -= float64((-2.25 / 0.1))
                var v42 string = ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "")
                _ = v42
            }
            _ = fi41
        }
    default:
        v43 := GenS7{F0: (v38 % (((v38 >> ((255) & 63))) | 1))}
        _ = v43
    }
    if false {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "a")
        }
        for fi44 := 0; fi44 < 4; fi44++ {
            fmt.Println(v40)
            if (false && (!false)) {
                fmt.Println(v5)
            }
            _ = fi44
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v45 := GenS8{F0: 255, F1: (v38 & v40)}
    _ = len(v1)
    var v46 float64 = (v2 / v39)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
