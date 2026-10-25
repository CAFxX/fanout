package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32) int32 {
    p0 += 2147483647
    var v4 string = "hello"
    _ = p0
    _ = v4
    return 2147483647
}

func main() {
    var v1 uint16 = 255
    var v2 int = 42
    v3 := make(map[string]int)
    if (0 < 42) {
        v2 = -444
    } else {
        if (!false) {
            v1 += (v1 ^ v1)
            if true {
                var v5 float64 = (0.1 - (0.1 + 3.14159))
                var v6 uint16 = 65535
                _ = v5
                _ = v6
            }
            for fi7 := 0; fi7 < 6; fi7++ {
                fmt.Println(v1, v1, v2)
                if (true || (255 < 0)) {
                    v8 := GenS1{F0: true, F1: 4294967295, F2: "foo bar"}
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := GenS2{F0: 2}
                        if (!false) {
                            var v11 uint = 417
                            if true {
                                v1 += (v1 & (v1 - (v1 << (((v11 | 402)) & 63))))
                            }
                            _ = v11
                        } else {
                            switch "hello" {
                            case "hello":
                                v2 += 32767
                                v12 := GenS2{F0: 578}
                                if true {
                                    v1 = (v1 ^ (v1 | (v1 >> ((4294967295) & 63))))
                                }
                                _ = v12
                            case "":
                                fmt.Println(fi9, v1)
                            case "a":
                                var v13 float64 = -2.25
                                _ = v13
                            }
                            _ = len(v3)
                        }
                        _ = fi9
                        _ = v10
                    }
                    v14 := v3[(("" + "\n\t") + "a")]
                    _ = v14
                    _ = v8
                }
                v2 = (v2 + (fi7 | fi7))
                _ = fi7
            }
        }
        v1 += (v1 / ((v1) | 1))
        v15 := []int{(v2 - v2), 4294967295, 32768, (v2 + (v2 * (^(v2))))}
        _ = v15
    }
    v16 := GenS3{F0: 810, F1: 42, F2: (("foo bar" + "hello") + "a")}
    v3["foo bar"] = v2
    var v17 uint16 = 435
    v17 += 42
    v1 = 604
    v1 = v1
    if (9223372036854775807 > v2) {
        v18 := []int{(v2 * 9223372036854775807), 42}
        _ = v18
    } else {
        if (true || ((float64(3.14159) * 0.0) <= 1.5)) {
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 = v17
                _ = fi19
            }
            switch "foo bar" {
            case "a":
                var v20 int32 = 1
                _ = v20
            }
            if (!true) {
                v21 := GenS2{F0: 379}
                v1 = (v17 ^ 0)
                fmt.Println(v1, v1)
                _ = v21
            }
        } else {
            var v22 int = v2
            var v23 uint16 = 2
            v24 := GenS4{F0: 4294967295}
            _ = v22
            _ = v23
            _ = v24
        }
    }
    v1 += (v17 - v17)
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := []int{(v2 / ((v2) | 1)), v2, (v2 / ((fi25) | 1)), (fi25 & -40)}
        _ = fi25
        _ = v26
    }
    v27 := []int{v2, v2, -2147483648}
    if true {
        var v28 uint16 = v1
        v2 = (v2 << ((1) & 63))
        fmt.Println(v2)
        _ = v28
    }
    v29 := GenS5{F0: -1}
    delete(v3, "")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v29
    _ = v3
}
