package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int; F2 string }
type GenS5 struct { F0 uint }
func genFn0(p0 int32, p1 uint) int {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            p1 = (p1 - (^(p1)))
            switch (fi7 | fi6) {
            case 486:
                v8 := GenS1{F0: (((("a" + ("" + "a")) + ("a" + "hello")) + "foo bar") + "a"), F1: 2}
                _ = v8
            default:
                p1 += 4294967295
            }
            v9 := []int{-9223372036854775808, 551, (fi6 ^ fi7)}
            var v10 int64 = 42
            v10 = (v10 & (-(v10)))
            _ = fi7
            _ = v10
            _ = v9
        }
        p0 = (-(p0))
        _ = fi6
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int16 = -32768
        for fi13 := 0; fi13 < 4; fi13++ {
            switch (fi13 / ((fi11) | 1)) {
            case 2147483647:
                v12 += (v12 >> ((18446744073709551615) & 63))
            case 32768:
                v12 = v12
            default:
                p1 -= (p1 % (((p1 & (^(p1)))) | 1))
            }
            var v14 int = (fi13 * (fi13 & fi13))
            _ = fi13
            _ = v14
        }
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    switch "foo bar" {
    case "\n\t":
        v15 := GenS2{F0: 0}
        _ = v15
    case "foo bar":
        var v16 float32 = float32(float64(0.0))
        _ = v16
    case "":
        if (96 > 824) {
            switch "hello" {
            case "a":
                if ((1.5 <= 1e10) && (-2.25 > -2.25)) {
                    v17 := []int{}
                    _ = v17
                }
            default:
                p1 += p1
            }
            if true {
                var v18 uint32 = 0
                _ = v18
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                p0 += p0
                _ = fi19
            }
        }
    default:
        v20 := []int{}
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS2{F0: -612}
        p1 = p1
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    return 123
}
func genFn1(p0 uint64, p1 string, p2 uint8) uint32 {
    var v23 int32 = 2147483647
    var v24 int64 = 42
    var v25 uint32 = 65535
    v26 := GenS1{F0: p1, F1: 1}
    v27 := []int{0, 4294967295, int(v25), 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    return 0
}

func main() {
    var v1 int32 = -130
    v2 := make(map[int]int)
    v3 := []uint32{65535}
    v4 := make(map[int]int)
    v5 := []string{}
    v28 := GenS3{F0: 18446744073709551615}
    v1 = (v1 + -990)
    v1 -= v1
    if (0 == 255) {
        v29 := GenS3{F0: 4294967295}
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 uint32 = 0
            v32 := GenS2{F0: -32768}
            _ = fi30
            _ = v31
            _ = v32
        }
        v1 = v1
        _ = v29
    } else {
        v33 := GenS4{F0: 66, F1: 2147483647, F2: "\n\t"}
        if (!false) {
            var v34 uint = 18446744073709551615
            _ = len(v5)
            var v35 uint64 = 4294967295
            _ = v34
            _ = v35
        } else {
            var v36 int = 42
            _ = v36
        }
        fmt.Println(v1)
        _ = v33
    }
    v1 += v1
    if false {
        for fi37 := 0; fi37 < 15; fi37++ {
            for fi38 := 0; fi38 < 14; fi38++ {
                v1 += v1
                if true {
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                    v39 := GenS4{F0: 271, F1: fi37, F2: ("foo bar" + "foo bar")}
                    _ = v39
                }
                _ = fi38
            }
            _ = fi37
        }
        delete(v2, -919)
    } else {
        v40 := GenS5{F0: 137}
        _ = v40
    }
    var v41 uint64 = 2
    v42 := GenS2{F0: 32767}
    for fi43 := 0; fi43 < 13; fi43++ {
        if (true || true) {
            v5 = append(v5, "")
        }
        _ = fi43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        if false {
            v41 -= (v41 & 730)
            v1 -= v1
            if ((^(v41)) > (v41 / ((v41) | 1))) {
                var v45 int = fi44
                if (false || true) {
                    for fi46 := 0; fi46 < 9; fi46++ {
                        var v47 int8 = 70
                        _ = fi46
                        _ = v47
                    }
                }
                _ = v45
            }
        } else {
            v41 += v41
        }
        _ = fi44
    }
    if (3.14159 == 0.0) {
        v1 -= v1
        var v48 int32 = v1
        switch "hello" {
        case "":
            var v49 int32 = (v48 * v48)
            _ = v49
        case "\n\t":
            for fi50 := 0; fi50 < 14; fi50++ {
                var v51 uint8 = 137
                var v52 int = fi50
                fmt.Println(v48, v41)
                _ = fi50
                _ = v51
                _ = v52
            }
        default:
            switch -415 {
            case 32768:
                v53 := GenS5{F0: 1}
                _ = v53
            case 0:
                var v54 uint8 = 124
                _ = v54
            }
        }
        _ = v48
    } else {
        if false {
            v2[634] = 9223372036854775807
        }
        v55 := GenS5{F0: 962}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
