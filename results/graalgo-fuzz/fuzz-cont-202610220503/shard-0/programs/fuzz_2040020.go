package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 uint) uint {
    if false {
        var v5 float32 = (1e10 / 1.0)
        _ = v5
    }
    var v6 int16 = p1
    var v7 int8 = 43
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 + 982)
}

func main() {
    v1 := []int{880}
    var v2 bool = (255 < 42)
    var v3 float32 = 1.5
    var v4 float32 = v3
    for fi8 := 0; fi8 < 6; fi8++ {
        v3 += (v3 + v3)
        if ((1.5 == ((-2.25 / 0.1) / float64(v4))) && true) {
            v3 += (v4 + 1e10)
            v4 -= float32((1e10 / (0.0 - ((((1.5 / ((-2.25 * 3.14159) - (0.1 / 3.14159))) + 0.1) / (((float64(v3) / 1.5) / 0.1) + -2.25)) - ((1.5 / (0.0 - (0.1 / 1e10))) / ((0.0 - 1e10) * (0.1 + 1.5)))))))
        }
        v9 := GenS1{F0: 0, F1: "\n\t", F2: 700}
        _ = fi8
        _ = v9
    }
    var v10 int64 = 42
    switch -635 {
    case -9223372036854775808:
        _ = len(v1)
    case 931:
        v3 -= (0.1 / float32(0.1))
    default:
        var v11 uint8 = 128
        _ = v11
    }
    var v12 int16 = -32768
    if ((("foo bar" + "a") > "foo bar") || (!((1e10 == (1e10 / ((((0.0 / float64(float32(3.14159))) - 0.0) - 1e10) / 3.14159))) || false))) {
        switch (("foo bar" + "a") + "") {
        case "a":
            if (int64(v12) <= (v10 & 1)) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi13
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    _ = len(v1)
                    _ = fi14
                }
                var v15 uint64 = 928
                _ = v15
            } else {
                v10 = (v10 / (((v10 | v10)) | 1))
                var v16 bool = (462 < 173)
                for fi17 := 0; fi17 < 10; fi17++ {
                    if ((0.0 + float64(v4)) > -2.25) {
                        var v18 uint16 = 434
                        if (float32((float64((((1e10 + (1e10 / v4)) * 0.0) + (v4 / 1.0))) * 0.0)) == float32(float64(v4))) {
                            var v19 uint64 = 1
                            _ = v19
                        }
                        _ = v18
                    } else {
                        v4 = float32(((3.14159 * 1.5) / ((1.5 - (0.0 - 0.1)) * 3.14159)))
                        _ = len(v1)
                    }
                    _ = fi17
                }
                _ = v16
            }
        case "":
            v20 := GenS2{F0: 4294967295, F1: 188}
            _ = v20
        case "\n\t":
            var v21 uint32 = 0
            _ = v21
        default:
            v22 := GenS2{F0: 1, F1: 1}
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            if v2 {
                v4 -= 0.0
                fmt.Println(fi23)
                var v24 uint8 = 1
                _ = v24
            } else {
                switch "a" {
                case "foo bar":
                    v25 := GenS1{F0: 1, F1: ("hello" + ("a" + (("a" + ("\n\t" + "a")) + ("a" + "hello")))), F2: 0}
                    _ = v25
                }
                for fi26 := 0; fi26 < 13; fi26++ {
                    switch (fi26 * (fi23 / ((fi26) | 1))) {
                    case 32767:
                        var v27 float32 = 0.1
                        _ = v27
                    default:
                        v10 = int64(fi23)
                    }
                    if v2 {
                        v4 -= v4
                        switch (fi23 & fi26) {
                        case 977:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    var v28 uint16 = 32767
                    fmt.Println(v28, v12)
                    if len(v1) > 0 {
                        v1[0] = (fi26 >> ((673) & 63))
                    }
                    _ = fi26
                    _ = v28
                }
                var v29 uint16 = 878
                _ = v29
            }
            v30 := GenS3{F0: v10}
            _ = fi23
            _ = v30
        }
    } else {
        var v31 float64 = ((float64(v3) - (1e10 / 1.5)) / 1.5)
        _ = v31
    }
    v1 = append(v1, 32768)
    fmt.Println(v12, v10, v12)
    v12 -= (v12 ^ v12)
    fmt.Println(v10, v4)
    fmt.Println(v2, v3)
    v3 = (v3 * 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
