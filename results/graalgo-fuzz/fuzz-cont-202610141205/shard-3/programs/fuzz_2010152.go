package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0() int {
    var v5 uint8 = 65
    if (525 <= 134) {
        fmt.Println(v5)
        if true {
            fmt.Println(v5)
            if false {
                for fi6 := 0; fi6 < 14; fi6++ {
                    v5 = v5
                    var v7 int = 2147483647
                    v8 := []int{v7, (fi6 + v7)}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            } else {
                if ((!false) || true) {
                    v9 := GenS1{F0: 42}
                    v5 -= v5
                    _ = v9
                }
            }
        }
    } else {
        v5 = (v5 ^ (^(v5)))
    }
    fmt.Println(v5)
    var v10 string = "hello"
    _ = v10
    _ = v5
    return 686
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    switch 32767 {
    case 9223372036854775807:
        v11 := []int{genFn0(), -986, 32767, 32768}
        _ = v11
    }
    var v12 int16 = 97
    if ((-(v1)) >= -1) {
        v12 = (v12 / ((v12) | 1))
    }
    v12 -= (-(v12))
    fmt.Println(v4)
    v4 = ((v4 / 0.1) + v4)
    v4 = 0.0
    var v13 float64 = float64(3.14159)
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 = (-(v1))
            fmt.Println(fi15)
            _ = fi15
        }
        _ = fi14
    }
    v16 := GenS1{F0: 42}
    var v17 uint16 = 706
    if (!false) {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 1}
            v20 := GenS2{F0: (-(fi18))}
            if (float64(-2.25) < (((v13 - -2.25) * float64(-2.25)) - ((v4 / float64((0.1 + 0.1))) + (v13 + ((v13 / ((((-2.25 * v4) + 1.5) * (float64((3.14159 * float32(v13))) - v13)) / 1e10)) / v13))))) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v3[(fi21 ^ fi21)] = 2147483648
                    _ = fi21
                }
                fmt.Println(v13, v12, fi18)
                switch -566 {
                case -674:
                    if (true && false) {
                        var v22 uint64 = 18446744073709551615
                        var v23 string = "foo bar"
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int = 4294967295
                        _ = v24
                    }
                }
            } else {
                v25 := GenS2{F0: 128}
                v26 := GenS3{F0: fi18, F1: 113}
                _ = v25
                _ = v26
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        v17 = v17
    }
    var v27 uint64 = 764
    v13 = v13
    fmt.Println(v17, v27)
    var v28 uint64 = (v27 + (v27 << ((255) & 63)))
    switch ("foo bar" + "\n\t") {
    case "a":
        if ((("hello" + "\n\t") + "") > ("foo bar" + "")) {
            v29 := []int{-966, -170}
            for fi30 := 0; fi30 < 6; fi30++ {
                v28 -= 0
                _ = fi30
            }
            _ = v29
        }
    case "hello":
        var v31 int32 = 0
        _ = v31
    default:
        v28 -= (^(v27))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
