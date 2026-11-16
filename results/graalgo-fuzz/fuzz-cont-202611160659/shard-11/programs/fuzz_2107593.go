package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 bool }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 255
    var v3 int16 = 42
    v4 := make(map[string]uint32)
    var v5 int64 = -1
    v6 := make(map[uint32]string)
    var v7 int8 = -1
    var v8 string = ("foo bar" + "\n\t")
    fmt.Println(v2)
    v1 = v1
    if (!false) {
        v8 = v8
        fmt.Println(v1, v2)
        switch v8 {
        case "":
            v2 -= 2
        case "\n\t":
            fmt.Println(v7, v8, v8)
        default:
            var v9 float32 = ((((1e10 - 3.14159) * 0.0) / float32((((1.5 * (0.0 + (0.1 / 1e10))) * 1.5) * ((-2.25 / 1.0) / 1.5)))) / (0.1 / (1e10 / 1.0)))
            _ = v9
        }
    }
    v10 := GenS1{F0: ("\n\t" + v8)}
    if true {
        var v11 int64 = v5
        _ = v11
    } else {
        var v12 float64 = 1.5
        _ = v12
    }
    v5 = (v5 / (((v5 * 42)) | 1))
    if ((v1 * v1) != (v1 % (((v1 << ((0) & 63))) | 1))) {
        if (!(-37 != -235)) {
            fmt.Println(v3)
            if ((v1 & (^(v1))) < (v1 - uint8(v2))) {
                v13 := []int{32768, 729, int(v2), 32767}
                _ = v13
            }
        } else {
            v5 = (v5 % (((v5 >> ((1) & 63))) | 1))
        }
        if false {
            v6[2147483647] = (((v8 + "") + (v8 + v8)) + "a")
            switch int(v1) {
            case 32768:
                if (float64(1e10) < 0.0) {
                    if true {
                        v14 := v4["foo bar"]
                        _ = v14
                    } else {
                        v15 := []int{0, 402}
                        var v16 uint64 = 4294967295
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v8, v2, v1)
                } else {
                    if true {
                        var v17 bool = true
                        v18 := GenS1{F0: ("hello" + v8)}
                        v19 := GenS1{F0: ""}
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v4[(v8 + "")] = 0
                    }
                    var v20 int64 = v5
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if true {
                            for fi22 := 0; fi22 < 18; fi22++ {
                                var v23 uint = 1
                                v24 := GenS2{F0: (fi21 | (-(fi22))), F1: v7, F2: true}
                                var v25 uint16 = v2
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        }
                        v8 = ("foo bar" + "\n\t")
                        _ = fi21
                    }
                    _ = v20
                }
            case 9223372036854775807:
                v5 = (^(v5))
            case 2147483647:
                v1 -= 128
            }
        }
    }
    v26 := []int{}
    switch int(v5) {
    case 484:
        v8 = v8
    case -9223372036854775808:
        v7 = int8(v3)
    }
    v3 -= -32768
    var v27 int64 = -598
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
