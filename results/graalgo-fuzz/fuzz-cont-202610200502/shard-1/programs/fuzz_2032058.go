package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 bool) uint {
    var v4 int8 = 127
    var v5 bool = (-22 == 9223372036854775807)
    v5 = v5
    var v6 int64 = 9223372036854775807
    v4 += (v4 * (^(v4)))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1() bool {
    var v7 int32 = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float64 = (3.14159 / 0.1)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 329, F2: -1}
    fmt.Println(v7)
    _ = v10
    _ = v7
    return (-670 >= -32768)
}

func main() {
    var v1 int = 32768
    var v2 float64 = float64(((0.0 + ((float32((float64(0.1) - 1e10)) - 3.14159) + -2.25)) - -2.25))
    var v3 uint64 = 0
    v11 := GenS1{F0: (v3 << ((59) & 63)), F1: 904, F2: 1}
    v2 += 0.0
    var v12 float64 = v2
    var v13 string = "hello"
    v13 = ((v13 + "a") + v13)
    v14 := []int{v1, v1, (-(v1))}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int = -580
        switch 830 {
        case 127:
            var v17 string = ("foo bar" + "\n\t")
            _ = v17
        case 4294967295:
            if (1.5 != v2) {
                v18 := GenS1{F0: v3, F1: 145, F2: 740}
                _ = v18
            } else {
                var v19 uint32 = 2147483648
                switch v1 {
                case 1:
                    for fi20 := 0; fi20 < 5; fi20++ {
                        switch (v1 & fi15) {
                        case 0:
                            switch v16 {
                            case 4294967295:
                                var v21 bool = ((v13 < v13) || true)
                                _ = v21
                            case 42:
                                if (v19 == (v19 % (((^(v19))) | 1))) {
                                    if (2147483647 <= v19) {
                                        var v22 int32 = -2147483648
                                        v3 = 42
                                        fmt.Println(v13)
                                        _ = v22
                                    }
                                }
                            default:
                                var v23 int = 2147483647
                                _ = v23
                            }
                        }
                        var v24 int = v1
                        _ = fi20
                        _ = v24
                    }
                case 738:
                    v12 = (float64(-2.25) + v12)
                }
                for fi25 := 0; fi25 < 20; fi25++ {
                    v2 = (v12 - 3.14159)
                    _ = fi25
                }
                _ = v19
            }
        }
        v26 := GenS1{F0: (v3 & (^(v3))), F1: 1, F2: -162}
        _ = fi15
        _ = v16
        _ = v26
    }
    fmt.Println(v1, v1, v12)
    v27 := GenS2{F0: v3, F1: 32767}
    for fi28 := 0; fi28 < 19; fi28++ {
        if (!(v13 <= "\n\t")) {
            v12 = v12
        }
        v29 := GenS3{F0: true, F1: 255}
        if false {
            v13 = (v13 + v13)
            var v30 uint16 = 32767
            if (198 > 2) {
                if ("foo bar" <= ("foo bar" + (((((v13 + "a") + (v13 + "foo bar")) + "a") + ((("" + "foo bar") + "a") + v13)) + "hello"))) {
                    v31 := GenS1{F0: v3, F1: 2147483648, F2: -9223372036854775808}
                    _ = v31
                } else {
                    var v32 float64 = float64(0.0)
                    _ = v32
                }
                v12 -= -2.25
            }
            _ = v30
        } else {
            for fi33 := 0; fi33 < 14; fi33++ {
                v1 = (fi33 * 209)
                if (92 > 103) {
                    v34 := GenS2{F0: v3, F1: 889}
                    var v35 int64 = 1
                    var v36 float32 = (3.14159 / ((0.0 * 0.1) - 0.1))
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v12 -= v2
                _ = fi33
            }
        }
        _ = fi28
        _ = v29
    }
    if false {
        switch v13 {
        case "a":
            fmt.Println(v3, v13)
        default:
            for fi37 := 0; fi37 < 15; fi37++ {
                v13 = "hello"
                _ = fi37
            }
        }
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            v3 += v3
            _ = fi38
        }
        for fi39 := 0; fi39 < 18; fi39++ {
            v2 = v2
            var v40 float64 = v2
            v41 := GenS4{F0: (v13 + "\n\t")}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    if ((!(!false)) && ((int16(v1) < 68) && (0 >= -14))) {
        if (2 > 328) {
            _ = len(v14)
            var v42 uint64 = (v3 - (v3 - v3))
            _ = v42
        }
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println(v13, v12)
    v12 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
}
