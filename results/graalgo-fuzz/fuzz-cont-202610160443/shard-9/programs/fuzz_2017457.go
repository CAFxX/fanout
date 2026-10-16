package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
func genFn0(p0 int8) int32 {
    if (1000 != -15) {
        switch "\n\t" {
        case "\n\t":
            v5 := []int{}
            _ = v5
        case "a":
            if false {
                if true {
                    var v6 float32 = ((1.5 + 1e10) * float32(1.5))
                    v6 = 1e10
                    var v7 int32 = 42
                    _ = v6
                    _ = v7
                }
                v8 := GenS1{F0: -1, F1: -1}
                _ = v8
            } else {
                if true {
                    var v9 uint16 = 1
                    _ = v9
                } else {
                    if ((((2147483647 > -2147483648) || (0 != 255)) && false) && (-9223372036854775808 <= -2147483648)) {
                        var v10 bool = false
                        var v11 uint32 = 2147483648
                        _ = v10
                        _ = v11
                    } else {
                        switch 2147483647 {
                        case 66:
                            var v12 uint16 = 1
                            _ = v12
                        case 32767:
                            var v13 int8 = p0
                            _ = v13
                        }
                    }
                    for fi14 := 0; fi14 < 7; fi14++ {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v16 := []int{(fi15 % ((377) | 1))}
                            if ((53 != 2) || ((2147483647 >= 302) && true)) {
                                var v17 bool = false
                                _ = v17
                            } else {
                                v18 := GenS1{F0: 2147483647, F1: 42}
                                var v19 int16 = 471
                                _ = v18
                                _ = v19
                            }
                            var v20 int64 = 1
                            _ = fi15
                            _ = v16
                            _ = v20
                        }
                        v21 := GenS1{F0: 42, F1: 0}
                        _ = fi14
                        _ = v21
                    }
                }
                fmt.Println(p0)
                var v22 float32 = float32(1e10)
                _ = v22
            }
        default:
            fmt.Println(p0)
        }
        p0 = 127
    }
    var v23 float32 = (1.5 * float32(((1.5 - -2.25) + (0.0 * 0.1))))
    v24 := GenS2{F0: 4294967295, F1: int16(p0)}
    p0 = (p0 + p0)
    var v25 uint8 = 110
    _ = p0
    _ = v23
    _ = v24
    _ = v25
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -629
    var v3 uint32 = 2147483648
    v4 := make(map[int]int)
    var v26 float32 = ((0.1 - (1.5 - 0.0)) - 1e10)
    if ((255 == 65535) && (128 >= 255)) {
        v27 := GenS3{F0: 2, F1: -9223372036854775808}
        v28 := GenS4{F0: 100}
        _ = v27
        _ = v28
    } else {
        switch "a" {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v2, v3)
            case "hello":
                v3 += 0
            case "":
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 float32 = v26
                    _ = fi29
                    _ = v30
                }
            }
        case "":
            var v31 uint = 65535
            _ = v31
        }
        switch 127 {
        case 9223372036854775807:
            v26 += (1e10 - (1.5 * (0.0 / v26)))
        case 0:
            if (!(1 > -1)) {
                var v32 int64 = 0
                v33 := []int{-2147483648, 0, -984}
                _ = v32
                _ = v33
            }
        }
        v34 := GenS4{F0: 0}
        _ = v34
    }
    v26 = (3.14159 - (((1.5 * 1e10) / v26) / 3.14159))
    v35 := []int{127, 32768, -560, int(v2)}
    fmt.Println(v3, v2, v2)
    v36 := GenS5{F0: (!(42 >= 99))}
    var v37 string = "a"
    switch 32767 {
    case 32767:
        var v38 float64 = -2.25
        _ = v38
    default:
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 bool = (173 != 927)
            _ = fi39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        var v42 uint64 = 914
        var v43 uint = 18446744073709551615
        var v44 int = (fi41 * -9223372036854775808)
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    var v45 uint64 = 0
    _ = len(v35)
    var v46 int = -1
    fmt.Println(v3)
    delete(v4, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v46
}
