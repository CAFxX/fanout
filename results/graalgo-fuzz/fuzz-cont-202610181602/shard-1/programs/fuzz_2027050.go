package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int64, p2 int64) uint8 {
    switch 128 {
    case 128:
        var v7 uint64 = (^(p0))
        _ = v7
    case 4294967295:
        var v8 float64 = -2.25
        _ = v8
    case 0:
        p1 -= p2
    }
    if true {
        v9 := GenS1{F0: true, F1: 65535}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint) int32 {
    fmt.Println(p0)
    var v10 float32 = 0.1
    var v11 int8 = -1
    for fi12 := 0; fi12 < 17; fi12++ {
        p0 = 352
        v13 := GenS1{F0: (true || true), F1: p0}
        if (!false) {
            if false {
                v11 = (v11 | v11)
                for fi14 := 0; fi14 < 10; fi14++ {
                    v10 = (0.1 * ((-2.25 * v10) * v10))
                    p0 = (p0 % (((p0 + (^(p0)))) | 1))
                    _ = fi14
                }
            }
        } else {
            v10 += v10
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return -2147483648
}

func main() {
    var v1 string = (("foo bar" + "hello") + "\n\t")
    var v2 string = v1
    v3 := []uint32{2147483648, 1, 820}
    v4 := make(map[string]int)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]uint32)
    for fi15 := 0; fi15 < 17; fi15++ {
        switch (fi15 - (fi15 - fi15)) {
        case -782:
            if false {
                if (false && true) {
                    delete(v6, ((v2 + "a") + "hello"))
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        fmt.Println(fi15)
                        fmt.Println(fi16)
                        fmt.Println(fi15, v2)
                        v17 := GenS2{F0: -1}
                        if (!true) {
                            fmt.Println(fi15, v1)
                        } else {
                            var v18 float64 = 1.5
                            v19 := v4[v2]
                            v5 = (^(v5))
                            _ = v18
                            _ = v19
                        }
                        v5 += v5
                        _ = fi16
                        _ = v17
                    }
                    var v20 int64 = 42
                    var v21 int64 = v20
                    _ = v20
                    _ = v21
                }
                v1 = v2
            }
        default:
            for fi22 := 0; fi22 < 15; fi22++ {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                _ = fi22
            }
        }
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                if (15 < 1) {
                    v24 := GenS1{F0: (226 <= 119), F1: 42}
                    switch (fi23 | (fi15 << ((42) & 63))) {
                    case 9223372036854775807:
                        v2 = "foo bar"
                    case 1:
                        var v25 float32 = (0.1 * (-2.25 / 1e10))
                        _ = v25
                    case 668:
                        v5 += v5
                    default:
                        v1 = v1
                    }
                    _ = v24
                } else {
                    _ = len(v6)
                }
                v26 := GenS3{F0: v2}
                if false {
                    delete(v6, v1)
                    fmt.Println(v2, v5, v2)
                }
                _ = fi23
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 13; fi27++ {
                fmt.Println(v5, v2)
                fmt.Println(v2, fi15, v5)
                v5 = v5
                _ = fi27
            }
            delete(v6, v1)
            var v28 int = (fi15 >> ((65535) & 63))
            _ = v28
        }
        _ = fi15
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint = 18446744073709551615
        fmt.Println(v5, v5)
        switch (fi29 | fi29) {
        case 0:
            fmt.Println(v1, fi29)
        case -680:
            var v31 uint64 = (v5 >> (((v30 * v30)) & 63))
            _ = v31
        default:
            var v32 uint = uint(v5)
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v33 := GenS1{F0: (616 < 127), F1: 42}
    if ((float64((0.1 / 1.0)) - (3.14159 + 0.0)) <= -2.25) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v34 float32 = ((1.5 / 1e10) + (1.5 / 1.5))
        var v35 int32 = int32(v5)
        _ = v34
        _ = v35
    }
    var v36 uint32 = 0
    if (65535 < 691) {
        v37 := GenS4{F0: (-2147483648 < 32768)}
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 int64 = 0
            v36 = uint32(v5)
            _ = fi38
            _ = v39
        }
        var v40 int = 128
        _ = v37
        _ = v40
    }
    v41 := GenS3{F0: v2}
    delete(v4, v2)
    _ = len(v4)
    switch v2 {
    case "a":
        var v42 int16 = 32767
        _ = v42
    case "\n\t":
        var v43 uint64 = 18446744073709551615
        _ = v43
    case "foo bar":
        var v44 uint64 = 433
        _ = v44
    default:
        if false {
            v45 := GenS4{F0: false}
            var v46 float32 = 1e10
            _ = v45
            _ = v46
        } else {
            var v47 string = "a"
            if true {
                v1 = v2
            } else {
                switch "" {
                case "\n\t":
                    if true {
                        fmt.Println(v1, v1, v47)
                        for fi48 := 0; fi48 < 18; fi48++ {
                            var v49 int32 = 0
                            delete(v6, v47)
                            _ = fi48
                            _ = v49
                        }
                    }
                case "":
                    v50 := GenS5{F0: (^(v5))}
                    _ = v50
                default:
                    if len(v3) > 0 {
                        v3[0] = v36
                    }
                }
            }
            fmt.Println(v5)
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
