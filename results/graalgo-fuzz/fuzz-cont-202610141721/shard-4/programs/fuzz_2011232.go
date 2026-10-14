package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint; F1 uint; F2 int16 }

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 uint32 = (v1 * v1)
    var v5 int64 = -996
    var v6 int32 = 2147483647
    if true {
        v7 := GenS1{F0: v5, F1: -258}
        fmt.Println(v6, v2, v5)
        fmt.Println(v5)
        _ = v7
    } else {
        v5 = v5
        switch 127 {
        case 9223372036854775807:
            var v8 uint8 = 42
            _ = v8
        case -2147483648:
            v9 := []int{int(v1), -291}
            _ = v9
        case -1000:
            v10 := GenS1{F0: -1, F1: 0}
            _ = v10
        default:
            fmt.Println(v1, v1)
        }
        if false {
            v11 := GenS1{F0: v5, F1: v6}
            switch (v2 + v2) {
            case "hello":
                v12 := GenS1{F0: v5, F1: v6}
                _ = v12
            }
            fmt.Println(v1)
            _ = v11
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                if (false && (455 == 2)) {
                    v14 := GenS2{F0: v5, F1: 224}
                    if (true && true) {
                        switch -9223372036854775808 {
                        case -146:
                            fmt.Println(v2)
                        default:
                            v15 := []int{(fi13 - (fi13 & (fi13 & (fi13 * fi13)))), (fi13 | (fi13 + 4294967295))}
                            _ = v15
                        }
                        var v16 uint64 = 808
                        _ = v16
                    } else {
                        var v17 int32 = (v6 % (((v6 & v6)) | 1))
                        v18 := []int{}
                        for fi19 := 0; fi19 < 14; fi19++ {
                            v20 := GenS3{F0: v6, F1: true}
                            _ = fi19
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v22 := GenS3{F0: (v6 / ((v6) | 1)), F1: (922 >= 629)}
                        _ = fi21
                        _ = v22
                    }
                    _ = v14
                } else {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := []int{(fi23 + fi13)}
                        _ = fi23
                        _ = v24
                    }
                    fmt.Println(v6, v1, v6)
                    v3[(v4 << ((287) & 63))] = (v1 % ((v1) | 1))
                }
                _ = fi13
            }
            switch v2 {
            case "":
                delete(v3, v1)
            case "\n\t":
                v25 := GenS4{F0: 42, F1: 107, F2: -471}
                _ = v25
            case "a":
                v26 := GenS3{F0: v6, F1: true}
                _ = v26
            }
        }
    }
    v27 := GenS2{F0: 1, F1: 118}
    if ((true && true) && false) {
        switch ((("foo bar" + v2) + "\n\t") + (v2 + v2)) {
        case "\n\t":
            v2 = v2
        }
    }
    v6 -= v6
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 string = (v2 + v2)
            switch (fi29 / (((^(fi29))) | 1)) {
            case 4294967295:
                v31 := v3[(v1 | 749)]
                _ = v31
            default:
                v32 := GenS2{F0: (-(v5)), F1: 121}
                _ = v32
            }
            v4 = v4
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v33 := []int{329, 1}
    var v34 float64 = float64(0.0)
    var v35 uint = 543
    fmt.Println(v4, v2, v4)
    v34 -= ((((1.5 * v34) / float64(1.5)) + v34) + v34)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
