package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint) int64 {
    var v4 uint8 = 204
    v5 := []int{-9223372036854775808, 4294967295}
    switch "a" {
    case "a":
        if true {
            if false {
                switch 4294967295 {
                case 9223372036854775807:
                    var v6 int32 = 2147483647
                    _ = v6
                case 32767:
                    for fi7 := 0; fi7 < 16; fi7++ {
                        var v8 uint8 = (v4 & v4)
                        for fi9 := 0; fi9 < 17; fi9++ {
                            p0 = p0
                            v10 := GenS1{F0: 42, F1: 1, F2: (true || (18446744073709551615 >= 2))}
                            _ = fi9
                            _ = v10
                        }
                        _ = fi7
                        _ = v8
                    }
                case 0:
                    if (uint32(v4) < 4294967295) {
                        var v11 float64 = 1e10
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v13 := GenS2{F0: 1000}
                            p0 -= 417
                            var v14 uint16 = 0
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        fmt.Println(p0, p0)
                        v15 := GenS3{F0: 1, F1: "a"}
                        _ = v15
                    }
                default:
                    var v16 float32 = ((((1.5 - float32(-2.25)) / 0.1) + (float32(1.5) * (0.0 / 1.0))) - 0.1)
                    _ = v16
                }
                if true {
                    p0 = p0
                    var v17 uint32 = 2147483648
                    _ = v17
                }
                var v18 int = 4294967295
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                v5 = append(v5, fi19)
                _ = fi19
            }
        }
    }
    fmt.Println(p0, v4)
    var v20 int32 = -1
    _ = p0
    _ = v20
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 bool, p1 uint16, p2 bool) int8 {
    var v21 uint8 = 2
    if (!p2) {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := []int{32767, (fi22 >> ((527) & 63)), fi22}
            v21 += v21
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS1{F0: -34, F1: 127, F2: (!(127 != 4294967295))}
    v25 := []int{}
    for fi26 := 0; fi26 < 20; fi26++ {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        if (1 <= v21) {
            var v27 int8 = int8(p1)
            p2 = (!((0.1 + 0.1) < 3.14159))
            switch (fi26 << ((0) & 63)) {
            case 32767:
                fmt.Println(p2, p1, p1)
            default:
                var v28 int = fi26
                _ = v28
            }
            _ = v27
        }
        _ = fi26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v24
    _ = v25
    return 42
}

func main() {
    var v1 int32 = 197
    v2 := make(map[int]uint32)
    var v3 float32 = (3.14159 - (1e10 - 3.14159))
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 int16 = 0
        v1 = (v1 + -330)
        var v31 bool = false
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := []int{(-(fi32)), fi32, 2147483647, (-(fi32))}
        _ = fi32
        _ = v33
    }
    fmt.Println(v1, v3)
    v34 := GenS2{F0: 1000}
    v1 -= (v1 - 2147483647)
    var v35 int16 = 42
    if (1e10 != float32((0.0 * 1e10))) {
        var v36 uint32 = 2147483648
        var v37 int8 = 100
        switch -888 {
        case 506:
            var v38 uint8 = 1
            _ = v38
        }
        _ = v36
        _ = v37
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            v1 = -2147483648
            _ = fi39
        }
        fmt.Println(v3, v3)
    }
    v40 := GenS4{F0: -101}
    switch -9223372036854775808 {
    case -2147483648:
        var v41 int32 = 0
        _ = v41
    default:
        switch 32767 {
        case -339:
            var v42 uint32 = 533
            _ = v42
        case 32767:
            var v43 uint16 = 1
            _ = v43
        case -2147483648:
            v44 := GenS4{F0: 127}
            _ = v44
        }
    }
    var v45 int16 = (v35 >> ((42) & 63))
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS5{F0: fi46}
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v40
    _ = v45
}
