package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int, p1 uint16, p2 int32) uint64 {
    fmt.Println(p1, p0, p2)
    v5 := []int{p0, p0, p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 65535
    var v3 float32 = (1.5 + 1.5)
    v4 := []uint32{2}
    v6 := GenS1{F0: 1, F1: 2147483648}
    switch (("hello" + "foo bar") + "a") {
    case "hello":
        var v7 uint = 1
        _ = v7
    case "":
        v3 += -2.25
    case "\n\t":
        var v8 int16 = 1
        _ = v8
    default:
        if (786 > genFn0(42, 187, -84)) {
            if false {
                _ = len(v1)
            }
        } else {
            if false {
                if (1 <= -540) {
                    v2 -= 255
                } else {
                    var v9 bool = true
                    v10 := GenS1{F0: 180, F1: 965}
                    v2 = v2
                    _ = v10
                    _ = v9
                }
            } else {
                if true {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v12 := v1[298]
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                if (false || true) {
                                    v13 := GenS1{F0: 2, F1: v12}
                                    v4 = append(v4, (v12 / ((2147483647) | 1)))
                                    v2 = (^(v2))
                                    _ = v13
                                }
                            } else {
                                fmt.Println(v12)
                                for fi14 := 0; fi14 < 14; fi14++ {
                                    if true {
                                        if (fi11 != int(v12)) {
                                            switch fi11 {
                                            case 2147483647:
                                                v15 := GenS1{F0: 218, F1: 0}
                                                _ = v15
                                            case -235:
                                                v16 := GenS2{F0: fi11, F1: 2147483647}
                                                _ = v16
                                            case -395:
                                                var v17 uint = 404
                                                _ = v17
                                            default:
                                                v3 += (v3 - v3)
                                            }
                                        }
                                        fmt.Println(v2)
                                        v12 += v12
                                    } else {
                                        v1[v12] = v12
                                        var v18 int64 = -9223372036854775808
                                        _ = v18
                                    }
                                    v4 = append(v4, 65535)
                                    v19 := GenS1{F0: 2, F1: v12}
                                    _ = fi14
                                    _ = v19
                                }
                                v20 := GenS2{F0: -668, F1: (^(fi11))}
                                _ = v20
                            }
                        case "hello":
                            for fi21 := 0; fi21 < 8; fi21++ {
                                switch 42 {
                                case 305:
                                    v2 = (^(v2))
                                }
                                _ = fi21
                            }
                        default:
                            fmt.Println(v3, fi11, fi11)
                        }
                        _ = fi11
                        _ = v12
                    }
                    fmt.Println(v3, v2)
                    fmt.Println(v3)
                }
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v2, v2, v2)
                if (true || true) {
                    var v23 string = "a"
                    v24 := GenS3{F0: 1, F1: (^(v2)), F2: 127}
                    _ = v23
                    _ = v24
                }
                _ = fi22
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                v3 = float32(3.14159)
                v1[2] = 154
                v26 := GenS1{F0: 42, F1: 2}
                _ = fi25
                _ = v26
            }
        }
    }
    if (((3.14159 - 1e10) / 3.14159) != 0.0) {
        if true {
            for fi27 := 0; fi27 < 16; fi27++ {
                for fi28 := 0; fi28 < 8; fi28++ {
                    v3 = float32(0.1)
                    _ = fi28
                }
                v4 = append(v4, 2)
                _ = fi27
            }
        }
    }
    switch "hello" {
    case "foo bar":
        v2 += 0
    case "hello":
        var v29 int32 = -1
        _ = v29
    case "":
        var v30 int = 2147483647
        _ = v30
    default:
        v4 = append(v4, 393)
    }
    fmt.Println(v2, v3)
    var v31 uint16 = 32767
    v32 := GenS3{F0: 399, F1: (v2 + 0), F2: 127}
    v33 := v1[1]
    switch "\n\t" {
    case "":
        for fi34 := 0; fi34 < 16; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                v36 := GenS4{F0: 42, F1: v31}
                var v37 int = (fi34 * -1)
                _ = fi35
                _ = v36
                _ = v37
            }
            fmt.Println(v31, fi34)
            _ = fi34
        }
    case "hello":
        v1[(^(v33))] = (v33 % (((v33 - v33)) | 1))
    }
    var v38 uint16 = (v31 + 0)
    for fi39 := 0; fi39 < 19; fi39++ {
        v31 += (^(v38))
        var v40 int8 = 100
        _ = fi39
        _ = v40
    }
    v31 += (v38 % (((^(v38))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v6
}
