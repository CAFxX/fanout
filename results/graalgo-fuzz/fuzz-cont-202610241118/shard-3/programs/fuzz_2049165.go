package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32, p1 uint64) bool {
    switch "hello" {
    case "hello":
        fmt.Println(p1, p0)
    default:
        var v5 int = -1
        _ = v5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        p1 = 2
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: false, F2: 1}
    switch 42 {
    case 32768:
        switch 689 {
        case 127:
            v9 := []int{-9223372036854775808, -230}
            _ = v9
        case 0:
            switch "" {
            case "\n\t":
                p0 += (p0 ^ p0)
            case "a":
                if (42 > 9223372036854775807) {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        if true {
                            if (65535 <= 133) {
                                v11 := GenS1{F0: 255, F1: false, F2: 2}
                                switch "a" {
                                case "\n\t":
                                    var v12 int = 1
                                    _ = v12
                                case "foo bar":
                                    p0 = (p0 - p0)
                                }
                                _ = v11
                            }
                            v13 := []int{fi10, fi10, (fi10 << ((922) & 63))}
                            v14 := GenS1{F0: 18446744073709551615, F1: true, F2: 18}
                            _ = v13
                            _ = v14
                        }
                        _ = fi10
                    }
                    if ((uint16(p1) != 255) && (!(120 >= 0))) {
                        p0 = (p0 * (p0 - (p0 + (p0 & -1))))
                    }
                    var v15 int8 = 1
                    _ = v15
                }
            case "hello":
                v16 := []int{-2147483648}
                _ = v16
            }
        case 628:
            if false {
                v17 := []int{-9223372036854775808}
                _ = v17
            }
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 42
    var v3 int32 = -2147483648
    v4 := make(map[uint32]string)
    var v18 uint16 = 32767
    if (42 >= 0) {
        fmt.Println(v1, v2, v3)
        delete(v4, 4294967295)
    } else {
        v18 = v18
        var v19 uint64 = 4294967295
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(fi20, v3)
            var v21 uint16 = 0
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    if false {
        v18 -= (v18 & v18)
    } else {
        if (!true) {
            v1 = (v1 ^ (v1 - 0))
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := []int{}
                var v24 uint8 = 128
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            if ((1.5 - ((1e10 * (0.0 / 0.1)) / (1e10 / float32((0.1 * ((-2.25 * (float64(1e10) + (3.14159 + 0.0))) + ((-2.25 + -2.25) - 0.1))))))) < (0.0 * 1e10)) {
                var v25 uint64 = 716
                v18 -= (^(v18))
                _ = v25
            }
            v2 -= (v2 - -9223372036854775808)
        }
        _ = len(v4)
        switch 4294967295 {
        case -1:
            v26 := GenS2{F0: 0, F1: 733}
            _ = v26
        default:
            var v27 uint8 = 255
            _ = v27
        }
    }
    fmt.Println(v1, v1)
    var v28 uint16 = v18
    switch ("a" + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v1)
    }
    v3 = -1
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 int64 = (v2 * (v2 << ((4294967295) & 63)))
            v32 := []int{}
            delete(v4, 1)
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 uint8 = 224
        var v34 float64 = (0.1 * (-2.25 * (0.0 / 1.5)))
        _ = fi29
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
