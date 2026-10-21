package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint8) uint64 {
    v5 := GenS1{F0: 1, F1: 42}
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            if (((0.0 * 0.1) + (-2.25 * 1e10)) < float32((0.0 + 1.5))) {
                v7 := GenS2{F0: 32767}
                _ = v7
            }
            if false {
                for fi8 := 0; fi8 < 15; fi8++ {
                    switch fi8 {
                    case 0:
                        var v9 int8 = 1
                        _ = v9
                    default:
                        var v10 uint = 907
                        _ = v10
                    }
                    v11 := []int{fi6, (fi6 % ((-467) | 1))}
                    v12 := GenS1{F0: 279, F1: 0}
                    _ = fi8
                    _ = v11
                    _ = v12
                }
                fmt.Println(p0, fi6)
            }
        }
        _ = fi6
    }
    var v13 bool = (0 <= 0)
    v14 := []int{-2147483648, -2147483648}
    _ = p0
    _ = v13
    _ = v14
    _ = v5
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    v15 := GenS2{F0: 42}
    switch "hello" {
    case "a":
        v1 += 0
    case "":
        if (!true) {
            v16 := []int{32768, 4294967295}
            _ = v16
        }
    case "\n\t":
        v1 = v1
    default:
        v1 = v1
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: 255}
        _ = fi17
        _ = v18
    }
    var v19 uint16 = 255
    v3[1] = 2147483647
    v1 = v1
    if true {
        v20 := GenS2{F0: (v19 % ((0) | 1))}
        fmt.Println(v19, v19)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := []int{fi21}
            fmt.Println(v1, v1, v1)
            for fi23 := 0; fi23 < 9; fi23++ {
                if len(v22) > 0 {
                    v22[0] = fi23
                }
                var v24 uint32 = 638
                if (1000 <= 0) {
                    if ("foo bar" <= "a") {
                        if (1 < 1000) {
                            if (!((v19 / (((v19 >> ((18446744073709551615) & 63))) | 1)) > v19)) {
                                var v25 int = (fi21 >> ((2) & 63))
                                v1 += -128
                                v26 := GenS1{F0: 42, F1: 1}
                                _ = v25
                                _ = v26
                            }
                            fmt.Println(fi23, v1, v19)
                            fmt.Println(fi21, v19, v19)
                        } else {
                            fmt.Println(v19, v1)
                            v27 := GenS2{F0: (v19 * 521)}
                            _ = v27
                        }
                        fmt.Println(fi21, fi23, v19)
                    }
                    fmt.Println(fi23, v24)
                } else {
                    if (1 != 1000) {
                        v28 := GenS1{F0: 42, F1: 295}
                        var v29 uint8 = 80
                        v24 = (v24 ^ v24)
                        _ = v28
                        _ = v29
                    } else {
                        v1 += (v1 % ((0) | 1))
                        var v30 string = ("\n\t" + ("foo bar" + ("hello" + "hello")))
                        v24 -= 381
                        _ = v30
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        if (1 >= 2147483647) {
            var v31 int8 = v1
            switch -946 {
            case -1:
                if (false && false) {
                    v1 += 1
                    v19 -= v19
                    v19 += 42
                }
            case 729:
                v32 := GenS1{F0: 18446744073709551615, F1: 2}
                _ = v32
            case -2147483648:
                var v33 string = "foo bar"
                _ = v33
            default:
                switch ("a" + (("\n\t" + "\n\t") + ("hello" + ("" + "\n\t")))) {
                case "a":
                    v31 += v1
                }
            }
            v34 := GenS3{F0: 366, F1: genFn0(83)}
            _ = v31
            _ = v34
        } else {
            v19 += v19
        }
    }
    v19 -= 0
    switch -2147483648 {
    case -311:
        switch -1 {
        case 85:
            v1 -= (v1 ^ (v1 % ((v1) | 1)))
        case 32767:
            var v35 uint = 0
            _ = v35
        default:
            var v36 string = ""
            _ = v36
        }
    case 9223372036854775807:
        v37 := GenS3{F0: 4294967295, F1: 4294967295}
        _ = v37
    case -9223372036854775808:
        v1 += 1
    }
    v38 := GenS2{F0: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
