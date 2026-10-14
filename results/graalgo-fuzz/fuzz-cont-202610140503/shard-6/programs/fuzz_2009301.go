package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint) int64 {
    v4 := GenS1{F0: 255, F1: 18446744073709551615}
    for fi5 := 0; fi5 < 15; fi5++ {
        switch fi5 {
        case -1:
            fmt.Println(p1, p0)
        case 127:
            p1 += (p1 + p1)
        default:
            p0 -= -9223372036854775808
        }
        _ = fi5
    }
    var v6 bool = false
    v7 := GenS2{F0: -84, F1: -1}
    for fi8 := 0; fi8 < 14; fi8++ {
        if (v6 && (p0 != 9223372036854775807)) {
            v9 := GenS1{F0: 620, F1: 455}
            _ = v9
        } else {
            var v10 string = "foo bar"
            v6 = v6
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{128, 40}
            var v13 float64 = (((((1e10 * (3.14159 * float64(3.14159))) / float64(3.14159)) + 0.1) * 3.14159) / -2.25)
            fmt.Println(fi11)
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int8 = int8(p0)
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return (p0 & 42)
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar", "foo bar"}
    v3 := make(map[string]int)
    var v15 bool = (180 < 2)
    if len(v2) > 0 {
        v2[0] = ("hello" + "foo bar")
    }
    if ((!v15) && v15) {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 int32 = -1
                _ = fi17
                _ = v18
            }
            v15 = true
            v2 = append(v2, "\n\t")
            _ = fi16
        }
    }
    v2 = append(v2, "\n\t")
    switch "\n\t" {
    case "\n\t":
        v19 := v3[("hello" + "a")]
        _ = v19
    }
    if v15 {
        if (!false) {
            var v20 int32 = -89
            var v21 int16 = -903
            _ = v20
            _ = v21
        } else {
            v15 = (v15 && v15)
            v2 = append(v2, ("" + "\n\t"))
        }
        v22 := GenS3{F0: 42}
        switch "hello" {
        case "foo bar":
            var v23 float32 = (3.14159 - (1.5 * float32(float64(0.0))))
            _ = v23
        case "hello":
            var v24 int8 = -4
            _ = v24
        case "a":
            v15 = (18446744073709551615 != 2)
        default:
            if v15 {
                var v25 int8 = -84
                for fi26 := 0; fi26 < 2; fi26++ {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (!(43 <= 42)) {
                            if true {
                                fmt.Println(v25, fi26)
                            }
                            fmt.Println(v15)
                        }
                        var v28 bool = (!((v25 * (v25 - -53)) < v25))
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
                _ = v25
            }
        }
        _ = v22
    } else {
        switch "foo bar" {
        case "foo bar":
            switch "\n\t" {
            case "a":
                var v29 float32 = (1.5 - -2.25)
                _ = v29
            default:
                var v30 int8 = -31
                _ = v30
            }
        default:
            fmt.Println(v15)
        }
    }
    v15 = (!(127 >= 90))
    for fi31 := 0; fi31 < 9; fi31++ {
        if v15 {
            v15 = (367 > 447)
            if (!(v15 || (42 < 65535))) {
                var v32 uint8 = 52
                _ = len(v3)
                v33 := GenS4{F0: (v32 >> ((1) & 63))}
                _ = v32
                _ = v33
            } else {
                v15 = (2 < 42)
                fmt.Println(fi31, v15)
                fmt.Println(v15, fi31)
            }
            var v34 bool = (v15 && v15)
            _ = v34
        } else {
            v35 := GenS3{F0: 0}
            var v36 uint = 255
            _ = v35
            _ = v36
        }
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
