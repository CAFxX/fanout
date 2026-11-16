package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int32; F1 string; F2 int }
func genFn0() int8 {
    var v5 int32 = 2147483647
    if true {
        v5 = v5
        var v6 int8 = 0
        if (v5 > v5) {
            var v7 string = "hello"
            v5 = -2147483648
            switch 4294967295 {
            case 9223372036854775807:
                v6 = -1
            case -273:
                v8 := []int{}
                _ = v8
            case -917:
                if false {
                    v9 := GenS1{F0: 2}
                    switch "" {
                    case "a":
                        v5 -= (v5 & v5)
                    default:
                        var v10 uint16 = 65535
                        _ = v10
                    }
                    v11 := []int{-1, 410, -9223372036854775808, 127}
                    _ = v11
                    _ = v9
                }
            default:
                v6 = v6
            }
            _ = v7
        } else {
            v5 -= v5
        }
        _ = v6
    }
    _ = v5
    return -1
}
func genFn1() uint32 {
    var v12 string = "a"
    var v13 bool = false
    if v13 {
        if true {
            var v14 bool = (2147483648 >= 649)
            _ = v14
        }
        if false {
            if v13 {
                fmt.Println(v12)
                v15 := []int{623}
                v13 = v13
                _ = v15
            }
            var v16 uint32 = 2147483647
            var v17 uint16 = 42
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 float32 = 0.1
                    switch (fi18 + fi18) {
                    case -30:
                        v21 := GenS2{F0: 655, F1: -9223372036854775808, F2: 2147483648}
                        _ = v21
                    case 268:
                        if true {
                            v22 := []int{4294967295, (fi19 | fi18), 652}
                            var v23 int8 = -80
                            v24 := GenS1{F0: 613}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    case 0:
                        switch 324 {
                        case 631:
                            v13 = false
                        }
                        v12 = v12
                    default:
                        var v25 int16 = 802
                        _ = v25
                    }
                    var v26 string = ((("" + v12) + v12) + v12)
                    _ = fi19
                    _ = v20
                    _ = v26
                }
                var v27 int = fi18
                _ = fi18
                _ = v27
            }
        }
    }
    v12 = (v12 + ((v12 + (v12 + ("foo bar" + (v12 + v12)))) + ""))
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 int = 147
    for fi28 := 0; fi28 < 18; fi28++ {
        fmt.Println(v1, v1, v4)
        for fi29 := 0; fi29 < 1; fi29++ {
            if (genFn0() == -128) {
                _ = len(v3)
            }
            if ((v1 * (v1 / ((v1) | 1))) >= -9223372036854775808) {
                v30 := []int{(fi29 / (((v4 / (((v4 >> ((18446744073709551615) & 63))) | 1))) | 1))}
                var v31 uint32 = 4294967295
                v32 := GenS2{F0: 32767, F1: v1, F2: (^(v31))}
                _ = v30
                _ = v31
                _ = v32
            } else {
                v4 += (fi29 + fi28)
                v1 = v1
                v33 := GenS3{F0: ((("a" + "\n\t") + "\n\t") + "\n\t"), F1: 223}
                _ = v33
            }
            _ = len(v3)
            fmt.Println(v1)
            v4 += (-(v4))
            _ = fi29
        }
        _ = fi28
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS4{F0: 42, F1: true}
        switch fi34 {
        case 404:
            v1 -= v1
        case -9223372036854775808:
            fmt.Println(v1, v1, v4)
        default:
            if false {
                var v36 uint16 = 42
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 uint = 682
            v38 = v38
            _ = fi37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    v39 := []int{47, v4, (v4 + v4), 32767}
    if true {
        fmt.Println(v1)
    } else {
        v4 -= -1
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 int = -9223372036854775808
            var v42 bool = ((v41 % ((128) | 1)) >= (v41 & 1))
            _ = fi40
            _ = v41
            _ = v42
        }
        switch v4 {
        case 795:
            var v43 uint8 = 42
            _ = v43
        case -708:
            var v44 uint16 = 255
            _ = v44
        default:
            var v45 int = v4
            _ = v45
        }
    }
    fmt.Println(v1, v1)
    v46 := GenS2{F0: -1, F1: 42, F2: 42}
    fmt.Println(v1, v4)
    _ = len(v39)
    fmt.Println(v1)
    v47 := GenS5{F0: -1, F1: ("" + "foo bar"), F2: (-(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
}
