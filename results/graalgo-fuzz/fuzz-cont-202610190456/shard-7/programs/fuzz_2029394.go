package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint, p1 uint8) string {
    p1 += 133
    var v6 float64 = 1.5
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}
func genFn1() uint32 {
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 float32 = (-2.25 * float32((((0.1 - 1e10) + float64(1e10)) + 0.0)))
            var v9 float64 = 0.1
            fmt.Println(v9, v9, v8)
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{42, 419}
        var v11 int32 = -840
        _ = v10
        _ = v11
    } else {
        v12 := []int{127}
        var v13 int64 = -75
        _ = v12
        _ = v13
    }
    if false {
        v14 := GenS1{F0: 290, F1: -1}
        _ = v14
    }
    v15 := []int{32768}
    v16 := GenS1{F0: 560, F1: 0}
    _ = v15
    _ = v16
    return 948
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    var v3 string = "\n\t"
    v4 := []int{-300, -9223372036854775808}
    var v5 int16 = 155
    var v17 uint = 532
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v3, v5)
        var v19 int64 = 673
        _ = fi18
        _ = v19
    }
    var v20 uint64 = v2
    v5 += v5
    switch "hello" {
    case "hello":
        v21 := GenS2{F0: 2147483647}
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 string = v3
            _ = fi22
            _ = v23
        }
    default:
        fmt.Println(v17)
    }
    fmt.Println(v20, v5, v17)
    v1[-1] = false
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: 0, F1: 100}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    switch 76 {
    case 0:
        var v27 int64 = -9223372036854775808
        _ = v27
    case 89:
        switch 1 {
        case -380:
            for fi28 := 0; fi28 < 1; fi28++ {
                switch ((v3 + v3) + v3) {
                case "\n\t":
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if true {
                            var v30 int64 = 9223372036854775807
                            if true {
                                fmt.Println(v3, fi28)
                                v2 += (v20 / (((v2 & 340)) | 1))
                            } else {
                                delete(v1, fi29)
                                fmt.Println(v5)
                                v17 -= (v17 * (v17 - (^(v17))))
                                v17 += v17
                                _ = len(v1)
                                fmt.Println(v30, fi29, fi28)
                            }
                            _ = v30
                        }
                        v17 = (v17 / ((v17) | 1))
                        switch int(v20) {
                        case -108:
                            var v31 int64 = 693
                            _ = v31
                        case -141:
                            var v32 int32 = 2147483647
                            _ = v32
                        case 42:
                            fmt.Println(fi29)
                        default:
                            var v33 uint8 = 131
                            _ = v33
                        }
                        _ = fi29
                    }
                case "hello":
                    switch (v3 + v3) {
                    case "foo bar":
                        fmt.Println(v3, fi28, fi28)
                    default:
                        if (-2.25 >= float32((float64(-2.25) / (1.5 / 1.5)))) {
                            for fi34 := 0; fi34 < 19; fi34++ {
                                v5 = (v5 << ((1) & 63))
                                v5 += (v5 ^ (v5 + 32767))
                                switch fi28 {
                                case 0:
                                    fmt.Println(fi34)
                                    var v35 int64 = 1
                                    var v36 float64 = 0.0
                                    switch v3 {
                                    case "hello":
                                        v36 -= 0.1
                                        v5 = (-(v5))
                                        switch v3 {
                                        case "a":
                                            var v37 uint64 = v2
                                            _ = v37
                                        default:
                                            var v38 uint = (^(v17))
                                            _ = v38
                                        }
                                    case "":
                                        var v39 int16 = -32768
                                        _ = v39
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi34
                            }
                            var v40 uint = (v17 & 441)
                            v5 = v5
                            _ = v40
                        }
                    }
                case "a":
                    v41 := GenS1{F0: 37, F1: -1}
                    _ = v41
                }
                fmt.Println(v17, v2)
                v1[(fi28 ^ fi28)] = true
                _ = fi28
            }
        }
    case 4294967295:
        for fi42 := 0; fi42 < 19; fi42++ {
            v43 := GenS1{F0: 597, F1: 42}
            _ = len(v4)
            var v44 int64 = -257
            _ = fi42
            _ = v43
            _ = v44
        }
    default:
        v3 = v3
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        switch v3 {
        case "":
            var v46 float64 = float64((-2.25 + 0.1))
            _ = v46
        }
        _ = fi45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
