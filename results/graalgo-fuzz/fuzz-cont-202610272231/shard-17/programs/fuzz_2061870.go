package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 int64 }
func genFn0(p0 int8, p1 int) int8 {
    var v7 uint = 42
    if true {
        var v8 float32 = 0.1
        v9 := []int{}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: p0, F1: (v7 ^ (v7 << (((^(v7))) & 63))), F2: (^(v7))}
    v11 := []int{}
    v12 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    v1 := []int{32768}
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 string = "a"
    v5 := make(map[int]string)
    v6 := make(map[int]string)
    var v13 int16 = 1
    var v14 uint8 = 57
    var v15 uint64 = 0
    switch (v4 + (("" + (v4 + "foo bar")) + "foo bar")) {
    case "foo bar":
        var v16 int64 = -46
        _ = v16
    case "":
        fmt.Println(v4)
    default:
        _ = len(v2)
    }
    var v17 float32 = -2.25
    if (65535 < 4294967295) {
        if (32767 != (v13 << ((668) & 63))) {
            v14 = (^(v14))
            for fi18 := 0; fi18 < 8; fi18++ {
                switch (v4 + v4) {
                case "a":
                    var v19 uint32 = 4294967295
                    _ = v19
                }
                for fi20 := 0; fi20 < 4; fi20++ {
                    if true {
                        fmt.Println(fi18, v13, fi18)
                        fmt.Println(v4, v4)
                        v1 = append(v1, (fi20 >> ((65535) & 63)))
                    } else {
                        var v21 int16 = (v13 % (((v13 >> ((770) & 63))) | 1))
                        switch fi20 {
                        case 9223372036854775807:
                            switch -78 {
                            case 42:
                                fmt.Println(v15, v14)
                                v21 = v13
                                var v22 bool = true
                                var v23 int64 = 9223372036854775807
                                v23 = v23
                                _ = v22
                                _ = v23
                            case 2147483647:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case -9223372036854775808:
                            var v24 int64 = 944
                            v25 := GenS2{F0: (fi18 + 9223372036854775807)}
                            if (!((!true) && (-1 <= 42))) {
                                fmt.Println(fi18)
                            } else {
                                var v26 uint = 902
                                fmt.Println(v26, v17)
                                if false {
                                    var v27 float32 = 1e10
                                    _ = v27
                                }
                                _ = v26
                            }
                            _ = v24
                            _ = v25
                        case 151:
                            v13 = v13
                        default:
                            v28 := GenS3{F0: 100, F1: -2147483648, F2: true}
                            _ = v28
                        }
                        _ = v21
                    }
                    v13 += (v13 | v13)
                    _ = fi20
                }
                if false {
                    v14 = (^(v14))
                    var v29 uint = uint(v15)
                    _ = v29
                }
                _ = fi18
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                for fi31 := 0; fi31 < 14; fi31++ {
                    fmt.Println(fi30, v14)
                    _ = len(v3)
                    _ = fi31
                }
                _ = fi30
            }
        }
        fmt.Println(v4, v15)
    } else {
        v4 = v4
        delete(v6, -640)
    }
    var v32 uint64 = v15
    v32 = 0
    fmt.Println(v32, v15)
    _ = len(v1)
    fmt.Println(v15, v14, v14)
    for fi33 := 0; fi33 < 15; fi33++ {
        if (2 != 0) {
            switch fi33 {
            case 42:
                for fi34 := 0; fi34 < 18; fi34++ {
                    v35 := GenS4{F0: (fi33 * (fi34 ^ fi33)), F1: -1}
                    fmt.Println(fi34, v13, fi34)
                    _ = fi34
                    _ = v35
                }
            case -21:
                _ = len(v5)
            default:
                v36 := GenS1{F0: 42, F1: 0, F2: 981}
                _ = v36
            }
            v15 -= v32
        } else {
            if len(v1) > 0 {
                v1[0] = fi33
            }
        }
        v14 -= 42
        _ = fi33
    }
    v1 = append(v1, 1)
    fmt.Println(v17)
    var v37 uint32 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
