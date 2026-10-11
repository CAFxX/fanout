package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    v1 := []int{128}
    var v2 uint = 859
    var v3 float32 = float32(0.1)
    var v4 uint64 = 42
    var v5 uint32 = 357
    var v6 string = "\n\t"
    _ = len(v1)
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v6, v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (fi7 | -2147483648) {
        case 560:
            for fi8 := 0; fi8 < 5; fi8++ {
                v5 = v5
                _ = len(v1)
                var v9 int8 = 22
                _ = fi8
                _ = v9
            }
        default:
            v10 := GenS1{F0: (v6 + v6), F1: v5}
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v3, v5)
    if (1 <= -2147483648) {
        fmt.Println(v4, v4, v2)
        switch (v6 + v6) {
        case "foo bar":
            if ((!(!(false && true))) || false) {
                switch (v6 + "a") {
                case "a":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := GenS2{F0: -60}
                        v1 = append(v1, (fi11 * (fi11 % ((fi11) | 1))))
                        v13 := GenS1{F0: ("hello" + v6), F1: (v5 + 4294967295)}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                case "foo bar":
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            v14 := GenS2{F0: 42}
                            v1 = append(v1, 4294967295)
                            _ = v14
                        }
                    } else {
                        v3 -= (1.5 * v3)
                        v2 = (v2 - v2)
                    }
                default:
                    v15 := GenS3{F0: v2, F1: 0}
                    _ = v15
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: v6, F1: (v5 << ((v2) & 63))}
                    v3 = v3
                    _ = fi16
                    _ = v17
                }
            }
        case "hello":
            if (!true) {
                fmt.Println(v3, v2)
                v18 := GenS4{F0: (v5 * (v5 % (((v5 | 1)) | 1)))}
                _ = v18
            }
        case "a":
            var v19 bool = false
            _ = v19
        default:
            var v20 uint16 = uint16(v4)
            _ = v20
        }
    } else {
        if (42 < 2147483647) {
            v5 += (v5 << (((v2 + v2)) & 63))
            _ = len(v1)
            var v21 int64 = 107
            _ = v21
        }
    }
    if (9223372036854775807 > 42) {
        var v22 float32 = 3.14159
        if len(v1) > 0 {
            v1[0] = -335
        }
        var v23 uint32 = v5
        _ = v22
        _ = v23
    } else {
        v4 -= (v4 * (v4 + (v4 & (v4 % ((v4) | 1)))))
        _ = len(v1)
    }
    fmt.Println(v2)
    for fi24 := 0; fi24 < 15; fi24++ {
        v5 -= (v5 | v5)
        var v25 string = v6
        _ = fi24
        _ = v25
    }
    fmt.Println(v3, v3, v6)
    _ = len(v1)
    if true {
        for fi26 := 0; fi26 < 15; fi26++ {
            v1 = append(v1, fi26)
            _ = fi26
        }
        v6 = (v6 + "foo bar")
    }
    if ((3.14159 * (v3 - (float32((float64(-2.25) - (1.5 * (1e10 / 1.0)))) * ((1e10 / v3) * v3)))) <= 3.14159) {
        v27 := GenS5{F0: (v2 % (((v2 / ((v2) | 1))) | 1)), F1: (!(65535 < 255))}
        fmt.Println(v4, v5, v5)
        v1 = append(v1, 2147483647)
        _ = v27
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v6 = v6
    var v28 int64 = 0
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
