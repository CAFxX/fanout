package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0() uint {
    v6 := []int{32767, -1}
    var v7 uint = 1
    var v8 string = (("\n\t" + "\n\t") + "")
    _ = v6
    _ = v7
    _ = v8
    return (v7 ^ v7)
}
func genFn1(p0 uint64) int16 {
    var v9 int32 = 2147483647
    if false {
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    var v3 int32 = 1
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "hello" {
        case "":
            switch fi11 {
            case 9223372036854775807:
                v12 := []int{(fi11 & (fi11 >> ((genFn0()) & 63)))}
                _ = v12
            case -629:
                switch "a" {
                case "foo bar":
                    for fi13 := 0; fi13 < 3; fi13++ {
                        if false {
                            var v14 uint8 = 255
                            if (false || true) {
                                var v15 int32 = (v3 - v3)
                                _ = v15
                            }
                            delete(v4, (fi13 % ((0) | 1)))
                            var v16 int8 = -13
                            _ = v14
                            _ = v16
                        } else {
                            var v17 uint16 = 621
                            switch ((("hello" + "foo bar") + ("hello" + "\n\t")) + "hello") {
                            case "a":
                                v18 := []int{-63}
                                _ = v18
                            }
                            var v19 float32 = ((-2.25 * (-2.25 + (3.14159 - 0.0))) * float32(-2.25))
                            var v20 int8 = 100
                            var v21 uint64 = 570
                            _ = v17
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS1{F0: 42, F1: 0, F2: 79}
                        _ = fi13
                        _ = v22
                    }
                default:
                    v3 += 0
                }
            default:
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    v3 -= v3
                default:
                    v4[(fi11 % ((fi11) | 1))] = 617
                }
            }
        case "a":
            v2 += 0.0
        case "foo bar":
            v23 := []int{fi11, -499, fi11}
            _ = v23
        default:
            var v24 int32 = v3
            _ = v24
        }
        _ = fi11
    }
    v3 += (v3 << ((326) & 63))
    var v25 uint16 = 1
    fmt.Println(v2, v2, v25)
    switch "foo bar" {
    case "a":
        v2 = (((3.14159 / -2.25) * (v2 + v2)) / 1e10)
    case "":
        switch -1 {
        case 0:
            v25 -= v25
        }
    case "foo bar":
        var v26 int32 = (v3 / ((652) | 1))
        _ = v26
    default:
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float64 = v2
            _ = fi27
            _ = v28
        }
    }
    var v29 float64 = v2
    v30 := []int{-1, 2147483647, -968, 834}
    var v31 uint32 = 1
    if (v3 <= v3) {
        var v32 uint8 = 0
        var v33 int8 = 127
        _ = v32
        _ = v33
    } else {
        var v34 int = 593
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if len(v30) > 0 {
            v30[0] = (fi35 / ((fi35) | 1))
        }
        _ = fi35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v2 = float64(1e10)
        if false {
            var v37 uint16 = v25
            var v38 string = "foo bar"
            var v39 int = (^(fi36))
            _ = v37
            _ = v38
            _ = v39
        } else {
            fmt.Println(v31, fi36)
        }
        _ = len(v30)
        _ = fi36
    }
    v40 := GenS1{F0: 2147483648, F1: 4294967295, F2: 18446744073709551615}
    var v41 float32 = (-2.25 * 3.14159)
    fmt.Println(v29)
    for fi42 := 0; fi42 < 18; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            v3 += v3
            v44 := GenS2{F0: (v25 & (v25 >> ((uint(v31)) & 63))), F1: 0, F2: (^(v31))}
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    v45 := v1[v31]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
}
