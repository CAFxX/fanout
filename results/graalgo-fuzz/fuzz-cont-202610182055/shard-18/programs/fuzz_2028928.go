package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
func genFn0(p0 int32) uint16 {
    var v7 uint16 = 32767
    switch "a" {
    case "hello":
        var v8 uint16 = v7
        _ = v8
    case "":
        v9 := GenS1{F0: 42, F1: "foo bar", F2: 255}
        _ = v9
    case "\n\t":
        switch -889 {
        case -2147483648:
            if ((v7 / ((1) | 1)) != v7) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        switch ("" + ("" + "foo bar")) {
                        case "foo bar":
                            switch "foo bar" {
                            case "a":
                                var v12 int32 = (p0 << ((0) & 63))
                                fmt.Println(p0)
                                _ = v12
                            case "\n\t":
                                fmt.Println(fi11)
                            case "foo bar":
                                v13 := GenS1{F0: 127, F1: ("a" + "foo bar"), F2: 42}
                                _ = v13
                            default:
                                var v14 string = "hello"
                                _ = v14
                            }
                        case "":
                            switch int(v7) {
                            case 0:
                                var v15 int32 = 819
                                _ = v15
                            case 32768:
                                fmt.Println(fi11)
                            default:
                                var v16 uint = 120
                                _ = v16
                            }
                        case "hello":
                            v17 := []int{fi11, -787, (fi10 - 42)}
                            _ = v17
                        }
                        fmt.Println(p0, p0, fi11)
                        _ = fi11
                    }
                    _ = fi10
                }
                fmt.Println(v7)
            }
        }
    default:
        switch 2147483647 {
        case 32767:
            var v18 int64 = -1
            _ = v18
        case -174:
            if (2147483647 < 32767) {
                v7 = v7
            }
        case -9223372036854775808:
            var v19 float64 = float64((0.1 / 1.0))
            _ = v19
        default:
            v20 := GenS1{F0: 255, F1: (("foo bar" + "hello") + "a"), F2: 42}
            _ = v20
        }
    }
    _ = p0
    _ = v7
    return (^(v7))
}
func genFn1(p0 int8, p1 int64) int {
    fmt.Println(p1)
    var v21 uint = 896
    p0 = (p0 << ((0) & 63))
    _ = p0
    _ = p1
    _ = v21
    return 4294967295
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 int32 = -2147483648
    var v5 int = 127
    v6 := make(map[string]uint32)
    v22 := GenS1{F0: 2, F1: "hello", F2: (^(v1))}
    fmt.Println(v1)
    var v23 uint32 = 590
    var v24 float32 = (0.1 / 0.1)
    if false {
        v25 := []int{(^(v5)), -973, (v5 % (((v5 << ((v1) & 63))) | 1))}
        _ = v25
    }
    v5 -= int(v4)
    v26 := []int{}
    v1 = (v1 << ((9) & 63))
    for fi27 := 0; fi27 < 18; fi27++ {
        v5 += (fi27 ^ (fi27 & v5))
        var v28 uint16 = 65535
        switch genFn1(114, 0) {
        case 4294967295:
            var v29 int64 = -612
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    if true {
        fmt.Println(v24, v4, v1)
        if true {
            if (!(genFn0(v4) <= 4)) {
                var v30 int32 = (-(v4))
                _ = v30
            }
            v24 -= v24
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                for fi32 := 0; fi32 < 8; fi32++ {
                    v4 = 1
                    _ = fi32
                }
                var v33 float64 = (1e10 + ((1.5 / 1.0) - 1e10))
                _ = fi31
                _ = v33
            }
        }
    }
    fmt.Println(v23, v1, v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
