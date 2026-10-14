package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS7 struct { F0 int; F1 uint8; F2 bool }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint, p1 bool, p2 uint32) int64 {
    v7 := GenS1{F0: -1}
    if p1 {
        p0 = (^(p0))
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8, p2, p2)
        p1 = true
        v9 := []int{127, 128}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 4294967295, F1: 2, F2: 879}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"foo bar", (("" + "foo bar") + "\n\t")}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 float32 = float32(((0.0 / -2.25) - -2.25))
    var v6 int16 = 1
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS3{F0: 382, F1: (fi11 / ((128) | 1))}
        if (!true) {
            for fi13 := 0; fi13 < 8; fi13++ {
                v6 = 1000
                if (!true) {
                    if false {
                        fmt.Println(fi11, v5)
                        if (false && (!false)) {
                            var v14 uint64 = 410
                            _ = v14
                        }
                    }
                    fmt.Println(v6)
                    _ = len(v2)
                }
                v15 := GenS3{F0: (v6 << ((2) & 63)), F1: fi13}
                _ = fi13
                _ = v15
            }
            var v16 bool = true
            _ = v16
        } else {
            fmt.Println(fi11, v5)
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int64 = 9223372036854775807
                switch (fi11 - fi11) {
                case -2147483648:
                    fmt.Println(v18)
                case -69:
                    var v19 float32 = v5
                    _ = v19
                case 346:
                    v6 = 1
                default:
                    delete(v3, fi11)
                }
                var v20 int = -171
                if (false && (!true)) {
                    switch (fi11 << ((865) & 63)) {
                    case 42:
                        if false {
                            v5 = -2.25
                            v20 -= (v20 % ((v20) | 1))
                        }
                    case 2147483647:
                        v2 = append(v2, "\n\t")
                    case 0:
                        delete(v1, 4294967295)
                    default:
                        if true {
                            var v21 float64 = 1e10
                            v18 = (v18 | (v18 * v18))
                            if false {
                                var v22 uint8 = 128
                                if (false && false) {
                                    switch "" {
                                    case "a":
                                        switch (fi17 | 190) {
                                        case 1:
                                            var v23 uint = 0
                                            _ = v23
                                        case -36:
                                            var v24 uint = 4294967295
                                            _ = v24
                                        default:
                                            v25 := GenS4{F0: (v20 / (((-(fi17))) | 1))}
                                            _ = v25
                                        }
                                        delete(v3, (v20 & fi17))
                                        v26 := GenS5{F0: v6, F1: -24}
                                        _ = v26
                                    case "hello":
                                        fmt.Println(v5)
                                    case "foo bar":
                                        v27 := GenS6{F0: 162, F1: 42, F2: (v6 << ((4294967295) & 63))}
                                        v22 += (v22 >> ((579) & 63))
                                        v2 = append(v2, "a")
                                        _ = v27
                                    }
                                }
                                v28 := v4[4294967295]
                                _ = v22
                                _ = v28
                            }
                            _ = v21
                        } else {
                            v20 = (v20 >> ((0) & 63))
                            var v29 uint32 = 42
                            var v30 float32 = v5
                            _ = v29
                            _ = v30
                        }
                    }
                    var v31 uint16 = 0
                    v31 = v31
                    _ = v31
                } else {
                    v5 += float32(1e10)
                    v32 := v3[v20]
                    _ = v32
                }
                var v33 uint = 1
                fmt.Println(fi11)
                fmt.Println(fi17, fi11)
                v20 = fi17
                switch (fi17 << ((474) & 63)) {
                case 4294967295:
                    v6 -= (v6 ^ (v6 - (v6 ^ 0)))
                case 42:
                    v18 += v18
                default:
                    v20 = 4294967295
                    v18 = v18
                }
                _ = fi17
                _ = v18
                _ = v20
                _ = v33
            }
        }
        v6 -= 341
        _ = fi11
        _ = v12
    }
    var v34 int8 = -20
    var v35 int = 1
    v5 += v5
    v6 -= (v6 % ((v6) | 1))
    for fi36 := 0; fi36 < 1; fi36++ {
        var v37 bool = (((v6 & (-(v6))) < v6) || (-9223372036854775808 == -958))
        v38 := GenS7{F0: 127, F1: 128, F2: (1 >= 1)}
        _ = fi36
        _ = v37
        _ = v38
    }
    if (!(false && true)) {
        var v39 int16 = v6
        if ((!false) || (v35 <= v35)) {
            fmt.Println(v6)
            var v40 uint8 = 127
            v6 = v6
            _ = v40
        } else {
            v41 := GenS8{F0: 9223372036854775807}
            if len(v2) > 0 {
                v2[0] = ("a" + "")
            }
            _ = v41
        }
        var v42 uint = 42
        _ = v39
        _ = v42
    }
    v6 = v6
    var v43 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
