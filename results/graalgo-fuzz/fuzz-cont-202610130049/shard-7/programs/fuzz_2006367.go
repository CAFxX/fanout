package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int8; F2 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int = 2147483647
    for fi4 := 0; fi4 < 14; fi4++ {
        v3 += 128
        for fi5 := 0; fi5 < 5; fi5++ {
            if true {
                var v6 int8 = -72
                v7 := []int{}
                _ = v6
                _ = v7
            }
            var v8 uint = 18446744073709551615
            _ = fi5
            _ = v8
        }
        switch (v3 >> ((65535) & 63)) {
        case 42:
            var v9 uint32 = 4294967295
            _ = v9
        case 32767:
            switch "\n\t" {
            case "foo bar":
                fmt.Println(fi4, v3)
            }
        case -874:
            v10 := GenS1{F0: 32767, F1: 488}
            _ = v10
        default:
            switch "foo bar" {
            case "\n\t":
                var v11 int32 = 1
                _ = v11
            }
        }
        _ = fi4
    }
    if true {
        v12 := GenS1{F0: 42, F1: 4294967295}
        var v13 uint8 = 24
        fmt.Println(v3)
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 42, F1: 0, F2: 42}
    var v15 uint16 = 42
    if true {
        v3 = 127
        v16 := GenS3{F0: 0}
        _ = v16
    }
    fmt.Println(v15, v15)
    var v17 string = ("hello" + "\n\t")
    if (!true) {
        var v18 int32 = -2147483648
        v19 := []int{}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v3)
            switch v17 {
            case "":
                switch (-(fi21)) {
                case 32767:
                    v15 += v15
                case 127:
                    if true {
                        var v22 int8 = 42
                        switch ("hello" + ("" + v17)) {
                        case "\n\t":
                            v15 = v15
                        case "foo bar":
                            _ = len(v2)
                        case "":
                            var v23 int64 = 42
                            var v24 int32 = 318
                            _ = v23
                            _ = v24
                        default:
                            fmt.Println(fi20)
                        }
                        v3 = fi20
                        _ = v22
                    } else {
                        if (v15 <= (v15 + (v15 - 979))) {
                            v25 := []int{v3}
                            v26 := GenS4{F0: false, F1: 42}
                            _ = v25
                            _ = v26
                        } else {
                            switch v17 {
                            case "a":
                                var v27 int8 = -33
                                _ = v27
                            }
                            switch v17 {
                            case "hello":
                                v28 := GenS5{F0: -128, F1: (9223372036854775807 != -660)}
                                _ = v28
                            case "":
                                var v29 float32 = ((0.1 + 3.14159) * 3.14159)
                                _ = v29
                            }
                            v17 = (((v17 + "foo bar") + v17) + v17)
                        }
                        switch v17 {
                        case "hello":
                            v3 += (-(fi20))
                        default:
                            v17 = v17
                        }
                        v30 := GenS3{F0: 4294967295}
                        var v31 uint8 = 1
                        switch 42 {
                        case -9223372036854775808:
                            v32 := GenS6{F0: v15, F1: (v31 ^ 1), F2: 42}
                            _ = v32
                        case 127:
                            v33 := GenS7{F0: 0, F1: 1, F2: 4294967295}
                            _ = v33
                        case 128:
                            var v34 int8 = int8(fi21)
                            _ = v34
                        default:
                            delete(v1, (fi20 + 4294967295))
                        }
                        _ = v30
                        _ = v31
                    }
                case 327:
                    var v35 float64 = 3.14159
                    _ = v35
                }
            }
            v15 -= v15
            var v36 uint16 = (v15 - 255)
            v17 = v17
            _ = fi21
            _ = v36
        }
        if (966 < 42) {
            for fi37 := 0; fi37 < 12; fi37++ {
                v38 := v2[(fi37 % ((v3) | 1))]
                v39 := GenS5{F0: -1, F1: (-904 == 32767)}
                fmt.Println(v38, v15)
                _ = fi37
                _ = v38
                _ = v39
            }
            switch ("foo bar" + v17) {
            case "foo bar":
                for fi40 := 0; fi40 < 13; fi40++ {
                    switch 4294967295 {
                    case 4294967295:
                        v41 := []int{(^(v3)), fi20, (fi40 & (v3 * (fi20 ^ 128))), (fi40 % ((32767) | 1))}
                        _ = v41
                    }
                    v17 = ""
                    _ = fi40
                }
            case "hello":
                v15 -= (^(v15))
            }
            if (255 >= 42) {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v43 := GenS1{F0: 0, F1: 4294967295}
                    _ = fi42
                    _ = v43
                }
                delete(v2, -9223372036854775808)
                v17 = "\n\t"
            }
        }
        v3 += (v3 << ((18446744073709551615) & 63))
        _ = fi20
    }
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
}
