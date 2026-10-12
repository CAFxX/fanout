package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 uint64 = 42
    v5 := []string{("a" + "hello"), ""}
    if ((false && (!(121 >= 0))) && true) {
        v6 := GenS1{F0: true, F1: (v1 - (-(v1))), F2: 2147483647}
        _ = v6
    }
    switch 32767 {
    case 2147483647:
        if true {
            var v7 float64 = (3.14159 * (((float64(0.1) * (1.5 / ((-2.25 / 0.1) / 1.5))) + ((1e10 - (0.0 + 3.14159)) - ((((-2.25 / (3.14159 + 3.14159)) * float64((3.14159 + 0.0))) - 3.14159) + (((float64(0.0) / (0.1 + 1.5)) + float64((0.1 + 0.1))) + (0.1 - -2.25))))) / 0.1))
            fmt.Println(v1, v4, v4)
            var v8 bool = (2 < 195)
            _ = v7
            _ = v8
        }
    default:
        v9 := GenS1{F0: true, F1: (-(v1)), F2: 42}
        _ = v9
    }
    var v10 int16 = 1
    if (-2.25 == (-2.25 - (-2.25 * 1e10))) {
        switch -778 {
        case 128:
            v4 = (v4 - (v4 + (v4 & v4)))
        case 137:
            v11 := GenS1{F0: false, F1: 9223372036854775807, F2: 4294967295}
            _ = v11
        }
        if (true || (18446744073709551615 < 2)) {
            v5 = append(v5, "")
        }
    }
    switch "foo bar" {
    case "":
        var v12 string = "\n\t"
        _ = v12
    case "a":
        switch "hello" {
        case "a":
            if ("" <= ("hello" + "a")) {
                var v13 float64 = (1.5 * (0.0 / ((0.1 / 1.5) + -2.25)))
                var v14 int8 = 127
                _ = v13
                _ = v14
            } else {
                v4 = (v4 & (v4 / ((4294967295) | 1)))
            }
        default:
            v10 += (v10 + (v10 ^ (v10 & (v10 - 508))))
        }
    case "\n\t":
        if (!true) {
            var v15 uint32 = 0
            _ = v15
        } else {
            if (false && (0.0 != -2.25)) {
                v4 -= (v4 << ((2) & 63))
                v16 := GenS1{F0: true, F1: (v1 >> ((0) & 63)), F2: 925}
                v1 = 1
                _ = v16
            }
            var v17 uint64 = (v4 - 0)
            _ = v17
        }
    }
    v4 = (v4 ^ v4)
    for fi18 := 0; fi18 < 9; fi18++ {
        v1 = v1
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint16 = 0
            _ = fi19
            _ = v20
        }
        var v21 uint64 = (^(v4))
        _ = fi18
        _ = v21
    }
    v5 = append(v5, ("foo bar" + ("a" + "\n\t")))
    v22 := v2[65535]
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 bool = true
        for fi25 := 0; fi25 < 7; fi25++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    var v26 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
