package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0() int16 {
    v7 := GenS1{F0: 42, F1: false}
    if (42 > 1) {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{}
            _ = len(v9)
            switch "foo bar" {
            case "foo bar":
                v10 := GenS1{F0: 1000, F1: (!(92 >= 907))}
                _ = v10
            case "hello":
                switch fi8 {
                case 9223372036854775807:
                    var v11 uint16 = 2
                    _ = v11
                case 32768:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            default:
                v12 := GenS2{F0: 18446744073709551615, F1: -85, F2: 1}
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        var v13 int64 = 0
        _ = v13
    } else {
        if (18446744073709551615 < 42) {
            for fi14 := 0; fi14 < 1; fi14++ {
                switch (^(fi14)) {
                case 0:
                    var v15 string = ("foo bar" + ("\n\t" + "foo bar"))
                    _ = v15
                case 9223372036854775807:
                    var v16 int32 = -2147483648
                    _ = v16
                case 4294967295:
                    var v17 float64 = ((0.1 + 3.14159) + 0.0)
                    _ = v17
                default:
                    if (127 != 42) {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint = 2
                            fmt.Println(fi14, fi14, v19)
                            v19 -= 65535
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        v20 := []int{128, (fi14 % (((-(fi14))) | 1)), 2147483647, fi14}
                        _ = v20
                    }
                }
                fmt.Println(fi14)
                var v21 uint64 = 0
                _ = fi14
                _ = v21
            }
            switch ("a" + "\n\t") {
            case "a":
                switch 128 {
                case -1:
                    v22 := GenS1{F0: -1, F1: false}
                    _ = v22
                case 1:
                    var v23 int8 = 0
                    _ = v23
                case 194:
                    v24 := []int{-236}
                    _ = v24
                }
            case "hello":
                switch 182 {
                case 32768:
                    if true {
                        var v25 int16 = -430
                        _ = v25
                    } else {
                        var v26 int64 = -1
                        v26 = (v26 - (v26 % ((v26) | 1)))
                        _ = v26
                    }
                case 32767:
                    v27 := GenS1{F0: 32767, F1: (2 > 133)}
                    _ = v27
                case 128:
                    var v28 uint8 = 0
                    _ = v28
                }
            }
        }
        var v29 int64 = 9223372036854775807
        if true {
            if ((!(int8(v29) > 0)) && true) {
                if ((3.14159 / (0.1 * 1e10)) < ((1.5 - 1e10) + (0.0 - 0.1))) {
                    if false {
                        v30 := GenS2{F0: 591, F1: 1, F2: 80}
                        if false {
                            var v31 int16 = -863
                            _ = v31
                        }
                        v29 += v29
                        _ = v30
                    } else {
                        var v32 uint = 65535
                        v33 := []int{}
                        v29 -= v29
                        _ = v32
                        _ = v33
                    }
                    v29 = (v29 | v29)
                }
                v34 := []int{32767, 0, -9223372036854775808, 0}
                _ = v34
            }
            fmt.Println(v29)
            if (!true) {
                v29 = (^(v29))
                var v35 uint16 = 158
                var v36 bool = (208 > 32)
                _ = v35
                _ = v36
            } else {
                fmt.Println(v29)
            }
        }
        _ = v29
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 bool = (0.1 == (1e10 * -2.25))
        var v39 float32 = 3.14159
        var v40 uint32 = 287
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    _ = v7
    return 1000
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "a"
    var v3 int32 = 1
    v4 := make(map[string]uint32)
    var v5 int32 = (^(v3))
    v6 := []string{"hello", ((("\n\t" + v2) + ((v2 + v2) + (("a" + v2) + v2))) + "\n\t"), "a"}
    v3 = v3
    var v41 int32 = v5
    var v42 int32 = (-(v5))
    for fi43 := 0; fi43 < 4; fi43++ {
        switch fi43 {
        case 128:
            switch "\n\t" {
            case "a":
                var v44 int = fi43
                _ = v44
            case "":
                var v45 uint = 255
                _ = v45
            default:
                fmt.Println(v2)
            }
        }
        switch "" {
        case "hello":
            v1 = -2.25
        case "a":
            _ = len(v4)
        case "foo bar":
            var v46 float64 = 1.5
            _ = v46
        }
        _ = fi43
    }
    v41 -= (v41 - (v5 * v5))
    v5 = v41
    v1 -= v1
    v41 = v5
    if len(v6) > 0 {
        v6[0] = ""
    }
    var v47 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v47
    _ = v5
    _ = v6
}
