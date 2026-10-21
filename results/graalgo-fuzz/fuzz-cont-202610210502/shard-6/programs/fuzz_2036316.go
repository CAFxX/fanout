package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int64; F2 int16 }
func genFn0(p0 string, p1 int) int16 {
    p1 += p1
    p1 += p1
    switch ("foo bar" + (p0 + ((((p0 + "hello") + "foo bar") + ("hello" + p0)) + (p0 + p0)))) {
    case "a":
        var v7 uint8 = 26
        _ = v7
    case "\n\t":
        p0 = p0
    default:
        v8 := GenS1{F0: false}
        _ = v8
    }
    if (42 <= 42) {
        var v9 int32 = 2147483647
        v10 := []int{0, p1, p1}
        switch ((p0 + "") + p0) {
        case "foo bar":
            p1 = -2147483648
        case "\n\t":
            var v11 int8 = 127
            _ = v11
        case "hello":
            p1 -= p1
        }
        _ = v10
        _ = v9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint = 1
        fmt.Println(v13, fi12)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return -131
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 uint32 = 2
    var v5 int = 499
    var v6 uint = 858
    fmt.Println(v5, v6)
    v6 = (^(v6))
    v6 -= (^(v6))
    v1 = (^(v1))
    for fi14 := 0; fi14 < 18; fi14++ {
        switch "" {
        case "\n\t":
            var v15 uint8 = 1
            _ = v15
        }
        _ = fi14
    }
    v3[65535] = v5
    if true {
        delete(v2, "")
    } else {
        switch "hello" {
        case "foo bar":
            v5 = (v5 % (((v5 * (v5 | (v5 + (^(v5)))))) | 1))
        case "\n\t":
            var v16 int64 = 1
            _ = v16
        default:
            v17 := GenS2{F0: -491, F1: (v5 ^ 1)}
            _ = v17
        }
    }
    switch 1 {
    case 599:
        var v18 int64 = -9223372036854775808
        _ = v18
    case -832:
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v19 := v3[v4]
            _ = v19
        default:
            var v20 uint8 = 3
            _ = v20
        }
    }
    switch "a" {
    case "hello":
        v5 -= v5
    case "":
        for fi21 := 0; fi21 < 11; fi21++ {
            switch fi21 {
            case 2147483647:
                if false {
                    if false {
                        var v22 uint16 = 1
                        var v23 int = v5
                        v1 += v1
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(v6, v1)
                        v4 = v4
                        if true {
                            var v24 float64 = ((0.1 / float64((3.14159 + ((3.14159 * 1.5) / 1.0)))) - (1e10 * float64(1.5)))
                            delete(v3, v4)
                            _ = v24
                        } else {
                            v4 += (v4 % ((v4) | 1))
                        }
                    }
                    v6 += v6
                }
            case 32767:
                if true {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS1{F0: true}
                        switch "hello" {
                        case "":
                            if (true && (true && false)) {
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            } else {
                                fmt.Println(v4, v1, v4)
                                var v28 uint8 = 73
                                var v29 int = fi25
                                if (v4 != (v4 % ((v4) | 1))) {
                                    var v30 uint64 = 0
                                    switch "\n\t" {
                                    case "a":
                                        v31 := GenS1{F0: false}
                                        _ = v31
                                    }
                                    _ = v30
                                }
                                _ = v28
                                _ = v29
                            }
                            v1 -= 42
                            v4 += 65535
                            if false {
                                var v32 uint8 = 1
                                var v33 uint16 = 1
                                _ = v32
                                _ = v33
                            }
                        }
                        v34 := GenS3{F0: v4}
                        _ = fi25
                        _ = v26
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 10; fi35++ {
                        fmt.Println(v6)
                        if true {
                            var v36 float32 = 3.14159
                            v1 = (v1 / ((genFn0("\n\t", (fi35 | 4294967295))) | 1))
                            v37 := []int{}
                            _ = v36
                            _ = v37
                        } else {
                            switch v5 {
                            case -84:
                                var v38 uint = 0
                                _ = v38
                            case -663:
                                var v39 int = (fi21 << (((v6 + 2)) & 63))
                                _ = v39
                            default:
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch (-(v5)) {
                                        case 32768:
                                            var v40 int64 = -1
                                            _ = v40
                                        case 0:
                                            v41 := GenS4{F0: 2, F1: 9223372036854775807}
                                            _ = v41
                                        default:
                                            fmt.Println(v6, fi35, v4)
                                        }
                                    }
                                    var v42 int8 = -86
                                    v5 = 0
                                    v43 := GenS5{F0: true}
                                    _ = v42
                                    _ = v43
                                } else {
                                    fmt.Println(fi21, v1)
                                }
                            }
                            v44 := GenS6{F0: "\n\t", F1: 9223372036854775807, F2: genFn0("hello", fi21)}
                            var v45 int32 = 686
                            v46 := GenS4{F0: 255, F1: 42}
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = fi35
                    }
                    switch 1 {
                    case 362:
                        switch (fi21 & fi21) {
                        case 4294967295:
                            v47 := v3[v4]
                            _ = v47
                        case 1:
                            v48 := GenS3{F0: 42}
                            _ = v48
                        case -721:
                            v49 := []int{(v5 | fi21)}
                            _ = v49
                        default:
                            fmt.Println(v6, v4)
                        }
                    default:
                        v1 += v1
                    }
                }
            }
            var v50 float32 = 1e10
            v51 := GenS4{F0: 255, F1: 9223372036854775807}
            _ = fi21
            _ = v50
            _ = v51
        }
    case "a":
        fmt.Println(v5)
    default:
        var v52 float64 = (0.0 - (0.0 * 3.14159))
        _ = v52
    }
    v6 -= (v6 % ((v6) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
