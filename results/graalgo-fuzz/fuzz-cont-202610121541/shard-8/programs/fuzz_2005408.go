package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 int }
type GenS5 struct { F0 int; F1 bool; F2 int64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint32 = 699
            if false {
                v9 = 635
                if true {
                    var v10 uint = 4294967295
                    fmt.Println(fi8, fi8)
                    if (!(!(!true))) {
                        var v11 uint = (v10 % ((v10) | 1))
                        v12 := GenS1{F0: int8(v11), F1: 32767, F2: (v9 >> (((v10 * v10)) & 63))}
                        _ = v11
                        _ = v12
                    } else {
                        var v13 uint8 = 2
                        _ = v13
                    }
                    _ = v10
                } else {
                    v9 = (v9 & v9)
                    if (!false) {
                        switch (fi7 ^ (fi7 / ((127) | 1))) {
                        case -9223372036854775808:
                            switch (fi7 ^ 148) {
                            case 32768:
                                v14 := GenS2{F0: 135}
                                _ = v14
                            case 9223372036854775807:
                                v9 -= v9
                                v9 -= (v9 & 65535)
                                v15 := GenS3{F0: -1, F1: -32768, F2: (("a" + "a") == "")}
                                _ = v15
                            }
                        case 32767:
                            var v16 uint32 = 2
                            _ = v16
                        default:
                            switch "hello" {
                            case "hello":
                                v9 = (^(v9))
                            case "\n\t":
                                var v17 float64 = (3.14159 + 0.0)
                                _ = v17
                            }
                        }
                        v9 -= 188
                    }
                }
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if (!((65535 != 32767) || true)) {
        var v18 uint32 = 930
        _ = v18
    }
    return 100
}

func main() {
    v1 := []int{42, -388}
    var v2 int64 = 685
    var v3 int = 0
    var v4 uint64 = 4294967295
    v5 := []uint32{1, 1, 2}
    var v6 uint16 = 440
    var v19 float64 = 0.1
    fmt.Println(v3, v2)
    v19 = ((3.14159 * v19) - v19)
    var v20 int32 = -1
    if true {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint8 = 203
            if (-2.25 == 1.5) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi21
            _ = v22
        }
        v5 = append(v5, 1)
    }
    var v23 float64 = v19
    switch ("\n\t" + "a") {
    case "hello":
        var v24 int16 = -32768
        _ = v24
    }
    v25 := GenS4{F0: 126, F1: true, F2: (v3 / ((-413) | 1))}
    v26 := GenS5{F0: (v3 | (v3 % ((v3) | 1))), F1: true, F2: (^(v2))}
    if false {
        var v27 int = 9223372036854775807
        var v28 uint32 = 2147483647
        v27 -= (^(v27))
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 bool = true
        v31 := GenS2{F0: 128}
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
