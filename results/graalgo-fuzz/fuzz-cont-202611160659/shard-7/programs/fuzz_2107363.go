package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    v3 := []string{""}
    var v4 int8 = 8
    var v5 uint = 255
    v6 := make(map[string]string)
    var v7 string = ""
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(fi8)
        fmt.Println(v7, v5)
        var v9 uint = v5
        _ = fi8
        _ = v9
    }
    if false {
        var v10 float32 = 3.14159
        fmt.Println(v4, v4)
        _ = len(v6)
        _ = v10
    }
    var v11 int = 127
    v7 = v7
    var v12 int64 = 0
    if ((("a" + "\n\t") + v7) <= v7) {
        v13 := GenS1{F0: (v7 + v7), F1: (v5 / ((625) | 1)), F2: 42}
        _ = v13
    }
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 19; fi15++ {
        switch fi15 {
        case -115:
            v16 := GenS2{F0: -165, F1: 0, F2: 235}
            _ = v16
        case -9223372036854775808:
            if true {
                var v17 string = v7
                v18 := GenS1{F0: v7, F1: v5, F2: 263}
                _ = v17
                _ = v18
            }
        case 9223372036854775807:
            if (!true) {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v12 += (v12 / (((v12 / ((v12) | 1))) | 1))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi19
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    if true {
                        _ = len(v2)
                        if (!true) {
                            var v21 uint64 = 281
                            v22 := GenS3{F0: -79}
                            var v23 int16 = 0
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS1{F0: v7, F1: 346, F2: 18446744073709551615}
                        _ = v24
                    }
                    if (127 == 42) {
                        v5 += 452
                        switch (v11 * fi15) {
                        case 0:
                            v25 := GenS1{F0: ((v7 + v7) + v7), F1: (v5 ^ (v5 & (v5 | (^(v5))))), F2: 868}
                            _ = v25
                        case -6:
                            v26 := GenS4{F0: (v12 / ((v12) | 1))}
                            _ = v26
                        case -9223372036854775808:
                            v3 = append(v3, v7)
                        }
                    }
                    _ = fi20
                }
            } else {
                var v27 int16 = 0
                _ = v27
            }
        default:
            var v28 float64 = 1.5
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            if false {
                var v30 uint32 = 42
                var v31 uint16 = (^(v14))
                v5 += 0
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + ("a" + v7))
                }
                _ = v30
                _ = v31
            }
            if true {
                fmt.Println(v7, v12, v7)
            }
            _ = fi29
        }
        var v32 uint64 = 604
        _ = fi15
        _ = v32
    }
    v33 := GenS4{F0: v12}
    switch v7 {
    case "hello":
        v14 = (v14 % ((v14) | 1))
    case "foo bar":
        var v34 float64 = 3.14159
        _ = v34
    case "a":
        v35 := GenS1{F0: v7, F1: (^(v5)), F2: 18446744073709551615}
        _ = v35
    default:
        var v36 string = v7
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
