package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 int, p1 uint32, p2 uint8) int64 {
    if (p0 < (p0 << ((729) & 63))) {
        v6 := GenS1{F0: (p0 ^ p0)}
        if true {
            if (true && ((-2.25 - ((3.14159 / 0.1) / 0.1)) > 1.5)) {
                var v7 uint8 = (p2 | 2)
                _ = v7
            } else {
                var v8 int8 = 100
                v9 := GenS1{F0: (p0 % (((p0 / (((p0 - p0)) | 1))) | 1))}
                _ = v8
                _ = v9
            }
            var v10 int16 = 1000
            _ = v10
        }
        _ = v6
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int64 = 0
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return -179
}
func genFn1(p0 uint16, p1 int8, p2 int8) int64 {
    if true {
        v13 := GenS2{F0: 122, F1: 42}
        v14 := []int{730, 0, -31}
        p2 = 1
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        switch (fi15 ^ fi15) {
        case 42:
            p2 += p1
        case -9223372036854775808:
            p2 = int8(fi15)
        case 32768:
            v16 := []int{}
            _ = v16
        default:
            p2 += -110
        }
        _ = fi15
    }
    switch "\n\t" {
    case "":
        v17 := []int{int(p2), 2147483647, 1}
        _ = v17
    default:
        p0 += (p0 * p0)
    }
    if (false || (!false)) {
        var v18 int8 = p1
        _ = v18
    } else {
        if (true && ((true && true) || false)) {
            for fi19 := 0; fi19 < 12; fi19++ {
                if (!true) {
                    fmt.Println(p1, p1, p2)
                }
                var v20 uint16 = p0
                _ = fi19
                _ = v20
            }
            p1 -= (p1 * p1)
            for fi21 := 0; fi21 < 17; fi21++ {
                switch fi21 {
                case 799:
                    var v22 float64 = 0.1
                    _ = v22
                case 1:
                    v23 := GenS3{F0: -1, F1: (fi21 % ((fi21) | 1)), F2: p0}
                    _ = v23
                default:
                    p0 -= (p0 % ((p0) | 1))
                }
                var v24 int8 = (p2 * (p1 | p2))
                _ = fi21
                _ = v24
            }
        } else {
            var v25 float32 = 1.5
            for fi26 := 0; fi26 < 10; fi26++ {
                v25 -= 1e10
                p1 += -42
                if ((p0 + p0) > (p0 + 1)) {
                    p0 = p0
                }
                _ = fi26
            }
            _ = v25
        }
        v27 := GenS4{F0: 32768, F1: 128, F2: (false && (786 < 65535))}
        v28 := GenS3{F0: -693, F1: 0, F2: p0}
        _ = v27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    var v4 float64 = 1.5
    var v5 int8 = v1
    var v29 int = 32767
    switch "foo bar" {
    case "a":
        var v30 int64 = 9223372036854775807
        _ = v30
    default:
        for fi31 := 0; fi31 < 11; fi31++ {
            v4 -= v4
            var v32 uint16 = 1
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS5{F0: (v29 | 20)}
    v34 := []int{(v29 & 0), 2147483647, -1, v29}
    switch v29 {
    case 770:
        var v35 int = 42
        _ = v35
    }
    var v36 float64 = (v4 * float64((1.5 * 0.0)))
    var v37 uint64 = 485
    var v38 uint32 = 841
    v4 = v36
    switch ("foo bar" + ("hello" + ("hello" + ""))) {
    case "":
        for fi39 := 0; fi39 < 19; fi39++ {
            v1 += (v5 ^ (v1 | (-(v1))))
            if true {
                v3 -= v37
            } else {
                var v40 float64 = v4
                _ = v40
            }
            _ = fi39
        }
    case "a":
        var v41 int16 = -32768
        _ = v41
    }
    _ = len(v2)
    v42 := GenS6{F0: (v5 ^ v5), F1: (!false), F2: 1}
    v4 -= (1e10 / ((v36 / (v36 - ((((1e10 / -2.25) / v4) - (0.0 + v36)) - ((-2.25 / (1e10 / (3.14159 + 1e10))) - float64(0.0))))) * v36))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
