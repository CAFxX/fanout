package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 13; fi4++ {
        if ((-(fi4)) == 0) {
            var v5 uint8 = 2
            _ = v5
        } else {
            if (42 == 42) {
                var v6 uint16 = 2
                v6 += v6
                _ = v6
            }
            fmt.Println(fi4)
        }
        _ = fi4
    }
    var v7 uint64 = 4294967295
    _ = v7
    return (v7 / (((v7 ^ 42)) | 1))
}
func genFn1() int16 {
    if true {
        switch -779 {
        case 4294967295:
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: 42, F1: -1}
                _ = fi8
                _ = v9
            }
        case -9223372036854775808:
            var v10 int = 4294967295
            _ = v10
        case -1:
            if (false || false) {
                var v11 uint16 = 42
                _ = v11
            } else {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 int = fi12
                    var v14 int16 = -112
                    var v15 int32 = -684
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint32 = 526
                    switch "" {
                    case "\n\t":
                        if true {
                            if (false || (42 > -1)) {
                                v17 = v17
                            } else {
                                v18 := GenS2{F0: fi16}
                                var v19 uint32 = (^(v17))
                                v20 := []int{fi16, fi16, 128, fi16}
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        } else {
                            v21 := []int{fi16, (-(fi16))}
                            var v22 uint8 = 128
                            _ = v21
                            _ = v22
                        }
                    }
                    v17 -= (v17 + (^(v17)))
                    _ = fi16
                    _ = v17
                }
                if true {
                    v23 := []int{-2147483648, 127, 42}
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    var v24 int16 = 0
                    _ = v23
                    _ = v24
                }
            }
        }
        var v25 bool = false
        for fi26 := 0; fi26 < 8; fi26++ {
            v25 = v25
            fmt.Println(v25)
            _ = fi26
        }
        _ = v25
    }
    var v27 uint32 = 862
    _ = v27
    return 42
}

func main() {
    v1 := []int{9223372036854775807, 127}
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 uint64 = 4294967295
        for fi30 := 0; fi30 < 6; fi30++ {
            v3[1] = false
            v31 := GenS2{F0: fi30}
            if true {
                v2 += v2
                v29 = v29
                switch fi28 {
                case 128:
                    fmt.Println(fi28, v29)
                }
            }
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    var v32 string = ("\n\t" + "a")
    var v33 uint = 601
    var v34 float32 = 1e10
    for fi35 := 0; fi35 < 4; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            fmt.Println(fi35, v34, v33)
            _ = fi36
        }
        v37 := GenS2{F0: 127}
        var v38 uint8 = 177
        _ = fi35
        _ = v37
        _ = v38
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v39 int8 = 64
    var v40 int64 = -129
    for fi41 := 0; fi41 < 7; fi41++ {
        v33 -= (v33 / (((v33 % ((18446744073709551615) | 1))) | 1))
        v34 = v34
        _ = fi41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        v32 = ("foo bar" + v32)
        _ = fi42
    }
    v32 = v32
    switch 32768 {
    case 127:
        var v43 uint32 = 0
        _ = v43
    default:
        fmt.Println(v32, v2, v33)
    }
    v33 = (v33 / (((v33 | v33)) | 1))
    var v44 int = 128
    for fi45 := 0; fi45 < 9; fi45++ {
        v3[172] = (42 == 255)
        _ = fi45
    }
    switch ("hello" + "") {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        v46 := GenS3{F0: 398}
        _ = v46
    case "hello":
        v34 = v34
    default:
        var v47 int8 = v39
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v40
    _ = v44
}
