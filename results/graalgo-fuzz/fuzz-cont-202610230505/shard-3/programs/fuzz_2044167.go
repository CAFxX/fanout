package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 int, p1 uint8) uint64 {
    var v7 int32 = 42
    var v8 uint = 48
    var v9 int = (^(p0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 42
    v2 := []int{748, -70}
    var v3 uint8 = 0
    v4 := []int{}
    var v5 int32 = 624
    v6 := make(map[string]int)
    v5 -= -842
    v5 = v5
    var v10 bool = true
    fmt.Println(v3)
    v11 := GenS1{F0: 127, F1: (65535 == 4294967295), F2: "foo bar"}
    if true {
        if (1 == 786) {
            v12 := GenS1{F0: 852, F1: true, F2: "hello"}
            if v10 {
                if v10 {
                    v13 := v6["foo bar"]
                    for fi14 := 0; fi14 < 15; fi14++ {
                        for fi15 := 0; fi15 < 19; fi15++ {
                            _ = len(v6)
                            if true {
                                fmt.Println(fi14, v1)
                                var v16 float32 = ((((-2.25 * -2.25) * 0.0) * (0.0 - 1e10)) * 1e10)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v17 := GenS2{F0: 0}
                                _ = v16
                                _ = v17
                            } else {
                                var v18 uint8 = v3
                                var v19 uint16 = 0
                                v19 = 536
                                _ = v18
                                _ = v19
                            }
                            fmt.Println(v3, v1)
                            switch v13 {
                            case 1:
                                var v20 int64 = 42
                                v21 := GenS1{F0: v13, F1: false, F2: (((("" + ("\n\t" + "\n\t")) + "foo bar") + "a") + "foo bar")}
                                _ = v20
                                _ = v21
                            case -67:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                delete(v6, "")
                            }
                            var v22 uint64 = v1
                            _ = fi15
                            _ = v22
                        }
                        _ = fi14
                    }
                    _ = v13
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v23 := GenS3{F0: -552, F1: -500}
                    var v24 int16 = 1000
                    _ = v23
                    _ = v24
                }
                v25 := GenS4{F0: (^(v1)), F1: -9223372036854775808}
                var v26 float64 = -2.25
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 int64 = 1
                    v10 = (42 > 2147483648)
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    v6["foo bar"] = fi29
                    _ = fi29
                }
                _ = len(v6)
            }
            switch "\n\t" {
            case "\n\t":
                switch 2147483647 {
                case 948:
                    v10 = false
                default:
                    v30 := GenS3{F0: 907, F1: 127}
                    _ = v30
                }
            }
            _ = v12
        } else {
            v31 := GenS1{F0: 42, F1: (float64(3.14159) == float64(0.1)), F2: ("\n\t" + "\n\t")}
            var v32 int = -9223372036854775808
            _ = v31
            _ = v32
        }
        v1 = v1
        v33 := GenS2{F0: 4294967295}
        _ = v33
    }
    if v10 {
        switch -9223372036854775808 {
        case 735:
            switch 1 {
            case 943:
                var v34 float64 = 0.0
                _ = v34
            default:
                v35 := GenS3{F0: 1, F1: 9223372036854775807}
                _ = v35
            }
        }
    }
    v1 -= (v1 % (((v1 << ((255) & 63))) | 1))
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 int64 = 9223372036854775807
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
