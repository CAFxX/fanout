package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 int64) uint32 {
    v6 := []int{479}
    p0 = (p0 | p0)
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(fi8, fi9)
            _ = fi9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            v6 = append(v6, fi8)
            _ = fi10
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 291
    v3 := make(map[uint32]bool)
    var v4 float64 = ((((1e10 - 0.1) + (float64(1.5) - ((1.5 - ((1e10 * 1e10) - ((1e10 - 1.5) - 3.14159))) + (0.0 + 1.5)))) - 1.5) + float64((-2.25 / (1.5 * (float32(1.5) / 1e10)))))
    v5 := make(map[uint32]bool)
    fmt.Println(v4)
    var v11 uint32 = 488
    if (!((9223372036854775807 > 0) || (2 >= uint16(v2)))) {
        var v12 int16 = 42
        v11 = (v11 - v11)
        var v13 uint16 = 187
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if false {
            v2 += 0
            fmt.Println(fi14, v4, fi14)
        } else {
            if (0 <= 65535) {
                switch fi14 {
                case 127:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v4 -= v4
                        var v16 float64 = v4
                        _ = fi15
                        _ = v16
                    }
                default:
                    v4 = v4
                }
            }
            if true {
                switch (fi14 << ((2) & 63)) {
                case -1:
                    var v17 int64 = 9223372036854775807
                    _ = v17
                case -786:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        fmt.Println(fi18, v11)
                        _ = fi18
                    }
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := GenS1{F0: 875, F1: -1}
                        var v21 int8 = 33
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                v22 := []int{fi14, 127, (fi14 % ((0) | 1)), (-(fi14))}
                _ = v22
            }
            var v23 uint = 18446744073709551615
            _ = v23
        }
        _ = fi14
    }
    var v24 int32 = 2147483647
    var v25 int32 = v24
    v26 := GenS1{F0: 255, F1: v25}
    if (-264 != 128) {
        v2 = (v2 - (^(v2)))
    } else {
        v2 = v2
        if true {
            v25 += v25
            if true {
                var v27 int64 = -543
                _ = v27
            } else {
                v28 := []int{}
                _ = v28
            }
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint64 = (v2 + uint64(v11))
        _ = fi29
        _ = v30
    }
    var v31 float64 = v4
    switch 32767 {
    case 128:
        v32 := []int{127, -2147483648}
        _ = v32
    }
    _ = len(v1)
    v4 += 0.0
    var v33 bool = false
    switch ("" + "foo bar") {
    case "":
        v2 += 42
    case "\n\t":
        var v34 float64 = ((v31 + v4) * float64((3.14159 + 1.5)))
        _ = v34
    case "a":
        for fi35 := 0; fi35 < 3; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                v37 := GenS2{F0: (v25 - (-(v25))), F1: -38}
                switch (("\n\t" + "") + ("a" + "\n\t")) {
                case "\n\t":
                    if v33 {
                        var v38 int8 = -54
                        if true {
                            v39 := []int{(fi36 / ((fi35) | 1)), 127, fi36, fi36}
                            var v40 int16 = 1000
                            _ = v39
                            _ = v40
                        }
                        _ = v38
                    }
                case "":
                    switch 9223372036854775807 {
                    case 936:
                        if (!(!false)) {
                            fmt.Println(v25, v25, v2)
                        }
                    }
                }
                var v41 int32 = (v25 / (((v25 | (v24 ^ v25))) | 1))
                _ = fi36
                _ = v37
                _ = v41
            }
            _ = fi35
        }
    default:
        v42 := GenS3{F0: "a", F1: "a"}
        _ = v42
    }
    switch (("foo bar" + "foo bar") + "a") {
    case "a":
        v43 := GenS2{F0: (v25 >> ((278) & 63)), F1: 42}
        _ = v43
    default:
        var v44 int16 = -346
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
}
