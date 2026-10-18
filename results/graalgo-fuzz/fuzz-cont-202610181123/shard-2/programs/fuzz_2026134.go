package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8) uint {
    v6 := []int{-85, 42, 9223372036854775807}
    var v7 int64 = 9223372036854775807
    switch 32768 {
    case -247:
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint16 = 946
                var v10 uint64 = 368
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    case -66:
        v11 := GenS1{F0: 127}
        _ = v11
    default:
        v7 = 0
    }
    _ = p0
    _ = v6
    _ = v7
    return 540
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 1
    var v4 int8 = 127
    var v5 uint64 = (v2 | (^(v2)))
    if ("\n\t" > "a") {
        v12 := GenS2{F0: 255, F1: 2147483647}
        if false {
            v13 := GenS1{F0: 127}
            if (2147483647 <= 4294967295) {
                if (!(true && (0 == 1000))) {
                    var v14 uint16 = 42
                    _ = v14
                } else {
                    v15 := []int{-9223372036854775808, 4294967295, -780}
                    var v16 float64 = 1e10
                    _ = v15
                    _ = v16
                }
            } else {
                v3 = (^(v3))
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch -894 {
                    case 0:
                        var v18 float64 = 0.0
                        _ = v18
                    case 32768:
                        if (0 != -1) {
                            var v19 bool = (2 == 0)
                            _ = v19
                        }
                    case -9223372036854775808:
                        var v20 float64 = 0.1
                        _ = v20
                    }
                    switch "foo bar" {
                    case "hello":
                        if false {
                            v2 = v2
                        } else {
                            v3 = 42
                        }
                    default:
                        fmt.Println(v5)
                    }
                    switch 1 {
                    case 42:
                        v2 = 4294967295
                    case 375:
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v3 += v3
                            _ = fi21
                        }
                    }
                    _ = fi17
                }
            }
            _ = v13
        } else {
            fmt.Println(v4, v3)
            switch -651 {
            case 127:
                fmt.Println(v3, v3)
            case 0:
                if (!true) {
                    v22 := []int{}
                    if true {
                        v1[2147483648] = true
                    }
                    fmt.Println(v5)
                    _ = v22
                } else {
                    v2 += 561
                    var v23 int = 127
                    v24 := GenS3{F0: 2147483647, F1: false}
                    _ = v23
                    _ = v24
                }
            default:
                v25 := []int{128, 0}
                _ = v25
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v28 := GenS4{F0: (v2 % ((0) | 1))}
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        }
        var v29 string = ("\n\t" + "foo bar")
        _ = v12
        _ = v29
    } else {
        v30 := GenS5{F0: v2}
        v31 := []int{0, 258, -9223372036854775808}
        if (true || (!(!false))) {
            for fi32 := 0; fi32 < 10; fi32++ {
                v5 = (^(v5))
                _ = fi32
            }
            v4 += (v4 % ((127) | 1))
            fmt.Println(v4, v5, v5)
        } else {
            var v33 int16 = 1
            _ = v33
        }
        _ = v30
        _ = v31
    }
    if true {
        v34 := GenS2{F0: 255, F1: 489}
        for fi35 := 0; fi35 < 5; fi35++ {
            v5 = (v5 - v5)
            _ = fi35
        }
        _ = v34
    } else {
        fmt.Println(v4, v3, v4)
        if false {
            for fi36 := 0; fi36 < 14; fi36++ {
                var v37 int16 = 108
                _ = fi36
                _ = v37
            }
            _ = len(v1)
            v38 := GenS5{F0: (v5 << ((354) & 63))}
            _ = v38
        }
    }
    v39 := []int{}
    v40 := GenS5{F0: (^(v2))}
    for fi41 := 0; fi41 < 12; fi41++ {
        if true {
            var v42 uint16 = 42
            _ = v42
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            v4 -= v4
            _ = fi43
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            v45 := GenS4{F0: (v2 / ((v2) | 1))}
            v46 := v1[807]
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi41
    }
    var v47 int8 = (-(v4))
    if true {
        if false {
            v48 := GenS3{F0: 0, F1: true}
            _ = v48
        } else {
            v4 -= (v47 % ((v4) | 1))
            var v49 float32 = ((0.0 * (-2.25 - 1.5)) / 3.14159)
            var v50 bool = false
            _ = v49
            _ = v50
        }
    }
    _ = len(v1)
    v2 -= 0
    var v51 int32 = (-(v3))
    v1[790] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v51
}
