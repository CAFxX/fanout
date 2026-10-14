package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0(p0 bool, p1 int16) int64 {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = 1.5
        fmt.Println(v7, fi6, p0)
        _ = fi6
        _ = v7
    }
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = v8
    return -1
}
func genFn1(p0 uint, p1 int16, p2 uint) int64 {
    var v9 float32 = -2.25
    v10 := GenS1{F0: p1, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{132}
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 int8 = 42
    var v5 uint64 = 1
    switch 0 {
    case 496:
        fmt.Println(v4, v5)
    case -476:
        for fi11 := 0; fi11 < 17; fi11++ {
            _ = len(v3)
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int32 = -1
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    case 693:
        v14 := GenS2{F0: 0, F1: 1}
        _ = v14
    default:
        v15 := GenS3{F0: "hello", F1: 32768, F2: (v5 & (v5 * v5))}
        _ = v15
    }
    var v16 int = 775
    for fi17 := 0; fi17 < 9; fi17++ {
        delete(v2, "a")
        for fi18 := 0; fi18 < 13; fi18++ {
            v4 = v4
            v16 += v16
            v16 += (fi17 - 127)
            _ = fi18
        }
        v1 = append(v1, 65535)
        _ = fi17
    }
    fmt.Println(v4)
    if ((((0 > 255) || true) && ((v16 == (v16 % (((v16 ^ (v16 * v16))) | 1))) || false)) && true) {
        fmt.Println(v5)
        if (!true) {
            v19 := GenS2{F0: 61, F1: 225}
            _ = v19
        } else {
            var v20 int64 = 42
            _ = v20
        }
    }
    fmt.Println(v16)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int32 = 1
        _ = fi21
        _ = v22
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v4, v5)
    case "foo bar":
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS4{F0: 9223372036854775807, F1: "\n\t"}
            v25 := v3[-9223372036854775808]
            if (0 < 42) {
                var v26 int16 = 42
                var v27 uint64 = (v5 | (v5 << ((0) & 63)))
                _ = v26
                _ = v27
            } else {
                var v28 bool = true
                v16 = (-(v16))
                v5 = v5
                _ = v28
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v4, v16, v16)
    v5 = 0
    for fi29 := 0; fi29 < 2; fi29++ {
        if len(v1) > 0 {
            v1[0] = 1
        }
        if true {
            if (!(false && (0.0 <= 1e10))) {
                v30 := GenS5{F0: uint(v5), F1: 0}
                _ = v30
            }
            v31 := GenS2{F0: 894, F1: 42}
            switch (fi29 >> ((30) & 63)) {
            case 128:
                if true {
                    v4 = 42
                }
            default:
                var v32 uint16 = 0
                _ = v32
            }
            _ = v31
        }
        _ = fi29
    }
    var v33 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
