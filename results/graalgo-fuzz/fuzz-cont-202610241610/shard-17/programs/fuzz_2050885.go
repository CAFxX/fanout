package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 int64; F1 uint; F2 bool }
type GenS5 struct { F0 uint32; F1 bool; F2 int16 }
type GenS6 struct { F0 int64; F1 uint32; F2 int16 }
type GenS7 struct { F0 uint32; F1 int64; F2 bool }
func genFn0(p0 uint, p1 uint8) uint {
    p0 += uint(p1)
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: ("\n\t" + "hello"), F1: ((42 > 0) || false), F2: (0 != 2)}
        if (42 < 32767) {
            v7 := []int{(fi5 ^ -813), 32768}
            _ = v7
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint32 = 2147483647
            v10 := []int{(^(fi8)), fi8}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v11 bool = ((p1 >> ((p0) & 63)) == (p1 ^ p1))
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = v11
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    var v3 int8 = -128
    var v4 int16 = 0
    v3 = (v3 + 127)
    var v12 int8 = -1
    switch "foo bar" {
    case "foo bar":
        v13 := []int{42}
        _ = v13
    default:
        var v14 int = 32768
        _ = v14
    }
    switch "hello" {
    case "hello":
        var v15 int64 = -1
        _ = v15
    case "\n\t":
        var v16 float32 = (1.5 + (float32(((-2.25 * (-2.25 * (1e10 * 3.14159))) - ((0.0 + ((0.1 + 1.5) - 3.14159)) - (((0.0 + (1.5 + 0.1)) - (0.0 * 0.1)) - ((0.1 / 1.0) * 3.14159))))) - -2.25))
        _ = v16
    case "foo bar":
        delete(v1, (("" + ("hello" + "\n\t")) + ""))
    }
    var v17 uint8 = 42
    if false {
        v18 := GenS2{F0: (v2 | 818), F1: 1}
        switch ("foo bar" + "") {
        case "a":
            fmt.Println(v2)
        case "foo bar":
            v19 := GenS2{F0: v2, F1: 634}
            _ = v19
        default:
            var v20 int64 = -427
            _ = v20
        }
        _ = v18
    } else {
        switch 32767 {
        case -2147483648:
            var v21 uint8 = (v17 << ((genFn0(356, v17)) & 63))
            _ = v21
        }
        if (-1 > v4) {
            v22 := []int{4, 127}
            v12 -= -26
            switch "" {
            case "hello":
                var v23 uint = 65535
                _ = v23
            case "foo bar":
                fmt.Println(v2, v3, v3)
            case "a":
                v24 := GenS3{F0: v2, F1: 2147483647, F2: "hello"}
                _ = v24
            default:
                var v25 int = -118
                _ = v25
            }
            _ = v22
        }
    }
    if (2 < uint32(v2)) {
        v26 := []int{-2147483648, 32768, 32767}
        var v27 float64 = (0.1 / 1.5)
        v28 := GenS4{F0: 1, F1: 358, F2: false}
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 int16 = (v4 & -265)
    if true {
        if false {
            _ = len(v1)
        } else {
            fmt.Println(v12, v29)
            var v30 string = "hello"
            _ = v30
        }
        switch "hello" {
        case "foo bar":
            if ((!(v2 > 42)) || (-2147483648 == -2147483648)) {
                v31 := GenS5{F0: 1, F1: false, F2: v4}
                _ = len(v1)
                v12 = (v3 / (((v12 - v3)) | 1))
                _ = v31
            }
        case "hello":
            var v32 int64 = -1
            _ = v32
        case "":
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 float64 = float64(0.0)
                switch (fi33 << ((255) & 63)) {
                case 0:
                    for fi35 := 0; fi35 < 15; fi35++ {
                        if true {
                            v36 := []int{}
                            v37 := GenS6{F0: -977, F1: 2, F2: -624}
                            if ((!false) || false) {
                                v12 -= 0
                                if (!(true || ((v17 - 255) < (v17 << ((172) & 63))))) {
                                    delete(v1, "hello")
                                }
                            } else {
                                if (v34 >= v34) {
                                    delete(v1, "")
                                    if true {
                                        v38 := GenS7{F0: 2, F1: -1, F2: false}
                                        _ = v38
                                    }
                                }
                                var v39 uint32 = 1
                                v12 = 63
                                _ = v39
                            }
                            _ = v36
                            _ = v37
                        } else {
                            _ = len(v1)
                            v4 = v4
                        }
                        v4 = v4
                        var v40 int8 = v3
                        _ = fi35
                        _ = v40
                    }
                default:
                    var v41 uint8 = v17
                    _ = v41
                }
                v2 = (^(v2))
                _ = fi33
                _ = v34
            }
        }
    }
    v42 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
}
