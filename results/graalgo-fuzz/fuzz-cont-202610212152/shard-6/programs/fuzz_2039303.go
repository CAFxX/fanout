package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 uint }
func genFn0(p0 int8) uint16 {
    if true {
        var v4 int64 = 42
        _ = v4
    }
    switch "\n\t" {
    case "":
        v5 := []int{0}
        _ = v5
    }
    v6 := []int{128, 127, 2147483647, 523}
    _ = p0
    _ = v6
    return 374
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = ((0.0 + 1.5) + -2.25)
    v3 := []uint32{4294967295, 42}
    switch -9223372036854775808 {
    case 784:
        var v7 int = -2147483648
        _ = v7
    }
    _ = len(v3)
    v2 = v2
    v2 += (v2 / v2)
    switch "a" {
    case "hello":
        v8 := GenS1{F0: 956, F1: 0, F2: 18446744073709551615}
        _ = v8
    case "\n\t":
        v9 := v1[2147483647]
        _ = v9
    }
    var v10 int16 = -853
    var v11 string = "\n\t"
    switch 128 {
    case 32767:
        var v12 int64 = 9223372036854775807
        _ = v12
    case -146:
        switch 9 {
        case -509:
            if len(v3) > 0 {
                v3[0] = 0
            }
        case -1:
            fmt.Println(v10)
        default:
            var v13 int16 = 0
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            v10 -= int16(fi14)
            switch fi14 {
            case -450:
                var v15 uint64 = 18446744073709551615
                _ = v15
            default:
                var v16 uint64 = 0
                _ = v16
            }
            _ = fi14
        }
    }
    if (-37 <= v10) {
        var v17 int = 42
        for fi18 := 0; fi18 < 13; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                switch 1 {
                case 2147483647:
                    if true {
                        v20 := GenS1{F0: fi19, F1: (v10 >> ((42) & 63)), F2: 1}
                        switch -591 {
                        case 128:
                            var v21 int = (^(fi19))
                            _ = v21
                        case 95:
                            var v22 float64 = 0.0
                            _ = v22
                        default:
                            v23 := GenS1{F0: (fi18 % ((fi19) | 1)), F1: (v10 - v10), F2: 499}
                            _ = v23
                        }
                        _ = v20
                    } else {
                        v10 -= (v10 * v10)
                        switch fi19 {
                        case -227:
                            fmt.Println(v10, v10)
                        }
                        switch ((v11 + (v11 + (v11 + "a"))) + v11) {
                        case "hello":
                            v11 = v11
                        case "\n\t":
                            switch (v17 % ((fi18) | 1)) {
                            case 127:
                                fmt.Println(fi19, fi18)
                            case 128:
                                v3 = append(v3, 42)
                            case -1:
                                v24 := GenS1{F0: (fi18 >> ((42) & 63)), F1: v10, F2: 1}
                                _ = v24
                            }
                        case "foo bar":
                            fmt.Println(v11, v17, v10)
                        }
                    }
                    fmt.Println(fi18, fi19)
                case 173:
                    v11 = v11
                    if false {
                        var v25 string = v11
                        v25 = v11
                        _ = v25
                    } else {
                        v26 := GenS1{F0: -9223372036854775808, F1: (v10 % ((-32768) | 1)), F2: 423}
                        v27 := GenS2{F0: 426}
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v10, v17)
                case 127:
                    if ((v11 + v11) < "\n\t") {
                        var v28 int8 = 100
                        _ = v28
                    }
                default:
                    fmt.Println(v11, fi18, v17)
                }
                _ = fi19
            }
            fmt.Println(v10)
            v10 -= (v10 | 706)
            _ = fi18
        }
        _ = v17
    }
    var v29 uint32 = 651
    v2 = v2
    fmt.Println(v10, v10, v10)
    var v30 int16 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
