package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    var v1 bool = false
    v2 := []string{"\n\t", "foo bar", "foo bar"}
    var v3 int32 = 1
    var v4 int8 = 100
    v5 := []int{int(v4), -1, 1}
    var v6 bool = v1
    var v7 bool = true
    v7 = (!(9223372036854775807 > 1))
    v1 = (!(v1 && v6))
    fmt.Println(v7)
    v1 = v7
    v3 -= (-(v3))
    if v7 {
        v3 -= (v3 >> ((65535) & 63))
        var v8 int8 = v4
        if len(v5) > 0 {
            v5[0] = -498
        }
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 235, F2: 1}
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 42, F1: 0, F2: -489}
        for fi12 := 0; fi12 < 4; fi12++ {
            v4 += (v4 + -87)
            var v13 string = ("\n\t" + (("\n\t" + "foo bar") + "\n\t"))
            fmt.Println(fi12, v7, fi12)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if (false || (32767 >= 1000)) {
        v7 = v7
        var v14 uint8 = 128
        for fi15 := 0; fi15 < 12; fi15++ {
            if v6 {
                v3 = 2147483647
                v16 := GenS1{F0: 319, F1: 523, F2: -9223372036854775808}
                v1 = (1 <= 519)
                _ = v16
            } else {
                v7 = (float32(((0.0 * 1.5) - 1e10)) <= -2.25)
                if false {
                    v14 += (^(v14))
                } else {
                    switch "" {
                    case "foo bar":
                        if (false && (-291 < (fi15 % ((fi15) | 1)))) {
                            var v17 uint64 = 18446744073709551615
                            _ = v17
                        }
                    case "hello":
                        for fi18 := 0; fi18 < 7; fi18++ {
                            switch "" {
                            case "hello":
                                var v19 uint64 = 2
                                _ = v19
                            case "a":
                                if (!(0 > 1)) {
                                    if (0.1 >= 0.1) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        v20 := GenS1{F0: 42, F1: 18446744073709551615, F2: -9223372036854775808}
                                        if v6 {
                                            if (-2.25 <= 1e10) {
                                                var v21 bool = false
                                                _ = v21
                                            }
                                            v1 = ((v14 | 220) < (v14 + (v14 >> ((18446744073709551615) & 63))))
                                        } else {
                                            _ = len(v2)
                                            switch "a" {
                                            case "":
                                                fmt.Println(v3, fi18, v7)
                                            case "a":
                                                var v22 int8 = -82
                                                _ = v22
                                            case "\n\t":
                                                if (!true) {
                                                    fmt.Println(v14, v1, fi18)
                                                } else {
                                                    v23 := GenS2{F0: 2147483647, F1: "hello"}
                                                    _ = v23
                                                }
                                            }
                                        }
                                        var v24 string = "\n\t"
                                        _ = v20
                                        _ = v24
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    v25 := GenS1{F0: -1, F1: 18446744073709551615, F2: 42}
                                    _ = v25
                                } else {
                                    var v26 int = -2147483648
                                    fmt.Println(v6)
                                    v27 := GenS1{F0: -32768, F1: 0, F2: -1}
                                    v1 = v7
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                var v28 uint32 = 2147483648
                                _ = v28
                            }
                            var v29 int = 42
                            switch "a" {
                            case "\n\t":
                                if ((fi15 ^ -962) <= (fi15 / ((fi15) | 1))) {
                                    fmt.Println(fi15, v14, v3)
                                } else {
                                    v4 -= v4
                                }
                            case "hello":
                                var v30 string = "hello"
                                _ = v30
                            }
                            _ = fi18
                            _ = v29
                        }
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v31 float64 = -2.25
            _ = fi15
            _ = v31
        }
        _ = v14
    }
    v7 = false
    v3 = (v3 << ((659) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
