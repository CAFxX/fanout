package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8 }
func genFn0() int64 {
    if true {
        if true {
            var v7 int8 = -128
            for fi8 := 0; fi8 < 11; fi8++ {
                v7 += 33
                v9 := GenS1{F0: 269, F1: 127}
                for fi10 := 0; fi10 < 12; fi10++ {
                    if false {
                        v11 := GenS1{F0: 258, F1: 1}
                        _ = v11
                    }
                    v7 = 1
                    v12 := GenS2{F0: 1, F1: 562}
                    _ = fi10
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            v13 := GenS2{F0: -1, F1: 0}
            _ = v13
            _ = v7
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS3{F0: -19, F1: 82}
                var v16 uint16 = 121
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        var v17 uint16 = 0
        switch 32768 {
        case 9223372036854775807:
            v18 := GenS4{F0: 521}
            _ = v18
        case 32768:
            v19 := GenS3{F0: -1, F1: 1}
            _ = v19
        case 32767:
            var v20 string = ""
            _ = v20
        }
        _ = v17
    }
    var v21 int16 = 682
    _ = v21
    return int64(v21)
}

func main() {
    v1 := []string{}
    var v2 int = 1
    var v3 float32 = 3.14159
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    var v6 uint = 146
    v22 := GenS2{F0: -1, F1: 2}
    for fi23 := 0; fi23 < 20; fi23++ {
        if ((!(!(-1 != 1))) && v5) {
            var v24 int8 = 109
            v2 = fi23
            v25 := GenS4{F0: v6}
            _ = v24
            _ = v25
        }
        _ = fi23
    }
    switch "" {
    case "":
        v6 -= v6
    case "foo bar":
        v26 := GenS1{F0: 338, F1: 1}
        _ = v26
    case "\n\t":
        delete(v4, 2)
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint64 = 0
        var v29 float64 = 1.5
        _ = len(v4)
        _ = fi27
        _ = v28
        _ = v29
    }
    v2 = 9223372036854775807
    switch ("foo bar" + ("" + "\n\t")) {
    case "\n\t":
        for fi30 := 0; fi30 < 13; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                switch fi31 {
                case 9223372036854775807:
                    v3 -= 0.1
                case 0:
                    if true {
                        if v5 {
                            switch (v2 / ((42) | 1)) {
                            case 0:
                                var v32 int64 = 42
                                _ = v32
                            case 32767:
                                _ = len(v4)
                            case 42:
                                fmt.Println(v5)
                            default:
                                v33 := GenS2{F0: genFn0(), F1: (v6 % ((279) | 1))}
                                _ = v33
                            }
                        } else {
                            var v34 bool = ((("a" + ("\n\t" + "foo bar")) != ("" + ("hello" + "a"))) && v5)
                            switch (v2 ^ fi30) {
                            case 32768:
                                var v35 uint64 = 1
                                _ = v35
                            default:
                                fmt.Println(v34)
                            }
                            var v36 int16 = 817
                            _ = v34
                            _ = v36
                        }
                        var v37 uint64 = 42
                        v38 := GenS5{F0: 2}
                        _ = v37
                        _ = v38
                    }
                case 32768:
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                if (313 < 626) {
                    if v5 {
                        v3 -= (float32(1.5) - (0.1 - (3.14159 + 3.14159)))
                        if (v5 && v5) {
                            v39 := GenS2{F0: -9223372036854775808, F1: (v6 & (v6 & (v6 ^ (v6 / ((18446744073709551615) | 1)))))}
                            _ = v39
                        }
                        fmt.Println(fi31)
                    }
                    if (!true) {
                        var v40 int8 = 0
                        v40 = v40
                        fmt.Println(v40)
                        _ = v40
                    } else {
                        v6 = v6
                        _ = len(v1)
                        switch (("\n\t" + "foo bar") + "a") {
                        case "a":
                            v6 += (v6 >> ((v6) & 63))
                        default:
                            v41 := GenS6{F0: 970, F1: -128}
                            _ = v41
                        }
                    }
                    var v42 int64 = -9223372036854775808
                    _ = v42
                } else {
                    var v43 float32 = (float32((1e10 / 1.0)) / v3)
                    _ = v43
                }
                var v44 int16 = 1
                _ = fi31
                _ = v44
            }
            if ((float32(1e10) + ((v3 + 1e10) + ((((v3 + (3.14159 * ((v3 - 3.14159) + v3))) * v3) - v3) / 0.1))) > (((((v3 / (float32((3.14159 / (0.1 / 1e10))) * v3)) - (0.1 / (v3 / v3))) * (0.1 / (v3 / v3))) + v3) - v3)) {
                v5 = ("" > ("a" + ("\n\t" + "")))
            } else {
                v5 = v5
                var v45 uint32 = 0
                _ = v45
            }
            _ = fi30
        }
    case "hello":
        var v46 int8 = 46
        _ = v46
    }
    v4[2147483648] = 0
    fmt.Println(v5, v5, v5)
    v3 = 3.14159
    var v47 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
