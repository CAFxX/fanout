package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 string; F2 uint8 }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int32, p2 string) string {
    fmt.Println(p1, p0, p1)
    for fi4 := 0; fi4 < 15; fi4++ {
        if false {
            v5 := GenS1{F0: 42}
            var v6 int64 = 9223372036854775807
            if (!(0 > 20)) {
                var v7 int16 = 264
                p1 = p1
                p1 = (p1 >> ((255) & 63))
                _ = v7
            }
            _ = v5
            _ = v6
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                p1 += (^(p1))
                _ = fi8
            }
            var v9 int8 = 42
            v10 := []int{(fi4 % ((fi4) | 1)), (fi4 * int(p1))}
            _ = v10
            _ = v9
        }
        fmt.Println(p1, p2, p2)
        v11 := []int{fi4, -9223372036854775808}
        _ = fi4
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1(p0 int16, p1 uint8, p2 uint8) uint64 {
    switch (("hello" + ("foo bar" + "a")) + "") {
    case "\n\t":
        var v12 bool = (0 < 42)
        _ = v12
    case "foo bar":
        p2 = 42
    case "":
        p1 -= p2
    default:
        fmt.Println(p2, p1, p0)
    }
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            p1 = (p2 * (p1 << ((42) & 63)))
            if (0 >= 1) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 uint64 = 4294967295
                    _ = fi14
                    _ = v15
                }
                switch (fi13 + fi13) {
                case -2147483648:
                    p0 += p0
                case 0:
                    var v16 int8 = 100
                    _ = v16
                default:
                    v17 := GenS2{F0: (fi13 ^ -9223372036854775808), F1: 42}
                    _ = v17
                }
            }
            switch "foo bar" {
            case "foo bar":
                v18 := []int{}
                _ = v18
            case "hello":
                var v19 int32 = 236
                _ = v19
            }
            _ = fi13
        }
        var v20 uint32 = 42
        fmt.Println(v20)
        _ = v20
    } else {
        fmt.Println(p2, p1, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 247
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 int64 = 9223372036854775807
    switch "a" {
    case "foo bar":
        for fi21 := 0; fi21 < 7; fi21++ {
            switch 9223372036854775807 {
            case 4294967295:
                var v22 uint64 = 42
                _ = v22
            case 1:
                v1[(fi21 % ((fi21) | 1))] = (fi21 << ((2) & 63))
            default:
                v3 += v3
            }
            fmt.Println(v3, fi21)
            _ = fi21
        }
    case "":
        v23 := GenS3{F0: 1, F1: (false && false), F2: false}
        _ = v23
    case "\n\t":
        v24 := GenS3{F0: 65535, F1: (-32768 <= -32768), F2: (false && true)}
        _ = v24
    default:
        v25 := GenS1{F0: 4294967295}
        _ = v25
    }
    var v26 int32 = 300
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := []int{(fi28 % (((fi28 << ((2) & 63))) | 1)), -9223372036854775808, fi28, fi27}
            if (true || true) {
                var v30 string = "a"
                v2[(fi28 * (fi27 / ((fi28) | 1)))] = 627
                _ = v30
            }
            if false {
                v3 += 0
            }
            _ = fi28
            _ = v29
        }
        v31 := GenS4{F0: 4294967295}
        _ = fi27
        _ = v31
    }
    var v32 int64 = v3
    v33 := GenS5{F0: 272, F1: 0, F2: (v26 - (v26 / ((v26) | 1)))}
    v34 := GenS2{F0: 42, F1: 990}
    var v35 uint64 = 1
    var v36 int = -2147483648
    switch "hello" {
    case "foo bar":
        v37 := GenS4{F0: genFn1(42, 42, 146)}
        _ = v37
    case "\n\t":
        if (843 < 737) {
            _ = len(v1)
            v35 = 18446744073709551615
            if (!(int64(v36) == v32)) {
                v36 -= 9223372036854775807
                if true {
                    v38 := GenS6{F0: (-(v32)), F1: (("\n\t" + "") + "foo bar"), F2: 42}
                    _ = v38
                } else {
                    var v39 float64 = ((((0.1 + (-2.25 * (1.5 - (-2.25 - 0.1)))) * -2.25) - 1.5) / float64(((((0.0 + 0.0) + (((0.1 / 3.14159) + 1e10) * -2.25)) + 0.0) + 0.0)))
                    v3 -= (-(v32))
                    _ = v39
                }
            }
        } else {
            for fi40 := 0; fi40 < 7; fi40++ {
                fmt.Println(v36)
                _ = fi40
            }
            var v41 uint16 = 65535
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 uint32 = 2147483647
                _ = fi42
                _ = v43
            }
            _ = v41
        }
    case "hello":
        v35 = (v35 >> ((26) & 63))
    }
    if true {
        var v44 uint64 = v35
        _ = v44
    }
    v45 := GenS7{F0: -46, F1: v36, F2: 0}
    v46 := []int{(v36 >> ((4294967295) & 63)), 1, 751}
    v26 -= -2147483648
    var v47 string = "foo bar"
    if true {
        var v48 string = (((v47 + v47) + v47) + "")
        if (((3.14159 + (3.14159 * 0.0)) >= -2.25) || true) {
            if (!true) {
                switch (v36 >> ((1) & 63)) {
                case 110:
                    v49 := GenS5{F0: 1, F1: 0, F2: v26}
                    _ = v49
                case -182:
                    v26 = (^(v26))
                case 1:
                    v36 = (v36 * (v36 << ((1) & 63)))
                default:
                    for fi50 := 0; fi50 < 6; fi50++ {
                        v51 := GenS8{F0: 288, F1: 121}
                        v52 := GenS1{F0: (v36 << ((255) & 63))}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
            } else {
                var v53 int64 = (v3 / (((v32 / (((v32 - v32)) | 1))) | 1))
                v26 = 1
                for fi54 := 0; fi54 < 17; fi54++ {
                    var v55 uint = 65535
                    var v56 int16 = 42
                    for fi57 := 0; fi57 < 15; fi57++ {
                        _ = len(v2)
                        fmt.Println(fi57, v55)
                        v58 := GenS5{F0: v55, F1: (v56 * v56), F2: (v26 / ((v26) | 1))}
                        _ = fi57
                        _ = v58
                    }
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                _ = v53
            }
            if (true || ((v26 != v26) && (false && (-1 != 1)))) {
                for fi59 := 0; fi59 < 18; fi59++ {
                    v48 = ("a" + v48)
                    v35 += (v35 | 2)
                    _ = fi59
                }
                fmt.Println(v35, v3, v36)
            } else {
                v35 = v35
            }
        }
        if (0 < 16) {
            var v60 int = int(v3)
            if len(v46) > 0 {
                v46[0] = (v60 * (v36 + v60))
            }
            v61 := GenS9{F0: (v35 << ((239) & 63))}
            _ = v60
            _ = v61
        } else {
            if len(v46) > 0 {
                v46[0] = v36
            }
        }
        _ = v48
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 uint = 255
        _ = fi62
        _ = v63
    }
    var v64 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v45
    _ = v46
    _ = v47
    _ = v64
}
