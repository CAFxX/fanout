package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 string }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int64, p2 uint) bool {
    switch 127 {
    case 115:
        fmt.Println(p0, p0)
    case 534:
        switch "hello" {
        case "hello":
            v5 := []int{-1, 32767, 0, 1}
            _ = v5
        default:
            p1 = (p1 << ((p2) & 63))
        }
    case 574:
        p1 = (p1 + 0)
    default:
        v6 := []int{0, 4294967295, 726, -235}
        _ = v6
    }
    v7 := []int{127, -9223372036854775808, -2147483648}
    var v8 int = 127
    fmt.Println(v8, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 uint8 = 30
    var v4 string = "\n\t"
    v4 = "hello"
    switch v4 {
    case "foo bar":
        fmt.Println(v4)
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(fi9, fi9)
            _ = fi9
        }
    }
    var v10 int64 = 42
    if false {
        v3 = (^(v3))
    } else {
        var v11 float64 = 1.5
        v12 := GenS1{F0: (v3 - (v3 + (v3 % (((^(v3))) | 1)))), F1: 1}
        _ = v11
        _ = v12
    }
    v13 := []int{866}
    var v14 uint32 = 4294967295
    v15 := v2[-2147483648]
    v10 -= (^(v10))
    switch 280 {
    case 32767:
        v15 = false
    case -440:
        v16 := GenS2{F0: v3}
        _ = v16
    case 0:
        var v17 uint8 = (v3 * 1)
        _ = v17
    }
    v18 := GenS3{F0: 517, F1: 1, F2: 296}
    fmt.Println(v3, v14)
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v13) > 0 {
            v13[0] = fi19
        }
        _ = fi19
    }
    v20 := GenS4{F0: 65535, F1: 1, F2: 252}
    v21 := GenS2{F0: v3}
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int8 = -31
            fmt.Println(v4, v10, v4)
            _ = fi23
            _ = v24
        }
        switch (((v4 + ("" + v4)) + v4) + "") {
        case "a":
            v13 = append(v13, 990)
        case "hello":
            for fi25 := 0; fi25 < 7; fi25++ {
                v13 = append(v13, (fi25 >> ((0) & 63)))
                v26 := GenS5{F0: 4294967295, F1: 100, F2: v4}
                _ = fi25
                _ = v26
            }
        case "":
            v14 = 416
        default:
            switch "foo bar" {
            case "\n\t":
                if (v15 && (-1 <= v10)) {
                    v27 := v2[fi22]
                    _ = v27
                }
            case "hello":
                v14 -= v14
            case "foo bar":
                if v15 {
                    v14 = v14
                    for fi28 := 0; fi28 < 14; fi28++ {
                        fmt.Println(v4)
                        v29 := GenS6{F0: v14}
                        _ = fi28
                        _ = v29
                    }
                }
            default:
                var v30 string = v4
                _ = v30
            }
        }
        if (693 >= 32767) {
            for fi31 := 0; fi31 < 11; fi31++ {
                fmt.Println(fi22, v15, v3)
                var v32 int64 = (v10 & v10)
                v15 = (!false)
                v4 = v4
                v4 = (("hello" + ("hello" + v4)) + ("" + (("\n\t" + v4) + "foo bar")))
                switch "a" {
                case "a":
                    var v33 bool = (v15 && (("\n\t" + "foo bar") != ""))
                    _ = v33
                case "hello":
                    fmt.Println(fi31)
                case "\n\t":
                    if (v14 > v14) {
                        v14 = (v14 + v14)
                        fmt.Println(v15, v10, fi22)
                    } else {
                        v34 := GenS6{F0: v14}
                        v32 = (^(v10))
                        v32 -= (v32 ^ 9223372036854775807)
                        _ = v34
                    }
                default:
                    v10 += (v32 + 9223372036854775807)
                }
                _ = fi31
                _ = v32
            }
        } else {
            v14 -= (^(v14))
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                switch (fi35 / ((1) | 1)) {
                case 346:
                    var v36 uint = uint(v3)
                    _ = v36
                default:
                    v10 = (v10 & v10)
                }
                _ = fi35
            }
        }
        _ = fi22
    }
    var v37 string = ""
    var v38 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
