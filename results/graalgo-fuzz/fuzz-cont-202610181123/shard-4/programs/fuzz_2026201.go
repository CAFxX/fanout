package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int = 42
    v4 := make(map[string]int)
    var v5 int = -2147483648
    if true {
        if ((1 < 18446744073709551615) && (!(446 < 1))) {
            switch "hello" {
            case "hello":
                if ((v5 ^ 32767) == v3) {
                    switch "\n\t" {
                    case "":
                        if true {
                            switch 0 {
                            case 574:
                                v6 := GenS1{F0: 128, F1: 706}
                                _ = v6
                            case 872:
                                v7 := GenS1{F0: 255, F1: -983}
                                _ = v7
                            }
                            var v8 int32 = 0
                            _ = v8
                        } else {
                            fmt.Println(v3)
                            v9 := []int{v5}
                            if len(v9) > 0 {
                                v9[0] = (v5 * (v5 ^ (v3 - (-(v5)))))
                            }
                            _ = v9
                        }
                    case "hello":
                        v10 := []int{v5, (v3 | (v3 % (((^(v3))) | 1))), (v3 / ((2147483647) | 1)), (v3 | v5)}
                        _ = v10
                    case "a":
                        v5 += (v3 / (((v3 << ((0) & 63))) | 1))
                    default:
                        v11 := GenS2{F0: -158, F1: 2}
                        _ = v11
                    }
                    v12 := GenS3{F0: 42, F1: -454}
                    v13 := GenS2{F0: 42, F1: 52}
                    _ = v12
                    _ = v13
                }
            case "":
                if (-1 == -1) {
                    if false {
                        var v14 int64 = 0
                        _ = v14
                    }
                    v3 = v5
                } else {
                    v15 := GenS3{F0: -1, F1: -1}
                    _ = v15
                }
            default:
                v16 := GenS2{F0: -205, F1: 0}
                _ = v16
            }
            var v17 int = (^(v5))
            _ = v17
        } else {
            for fi18 := 0; fi18 < 19; fi18++ {
                v3 = 9223372036854775807
                _ = fi18
            }
        }
        var v19 int = (v5 % ((v5) | 1))
        var v20 int32 = 0
        _ = v19
        _ = v20
    }
    var v21 bool = false
    var v22 uint32 = 4294967295
    if true {
        v23 := GenS4{F0: -32768, F1: 28}
        _ = v23
    } else {
        switch (v5 % (((v5 ^ (v5 / ((v3) | 1)))) | 1)) {
        case 127:
            v24 := v1["\n\t"]
            _ = v24
        case 150:
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := []int{114, fi25}
                _ = fi25
                _ = v26
            }
        }
        var v27 int32 = -2147483648
        var v28 uint = 657
        _ = v27
        _ = v28
    }
    if ((v3 << ((18446744073709551615) & 63)) < v5) {
        v29 := GenS2{F0: 2147483647, F1: 127}
        _ = v29
    } else {
        var v30 float64 = 0.0
        fmt.Println(v5, v22, v5)
        _ = v30
    }
    v1[(("a" + "") + "a")] = "foo bar"
    switch "\n\t" {
    case "foo bar":
        if ((42 > 1) || (255 < 25)) {
            v3 = v5
            fmt.Println(v22, v5, v22)
            fmt.Println(v21, v21, v21)
        }
    default:
        v31 := GenS1{F0: 130, F1: -482}
        _ = v31
    }
    fmt.Println(v21, v22)
    v32 := []int{}
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(fi33, v3, v5)
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(fi33)
            v35 := GenS3{F0: 42, F1: 1000}
            _ = fi34
            _ = v35
        }
        var v36 string = "hello"
        _ = fi33
        _ = v36
    }
    var v37 string = ("a" + "a")
    for fi38 := 0; fi38 < 19; fi38++ {
        switch v37 {
        case "foo bar":
            if v21 {
                for fi39 := 0; fi39 < 5; fi39++ {
                    if false {
                        fmt.Println(v21)
                    } else {
                        v22 -= v22
                    }
                    _ = fi39
                }
            } else {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v5 -= v5
                    _ = fi40
                }
            }
        case "":
            var v41 uint = 4294967295
            _ = v41
        case "a":
            for fi42 := 0; fi42 < 7; fi42++ {
                if ((3.14159 != (-2.25 + 0.0)) && v21) {
                    delete(v4, "hello")
                    if (520 == 420) {
                        fmt.Println(fi42, v5, fi38)
                        var v43 string = v37
                        _ = v43
                    } else {
                        fmt.Println(v37)
                        if len(v32) > 0 {
                            v32[0] = -2147483648
                        }
                    }
                }
                v44 := GenS1{F0: 249, F1: -91}
                _ = fi42
                _ = v44
            }
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
