package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 bool, p1 int64, p2 int64) uint64 {
    p2 = p2
    v7 := []int{}
    v8 := GenS1{F0: false, F1: 1, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool) uint64 {
    var v9 uint = 4294967295
    v9 -= v9
    if (!(!false)) {
        var v10 uint64 = 2
        fmt.Println(v10)
        switch 191 {
        case 4294967295:
            var v11 uint64 = (v10 % (((v10 << (((^(v9))) & 63))) | 1))
            _ = v11
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(fi12, v9, fi12)
                _ = fi12
            }
        }
        _ = v10
    } else {
        p0 = (1 >= 240)
        if (2147483647 >= 42) {
            v13 := GenS1{F0: (1 == 2147483647), F1: v9, F2: 1}
            _ = v13
        }
        if ((((0.0 / (((3.14159 + 1e10) / (-2.25 + float32(1e10))) + 0.0)) + 0.0) + (3.14159 / 1e10)) < 0.1) {
            var v14 bool = false
            switch 128 {
            case 32767:
                v15 := []int{42}
                _ = v15
            case 70:
                v16 := []int{2147483647, 32768, 575, 0}
                _ = v16
            }
            _ = v14
        }
    }
    var v17 float64 = -2.25
    p0 = p0
    _ = p0
    _ = v17
    _ = v9
    return genFn0(p0, 9223372036854775807, 1)
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 float32 = float32(0.1)
    v5 := []string{"", "foo bar", ("foo bar" + "hello"), "hello"}
    v6 := make(map[int]int)
    for fi18 := 0; fi18 < 8; fi18++ {
        if false {
            fmt.Println(fi18, fi18, v4)
            for fi19 := 0; fi19 < 11; fi19++ {
                v4 = 1e10
                fmt.Println(v4, v1)
                if ((0 < int32(fi18)) && false) {
                    switch fi18 {
                    case 128:
                        var v20 uint16 = 42
                        _ = v20
                    case 0:
                        v4 = (0.1 * -2.25)
                    case -184:
                        v1 = (v1 & 420)
                        var v21 string = "foo bar"
                        switch v21 {
                        case "a":
                            var v22 float64 = (((1e10 - 0.1) + (3.14159 + ((0.1 * -2.25) / 1.0))) - 3.14159)
                            _ = v22
                        }
                        var v23 uint32 = 984
                        _ = v21
                        _ = v23
                    default:
                        v4 = (3.14159 - (v4 - (v4 - v4)))
                        if true {
                            var v24 int32 = 1
                            if ((!false) || false) {
                                var v25 uint64 = genFn0((18446744073709551615 <= 18446744073709551615), 42, -1)
                                fmt.Println(v25, v4)
                                var v26 int8 = -1
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int16 = -533
                                v28 := GenS2{F0: fi18}
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        if (42 >= 388) {
                                            v24 -= v24
                                            v24 -= (v24 ^ (v24 | v24))
                                        }
                                        v4 += (-2.25 + 1e10)
                                        var v29 uint64 = 345
                                        _ = v29
                                    case "hello":
                                        fmt.Println(v4, v24)
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                _ = v27
                                _ = v28
                            }
                            if false {
                                delete(v6, (fi18 ^ fi19))
                                var v30 float32 = 1e10
                                _ = v30
                            }
                            _ = v24
                        } else {
                            fmt.Println(fi19, v1)
                        }
                        v31 := GenS3{F0: "foo bar", F1: 9223372036854775807}
                        _ = v31
                    }
                    var v32 uint16 = 1
                    var v33 uint64 = 18446744073709551615
                    _ = v32
                    _ = v33
                }
                var v34 int = fi18
                _ = fi19
                _ = v34
            }
            var v35 int16 = 993
            _ = v35
        } else {
            v36 := GenS4{F0: 100, F1: -1, F2: 809}
            fmt.Println(v4)
            _ = v36
        }
        if (false || (false && false)) {
            var v37 int = fi18
            v1 = v1
            _ = v37
        } else {
            v4 = ((v4 / 3.14159) - (float32(((((float64(((v4 + (v4 + (v4 - v4))) / 1.0)) + 0.1) / 3.14159) / -2.25) / 1.5)) + (float32((1.5 / 1e10)) + v4)))
        }
        _ = fi18
    }
    fmt.Println(v4, v4)
    v1 = (^(v1))
    var v38 string = "hello"
    if true {
        delete(v6, 0)
    }
    v1 -= 1
    var v39 uint64 = 4294967295
    _ = len(v6)
    switch v38 {
    case "\n\t":
        switch 128 {
        case 32768:
            v40 := v3[(((v38 + v38) + "hello") + (v38 + v38))]
            _ = v40
        case 128:
            v41 := GenS3{F0: "foo bar", F1: 9223372036854775807}
            _ = v41
        case 32767:
            v39 = v39
        default:
            fmt.Println(v4)
        }
    case "a":
        switch -89 {
        case -1:
            v42 := GenS5{F0: 100, F1: 1}
            _ = v42
        case 2147483647:
            if true {
                v39 -= genFn1(true)
                v4 = v4
                var v43 int8 = -8
                _ = v43
            } else {
                v5 = append(v5, v38)
                v44 := GenS6{F0: 646}
                v39 += (v39 & (^(v39)))
                _ = v44
            }
        }
    default:
        var v45 uint = 2
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
