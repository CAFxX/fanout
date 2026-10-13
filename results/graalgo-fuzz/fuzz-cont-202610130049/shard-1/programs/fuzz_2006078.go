package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    v1 := []int{42}
    v2 := make(map[int]int)
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]string)
    var v5 int16 = 129
    var v6 float64 = (1.5 - 3.14159)
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v7 uint64 = 4294967295
        var v8 uint = (v3 ^ v3)
        var v9 uint = v3
        _ = v7
        _ = v8
        _ = v9
    }
    v6 += v6
    switch 42 {
    case 42:
        switch "a" {
        case "hello":
            var v10 uint = v3
            _ = v10
        case "":
            v11 := GenS1{F0: true, F1: true}
            _ = v11
        default:
            if true {
                var v12 float32 = ((0.0 + 1e10) * float32((v6 + float64(((1.5 / 1e10) + 0.0)))))
                switch "" {
                case "":
                    var v13 int8 = -72
                    _ = v13
                case "foo bar":
                    if true {
                        v6 = v6
                        v12 -= v12
                    } else {
                        fmt.Println(v12)
                        var v14 uint8 = 127
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v14
                    }
                case "\n\t":
                    v6 = v6
                default:
                    v5 = 32767
                }
                _ = v12
            } else {
                v5 = (v5 % ((v5) | 1))
                var v15 uint16 = 32767
                var v16 int8 = 93
                _ = v15
                _ = v16
            }
        }
    default:
        if (false && (!(true && (v3 <= (v3 / ((v3) | 1)))))) {
            var v17 uint16 = 1
            var v18 int16 = (v5 ^ (v5 - (v5 + int16(v17))))
            _ = v17
            _ = v18
        } else {
            var v19 int64 = 9223372036854775807
            _ = v19
        }
    }
    if (!(32767 >= 32767)) {
        v5 += 0
        var v20 uint32 = 42
        var v21 uint = (v3 ^ (^(v3)))
        _ = v20
        _ = v21
    } else {
        if ((v6 < float64(0.1)) && (222 >= 0)) {
            for fi22 := 0; fi22 < 19; fi22++ {
                v6 -= -2.25
                v23 := GenS1{F0: (((42 != 234) || (528 > 0)) && false), F1: (4294967295 == 42)}
                v24 := GenS1{F0: true, F1: false}
                _ = fi22
                _ = v23
                _ = v24
            }
            switch "foo bar" {
            case "\n\t":
                for fi25 := 0; fi25 < 19; fi25++ {
                    v6 = v6
                    _ = fi25
                }
            case "":
                var v26 int32 = -2147483648
                _ = v26
            }
            v6 -= v6
        }
        var v27 bool = (9223372036854775807 < -927)
        var v28 int16 = v5
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 uint8 = 255
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 int16 = v5
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4[2147483647] = ""
    v6 += -2.25
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
