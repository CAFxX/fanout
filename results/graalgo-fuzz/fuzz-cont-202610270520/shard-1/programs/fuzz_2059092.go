package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{-2147483648, 9223372036854775807, -9223372036854775808}
    var v2 uint32 = 0
    var v3 string = "a"
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(v2, v3)
        if true {
            fmt.Println(v2)
            v2 += 114
            if true {
                v3 = v3
            }
        }
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        _ = fi4
    }
    var v5 int64 = 42
    if true {
        fmt.Println(v5, v5)
    }
    var v6 int8 = -1
    for fi7 := 0; fi7 < 12; fi7++ {
        switch "" {
        case "a":
            if false {
                v6 = (v6 / ((v6) | 1))
                switch (-(fi7)) {
                case 42:
                    if len(v1) > 0 {
                        v1[0] = fi7
                    }
                default:
                    var v8 string = v3
                    _ = v8
                }
            }
        case "\n\t":
            var v9 uint16 = 2
            _ = v9
        case "":
            v10 := GenS1{F0: v6, F1: -1, F2: v2}
            _ = v10
        }
        v3 = ""
        var v11 int32 = 0
        _ = fi7
        _ = v11
    }
    var v12 int16 = 1
    if (false || true) {
        v6 += (v6 & v6)
        var v13 int16 = v12
        _ = v13
    }
    if true {
        switch v3 {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = -295
            }
        case "a":
            for fi14 := 0; fi14 < 14; fi14++ {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v3 = "hello"
                    var v16 uint = 787
                    _ = fi15
                    _ = v16
                }
                fmt.Println(fi14, v2, v5)
                var v17 int16 = 1000
                _ = fi14
                _ = v17
            }
        case "\n\t":
            v12 -= (-(v12))
        }
        v1 = append(v1, 0)
    } else {
        switch 393 {
        case -343:
            var v18 uint32 = (v2 / ((2147483648) | 1))
            _ = v18
        default:
            for fi19 := 0; fi19 < 7; fi19++ {
                v2 += v2
                _ = fi19
            }
        }
        v6 = (v6 % (((^(v6))) | 1))
    }
    switch ((("foo bar" + (((("foo bar" + ("a" + "a")) + v3) + v3) + ("\n\t" + v3))) + v3) + v3) {
    case "a":
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint64 = 1
            _ = fi20
            _ = v21
        }
    case "hello":
        if (635 != 18446744073709551615) {
            if true {
                v22 := GenS2{F0: v2}
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi23 := 0; fi23 < 10; fi23++ {
                        v24 := GenS2{F0: 2147483647}
                        v25 := GenS2{F0: (^(v2))}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 float32 = 1.5
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        if (!false) {
                            v3 = "\n\t"
                        } else {
                            var v28 string = v3
                            var v29 int16 = -724
                            if (101 > 94) {
                                var v30 uint64 = 2
                                v31 := GenS2{F0: 701}
                                v6 -= v6
                                _ = v30
                                _ = v31
                            } else {
                                var v32 int64 = (v5 / (((v5 + v5)) | 1))
                                var v33 int32 = -2147483648
                                v1 = append(v1, (fi27 % (((fi27 << ((761) & 63))) | 1)))
                                _ = v32
                                _ = v33
                            }
                            _ = v28
                            _ = v29
                        }
                        v6 = v6
                        _ = fi27
                    }
                    if ((3.14159 <= 0.0) || ((!false) && (v3 <= "hello"))) {
                        fmt.Println(v6, v5)
                        switch 42 {
                        case 4294967295:
                            v3 = v3
                        }
                    }
                }
                _ = v22
            } else {
                var v34 float32 = ((((1.5 + ((0.1 - (0.1 / 3.14159)) / float32((0.1 / float64((1.5 - (((0.0 + (1e10 / float32(1.5))) + 3.14159) + 0.1))))))) - 3.14159) * 0.0) / 1.5)
                var v35 string = v3
                var v36 int = -2147483648
                _ = v34
                _ = v35
                _ = v36
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
