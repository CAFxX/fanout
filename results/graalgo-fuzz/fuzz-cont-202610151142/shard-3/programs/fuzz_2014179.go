package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 string; F2 string }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 bool; F1 bool; F2 bool }
type GenS6 struct { F0 int64; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int8 = 42
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 float64 = -2.25
        if false {
            switch (fi4 | 0) {
            case 2147483647:
                v6 := GenS1{F0: (fi4 * 127)}
                _ = v6
            case 1:
                v7 := GenS1{F0: fi4}
                _ = v7
            case 127:
                if true {
                    v8 := GenS2{F0: 42, F1: ("a" + "hello"), F2: "foo bar"}
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        fmt.Println(v3, fi4)
                        fmt.Println(v5)
                        _ = fi9
                    }
                    if true {
                        v10 := GenS3{F0: (-(fi4)), F1: 2147483648, F2: 42}
                        v3 -= (v3 - v3)
                        for fi11 := 0; fi11 < 19; fi11++ {
                            var v12 uint16 = 65535
                            v13 := GenS2{F0: 42, F1: (("\n\t" + "\n\t") + "\n\t"), F2: "hello"}
                            v1[(fi4 + (fi4 / ((fi4) | 1)))] = "hello"
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        _ = v10
                    }
                }
            default:
                if (547 <= 9223372036854775807) {
                    v5 = v5
                } else {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 bool = (!true)
                        v16 := v1[fi4]
                        var v17 uint16 = 1
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v18 uint16 = 32767
                    _ = v18
                }
            }
            v3 += (v3 | -1)
        }
        v3 -= 0
        _ = fi4
        _ = v5
    }
    if true {
        var v19 uint64 = 42
        fmt.Println(v3)
        _ = v19
    } else {
        v3 = (v3 << ((856) & 63))
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS4{F0: 2, F1: 2, F2: 32767}
            delete(v1, fi20)
            v22 := []int{fi20, fi20, fi20}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 int64 = -214
    v23 = (^(v23))
    var v24 uint32 = 422
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            v24 += (v24 / ((2147483648) | 1))
            v3 -= -1
            if (!false) {
                switch "a" {
                case "":
                    if true {
                        v27 := GenS5{F0: (false && (true || false)), F1: false, F2: (!false)}
                        switch (-(fi26)) {
                        case 4294967295:
                            var v28 bool = ("" >= "")
                            _ = v28
                        default:
                            if (0.0 > 1.5) {
                                v29 := v1[fi25]
                                _ = v29
                            } else {
                                fmt.Println(v3, v23, v23)
                                var v30 uint64 = 4294967295
                                var v31 uint8 = 128
                                _ = v30
                                _ = v31
                            }
                        }
                        _ = v27
                    }
                case "foo bar":
                    if (26 < 931) {
                        v3 += (v3 | (v3 | 127))
                    }
                case "a":
                    delete(v1, (fi26 ^ -372))
                }
                if true {
                    if (false && true) {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi25, fi26)
                        case "":
                            var v32 uint32 = (^(v24))
                            _ = v32
                        case "hello":
                            v24 -= 0
                        }
                    }
                    v33 := GenS2{F0: v23, F1: "hello", F2: "\n\t"}
                    _ = v33
                } else {
                    v34 := GenS6{F0: v23, F1: v24, F2: fi26}
                    v35 := GenS3{F0: (fi26 * fi26), F1: v24, F2: 1}
                    if (!(0 <= 0)) {
                        switch "" {
                        case "":
                            var v36 int8 = -1
                            _ = v36
                        case "hello":
                            v24 = v24
                        }
                        if true {
                            if false {
                                fmt.Println(v23, v23)
                                fmt.Println(v24)
                            }
                            var v37 int32 = 42
                            _ = v37
                        } else {
                            fmt.Println(v24, v23)
                            var v38 uint8 = 1
                            v3 = (v3 + int8(fi25))
                            _ = v38
                        }
                        v3 += v3
                    }
                    v39 := GenS4{F0: 2, F1: 2, F2: 2}
                    _ = v34
                    _ = v35
                    _ = v39
                }
                fmt.Println(fi25, fi26, v24)
            } else {
                var v40 int64 = v23
                _ = v40
            }
            _ = fi26
        }
        _ = fi25
    }
    v1[127] = ("foo bar" + "foo bar")
    v41 := GenS5{F0: (false && false), F1: true, F2: ((false && true) && false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v41
}
