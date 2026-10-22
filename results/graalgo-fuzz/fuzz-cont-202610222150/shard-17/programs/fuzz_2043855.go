package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 string, p1 string, p2 int16) int64 {
    p2 -= (p2 * -32768)
    p1 = p1
    v6 := []int{-521}
    v6 = append(v6, 2147483647)
    v6 = append(v6, 2147483647)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := []int{42, -9223372036854775808}
    v5 := make(map[string]uint32)
    v7 := GenS1{F0: -13}
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 1}
        var v11 uint64 = 2
        v11 += 18446744073709551615
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if (0 < 2) {
            var v13 int16 = -275
            _ = v13
        } else {
            v8 -= v8
            if (-128 > 100) {
                v8 += v8
                var v14 bool = (((v8 - v8) < (^(v8))) && false)
                _ = v14
            } else {
                var v15 int16 = 1
                if true {
                    var v16 uint8 = (v8 + 42)
                    v17 := GenS1{F0: 210}
                    var v18 string = ""
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
            if (!false) {
                v8 += 2
                if false {
                    var v19 string = "a"
                    v20 := GenS2{F0: -28, F1: -32768}
                    _ = v19
                    _ = v20
                } else {
                    if true {
                        v8 += 42
                        v8 -= v8
                        fmt.Println(fi12, v8)
                    }
                    v8 += (v8 << ((2) & 63))
                    v8 += (^(v8))
                }
            }
        }
        var v21 uint32 = 949
        fmt.Println(v21, v8)
        _ = fi12
        _ = v21
    }
    var v22 uint64 = 0
    var v23 uint16 = 65535
    v8 += 42
    if ("a" >= "foo bar") {
        if (42 == -2147483648) {
            v8 += (v8 >> ((2) & 63))
            var v24 int16 = -134
            fmt.Println(v22)
            _ = v24
        } else {
            _ = len(v4)
        }
        if (true && (!true)) {
            if false {
                v8 += v8
                v25 := GenS1{F0: 625}
                _ = v25
            }
            var v26 int64 = 812
            _ = v26
        }
        v27 := v2[0]
        _ = v27
    } else {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v8)
        case "hello":
            switch -543 {
            case 4294967295:
                for fi28 := 0; fi28 < 19; fi28++ {
                    delete(v2, 467)
                    var v29 int16 = 0
                    _ = fi28
                    _ = v29
                }
            }
        case "foo bar":
            v8 = (^(v8))
        }
        switch "hello" {
        case "hello":
            switch "" {
            case "":
                delete(v5, "")
            case "a":
                fmt.Println(v8)
            case "hello":
                if (0 == 452) {
                    v22 = v22
                } else {
                    if true {
                        if ((v22 >> ((4294967295) & 63)) != (v22 / ((1) | 1))) {
                            v23 = (^(v23))
                        }
                    }
                }
            default:
                var v30 uint8 = 127
                _ = v30
            }
        case "foo bar":
            var v31 string = ("foo bar" + ("foo bar" + "a"))
            _ = v31
        case "":
            v23 += v23
        }
    }
    v22 -= (v22 + v22)
    var v32 uint8 = (v8 ^ 108)
    if false {
        v32 = v8
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        v22 = v22
        _ = fi33
    }
    fmt.Println(v32)
    if (1 <= 896) {
        v34 := GenS2{F0: 100, F1: -1}
        var v35 float64 = 0.0
        v22 = v22
        _ = v34
        _ = v35
    }
    var v36 uint64 = (v22 | 1)
    v36 = v22
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
