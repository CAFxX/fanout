package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 int16, p2 uint32) uint16 {
    p2 += p2
    v4 := GenS1{F0: 686, F1: 32767, F2: 42}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 895
}
func genFn1() string {
    switch "hello" {
    case "foo bar":
        for fi5 := 0; fi5 < 10; fi5++ {
            if false {
                for fi6 := 0; fi6 < 8; fi6++ {
                    fmt.Println(fi5, fi6)
                    v7 := []int{}
                    v8 := GenS1{F0: 42, F1: 32767, F2: -1}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                if (1e10 < 1e10) {
                    var v9 int = (fi5 & 32768)
                    for fi10 := 0; fi10 < 1; fi10++ {
                        fmt.Println(v9, v9, v9)
                        if (42 < -4) {
                            var v11 float32 = 0.0
                            _ = v11
                        }
                        _ = fi10
                    }
                    var v12 float64 = (1.5 / 1.0)
                    _ = v12
                    _ = v9
                }
            } else {
                if true {
                    v13 := []int{fi5}
                    if len(v13) > 0 {
                        v13[0] = fi5
                    }
                    switch fi5 {
                    case 632:
                        var v14 uint32 = 604
                        _ = v14
                    }
                    _ = v13
                }
                v15 := GenS1{F0: 0, F1: 255, F2: 42}
                _ = v15
            }
            _ = fi5
        }
    }
    if true {
        var v16 int = -530
        _ = v16
    }
    var v17 float64 = (float64((((3.14159 + 1.5) / 1.5) * 1.5)) * (0.0 * (3.14159 / 0.1)))
    _ = v17
    return ("hello" + (("a" + "") + "a"))
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v2[32767] = -265
    v18 := []int{1, 128, -9223372036854775808, -1}
    v1 -= (v1 / ((127) | 1))
    v1 -= 163
    v1 -= (v1 ^ (^(v1)))
    fmt.Println(v1)
    switch 457 {
    case 128:
        delete(v2, 9223372036854775807)
    default:
        if (42 != 999) {
            var v19 int64 = -1
            switch -1 {
            case -1:
                fmt.Println(v19, v1)
            case -9223372036854775808:
                switch (("hello" + "\n\t") + "a") {
                case "foo bar":
                    var v20 float32 = (float32(((0.1 / -2.25) / 1.0)) + (-2.25 - (0.0 + ((((3.14159 - 1.5) + float32(0.1)) - ((3.14159 + (-2.25 / float32(float64(1.5)))) + 0.0)) + float32(3.14159)))))
                    _ = v20
                case "a":
                    if len(v18) > 0 {
                        v18[0] = -2147483648
                    }
                case "":
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 int = (fi21 / (((fi21 + 128)) | 1))
                        _ = fi21
                        _ = v22
                    }
                }
            default:
                var v23 uint16 = 221
                _ = v23
            }
            fmt.Println(v19, v19)
            _ = v19
        }
    }
    fmt.Println(v1)
    v1 = v1
    if (255 <= v1) {
        v1 = v1
    } else {
        if (!false) {
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
