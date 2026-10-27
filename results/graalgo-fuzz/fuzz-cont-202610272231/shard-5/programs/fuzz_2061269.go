package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 uint8 = 251
    v5 := make(map[int]uint32)
    v4 -= 127
    fmt.Println(v4)
    var v6 uint16 = 442
    v7 := []int{0}
    v8 := GenS1{F0: 171}
    v4 -= (v4 | (v4 / ((uint8(v6)) | 1)))
    v4 += v4
    var v9 int16 = 1000
    v9 = 0
    var v10 string = ("foo bar" + "")
    fmt.Println(v4, v4)
    _ = len(v7)
    if ((float64(-2.25) + (3.14159 * (-2.25 * ((1e10 / 0.1) - (-2.25 / 1.5))))) == 3.14159) {
        v11 := v2[145]
        v12 := GenS1{F0: 174}
        _ = v11
        _ = v12
    } else {
        if (!(!((!((-9223372036854775808 != 9223372036854775807) || (v4 >= v4))) && false))) {
            for fi13 := 0; fi13 < 20; fi13++ {
                if (!true) {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v15 := GenS1{F0: 2}
                        if false {
                            fmt.Println(fi14)
                            v16 := GenS1{F0: 42}
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    v17 := GenS2{F0: (224 < 287), F1: (fi13 / ((32767) | 1))}
                    _ = v17
                } else {
                    v9 += 42
                    if false {
                        v4 -= v4
                        v18 := GenS1{F0: 147}
                        _ = v18
                    }
                }
                fmt.Println(v9, v10, fi13)
                if true {
                    var v19 float32 = 0.0
                    v7 = append(v7, (-(fi13)))
                    var v20 uint = 297
                    _ = v19
                    _ = v20
                } else {
                    switch (fi13 ^ -1) {
                    case 9223372036854775807:
                        if (127 < 1) {
                            _ = len(v3)
                            v9 += 32767
                            v7 = append(v7, fi13)
                        } else {
                            var v21 int = (fi13 / ((74) | 1))
                            _ = v21
                        }
                    default:
                        var v22 float32 = (float32(1e10) + float32((float64((float32((1.5 * 1.5)) * float32(0.1))) + -2.25)))
                        _ = v22
                    }
                    var v23 float32 = 0.1
                    v4 += (v4 | v4)
                    _ = v23
                }
                _ = fi13
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                switch (fi24 % ((fi24) | 1)) {
                case 624:
                    if (false || false) {
                        v6 += v6
                        var v25 uint = 0
                        _ = v25
                    } else {
                        v4 += (v4 / ((v4) | 1))
                        var v26 uint64 = 0
                        v27 := GenS1{F0: 197}
                        _ = v26
                        _ = v27
                    }
                case 0:
                    v10 = ((v10 + v10) + "foo bar")
                case 9223372036854775807:
                    var v28 uint8 = 255
                    _ = v28
                }
                _ = fi24
            }
        }
        v9 -= (v9 & (v9 + 0))
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint8 = 26
            _ = fi29
            _ = v30
        }
    }
    if (128 > 1) {
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(fi31, fi31, v9)
            _ = fi31
        }
        fmt.Println(v4)
    } else {
        switch ("hello" + (v10 + "")) {
        case "foo bar":
            v9 = v9
        }
        v10 = (v10 + "a")
    }
    v32 := v5[-1]
    fmt.Println(v10, v10, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
