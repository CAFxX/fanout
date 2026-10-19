package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 int64, p1 string, p2 uint16) uint8 {
    v5 := GenS1{F0: 0, F1: 32768, F2: 100}
    var v6 uint = 611
    var v7 int8 = 1
    fmt.Println(v6, v7, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 string = "foo bar"
    var v3 int = 127
    v4 := make(map[string]uint32)
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(fi8, v3)
                v1 = (^(v1))
                _ = fi9
            }
        } else {
            switch v3 {
            case 42:
                v10 := []int{v3, fi8, (fi8 % ((fi8) | 1))}
                _ = v10
            default:
                var v11 float64 = 1e10
                _ = v11
            }
            v12 := GenS1{F0: genFn0(-449, (v2 + "foo bar"), 0), F1: (v3 & (fi8 & -135)), F2: -128}
            v13 := GenS2{F0: 4294967295, F1: 42}
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    if true {
        var v14 float32 = 0.0
        var v15 uint8 = 0
        _ = v14
        _ = v15
    }
    if false {
        switch v3 {
        case -1:
            switch (((((v2 + (v2 + v2)) + ("foo bar" + "hello")) + "a") + ((v2 + v2) + "")) + ("\n\t" + "foo bar")) {
            case "":
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 int16 = 32767
                    v1 += v1
                    var v18 uint64 = 2
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                var v19 float64 = 0.1
                _ = v19
            }
        case 32768:
            fmt.Println(v3, v1, v3)
        }
        switch ("" + "\n\t") {
        case "\n\t":
            fmt.Println(v2, v1, v2)
        case "hello":
            switch v3 {
            case -2147483648:
                var v20 uint16 = 723
                _ = v20
            case 9223372036854775807:
                for fi21 := 0; fi21 < 18; fi21++ {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        if true {
                            v23 := GenS1{F0: 165, F1: 32767, F2: -1}
                            _ = v23
                        }
                        v3 += fi21
                        _ = fi22
                    }
                    _ = fi21
                }
            case -287:
                v3 += 4294967295
            }
        case "foo bar":
            v24 := GenS1{F0: 255, F1: v3, F2: -126}
            _ = v24
        }
        v4[v2] = (v1 >> ((255) & 63))
    }
    v25 := GenS3{F0: v3, F1: 65535, F2: ((v2 + "") + "a")}
    v3 -= (v3 - (-(v3)))
    if (916 == 65535) {
        var v26 int64 = 42
        v27 := []int{89, (^(v3))}
        v2 = ((v2 + v2) + v2)
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 128, F1: (v3 / ((v3) | 1)), F2: -100}
    switch v3 {
    case 1:
        var v29 float64 = 0.1
        _ = v29
    case 127:
        for fi30 := 0; fi30 < 11; fi30++ {
            v1 = 2
            _ = fi30
        }
    default:
        var v31 uint32 = (v1 | (v1 | (v1 << ((255) & 63))))
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint32 = 806
        var v34 uint32 = (v33 + 352)
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 float64 = (-2.25 - 0.1)
    if (358 != (v1 >> ((18446744073709551615) & 63))) {
        fmt.Println(v35, v2, v1)
    } else {
        for fi36 := 0; fi36 < 6; fi36++ {
            v37 := GenS3{F0: (v3 % ((v3) | 1)), F1: 255, F2: v2}
            _ = fi36
            _ = v37
        }
    }
    switch (((("foo bar" + v2) + "hello") + v2) + ("foo bar" + v2)) {
    case "foo bar":
        v38 := GenS4{F0: 1, F1: (!true)}
        _ = v38
    default:
        switch (v3 << ((42) & 63)) {
        case -9223372036854775808:
            if false {
                if true {
                    switch (v2 + (v2 + v2)) {
                    case "foo bar":
                        var v39 int16 = -315
                        _ = v39
                    default:
                        fmt.Println(v2, v2)
                    }
                    v3 = (v3 ^ 4294967295)
                    var v40 uint32 = 42
                    _ = v40
                } else {
                    v1 += v1
                }
            } else {
                fmt.Println(v1, v3, v35)
                switch v3 {
                case 9223372036854775807:
                    v41 := []int{(v3 & v3), v3, 32768}
                    _ = v41
                case 1:
                    v42 := GenS5{F0: 1, F1: 45, F2: -444}
                    _ = v42
                default:
                    switch (v3 + 251) {
                    case 1:
                        for fi43 := 0; fi43 < 10; fi43++ {
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v45 := GenS1{F0: 255, F1: (fi44 << ((42) & 63)), F2: 100}
                                _ = fi44
                                _ = v45
                            }
                            if true {
                                v3 = (v3 % ((fi43) | 1))
                                switch v2 {
                                case "\n\t":
                                    v1 += (v1 >> ((4294967295) & 63))
                                case "foo bar":
                                    for fi46 := 0; fi46 < 18; fi46++ {
                                        var v47 int64 = -98
                                        var v48 uint64 = 42
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                fmt.Println(v3, fi43)
                            }
                            _ = fi43
                        }
                    case 2147483647:
                        var v49 uint64 = 484
                        _ = v49
                    case 42:
                        switch (v3 << ((2) & 63)) {
                        case 42:
                            var v50 uint8 = 96
                            _ = v50
                        }
                    default:
                        var v51 uint16 = 0
                        _ = v51
                    }
                }
                switch ((v2 + v2) + v2) {
                case "foo bar":
                    var v52 int64 = 476
                    _ = v52
                }
            }
        case 32768:
            v53 := GenS3{F0: -745, F1: 271, F2: v2}
            _ = v53
        }
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v35
    _ = v4
}
