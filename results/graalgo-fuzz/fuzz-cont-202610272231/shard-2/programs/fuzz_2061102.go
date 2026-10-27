package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    var v3 int16 = 42
    var v4 int = -244
    var v5 int64 = 199
    v6 := []uint32{1, 821, 1, 443}
    if (uint16(v2) <= 663) {
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(fi7)
            switch (-(fi7)) {
            case 32767:
                for fi8 := 0; fi8 < 10; fi8++ {
                    v4 += fi8
                    if len(v6) > 0 {
                        v6[0] = 896
                    }
                    var v9 uint16 = 1
                    _ = fi8
                    _ = v9
                }
            }
            v2 = v2
            _ = fi7
        }
    } else {
        var v10 bool = (-1 != -115)
        _ = v10
    }
    v5 -= 92
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v5 += (v5 / ((9223372036854775807) | 1))
            var v13 float64 = 1.5
            _ = fi12
            _ = v13
        }
        var v14 float32 = -2.25
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 float64 = (float64((1e10 * (v14 + (float32((3.14159 + 1.5)) + v14)))) - ((-2.25 / -2.25) * ((((((0.1 * 0.0) / 3.14159) / ((0.1 + 0.0) * (0.1 + (0.1 - 0.1)))) / 1.5) / -2.25) + (0.0 * (float64(1e10) + 0.1)))))
            _ = fi15
            _ = v16
        }
        _ = fi11
        _ = v14
    }
    var v17 uint64 = 42
    _ = len(v6)
    if (!(0 == 2147483648)) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            switch (("hello" + ("" + "a")) + ("" + "foo bar")) {
            case "":
                if (!(127 == 1)) {
                    switch fi18 {
                    case 681:
                        var v19 uint16 = 1
                        _ = v19
                    case 1:
                        v3 = (^(v3))
                    }
                    v20 := GenS1{F0: 1, F1: 42, F2: 42}
                    var v21 int = (v4 & (v4 + fi18))
                    _ = v20
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        _ = len(v6)
                        v17 -= v17
                        var v23 uint32 = 2
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v5, v17)
                }
            case "hello":
                fmt.Println(v4, v3, fi18)
            default:
                if false {
                    v4 += 2147483647
                }
            }
            delete(v1, "\n\t")
            _ = fi18
        }
    }
    var v24 string = "a"
    _ = len(v6)
    var v25 uint = 65535
    v25 -= v25
    var v26 int16 = (v3 << ((v25) & 63))
    var v27 int8 = 100
    if (v5 > (v5 % (((v5 % ((-9223372036854775808) | 1))) | 1))) {
        var v28 uint16 = 32767
        _ = len(v6)
        _ = v28
    }
    v1["\n\t"] = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
