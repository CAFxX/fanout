package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 uint16) uint16 {
    v6 := []int{1, -2147483648}
    if (0.1 != 3.14159) {
        for fi7 := 0; fi7 < 19; fi7++ {
            if len(v6) > 0 {
                v6[0] = fi7
            }
            var v8 uint8 = 15
            _ = fi7
            _ = v8
        }
        fmt.Println(p0)
        var v9 int64 = 854
        _ = v9
    }
    _ = len(v6)
    _ = p0
    _ = v6
    return (p0 << ((4294967295) & 63))
}
func genFn1(p0 bool, p1 string, p2 uint16) uint {
    p2 = 42
    switch 9223372036854775807 {
    case 9223372036854775807:
        if (1 != 11) {
            v10 := GenS1{F0: ((false && p0) && (2 == 1))}
            if false {
                if (1.5 == 0.0) {
                    switch 32767 {
                    case 32767:
                        var v11 bool = true
                        _ = v11
                    }
                    v12 := GenS1{F0: p0}
                    v13 := []int{}
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        p2 = genFn0((p2 >> ((42) & 63)))
                        if p0 {
                            switch "hello" {
                            case "a":
                                p1 = p1
                            case "":
                                var v15 int = (fi14 % ((fi14) | 1))
                                _ = v15
                            case "hello":
                                fmt.Println(p0, p2, p2)
                            }
                            fmt.Println(fi14, p2)
                            var v16 float64 = 3.14159
                            _ = v16
                        } else {
                            var v17 float32 = (0.1 - float32((1.5 * 0.1)))
                            _ = v17
                        }
                        _ = fi14
                    }
                    if p0 {
                        var v18 float64 = 0.1
                        switch -122 {
                        case 197:
                            var v19 bool = (1 != 255)
                            _ = v19
                        case -2147483648:
                            v20 := []int{-65, 32767, 32767, 127}
                            _ = v20
                        default:
                            var v21 int8 = 41
                            _ = v21
                        }
                        _ = v18
                    }
                }
                p2 = 570
                if (127 == 53) {
                    var v22 int = 2147483647
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := GenS1{F0: (32767 >= 32767)}
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
            }
            _ = v10
        } else {
            v25 := GenS1{F0: true}
            v26 := []int{32767, -9223372036854775808}
            v26 = append(v26, 382)
            _ = v25
            _ = v26
        }
    case 455:
        var v27 string = ""
        _ = v27
    default:
        if true {
            var v28 string = ""
            v29 := GenS1{F0: true}
            _ = v28
            _ = v29
        } else {
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []uint32{42, 2, 2}
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 string = "a"
    var v30 uint8 = 2
    v30 += v30
    v4[v2] = (v2 & v2)
    fmt.Println(v2, v2, v5)
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 float32 = 1.5
        var v33 int32 = 0
        _ = fi31
        _ = v32
        _ = v33
    }
    v1 = append(v1, 43)
    v34 := GenS2{F0: 28, F1: 288}
    v30 -= (v30 / (((v30 * (v30 * v30))) | 1))
    for fi35 := 0; fi35 < 14; fi35++ {
        if true {
            v5 = v5
        } else {
            var v36 bool = false
            _ = v36
        }
        fmt.Println(v30, v2)
        _ = fi35
    }
    var v37 int32 = 0
    var v38 float32 = (((float32(((-2.25 + -2.25) - (float64(-2.25) - 0.1))) + (1.5 / 1.0)) + 1.5) - -2.25)
    v5 = "hello"
    v4[v2] = v2
    if (false || false) {
        switch 2147483647 {
        case -9223372036854775808:
            v2 += v2
        case 1:
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 uint8 = v30
                _ = fi39
                _ = v40
            }
        case 4294967295:
            v2 = v2
        default:
            v41 := GenS1{F0: ((v5 <= v5) || (!true))}
            _ = v41
        }
        fmt.Println(v30, v38)
    }
    if (282 <= 1) {
        var v42 float64 = 1e10
        _ = v42
    } else {
        if (v30 >= v30) {
            switch ((v5 + v5) + "a") {
            case "hello":
                v5 = "a"
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
