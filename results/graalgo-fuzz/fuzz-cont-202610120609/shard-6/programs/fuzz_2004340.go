package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
func genFn0(p0 int16) bool {
    if false {
        v7 := []int{-2147483648, 0, 32767, 32767}
        p0 += 1
        _ = v7
    }
    v8 := []int{32767, 9223372036854775807}
    if len(v8) > 0 {
        v8[0] = 32767
    }
    _ = p0
    _ = v8
    return (32767 >= 42)
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"foo bar"}
    var v3 float32 = 1e10
    v4 := make(map[string]bool)
    var v5 uint8 = 94
    var v6 int = 471
    if (65535 <= 0) {
        v9 := GenS1{F0: 1, F1: v6}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 bool = false
                var v13 int8 = -1
                _ = fi11
                _ = v12
                _ = v13
            }
            v5 += 74
            v14 := GenS1{F0: 0, F1: fi10}
            _ = fi10
            _ = v14
        }
        if false {
            _ = len(v2)
            var v15 uint16 = 2
            _ = v15
        } else {
            switch v6 {
            case -1:
                fmt.Println(v3, v5)
            default:
                fmt.Println(v6, v3, v5)
            }
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 bool = false
            for fi18 := 0; fi18 < 2; fi18++ {
                v3 = 1.5
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            if ((0.1 / float32((1.5 + (0.0 / 1.0)))) < (((v3 - float32((float64(v3) / -2.25))) * (v3 + float32(((0.0 * (3.14159 - (float64(float32(float64(v3))) * 1.5))) - ((0.1 + float64(1.5)) - (0.1 / 1e10)))))) * v3)) {
                for fi20 := 0; fi20 < 8; fi20++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v6 = (v6 - 494)
                    var v21 int32 = -1
                    _ = len(v1)
                    _ = fi20
                    _ = v21
                }
            }
            var v22 string = ("a" + "")
            if ((int16(fi19) <= -1) && true) {
                v23 := GenS1{F0: 42, F1: v6}
                _ = v23
            } else {
                _ = len(v2)
            }
            _ = v22
        }
        fmt.Println(v3)
        _ = fi19
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int64 = -195
        _ = fi24
        _ = v25
    }
    var v26 string = ("\n\t" + ("\n\t" + ""))
    switch v26 {
    case "foo bar":
        var v27 uint64 = 42
        _ = v27
    case "":
        var v28 string = (((v26 + v26) + v26) + v26)
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 18; fi29++ {
            v5 += v5
            v26 = ((((("" + (v26 + ("foo bar" + v26))) + (v26 + v26)) + v26) + v26) + v26)
            if (!false) {
                var v30 int32 = 42
                var v31 int32 = (v30 ^ v30)
                fmt.Println(fi29, v5, v31)
                _ = v30
                _ = v31
            } else {
                fmt.Println(v3, v26)
            }
            _ = fi29
        }
    default:
        if (!(false || ((!(!(255 < 129))) || (3.14159 == 1.5)))) {
            if true {
                var v32 uint64 = 132
                _ = v32
            }
        } else {
            switch (^(v6)) {
            case -2147483648:
                v33 := GenS1{F0: 9223372036854775807, F1: v6}
                _ = v33
            case 32768:
                v34 := GenS1{F0: 42, F1: v6}
                _ = v34
            case 4294967295:
                v3 = v3
            default:
                v2 = append(v2, "a")
            }
            switch "" {
            case "foo bar":
                fmt.Println(v6, v6)
            }
        }
    }
    v3 -= v3
    _ = len(v2)
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint16 = 32767
        _ = fi35
        _ = v36
    }
    v26 = v26
    if false {
        if true {
            v37 := GenS2{F0: (v5 >> ((65535) & 63)), F1: 1, F2: (v6 << ((499) & 63))}
            _ = v37
        }
        fmt.Println(v6, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
