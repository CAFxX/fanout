package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint16) int64 {
    v4 := GenS1{F0: 2, F1: true, F2: p0}
    p0 -= p0
    for fi5 := 0; fi5 < 2; fi5++ {
        switch fi5 {
        case 338:
            for fi6 := 0; fi6 < 15; fi6++ {
                p0 = p0
                p0 += p0
                v7 := GenS1{F0: 316, F1: (("" > "a") || (42 > 442)), F2: p0}
                switch 0 {
                case 273:
                    switch (("hello" + "\n\t") + "\n\t") {
                    case "":
                        p0 += 42
                    }
                case 0:
                    v8 := GenS2{F0: (((^(fi5)) >= -9223372036854775808) && false), F1: 2}
                    _ = v8
                default:
                    var v9 uint64 = 129
                    _ = v9
                }
                fmt.Println(fi6)
                _ = fi6
                _ = v7
            }
        }
        p0 = p0
        _ = fi5
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    var v10 int16 = -662
    var v11 uint32 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        switch "" {
        case "":
            var v13 int16 = v10
            _ = v13
        }
        if len(v2) > 0 {
            v2[0] = v11
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v2 = append(v2, (^(v11)))
            _ = fi14
        }
        _ = fi12
    }
    if (718 > 42) {
        var v15 uint64 = 42
        v15 = v15
        v1 += float64((1e10 * (float32(v1) / (((1.5 - float32(-2.25)) - 0.0) - ((float32(((v1 / v1) / (v1 * float64(0.1)))) * (1e10 + 0.1)) + (0.1 - 3.14159))))))
        _ = v15
    } else {
        if (((v1 * v1) == (v1 * 3.14159)) || true) {
            if true {
                _ = len(v2)
                for fi16 := 0; fi16 < 12; fi16++ {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        switch "foo bar" {
                        case "foo bar":
                            v1 -= 3.14159
                        }
                        switch fi17 {
                        case 32768:
                            var v18 int = (fi16 >> ((1) & 63))
                            _ = v18
                        default:
                            v19 := GenS2{F0: (!((fi17 >> ((410) & 63)) > fi17)), F1: 0}
                            _ = v19
                        }
                        _ = fi17
                    }
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v11 -= v11
                        v21 := GenS2{F0: false, F1: 2}
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 bool = false
                        _ = fi22
                        _ = v23
                    }
                    _ = fi16
                }
            } else {
                v24 := GenS3{F0: 127}
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = fi25
        }
    }
    var v26 bool = true
    var v27 string = (((("" + "hello") + (("" + "hello") + ("\n\t" + "foo bar"))) + "\n\t") + "")
    for fi28 := 0; fi28 < 8; fi28++ {
        v10 = 148
        fmt.Println(v27, fi28, v10)
        fmt.Println(v11, v11, v10)
        _ = fi28
    }
    if true {
        v1 = v1
        if (42 < 2) {
            v1 = v1
            v2 = append(v2, v11)
        }
        v10 = -378
    } else {
        if (false || (42 == 255)) {
            delete(v3, v11)
            if len(v2) > 0 {
                v2[0] = v11
            }
            if (v1 != (0.1 * (v1 - ((float64((1.5 - 0.1)) + ((0.1 * 1e10) * v1)) / ((1.5 * v1) / v1))))) {
                for fi29 := 0; fi29 < 17; fi29++ {
                    if ((genFn0(0) < 419) || ((v11 / ((v11) | 1)) < v11)) {
                        if ((-2.25 - v1) == v1) {
                            v30 := GenS1{F0: 4294967295, F1: (154 > 4294967295), F2: 32767}
                            for fi31 := 0; fi31 < 9; fi31++ {
                                v1 = v1
                                _ = fi31
                            }
                            _ = v30
                        } else {
                            _ = len(v2)
                        }
                        v26 = v26
                        var v32 float64 = (1.5 - float64(3.14159))
                        _ = v32
                    }
                    _ = fi29
                }
                fmt.Println(v1)
            }
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            v26 = (374 < 0)
            _ = fi33
        }
        var v34 uint8 = 42
        _ = v34
    }
    var v35 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
}
