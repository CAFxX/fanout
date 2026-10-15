package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32) uint32 {
    var v6 int16 = 234
    v7 := GenS1{F0: 644}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 42
        if false {
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(v9)
                if (true && false) {
                    if true {
                        v11 := []int{fi8, (fi8 * -9223372036854775808)}
                        _ = v11
                    }
                }
                v12 := GenS1{F0: 2}
                _ = fi10
                _ = v12
            }
            v13 := GenS2{F0: -9223372036854775808, F1: false}
            _ = v13
        } else {
            var v14 uint64 = 18446744073709551615
            v14 = 758
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{(fi15 << ((909) & 63)), 4294967295, -1, (^(fi8))}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    v17 := []int{128, 42, 127, 32768}
    _ = p0
    _ = v17
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 int16 = 1
    var v4 int64 = 294
    var v5 float32 = -2.25
    v18 := []int{-505, -9223372036854775808}
    fmt.Println(v4, v4, v5)
    v19 := GenS1{F0: 2147483647}
    switch (v1 + (v1 + ("foo bar" + v1))) {
    case "foo bar":
        switch 0 {
        case -224:
            v20 := GenS3{F0: 1}
            _ = v20
        case 644:
            v4 = (v4 - v4)
        case 4294967295:
            for fi21 := 0; fi21 < 5; fi21++ {
                v4 += v4
                fmt.Println(v3, v4, v4)
                _ = fi21
            }
        default:
            v4 = (v4 & v4)
        }
    case "hello":
        if len(v18) > 0 {
            v18[0] = -9223372036854775808
        }
    default:
        var v22 int16 = (v3 * 14)
        _ = v22
    }
    switch v1 {
    case "hello":
        switch 32767 {
        case 4294967295:
            if true {
                if (!false) {
                    var v23 int32 = 2147483647
                    _ = v23
                } else {
                    var v24 bool = true
                    _ = v24
                }
                fmt.Println(v3, v4)
                if ((true && (-32768 <= (v3 % (((v3 >> ((444) & 63))) | 1)))) && true) {
                    v3 -= 750
                    var v25 uint32 = 2147483647
                    fmt.Println(v3)
                    _ = v25
                } else {
                    v26 := GenS2{F0: (-(v4)), F1: false}
                    v4 -= v4
                    _ = v26
                }
            } else {
                fmt.Println(v4, v5, v5)
                v1 = v1
            }
        default:
            fmt.Println(v3, v4, v1)
        }
    case "":
        var v27 uint = 255
        _ = v27
    default:
        v4 = v4
    }
    var v28 int64 = (-(v4))
    for fi29 := 0; fi29 < 9; fi29++ {
        if (3.14159 < 0.1) {
            v30 := GenS2{F0: 9223372036854775807, F1: (!false)}
            _ = v30
        } else {
            v28 -= v4
            var v31 uint32 = 2147483648
            _ = v31
        }
        _ = fi29
    }
    _ = len(v18)
    if ((((-2.25 + float64((v5 * (v5 * v5)))) - (float64(v5) + 1e10)) / 1.0) >= 0.0) {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    } else {
        if false {
            v5 = v5
            v32 := GenS4{F0: 696}
            var v33 uint = 1
            _ = v32
            _ = v33
        }
        v3 += v3
    }
    var v34 float32 = v5
    v5 += v34
    for fi35 := 0; fi35 < 9; fi35++ {
        _ = len(v18)
        var v36 int64 = (^(v4))
        var v37 int8 = 123
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint16 = 2
    v18 = append(v18, -1)
    var v39 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
