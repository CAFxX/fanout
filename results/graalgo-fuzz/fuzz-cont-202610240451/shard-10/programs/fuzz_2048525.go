package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = float64((((0.1 * (3.14159 - (3.14159 - 1e10))) / 0.1) * 1e10))
    var v2 int16 = 0
    var v3 uint = 4294967295
    var v4 int16 = -830
    v5 := []int{672, 2147483647}
    var v6 uint32 = 161
    v5 = append(v5, -883)
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (fi7 + (fi7 / (((fi7 << ((uint(v6)) & 63))) | 1))) {
        case -830:
            fmt.Println(v2)
        case 612:
            for fi8 := 0; fi8 < 8; fi8++ {
                _ = len(v5)
                _ = len(v5)
                _ = fi8
            }
        default:
            if false {
                _ = len(v5)
                switch "a" {
                case "":
                    var v9 uint16 = 0
                    _ = v9
                case "hello":
                    v10 := GenS1{F0: (v3 << (((v3 | (v3 * (v3 * v3)))) & 63))}
                    _ = v10
                default:
                    v4 += v4
                }
                var v11 int64 = 42
                _ = v11
            }
        }
        _ = fi7
    }
    switch "" {
    case "\n\t":
        var v12 int8 = 42
        _ = v12
    case "foo bar":
        v2 = 42
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v13 int16 = (-(v4))
    fmt.Println(v4)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if (!(!((fi14 - (fi14 ^ -1)) != fi14))) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, fi14)
        }
        _ = fi14
    }
    v13 = -1
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            v6 += 132
            fmt.Println(v1)
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v3 = (v3 ^ (^(v3)))
                    _ = fi16
                }
                _ = fi15
            }
        }
        if (!true) {
            var v17 int16 = v13
            for fi18 := 0; fi18 < 2; fi18++ {
                v4 += (v2 >> ((v3) & 63))
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 float64 = (v1 / v1)
                    v20 = 1e10
                    var v21 uint32 = 42
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                if false {
                    var v22 string = (("hello" + "a") + "\n\t")
                    _ = v22
                } else {
                    switch int(v3) {
                    case -100:
                        v23 := GenS2{F0: 42, F1: 1, F2: -9223372036854775808}
                        _ = v23
                    case 893:
                        v13 += (v4 - v13)
                    }
                    if true {
                        if ((uint8(v6) == 215) && (204 != 2)) {
                            var v24 int16 = -319
                            _ = v24
                        } else {
                            v25 := GenS3{F0: 65535}
                            fmt.Println(v3, v4)
                            _ = v25
                        }
                        v26 := GenS4{F0: 62, F1: "a"}
                        _ = v26
                    }
                    var v27 uint = v3
                    _ = v27
                }
                _ = fi18
            }
            switch 0 {
            case 84:
                v28 := GenS2{F0: 796, F1: 100, F2: 0}
                _ = v28
            }
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
