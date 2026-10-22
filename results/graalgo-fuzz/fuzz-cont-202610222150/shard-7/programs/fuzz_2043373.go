package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    var v4 int32 = 756
    v5 := []int{0, -89, 2147483647}
    v4 += (v4 ^ (v4 | v4))
    var v6 float32 = (1e10 * (3.14159 - (float32((float64(((-2.25 - 1e10) + 0.0)) - ((0.1 - 1.5) + 1e10))) - 3.14159)))
    if (((-355 == 42) || false) && (!(4294967295 == 4294967295))) {
        var v7 uint8 = 255
        var v8 uint64 = 894
        _ = v7
        _ = v8
    } else {
        v4 += v4
    }
    v6 -= 3.14159
    var v9 uint16 = 176
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            v6 -= v6
            if false {
                v12 := GenS1{F0: 9223372036854775807}
                v2[42] = uint32(v9)
                v13 := GenS2{F0: int(v4), F1: 42}
                _ = v12
                _ = v13
            } else {
                var v14 uint32 = 874
                _ = v14
            }
            v15 := GenS1{F0: 9223372036854775807}
            var v16 string = "\n\t"
            _ = fi11
            _ = v15
            _ = v16
        }
        v4 = v4
        v4 = 3
        _ = fi10
    }
    var v17 bool = (!(866 > 255))
    if (0 == 42) {
        switch 0 {
        case 235:
            var v18 uint64 = 2
            _ = v18
        case -9223372036854775808:
            fmt.Println(v9, v9)
        }
        v17 = (683 <= 4294967295)
    }
    switch 127 {
    case 42:
        var v19 int64 = 907
        _ = v19
    case 32768:
        for fi20 := 0; fi20 < 14; fi20++ {
            v6 = v6
            v21 := v2[984]
            _ = fi20
            _ = v21
        }
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if (((v6 / ((3.14159 + v6) / (v6 / v6))) * float32(1e10)) == float32(((1e10 * 1e10) + 0.1))) {
        var v22 int = 128
        v23 := GenS1{F0: 823}
        for fi24 := 0; fi24 < 7; fi24++ {
            v17 = true
            _ = fi24
        }
        _ = v22
        _ = v23
    }
    switch "" {
    case "hello":
        v25 := GenS2{F0: 4294967295, F1: 18446744073709551615}
        _ = v25
    case "\n\t":
        for fi26 := 0; fi26 < 18; fi26++ {
            if ((true || (true && v17)) && ((true && v17) && (("foo bar" <= "a") && (true && ((118 != 700) && v17))))) {
                v9 += v9
                for fi27 := 0; fi27 < 16; fi27++ {
                    switch "a" {
                    case "hello":
                        if (2147483648 != 0) {
                            v28 := GenS2{F0: (fi26 + 1), F1: 1}
                            v9 = 2
                            _ = v28
                        }
                    }
                    _ = fi27
                }
            }
            _ = fi26
        }
    }
    var v29 int8 = 0
    for fi30 := 0; fi30 < 2; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 bool = false
            v9 -= (v9 | (v9 % ((v9) | 1)))
            v33 := GenS1{F0: 42}
            _ = fi31
            _ = v32
            _ = v33
        }
        v17 = v17
        _ = fi30
    }
    var v34 int16 = 42
    if (0.0 < v6) {
        v17 = false
    } else {
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 int32 = 2147483647
            v4 = -769
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
