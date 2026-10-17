package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 string; F1 uint8; F2 int }

func main() {
    var v1 float32 = (1e10 * ((1.5 - 0.0) / 1.0))
    var v2 int8 = 0
    v3 := make(map[int]bool)
    var v4 string = ("foo bar" + "\n\t")
    var v5 uint32 = 312
    v6 := GenS1{F0: (v4 + "foo bar")}
    for fi7 := 0; fi7 < 14; fi7++ {
        v5 += v5
        var v8 uint64 = 4294967295
        delete(v3, 4294967295)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "foo bar"}
    v10 := GenS1{F0: (((v4 + ((v4 + v4) + v4)) + v4) + "a")}
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 = v1
        _ = fi11
    }
    if (118 <= 0) {
        v12 := []int{}
        fmt.Println(v5, v1, v4)
        switch v4 {
        case "":
            switch v4 {
            case "hello":
                v13 := GenS2{F0: (0.1 == 1.5), F1: 42, F2: v4}
                _ = v13
            case "\n\t":
                v14 := GenS2{F0: (false || (255 < 2)), F1: 0, F2: ("" + v4)}
                _ = v14
            default:
                for fi15 := 0; fi15 < 14; fi15++ {
                    if (fi15 < 9223372036854775807) {
                        var v16 uint16 = 355
                        v3[42] = (125 != 19)
                        _ = v16
                    }
                    switch (fi15 ^ (fi15 % (((fi15 % (((fi15 / (((fi15 / ((fi15) | 1))) | 1))) | 1))) | 1))) {
                    case 2147483647:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case -241:
                        delete(v3, fi15)
                    }
                    _ = fi15
                }
            }
        case "\n\t":
            fmt.Println(v2, v1)
        default:
            v17 := GenS2{F0: (v4 == v4), F1: 42, F2: (v4 + "hello")}
            _ = v17
        }
        _ = v12
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    switch 128 {
    case 1:
        v19 := []int{-2147483648, -743, 1}
        _ = v19
    default:
        switch v4 {
        case "a":
            _ = len(v3)
        default:
            v20 := []int{}
            _ = v20
        }
    }
    v21 := GenS1{F0: v4}
    if (-1 < 1) {
        v1 -= v1
    } else {
        var v22 int = -1
        v23 := []int{(v22 / (((v22 ^ (v22 << ((255) & 63)))) | 1))}
        _ = v22
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 13; fi24++ {
            if (v2 < v2) {
                var v25 int16 = 42
                var v26 int16 = 0
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 int16 = int16(v5)
                    if (false && true) {
                        v29 := []int{(-(fi24)), (fi24 + fi24), fi27}
                        fmt.Println(v26, v25)
                        v30 := GenS1{F0: v4}
                        _ = v29
                        _ = v30
                    } else {
                        var v31 float32 = (v1 + 0.1)
                        v32 := v3[0]
                        _ = v31
                        _ = v32
                    }
                    var v33 uint = 255
                    _ = fi27
                    _ = v28
                    _ = v33
                }
                _ = v25
                _ = v26
            } else {
                v5 = 42
                if false {
                    v5 += 42
                    switch "hello" {
                    case "":
                        v1 -= (3.14159 + v1)
                    case "hello":
                        if (32767 > 255) {
                            v34 := []int{}
                            fmt.Println(v4)
                            _ = v34
                        } else {
                            var v35 bool = false
                            var v36 float32 = 0.0
                            _ = v35
                            _ = v36
                        }
                    default:
                        v37 := GenS3{F0: (v4 + "a"), F1: 255, F2: 9223372036854775807}
                        _ = v37
                    }
                }
            }
            _ = fi24
        }
    }
    switch "foo bar" {
    case "a":
        v38 := GenS1{F0: v4}
        _ = v38
    case "hello":
        var v39 int32 = 42
        _ = v39
    default:
        v3[42] = ((v5 >> ((4294967295) & 63)) < v5)
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v2 = (v2 / ((v2) | 1))
        v41 := v3[9223372036854775807]
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
