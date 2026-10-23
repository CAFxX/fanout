package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 int16, p2 int32) uint32 {
    p1 += 1000
    var v6 int64 = 1
    var v7 int = 938
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{int(v6), v7, fi8, fi8}
        v9 = append(v9, v7)
        v10 := GenS1{F0: ((1.5 * ((-2.25 * 0.1) * 0.0)) == 1e10), F1: 2}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 string = ""
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    var v5 bool = true
    v5 = v5
    fmt.Println(v3, v3, v5)
    if true {
        var v11 uint16 = (v3 + (v3 % ((v3) | 1)))
        v11 = 2
        var v12 uint16 = (v3 ^ 2)
        _ = v11
        _ = v12
    } else {
        var v13 uint8 = 2
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int = fi14
            for fi16 := 0; fi16 < 6; fi16++ {
                v13 -= (v13 | 42)
                _ = fi16
            }
            v17 := GenS2{F0: 995}
            _ = fi14
            _ = v15
            _ = v17
        }
        _ = v13
    }
    if (true && false) {
        if v5 {
            var v18 int = 9223372036854775807
            var v19 bool = (711 == 1)
            _ = v18
            _ = v19
        } else {
            fmt.Println(v3, v3)
            v2 = v2
        }
        var v20 int8 = -128
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        _ = len(v1)
        if v5 {
            v3 -= v3
            if v5 {
                var v22 int = 127
                _ = v22
            } else {
                v23 := GenS3{F0: 65535, F1: 381, F2: 0}
                if v5 {
                    var v24 int = (fi21 ^ 4294967295)
                    _ = v24
                } else {
                    v5 = (-1 == 42)
                }
                _ = v23
            }
        }
        _ = fi21
    }
    var v25 int16 = -32768
    v4[-525] = 608
    v26 := GenS4{F0: (v5 && false)}
    _ = len(v1)
    v2 = v2
    if (!(-1 == 2147483647)) {
        if v5 {
            v5 = v5
            var v27 uint64 = 7
            switch 42 {
            case -390:
                _ = len(v4)
            }
            _ = v27
        }
    }
    v28 := GenS5{F0: ((-1 >= -66) || ((-276 <= -27) || v5)), F1: 1}
    var v29 int32 = -1
    v30 := GenS5{F0: (!(v3 < 255)), F1: 242}
    var v31 string = v2
    switch (v2 + (v2 + (("foo bar" + v2) + ("" + "")))) {
    case "\n\t":
        if false {
            var v32 bool = (((-2.25 * 1e10) + -2.25) != 0.0)
            fmt.Println(v3, v2, v2)
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 string = (v31 + "a")
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
    case "":
        var v35 uint8 = 145
        _ = v35
    case "a":
        var v36 uint32 = genFn0(uint8(v3), (v25 ^ v25), (v29 - 42))
        _ = v36
    default:
        if v5 {
            switch "foo bar" {
            case "a":
                v5 = v5
            case "\n\t":
                if (!(!(v5 || false))) {
                    var v37 int = -2147483648
                    v5 = (2147483648 < 935)
                    _ = v37
                }
            case "foo bar":
                v5 = v5
            }
            if (255 < 255) {
                v5 = (true || true)
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                v39 := GenS3{F0: 209, F1: v25, F2: -29}
                _ = fi38
                _ = v39
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
