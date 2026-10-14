package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 uint16, p1 int16) int {
    var v6 uint = 42
    v7 := []int{695, 32768, 2147483647}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint32 = 4294967295
    v3 := []string{"foo bar", "", ("hello" + "a"), "\n\t"}
    v4 := make(map[int]string)
    var v5 int = 127
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            v5 = (v5 | v5)
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint64 = 669
                if (46 == 158) {
                    var v11 int64 = 22
                    var v12 uint16 = 42
                    _ = v11
                    _ = v12
                } else {
                    v10 = v10
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        var v13 uint32 = v2
        _ = v13
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v4)
            _ = fi14
        }
        switch 32768 {
        case 9223372036854775807:
            fmt.Println(v1, v2)
        case 1:
            var v15 float64 = (0.0 * float64((float32(1.5) / float32(0.1))))
            _ = v15
        case -1:
            var v16 float64 = 1.5
            _ = v16
        default:
            var v17 int32 = 396
            _ = v17
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(fi18)
            fmt.Println(v2, fi18)
            _ = fi18
        }
    }
    v19 := GenS1{F0: (false && false), F1: -1}
    var v20 float64 = (0.0 - -2.25)
    v20 = (v20 / v20)
    v21 := v4[(v5 & v5)]
    fmt.Println(v2, v1)
    v22 := GenS1{F0: (953 != -209), F1: 42}
    for fi23 := 0; fi23 < 14; fi23++ {
        switch "\n\t" {
        case "\n\t":
            var v24 int8 = -128
            _ = v24
        default:
            var v25 uint = 614
            _ = v25
        }
        if false {
            switch -2147483648 {
            case -9223372036854775808:
                var v26 string = v21
                _ = v26
            case -635:
                v21 = v21
            case -436:
                switch (v21 + v21) {
                case "hello":
                    var v27 int32 = 42
                    _ = v27
                case "\n\t":
                    var v28 uint64 = 4294967295
                    _ = v28
                }
            }
            v29 := GenS1{F0: (true && (!true)), F1: -226}
            fmt.Println(v1, fi23)
            _ = v29
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
