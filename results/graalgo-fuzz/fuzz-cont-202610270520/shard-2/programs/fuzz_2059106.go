package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -1
    v3 := []int{-99, -967}
    v4 := make(map[string]uint32)
    var v5 int8 = 0
    v5 -= -105
    var v6 int = 32768
    v5 = v5
    delete(v1, 1)
    if (!(!(113 == 836))) {
        var v7 float64 = -2.25
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(fi8)
            v9 := GenS1{F0: (v5 % (((v5 ^ v5)) | 1)), F1: v2, F2: 0}
            _ = fi8
            _ = v9
        }
        var v10 float64 = (((v7 / v7) + v7) * (v7 / (((0.1 - (1.5 / ((((v7 / v7) * v7) + (((v7 - float64(-2.25)) * v7) + float64((1.5 - (1.5 - ((float32(v7) / 1.0) - 0.0)))))) + v7))) * (float64(float32(float64((1e10 + -2.25)))) / v7)) / float64((3.14159 * 0.1)))))
        _ = v10
        _ = v7
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            v6 += v6
            _ = fi11
        }
    }
    if true {
        v5 -= v5
        switch 2147483647 {
        case 42:
            var v12 int32 = 2147483647
            _ = v12
        case 1:
            if false {
                for fi13 := 0; fi13 < 12; fi13++ {
                    if len(v3) > 0 {
                        v3[0] = (fi13 & fi13)
                    }
                    _ = fi13
                }
                v5 += (v5 >> ((2) & 63))
                v14 := v4["foo bar"]
                _ = v14
            }
        case -9223372036854775808:
            var v15 uint = 65535
            _ = v15
        default:
            switch ((("hello" + "hello") + "") + "\n\t") {
            case "foo bar":
                v2 = -9223372036854775808
            case "hello":
                if (0.1 == 1e10) {
                    if false {
                        var v16 float64 = -2.25
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v2 += v2
                            switch (v6 & fi17) {
                            case 1:
                                if true {
                                    fmt.Println(fi17, v2)
                                }
                            case 569:
                                var v18 float32 = ((1.5 * (1e10 - 3.14159)) - float32((((1e10 / 3.14159) - 0.0) / (3.14159 + (3.14159 + 1.5)))))
                                _ = v18
                            case -1:
                                for fi19 := 0; fi19 < 8; fi19++ {
                                    var v20 uint64 = 1
                                    fmt.Println(v2, fi19)
                                    _ = fi19
                                    _ = v20
                                }
                            default:
                                var v21 float32 = 1e10
                                _ = v21
                            }
                            _ = fi17
                        }
                        v22 := v1[443]
                        _ = v22
                    }
                    delete(v4, "\n\t")
                    var v23 int16 = 42
                    _ = v23
                }
            case "":
                var v24 uint16 = 778
                _ = v24
            default:
                for fi25 := 0; fi25 < 17; fi25++ {
                    fmt.Println(v2, v5, v2)
                    var v26 bool = (!false)
                    fmt.Println(v2, v5)
                    _ = fi25
                    _ = v26
                }
            }
        }
        v2 = v2
    }
    var v27 uint = 93
    if true {
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(v2)
            _ = fi28
        }
        v29 := v4["foo bar"]
        _ = v29
    } else {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v30 := GenS1{F0: (v5 % (((v5 / (((v5 >> ((v27) & 63))) | 1))) | 1)), F1: v2, F2: v27}
            _ = v30
        case "\n\t":
            for fi31 := 0; fi31 < 1; fi31++ {
                if false {
                    v6 = fi31
                    switch ("a" + ("" + "")) {
                    case "a":
                        if true {
                            v5 += v5
                        }
                    case "foo bar":
                        var v32 uint64 = 2
                        _ = v32
                    case "\n\t":
                        var v33 uint32 = 1
                        _ = v33
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v2 -= v2
                }
                v34 := GenS1{F0: 0, F1: v2, F2: (v27 - v27)}
                _ = fi31
                _ = v34
            }
        case "":
            v35 := GenS1{F0: 119, F1: 0, F2: v27}
            _ = v35
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 bool = (559 > -2147483648)
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    if (!((1 < -2147483648) || false)) {
        v39 := GenS1{F0: (v5 & 1), F1: (^(v2)), F2: (v27 - v27)}
        _ = v39
    }
    v5 = -26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
