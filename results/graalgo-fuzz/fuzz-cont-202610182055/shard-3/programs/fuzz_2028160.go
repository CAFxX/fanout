package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 int64; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0() uint16 {
    if (!false) {
        v6 := GenS1{F0: -42, F1: 819, F2: "\n\t"}
        _ = v6
    } else {
        if (!true) {
            var v7 uint16 = 767
            if true {
                for fi8 := 0; fi8 < 7; fi8++ {
                    var v9 uint8 = 0
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v9 += v9
                        v7 -= (v7 * v7)
                        _ = fi10
                    }
                    v11 := []int{(fi8 << ((65535) & 63)), 0, (-(fi8)), (fi8 % (((fi8 >> ((2) & 63))) | 1))}
                    _ = fi8
                    _ = v11
                    _ = v9
                }
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := []int{fi12}
                    if ((0.1 / -2.25) > (float64(1.5) * 1e10)) {
                        var v14 uint16 = (^(v7))
                        _ = v14
                    } else {
                        switch (fi12 | (fi12 % ((32767) | 1))) {
                        case 187:
                            v7 = 2
                        default:
                            fmt.Println(v7, fi12)
                        }
                    }
                    var v15 uint = 18446744073709551615
                    _ = fi12
                    _ = v13
                    _ = v15
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    if (0 < 0) {
                        if (0 >= 1) {
                            v7 -= v7
                        } else {
                            v17 := []int{fi16, 127, fi16}
                            _ = v17
                        }
                    } else {
                        v7 = (v7 + v7)
                        if true {
                            for fi18 := 0; fi18 < 3; fi18++ {
                                v19 := []int{-1, (fi18 & fi16)}
                                v7 = 859
                                v20 := GenS1{F0: -128, F1: v7, F2: "foo bar"}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            v21 := []int{32767, (fi16 >> ((42) & 63))}
                            for fi22 := 0; fi22 < 1; fi22++ {
                                if false {
                                    switch fi22 {
                                    case 32767:
                                        var v23 int8 = 1
                                        _ = v23
                                    case 127:
                                        var v24 string = "foo bar"
                                        _ = v24
                                    }
                                } else {
                                    fmt.Println(fi22, fi22, v7)
                                    var v25 uint16 = (^(v7))
                                    _ = v25
                                }
                                _ = fi22
                            }
                            _ = v21
                        }
                        var v26 int = (fi16 - 4294967295)
                        _ = v26
                    }
                    _ = fi16
                }
            }
            v27 := []int{4294967295}
            _ = v27
            _ = v7
        }
        var v28 uint64 = 18446744073709551615
        _ = v28
    }
    var v29 int8 = 42
    _ = v29
    return 255
}

func main() {
    var v1 uint8 = 42
    v2 := []string{("a" + "hello"), "a"}
    var v3 uint = uint(v1)
    v4 := make(map[int]string)
    var v5 string = "a"
    v30 := GenS1{F0: 100, F1: 42, F2: v5}
    var v31 int32 = -404
    if ((!true) && false) {
        var v32 uint64 = 2
        switch "\n\t" {
        case "\n\t":
            delete(v4, 707)
        case "":
            fmt.Println(v1)
        }
        v33 := GenS2{F0: 1, F1: -9223372036854775808, F2: -702}
        _ = v32
        _ = v33
    }
    var v34 uint64 = 2
    var v35 int8 = -35
    v36 := GenS3{F0: v1}
    var v37 int8 = v35
    v38 := GenS1{F0: v35, F1: 0, F2: v5}
    switch -581 {
    case -9223372036854775808:
        v5 = (v5 + v5)
    }
    fmt.Println(v3, v31)
    fmt.Println(v3, v37, v5)
    if (!(18446744073709551615 < v34)) {
        v39 := GenS2{F0: -571, F1: 0, F2: 157}
        _ = v39
    } else {
        v40 := v4[int(v35)]
        v5 = v40
        fmt.Println(v5, v40)
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
