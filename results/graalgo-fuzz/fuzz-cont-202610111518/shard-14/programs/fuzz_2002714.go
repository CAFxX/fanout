package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64) int32 {
    v5 := GenS1{F0: 1, F1: (-(p0)), F2: -1}
    var v6 int16 = -528
    v7 := GenS1{F0: 128, F1: p0, F2: 1000}
    var v8 uint = 18446744073709551615
    v6 -= (v6 | (v6 & v6))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 840
}

func main() {
    v1 := []int{2147483647}
    var v2 string = "a"
    v3 := []string{("a" + v2), (v2 + "foo bar"), (((((((("\n\t" + (v2 + v2)) + v2) + v2) + (v2 + (v2 + "hello"))) + v2) + (v2 + "")) + v2) + "foo bar")}
    var v4 uint16 = 42
    var v9 bool = false
    var v10 int16 = -1
    fmt.Println(v4, v10, v10)
    if v9 {
        if v9 {
            for fi11 := 0; fi11 < 16; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    if v9 {
                        if len(v1) > 0 {
                            v1[0] = (fi11 | (fi11 / ((32767) | 1)))
                        }
                        if (true || (0 >= 0)) {
                            fmt.Println(fi12, v4, v9)
                            if (842 < (v4 >> ((923) & 63))) {
                                var v13 int16 = (v10 / ((v10) | 1))
                                _ = v13
                            } else {
                                v9 = v9
                                v2 = v2
                                var v14 int8 = 0
                                if v9 {
                                    _ = len(v3)
                                } else {
                                    v15 := GenS2{F0: 42, F1: 689}
                                    var v16 uint8 = 255
                                    _ = v15
                                    _ = v16
                                }
                                fmt.Println(fi11)
                                _ = v14
                            }
                            var v17 int32 = 145
                            _ = v17
                        }
                    } else {
                        switch (v2 + v2) {
                        case "hello":
                            v18 := GenS3{F0: "a", F1: ((0 != 330) || (!false))}
                            _ = v18
                        case "":
                            var v19 uint16 = v4
                            _ = v19
                        }
                        v20 := GenS1{F0: (fi11 ^ 507), F1: -9223372036854775808, F2: 83}
                        var v21 uint8 = 127
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(fi11)
                    _ = fi12
                }
                _ = fi11
            }
            if v9 {
                v10 -= v10
                fmt.Println(v9, v4, v2)
            } else {
                _ = len(v1)
                var v22 int32 = int32(v10)
                _ = v22
            }
        }
        if (false || false) {
            v10 -= v10
            v9 = ((false || v9) && (127 < 255))
        }
        v3 = append(v3, (v2 + v2))
    }
    switch v2 {
    case "\n\t":
        v2 = v2
    case "":
        var v23 uint16 = (v4 * v4)
        _ = v23
    case "a":
        fmt.Println(v9, v10)
    }
    v24 := GenS4{F0: 0}
    fmt.Println(v9, v2)
    var v25 int32 = -2147483648
    var v26 uint32 = 4294967295
    v4 = (v4 - (v4 | v4))
    v27 := GenS1{F0: 36, F1: -536, F2: (v10 | v10)}
    if true {
        v10 += v10
        var v28 uint64 = 42
        v2 = v2
        _ = v28
    } else {
        switch "\n\t" {
        case "":
            var v29 uint64 = 4294967295
            _ = v29
        case "foo bar":
            if (1.5 <= ((float64((1.5 / (-2.25 / (1.5 / 3.14159)))) * (((0.0 * -2.25) * (((3.14159 + -2.25) / ((((1.5 / -2.25) * (0.1 + 0.0)) - 1.5) + (0.0 * (1e10 / (((((3.14159 * float64((3.14159 * (0.1 / (((float32(0.0) * float32(1.5)) + 1.5) / 1e10))))) + 1e10) * -2.25) - 3.14159) * float64(1.5)))))) * 0.1)) - 0.1)) / 1e10)) {
                v9 = ((32767 > 127) || v9)
                for fi30 := 0; fi30 < 6; fi30++ {
                    v10 = v10
                    _ = fi30
                }
            }
        case "a":
            v3 = append(v3, ("foo bar" + v2))
        default:
            v9 = (v9 && ((!v9) && v9))
        }
        v31 := GenS1{F0: -9223372036854775808, F1: 0, F2: (v10 - 32767)}
        var v32 int8 = -37
        _ = v31
        _ = v32
    }
    v33 := GenS2{F0: 2, F1: 4294967295}
    v25 += v25
    var v34 float64 = 1e10
    var v35 bool = (((!(103 <= 53)) || ((v2 + (("hello" + v2) + v2)) < ((v2 + "foo bar") + "\n\t"))) || false)
    v36 := GenS2{F0: 873, F1: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v9
}
