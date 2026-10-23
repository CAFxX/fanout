package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint32) uint64 {
    if (1 >= 704) {
        p0 = 484
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 float32 = (1.5 / 0.1)
            v7 = v7
            switch "hello" {
            case "":
                p0 -= p0
            }
            _ = fi6
            _ = v7
        }
        if (-79 != -1) {
            if false {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v9 := []int{687, fi8, fi8, (fi8 & 32767)}
                    v10 := GenS1{F0: p0, F1: (fi8 * fi8), F2: "hello"}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        } else {
            p0 -= p0
            if false {
                p0 = p0
                p0 = 0
            }
        }
    } else {
        fmt.Println(p0)
        v11 := GenS1{F0: p0, F1: 705, F2: "\n\t"}
        _ = v11
    }
    p0 -= p0
    v12 := GenS1{F0: p0, F1: -9223372036854775808, F2: "foo bar"}
    v13 := GenS1{F0: p0, F1: 0, F2: "a"}
    var v14 float64 = ((((((0.0 - float64(1.5)) + ((float64(1.5) / 1.0) + 0.1)) - (3.14159 / -2.25)) / 1.0) / 1.5) - 0.1)
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 uint64) uint8 {
    var v15 int64 = -9223372036854775808
    if true {
        var v16 string = "foo bar"
        _ = v16
    }
    var v17 bool = (false || false)
    _ = p0
    _ = v15
    _ = v17
    return 183
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -9223372036854775808
    var v3 int32 = 2147483647
    var v4 uint32 = 265
    v5 := make(map[int]int)
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v4)
        var v19 int64 = v2
        v20 := GenS1{F0: (^(v4)), F1: fi18, F2: "a"}
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 -= (v4 | v4)
    fmt.Println(v2, v2)
    var v21 int64 = 0
    v4 -= v4
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int64 = (v21 | v2)
        for fi24 := 0; fi24 < 20; fi24++ {
            if true {
                switch "\n\t" {
                case "a":
                    v3 = 754
                default:
                    v4 -= (v4 % ((65535) | 1))
                }
                var v25 int64 = (v2 << ((4294967295) & 63))
                var v26 uint = 18446744073709551615
                _ = v25
                _ = v26
            } else {
                v4 += 781
                v3 -= v3
            }
            fmt.Println(v3)
            if false {
                var v27 uint16 = 42
                v28 := GenS2{F0: 255, F1: 0}
                _ = v27
                _ = v28
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    v4 = (v4 + 437)
    v4 = v4
    var v29 bool = (float64(((1e10 - 3.14159) + 1e10)) >= 1.5)
    if ("hello" == "foo bar") {
        v4 -= v4
        v4 = 553
    } else {
        switch "" {
        case "foo bar":
            var v30 int8 = -120
            _ = v30
        default:
            var v31 int8 = 100
            _ = v31
        }
    }
    var v32 string = (("a" + "") + "foo bar")
    fmt.Println(v4, v2, v3)
    v33 := GenS2{F0: 42, F1: 42}
    v34 := GenS1{F0: (v4 >> ((65535) & 63)), F1: -487, F2: (v32 + v32)}
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := []int{fi35, (fi35 - fi35), 210, 2147483647}
        _ = fi35
        _ = v36
    }
    if ((!(v29 || (!true))) || (v29 && (127 > -126))) {
        if (42 != -3) {
            var v37 float32 = (-2.25 - 1e10)
            _ = v37
        } else {
            var v38 int32 = v3
            var v39 int8 = 127
            _ = v38
            _ = v39
        }
    } else {
        v40 := GenS1{F0: (v4 >> ((65535) & 63)), F1: 32768, F2: v32}
        v41 := []int{int(v3), -381, 128}
        v21 = v21
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
