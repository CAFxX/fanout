package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }

func main() {
    v1 := []uint32{436, 264, 42, 2147483647}
    v2 := make(map[string]bool)
    v3 := []uint32{2147483648, 900}
    var v4 uint32 = 2
    var v5 int64 = -9223372036854775808
    v6 := []int{32767, -1, -459}
    v4 = v4
    _ = len(v2)
    if ((!false) || false) {
        var v7 float64 = float64(1.5)
        v7 = v7
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v4 = v4
        _ = len(v3)
        switch 42 {
        case 32768:
            v2["hello"] = true
        }
        _ = fi8
    }
    fmt.Println(v5, v5)
    v9 := GenS1{F0: 42, F1: v5, F2: 33}
    for fi10 := 0; fi10 < 19; fi10++ {
        v5 = (v5 - (v5 - (v5 >> ((42) & 63))))
        if (!((v4 >> ((675) & 63)) > (^(v4)))) {
            for fi11 := 0; fi11 < 2; fi11++ {
                v4 = (^(v4))
                v4 = 2147483648
                _ = fi11
            }
            if (-128 > 42) {
                switch "a" {
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        var v12 int16 = -876
                        _ = v12
                    }
                case "foo bar":
                    v5 = v5
                case "a":
                    v13 := v2["foo bar"]
                    _ = v13
                default:
                    v14 := GenS1{F0: (^(v5)), F1: v5, F2: 1}
                    _ = v14
                }
                _ = len(v2)
            } else {
                switch ((("foo bar" + "hello") + "") + "\n\t") {
                case "a":
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 int64 = (v5 >> ((18446744073709551615) & 63))
                        _ = fi15
                        _ = v16
                    }
                case "hello":
                    v17 := GenS1{F0: (v5 * v5), F1: v5, F2: 0}
                    _ = v17
                }
                if false {
                    v4 += v4
                    _ = len(v2)
                } else {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if len(v6) > 0 {
                            v6[0] = fi10
                        }
                        _ = fi18
                    }
                    if ((((0.0 + -2.25) - 0.0) * 0.0) >= 0.1) {
                        for fi19 := 0; fi19 < 4; fi19++ {
                            var v20 int64 = (v5 & (v5 ^ v5))
                            v21 := GenS1{F0: 383, F1: -9223372036854775808, F2: -1}
                            var v22 int = (fi19 << ((18446744073709551615) & 63))
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS1{F0: -510, F1: 734, F2: 1}
                        v24 := GenS1{F0: v5, F1: v5, F2: 48}
                        _ = v23
                        _ = v24
                    }
                    v4 += v4
                }
            }
            v5 = (^(v5))
        } else {
            v5 += 9223372036854775807
            for fi25 := 0; fi25 < 14; fi25++ {
                if false {
                    v26 := GenS2{F0: 1, F1: "\n\t", F2: 2}
                    v27 := v2["hello"]
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v28 := GenS3{F0: v27}
                    v5 += (v5 - 790)
                    fmt.Println(v27)
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v4, fi25)
                    v4 = 48
                    var v29 int8 = int8(fi25)
                    _ = v29
                }
                v30 := GenS4{F0: 1}
                _ = fi25
                _ = v30
            }
            v5 += v5
        }
        var v31 int32 = int32(fi10)
        _ = fi10
        _ = v31
    }
    v32 := GenS3{F0: (!true)}
    v3 = append(v3, (v4 | (v4 * (v4 ^ (v4 - 2147483648)))))
    fmt.Println(v4, v5)
    _ = len(v3)
    var v33 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
