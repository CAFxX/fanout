package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int16, p1 int16, p2 bool) uint8 {
    var v4 float64 = (0.0 - (float64(0.1) / 1.5))
    v4 = v4
    v5 := []int{2147483647, 32768}
    v4 = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 168
}
func genFn1() int {
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            var v7 int = (^(fi6))
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v7)
                _ = fi8
            }
            v7 -= 0
            _ = v7
        } else {
            var v9 float32 = 0.0
            _ = v9
        }
        var v10 int64 = -874
        _ = fi6
        _ = v10
    }
    var v11 int = 32767
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 int32 = 42
    var v2 string = ""
    v3 := []int{}
    v2 = ((v2 + v2) + ((v2 + v2) + ((v2 + (("foo bar" + v2) + (((v2 + v2) + "foo bar") + "\n\t"))) + v2)))
    v3 = append(v3, genFn1())
    var v12 uint8 = 3
    var v13 bool = false
    var v14 float64 = (3.14159 + 3.14159)
    if (!v13) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 32767 {
        case 4294967295:
            if v13 {
                if ((v14 * 0.1) != ((v14 / (3.14159 / 3.14159)) + (1.5 + v14))) {
                    if true {
                        fmt.Println(v14, v2)
                        var v15 uint8 = v12
                        _ = v15
                    }
                    var v16 int = 32767
                    switch (v2 + ("\n\t" + "\n\t")) {
                    case "":
                        if (!v13) {
                            v17 := GenS1{F0: 1}
                            _ = v17
                        } else {
                            var v18 float32 = 0.0
                            v1 = (v1 & (v1 ^ v1))
                            _ = v18
                        }
                    default:
                        v3 = append(v3, (v16 / (((v16 % ((v16) | 1))) | 1)))
                    }
                    _ = v16
                } else {
                    fmt.Println(v12, v2)
                    v19 := GenS1{F0: 255}
                    _ = v19
                }
                v12 += v12
                switch 0 {
                case 127:
                    if (v1 > 71) {
                        fmt.Println(v14, v14)
                    }
                case -2147483648:
                    _ = len(v3)
                }
            } else {
                fmt.Println(v1)
            }
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                v12 = (^(v12))
                _ = fi20
            }
        }
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS2{F0: 944, F1: 2, F2: v2}
            var v23 int32 = v1
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v12, v2)
    }
    v24 := GenS3{F0: false}
    var v25 string = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 uint16 = 800
    v1 += 990
    var v27 uint64 = 31
    fmt.Println(v1)
    fmt.Println(v25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 uint64 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
