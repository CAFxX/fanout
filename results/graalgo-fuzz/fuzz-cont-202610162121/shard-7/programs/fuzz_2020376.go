package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 uint16, p2 int) uint8 {
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, (fi7 % ((p2) | 1))}
        _ = fi7
        _ = v8
    }
    if (!true) {
        v9 := []int{p2, (p2 * -721)}
        p2 -= (^(p2))
        if false {
            var v10 int8 = 100
            switch p0 {
            case "hello":
                var v11 bool = false
                _ = v11
            default:
                v12 := GenS1{F0: 1000, F1: 0, F2: 2147483647}
                _ = v12
            }
            _ = v10
        } else {
            var v13 string = p0
            _ = v13
        }
        _ = v9
    }
    v14 := []int{p2, (p2 - -1), (p2 & -2147483648)}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 128
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]bool)
    var v3 int8 = 42
    var v4 uint = 65535
    var v5 uint8 = 127
    v15 := []int{2147483647, -40}
    var v16 uint8 = (v5 << ((18446744073709551615) & 63))
    v3 = (v3 + 1)
    if ((65535 >= 2) || (!false)) {
        var v17 int64 = 1
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: v3}
    if true {
        var v19 string = "a"
        v20 := GenS1{F0: 596, F1: 1000, F2: v1}
        v4 = v4
        _ = v19
        _ = v20
    }
    var v21 uint32 = 42
    if (1.5 != -2.25) {
        var v22 uint32 = (v21 * (v21 % ((v21) | 1)))
        var v23 int8 = v3
        var v24 int8 = v23
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 int16 = 1000
        _ = v25
    }
    switch "\n\t" {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            v4 -= (v4 | (v4 & (v4 * (v4 & v4))))
        case "":
            v2[42] = (true && false)
        case "foo bar":
            switch "a" {
            case "foo bar":
                for fi26 := 0; fi26 < 1; fi26++ {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    switch 1 {
                    case 32768:
                        switch fi26 {
                        case -826:
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 int8 = v3
                                v3 += v28
                                if true {
                                    v2[1] = (9223372036854775807 <= 1)
                                }
                                v3 = v3
                                fmt.Println(v1)
                                _ = fi27
                                _ = v28
                            }
                        }
                    }
                    _ = fi26
                }
            default:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
        default:
            v16 -= 42
        }
    case "":
        var v29 bool = (!true)
        _ = v29
    case "\n\t":
        v16 += v5
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 bool = (!((false || false) || true))
                v33 := v2[1]
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: (v21 / (((v21 / ((v21) | 1))) | 1)), F1: 42}
            v35 := GenS3{F0: v4, F1: 1000, F2: 65535}
            _ = fi30
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
