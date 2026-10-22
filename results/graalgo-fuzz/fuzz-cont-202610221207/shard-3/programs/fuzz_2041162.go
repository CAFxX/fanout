package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int32, p2 uint) int16 {
    p1 += -2147483648
    fmt.Println(p2, p2, p0)
    v7 := []int{622, 42, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 390
}
func genFn1(p0 int, p1 string) uint32 {
    p1 = ""
    var v8 int8 = 42
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    var v3 uint64 = (v1 & v1)
    v4 := []string{"foo bar", "a"}
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    if true {
        v3 += (v1 / (((^(v3))) | 1))
        if (-540 >= 1) {
            if true {
                v1 = (v3 * (v1 >> ((1) & 63)))
                v1 -= v3
            } else {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1, v3)
                case "\n\t":
                    fmt.Println(v3)
                case "a":
                    var v9 int16 = -32768
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 1; fi10++ {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v12 float64 = 0.0
                            var v13 float32 = (((((-2.25 - -2.25) - 0.1) - 1e10) / 3.14159) / float32(0.1))
                            v14 := GenS1{F0: 1}
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v1 += (v1 | v1)
                        _ = len(v2)
                        _ = fi10
                    }
                }
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    fmt.Println(v1, fi15)
                    fmt.Println(v3, v1, fi16)
                    _ = fi16
                }
                switch 128 {
                case -1:
                    v3 += 42
                case 32767:
                    switch (fi15 / ((fi15) | 1)) {
                    case 32767:
                        fmt.Println(v1)
                    }
                case 0:
                    var v17 float64 = (((3.14159 / 1.5) - 0.1) / -2.25)
                    _ = v17
                }
                if (65535 < 4294967295) {
                    var v18 uint = 4294967295
                    v1 += 638
                    _ = v18
                } else {
                    v19 := GenS1{F0: -577}
                    v3 = v3
                    v20 := GenS2{F0: 127, F1: 42}
                    _ = v19
                    _ = v20
                }
                _ = fi15
            }
        }
    } else {
        fmt.Println(v1, v1)
    }
    v4 = append(v4, "a")
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS2{F0: -128, F1: 18446744073709551615}
            v3 = 4294967295
            _ = fi22
            _ = v23
        }
        if ("" != (("a" + "hello") + (("hello" + "hello") + "\n\t"))) {
            v1 -= v1
            if (988 <= 350) {
                v3 = v3
                v24 := GenS3{F0: 65535, F1: fi21}
                var v25 float64 = 0.0
                _ = v24
                _ = v25
            }
        }
        v3 += v1
        _ = fi21
    }
    if (-1 <= 42) {
        fmt.Println(v1, v3)
    } else {
        var v26 bool = true
        var v27 int16 = 217
        _ = v26
        _ = v27
    }
    fmt.Println(v3)
    switch -295 {
    case 42:
        var v28 float32 = 1e10
        _ = v28
    case 9223372036854775807:
        v1 = (v1 ^ (v3 ^ (v1 << ((1) & 63))))
    case 2147483647:
        v4 = append(v4, "a")
    default:
        v1 -= (v1 - (v1 >> ((42) & 63)))
    }
    if ("hello" > ("" + "")) {
        fmt.Println(v1, v1)
        var v29 float64 = 1.5
        _ = v29
    } else {
        _ = len(v6)
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 string = "\n\t"
        v32 := GenS1{F0: 668}
        fmt.Println(v3)
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS4{F0: 9223372036854775807, F1: -1, F2: 127}
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
