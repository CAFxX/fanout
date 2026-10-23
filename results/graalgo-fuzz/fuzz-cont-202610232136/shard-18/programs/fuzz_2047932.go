package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 string; F2 int64 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint64 = 46
    var v2 uint16 = 0
    v3 := []uint32{4294967295, uint32(v2), 806}
    var v4 int64 = 0
    var v5 int = -691
    v6 := make(map[string]int)
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 6; fi7++ {
            v5 = (fi7 >> ((1) & 63))
            if (0 >= 100) {
                v2 = v2
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(fi7)
                    _ = fi8
                }
            }
            _ = fi7
        }
    case "a":
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: v1}
            v6[((("foo bar" + "foo bar") + "a") + "hello")] = (v5 | -1)
            v1 += 42
            _ = fi9
            _ = v10
        }
    default:
        var v11 uint32 = uint32(v2)
        _ = v11
    }
    fmt.Println(v1, v2)
    v2 -= (^(v2))
    v1 += v1
    _ = len(v3)
    var v12 int8 = 0
    var v13 uint64 = v1
    if true {
        if (!((v12 / ((v12) | 1)) == v12)) {
            delete(v6, "hello")
        } else {
            fmt.Println(v2)
            var v14 string = "hello"
            v1 += (v1 + 0)
            _ = v14
        }
        if ((v12 + 100) <= v12) {
            switch v5 {
            case 42:
                if (86 != 674) {
                    v5 = (v5 >> ((255) & 63))
                    var v15 float32 = ((1.5 + 1e10) + ((3.14159 * 1e10) - float32(3.14159)))
                    _ = v15
                } else {
                    var v16 uint32 = 403
                    v12 -= int8(v13)
                    _ = v16
                }
            }
        } else {
            v17 := GenS1{F0: v1}
            switch "a" {
            case "foo bar":
                var v18 string = ("hello" + "foo bar")
                _ = v18
            case "hello":
                var v19 uint32 = 65535
                _ = v19
            case "":
                fmt.Println(v1, v4)
            }
            var v20 int = (v5 & v5)
            _ = v17
            _ = v20
        }
    }
    if true {
        if ((float64(-2.25) <= (3.14159 + -2.25)) || false) {
            if ((v1 << ((2) & 63)) < 0) {
                if (127 == 1) {
                    v21 := GenS2{F0: v1, F1: true, F2: 464}
                    _ = v21
                } else {
                    fmt.Println(v5)
                }
                v22 := GenS3{F0: v5}
                for fi23 := 0; fi23 < 11; fi23++ {
                    if (!((fi23 % ((-2147483648) | 1)) < (^(v5)))) {
                        if false {
                            v24 := GenS2{F0: v13, F1: ((70 >= 83) && ((v4 / (((v4 << ((4294967295) & 63))) | 1)) <= v4)), F2: 18446744073709551615}
                            _ = len(v3)
                            _ = v24
                        }
                        delete(v6, (("foo bar" + "hello") + ""))
                        var v25 uint16 = (^(v2))
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v27 := GenS1{F0: v13}
                        if true {
                            fmt.Println(fi26)
                            fmt.Println(v2, v2, v2)
                            v28 := GenS4{F0: 42, F1: "", F2: v4}
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = fi23
                }
                _ = v22
            } else {
                _ = len(v6)
            }
            v2 += v2
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 int16 = -32768
                for fi31 := 0; fi31 < 19; fi31++ {
                    v32 := GenS3{F0: 9223372036854775807}
                    v33 := GenS1{F0: (v13 * 504)}
                    v5 -= 9223372036854775807
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = len(v3)
                _ = fi29
                _ = v30
            }
        } else {
            v34 := GenS5{F0: 255, F1: 1, F2: 65535}
            fmt.Println(v2, v1)
            _ = v34
        }
    } else {
        v5 = (v5 >> ((811) & 63))
        var v35 uint8 = 128
        _ = v35
    }
    fmt.Println(v13)
    v36 := GenS6{F0: (-1 != 42)}
    var v37 float32 = float32(-2.25)
    for fi38 := 0; fi38 < 20; fi38++ {
        v2 -= 1
        v37 -= (v37 - ((v37 + ((v37 / v37) - -2.25)) * v37))
        _ = fi38
    }
    v13 += (^(v1))
    var v39 float32 = float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
