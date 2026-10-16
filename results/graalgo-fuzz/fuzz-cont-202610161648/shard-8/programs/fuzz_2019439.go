package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int = -2147483648
    var v2 uint16 = 32767
    var v3 uint64 = 42
    fmt.Println(v3, v3)
    fmt.Println(v3)
    fmt.Println(v1)
    v2 = (v2 + v2)
    if true {
        switch 0 {
        case 1:
            if ((v1 * -84) <= v1) {
                var v4 uint = 2
                v2 += 255
                _ = v4
            } else {
                if (255 == 147) {
                    v5 := []int{-9223372036854775808}
                    _ = v5
                }
                v6 := GenS1{F0: int32(v2), F1: (^(v2)), F2: false}
                var v7 int16 = 92
                _ = v6
                _ = v7
            }
        case 0:
            var v8 float64 = 1e10
            _ = v8
        case -9223372036854775808:
            v1 = -1
        default:
            v9 := []int{v1}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: -1, F1: v2, F2: (!(1 <= -876))}
            _ = fi10
            _ = v11
        }
        switch "\n\t" {
        case "":
            v12 := GenS1{F0: int32(v1), F1: (v2 / ((1) | 1)), F2: ("foo bar" != "a")}
            _ = v12
        case "\n\t":
            v13 := GenS1{F0: 2147483647, F1: v2, F2: (2 < 42)}
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 14; fi14++ {
                v1 -= 127
                switch "\n\t" {
                case "hello":
                    if ("\n\t" == "hello") {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            fmt.Println(v3, fi15)
                            _ = fi15
                        }
                        v16 := GenS1{F0: 321, F1: (v2 * v2), F2: (281 == (v2 % (((^(v2))) | 1)))}
                        _ = v16
                    }
                case "foo bar":
                    var v17 int32 = -2147483648
                    _ = v17
                case "":
                    v2 += uint16(v3)
                default:
                    switch "hello" {
                    case "\n\t":
                        if true {
                            switch "a" {
                            case "hello":
                                var v18 uint = 2
                                _ = v18
                            default:
                                for fi19 := 0; fi19 < 15; fi19++ {
                                    v3 = (v3 % ((v3) | 1))
                                    _ = fi19
                                }
                            }
                            v20 := []int{-1, (v1 << ((4294967295) & 63))}
                            v21 := GenS1{F0: 1, F1: v2, F2: (fi14 > fi14)}
                            _ = v20
                            _ = v21
                        } else {
                            switch "a" {
                            case "hello":
                                v22 := []int{(v1 << ((255) & 63)), v1}
                                _ = v22
                            }
                        }
                    case "":
                        var v23 uint16 = v2
                        _ = v23
                    }
                }
                v24 := GenS1{F0: 2147483647, F1: (^(v2)), F2: ((float32((0.0 / ((1e10 * (float64(((0.0 * (0.1 * 1.5)) - 3.14159)) * float64(1.5))) - -2.25))) / -2.25) <= -2.25)}
                _ = fi14
                _ = v24
            }
        }
    }
    var v25 uint16 = (v2 + v2)
    v26 := GenS1{F0: 977, F1: (v25 / ((v25) | 1)), F2: (((104 != -10) || ((-(v1)) == 32767)) && true)}
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: 880, F1: (v25 * (v25 >> ((4294967295) & 63))), F2: ("\n\t" >= "hello")}
        var v29 int8 = 120
        v3 -= (v3 | (v3 + (v3 << ((18446744073709551615) & 63))))
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 string = "\n\t"
    var v31 int8 = -31
    fmt.Println(v2)
    if (1e10 == 0.0) {
        fmt.Println(v1, v3, v31)
        if true {
            var v32 float32 = 1.5
            _ = v32
        }
    }
    if true {
        v33 := GenS2{F0: 1000, F1: (v3 / ((v3) | 1))}
        _ = v33
    } else {
        switch (v1 & v1) {
        case 2147483647:
            var v34 float32 = 1.5
            _ = v34
        case 42:
            if true {
                fmt.Println(v1, v30)
            }
        case 1:
            var v35 int64 = 42
            _ = v35
        default:
            switch v1 {
            case -678:
                var v36 float32 = 1.5
                _ = v36
            case 0:
                v31 = (v31 / ((v31) | 1))
            case -480:
                var v37 uint64 = v3
                _ = v37
            default:
                fmt.Println(v30)
            }
        }
        var v38 int64 = 42
        _ = v38
    }
    var v39 int = v1
    switch (v39 & (v39 + -9223372036854775808)) {
    case 571:
        v40 := []int{(v1 | v39), (v1 / ((v39) | 1)), -2147483648}
        _ = v40
    default:
        v41 := []int{(v39 / ((v1) | 1)), v1}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v39
}
