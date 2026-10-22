package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0(p0 uint, p1 uint16, p2 uint8) uint32 {
    var v4 float64 = 0.1
    p0 -= 1
    p1 = 0
    switch -790 {
    case 9223372036854775807:
        switch 127 {
        case -644:
            for fi5 := 0; fi5 < 13; fi5++ {
                if (0 < 42) {
                    for fi6 := 0; fi6 < 12; fi6++ {
                        if false {
                            v7 := []int{fi6, (fi5 - fi6), (fi5 / ((fi6) | 1))}
                            _ = v7
                        }
                        _ = fi6
                    }
                    fmt.Println(p1)
                }
                var v8 bool = true
                var v9 string = "foo bar"
                _ = fi5
                _ = v8
                _ = v9
            }
        case 2147483647:
            var v10 float64 = 0.0
            _ = v10
        }
    case -748:
        var v11 int16 = 1
        _ = v11
    case -729:
        v4 += 0.0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1() uint16 {
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: -743}
        _ = fi12
        _ = v13
    }
    var v14 float32 = 3.14159
    fmt.Println(v14)
    v15 := GenS2{F0: 817, F1: 1, F2: "foo bar"}
    v14 = v14
    _ = v14
    _ = v15
    return 811
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{2147483648, 937}
    v3 := make(map[uint32]bool)
    switch "\n\t" {
    case "hello":
        var v16 string = "foo bar"
        _ = v16
    case "":
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 127}
            var v19 uint8 = 127
            _ = fi17
            _ = v18
            _ = v19
        }
    case "a":
        var v20 int8 = 127
        _ = v20
    default:
        if ((!false) || (!((!true) && true))) {
            var v21 uint8 = 214
            var v22 int32 = -2147483648
            fmt.Println(v21, v21)
            _ = v21
            _ = v22
        }
    }
    v3[944] = true
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int = fi23
        _ = fi23
        _ = v24
    }
    var v25 int8 = 100
    fmt.Println(v25)
    var v26 uint64 = 468
    var v27 uint64 = (^(v26))
    fmt.Println(v25, v25)
    fmt.Println(v26, v25)
    v26 = (v27 >> ((1) & 63))
    switch 127 {
    case 1:
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 int64 = 43
            for fi30 := 0; fi30 < 1; fi30++ {
                fmt.Println(v29, v29, v27)
                v2 = append(v2, 4294967295)
                switch "foo bar" {
                case "\n\t":
                    v27 -= 329
                case "hello":
                    if true {
                        fmt.Println(v29)
                    } else {
                        v27 += 0
                    }
                }
                v31 := GenS1{F0: -866}
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
    case -1:
        v32 := GenS2{F0: 96, F1: 42, F2: "\n\t"}
        _ = v32
    case 127:
        var v33 int32 = 1
        _ = v33
    default:
        v34 := GenS3{F0: 65535, F1: true, F2: 127}
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        if ((float64(0.1) + (0.1 - 1.5)) < float64(-2.25)) {
            fmt.Println(v26)
            if (2147483648 <= 2) {
                for fi36 := 0; fi36 < 1; fi36++ {
                    v27 += 1
                    _ = fi36
                }
            }
            switch 32768 {
            case 9223372036854775807:
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 int16 = -785
                    _ = len(v1)
                    v39 := GenS2{F0: 10, F1: 2147483647, F2: "a"}
                    v38 -= 42
                    v26 = (v27 & (v26 + (v26 << ((255) & 63))))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
        }
        _ = fi35
    }
    if true {
        if ((v27 - (v27 / ((2) | 1))) == v26) {
            switch "foo bar" {
            case "a":
                v25 += (v25 << ((369) & 63))
            case "hello":
                switch "a" {
                case "hello":
                    for fi40 := 0; fi40 < 4; fi40++ {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v25, v27)
                        case "":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi40
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case "\n\t":
                v27 = (^(v27))
            }
            delete(v1, "\n\t")
        }
        fmt.Println(v27, v26, v25)
        fmt.Println(v25, v25)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
