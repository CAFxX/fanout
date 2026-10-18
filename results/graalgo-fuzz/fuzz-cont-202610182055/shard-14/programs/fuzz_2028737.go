package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int32 }

func main() {
    v1 := []int{-967, 4294967295, -9223372036854775808, 127}
    var v2 uint = 0
    v3 := make(map[string]int)
    var v4 float64 = 0.0
    if false {
        var v5 uint64 = 42
        _ = v5
    }
    switch "" {
    case "\n\t":
        var v6 string = ""
        _ = v6
    default:
        var v7 uint = v2
        _ = v7
    }
    v3["\n\t"] = 42
    v8 := GenS1{F0: 331}
    switch "foo bar" {
    case "foo bar":
        var v9 bool = false
        _ = v9
    case "a":
        v4 += (v4 - ((((v4 * v4) - ((float64(-2.25) * ((3.14159 / float64(-2.25)) - float64(1e10))) + float64((-2.25 * 1e10)))) - -2.25) / (((-2.25 + v4) / float64(1e10)) * 1.5)))
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            if (2 >= 32767) {
                switch (^(fi10)) {
                case 127:
                    v4 = v4
                case 612:
                    switch fi10 {
                    case 364:
                        fmt.Println(v2)
                    case -1:
                        v4 -= v4
                    case -270:
                        if (-2.25 <= 0.1) {
                            if true {
                                v11 := GenS1{F0: 544}
                                _ = v11
                            }
                        }
                    }
                default:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 uint32 = 211
                        _ = fi12
                        _ = v13
                    }
                }
                var v14 bool = false
                v2 = (v2 << ((841) & 63))
                _ = v14
            }
            delete(v3, "a")
            _ = fi10
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 42, F1: (fi15 << ((v2) & 63)), F2: -1}
        _ = fi15
        _ = v16
    }
    v2 = v2
    var v17 int16 = 473
    var v18 uint64 = 2
    switch "" {
    case "\n\t":
        var v19 int64 = 0
        _ = v19
    case "":
        var v20 int64 = -1
        _ = v20
    case "foo bar":
        v21 := v3["\n\t"]
        _ = v21
    }
    var v22 int32 = 370
    fmt.Println(v2, v18)
    v22 += v22
    v1 = append(v1, -774)
    if true {
        _ = len(v1)
        switch -884 {
        case -9223372036854775808:
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(v17, v18, v18)
                _ = fi23
            }
        case 1:
            var v24 float32 = float32((0.1 * 3.14159))
            _ = v24
        default:
            var v25 float64 = 0.1
            _ = v25
        }
    } else {
        var v26 string = (("a" + "\n\t") + ("a" + ""))
        _ = v26
    }
    if (true || false) {
        v17 -= v17
        switch "\n\t" {
        case "a":
            switch "hello" {
            case "":
                for fi27 := 0; fi27 < 11; fi27++ {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        if true {
                            _ = len(v1)
                        } else {
                            if (!((!(false && false)) || (((v4 < v4) && (true && true)) && false))) {
                                v29 := GenS3{F0: fi27, F1: (v17 | v17)}
                                if (!false) {
                                    v4 = (v4 / (float64(0.1) - 1e10))
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi28))
                                    }
                                    if (((("hello" + "hello") + ("foo bar" + "hello")) + (("\n\t" + "a") + (("a" + ("" + ("hello" + "a"))) + "\n\t"))) <= "") {
                                        v2 -= 4294967295
                                        _ = len(v3)
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = -1
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            v4 = v4
                                        case "":
                                            switch 128 {
                                            case -180:
                                                v30 := GenS2{F0: 42, F1: fi27, F2: fi27}
                                                _ = v30
                                            case -1:
                                                v31 := GenS4{F0: 2}
                                                _ = v31
                                            }
                                        default:
                                            switch fi27 {
                                            case -544:
                                                var v32 bool = false
                                                _ = v32
                                            case 32767:
                                                var v33 int64 = 42
                                                _ = v33
                                            default:
                                                switch (fi28 & (fi28 * (fi27 ^ 1))) {
                                                case 2147483647:
                                                    v34 := GenS3{F0: (^(fi27)), F1: (-(v17))}
                                                    _ = v34
                                                default:
                                                    v35 := v3["hello"]
                                                    var v36 uint = 65535
                                                    var v37 uint64 = v18
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v29
                            } else {
                                v4 = v4
                                var v38 int16 = v17
                                fmt.Println(fi27, v38)
                                _ = v38
                            }
                            v2 -= (v2 / ((v2) | 1))
                            var v39 bool = (42 == -2147483648)
                            _ = v39
                        }
                        fmt.Println(v2, v2)
                        v40 := GenS5{F0: (v2 - v2), F1: int32(v17)}
                        var v41 string = "a"
                        v42 := GenS3{F0: fi28, F1: (v17 + 1000)}
                        v43 := GenS2{F0: int8(v18), F1: 913, F2: (fi28 << (((v2 % ((0) | 1))) & 63))}
                        _ = fi28
                        _ = v40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi27
                }
            case "foo bar":
                var v44 int8 = 0
                _ = v44
            case "hello":
                fmt.Println(v2, v18, v22)
            }
        case "hello":
            _ = len(v3)
        case "foo bar":
            var v45 uint = v2
            _ = v45
        default:
            v1 = append(v1, -2147483648)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
