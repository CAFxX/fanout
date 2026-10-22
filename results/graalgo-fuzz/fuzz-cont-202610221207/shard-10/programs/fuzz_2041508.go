package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 uint, p1 uint8, p2 int) int64 {
    p1 = p1
    var v5 int8 = 57
    v6 := []int{2147483647, p2, (^(p2)), p2}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return int64(p1)
}

func main() {
    var v1 uint = 567
    v2 := make(map[int]int)
    var v3 int32 = 42
    var v4 bool = (!(true && true))
    switch "foo bar" {
    case "\n\t":
        var v7 uint64 = 275
        _ = v7
    case "foo bar":
        var v8 int = -386
        _ = v8
    case "hello":
        var v9 int16 = -1
        _ = v9
    default:
        var v10 bool = (1000 >= -1)
        _ = v10
    }
    v11 := []int{4294967295, 32768, 479, 42}
    switch "hello" {
    case "hello":
        switch 697 {
        case -170:
            delete(v2, 1)
        case 4294967295:
            var v12 int8 = 0
            _ = v12
        case 127:
            for fi13 := 0; fi13 < 15; fi13++ {
                if ("a" < (("" + "\n\t") + "hello")) {
                    if false {
                        fmt.Println(v1)
                        for fi14 := 0; fi14 < 17; fi14++ {
                            if v4 {
                                var v15 bool = (!v4)
                                _ = v15
                            } else {
                                v4 = (296 < 2147483647)
                                var v16 string = "a"
                                v17 := GenS1{F0: 795}
                                _ = v16
                                _ = v17
                            }
                            _ = fi14
                        }
                    }
                    var v18 bool = false
                    fmt.Println(v1, v3)
                    _ = v18
                } else {
                    if (0 == 0) {
                        fmt.Println(v3)
                        v19 := GenS1{F0: 275}
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 string = ""
                        fmt.Println(fi20, fi13)
                        v21 = v21
                        _ = fi20
                        _ = v21
                    }
                    if v4 {
                        v3 += v3
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v4, v4)
                        default:
                            if (!(!v4)) {
                                v3 -= (^(v3))
                                var v22 uint16 = 42
                                v23 := GenS1{F0: 442}
                                _ = v22
                                _ = v23
                            } else {
                                _ = len(v2)
                                var v24 int = fi13
                                _ = v24
                            }
                        }
                    }
                }
                _ = fi13
            }
        }
    case "a":
        v3 = (v3 * v3)
    }
    fmt.Println(v3, v1, v3)
    fmt.Println(v3, v4)
    var v25 uint64 = 221
    for fi26 := 0; fi26 < 18; fi26++ {
        v4 = true
        _ = fi26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint8 = 0
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int16 = int16(v3)
            var v31 float64 = 0.0
            delete(v2, fi27)
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 float64 = -2.25
        _ = fi27
        _ = v28
        _ = v32
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        if len(v11) > 0 {
            v11[0] = 931
        }
        _ = fi33
    }
    var v34 uint = v1
    if v4 {
        if ((v25 >= 400) && v4) {
            delete(v2, 32767)
        } else {
            if v4 {
                v11 = append(v11, 9223372036854775807)
                for fi35 := 0; fi35 < 18; fi35++ {
                    v3 = v3
                    if (802 <= 0) {
                        if (((1e10 - (1e10 + 3.14159)) - 1.5) <= ((-2.25 / 1.5) / ((0.1 - 0.0) - -2.25))) {
                            fmt.Println(fi35)
                            v36 := GenS1{F0: 2147483647}
                            fmt.Println(v34, v25, fi35)
                            _ = v36
                        }
                    }
                    var v37 int16 = -351
                    _ = fi35
                    _ = v37
                }
            }
        }
    }
    v34 -= (v34 & (v34 ^ v34))
    var v38 int16 = 1000
    switch 2147483647 {
    case -1:
        var v39 int64 = 9223372036854775807
        _ = v39
    case 32768:
        v1 = v1
    default:
        v38 = (v38 << (((v1 * v34)) & 63))
    }
    switch ("a" + "hello") {
    case "hello":
        v40 := GenS2{F0: 18446744073709551615, F1: v3}
        _ = v40
    case "foo bar":
        for fi41 := 0; fi41 < 7; fi41++ {
            fmt.Println(fi41, v4)
            _ = fi41
        }
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v38
    _ = v4
}
