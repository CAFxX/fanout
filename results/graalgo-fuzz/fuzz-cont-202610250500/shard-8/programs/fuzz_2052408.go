package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint64; F1 bool; F2 int }
func genFn0(p0 uint64, p1 uint64, p2 uint32) int64 {
    v6 := []int{4294967295, 32768, 42}
    p0 = (^(p1))
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p0, fi7, fi7)
        var v8 uint = 1
        var v9 int64 = 9223372036854775807
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = 1000
    var v11 int16 = v10
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v6
    return 1
}

func main() {
    var v1 float64 = (1.5 * 3.14159)
    var v2 uint8 = 228
    var v3 uint8 = (v2 * (v2 >> ((963) & 63)))
    v4 := make(map[string]string)
    var v5 float32 = 1e10
    v1 -= 1e10
    if (343 >= -1) {
        v12 := GenS1{F0: 2}
        fmt.Println(v1)
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint32 = 523
                var v16 int8 = -1
                if (v2 == (v2 * (v2 >> ((65) & 63)))) {
                    v1 -= ((3.14159 * float64(1.5)) / v1)
                    v17 := GenS2{F0: (^(v16)), F1: v2, F2: v16}
                    v1 = -2.25
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            v4[("\n\t" + "a")] = ((("foo bar" + "foo bar") + (("\n\t" + "") + "\n\t")) + (("\n\t" + "foo bar") + (("hello" + "\n\t") + "hello")))
            _ = fi13
        }
        _ = v12
    } else {
        v18 := GenS3{F0: -34}
        _ = v18
    }
    v1 -= v1
    v2 -= v3
    if (!((65535 > 141) && (913 >= 749))) {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := []int{(-(fi19)), fi19, (fi19 ^ (fi19 + fi19))}
            var v21 uint8 = (v3 | v2)
            _ = fi19
            _ = v20
            _ = v21
        }
        switch 128 {
        case 193:
            v22 := GenS1{F0: 1}
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                if true {
                    var v25 uint8 = 127
                    _ = v25
                }
                _ = fi24
            }
            var v26 int16 = -702
            _ = fi23
            _ = v26
        }
    }
    if ((v2 / (((^(v3))) | 1)) >= v3) {
        if false {
            switch "" {
            case "":
                if (0 > genFn0(808, 4294967295, 42)) {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := GenS4{F0: 922, F1: 0, F2: 65535}
                        if true {
                            for fi29 := 0; fi29 < 13; fi29++ {
                                fmt.Println(fi27, v5)
                                _ = fi29
                            }
                            v5 = (v5 / v5)
                            fmt.Println(v5)
                        }
                        if true {
                            var v30 uint16 = 1
                            v30 += v30
                            _ = v30
                        } else {
                            var v31 string = "hello"
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(v5, v5)
                    v32 := GenS5{F0: 1, F1: 42, F2: -9223372036854775808}
                    _ = v32
                } else {
                    var v33 int16 = 0
                    _ = v33
                }
            case "a":
                if (0 <= -415) {
                    v5 = ((0.0 * v5) + (v5 * (float32(((v1 + -2.25) / v1)) / v5)))
                } else {
                    var v34 float32 = ((((v5 / v5) - v5) * v5) - v5)
                    v35 := GenS5{F0: 651, F1: 0, F2: genFn0(770, 0, 4294967295)}
                    fmt.Println(v3, v1, v2)
                    _ = v34
                    _ = v35
                }
            default:
                v36 := GenS6{F0: 367, F1: false, F2: 128}
                _ = v36
            }
        }
    } else {
        v2 += v3
        v37 := GenS3{F0: 0}
        _ = v37
    }
    fmt.Println(v2, v5, v3)
    _ = len(v4)
    v5 += v5
    var v38 uint8 = v2
    var v39 uint8 = 255
    if ((1 > 780) && (0 > -409)) {
        switch -15 {
        case 4294967295:
            switch ("\n\t" + "foo bar") {
            case "a":
                switch ("" + "hello") {
                case "\n\t":
                    for fi40 := 0; fi40 < 12; fi40++ {
                        v38 = (v38 - 42)
                        _ = fi40
                    }
                case "hello":
                    if false {
                        v1 -= (3.14159 * ((v1 - 3.14159) + (v1 - 1.5)))
                    }
                }
            case "hello":
                var v41 float32 = float32(v1)
                _ = v41
            case "foo bar":
                v42 := []int{}
                _ = v42
            }
        }
    } else {
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 int32 = 42
            for fi45 := 0; fi45 < 12; fi45++ {
                v46 := []int{690, 2147483647}
                _ = fi45
                _ = v46
            }
            switch "\n\t" {
            case "foo bar":
                v2 -= (^(v3))
            case "a":
                var v47 bool = (!(-128 >= -62))
                _ = v47
            case "hello":
                v44 += 42
            default:
                v48 := GenS1{F0: 202}
                _ = v48
            }
            _ = fi43
            _ = v44
        }
        if true {
            v5 = float32((v1 * 1e10))
        } else {
            fmt.Println(v2, v39)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
