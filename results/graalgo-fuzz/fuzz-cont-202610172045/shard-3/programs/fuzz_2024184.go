package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint8, p1 int32) int16 {
    if (p1 <= p1) {
        var v7 uint = 255
        var v8 uint32 = 0
        var v9 int = 4294967295
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 65535, F1: 1}
        _ = v10
    }
    p0 = (^(p0))
    fmt.Println(p0, p0)
    if (1 == 582) {
        v11 := GenS1{F0: 2147483648, F1: 4294967295}
        v12 := GenS1{F0: 355, F1: 2}
        p0 = (p0 / ((p0) | 1))
        _ = v11
        _ = v12
    }
    if (true || true) {
        var v13 uint = 4294967295
        v14 := []int{}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int, p1 int8) uint8 {
    for fi15 := 0; fi15 < 8; fi15++ {
        p1 -= p1
        var v16 float64 = 3.14159
        _ = fi15
        _ = v16
    }
    var v17 float64 = float64(0.1)
    for fi18 := 0; fi18 < 20; fi18++ {
        p1 -= 37
        v19 := []int{-1, (fi18 + (fi18 & fi18)), (fi18 ^ fi18)}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        switch fi20 {
        case 128:
            var v21 uint64 = 992
            _ = v21
        case 1:
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS2{F0: ((p1 < (-(p1))) || (true && (787 != 285))), F1: 2147483647}
                var v24 uint = 65535
                p0 = 127
                _ = fi22
                _ = v23
                _ = v24
            }
        case -2147483648:
            for fi25 := 0; fi25 < 15; fi25++ {
                p0 -= -2147483648
                if (!true) {
                    if true {
                        fmt.Println(fi25, p1, v17)
                        v26 := []int{(p0 / ((fi20) | 1)), (p0 + -962), 32767}
                        _ = v26
                    }
                    p0 += fi20
                } else {
                    fmt.Println(fi25, p0, v17)
                }
                _ = fi25
            }
        }
        _ = fi20
    }
    v27 := []int{p0, 32767}
    _ = p0
    _ = p1
    _ = v17
    _ = v27
    return 1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 float32 = float32(3.14159)
    v4 := []string{"foo bar", (("" + "\n\t") + ("" + "hello"))}
    v5 := []uint32{0, 654, 42, 4294967295}
    v6 := make(map[uint32]bool)
    v1 = v1
    if (-1 > 42) {
        v1 = v1
        v28 := GenS1{F0: 93, F1: 42}
        _ = v28
    } else {
        v3 = 0.0
    }
    var v29 float32 = v3
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    var v30 int64 = 42
    var v31 string = "foo bar"
    v30 = (v30 - v30)
    _ = len(v6)
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(v30)
        _ = len(v4)
        _ = fi32
    }
    fmt.Println(v29)
    if false {
        var v33 int8 = 100
        var v34 int8 = v33
        var v35 int16 = 42
        _ = v33
        _ = v34
        _ = v35
    }
    if (2 != 42) {
        v1 = (v1 & v1)
        switch -23 {
        case -797:
            for fi36 := 0; fi36 < 1; fi36++ {
                for fi37 := 0; fi37 < 11; fi37++ {
                    v29 -= (((1.5 * ((v29 + v3) * float32((1e10 - (float64(3.14159) - 0.1))))) + v29) + ((v3 - float32(float64(v29))) - v29))
                    _ = fi37
                }
                var v38 int16 = -32768
                _ = fi36
                _ = v38
            }
        case 42:
            v39 := GenS3{F0: (-(v1)), F1: 4294967295}
            _ = v39
        default:
            if ((!(!(!true))) && true) {
                var v40 uint = 2
                _ = v40
            }
        }
        v29 += ((float32(((0.1 - (1e10 + float64(v29))) + -2.25)) + (v29 - (v29 + (((1.5 / float32(1.5)) / (v29 + v3)) + ((float32(-2.25) * v3) * v3))))) * v29)
    }
    var v41 int = (v1 << ((42) & 63))
    v41 += 2147483647
    var v42 string = "foo bar"
    var v43 uint64 = 18446744073709551615
    v44 := GenS1{F0: 303, F1: 18446744073709551615}
    switch (v1 + v41) {
    case 127:
        v1 += (v41 - (-(v1)))
    case 673:
        v45 := v2[65535]
        _ = v45
    default:
        var v46 int = int(v30)
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
