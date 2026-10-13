package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int) string {
    p0 = p0
    p0 -= (p0 - (p0 << ((65535) & 63)))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint8 = 37
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += 988
        _ = fi7
    }
    switch "a" {
    case "a":
        if (false && true) {
            p0 += 127
            if true {
                for fi8 := 0; fi8 < 8; fi8++ {
                    fmt.Println(p0, fi8)
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 int64 = -1
                        var v11 float64 = 1e10
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    _ = fi8
                }
            } else {
                fmt.Println(p0)
            }
        } else {
            for fi12 := 0; fi12 < 4; fi12++ {
                if false {
                    p0 += (p0 >> ((4294967295) & 63))
                    v13 := GenS1{F0: 250, F1: 225, F2: 1}
                    p0 -= p0
                    _ = v13
                } else {
                    var v14 uint64 = 1
                    v15 := []int{fi12}
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        }
    }
    _ = p0
    return (("a" + "foo bar") + "hello")
}
func genFn1(p0 int64) uint16 {
    fmt.Println(p0)
    p0 -= p0
    var v16 string = ""
    _ = p0
    _ = v16
    return 640
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    var v3 int = 9223372036854775807
    var v4 int32 = 1
    v3 += v3
    v3 += (v3 << ((494) & 63))
    v4 -= (-(v4))
    if (18446744073709551615 > 344) {
        var v17 int64 = -1
        v18 := []int{}
        _ = v17
        _ = v18
    } else {
        v2 = v2
    }
    v19 := GenS1{F0: 1, F1: 18446744073709551615, F2: 18446744073709551615}
    switch 32767 {
    case 0:
        for fi20 := 0; fi20 < 8; fi20++ {
            v3 += (fi20 - v3)
            v4 = v4
            var v21 uint8 = v2
            _ = fi20
            _ = v21
        }
    case -1:
        v22 := GenS2{F0: 376, F1: -128}
        _ = v22
    }
    fmt.Println(v4, v2)
    fmt.Println(v3, v4)
    var v23 int32 = (v4 / (((v4 % ((v4) | 1))) | 1))
    v2 = (v2 & v2)
    v24 := GenS2{F0: 1, F1: 0}
    fmt.Println(v4, v23, v23)
    switch (v3 | (v3 >> ((18446744073709551615) & 63))) {
    case -859:
        if false {
            v3 += v3
            v25 := GenS2{F0: 2, F1: 127}
            v1["a"] = (!(!((-(v3)) < (v3 ^ (v3 * (v3 * v3))))))
            _ = v25
        } else {
            v26 := GenS2{F0: 415, F1: 0}
            v27 := GenS3{F0: 42, F1: 625}
            _ = v26
            _ = v27
        }
    case -1:
        v2 = v2
    case -759:
        switch ("a" + "") {
        case "a":
            v2 = 255
        case "foo bar":
            if ((-(v23)) <= 42) {
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 string = "\n\t"
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := GenS4{F0: (v2 | 252)}
                    for fi32 := 0; fi32 < 4; fi32++ {
                        fmt.Println(fi32)
                        v3 -= (fi30 << ((908) & 63))
                        switch fi30 {
                        case 262:
                            var v33 float64 = 1e10
                            _ = v33
                        case 792:
                            switch fi32 {
                            case 32768:
                                fmt.Println(v4, v4, fi32)
                            default:
                                v23 = (v23 * v23)
                            }
                        default:
                            var v34 int32 = (v4 - (v23 + -711))
                            _ = v34
                        }
                        _ = fi32
                    }
                    _ = fi30
                    _ = v31
                }
                fmt.Println(v3, v23)
            }
        case "\n\t":
            for fi35 := 0; fi35 < 2; fi35++ {
                v23 += v23
                v36 := v1["foo bar"]
                _ = fi35
                _ = v36
            }
        }
    }
    v37 := GenS4{F0: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v4
}
