package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
func genFn0() string {
    if (true && (true || (-128 > 42))) {
        var v7 float64 = 3.14159
        v8 := []int{128, 42, -1}
        v7 -= v7
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 127
        fmt.Println(v9)
        _ = v9
    }
    var v10 int64 = 0
    _ = v10
    return (("a" + ("foo bar" + ("foo bar" + "a"))) + "foo bar")
}
func genFn1(p0 uint16, p1 uint8) uint8 {
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v11 := []int{42}
        _ = v11
    default:
        p0 -= (p0 ^ p0)
    }
    v12 := []int{186, 249, 4294967295}
    _ = len(v12)
    switch (("" + ("\n\t" + "hello")) + "") {
    case "foo bar":
        v13 := GenS1{F0: (p1 | (p1 % ((p1) | 1)))}
        _ = v13
    case "":
        if ((("hello" + ("a" + "")) + "\n\t") < "\n\t") {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS1{F0: p1}
                if (864 == 1) {
                    p1 += 255
                    var v16 int32 = -2147483648
                    _ = v16
                }
                if (0 >= -985) {
                    var v17 float64 = -2.25
                    _ = v17
                } else {
                    p0 = 2
                    var v18 uint64 = 42
                    var v19 uint16 = p0
                    _ = v18
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            if ((("\n\t" + "\n\t") + "\n\t") != "\n\t") {
                if (!true) {
                    p1 += (p1 & 0)
                    switch "foo bar" {
                    case "":
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case "a":
                        if (("\n\t" + "\n\t") != "a") {
                            p1 = (p1 ^ (p1 | (p1 << ((102) & 63))))
                            p1 = (^(p1))
                        }
                    }
                }
                var v20 int16 = 1000
                if len(v12) > 0 {
                    v12[0] = -2147483648
                }
                _ = v20
            }
            var v21 bool = (!false)
            _ = v21
        }
    case "a":
        var v22 int64 = 0
        _ = v22
    }
    p1 += 2
    _ = p0
    _ = p1
    _ = v12
    return p1
}

func main() {
    var v1 float64 = (0.0 / 1.0)
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[int]bool)
    v4 := []uint32{}
    var v5 int32 = -2147483648
    v6 := make(map[int]bool)
    switch "\n\t" {
    case "\n\t":
        v23 := GenS2{F0: 42, F1: 897, F2: 728}
        _ = v23
    }
    switch 923 {
    case -453:
        if (18446744073709551615 <= 0) {
            switch 867 {
            case 32768:
                _ = len(v4)
            }
        }
    default:
        for fi24 := 0; fi24 < 15; fi24++ {
            if true {
                v5 = (-(v5))
                var v25 int8 = 25
                _ = v25
            } else {
                v2 = ("\n\t" + "")
                v2 = v2
            }
            _ = fi24
        }
    }
    var v26 float64 = 0.0
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS1{F0: 255}
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int8 = 2
            v31 := GenS1{F0: 0}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v1, v1)
    v5 = v5
    fmt.Println(v26)
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS1{F0: 127}
        var v34 float32 = 1.5
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: genFn1(65535, 1), F1: 42}
    v5 = 0
    switch -764 {
    case -99:
        for fi36 := 0; fi36 < 3; fi36++ {
            fmt.Println(v5)
            _ = fi36
        }
    case 377:
        switch 32768 {
        case 1:
            switch v2 {
            case "\n\t":
                v26 = (((1.5 * (v1 + (-2.25 + (0.1 * 0.1)))) + (float64((1e10 * (0.1 + 1e10))) / (v1 * v26))) + 3.14159)
            default:
                v37 := GenS4{F0: 0}
                _ = v37
            }
        }
    case 2147483647:
        v38 := GenS5{F0: 32767}
        _ = v38
    default:
        var v39 int32 = (-(v5))
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS2{F0: 0, F1: 255, F2: 42}
            v42 := GenS6{F0: ((!(726 == 570)) || (false && false))}
            v4 = append(v4, 200)
            _ = fi40
            _ = v41
            _ = v42
        }
    } else {
        for fi43 := 0; fi43 < 18; fi43++ {
            var v44 int32 = (-(v5))
            if (0 != -288) {
                if true {
                    v26 = 1e10
                    v44 = -926
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v45 := GenS3{F0: 25, F1: 255}
                    var v46 uint16 = 255
                    v26 += v26
                    _ = v45
                    _ = v46
                }
            } else {
                var v47 float64 = ((1e10 - (v1 / (float64(((1e10 - 0.0) - 3.14159)) / (v26 / (v26 + 1e10))))) / v26)
                _ = v47
            }
            _ = fi43
            _ = v44
        }
        v48 := v6[-244]
        var v49 uint = 4294967295
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
