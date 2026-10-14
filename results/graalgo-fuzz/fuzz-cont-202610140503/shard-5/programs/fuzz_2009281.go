package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
func genFn0() int64 {
    var v4 uint = 101
    v4 = 876
    if false {
        var v5 bool = ((v4 / (((v4 >> ((4294967295) & 63))) | 1)) == 42)
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            v4 -= 0
            _ = fi6
        }
        var v7 uint8 = 42
        _ = v7
    }
    v4 += 733
    var v8 int8 = 127
    _ = v4
    _ = v8
    return 465
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    var v3 int16 = 42
    if true {
        var v9 uint64 = 248
        v10 := []int{32768, 128, 42, 32768}
        _ = v10
        _ = v9
    } else {
        v2 = 1
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "hello" {
            case "":
                if (-705 >= -2147483648) {
                    v12 := []int{(fi11 ^ (^(fi11))), fi11, (fi11 + (fi11 - (^(fi11)))), fi11}
                    for fi13 := 0; fi13 < 9; fi13++ {
                        fmt.Println(v2, v2)
                        v14 := GenS1{F0: int64(fi11), F1: (2 >= 371)}
                        _ = fi13
                        _ = v14
                    }
                    v2 = v2
                    _ = v12
                } else {
                    fmt.Println(fi11, v1)
                }
            case "a":
                for fi15 := 0; fi15 < 18; fi15++ {
                    switch fi11 {
                    case 2147483647:
                        v16 := []int{(^(fi15)), fi11, (fi15 / ((fi11) | 1))}
                        _ = v16
                    case 9223372036854775807:
                        v3 = (v3 >> ((307) & 63))
                    }
                    _ = fi15
                }
            case "foo bar":
                v1 -= 859
            default:
                v17 := []int{fi11, 32768}
                _ = v17
            }
            fmt.Println(v3, v1, v3)
            v3 += (^(v3))
            _ = fi11
        }
    }
    fmt.Println(v3, v1, v2)
    fmt.Println(v2, v2, v3)
    if (("a" >= "a") && true) {
        fmt.Println(v2, v3, v2)
    }
    if true {
        v2 += 0
    } else {
        var v18 uint16 = 42
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := []int{(fi20 & fi19), fi20}
                if (!true) {
                    var v22 int16 = (v3 ^ -413)
                    _ = v22
                }
                switch "a" {
                case "a":
                    var v23 uint = 4294967295
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float64 = 3.14159
                var v26 uint8 = 128
                fmt.Println(v2, v1)
                var v27 int8 = 1
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 int8 = 36
            _ = fi19
            _ = v28
        }
        _ = v18
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint = 99
        if false {
            v31 := GenS2{F0: uint32(v30)}
            var v32 int32 = 361
            _ = v31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    var v33 uint64 = 156
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 int16 = 1
        if false {
            if true {
                var v36 string = (("\n\t" + "foo bar") + "hello")
                _ = v36
            } else {
                v37 := GenS2{F0: 4294967295}
                switch "foo bar" {
                case "foo bar":
                    v38 := GenS3{F0: "\n\t"}
                    _ = v38
                case "hello":
                    v39 := []int{fi34, (^(fi34)), (fi34 / (((fi34 * (fi34 << ((65535) & 63)))) | 1)), fi34}
                    _ = v39
                }
                v33 -= 18446744073709551615
                _ = v37
            }
        } else {
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 uint64 = (v33 | 923)
                switch "" {
                case "hello":
                    var v42 int64 = -9223372036854775808
                    _ = v42
                case "a":
                    v43 := GenS4{F0: "hello"}
                    _ = v43
                default:
                    v44 := []int{fi34, int(v35), -9223372036854775808, fi40}
                    _ = v44
                }
                _ = fi40
                _ = v41
            }
            fmt.Println(v33)
            if (42 != 691) {
                var v45 string = "hello"
                _ = v45
            }
        }
        _ = fi34
        _ = v35
    }
    v33 += v33
    v46 := []int{-30, 670}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v46
}
