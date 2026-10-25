package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int; F1 int; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 int16) bool {
    v7 := []int{}
    if true {
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(p1, fi8, p0)
            fmt.Println(p1)
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return (-128 <= 42)
}
func genFn1(p0 uint8, p1 int8, p2 int64) uint64 {
    v9 := []int{32768}
    var v10 int32 = 1
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint32 = 807
        var v13 int = (-(fi11))
        fmt.Println(v12, v10)
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v10 += v10
        if false {
            switch ("" + "a") {
            case "a":
                v15 := GenS1{F0: (!(true || (p1 >= p1))), F1: 1, F2: 42}
                _ = v15
            }
            fmt.Println(v10, v10, p1)
            if len(v9) > 0 {
                v9[0] = -257
            }
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 237
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := []uint32{841, 2147483648}
    var v5 int32 = -2147483648
    v6 := make(map[int]uint32)
    if false {
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS1{F0: (1 >= 4294967295), F1: fi16, F2: 485}
            if true {
                if (true || ((((!((false && true) || (0.1 != (1e10 - 3.14159)))) || true) || (v5 >= (-(v5)))) && false)) {
                    var v18 float32 = 0.1
                    v18 = v18
                    v18 += v18
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v3[fi16] = 72
                    _ = fi19
                }
                switch (fi16 + 32767) {
                case 643:
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 uint64 = 617
                        v22 := GenS2{F0: (fi16 + (fi16 % ((fi20) | 1))), F1: 2147483647}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    var v23 string = ("a" + "")
                    _ = v23
                }
            }
            _ = fi16
            _ = v17
        }
        var v24 string = ((("\n\t" + "a") + "\n\t") + ("hello" + "hello"))
        _ = v24
    }
    if false {
        v25 := GenS3{F0: -911, F1: -356, F2: -9223372036854775808}
        v6[1] = 0
        _ = v25
    } else {
        v5 += (v5 & (v5 * 42))
    }
    v26 := GenS4{F0: (true || false)}
    v5 += v5
    v27 := GenS4{F0: (42 >= 127)}
    switch "foo bar" {
    case "hello":
        if (1 == 42) {
            fmt.Println(v5)
            v28 := GenS4{F0: true}
            v29 := v1[("a" + "\n\t")]
            _ = v28
            _ = v29
        }
    }
    v4 = append(v4, 65535)
    v5 -= v5
    for fi30 := 0; fi30 < 13; fi30++ {
        v5 += 42
        v5 = v5
        _ = fi30
    }
    var v31 int64 = -1
    if (1 >= 32767) {
        for fi32 := 0; fi32 < 7; fi32++ {
            fmt.Println(v31)
            var v33 float64 = (0.1 / 1e10)
            v31 = (v31 / ((0) | 1))
            _ = fi32
            _ = v33
        }
        var v34 uint64 = 1
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        v6[-39] = 65535
        delete(v2, "hello")
        v36 := GenS5{F0: 128, F1: v5}
        _ = fi35
        _ = v36
    }
    switch "\n\t" {
    case "a":
        if true {
            for fi37 := 0; fi37 < 3; fi37++ {
                if (true || (259 >= 2)) {
                    v38 := GenS1{F0: false, F1: 0, F2: 973}
                    _ = v38
                }
                if false {
                    var v39 uint8 = 127
                    v39 = 128
                    _ = v39
                } else {
                    var v40 int16 = -1
                    v41 := GenS4{F0: (!(false || false))}
                    _ = v40
                    _ = v41
                }
                for fi42 := 0; fi42 < 15; fi42++ {
                    v5 = (^(v5))
                    _ = fi42
                }
                _ = fi37
            }
            var v43 int = 1
            v31 = (v31 << ((870) & 63))
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 uint32 = 2147483648
        v45 -= 2
        v46 := GenS1{F0: true, F1: 42, F2: 255}
        _ = fi44
        _ = v45
        _ = v46
    }
    v5 -= v5
    var v47 int32 = (v5 / ((v5) | 1))
    v47 = v47
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
