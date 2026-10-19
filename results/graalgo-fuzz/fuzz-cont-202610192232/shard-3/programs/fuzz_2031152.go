package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    if (1 == 42) {
        var v7 int = 128
        v7 -= 42
        _ = v7
    } else {
        if ((!false) || true) {
            v8 := []int{-9223372036854775808, 370}
            if (!false) {
                var v9 float32 = ((float32(((float64(-2.25) - 1e10) / float64(-2.25))) - (0.1 * 1.5)) - 1.5)
                if false {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    if (false && ((!(2147483648 >= 71)) && false)) {
                        v9 += 0.1
                        v9 = v9
                    } else {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v9 = (v9 / -2.25)
                            _ = fi10
                        }
                        v9 = float32(3.14159)
                    }
                    fmt.Println(v9)
                    var v11 uint32 = 2147483647
                    _ = v11
                }
                _ = v9
            } else {
                v12 := GenS1{F0: 9223372036854775807}
                _ = v12
            }
            _ = v8
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 string = (("" + "a") + "\n\t")
                v16 := []int{fi13, fi13, (fi14 << ((255) & 63)), fi13}
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 1}
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(fi13, fi13)
                var v19 float64 = (1.5 * (3.14159 / 3.14159))
                _ = fi18
                _ = v19
            }
            _ = fi13
            _ = v17
        }
        v20 := []int{32768, -2147483648, -9223372036854775808, 4294967295}
        _ = v20
    }
    var v21 float32 = -2.25
    _ = v21
    return 244
}
func genFn1() string {
    var v22 int64 = -729
    if true {
        switch "\n\t" {
        case "a":
            v22 -= v22
        case "foo bar":
            fmt.Println(v22)
        case "":
            fmt.Println(v22)
        default:
            fmt.Println(v22)
        }
        v23 := []int{-2147483648, 903, -412, 127}
        _ = v23
    }
    var v24 int = 127
    _ = v22
    _ = v24
    return "\n\t"
}

func main() {
    var v1 int16 = -1
    var v2 uint = 590
    v3 := make(map[int]bool)
    var v4 int = -380
    var v5 int = int(v1)
    var v25 uint32 = 610
    v2 = (^(v2))
    v2 = 2
    v2 = (v2 * (v2 ^ v2))
    v26 := GenS2{F0: -543, F1: 2147483647}
    v4 += v4
    var v27 uint = v2
    var v28 uint8 = 121
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int8 = 42
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS3{F0: (127 < (v28 + v28))}
            v33 := []int{495, fi31, 32767}
            if (!("a" <= "a")) {
                v4 += fi29
                var v34 int64 = 0
                _ = v34
            } else {
                if (true || false) {
                    var v35 int16 = v1
                    _ = v35
                } else {
                    fmt.Println(v28, v2, v2)
                    v36 := GenS4{F0: ((v30 >= v30) && true)}
                    v37 := GenS1{F0: 42}
                    _ = v36
                    _ = v37
                }
                v5 = (fi29 & 128)
                if false {
                    v25 = 42
                    v1 -= v1
                }
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            if false {
                v27 += v27
            }
            fmt.Println(v27)
            _ = fi38
        }
        _ = fi29
        _ = v30
    }
    var v39 uint16 = 2
    v40 := GenS4{F0: true}
    var v41 uint32 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
