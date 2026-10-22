package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 int16, p1 int32, p2 bool) uint32 {
    var v6 int64 = 649
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    v1 := []int{703}
    v2 := make(map[string]int)
    var v3 uint16 = 1
    v4 := []string{("\n\t" + (("hello" + "") + (("foo bar" + "\n\t") + ("\n\t" + "\n\t")))), "\n\t"}
    var v5 int8 = int8(v3)
    fmt.Println(v3)
    var v7 uint32 = 599
    for fi8 := 0; fi8 < 1; fi8++ {
        if true {
            v3 = (v3 >> ((4294967295) & 63))
        } else {
            v5 = 100
            switch "foo bar" {
            case "foo bar":
                var v9 float32 = (3.14159 + float32(1e10))
                _ = v9
            case "":
                switch (fi8 / ((fi8) | 1)) {
                case 566:
                    v3 = (v3 << ((0) & 63))
                case -11:
                    v5 += v5
                default:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if ((v3 >> ((255) & 63)) < (v3 & (v3 >> ((42) & 63)))) {
                            switch (fi8 - -9223372036854775808) {
                            case 32768:
                                switch (fi10 / (((fi10 >> ((2) & 63))) | 1)) {
                                case 9223372036854775807:
                                    var v11 int16 = -783
                                    var v12 uint16 = v3
                                    _ = v11
                                    _ = v12
                                case 128:
                                    fmt.Println(v7, v7, v7)
                                case 42:
                                    switch "\n\t" {
                                    case "":
                                        var v13 int8 = (v5 - v5)
                                        _ = v13
                                    }
                                }
                            case -135:
                                var v14 uint16 = (v3 * 660)
                                _ = v14
                            default:
                                v3 = 262
                                switch "hello" {
                                case "hello":
                                    _ = len(v2)
                                case "foo bar":
                                    v7 += genFn0(-1, -183, false)
                                }
                                var v15 int16 = 42
                                _ = v15
                            }
                        } else {
                            v7 = (v7 << ((255) & 63))
                            if (fi10 > (fi10 << ((752) & 63))) {
                                fmt.Println(fi10)
                            } else {
                                v16 := GenS1{F0: 0, F1: 113}
                                var v17 float32 = 0.1
                                _ = v16
                                _ = v17
                            }
                            if (405 <= 42) {
                                delete(v2, "foo bar")
                            }
                        }
                        _ = fi10
                    }
                }
            default:
                fmt.Println(fi8, fi8)
            }
            _ = len(v2)
        }
        var v18 uint8 = 0
        _ = fi8
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(fi19, fi19)
        v20 := GenS1{F0: (v3 ^ (v3 & (v3 * (v3 / ((v3) | 1))))), F1: 255}
        _ = fi19
        _ = v20
    }
    if true {
        var v21 bool = false
        _ = v21
    } else {
        v22 := GenS2{F0: 0, F1: v3, F2: 1}
        var v23 uint = 1
        var v24 string = "a"
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 float64 = (3.14159 + (3.14159 * 0.1))
    delete(v2, "\n\t")
    var v26 uint = 1
    var v27 float64 = ((0.1 - float64(-2.25)) + (v25 - (((((((((((v25 / v25) - (3.14159 - 3.14159)) * 0.1) * v25) + 1.5) + v25) / (v25 * ((v25 - v25) - ((1.5 * v25) / ((float64(3.14159) + 1.5) - v25))))) - (0.0 / -2.25)) * (v25 * v25)) * v25) / v25)))
    switch 561 {
    case -334:
        v5 -= v5
    case 32768:
        if (v27 == -2.25) {
            if true {
                var v28 int16 = -32768
                v25 += 1e10
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v28
            } else {
                var v29 float32 = 0.0
                switch 2147483647 {
                case 387:
                    delete(v2, "")
                }
                _ = v29
            }
        }
    default:
        var v30 bool = false
        _ = v30
    }
    v7 -= (v7 * (v7 ^ v7))
    fmt.Println(v5)
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
