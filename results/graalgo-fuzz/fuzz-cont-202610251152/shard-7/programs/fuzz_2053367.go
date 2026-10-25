package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }

func main() {
    var v1 int8 = -1
    var v2 int = 1
    v3 := make(map[int]bool)
    var v4 uint = 0
    var v5 uint = (v4 ^ (v4 << ((v4) & 63)))
    var v6 bool = (v2 == (v2 >> (((^(v4))) & 63)))
    var v7 uint32 = 99
    v2 += (-(v2))
    switch ("a" + "foo bar") {
    case "hello":
        v6 = v6
    case "foo bar":
        var v8 uint = 18446744073709551615
        _ = v8
    case "a":
        v1 = (v1 | (v1 / ((42) | 1)))
    default:
        var v9 int64 = -859
        _ = v9
    }
    switch (^(v2)) {
    case 1:
        v10 := GenS1{F0: 127, F1: 42, F2: (true && v6)}
        _ = v10
    case 32767:
        var v11 int32 = 0
        _ = v11
    case 2147483647:
        v3[-414] = v6
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        delete(v3, (fi12 - (v2 - v2)))
        var v13 uint64 = 206
        var v14 int64 = 9223372036854775807
        _ = fi12
        _ = v13
        _ = v14
    }
    if (1 <= 2) {
        v2 = v2
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                if v6 {
                    if v6 {
                        if false {
                            var v17 bool = (-574 == -9223372036854775808)
                            var v18 float32 = (float32(1.5) - float32(1e10))
                            var v19 int8 = (v1 + v1)
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{}
                        var v21 int32 = -2147483648
                        v22 := v3[(fi16 >> ((4294967295) & 63))]
                        v3[fi16] = (v6 && v22)
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        v6 = v6
                    }
                    if v6 {
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                var v23 float64 = float64(0.1)
                                _ = v23
                            case "a":
                                v5 = (^(v4))
                            }
                            var v24 int = (-(fi16))
                            v25 := GenS1{F0: 40, F1: (v5 << ((18446744073709551615) & 63)), F2: (1e10 > 3.14159)}
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v6, fi16, fi15)
                        }
                    }
                }
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    if (v6 || (4294967295 > int(v4))) {
                        fmt.Println(v6, v4, fi16)
                    } else {
                        switch fi15 {
                        case 32768:
                            if (!(42 == 1)) {
                                v1 -= (v1 | (v1 ^ v1))
                            } else {
                                v5 = v4
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            }
                        case 517:
                            if ("" >= "foo bar") {
                                var v27 bool = (!v6)
                                _ = v27
                            }
                        }
                    }
                case "foo bar":
                    v4 = 266
                case "hello":
                    v3[fi16] = v6
                default:
                    v28 := v3[(fi15 % (((v2 - (^(fi16)))) | 1))]
                    _ = v28
                }
                _ = fi16
            }
            _ = fi15
        }
        if false {
            if true {
                v4 += (v4 & 824)
            } else {
                delete(v3, (v2 * v2))
            }
            _ = len(v3)
        } else {
            v5 = 657
            delete(v3, v2)
            var v29 float64 = float64((((-2.25 * 0.1) / 0.1) / 1e10))
            _ = v29
        }
    } else {
        var v30 int8 = (v1 & (-(v1)))
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v4 += v4
        _ = fi31
    }
    v32 := GenS1{F0: 1, F1: v4, F2: (!(243 >= 598))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
