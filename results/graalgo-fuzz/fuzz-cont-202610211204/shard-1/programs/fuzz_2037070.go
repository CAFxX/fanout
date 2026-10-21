package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int32, p2 int8) uint64 {
    var v7 int8 = (p2 << ((4294967295) & 63))
    switch 127 {
    case 128:
        p1 = p1
    default:
        if ((("" + "") + "hello") == "\n\t") {
            var v8 uint8 = 94
            _ = v8
        } else {
            if (p0 == p0) {
                var v9 uint64 = 4294967295
                v10 := GenS1{F0: "", F1: 677, F2: -743}
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS2{F0: 42, F1: p0, F2: 42}
                v13 := []int{(-(fi11)), -1, (fi11 % ((fi11) | 1)), fi11}
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v7, v7)
        }
    }
    if false {
        v14 := []int{2147483647}
        var v15 uint16 = (p0 - 65535)
        var v16 float64 = -2.25
        _ = v14
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: "", F1: (p0 % ((p0) | 1)), F2: 0}
        _ = v17
    }
    p2 += 127
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, -2147483648}
    var v3 int = 2147483647
    v4 := make(map[int]int)
    var v5 int32 = 2147483647
    var v6 bool = false
    var v18 int16 = -1
    switch "hello" {
    case "":
        v6 = v6
    case "hello":
        var v19 uint8 = 118
        _ = v19
    }
    fmt.Println(v6, v5)
    switch v3 {
    case 2147483647:
        v3 = 2147483647
    default:
        if (!v6) {
            v20 := GenS3{F0: "hello", F1: "hello"}
            for fi21 := 0; fi21 < 16; fi21++ {
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (-2.25 == (float32(3.14159) / 1.0)) {
                        v6 = (v6 && ((366 >= 65535) || (100 == 127)))
                        if false {
                            var v23 uint64 = 981
                            var v24 uint16 = 65535
                            var v25 uint32 = 1
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        var v26 string = "hello"
                        _ = v26
                    }
                    _ = fi22
                }
                _ = fi21
            }
            var v27 int = (v3 - (v3 >> ((4294967295) & 63)))
            _ = v20
            _ = v27
        } else {
            var v28 uint8 = 255
            var v29 int64 = 671
            _ = v28
            _ = v29
        }
    }
    switch "foo bar" {
    case "a":
        v3 = 42
    case "":
        var v30 uint = 220
        _ = v30
    case "\n\t":
        for fi31 := 0; fi31 < 19; fi31++ {
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi31
        }
    }
    v6 = v6
    var v32 uint64 = 1
    var v33 uint = 18446744073709551615
    if v6 {
        v32 += (^(v32))
    } else {
        v3 += (v3 >> ((v33) & 63))
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(fi34, v3, v5)
            for fi35 := 0; fi35 < 1; fi35++ {
                var v36 uint8 = 42
                var v37 float64 = 1e10
                if ((^(v32)) > (v32 % ((2) | 1))) {
                    var v38 uint16 = 1
                    if v6 {
                        switch (fi35 | fi35) {
                        case -860:
                            var v39 float64 = v37
                            _ = v39
                        case 1:
                            var v40 int = (v3 >> ((v33) & 63))
                            _ = v40
                        default:
                            var v41 string = "a"
                            _ = v41
                        }
                        fmt.Println(v36, fi34)
                    }
                    _ = v38
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
    }
    if (false && v6) {
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 uint64 = (v32 << (((v33 >> ((v33) & 63))) & 63))
            var v44 bool = (!(!v6))
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    if (0 < -9223372036854775808) {
        v45 := GenS4{F0: 2}
        v4[(v3 << ((v33) & 63))] = v3
        if false {
            v5 -= (^(v5))
        }
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
