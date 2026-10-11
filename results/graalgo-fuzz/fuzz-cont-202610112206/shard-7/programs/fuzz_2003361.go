package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64; F2 int }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 bool; F1 int8; F2 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint16 = 0
    var v2 bool = false
    var v3 int32 = -2147483648
    var v4 uint64 = 18446744073709551615
    var v5 int8 = -46
    switch "" {
    case "\n\t":
        v6 := []int{9223372036854775807, 931, -689, -829}
        _ = v6
    case "foo bar":
        v7 := []int{130, 0}
        _ = v7
    }
    if (1 == uint64(v1)) {
        switch 136 {
        case 582:
            var v8 uint16 = (v1 ^ v1)
            _ = v8
        default:
            var v9 float32 = 1.5
            _ = v9
        }
    } else {
        v5 += 100
    }
    v10 := GenS1{F0: 1}
    if v2 {
        v11 := GenS2{F0: 128, F1: 0, F2: 26}
        _ = v11
    }
    v1 -= 498
    var v12 uint16 = v1
    if v2 {
        v13 := []int{}
        fmt.Println(v2)
        v14 := GenS3{F0: (v4 ^ v4), F1: -9223372036854775808, F2: 2}
        _ = v13
        _ = v14
    } else {
        v3 = v3
        switch "" {
        case "foo bar":
            v5 = v5
        case "a":
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := []int{0}
                _ = fi15
                _ = v16
            }
        case "\n\t":
            var v17 string = "hello"
            _ = v17
        default:
            for fi18 := 0; fi18 < 15; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    v12 = (v1 | v12)
                    v20 := GenS4{F0: (-514 > -930), F1: (v5 - (v5 << ((2) & 63))), F2: 132}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
        v3 -= int32(v1)
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint16 = v12
        if (v2 && (!false)) {
            for fi23 := 0; fi23 < 4; fi23++ {
                switch (fi23 / ((fi21) | 1)) {
                case 259:
                    fmt.Println(v3, v3)
                case 723:
                    var v24 int8 = (v5 | v5)
                    _ = v24
                default:
                    v12 += 1
                }
                _ = fi23
            }
        } else {
            var v25 uint16 = v1
            _ = v25
        }
        v26 := GenS5{F0: -910, F1: (-(v5))}
        _ = fi21
        _ = v22
        _ = v26
    }
    if v2 {
        v27 := GenS5{F0: 1000, F1: -68}
        switch "" {
        case "foo bar":
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 int8 = int8(v3)
                switch "" {
                case "\n\t":
                    if v2 {
                        fmt.Println(v29, v29, v3)
                    }
                default:
                    v30 := []int{(fi28 ^ 642), (fi28 & (fi28 | (fi28 ^ fi28))), (fi28 % ((2147483647) | 1)), fi28}
                    _ = v30
                }
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := []int{fi31, -20, 9223372036854775807}
                    switch "a" {
                    case "hello":
                        switch ((("\n\t" + "a") + "\n\t") + "") {
                        case "":
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        case "hello":
                            if len(v32) > 0 {
                                v32[0] = fi28
                            }
                        case "foo bar":
                            v33 := GenS6{F0: 2147483647}
                            v29 += 59
                            _ = v33
                        default:
                            var v34 string = "\n\t"
                            _ = v34
                        }
                        fmt.Println(v12)
                    }
                    _ = fi31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
        case "a":
            if (!false) {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := []int{fi35, fi35, (fi35 / ((fi35) | 1))}
                    v3 = -996
                    var v37 bool = (!(("foo bar" + "hello") < "a"))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                fmt.Println(v3, v5)
            } else {
                v38 := []int{2147483647, 42, -933, 9223372036854775807}
                _ = v38
            }
        case "hello":
            fmt.Println(v5, v2)
        default:
            var v39 int16 = -32
            _ = v39
        }
        v40 := GenS2{F0: 1, F1: -352, F2: 484}
        _ = v27
        _ = v40
    } else {
        fmt.Println(v2)
        v41 := GenS3{F0: v4, F1: int64(v5), F2: 255}
        _ = v41
    }
    var v42 float32 = (0.1 / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
