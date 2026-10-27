package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint16) bool {
    p0 -= (p0 >> ((255) & 63))
    var v6 int16 = 886
    v7 := []int{128, 2147483647, -794, -939}
    if false {
        var v8 int16 = v6
        var v9 uint = 18446744073709551615
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: "foo bar", F1: p0}
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 int64 = 688
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    v4 := []int{90}
    var v5 uint8 = v2
    if (!(!(1 < -128))) {
        var v11 float64 = (3.14159 + (((((float64((1.5 * 0.0)) + -2.25) - 1e10) / ((float64(((1e10 / 1e10) - (float32((float64(3.14159) + (0.0 + (0.0 / 1e10)))) / 0.1))) + 0.0) * 1.5)) / ((-2.25 + ((0.1 - (float64(1e10) - 1e10)) / (((3.14159 / 1.5) / 3.14159) / -2.25))) + 0.1)) / (0.1 + (3.14159 + (float64(0.1) / (1.5 + 3.14159))))))
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 float64 = 1e10
        v1 = -1
        var v14 float64 = v13
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v2)
    v15 := GenS2{F0: 42, F1: 751}
    v5 = (v5 << ((2) & 63))
    _ = len(v4)
    v16 := v3[""]
    var v17 int = -1
    if false {
        v18 := GenS3{F0: 255, F1: 2147483647}
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(v1)
            _ = fi19
        }
        _ = len(v4)
        _ = v18
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            v1 = v1
            fmt.Println(v1)
            _ = fi20
        }
        v1 += (v1 % (((v1 | int64(v2))) | 1))
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint8 = 127
        v23 := GenS1{F0: "", F1: 42}
        if v16 {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 uint8 = (^(v2))
                var v26 bool = false
                v2 -= 255
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 uint8 = v5
            fmt.Println(v1, v16)
            _ = v27
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if (v16 || (18446744073709551615 <= 1)) {
        v28 := GenS4{F0: 2, F1: v1}
        _ = v28
    }
    switch "foo bar" {
    case "a":
        switch ("foo bar" + "") {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v29 string = "\n\t"
            _ = v29
        }
    case "":
        v16 = false
    case "\n\t":
        v1 = 1
    }
    if (v1 == (v1 ^ 9223372036854775807)) {
        var v30 uint32 = 748
        for fi31 := 0; fi31 < 7; fi31++ {
            if v16 {
                v32 := GenS2{F0: (v17 & (^(fi31))), F1: 0}
                for fi33 := 0; fi33 < 16; fi33++ {
                    v30 = v30
                    var v34 uint32 = v30
                    v16 = v16
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
            _ = fi31
        }
        _ = v30
    } else {
        switch (^(v17)) {
        case -159:
            fmt.Println(v16, v1, v17)
        case 127:
            fmt.Println(v2, v16)
        case 1:
            v16 = v16
        default:
            switch -402 {
            case 128:
                switch "hello" {
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v5 -= 160
                        if (0 < 721) {
                            for fi36 := 0; fi36 < 3; fi36++ {
                                if (182 != 255) {
                                    switch "hello" {
                                    case "a":
                                        v2 = (^(v5))
                                    }
                                    v16 = false
                                } else {
                                    v37 := GenS3{F0: 1, F1: -2147483648}
                                    _ = v37
                                }
                                var v38 float32 = (0.1 + (-2.25 * 0.0))
                                var v39 int32 = -1
                                _ = fi36
                                _ = v38
                                _ = v39
                            }
                            v5 -= (^(v2))
                            for fi40 := 0; fi40 < 7; fi40++ {
                                v2 -= v5
                                if (!v16) {
                                    var v41 int16 = -1
                                    _ = v41
                                }
                                fmt.Println(v16)
                                var v42 int32 = 86
                                var v43 string = "foo bar"
                                _ = fi40
                                _ = v42
                                _ = v43
                            }
                        }
                        var v44 string = ("foo bar" + "")
                        _ = fi35
                        _ = v44
                    }
                }
            case 2147483647:
                v45 := GenS3{F0: 924, F1: -1}
                _ = v45
            default:
                var v46 bool = v16
                _ = v46
            }
        }
    }
    var v47 float64 = ((((0.1 / 3.14159) + (1e10 + float64((1e10 + (((1.5 + 1.5) + -2.25) + -2.25))))) / -2.25) + -2.25)
    var v48 int16 = 1
    var v49 int16 = (v48 - v48)
    v47 += (v47 - v47)
    fmt.Println(v48, v17, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
