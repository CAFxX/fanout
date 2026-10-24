package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int8; F1 uint64; F2 int64 }
type GenS10 struct { F0 int; F1 uint }
func genFn0(p0 uint8, p1 uint32, p2 bool) uint32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p2)
        p1 = p1
        var v8 string = "a"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 0, F1: 255}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 uint) uint32 {
    switch 128 {
    case 430:
        var v11 uint64 = 42
        _ = v11
    default:
        var v12 uint8 = 37
        _ = v12
    }
    var v13 int = -618
    switch "hello" {
    case "foo bar":
        var v14 int8 = 6
        _ = v14
    case "a":
        v15 := []int{v13, (v13 >> (((p0 << ((p0) & 63))) & 63)), v13}
        _ = v15
    case "\n\t":
        p0 = 934
    default:
        switch int(p0) {
        case 32767:
            switch 2147483647 {
            case -48:
                switch "" {
                case "a":
                    switch "" {
                    case "hello":
                        var v16 float64 = 1e10
                        _ = v16
                    case "a":
                        v17 := GenS2{F0: 885}
                        _ = v17
                    case "":
                        var v18 uint16 = 1
                        _ = v18
                    default:
                        switch (-(v13)) {
                        case 4294967295:
                            switch (-(v13)) {
                            case 958:
                                v13 -= (v13 * v13)
                            case -2147483648:
                                fmt.Println(v13)
                            default:
                                p0 = (p0 * (p0 | 0))
                            }
                        default:
                            v19 := []int{(v13 << ((p0) & 63)), -696, -500, v13}
                            _ = v19
                        }
                    }
                case "hello":
                    v20 := []int{(v13 / ((-1) | 1)), 9223372036854775807, v13}
                    _ = v20
                case "":
                    v21 := []int{(v13 ^ v13), (v13 / (((v13 << ((p0) & 63))) | 1))}
                    _ = v21
                default:
                    fmt.Println(v13, p0)
                }
            case -9223372036854775808:
                var v22 uint16 = 42
                _ = v22
            }
        }
    }
    var v23 bool = ((!false) && (true || true))
    var v24 uint16 = 0
    _ = p0
    _ = v13
    _ = v23
    _ = v24
    return 965
}

func main() {
    var v1 uint16 = 737
    v2 := make(map[int]bool)
    var v3 uint64 = 865
    v4 := make(map[uint32]int)
    var v5 int = 9223372036854775807
    v6 := make(map[string]int)
    if ((v3 - 42) != (v3 ^ 18446744073709551615)) {
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 float64 = (1e10 * -2.25)
            v27 := GenS3{F0: 4294967295, F1: -1}
            v3 = v3
            _ = fi25
            _ = v26
            _ = v27
        }
        v1 = v1
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint32 = 4294967295
        _ = fi28
        _ = v29
    }
    v30 := GenS4{F0: (v5 + -680)}
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS5{F0: 516}
        v33 := v2[fi31]
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 float64 = 1.5
        _ = fi34
        _ = v35
    }
    v3 -= 4294967295
    v5 = v5
    var v36 float64 = 1.5
    for fi37 := 0; fi37 < 5; fi37++ {
        if (!true) {
            fmt.Println(v36)
            var v38 uint64 = (v3 * (v3 >> ((42) & 63)))
            var v39 int16 = 42
            _ = v38
            _ = v39
        }
        var v40 float64 = v36
        _ = fi37
        _ = v40
    }
    if true {
        if (false && false) {
            v36 = -2.25
            var v41 float64 = v36
            _ = v41
        } else {
            v42 := GenS1{F0: v1, F1: 126}
            for fi43 := 0; fi43 < 3; fi43++ {
                for fi44 := 0; fi44 < 2; fi44++ {
                    v45 := v2[fi43]
                    _ = fi44
                    _ = v45
                }
                v46 := []int{87, v5, v5}
                _ = fi43
                _ = v46
            }
            _ = v42
        }
        switch ("" + "a") {
        case "hello":
            v47 := GenS6{F0: -188, F1: 2147483648, F2: (^(v3))}
            _ = v47
        }
        v48 := GenS7{F0: 1}
        _ = v48
    }
    v49 := GenS6{F0: 1, F1: 42, F2: (v3 % ((18446744073709551615) | 1))}
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if ((v3 & v3) <= 42) {
            for fi50 := 0; fi50 < 8; fi50++ {
                v51 := GenS8{F0: 763, F1: 2}
                _ = fi50
                _ = v51
            }
            if (0 != (v5 & v5)) {
                var v52 bool = (-639 == 42)
                var v53 float64 = 0.0
                _ = v52
                _ = v53
            } else {
                var v54 bool = true
                v3 -= v3
                _ = v54
            }
            v3 = 210
        }
    case "":
        var v55 uint = 944
        _ = v55
    case "\n\t":
        v56 := GenS9{F0: 127, F1: v3, F2: -1}
        _ = v56
    default:
        switch "hello" {
        case "":
            v57 := GenS10{F0: v5, F1: 0}
            _ = v57
        case "foo bar":
            for fi58 := 0; fi58 < 8; fi58++ {
                for fi59 := 0; fi59 < 12; fi59++ {
                    v3 += 18446744073709551615
                    _ = fi59
                }
                v60 := []int{(fi58 ^ fi58), fi58, fi58}
                _ = fi58
                _ = v60
            }
        case "\n\t":
            fmt.Println(v1, v3)
        default:
            v3 = (v3 >> ((42) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
