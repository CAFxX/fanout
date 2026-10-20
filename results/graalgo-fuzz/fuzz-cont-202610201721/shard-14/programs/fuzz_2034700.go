package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 string, p1 bool) bool {
    p1 = true
    var v6 int16 = 580
    if (-2147483648 >= 0) {
        v6 = v6
        for fi7 := 0; fi7 < 1; fi7++ {
            v6 = (-(v6))
            fmt.Println(p0, p1)
            _ = fi7
        }
    }
    switch p0 {
    case "hello":
        switch p0 {
        case "\n\t":
            fmt.Println(p1, p0)
        case "foo bar":
            for fi8 := 0; fi8 < 19; fi8++ {
                if (945 <= 65535) {
                    v9 := []int{fi8, fi8}
                    if (p1 || p1) {
                        v6 = v6
                        if true {
                            if p1 {
                                p0 = p0
                            }
                            p1 = p1
                        } else {
                            var v10 uint32 = 580
                            _ = v10
                        }
                        if ((((255 < 29) || (-9223372036854775808 > -593)) || (100 == 18446744073709551615)) && p1) {
                            if p1 {
                                fmt.Println(v6, p0, p1)
                                switch fi8 {
                                case -491:
                                    var v11 int32 = 544
                                    _ = v11
                                case -1:
                                    fmt.Println(p1, p0)
                                }
                                v6 = v6
                            }
                            for fi12 := 0; fi12 < 1; fi12++ {
                                var v13 int32 = -2147483648
                                fmt.Println(v13, v13, fi8)
                                fmt.Println(fi8)
                                _ = fi12
                                _ = v13
                            }
                            if false {
                                var v14 uint = 1
                                var v15 int64 = 9223372036854775807
                                var v16 bool = false
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        }
                    } else {
                        fmt.Println(v6, p0)
                    }
                    fmt.Println(p1, fi8, v6)
                    _ = v9
                } else {
                    var v17 uint8 = 255
                    _ = v17
                }
                v18 := []int{fi8, (-(fi8)), fi8, fi8}
                _ = fi8
                _ = v18
            }
        default:
            var v19 uint = 0
            _ = v19
        }
    case "foo bar":
        var v20 string = p0
        _ = v20
    case "":
        var v21 int64 = 9223372036854775807
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v6
    return (896 > 255)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -693
    var v3 float32 = float32(float64(0.1))
    var v4 uint64 = 1
    v5 := make(map[int]uint32)
    if false {
        v3 = (3.14159 - -2.25)
    }
    if (!true) {
        v22 := GenS1{F0: 0, F1: -104}
        v3 -= -2.25
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 float32 = (v3 + 1.5)
        _ = fi23
        _ = v24
    }
    fmt.Println(v3)
    v25 := GenS2{F0: "foo bar", F1: -32768, F2: 2147483648}
    for fi26 := 0; fi26 < 16; fi26++ {
        v5[(-(fi26))] = 2147483648
        if (42 == 580) {
            switch fi26 {
            case 1:
                v27 := []int{}
                _ = v27
            case 32767:
                v28 := []int{(fi26 + fi26), fi26, fi26}
                _ = v28
            case 32768:
                for fi29 := 0; fi29 < 9; fi29++ {
                    if (255 < 65535) {
                        v30 := GenS1{F0: 127, F1: -128}
                        v2 += -9223372036854775808
                        var v31 float64 = 0.0
                        _ = v30
                        _ = v31
                    } else {
                        v3 = v3
                        v32 := []int{fi26, int(v4)}
                        v33 := GenS2{F0: "hello", F1: 1000, F2: 2147483648}
                        var v34 uint8 = 127
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(v3, fi26)
                    _ = fi29
                }
            }
        }
        v2 = v2
        _ = fi26
    }
    var v35 int64 = 839
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
