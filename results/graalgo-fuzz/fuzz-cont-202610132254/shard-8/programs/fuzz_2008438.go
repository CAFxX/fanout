package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 uint16) bool {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{(fi4 & -2147483648), 956, (-(fi4)), fi4}
        if true {
            switch ("a" + "\n\t") {
            case "":
                switch "hello" {
                case "":
                    fmt.Println(p0, fi4)
                case "foo bar":
                    for fi6 := 0; fi6 < 18; fi6++ {
                        var v7 uint64 = 18446744073709551615
                        _ = fi6
                        _ = v7
                    }
                }
            }
            v8 := GenS1{F0: 78}
            switch (fi4 - 32767) {
            case -293:
                for fi9 := 0; fi9 < 17; fi9++ {
                    switch "a" {
                    case "foo bar":
                        var v10 uint = 2
                        _ = v10
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "":
                        fmt.Println(fi9, fi4, fi9)
                    }
                    p0 = p0
                    fmt.Println(fi9, fi4, p0)
                    v5 = append(v5, fi9)
                    fmt.Println(p0)
                    p0 = p0
                    _ = fi9
                }
            case 32768:
                v11 := GenS1{F0: -128}
                _ = v11
            case -9223372036854775808:
                fmt.Println(fi4, fi4)
            default:
                if true {
                    v12 := GenS1{F0: 0}
                    if (1 != 2) {
                        var v13 int64 = -478
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v13
                    }
                    p0 = p0
                    _ = v12
                }
            }
            _ = v8
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                p0 += 255
                _ = fi14
            }
        }
        p0 -= p0
        _ = fi4
        _ = v5
    }
    var v15 int16 = -1
    _ = p0
    _ = v15
    return ((127 > -6) || false)
}

func main() {
    v1 := []string{"", "hello"}
    v2 := []uint32{2, 2147483647, 2147483647, 833}
    v3 := make(map[int]int)
    v16 := v3[127]
    fmt.Println(v16)
    if (-786 <= 9223372036854775807) {
        v17 := GenS1{F0: 54}
        if len(v1) > 0 {
            v1[0] = ""
        }
        v16 += 9223372036854775807
        _ = v17
    }
    v16 += (-(v16))
    v18 := GenS1{F0: -1}
    var v19 float64 = 1.5
    var v20 uint32 = 359
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 10; fi21++ {
            if false {
                v20 = 2147483647
                if (!true) {
                    var v22 int64 = -1
                    _ = v22
                } else {
                    v19 = 3.14159
                    _ = len(v1)
                }
            } else {
                v23 := GenS2{F0: "hello", F1: 443}
                if ((!(v20 > v20)) && false) {
                    if (v16 > -377) {
                        switch "" {
                        case "hello":
                            var v24 int8 = 15
                            _ = v24
                        case "foo bar":
                            var v25 int8 = 42
                            _ = v25
                        }
                        v20 += v20
                        if ((v20 - (v20 & (v20 * (v20 * v20)))) <= 2147483648) {
                            v26 := v3[714]
                            _ = v26
                        } else {
                            v2 = append(v2, (v20 - v20))
                            switch "foo bar" {
                            case "\n\t":
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    var v28 float32 = (float32((float64(((1.5 * (float32(v19) + ((0.0 - (1.5 + (0.1 - (1.5 - 0.0)))) / 1.5))) / ((-2.25 / 1.0) * -2.25))) / 3.14159)) / (-2.25 + (-2.25 / (0.1 * (3.14159 - 0.1)))))
                                    _ = fi27
                                    _ = v28
                                }
                            case "foo bar":
                                var v29 uint32 = v20
                                _ = v29
                            }
                        }
                    }
                }
                switch v16 {
                case -2147483648:
                    v16 = (^(fi21))
                }
                _ = v23
            }
            var v30 uint16 = 528
            v20 -= 385
            _ = fi21
            _ = v30
        }
    default:
        v16 = (v16 * v16)
    }
    var v31 int8 = 76
    v16 += (^(v16))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
}
