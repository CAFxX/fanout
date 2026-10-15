package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0(p0 uint8, p1 int, p2 int16) int64 {
    v4 := []int{-2147483648, p1, p1, p1}
    v5 := GenS1{F0: 4294967295}
    var v6 float32 = float32(3.14159)
    switch (-(p1)) {
    case 852:
        var v7 float64 = 0.0
        _ = v7
    case -9223372036854775808:
        p0 += 2
    }
    if ((1 != 0) || (true && false)) {
        fmt.Println(p1, v6, p0)
        switch "" {
        case "hello":
            fmt.Println(v6, p1)
        case "":
            if (0 < 18446744073709551615) {
                for fi8 := 0; fi8 < 16; fi8++ {
                    p0 += 2
                    if true {
                        v9 := GenS1{F0: 42}
                        if ((542 < 0) && false) {
                            var v10 int8 = -128
                            var v11 int32 = 1
                            _ = v10
                            _ = v11
                        } else {
                            fmt.Println(p2, v6, p2)
                        }
                        _ = v9
                    } else {
                        p1 = p1
                        v4 = append(v4, p1)
                        p0 -= (p0 >> ((18446744073709551615) & 63))
                    }
                    if false {
                        if true {
                            v12 := GenS2{F0: 745, F1: 0, F2: -25}
                            switch (("a" + "a") + "a") {
                            case "a":
                                var v13 uint16 = 255
                                _ = v13
                            }
                            _ = v12
                        } else {
                            var v14 float32 = v6
                            v15 := GenS3{F0: 129, F1: p2}
                            fmt.Println(p1, fi8)
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS4{F0: (fi8 - (p1 ^ (fi8 + p1))), F1: true}
                        _ = v16
                    }
                    _ = fi8
                }
            }
        case "foo bar":
            var v17 float64 = ((0.1 - float64((float32(1.5) * (v6 + 0.0)))) + (3.14159 + (((0.1 * -2.25) / 3.14159) / -2.25)))
            _ = v17
        }
    } else {
        switch "foo bar" {
        case "":
            for fi18 := 0; fi18 < 19; fi18++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi18
            }
        case "foo bar":
            var v19 bool = false
            _ = v19
        }
        var v20 float64 = -2.25
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 uint64, p2 uint64) uint8 {
    switch (("a" + "a") + "\n\t") {
    case "a":
        fmt.Println(p0, p1, p0)
    case "":
        p1 += (p0 << ((18446744073709551615) & 63))
    }
    var v21 int64 = 1
    p0 -= 725
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 255
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 18446744073709551615
    v3 := []string{("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))), "a", ("" + "")}
    var v22 bool = ((1e10 / 3.14159) < 0.1)
    v2 += (v2 | (v2 | v2))
    var v23 int32 = 2147483647
    v1 = -406
    switch 32768 {
    case -2147483648:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    case -1:
        v24 := GenS5{F0: 32767, F1: (v23 * v23)}
        _ = v24
    default:
        v2 = 4294967295
    }
    var v25 int8 = -123
    fmt.Println(v25, v2)
    var v26 int32 = (v1 ^ v1)
    v2 = v2
    for fi27 := 0; fi27 < 5; fi27++ {
        if v22 {
            v25 -= v25
        } else {
            var v28 int8 = (v25 + (v25 & (v25 & v25)))
            _ = v28
        }
        _ = fi27
    }
    v29 := GenS3{F0: -2147483648, F1: 1}
    fmt.Println(v2, v26, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v29
    _ = v3
}
