package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16; F2 uint64 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 127
    var v3 string = "\n\t"
    var v4 int = (v2 - (-(v2)))
    switch (v4 / (((v2 - (v4 >> ((1) & 63)))) | 1)) {
    case 492:
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := []int{123}
            _ = len(v1)
            v4 = (^(v4))
            _ = fi5
            _ = v6
        }
    case 127:
        v4 = -9223372036854775808
    case -882:
        if (877 == 366) {
            v7 := GenS1{F0: 877, F1: v3}
            _ = v7
        }
    }
    v8 := GenS2{F0: 42}
    delete(v1, 354)
    v9 := GenS2{F0: 533}
    v2 += (v2 & v2)
    fmt.Println(v2, v2, v4)
    v4 += (v4 / (((-(v2))) | 1))
    v10 := []int{}
    var v11 int8 = -1
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int16 = -32768
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(v4, v3)
            if false {
                v15 := GenS1{F0: 703, F1: v3}
                _ = v15
            } else {
                var v16 int8 = int8(v13)
                if ((-(v13)) >= v13) {
                    v1[65535] = false
                    var v17 int = -9223372036854775808
                    _ = v17
                }
                _ = v16
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 4294967295}
        v4 = (-(v4))
        _ = fi18
        _ = v19
    }
    if (float32(3.14159) < 0.0) {
        fmt.Println(v3, v3)
        v20 := GenS3{F0: false, F1: 0, F2: 57}
        _ = v20
    } else {
        v21 := GenS4{F0: v11}
        var v22 float32 = 1.5
        var v23 int8 = (v11 ^ (-(v11)))
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int = v2
    var v25 bool = true
    switch v24 {
    case 4294967295:
        v4 += v4
    case 2147483647:
        v3 = v3
    case 309:
        v24 = v4
    default:
        switch "\n\t" {
        case "":
            if v25 {
                switch v3 {
                case "":
                    var v26 int32 = -1
                    _ = v26
                case "\n\t":
                    for fi27 := 0; fi27 < 12; fi27++ {
                        switch v3 {
                        case "\n\t":
                            fmt.Println(v11, v11, v24)
                        case "":
                            v24 += fi27
                        case "a":
                            v1[2147483647] = (((6 >= 4294967295) || (4294967295 > 42)) && v25)
                        }
                        for fi28 := 0; fi28 < 10; fi28++ {
                            fmt.Println(v2, v3, fi27)
                            _ = fi28
                        }
                        for fi29 := 0; fi29 < 10; fi29++ {
                            _ = len(v1)
                            v10 = append(v10, (v4 % (((-(fi29))) | 1)))
                            if v25 {
                                v25 = (18446744073709551615 < 18446744073709551615)
                                v30 := v1[233]
                                var v31 int16 = -1
                                v24 += (-(v4))
                                var v32 int8 = v11
                                var v33 float64 = float64(3.14159)
                                var v34 float64 = v33
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                v11 = v11
                                var v35 int64 = 1
                                _ = v35
                            }
                            _ = fi29
                        }
                        _ = fi27
                    }
                case "foo bar":
                    v4 = int(v11)
                default:
                    var v36 uint16 = 157
                    _ = v36
                }
                if (v25 || true) {
                    v37 := GenS2{F0: 952}
                    _ = v37
                }
            } else {
                if (((0.1 * 1e10) + (1.5 * 0.1)) < 0.1) {
                    var v38 bool = v25
                    var v39 int = (v4 - (v4 ^ v2))
                    _ = v38
                    _ = v39
                }
                delete(v1, 2)
            }
        case "\n\t":
            if (true || (!((^(v4)) <= (v2 >> ((451) & 63))))) {
                var v40 uint32 = 0
                var v41 uint32 = (v40 % (((^(v40))) | 1))
                fmt.Println(v41, v3)
                _ = v40
                _ = v41
            }
        default:
            if (567 != 2147483647) {
                v24 -= v24
            } else {
                var v42 string = (v3 + v3)
                _ = v42
            }
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
