package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int16; F1 uint; F2 int8 }

func main() {
    v1 := []int{32768, -899, -18, 32768}
    var v2 int32 = -871
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 2; fi4++ {
        switch (-(fi4)) {
        case 32767:
            for fi5 := 0; fi5 < 5; fi5++ {
                v6 := v3[2147483647]
                if (true || (0.1 >= (3.14159 * (3.14159 / 1.0)))) {
                    if (true && (1 != 1)) {
                        switch "\n\t" {
                        case "":
                            v1 = append(v1, -9223372036854775808)
                        }
                        if false {
                            fmt.Println(v6, fi4, fi5)
                            if false {
                                v7 := GenS1{F0: (v6 + v6), F1: 42}
                                v8 := GenS2{F0: true, F1: 445, F2: 170}
                                _ = v7
                                _ = v8
                            } else {
                                var v9 uint = 2
                                v2 -= 42
                                var v10 uint64 = 18446744073709551615
                                _ = v10
                                _ = v9
                            }
                            v2 -= (v2 >> ((4294967295) & 63))
                        }
                        v6 = (v6 + "a")
                    }
                }
                if true {
                    _ = len(v1)
                    var v11 int = (fi5 >> ((705) & 63))
                    _ = v11
                } else {
                    switch "a" {
                    case "":
                        v2 += 2147483647
                    default:
                        v2 = (v2 % (((v2 * -1)) | 1))
                    }
                    fmt.Println(fi4, v2, v6)
                }
                _ = fi5
                _ = v6
            }
        case 32768:
            for fi12 := 0; fi12 < 10; fi12++ {
                switch "" {
                case "hello":
                    if false {
                        switch fi4 {
                        case 0:
                            if ((v2 >> ((4294967295) & 63)) == (v2 & (v2 & -196))) {
                                v2 += v2
                                var v13 uint = 211
                                _ = v13
                            }
                            v14 := GenS1{F0: ("foo bar" + "a"), F1: 140}
                            _ = v14
                        }
                    }
                case "foo bar":
                    var v15 int32 = v2
                    _ = v15
                case "\n\t":
                    var v16 uint = 4294967295
                    _ = v16
                }
                v2 = 2147483647
                if len(v1) > 0 {
                    v1[0] = (^(fi12))
                }
                _ = fi12
            }
        case 9223372036854775807:
            if (902 > 24) {
                v2 = (v2 / ((-1) | 1))
                v17 := GenS3{F0: 509}
                if true {
                    if (1 < 542) {
                        v2 = (v2 + 1)
                        for fi18 := 0; fi18 < 7; fi18++ {
                            if true {
                                var v19 int = (-(fi18))
                                _ = v19
                            }
                            v1 = append(v1, -2147483648)
                            fmt.Println(fi4)
                            _ = fi18
                        }
                        if len(v1) > 0 {
                            v1[0] = fi4
                        }
                    }
                    var v20 uint32 = 2
                    v21 := GenS4{F0: -9223372036854775808, F1: 65535}
                    _ = v20
                    _ = v21
                }
                _ = v17
            } else {
                v2 += (-(v2))
                v2 -= -1
            }
        }
        v2 -= (v2 ^ int32(fi4))
        var v22 bool = false
        _ = fi4
        _ = v22
    }
    v23 := GenS2{F0: (-1 <= 32767), F1: 1, F2: -1}
    v24 := v3[1]
    var v25 uint = 340
    v25 = 65535
    var v26 float64 = -2.25
    var v27 int16 = int16(v25)
    _ = len(v1)
    switch (v24 + v24) {
    case "":
        v27 = v27
    case "\n\t":
        var v28 int8 = -128
        _ = v28
    }
    v24 = ((v24 + (v24 + "\n\t")) + (v24 + (v24 + (v24 + v24))))
    v29 := GenS5{F0: v27, F1: v25, F2: -42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
