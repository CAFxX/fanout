package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint8; F1 string; F2 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 int32, p1 uint32, p2 uint8) bool {
    var v6 float32 = ((1e10 - (((float32((-2.25 - 1.5)) * 0.0) + 3.14159) / (1.5 * -2.25))) * 0.1)
    if (!false) {
        for fi7 := 0; fi7 < 7; fi7++ {
            if (false && true) {
                switch 0 {
                case -453:
                    v8 := []int{fi7, int(p1), 2147483647, fi7}
                    _ = v8
                }
                if false {
                    v9 := []int{1, -934, 32768}
                    _ = v9
                }
                v6 = v6
            } else {
                v10 := GenS1{F0: "\n\t"}
                v11 := GenS2{F0: 111, F1: ((false || (false || true)) && ((true && false) || true))}
                if true {
                    v12 := []int{fi7, (fi7 ^ fi7), (fi7 & fi7)}
                    if (1.5 > (v6 * (1e10 + (3.14159 * v6)))) {
                        var v13 uint64 = 1
                        _ = v13
                    }
                    _ = v12
                } else {
                    fmt.Println(fi7, p2, p1)
                }
                _ = v10
                _ = v11
            }
            _ = fi7
        }
        p2 -= p2
    } else {
        var v14 string = ""
        p2 -= (p2 >> ((828) & 63))
        v15 := GenS2{F0: -28, F1: true}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}
func genFn1(p0 uint32, p1 int64) string {
    v16 := GenS3{F0: 42, F1: "foo bar", F2: 503}
    var v17 uint = 2
    var v18 int8 = -79
    if (true || true) {
        fmt.Println(p1, v18, v17)
        fmt.Println(v17, p0)
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := []int{(fi19 % ((fi19) | 1)), fi19, 1}
            if (false && false) {
                v18 = v18
                v21 := GenS4{F0: 9223372036854775807, F1: (p1 & (p1 % ((p1) | 1))), F2: (p1 ^ (-(p1)))}
                _ = len(v20)
                _ = v21
            }
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := GenS5{F0: 0}
                fmt.Println(v18, p0)
                _ = fi22
                _ = v23
            }
            _ = fi19
            _ = v20
        }
    } else {
        var v24 string = ("hello" + "a")
        var v25 int8 = 1
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return ("foo bar" + "hello")
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 148
    var v3 string = "a"
    v4 := make(map[int]uint32)
    var v5 int = 4294967295
    fmt.Println(v3, v1)
    var v26 string = ((v3 + ("a" + ("a" + ("a" + (v3 + v3))))) + (v3 + "a"))
    var v27 uint16 = 42
    v1 += v1
    v26 = v3
    switch "a" {
    case "a":
        var v28 int = 279
        _ = v28
    default:
        delete(v4, v5)
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(v2)
        var v30 int64 = 1
        switch v26 {
        case "hello":
            var v31 int8 = v1
            _ = v31
        case "":
            var v32 bool = false
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS6{F0: 18446744073709551615}
        _ = fi33
        _ = v34
    }
    var v35 int8 = (v1 & 100)
    var v36 int32 = 1
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 float64 = 3.14159
        var v39 int64 = 1
        v40 := v4[(v5 >> ((1) & 63))]
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    v41 := []int{-1, (-(v5))}
    v42 := GenS5{F0: 27}
    if (1 <= 0) {
        v43 := GenS7{F0: (!true)}
        _ = v43
    } else {
        var v44 uint64 = 1
        var v45 float64 = (1.5 * 0.1)
        var v46 uint16 = 0
        _ = v44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 8; fi47++ {
        v48 := GenS6{F0: 18446744073709551615}
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
