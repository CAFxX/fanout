package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 1000
    v3 := []uint32{4294967295, 962, 2147483648}
    var v4 uint = 1
    v5 := make(map[int]int)
    var v6 bool = false
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v2, fi7)
        var v8 float64 = 1e10
        var v9 string = "hello"
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 197
    switch 32768 {
    case 511:
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                v10 = (v10 + -1)
                _ = fi12
            }
            if v6 {
                v13 := GenS1{F0: (v1 / (((v1 - v1)) | 1)), F1: 1, F2: 1}
                v6 = true
                switch 0 {
                case 42:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 uint8 = 2
                        switch "foo bar" {
                        case "foo bar":
                            v16 := GenS1{F0: (v1 * v1), F1: 18446744073709551615, F2: uint8(v4)}
                            _ = v16
                        case "":
                            v3 = append(v3, 42)
                        case "\n\t":
                            fmt.Println(fi11)
                        default:
                            if ((!v6) || (4294967295 <= v4)) {
                                v17 := v5[9223372036854775807]
                                v4 += (v4 >> ((v4) & 63))
                                _ = v17
                            } else {
                                var v18 int8 = 100
                                _ = v18
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v13
            }
            _ = fi11
        }
    case 260:
        var v19 string = (("foo bar" + "") + (("a" + "\n\t") + "\n\t"))
        _ = v19
    default:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    v4 += 4294967295
    v2 = -1
    v20 := GenS1{F0: (v1 + (v1 * 1)), F1: 18446744073709551615, F2: uint8(v4)}
    fmt.Println(v10)
    v21 := GenS1{F0: (v1 + (v1 + 1)), F1: 1, F2: 159}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v10, v10, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
