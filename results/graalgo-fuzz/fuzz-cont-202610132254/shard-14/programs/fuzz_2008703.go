package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 int32, p1 int32, p2 int) int64 {
    if false {
        p1 = int32(p2)
        v4 := GenS1{F0: 100, F1: 21, F2: (p1 >> ((18446744073709551615) & 63))}
        v5 := []int{p2, (p2 % ((p2) | 1))}
        _ = v4
        _ = v5
    }
    var v6 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 uint8) uint8 {
    var v7 float64 = 1e10
    if true {
        fmt.Println(p0, p0)
        p0 -= p0
        v8 := []int{956, 0, -1}
        _ = v8
    } else {
        v9 := GenS1{F0: -1, F1: 2, F2: 0}
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = -1
            _ = fi10
            _ = v11
        }
        var v12 uint64 = 405
        _ = v12
        _ = v9
    }
    v13 := []int{-809, 361, 32768, 9223372036854775807}
    var v14 uint = 68
    v7 -= (0.1 / float64((float32(v7) * (-2.25 + 0.0))))
    _ = p0
    _ = v13
    _ = v14
    _ = v7
    return (^(p0))
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    var v3 float32 = float32(float64((1e10 + 0.0)))
    var v15 uint32 = 1
    v16 := GenS2{F0: 1}
    v17 := GenS3{F0: 667}
    var v18 int8 = -34
    for fi19 := 0; fi19 < 19; fi19++ {
        v18 = 118
        for fi20 := 0; fi20 < 12; fi20++ {
            v2[(v15 << ((973) & 63))] = ""
            _ = fi20
        }
        _ = fi19
    }
    switch "a" {
    case "\n\t":
        var v21 uint16 = (v1 ^ (v1 ^ 360))
        _ = v21
    case "foo bar":
        switch -1 {
        case 0:
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int = -223
                _ = fi22
                _ = v23
            }
        case -820:
            _ = len(v2)
        case 32768:
            switch 4294967295 {
            case -539:
                v24 := []int{2147483647, 42}
                _ = v24
            }
        default:
            v25 := GenS4{F0: -294, F1: 32767, F2: (("" + "a") + "")}
            _ = v25
        }
    default:
        v18 = v18
    }
    var v26 uint32 = (v15 >> ((65535) & 63))
    if true {
        v27 := GenS2{F0: genFn0(-1, 2147483647, 127)}
        v2[v26] = "\n\t"
        v1 = (v1 | v1)
        _ = v27
    }
    v28 := v2[(^(v26))]
    fmt.Println(v28, v15)
    v29 := []int{-2147483648, 128}
    switch 531 {
    case 42:
        if len(v29) > 0 {
            v29[0] = 0
        }
    }
    switch "foo bar" {
    case "\n\t":
        v3 += -2.25
    case "":
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
    case "foo bar":
        v28 = (v28 + v28)
    default:
        v30 := GenS3{F0: 95}
        _ = v30
    }
    v18 -= v18
    fmt.Println(v15)
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            switch v28 {
            case "hello":
                var v33 uint32 = (v26 / ((0) | 1))
                _ = v33
            default:
                var v34 uint = 255
                _ = v34
            }
            if true {
                if (v18 < -1) {
                    switch fi32 {
                    case 0:
                        var v35 float64 = 3.14159
                        var v36 string = "\n\t"
                        _ = v35
                        _ = v36
                    case 32768:
                        var v37 uint32 = v15
                        _ = v37
                    default:
                        if (v1 != (^(v1))) {
                            fmt.Println(v1, v18)
                        } else {
                            v18 = (-(v18))
                            v2[2147483648] = v28
                        }
                    }
                    v2[65535] = v28
                    v26 -= v15
                    var v38 int16 = -53
                    v39 := GenS2{F0: genFn0(-877, 271, (fi31 - -9223372036854775808))}
                    _ = v38
                    _ = v39
                } else {
                    v28 = ((("foo bar" + v28) + v28) + "a")
                    v26 += (^(v26))
                }
                v26 = 409
            } else {
                var v40 string = ("\n\t" + "hello")
                _ = v40
            }
            _ = fi32
        }
        _ = fi31
    }
    v26 += v26
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
}
