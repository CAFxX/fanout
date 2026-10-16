package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int16, p1 int8) int {
    if false {
        p1 = (p1 << ((140) & 63))
        for fi6 := 0; fi6 < 13; fi6++ {
            p0 = (p0 + (p0 | (p0 & (p0 >> ((42) & 63)))))
            _ = fi6
        }
        v7 := GenS1{F0: 197, F1: 1}
        _ = v7
    }
    fmt.Println(p0, p1)
    var v8 int64 = 1
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int = 128
    var v2 uint8 = 127
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{42, 42, 598}
    if true {
        if false {
            var v10 bool = true
            _ = len(v4)
            v2 = v2
            _ = v10
        } else {
            v11 := GenS1{F0: 252, F1: 1}
            if false {
                v5 = append(v5, 2)
                if len(v5) > 0 {
                    v5[0] = 1
                }
            } else {
                v3[467] = "foo bar"
            }
            _ = v11
        }
    }
    if (2147483647 != 42) {
        switch "\n\t" {
        case "a":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 float64 = 0.1
                v14 := GenS1{F0: (v2 | v2), F1: v1}
                var v15 uint32 = 1
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case "\n\t":
            var v16 float32 = 1e10
            _ = v16
        case "foo bar":
            if (!((0.1 != 3.14159) && ((v1 >> ((4294967295) & 63)) < (v1 ^ (^(v1)))))) {
                if (((0 >= 255) || true) && (false || true)) {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v2 -= 205
                        _ = fi17
                    }
                    v18 := GenS1{F0: 0, F1: (v1 * 42)}
                    _ = v18
                }
                v19 := GenS1{F0: (^(v2)), F1: genFn0(32767, 42)}
                v1 -= (-(v1))
                _ = v19
            }
        }
        fmt.Println(v1, v2)
        var v20 int8 = 106
        _ = v20
    } else {
        var v21 uint8 = 1
        v22 := GenS2{F0: 42, F1: 42}
        if (!true) {
            v23 := GenS2{F0: 42, F1: 546}
            if (4294967295 != 4294967295) {
                var v24 uint16 = 1
                fmt.Println(v21, v1)
                _ = v24
            } else {
                var v25 float32 = float32(1.5)
                var v26 bool = true
                for fi27 := 0; fi27 < 13; fi27++ {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        switch (v1 | (fi28 ^ (fi27 | (fi27 & v1)))) {
                        case 4294967295:
                            if (!(-1 >= 2147483647)) {
                                var v29 uint32 = 17
                                v25 = 1.5
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v29, v29, v29)
                                }
                                _ = v29
                            } else {
                                v30 := GenS1{F0: (v21 - v21), F1: (v1 >> ((4294967295) & 63))}
                                _ = v30
                            }
                        case 1:
                            fmt.Println(v25, v21)
                            v31 := GenS3{F0: 2147483647, F1: 42}
                            v32 := GenS3{F0: 2147483648, F1: -128}
                            _ = v31
                            _ = v32
                        case 42:
                            v33 := v4[65535]
                            _ = v33
                        }
                        _ = fi28
                    }
                    fmt.Println(v26, v26)
                    _ = fi27
                }
                _ = v25
                _ = v26
            }
            _ = v23
        }
        _ = v21
        _ = v22
    }
    v1 = (v1 >> ((0) & 63))
    var v34 uint = 311
    v5 = append(v5, 2147483647)
    var v35 int = (v1 + 32767)
    switch "foo bar" {
    case "hello":
        v36 := GenS3{F0: 42, F1: 127}
        _ = v36
    default:
        v37 := GenS2{F0: -1, F1: -51}
        _ = v37
    }
    v38 := GenS3{F0: 195, F1: 72}
    if (false && false) {
        v1 = (v1 ^ v1)
        v39 := GenS2{F0: int64(v35), F1: 32767}
        _ = v39
    } else {
        v3[1] = ("a" + "hello")
    }
    v40 := GenS4{F0: v34, F1: 203}
    var v41 int = (v1 - genFn0(-380, -86))
    fmt.Println(v2)
    if ((!false) && true) {
        v4[0] = 2147483647
    } else {
        fmt.Println(v34)
        var v42 uint = (^(v34))
        _ = v42
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v43 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v5
}
