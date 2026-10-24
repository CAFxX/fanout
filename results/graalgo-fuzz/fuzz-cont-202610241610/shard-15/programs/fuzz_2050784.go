package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 string; F1 int64; F2 string }
func genFn0(p0 uint32, p1 uint64) uint {
    if false {
        switch "foo bar" {
        case "":
            switch -1 {
            case 127:
                for fi7 := 0; fi7 < 16; fi7++ {
                    if false {
                        v8 := GenS1{F0: (p0 | (p0 | p0)), F1: 127, F2: true}
                        p0 -= p0
                        _ = v8
                    }
                    p1 += p1
                    fmt.Println(p1)
                    _ = fi7
                }
            default:
                fmt.Println(p1)
            }
        case "a":
            v9 := []int{-669, -9223372036854775808}
            _ = v9
        case "foo bar":
            fmt.Println(p0)
        }
    } else {
        p0 -= (p0 + (p0 ^ p0))
        var v10 bool = (true || ((1 <= 32767) && (2 < uint(p1))))
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS2{F0: 177, F1: "foo bar"}
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := []int{(fi13 + fi13)}
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    p1 = (p1 >> ((2) & 63))
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int = (-(fi15))
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 float64 = (0.1 + (3.14159 / 1.5))
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := GenS2{F0: 32767, F1: ""}
    _ = p0
    _ = p1
    _ = v19
    return 0
}
func genFn1() uint32 {
    var v20 uint8 = 42
    var v21 int = -1
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int64 = 0
            v20 += 127
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    _ = v20
    _ = v21
    return 657
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    var v4 bool = true
    var v5 uint8 = 42
    v6 := make(map[int]string)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int = fi25
        if v4 {
            if false {
                v27 := GenS1{F0: 325, F1: 0, F2: v4}
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 int64 = 9223372036854775807
                    v30 := []int{1, 32767, (^(fi28)), v26}
                    _ = len(v30)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
        } else {
            var v31 uint64 = 18446744073709551615
            v32 := GenS1{F0: v3, F1: 127, F2: (!true)}
            _ = v31
            _ = v32
        }
        v33 := GenS1{F0: genFn1(), F1: v5, F2: (!(("a" + "a") < "a"))}
        _ = fi25
        _ = v26
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        switch "" {
        case "a":
            _ = len(v2)
        case "":
            v35 := GenS3{F0: 765, F1: 137}
            _ = v35
        default:
            var v36 uint8 = (v5 & (v5 * (^(v5))))
            _ = v36
        }
        switch ("a" + "a") {
        case "foo bar":
            fmt.Println(v3, v1, v5)
        case "hello":
            var v37 float64 = v1
            _ = v37
        case "\n\t":
            if (v4 || (!(-477 <= -758))) {
                var v38 uint = 1
                var v39 int16 = 0
                for fi40 := 0; fi40 < 5; fi40++ {
                    v41 := GenS4{F0: "a", F1: 0, F2: ("a" + "a")}
                    if (v39 < (v39 / (((v39 | v39)) | 1))) {
                        v1 = (((v1 / 0.1) - v1) * float64(-2.25))
                        fmt.Println(fi40, v39)
                        fmt.Println(v4)
                    } else {
                        var v42 float32 = -2.25
                        _ = v42
                    }
                    var v43 int32 = 194
                    _ = fi40
                    _ = v41
                    _ = v43
                }
                _ = v38
                _ = v39
            } else {
                for fi44 := 0; fi44 < 15; fi44++ {
                    v3 -= 0
                    fmt.Println(v1)
                    _ = fi44
                }
            }
        default:
            v4 = v4
        }
        _ = fi34
    }
    v4 = false
    fmt.Println(v5)
    v3 = v3
    v5 = v5
    for fi45 := 0; fi45 < 6; fi45++ {
        for fi46 := 0; fi46 < 16; fi46++ {
            var v47 uint64 = 42
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    var v48 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
