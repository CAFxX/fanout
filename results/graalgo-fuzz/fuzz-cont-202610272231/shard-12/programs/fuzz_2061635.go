package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 uint16 = 434
    var v3 uint16 = v2
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{fi4, 115}
        v6 := GenS1{F0: 42, F1: (2 <= 255), F2: false}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-1, -638}
    if (((float32(1.5) * 0.0) + 0.0) != float32(((0.0 + 0.0) * (-2.25 + (-2.25 * ((1e10 / 3.14159) + 1e10)))))) {
        var v8 int8 = 100
        var v9 uint32 = 65535
        v3 = (v2 | (v2 | 65535))
        _ = v8
        _ = v9
    }
    v3 += v3
    _ = len(v7)
    var v10 int = 2147483647
    for fi11 := 0; fi11 < 17; fi11++ {
        v3 -= (v2 - (^(v2)))
        _ = fi11
    }
    v2 = v2
    if false {
        v2 -= v3
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: 18446744073709551615, F1: (!true), F2: true}
            _ = fi12
            _ = v13
        }
    } else {
        if (42 < 808) {
            if (0.1 >= 0.0) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v1 = v1
                    switch 920 {
                    case -9223372036854775808:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    v2 = v3
                    _ = fi14
                }
                for fi15 := 0; fi15 < 6; fi15++ {
                    v1 = "a"
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v3 = (v2 + (v3 - (v2 ^ v2)))
                        _ = fi16
                    }
                    if false {
                        _ = len(v7)
                        v10 -= -473
                        if (false || ((305 >= 26) && (2147483647 >= (fi15 << ((876) & 63))))) {
                            if len(v7) > 0 {
                                v7[0] = 9223372036854775807
                            }
                            var v17 int16 = 32767
                            _ = v17
                        }
                    }
                    _ = fi15
                }
            }
        } else {
            var v18 int8 = -128
            v19 := GenS1{F0: 42, F1: true, F2: false}
            switch (v10 ^ -17) {
            case 9223372036854775807:
                switch v1 {
                case "foo bar":
                    var v20 int64 = 1
                    _ = v20
                case "\n\t":
                    fmt.Println(v3)
                }
            case 127:
                if (673 <= 0) {
                    v10 -= 32768
                    v21 := GenS2{F0: 0}
                    _ = v21
                }
            default:
                var v22 uint8 = 42
                _ = v22
            }
            _ = v18
            _ = v19
        }
        v10 = (^(v10))
    }
    var v23 bool = false
    if len(v7) > 0 {
        v7[0] = v10
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if len(v7) > 0 {
            v7[0] = 889
        }
        v1 = v1
        var v25 float64 = ((1.5 / 3.14159) + 0.1)
        _ = fi24
        _ = v25
    }
    v3 = 826
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v26 int = v10
    var v27 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v7
}
