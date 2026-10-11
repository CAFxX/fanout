package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int; F1 int64 }
func genFn0(p0 int8, p1 string, p2 uint16) int16 {
    v4 := GenS1{F0: 42, F1: p1, F2: 221}
    var v5 int16 = -327
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: int32(fi6), F1: (p1 + ("foo bar" + p1)), F2: 0}
        if (!false) {
            v8 := GenS2{F0: -9223372036854775808}
            _ = v8
        } else {
            var v9 float32 = 0.1
            _ = v9
        }
        p1 = ((p1 + p1) + "\n\t")
        _ = fi6
        _ = v7
    }
    v10 := []int{42, 2147483647, -9223372036854775808, 32767}
    p1 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 uint32, p1 bool) uint64 {
    for fi11 := 0; fi11 < 16; fi11++ {
        p0 += (p0 - p0)
        _ = fi11
    }
    var v12 float64 = 3.14159
    if true {
        v13 := GenS1{F0: -996, F1: "a", F2: 156}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                p0 += (p0 & p0)
            }
            if true {
                var v15 int32 = -1
                v16 := []int{-776, (fi14 << ((0) & 63)), (fi14 + (^(fi14))), fi14}
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                p1 = false
                p1 = (-435 >= -1)
                _ = fi17
            }
            _ = fi14
        }
        if ((0.0 + (-2.25 - 1e10)) <= 3.14159) {
            var v18 float32 = 1.5
            var v19 float64 = 0.0
            for fi20 := 0; fi20 < 5; fi20++ {
                if (("foo bar" + "\n\t") < "hello") {
                    p0 -= p0
                    switch "\n\t" {
                    case "\n\t":
                        v12 += (v12 * 1e10)
                    }
                } else {
                    var v21 uint64 = 1
                    if (!(-778 == -1)) {
                        var v22 float32 = v18
                        if p1 {
                            fmt.Println(v21, p1)
                        }
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v21 -= v21
                        var v24 uint16 = 887
                        var v25 uint32 = (p0 >> ((739) & 63))
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                }
                switch 32767 {
                case 888:
                    var v26 float64 = v19
                    _ = v26
                case 4294967295:
                    var v27 uint8 = 128
                    _ = v27
                default:
                    switch ("a" + ("" + ("a" + "hello"))) {
                    case "":
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v29 := GenS3{F0: 282, F1: fi28}
                            _ = fi28
                            _ = v29
                        }
                    case "\n\t":
                        var v30 float64 = float64((v18 + -2.25))
                        _ = v30
                    }
                }
                _ = fi20
            }
            _ = v18
            _ = v19
        } else {
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := []int{(fi31 ^ fi31), fi31}
                _ = fi31
                _ = v32
            }
        }
        v33 := GenS4{F0: 968, F1: 42, F2: 1}
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v12
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 906
    var v3 uint16 = 255
    switch -395 {
    case 4294967295:
        v3 -= v3
    case -2147483648:
        switch "foo bar" {
        case "":
            v34 := []int{-69, -9223372036854775808, 32768, -98}
            _ = v34
        case "foo bar":
            v2 = v2
        case "hello":
            v35 := []int{42, -1}
            _ = v35
        }
    case -426:
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 uint16 = (v3 << ((1) & 63))
            v3 -= v3
            v3 -= (v37 ^ 699)
            _ = fi36
            _ = v37
        }
    }
    v38 := []int{-2147483648, 127, -2147483648}
    switch 377 {
    case 1:
        v2 = v2
    case 32767:
        if len(v38) > 0 {
            v38[0] = 24
        }
    case -2147483648:
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 int64 = 876
            if true {
                v40 += (v40 | 42)
                var v41 int64 = (v40 ^ (v40 % (((v40 & (v40 | 0))) | 1)))
                _ = v41
            }
            _ = fi39
            _ = v40
        }
    default:
        var v42 int32 = 1
        _ = v42
    }
    v43 := GenS3{F0: 781, F1: 127}
    v44 := GenS5{F0: 100, F1: 42}
    var v45 int32 = -2147483648
    v45 -= -996
    if (-1 >= -68) {
        delete(v1, 1)
        v46 := GenS6{F0: 9223372036854775807, F1: -1}
        _ = v46
    } else {
        fmt.Println(v3, v3)
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 bool = false
            _ = fi47
            _ = v48
        }
        fmt.Println(v3, v2, v3)
    }
    var v49 int8 = 42
    var v50 int8 = (v49 % (((-(v49))) | 1))
    fmt.Println(v50)
    if (0 >= 1) {
        if (true && true) {
            var v51 string = "\n\t"
            _ = v51
        }
        fmt.Println(v45, v2, v50)
        v50 = v49
    }
    switch -2147483648 {
    case -2147483648:
        if ((v2 / (((v2 >> ((4294967295) & 63))) | 1)) >= (v2 ^ 0)) {
            fmt.Println(v50)
            v52 := GenS7{F0: 849, F1: 558}
            v3 += (^(v3))
            _ = v52
        } else {
            fmt.Println(v50, v3, v3)
        }
    case -9223372036854775808:
        v38 = append(v38, 9223372036854775807)
    default:
        v50 = (^(v50))
    }
    switch "" {
    case "\n\t":
        v38 = append(v38, 2147483647)
    case "":
        if len(v38) >= 2 {
            v38 = v38[1:len(v38)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v43
    _ = v44
    _ = v45
    _ = v49
    _ = v50
}
