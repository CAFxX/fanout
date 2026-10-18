package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 bool }
type GenS3 struct { F0 uint; F1 int8; F2 bool }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int }
func genFn0() int16 {
    switch -2147483648 {
    case -1:
        if false {
            var v7 int8 = 42
            _ = v7
        }
    case 9223372036854775807:
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{(fi8 << ((4294967295) & 63))}
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 bool = true
                v12 := GenS1{F0: 987}
                v11 = (!false)
                if v11 {
                    var v13 bool = v11
                    v11 = v11
                    _ = v13
                } else {
                    if (!false) {
                        v9 = append(v9, (fi8 >> ((255) & 63)))
                        var v14 uint = 65535
                        v9 = append(v9, 280)
                        _ = v14
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    case 597:
        var v15 int8 = -1
        _ = v15
    }
    var v16 int64 = 42
    v16 += v16
    _ = v16
    return 42
}
func genFn1(p0 bool) uint8 {
    var v17 uint8 = 0
    var v18 int32 = -2147483648
    var v19 uint = 65535
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 1
    var v3 string = (("a" + "foo bar") + "\n\t")
    var v4 int16 = 851
    v5 := make(map[uint32]uint32)
    var v6 int16 = 1000
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int = (fi20 >> ((185) & 63))
        _ = fi20
        _ = v21
    }
    if (-9223372036854775808 != 42) {
        v22 := []int{}
        switch 32768 {
        case 4294967295:
            fmt.Println(v4, v2)
        case 9223372036854775807:
            switch -9223372036854775808 {
            case 747:
                fmt.Println(v4, v6)
            case 32767:
                var v23 bool = true
                _ = v23
            default:
                if ("" >= v3) {
                    v24 := GenS2{F0: v4, F1: (2147483647 > 858), F2: true}
                    switch (((("foo bar" + (v3 + "\n\t")) + v3) + (v3 + ("" + v3))) + "hello") {
                    case "":
                        v6 += (v6 ^ v4)
                    case "a":
                        if len(v22) > 0 {
                            v22[0] = 2147483647
                        }
                    case "\n\t":
                        var v25 int32 = -1
                        _ = v25
                    default:
                        v4 = v4
                    }
                    v2 = (v2 * v2)
                    _ = v24
                } else {
                    v26 := GenS3{F0: 0, F1: 0, F2: (v3 != v3)}
                    _ = v26
                }
            }
        case 42:
            for fi27 := 0; fi27 < 20; fi27++ {
                v6 -= v6
                _ = fi27
            }
        default:
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := v1[2]
                if false {
                    if false {
                        var v30 float64 = 1.5
                        fmt.Println(v4)
                        fmt.Println(v2, v29)
                        _ = v30
                    }
                    v2 += (v2 * (^(v2)))
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    _ = len(v22)
                    fmt.Println(fi31, fi28, v6)
                    _ = fi31
                }
                _ = fi28
                _ = v29
            }
        }
        switch 0 {
        case 337:
            var v32 int16 = v6
            _ = v32
        case 1000:
            v33 := v5[2147483648]
            _ = v33
        case -2147483648:
            v34 := GenS4{F0: (v2 - v2), F1: 924}
            _ = v34
        default:
            var v35 uint32 = 745
            _ = v35
        }
        _ = v22
    } else {
        if ((1.5 + (((1.5 / 1e10) + -2.25) + -2.25)) == 1.5) {
            var v36 uint32 = uint32(v2)
            var v37 uint8 = 143
            var v38 int64 = 9223372036854775807
            _ = v36
            _ = v37
            _ = v38
        }
    }
    switch 9223372036854775807 {
    case 32768:
        for fi39 := 0; fi39 < 19; fi39++ {
            v6 -= v4
            var v40 int = (-(fi39))
            var v41 float64 = float64((1.5 - (0.0 * 1.5)))
            _ = fi39
            _ = v40
            _ = v41
        }
    case 9223372036854775807:
        v2 = (v2 + (v2 << ((4294967295) & 63)))
    default:
        if false {
            v42 := GenS5{F0: -25, F1: 4294967295}
            for fi43 := 0; fi43 < 13; fi43++ {
                fmt.Println(v3, fi43, v6)
                for fi44 := 0; fi44 < 3; fi44++ {
                    fmt.Println(v2, v3, fi44)
                    var v45 uint = 18446744073709551615
                    var v46 int = 482
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                _ = fi43
            }
            var v47 float64 = ((0.1 + (float64(0.1) / -2.25)) - (0.0 / (float64(((0.1 * 0.0) + float32((0.1 + 1e10)))) + 0.0)))
            _ = v42
            _ = v47
        } else {
            v4 = v6
            for fi48 := 0; fi48 < 17; fi48++ {
                v49 := []int{}
                _ = fi48
                _ = v49
            }
        }
    }
    var v50 uint64 = 2
    for fi51 := 0; fi51 < 16; fi51++ {
        v52 := []int{(fi51 >> ((866) & 63)), (fi51 >> ((769) & 63)), 32768, (fi51 - (^(fi51)))}
        _ = fi51
        _ = v52
    }
    v3 = v3
    v53 := GenS6{F0: -2147483648}
    v3 = (v3 + ((v3 + v3) + v3))
    v54 := []int{4294967295, -9223372036854775808}
    for fi55 := 0; fi55 < 9; fi55++ {
        for fi56 := 0; fi56 < 17; fi56++ {
            fmt.Println(v4, fi55)
            if len(v54) > 0 {
                v54[0] = fi56
            }
            var v57 int8 = 127
            v58 := GenS2{F0: (v6 * v4), F1: false, F2: (18446744073709551615 > v50)}
            _ = fi56
            _ = v57
            _ = v58
        }
        fmt.Println(v4, v4)
        fmt.Println(v3)
        _ = fi55
    }
    v6 += genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v6
}
