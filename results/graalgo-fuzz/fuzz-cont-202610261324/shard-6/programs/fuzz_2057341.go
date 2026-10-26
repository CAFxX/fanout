package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 string }
func genFn0() int32 {
    var v4 int32 = 736
    v5 := GenS1{F0: 255}
    v6 := []int{138, -925}
    v4 = v4
    var v7 int = 2147483647
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v4
}

func main() {
    v1 := []uint32{65535, 42, 556}
    var v2 uint8 = 0
    v3 := []int{1, 42}
    var v8 string = ("\n\t" + "")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v8)
    v2 = v2
    var v9 uint64 = 42
    v9 -= uint64(v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float64 = float64(1e10)
        v8 = v8
        _ = fi10
        _ = v11
    }
    var v12 float64 = -2.25
    if (false && true) {
        switch (("a" + v8) + v8) {
        case "\n\t":
            var v13 int = 495
            _ = v13
        case "":
            switch v8 {
            case "a":
                var v14 int32 = 42
                _ = v14
            case "foo bar":
                var v15 bool = (0 >= int32(v9))
                _ = v15
            default:
                if true {
                    if len(v3) > 0 {
                        v3[0] = -968
                    }
                    v16 := GenS2{F0: -832, F1: -7}
                    var v17 uint32 = 65535
                    _ = v16
                    _ = v17
                } else {
                    if false {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi18 := 0; fi18 < 19; fi18++ {
                                v2 = (v2 % (((v2 << ((255) & 63))) | 1))
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    switch 1 {
                                    case 869:
                                        if (1 != -1) {
                                            var v20 float32 = (3.14159 - 3.14159)
                                            v9 += 1
                                            switch -9223372036854775808 {
                                            case -1:
                                                v9 += (v9 | v9)
                                            }
                                            var v21 uint32 = 0
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v22 := GenS3{F0: 1000}
                                            _ = v22
                                        }
                                    }
                                    v23 := GenS2{F0: 0, F1: 0}
                                    v9 = (v9 & (v9 & 523))
                                    v1 = append(v1, 743)
                                    _ = fi19
                                    _ = v23
                                }
                                _ = fi18
                            }
                        }
                    }
                    fmt.Println(v9)
                }
            }
        case "foo bar":
            if (int8(v2) == 100) {
                v8 = v8
            }
        default:
            if true {
                if len(v1) > 0 {
                    v1[0] = 293
                }
                var v24 uint = 2
                v8 = v8
                _ = v24
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 int64 = 858
                    switch (-(fi25)) {
                    case 9223372036854775807:
                        switch (^(fi25)) {
                        case 42:
                            fmt.Println(v12, v12)
                        case -9223372036854775808:
                            v2 += v2
                        case 599:
                            var v27 uint = 18446744073709551615
                            _ = v27
                        default:
                            var v28 int = fi25
                            _ = v28
                        }
                    case -424:
                        if (true && (true && (true || false))) {
                            for fi29 := 0; fi29 < 11; fi29++ {
                                v30 := GenS4{F0: v9, F1: 113, F2: (v8 + (v8 + v8))}
                                _ = fi29
                                _ = v30
                            }
                        }
                    case 273:
                        var v31 uint16 = 2
                        _ = v31
                    }
                    switch fi25 {
                    case -287:
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    case 2147483647:
                        if ((v9 % ((18446744073709551615) | 1)) <= v9) {
                            fmt.Println(fi25)
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                v8 = "foo bar"
            }
        }
        v9 = 581
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
