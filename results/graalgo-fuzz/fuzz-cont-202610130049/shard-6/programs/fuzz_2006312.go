package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint32, p1 uint16, p2 int) int8 {
    switch "" {
    case "a":
        var v7 int32 = 2147483647
        _ = v7
    case "":
        if false {
            fmt.Println(p2, p2)
            p0 = 320
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v10 := []int{}
                    switch (fi8 - 2147483647) {
                    case -611:
                        fmt.Println(p2)
                    default:
                        var v11 int64 = -1
                        _ = v11
                    }
                    if (!false) {
                        var v12 int8 = 127
                        var v13 uint32 = (p0 >> ((579) & 63))
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(p2, fi8)
                    }
                    _ = fi9
                    _ = v10
                }
                if false {
                    var v14 int8 = -1
                    _ = v14
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        v15 := GenS1{F0: (p0 << ((18446744073709551615) & 63))}
                        _ = v15
                    case "":
                        v16 := []int{(p2 << ((0) & 63)), 9223372036854775807}
                        _ = v16
                    default:
                        for fi17 := 0; fi17 < 3; fi17++ {
                            p2 = p2
                            _ = fi17
                        }
                    }
                }
                _ = fi8
            }
        }
    case "hello":
        if false {
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := []int{fi18}
                _ = fi18
                _ = v19
            }
        }
    default:
        fmt.Println(p1, p2)
    }
    switch "" {
    case "":
        if false {
            var v20 uint32 = 877
            switch (p2 ^ (p2 | (-(p2)))) {
            case 4294967295:
                v21 := GenS2{F0: -1, F1: (p2 << ((0) & 63)), F2: -2147483648}
                _ = v21
            default:
                if true {
                    var v22 float32 = (((0.0 * (1e10 / 1.0)) * 1.5) * 0.1)
                    fmt.Println(p1, v22)
                    _ = v22
                } else {
                    if (p1 != p1) {
                        switch p2 {
                        case 32768:
                            var v23 bool = (int(p1) != (-(p2)))
                            _ = v23
                        case 9223372036854775807:
                            fmt.Println(v20, v20, v20)
                        }
                    }
                    var v24 float64 = float64(1e10)
                    v25 := []int{}
                    _ = v24
                    _ = v25
                }
            }
            _ = v20
        } else {
            v26 := GenS1{F0: (p0 / ((0) | 1))}
            p2 += (p2 - p2)
            v27 := GenS2{F0: -736, F1: (p2 & 128), F2: -1}
            _ = v26
            _ = v27
        }
    case "\n\t":
        var v28 int64 = 42
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    var v3 bool = (123 != 335)
    var v4 string = "a"
    v5 := make(map[uint32]bool)
    var v6 uint = 680
    v29 := GenS2{F0: -935, F1: 128, F2: -1}
    v30 := GenS2{F0: 32767, F1: -1, F2: 1}
    if v3 {
        var v31 int8 = 69
        fmt.Println(v6)
        v31 = (v31 >> ((v6) & 63))
        _ = v31
    } else {
        var v32 string = ("\n\t" + ("hello" + "a"))
        _ = len(v5)
        _ = v32
    }
    var v33 float64 = (0.0 + (1e10 + 0.1))
    var v34 int = -1
    if (false && v3) {
        for fi35 := 0; fi35 < 14; fi35++ {
            v33 = 1.5
            var v36 int8 = 103
            _ = fi35
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 5; fi37++ {
            v33 = ((((v33 / v33) + (v33 - v33)) * v33) - v33)
            v38 := v5[677]
            fmt.Println(v2, v33, v6)
            _ = fi37
            _ = v38
        }
        delete(v5, 65535)
        v4 = (v4 + (v4 + ("a" + "")))
    }
    var v39 int8 = 0
    for fi40 := 0; fi40 < 7; fi40++ {
        v41 := GenS3{F0: 223, F1: 127}
        v34 -= v34
        var v42 uint = 1
        _ = fi40
        _ = v41
        _ = v42
    }
    switch ((v4 + v4) + v4) {
    case "a":
        v5[65535] = true
    case "foo bar":
        if v3 {
            switch 0 {
            case 42:
                fmt.Println(v39, v3, v39)
            default:
                v39 += v39
            }
        }
    case "hello":
        var v43 uint = 822
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
