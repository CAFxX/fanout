package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 uint8 = 128
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: true, F1: ("hello" + ""), F2: -265}
    v3 -= v3
    if false {
        switch "a" {
        case "":
            var v8 int8 = 127
            _ = v8
        case "hello":
            switch "hello" {
            case "hello":
                v9 := []int{1, 578, -184}
                _ = v9
            case "\n\t":
                fmt.Println(v3)
            default:
                v3 -= (^(v3))
            }
        }
        if (0 < uint64(v3)) {
            v10 := GenS2{F0: v3, F1: 42}
            v3 += 177
            _ = v10
        } else {
            v11 := GenS1{F0: ((0 > (v3 - (v3 % ((128) | 1)))) && false), F1: "hello", F2: -974}
            var v12 bool = (128 <= 0)
            if (644 != 0) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    if v12 {
                        v3 = v3
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 bool = v12
                            _ = fi14
                            _ = v15
                        }
                        fmt.Println(v12)
                    } else {
                        v16 := GenS1{F0: ((3.14159 <= 1.5) || false), F1: "hello", F2: 683}
                        _ = v16
                    }
                    _ = fi13
                }
                v17 := GenS3{F0: (v3 | 247), F1: "\n\t", F2: ""}
                _ = v17
            }
            _ = v11
            _ = v12
        }
    }
    if false {
        fmt.Println(v3)
    } else {
        var v18 int32 = 1
        _ = v18
    }
    var v19 int32 = -2147483648
    fmt.Println(v19, v3)
    delete(v5, 815)
    if false {
        v20 := []int{422, 4294967295, -209}
        var v21 int64 = -718
        switch "" {
        case "":
            v22 := v2[32767]
            _ = v22
        }
        _ = v20
        _ = v21
    }
    v19 = (v19 ^ v19)
    if (!false) {
        v3 = (^(v3))
        fmt.Println(v3, v19)
    } else {
        v23 := GenS3{F0: (v3 - (v3 & v3)), F1: "\n\t", F2: "\n\t"}
        if (896 == 374) {
            var v24 uint32 = 593
            for fi25 := 0; fi25 < 20; fi25++ {
                _ = len(v4)
                _ = len(v6)
                _ = fi25
            }
            fmt.Println(v3, v24)
            _ = v24
        }
        switch 9223372036854775807 {
        case -2147483648:
            v26 := []int{32767, 980, 9223372036854775807, -9223372036854775808}
            _ = v26
        case -124:
            v27 := GenS4{F0: -750}
            _ = v27
        }
        _ = v23
    }
    v28 := []int{9223372036854775807}
    v29 := GenS5{F0: int16(v19), F1: v3, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
