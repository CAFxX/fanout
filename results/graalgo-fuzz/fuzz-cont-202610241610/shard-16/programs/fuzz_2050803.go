package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint) uint16 {
    p0 += p0
    var v6 uint16 = 884
    var v7 float64 = 3.14159
    v8 := GenS1{F0: 127, F1: 172, F2: 10}
    if ((("\n\t" + "\n\t") != "a") && true) {
        v9 := []int{42, 32767}
        var v10 float32 = 1.5
        _ = v10
        _ = v9
    } else {
        switch "foo bar" {
        case "":
            var v11 int8 = 112
            _ = v11
        case "\n\t":
            v7 = v7
        case "hello":
            if (float64(3.14159) >= (1e10 * -2.25)) {
                if true {
                    var v12 float32 = float32((-2.25 + v7))
                    var v13 int64 = -450
                    var v14 uint64 = 2
                    _ = v12
                    _ = v13
                    _ = v14
                }
                switch "" {
                case "foo bar":
                    switch 127 {
                    case 4294967295:
                        for fi15 := 0; fi15 < 6; fi15++ {
                            for fi16 := 0; fi16 < 15; fi16++ {
                                if (!true) {
                                    if true {
                                        v17 := []int{}
                                        _ = v17
                                    }
                                } else {
                                    fmt.Println(p1, p1)
                                    v6 = v6
                                }
                                _ = fi16
                            }
                            var v18 int64 = -430
                            fmt.Println(p0, v18, p1)
                            _ = fi15
                            _ = v18
                        }
                    case 32767:
                        v19 := GenS2{F0: 2, F1: 0, F2: v6}
                        _ = v19
                    case -9223372036854775808:
                        fmt.Println(v7, v6, p0)
                    }
                default:
                    var v20 uint64 = 1
                    _ = v20
                }
                v21 := GenS3{F0: 680}
                _ = v21
            }
        default:
            fmt.Println(p0)
        }
        var v22 uint16 = 404
        p1 = p1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (v6 / ((v6) | 1))
}
func genFn1(p0 int8) uint16 {
    v23 := GenS1{F0: p0, F1: 659, F2: p0}
    var v24 uint32 = 0
    v24 -= 694
    switch 2147483647 {
    case -2147483648:
        switch "a" {
        case "a":
            p0 = (p0 >> ((0) & 63))
        case "hello":
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 int16 = 516
                _ = fi25
                _ = v26
            }
        case "\n\t":
            v24 += v24
        }
    }
    _ = p0
    _ = v23
    _ = v24
    return genFn0(32767, 165)
}

func main() {
    var v1 bool = false
    var v2 uint16 = 410
    var v3 uint16 = v2
    var v4 uint16 = (v3 ^ (v2 % ((v3) | 1)))
    v5 := make(map[string]uint32)
    v1 = (false || (!true))
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS1{F0: -127, F1: -1, F2: 100}
        var v29 uint16 = 159
        v2 += v2
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 42
    v30 = 1
    var v31 int = -1
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 bool = (1 == 0)
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    if (!v1) {
        fmt.Println(v31, v30, v4)
        var v35 float32 = ((0.1 / 1.5) - 0.1)
        v3 -= (^(v2))
        _ = v35
    } else {
        var v36 int = v31
        var v37 uint16 = 42
        _ = v36
        _ = v37
    }
    var v38 uint64 = 18446744073709551615
    var v39 bool = v1
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        fmt.Println(v1, v1)
    case "hello":
        if false {
            v40 := GenS1{F0: 127, F1: 1, F2: -1}
            var v41 uint64 = 18446744073709551615
            _ = v40
            _ = v41
        } else {
            var v42 bool = ((((-2.25 + (-2.25 - (0.1 / (3.14159 / 1.5)))) * 1.5) + (0.1 - 1.5)) > ((0.0 - 0.1) * 0.0))
            _ = v42
        }
    default:
        v43 := GenS4{F0: 255}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
