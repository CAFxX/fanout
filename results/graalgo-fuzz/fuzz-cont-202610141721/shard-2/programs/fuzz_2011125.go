package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 bool; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint8 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]bool)
    v3 := []int{32768, 32768, 1, 361}
    var v4 int16 = v1
    var v5 string = ("\n\t" + "")
    v6 := []string{("foo bar" + (v5 + v5)), (v5 + (v5 + "a")), (v5 + "foo bar"), "foo bar"}
    v7 := GenS1{F0: -9223372036854775808, F1: 1, F2: 955}
    var v8 uint16 = 1
    if (-38 <= 492) {
        switch v5 {
        case "a":
            var v9 float32 = -2.25
            _ = v9
        case "":
            v4 = (-(v1))
        case "\n\t":
            fmt.Println(v8, v8)
        }
        var v10 string = (((v5 + "hello") + v5) + "")
        switch v10 {
        case "foo bar":
            if (0 == 103) {
                var v11 int32 = -669
                _ = v11
            }
        default:
            if true {
                for fi12 := 0; fi12 < 10; fi12++ {
                    v10 = v5
                    var v13 int64 = 645
                    _ = fi12
                    _ = v13
                }
                fmt.Println(v1, v5)
            } else {
                v14 := GenS2{F0: 872, F1: 0, F2: (v8 ^ (v8 % ((1) | 1)))}
                _ = v14
            }
        }
        _ = v10
    }
    v15 := GenS3{F0: 120}
    v1 += v4
    delete(v2, 234)
    v1 -= 490
    if ("foo bar" < "\n\t") {
        fmt.Println(v4, v1, v5)
        v5 = (v5 + "hello")
    }
    _ = len(v2)
    for fi16 := 0; fi16 < 7; fi16++ {
        if len(v6) > 0 {
            v6[0] = ""
        }
        if false {
            v4 += v4
            var v17 float32 = 0.1
            var v18 uint8 = 128
            _ = v17
            _ = v18
        } else {
            var v19 uint8 = 1
            _ = v19
        }
        var v20 bool = false
        _ = fi16
        _ = v20
    }
    v5 = v5
    if true {
        var v21 float64 = -2.25
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            switch -1 {
            case 42:
                v22 := GenS1{F0: 364, F1: 100, F2: 2}
                _ = v22
            case 7:
                v1 += 42
            case 32767:
                for fi23 := 0; fi23 < 2; fi23++ {
                    if false {
                        var v24 uint64 = 25
                        if ((v5 < v5) || false) {
                            var v25 uint16 = 42
                            switch fi23 {
                            case -797:
                                switch 127 {
                                case 316:
                                    fmt.Println(v8, v5)
                                }
                            case 510:
                                v26 := GenS4{F0: 0, F1: 85, F2: -32}
                                _ = v26
                            case 0:
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    fmt.Println(v25, v1, v24)
                                    if (v24 <= (v24 + 2)) {
                                        v3 = append(v3, fi23)
                                        var v28 int8 = 115
                                        var v29 float64 = v21
                                        _ = v28
                                        _ = v29
                                    }
                                    v30 := GenS4{F0: 42, F1: 2, F2: 1}
                                    _ = fi27
                                    _ = v30
                                }
                            }
                            _ = v25
                        } else {
                            var v31 uint8 = 40
                            _ = v31
                        }
                        _ = v24
                    }
                    switch ("" + v5) {
                    case "foo bar":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "":
                        var v32 uint8 = 42
                        _ = v32
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            if false {
                                switch (v5 + v5) {
                                case "foo bar":
                                    for fi33 := 0; fi33 < 1; fi33++ {
                                        v5 = v5
                                        v34 := GenS5{F0: 402}
                                        var v35 int64 = -824
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                case "hello":
                                    if (1 <= 1) {
                                        v36 := GenS4{F0: 2, F1: 2, F2: 0}
                                        for fi37 := 0; fi37 < 6; fi37++ {
                                            var v38 int8 = -97
                                            switch "foo bar" {
                                            case "\n\t":
                                                v21 = v21
                                            }
                                            _ = fi37
                                            _ = v38
                                        }
                                        _ = v36
                                    }
                                default:
                                    v21 += ((1.5 - v21) / (float64(0.0) * (v21 + float64((0.0 / 1.5)))))
                                }
                            } else {
                                var v39 float32 = 1.5
                                for fi40 := 0; fi40 < 14; fi40++ {
                                    var v41 int8 = -1
                                    v42 := GenS6{F0: 761, F1: (true || (42 > 1)), F2: -1}
                                    v43 := GenS7{F0: 1, F1: 2, F2: 6}
                                    v44 := GenS8{F0: 954, F1: -1, F2: 255}
                                    var v45 uint16 = (v8 >> ((34) & 63))
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                var v46 uint = 255
                                _ = v39
                                _ = v46
                            }
                        default:
                            v2[fi23] = ((275 < 4294967295) || (1 == 145))
                        }
                    }
                    var v47 uint64 = 24
                    _ = fi23
                    _ = v47
                }
            }
            var v48 float64 = (((v21 / v21) + v21) * (v21 * v21))
            var v49 string = (("\n\t" + "hello") + (v5 + "a"))
            _ = v48
            _ = v49
        }
        _ = v21
    } else {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
