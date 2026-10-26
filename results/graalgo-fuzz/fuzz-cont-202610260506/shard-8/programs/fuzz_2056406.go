package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }

func main() {
    var v1 bool = (false || false)
    var v2 int = 0
    var v3 int32 = 0
    v4 := make(map[int]string)
    var v5 int64 = 42
    switch "" {
    case "\n\t":
        if (1 < -9223372036854775808) {
            var v6 int8 = -68
            _ = v6
        }
    case "a":
        v5 = v5
    default:
        if false {
            var v7 int8 = -113
            fmt.Println(v3)
            _ = v7
        } else {
            v8 := GenS1{F0: ((-842 < 32767) || (v3 < (v3 % ((v3) | 1)))), F1: -591, F2: -76}
            var v9 int64 = -9223372036854775808
            _ = v8
            _ = v9
        }
    }
    if (0 == 456) {
        v1 = (1e10 <= 1.5)
        _ = len(v4)
    } else {
        if ((1.5 - float32((-2.25 * 0.0))) < 1e10) {
            if (v1 || (v5 > v5)) {
                var v10 uint16 = 512
                _ = v10
            }
            v11 := []int{(v2 + (v2 | (v2 + (v2 + v2)))), v2, v2}
            switch (v2 + (-(v2))) {
            case 2147483647:
                v12 := GenS1{F0: v1, F1: 1000, F2: -1}
                _ = v12
            default:
                v13 := GenS1{F0: (!true), F1: 1000, F2: -551}
                _ = v13
            }
            _ = v11
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 bool = ((fi14 - fi14) != v2)
                    var v17 uint8 = 1
                    v18 := v4[127]
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := []int{fi19, (fi19 >> ((255) & 63)), 128}
                var v21 int16 = 700
                fmt.Println(v1, v2, v3)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    if (1e10 <= 1.5) {
        var v22 float32 = -2.25
        _ = v22
    } else {
        v23 := GenS1{F0: (124 <= 1), F1: 178, F2: -32768}
        if (!(2 <= 0)) {
            fmt.Println(v3, v1)
            var v24 int64 = -9223372036854775808
            _ = v24
        } else {
            fmt.Println(v2, v1)
            v2 -= (v2 * int(v3))
            for fi25 := 0; fi25 < 17; fi25++ {
                v1 = (201 >= 1)
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 int64 = v5
                    _ = fi26
                    _ = v27
                }
                v28 := []int{(v2 / (((fi25 % ((-815) | 1))) | 1))}
                _ = fi25
                _ = v28
            }
        }
        _ = v23
    }
    if true {
        switch 4294967295 {
        case 765:
            fmt.Println(v1)
        case 293:
            fmt.Println(v5, v1, v1)
        default:
            var v29 float32 = ((((0.0 - ((0.1 + 3.14159) / 0.1)) - 0.1) - 0.0) / float32(1.5))
            _ = v29
        }
    }
    v1 = v1
    v1 = true
    for fi30 := 0; fi30 < 19; fi30++ {
        v31 := []int{(v2 * int(v5)), 127}
        if v1 {
            v5 -= int64(v3)
            var v32 int = v2
            fmt.Println(v32)
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    var v33 int8 = 100
    fmt.Println(v3, v1)
    var v34 float32 = (-2.25 * 0.0)
    var v35 bool = v1
    v36 := GenS1{F0: true, F1: 1000, F2: 1}
    for fi37 := 0; fi37 < 8; fi37++ {
        if (18446744073709551615 != 1) {
            if (v1 && v35) {
                for fi38 := 0; fi38 < 11; fi38++ {
                    v33 = v33
                    fmt.Println(v1, v34)
                    _ = fi38
                }
            } else {
                var v39 float64 = 1e10
                fmt.Println(v2)
                delete(v4, (fi37 ^ -9223372036854775808))
                _ = v39
            }
            v4[32768] = "\n\t"
        } else {
            var v40 int16 = -1
            fmt.Println(v40, v34)
            _ = v40
        }
        _ = fi37
    }
    fmt.Println(v5)
    var v41 float32 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
