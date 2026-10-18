package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int64 }
func genFn0(p0 uint8) uint {
    if true {
        v7 := []int{9223372036854775807, -9223372036854775808, 42, 127}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 float32 = 0.0
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: (!false), F2: int8(p0)}
        _ = v10
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := []int{42, 711}
    var v2 int = -900
    var v3 int32 = -431
    var v4 int32 = (v3 | 282)
    v5 := make(map[int]bool)
    var v6 int = v2
    if true {
        v3 = v4
        if false {
            v1 = append(v1, -140)
            var v11 uint8 = 255
            _ = v11
        } else {
            if (false || true) {
                if (!false) {
                    fmt.Println(v2, v3)
                    if (0 <= v3) {
                        var v12 int = (v2 / ((v6) | 1))
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v14 := GenS2{F0: -97, F1: 4294967295, F2: 0}
                            if true {
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v3 = v4
                                    var v16 int = (v12 / (((v6 >> ((985) & 63))) | 1))
                                    v17 := GenS1{F0: 347, F1: true, F2: -108}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int16 = -432
                                switch "foo bar" {
                                case "foo bar":
                                    v19 := GenS3{F0: 2, F1: 65535}
                                    _ = v19
                                case "\n\t":
                                    for fi20 := 0; fi20 < 14; fi20++ {
                                        _ = len(v5)
                                        v21 := GenS4{F0: ("" + (("a" + (("\n\t" + "") + "foo bar")) + ""))}
                                        v5[(fi13 - 167)] = false
                                        _ = fi20
                                        _ = v21
                                    }
                                case "hello":
                                    v22 := GenS2{F0: 0, F1: 0, F2: 1}
                                    _ = v22
                                }
                                _ = v18
                            }
                            for fi23 := 0; fi23 < 8; fi23++ {
                                v24 := GenS4{F0: ("a" + "\n\t")}
                                _ = fi23
                                _ = v24
                            }
                            _ = fi13
                            _ = v14
                        }
                        _ = v12
                    }
                    v1 = append(v1, (v6 - v2))
                } else {
                    var v25 float64 = 0.1
                    var v26 float64 = 0.0
                    if len(v1) > 0 {
                        v1[0] = -161
                    }
                    _ = v25
                    _ = v26
                }
                var v27 uint64 = 18446744073709551615
                if true {
                    if false {
                        v5[v6] = true
                        v3 -= v3
                    } else {
                        v27 = v27
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if true {
                        if (!false) {
                            var v28 uint64 = (v27 >> ((0) & 63))
                            _ = v28
                        } else {
                            if (!(!true)) {
                                var v29 int16 = 32767
                                _ = v29
                            } else {
                                _ = len(v1)
                            }
                            v30 := GenS1{F0: 42, F1: true, F2: 1}
                            _ = v30
                        }
                        var v31 uint32 = 0
                        if (!true) {
                            var v32 uint8 = uint8(v31)
                            _ = v32
                        }
                        _ = v31
                    } else {
                        var v33 uint64 = v27
                        _ = v33
                    }
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        switch "foo bar" {
                        case "":
                            v34 := GenS1{F0: genFn0(128), F1: false, F2: 1}
                            _ = v34
                        default:
                            v35 := v5[v6]
                            _ = v35
                        }
                    }
                } else {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        v1 = append(v1, (fi36 & 4294967295))
                        _ = fi36
                    }
                    for fi37 := 0; fi37 < 1; fi37++ {
                        switch (v6 % ((fi37) | 1)) {
                        case 128:
                            fmt.Println(v27, v27)
                        case 574:
                            var v38 int16 = 32767
                            _ = v38
                        default:
                            v3 = 42
                        }
                        v39 := GenS5{F0: int64(v27)}
                        var v40 uint = 4294967295
                        _ = fi37
                        _ = v39
                        _ = v40
                    }
                }
                _ = v27
            }
            var v41 int64 = -9223372036854775808
            var v42 float64 = 1.5
            _ = v41
            _ = v42
        }
    }
    v43 := GenS4{F0: ("" + "a")}
    v44 := GenS6{F0: 2147483647, F1: 18446744073709551615, F2: 936}
    var v45 string = "\n\t"
    v45 = ((v45 + ("" + "\n\t")) + v45)
    for fi46 := 0; fi46 < 19; fi46++ {
        var v47 int64 = int64(v4)
        _ = fi46
        _ = v47
    }
    v3 = v4
    fmt.Println(v3, v45, v3)
    var v48 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}
