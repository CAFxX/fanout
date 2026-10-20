package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0() bool {
    v5 := []int{1}
    var v6 int8 = -1
    for fi7 := 0; fi7 < 9; fi7++ {
        _ = len(v5)
        _ = fi7
    }
    _ = len(v5)
    _ = v5
    _ = v6
    return (0 == -919)
}

func main() {
    v1 := []int{-791, 32768, 411}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    v4 := []int{42, -153}
    v8 := GenS1{F0: -2147483648}
    v2 = v2
    var v9 uint32 = 1
    _ = len(v3)
    v9 = 934
    var v10 bool = false
    var v11 bool = v10
    if len(v4) > 0 {
        v4[0] = -1
    }
    var v12 uint64 = 422
    for fi13 := 0; fi13 < 8; fi13++ {
        v9 = 843
        _ = fi13
    }
    switch -812 {
    case 288:
        var v14 int64 = v2
        _ = v14
    case 32768:
        switch 0 {
        case 138:
            fmt.Println(v9, v10)
        default:
            if (0.0 >= (-2.25 + 1e10)) {
                switch -628 {
                case 41:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v16 := GenS2{F0: "\n\t"}
                        _ = fi15
                        _ = v16
                    }
                case -9223372036854775808:
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 uint64 = v12
                        _ = fi17
                        _ = v18
                    }
                case -782:
                    if v11 {
                        v12 -= v12
                    } else {
                        if (!true) {
                            switch ("foo bar" + (("" + "\n\t") + "")) {
                            case "hello":
                                fmt.Println(v10, v10, v12)
                            default:
                                var v19 bool = false
                                _ = v19
                            }
                            fmt.Println(v2)
                        } else {
                            fmt.Println(v10, v2, v9)
                            if (-9223372036854775808 == -9223372036854775808) {
                                v9 += (^(v9))
                            }
                            var v20 uint64 = (v12 ^ (^(v12)))
                            _ = v20
                        }
                        v10 = ((65535 >= 1) || false)
                        switch "hello" {
                        case "":
                            delete(v3, -1)
                        case "hello":
                            for fi21 := 0; fi21 < 14; fi21++ {
                                _ = len(v3)
                                v22 := v3[(^(fi21))]
                                _ = fi21
                                _ = v22
                            }
                        case "a":
                            if false {
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    for fi24 := 0; fi24 < 7; fi24++ {
                                        fmt.Println(v9, fi23, v12)
                                        v12 = (v12 - (v12 / ((v12) | 1)))
                                        _ = fi24
                                    }
                                    fmt.Println(v10, v9)
                                    _ = fi23
                                }
                                v9 -= 2147483647
                                var v25 int16 = -681
                                _ = v25
                            }
                        }
                    }
                }
                switch 128 {
                case 535:
                    v9 = v9
                case -2147483648:
                    var v26 uint64 = 4294967295
                    _ = v26
                default:
                    var v27 int64 = (-(v2))
                    _ = v27
                }
                v12 -= 42
            } else {
                var v28 float32 = 1e10
                v29 := GenS3{F0: 2147483647, F1: v12}
                _ = v28
                _ = v29
            }
        }
    case -960:
        fmt.Println(v11, v11, v2)
    }
    var v30 int64 = (v2 / ((-576) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
    _ = v9
}
