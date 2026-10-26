package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint8) uint64 {
    v7 := []int{-683, 9223372036854775807, 127}
    var v8 uint64 = 42
    _ = p0
    _ = v7
    _ = v8
    return (^(v8))
}
func genFn1(p0 int, p1 string, p2 int8) uint {
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int = (p0 / ((p0) | 1))
        if false {
            var v11 int8 = (p2 * p2)
            _ = v11
        } else {
            p2 = p2
            v12 := []int{44}
            if true {
                var v13 uint = 734
                _ = v13
            } else {
                if true {
                    p1 = ""
                } else {
                    if true {
                        v14 := GenS1{F0: 117, F1: 65535, F2: 128}
                        v12 = append(v12, p0)
                        switch 178 {
                        case 927:
                            switch (p1 + p1) {
                            case "hello":
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    if ((false && false) && true) {
                                        if len(v12) > 0 {
                                            v12[0] = fi9
                                        }
                                        var v16 uint8 = 226
                                        v10 -= (fi9 & 127)
                                        switch 32767 {
                                        case 4294967295:
                                            _ = len(v12)
                                        default:
                                            if ((v16 + v16) == 5) {
                                                var v17 int32 = 1
                                                fmt.Println(v16, p2)
                                                _ = v17
                                            }
                                        }
                                        _ = v16
                                    }
                                    _ = fi15
                                }
                            }
                        default:
                            p2 += p2
                        }
                        _ = v14
                    }
                    v18 := GenS2{F0: true}
                    v19 := GenS1{F0: 1, F1: 4294967295, F2: 212}
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := GenS3{F0: ((1 >= 65535) && false), F1: p1, F2: (v10 - (fi9 % (((fi20 >> ((42) & 63))) | 1)))}
                    fmt.Println(fi20, fi20)
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int64 = -1
                    var v24 uint8 = 248
                    fmt.Println(p1, v23, fi9)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = v12
        }
        v25 := GenS4{F0: p1}
        _ = fi9
        _ = v10
        _ = v25
    }
    switch (p0 | p0) {
    case -2147483648:
        var v26 int64 = 9223372036854775807
        _ = v26
    case -771:
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint32 = 2147483648
            fmt.Println(v28)
            fmt.Println(p1, fi27, p0)
            _ = fi27
            _ = v28
        }
    default:
        var v29 float64 = float64(1.5)
        _ = v29
    }
    var v30 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 uint = 65535
    v4 := make(map[uint32]uint32)
    v5 := []string{("" + "")}
    var v6 uint = 18446744073709551615
    v3 -= (^(v3))
    v3 += (^(v3))
    v6 = 18446744073709551615
    var v31 uint32 = 168
    v32 := GenS1{F0: 873, F1: (v3 | 4294967295), F2: 42}
    var v33 bool = false
    fmt.Println(v6)
    switch "\n\t" {
    case "a":
        if (!(67 >= (v6 << ((65535) & 63)))) {
            v34 := GenS5{F0: 0, F1: genFn0(49), F2: 51}
            for fi35 := 0; fi35 < 20; fi35++ {
                if ((fi35 / ((fi35) | 1)) >= (fi35 * -9223372036854775808)) {
                    fmt.Println(fi35, v31, v33)
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v33 = (1.5 == (1e10 * (float32(0.0) - 3.14159)))
                        _ = fi36
                    }
                } else {
                    switch "hello" {
                    case "":
                        var v37 uint = (^(v6))
                        _ = v37
                    case "foo bar":
                        var v38 uint16 = 255
                        _ = v38
                    default:
                        v5 = append(v5, "hello")
                    }
                }
                v31 = (v31 - 846)
                var v39 int32 = -234
                _ = fi35
                _ = v39
            }
            var v40 uint16 = 32767
            _ = v34
            _ = v40
        } else {
            v3 = 1
        }
    default:
        if (20 == 18446744073709551615) {
            var v41 bool = v33
            if (1 >= 0) {
                v42 := GenS4{F0: "foo bar"}
                _ = v42
            } else {
                v43 := GenS6{F0: -1}
                _ = v43
            }
            _ = v41
        } else {
            v6 -= genFn1(4294967295, "\n\t", 127)
            v44 := GenS1{F0: 32767, F1: genFn1(1, "a", 42), F2: 42}
            _ = v44
        }
    }
    var v45 uint = v3
    v3 -= v6
    for fi46 := 0; fi46 < 19; fi46++ {
        v45 += (v45 >> ((v6) & 63))
        switch "\n\t" {
        case "\n\t":
            for fi47 := 0; fi47 < 15; fi47++ {
                var v48 int8 = 21
                fmt.Println(v31, v48, v33)
                v48 += (v48 / (((^(v48))) | 1))
                v6 += (v6 / (((^(v6))) | 1))
                v49 := GenS5{F0: 0, F1: 184, F2: 255}
                _ = fi47
                _ = v48
                _ = v49
            }
        case "a":
            var v50 uint16 = 986
            _ = v50
        }
        _ = fi46
    }
    switch 32767 {
    case 952:
        var v51 int32 = -1
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
