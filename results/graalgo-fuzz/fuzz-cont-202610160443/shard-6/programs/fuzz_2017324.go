package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int }

func main() {
    var v1 int = 4294967295
    var v2 int64 = 58
    var v3 float32 = (3.14159 + (0.0 - float32(float64(float32((1.5 * 1e10))))))
    v4 := make(map[int]bool)
    v5 := []int{v1, (^(v1))}
    v2 -= v2
    fmt.Println(v3, v3)
    var v6 int16 = -535
    delete(v4, v1)
    var v7 int8 = 35
    v8 := GenS1{F0: (v2 | (v2 / ((v2) | 1))), F1: 42, F2: 256}
    if len(v5) > 0 {
        v5[0] = v1
    }
    switch "" {
    case "a":
        var v9 float32 = (-2.25 - (-2.25 + v3))
        _ = v9
    case "":
        v7 += v7
    }
    if true {
        if true {
            if (true && (!true)) {
                v10 := GenS2{F0: (v1 ^ (v1 >> ((1) & 63))), F1: "a", F2: "a"}
                if (32767 < 255) {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v7 += 1
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 int64 = (v2 >> ((255) & 63))
                        delete(v4, int(v7))
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 19; fi14++ {
                        switch v1 {
                        case 42:
                            v15 := v4[(fi14 * v1)]
                            _ = v15
                        case 32767:
                            switch fi14 {
                            case 0:
                                delete(v4, v1)
                            case 32768:
                                v16 := v4[1]
                                _ = v16
                            default:
                                v4[(^(v1))] = false
                            }
                        case 128:
                            fmt.Println(v2)
                        }
                        v17 := GenS1{F0: 42, F1: 42, F2: 42}
                        v2 -= v2
                        _ = fi14
                        _ = v17
                    }
                }
                _ = v10
            } else {
                fmt.Println(v3, v2)
            }
            delete(v4, 4294967295)
        } else {
            fmt.Println(v7)
            v6 = v6
            for fi18 := 0; fi18 < 1; fi18++ {
                v1 -= (v1 + 4294967295)
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS3{F0: (-(v1)), F1: v2, F2: "hello"}
                    _ = fi19
                    _ = v20
                }
                if ((true && ((false || false) && ((!false) && true))) || true) {
                    v5 = append(v5, 32767)
                    var v21 uint = 712
                    if true {
                        fmt.Println(v2, v7)
                        for fi22 := 0; fi22 < 18; fi22++ {
                            switch "a" {
                            case "\n\t":
                                _ = len(v5)
                            }
                            var v23 int32 = 0
                            switch "hello" {
                            case "\n\t":
                                v21 = (v21 & v21)
                            case "hello":
                                v24 := GenS1{F0: (v2 - (v2 % ((42) | 1))), F1: 0, F2: (v23 >> ((v21) & 63))}
                                v6 = v6
                                _ = v24
                            default:
                                if ((v7 % ((v7) | 1)) < v7) {
                                    v25 := GenS4{F0: 42}
                                    switch (fi18 - int(v21)) {
                                    case 4294967295:
                                        if (!true) {
                                            v21 += v21
                                        } else {
                                            var v26 int16 = 0
                                            v27 := GenS5{F0: v26, F1: v1}
                                            _ = v26
                                            _ = v27
                                        }
                                    default:
                                        fmt.Println(v6)
                                    }
                                    var v28 uint = v21
                                    _ = v25
                                    _ = v28
                                }
                            }
                            _ = fi22
                            _ = v23
                        }
                        var v29 uint = v21
                        _ = v29
                    }
                    _ = v21
                }
                _ = fi18
            }
        }
        v30 := GenS4{F0: v7}
        v31 := GenS1{F0: (^(v2)), F1: -1, F2: 1}
        _ = v30
        _ = v31
    } else {
        var v32 string = "foo bar"
        for fi33 := 0; fi33 < 11; fi33++ {
            v32 = "hello"
            _ = fi33
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
