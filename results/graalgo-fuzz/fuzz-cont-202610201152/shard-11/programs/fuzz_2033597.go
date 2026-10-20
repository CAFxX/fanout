package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int16, p1 int) int32 {
    v4 := []int{2147483647}
    var v5 uint64 = 18446744073709551615
    var v6 bool = false
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{240, 42, 161}
    v3 := make(map[uint32]int)
    for fi7 := 0; fi7 < 11; fi7++ {
        if len(v2) > 0 {
            v2[0] = 42
        }
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: 42}
            _ = fi8
            _ = v9
        }
        _ = len(v3)
        _ = fi7
    }
    var v10 int16 = 1
    var v11 int64 = (v1 + v1)
    var v12 float64 = (-2.25 - 0.1)
    fmt.Println(v1)
    v13 := v3[359]
    if false {
        switch ("" + "foo bar") {
        case "a":
            if len(v2) > 0 {
                v2[0] = 2
            }
        case "\n\t":
            if false {
                if true {
                    var v14 int16 = (v10 / ((631) | 1))
                    _ = v14
                }
            } else {
                v15 := GenS2{F0: (v11 != -958), F1: false, F2: v10}
                if (527 > v10) {
                    if ((true || (true || (-1 > 0))) && (((v12 + v12) / (v12 + -2.25)) == v12)) {
                        if true {
                            v13 = 42
                            var v16 uint = 18446744073709551615
                            _ = v16
                        } else {
                            var v17 float64 = ((v12 * float64(-2.25)) / v12)
                            if (!(-1 == -2147483648)) {
                                v3[4294967295] = v13
                                v10 -= (v10 | v10)
                                v18 := GenS2{F0: true, F1: (!false), F2: (v10 & (-(v10)))}
                                _ = v18
                            } else {
                                var v19 uint16 = 2
                                _ = v19
                            }
                            _ = v17
                        }
                    } else {
                        v20 := GenS1{F0: 0}
                        v21 := GenS1{F0: 65535}
                        v10 -= v10
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v13)
                    v1 -= (v1 | v11)
                }
                _ = v15
            }
        }
        var v22 int16 = v10
        _ = v22
    }
    var v23 int8 = -1
    fmt.Println(v11, v12, v13)
    fmt.Println(v23, v23)
    v24 := GenS3{F0: 2147483648, F1: (-(v23))}
    if true {
        fmt.Println(v11)
        var v25 float64 = 0.0
        switch -1 {
        case 4294967295:
            v26 := GenS4{F0: -333}
            _ = v26
        case 127:
            var v27 int8 = v23
            _ = v27
        case 955:
            var v28 uint64 = 4294967295
            _ = v28
        default:
            v29 := GenS4{F0: genFn0((v10 << ((255) & 63)), v13)}
            _ = v29
        }
        _ = v25
    } else {
        var v30 uint32 = 2147483648
        v11 = -9223372036854775808
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v11 -= (v1 >> ((255) & 63))
        fmt.Println(v13)
        fmt.Println(v1, fi31, v23)
        _ = fi31
    }
    v2 = append(v2, 4294967295)
    var v32 int32 = -2147483648
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 int64 = -9223372036854775808
        _ = fi33
        _ = v34
    }
    switch "foo bar" {
    case "":
        v11 = -930
    default:
        var v35 int32 = (v32 / ((v32) | 1))
        _ = v35
    }
    switch "a" {
    case "\n\t":
        for fi36 := 0; fi36 < 20; fi36++ {
            if true {
                var v37 int = (fi36 | fi36)
                for fi38 := 0; fi38 < 13; fi38++ {
                    switch ("a" + "foo bar") {
                    case "a":
                        if true {
                            var v39 int8 = v23
                            if false {
                                v13 -= -1
                                var v40 uint64 = 1
                                switch ("a" + "") {
                                case "\n\t":
                                    var v41 uint = 739
                                    _ = v41
                                }
                                _ = v40
                            }
                            var v42 uint8 = 127
                            _ = v39
                            _ = v42
                        }
                        var v43 int16 = v10
                        _ = v43
                    case "":
                        fmt.Println(v32, fi38)
                    }
                    v44 := GenS3{F0: 2, F1: v23}
                    var v45 string = "hello"
                    _ = fi38
                    _ = v44
                    _ = v45
                }
                v46 := GenS4{F0: v32}
                _ = v37
                _ = v46
            } else {
                var v47 uint16 = 32767
                for fi48 := 0; fi48 < 12; fi48++ {
                    var v49 uint64 = 0
                    v50 := v3[4294967295]
                    var v51 float32 = float32(0.1)
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = v47
            }
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
}
