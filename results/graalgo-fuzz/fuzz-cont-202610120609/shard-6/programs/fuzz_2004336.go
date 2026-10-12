package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 float64 = (0.1 - ((1.5 * -2.25) / 1e10))
    var v4 float64 = -2.25
    var v5 uint = 2
    var v6 int16 = 135
    if (4294967295 <= 2147483647) {
        switch -2147483648 {
        case 585:
            if (0 > (-(v6))) {
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 int64 = -9223372036854775808
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 string = ""
                        if (!(42 != 8)) {
                            v11 := GenS1{F0: 1, F1: (fi9 + (fi9 / ((9223372036854775807) | 1))), F2: 65535}
                            if true {
                                v12 := []int{fi7, fi9, fi9}
                                _ = v12
                            }
                            _ = v11
                        }
                        _ = fi9
                        _ = v10
                    }
                    if true {
                        var v13 bool = true
                        _ = v13
                    }
                    _ = fi7
                    _ = v8
                }
                var v14 float32 = float32(((((v4 / v3) / 1.0) - float64(((1.5 / 1e10) / 0.1))) / v4))
                _ = v14
            } else {
                v4 = 1e10
                switch (("hello" + ((("hello" + "foo bar") + ("hello" + "hello")) + "foo bar")) + ("foo bar" + (("a" + "hello") + "foo bar"))) {
                case "\n\t":
                    if (!(42 == -128)) {
                        var v15 uint32 = 470
                        _ = v15
                    }
                }
            }
        case 42:
            var v16 uint = 1
            _ = v16
        case 1:
            v4 -= v3
        default:
            var v17 int32 = 0
            _ = v17
        }
        v1[-1] = false
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint32 = 4294967295
            var v20 uint64 = 0
            v4 = (1e10 * (float64((1e10 * (0.1 * ((1e10 - 3.14159) + (3.14159 * (float32(float64(-2.25)) * 3.14159)))))) - float64(0.0)))
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        v21 := GenS1{F0: 352, F1: -785, F2: 65535}
        _ = v21
    }
    var v22 uint64 = 990
    var v23 int32 = 293
    var v24 string = "a"
    if ((v24 < v24) || ((v2 << (((^(v5))) & 63)) >= v2)) {
        if (-725 != -1) {
            for fi25 := 0; fi25 < 13; fi25++ {
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 int16 = v6
                    _ = fi26
                    _ = v27
                }
                v1[(fi25 - (fi25 ^ -463))] = false
                _ = fi25
            }
            v6 += (v6 | v6)
        }
        v2 += -9223372036854775808
    }
    v28 := GenS1{F0: 42, F1: -2147483648, F2: 738}
    var v29 float64 = float64((3.14159 / (3.14159 / (1e10 / 1e10))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
