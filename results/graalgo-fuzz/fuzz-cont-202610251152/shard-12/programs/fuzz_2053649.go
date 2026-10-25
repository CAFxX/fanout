package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 bool, p2 int8) uint32 {
    p0 += (p0 ^ (-(p0)))
    v6 := []int{127, 9223372036854775807, 0, -836}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 782
}

func main() {
    v1 := []int{128, 0, -578}
    v2 := make(map[uint32]uint32)
    var v3 int64 = -1
    var v4 float32 = 1e10
    v5 := []uint32{0, 930, 1}
    var v7 uint8 = 219
    if false {
        fmt.Println(v7)
        switch "a" {
        case "hello":
            v8 := GenS1{F0: "\n\t", F1: -1, F2: 2147483647}
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v1) > 0 {
                v1[0] = (fi9 >> ((259) & 63))
            }
            _ = fi9
        }
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 float64 = 0.1
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v7, v4)
    _ = len(v2)
    if ((-(v3)) != v3) {
        if true {
            var v12 int16 = 0
            var v13 uint16 = 32767
            _ = v12
            _ = v13
        } else {
            v4 -= (v4 - v4)
            if true {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 bool = true
                    _ = fi14
                    _ = v15
                }
                if true {
                    if len(v1) > 0 {
                        v1[0] = 650
                    }
                    for fi16 := 0; fi16 < 10; fi16++ {
                        _ = len(v2)
                        v7 = v7
                        _ = fi16
                    }
                    delete(v2, 2147483647)
                }
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                if (!true) {
                    v5 = append(v5, 445)
                    var v18 uint = 42
                    if (v18 > v18) {
                        var v19 int8 = 42
                        fmt.Println(v4)
                        switch 32768 {
                        case 32767:
                            v7 += (^(v7))
                        case 760:
                            v20 := GenS1{F0: "a", F1: (^(v19)), F2: 9223372036854775807}
                            _ = v20
                        case -387:
                            var v21 float32 = (v4 + -2.25)
                            _ = v21
                        default:
                            var v22 uint = v18
                            _ = v22
                        }
                        _ = v19
                    }
                    _ = v18
                }
                _ = fi17
            }
        }
    } else {
        if false {
            if false {
                v2[330] = 0
                v3 += (-(v3))
                if true {
                    fmt.Println(v3, v7)
                    if true {
                        _ = len(v2)
                        v3 += 510
                    }
                }
            }
            var v23 int16 = 1000
            var v24 int8 = int8(v7)
            _ = v23
            _ = v24
        }
        if ((v7 ^ v7) >= (v7 * v7)) {
            v7 = 1
            switch ((("" + "") + ("a" + "hello")) + "\n\t") {
            case "hello":
                if false {
                    v7 += (^(v7))
                    var v25 float64 = 3.14159
                    _ = v25
                } else {
                    var v26 float64 = ((1e10 * (0.1 + 1.5)) * (3.14159 - 0.0))
                    v2[42] = uint32(v7)
                    switch -218 {
                    case 609:
                        if (1.5 <= v4) {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            var v27 float64 = v26
                            _ = v27
                        } else {
                            var v28 int64 = v3
                            _ = v28
                        }
                    case 544:
                        var v29 int8 = 42
                        _ = v29
                    case 0:
                        switch "foo bar" {
                        case "\n\t":
                            v3 -= (v3 | v3)
                        case "foo bar":
                            fmt.Println(v4, v3)
                        case "a":
                            v30 := GenS2{F0: -193, F1: 317, F2: 2}
                            _ = v30
                        default:
                            v26 = 1e10
                        }
                    }
                    _ = v26
                }
            default:
                var v31 uint32 = genFn0(1, false, -72)
                _ = v31
            }
        }
    }
    v3 += 6
    var v32 bool = (4294967295 > 2)
    v3 += (v3 / ((9223372036854775807) | 1))
    v33 := GenS3{F0: 0}
    v34 := GenS1{F0: ("foo bar" + "a"), F1: 72, F2: -9223372036854775808}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 string = "a"
    var v36 uint = 127
    switch "\n\t" {
    case "foo bar":
        v37 := GenS1{F0: "\n\t", F1: 127, F2: 1}
        _ = v37
    case "":
        fmt.Println(v3)
    default:
        var v38 bool = (0 > 695)
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        switch (fi39 % ((32768) | 1)) {
        case 42:
            switch 42 {
            case 32768:
                var v40 uint = 255
                _ = v40
            }
        default:
            v32 = (102 != 100)
        }
        _ = fi39
    }
    v41 := GenS2{F0: -890, F1: 65535, F2: v36}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
    _ = v7
}
