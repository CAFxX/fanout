package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }

func main() {
    var v1 float32 = (((0.0 - -2.25) - ((3.14159 + 1.5) - 3.14159)) * (1e10 / -2.25))
    var v2 int8 = -128
    var v3 uint64 = 42
    var v4 uint32 = 0
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(v4, fi6)
            v7 := []int{32768, fi6, (fi5 & fi6), 2147483647}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v1 = v1
    v8 := GenS1{F0: v2, F1: 18446744073709551615, F2: ("foo bar" != "hello")}
    var v9 uint32 = 2147483647
    v10 := []int{-303, -1}
    v11 := GenS2{F0: -1, F1: 65535, F2: 1}
    for fi12 := 0; fi12 < 5; fi12++ {
        v9 += 964
        _ = fi12
    }
    var v13 int8 = 9
    switch "\n\t" {
    case "foo bar":
        v2 += v13
    case "hello":
        if len(v10) > 0 {
            v10[0] = 843
        }
    case "\n\t":
        switch -494 {
        case -9223372036854775808:
            if true {
                v9 += v9
            } else {
                var v14 int64 = -9223372036854775808
                v10 = append(v10, 127)
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint16 = uint16(v3)
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        case 0:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        case 821:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v2 = int8(fi17)
        var v18 float64 = 0.1
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v3, v4)
    switch -9223372036854775808 {
    case 146:
        for fi19 := 0; fi19 < 16; fi19++ {
            switch fi19 {
            case 724:
                for fi20 := 0; fi20 < 9; fi20++ {
                    switch "foo bar" {
                    case "a":
                        if ((v13 >> ((1) & 63)) >= v13) {
                            var v21 float64 = 1.5
                            if false {
                                v3 -= (v3 & (v3 << ((18446744073709551615) & 63)))
                            }
                            _ = v21
                        }
                    case "":
                        v13 = -128
                    default:
                        v13 -= v2
                    }
                    v22 := GenS1{F0: v2, F1: 2, F2: false}
                    var v23 float32 = float32(1e10)
                    var v24 uint32 = v9
                    if (!(v23 != float32((0.1 + 0.0)))) {
                        var v25 uint64 = 550
                        _ = v25
                    }
                    _ = fi20
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                if false {
                    var v27 int16 = -658
                    _ = v27
                }
                var v28 bool = true
                _ = fi26
                _ = v28
            }
            _ = fi19
        }
    }
    v13 = v13
    v9 = v4
    v4 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
