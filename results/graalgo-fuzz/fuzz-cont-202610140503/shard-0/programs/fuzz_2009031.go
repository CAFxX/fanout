package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int, p1 uint8, p2 uint64) uint16 {
    var v5 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            p2 += (^(p2))
            _ = fi7
        }
        _ = fi6
    }
    if false {
        v8 := []int{(^(p0)), -2147483648}
        fmt.Println(p2)
        v9 := GenS1{F0: 42, F1: (p2 & 850), F2: 1}
        _ = v8
        _ = v9
    } else {
        p2 -= 63
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1() bool {
    v10 := []int{-797, 127, -9223372036854775808}
    var v11 uint16 = 255
    _ = v10
    _ = v11
    return (!(((1.5 / 3.14159) + -2.25) != (-2.25 + 3.14159)))
}

func main() {
    var v1 float32 = float32((0.0 - float64((1e10 / 1.0))))
    var v2 bool = false
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    var v12 string = ""
    for fi13 := 0; fi13 < 9; fi13++ {
        v3[(fi13 >> ((2) & 63))] = 2147483647
        v2 = true
        v2 = v2
        _ = fi13
    }
    v14 := []int{-920}
    if v2 {
        var v15 float32 = v1
        v16 := GenS1{F0: 127, F1: 2, F2: 1}
        v1 -= 1e10
        _ = v15
        _ = v16
    } else {
        var v17 int32 = -1
        _ = v17
    }
    v12 = v12
    v12 = v12
    for fi18 := 0; fi18 < 14; fi18++ {
        v3[fi18] = 0
        v14 = append(v14, (fi18 << ((134) & 63)))
        _ = fi18
    }
    v1 += 3.14159
    v19 := GenS1{F0: 0, F1: 54, F2: 255}
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            if (false || (80 == 104)) {
                v1 -= 1.5
                v22 := GenS1{F0: 100, F1: 2, F2: 2}
                if len(v14) > 0 {
                    v14[0] = 1
                }
                _ = v22
            } else {
                v23 := GenS2{F0: "", F1: 1, F2: 128}
                var v24 int = (fi20 | (^(fi20)))
                _ = v23
                _ = v24
            }
            switch (v12 + (((v12 + ("foo bar" + "\n\t")) + "") + v12)) {
            case "a":
                fmt.Println(fi21)
            default:
                fmt.Println(v2, v12)
            }
            fmt.Println(fi21)
            _ = fi21
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 string = v12
            var v27 int16 = 1000
            v1 += v1
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi20
    }
    if (v2 || v2) {
        var v28 float64 = (0.1 * 3.14159)
        _ = v28
    }
    if false {
        if len(v14) > 0 {
            v14[0] = 0
        }
        var v29 string = v12
        var v30 bool = false
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS3{F0: 100, F1: fi31}
        var v33 int32 = 2147483647
        if v2 {
            var v34 float64 = 1e10
            v35 := GenS3{F0: 127, F1: fi31}
            for fi36 := 0; fi36 < 4; fi36++ {
                v34 -= v34
                var v37 int = 384
                var v38 uint64 = 18446744073709551615
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v34
            _ = v35
        } else {
            var v39 uint = 211
            var v40 int8 = 0
            v41 := GenS2{F0: v12, F1: 42, F2: uint8(v39)}
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(fi42)
        v1 = v1
        _ = fi42
    }
    var v43 string = v12
    if ((1.5 * float32((1.5 / -2.25))) == (-2.25 * ((v1 + v1) + v1))) {
        for fi44 := 0; fi44 < 17; fi44++ {
            fmt.Println(fi44, v12, v1)
            v14 = append(v14, (fi44 & fi44))
            _ = fi44
        }
        if ((18446744073709551615 >= 0) || v2) {
            v2 = true
            v45 := GenS1{F0: 6, F1: 417, F2: 255}
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 uint = 2
                fmt.Println(v1)
                _ = fi46
                _ = v47
            }
            _ = v45
        } else {
            v48 := GenS2{F0: "a", F1: 575, F2: 127}
            _ = v48
        }
        var v49 int32 = -2147483648
        _ = v49
    } else {
        var v50 int32 = 1
        for fi51 := 0; fi51 < 15; fi51++ {
            var v52 uint32 = 47
            var v53 int64 = -1
            var v54 uint64 = 0
            _ = fi51
            _ = v52
            _ = v53
            _ = v54
        }
        _ = v50
    }
    v55 := v3[127]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v12
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v55
}
