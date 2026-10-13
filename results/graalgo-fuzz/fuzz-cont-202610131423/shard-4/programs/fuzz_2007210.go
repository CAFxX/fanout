package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 int16) uint {
    var v5 int64 = 1
    if (!true) {
        v5 += v5
        v5 = v5
    } else {
        p0 = p0
        v6 := []int{-9223372036854775808, 128, 0}
        fmt.Println(v5, p0, v5)
        _ = v6
    }
    v7 := []int{9223372036854775807, 44, 32768, 2147483647}
    if ((!(p1 >= p1)) || p0) {
        v8 := GenS1{F0: 1, F1: 0, F2: 18446744073709551615}
        var v9 int = 0
        for fi10 := 0; fi10 < 12; fi10++ {
            v5 += v5
            switch (((("a" + "\n\t") + "\n\t") + "foo bar") + "\n\t") {
            case "":
                for fi11 := 0; fi11 < 7; fi11++ {
                    if ((100 <= -128) && false) {
                        if ((-(v5)) > (v5 & -382)) {
                            var v12 bool = p0
                            v9 = 42
                            _ = v12
                        }
                        if false {
                            var v13 string = "foo bar"
                            _ = v13
                        } else {
                            switch (v9 - (v9 + 9223372036854775807)) {
                            case 128:
                                var v14 uint8 = 1
                                _ = v14
                            case -799:
                                switch "a" {
                                case "":
                                    var v15 float32 = 0.0
                                    _ = v15
                                case "foo bar":
                                    p0 = (1 == 18446744073709551615)
                                case "\n\t":
                                    var v16 int32 = 0
                                    _ = v16
                                }
                            default:
                                v5 += (-(v5))
                            }
                        }
                        if p0 {
                            var v17 uint64 = 619
                            fmt.Println(fi10, fi10, v9)
                            _ = v17
                        }
                    }
                    if (false || (0 <= -1)) {
                        var v18 uint32 = 65535
                        p0 = false
                        _ = v18
                    }
                    _ = fi11
                }
            case "\n\t":
                p0 = (false || p0)
            case "foo bar":
                if false {
                    if len(v7) > 0 {
                        v7[0] = (v9 + 350)
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            default:
                v19 := GenS2{F0: 2}
                _ = v19
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS3{F0: 2}
            if ("foo bar" == ("foo bar" + ("\n\t" + "a"))) {
                for fi22 := 0; fi22 < 20; fi22++ {
                    if (65535 == 255) {
                        fmt.Println(p1)
                        var v23 int64 = 42
                        _ = v23
                    } else {
                        v5 = 9223372036854775807
                        var v24 uint = 18446744073709551615
                        fmt.Println(p0, fi22, v24)
                        _ = v24
                    }
                    v5 = v5
                    switch 42 {
                    case 4294967295:
                        if true {
                            v25 := GenS3{F0: 2}
                            _ = v25
                        }
                    }
                    _ = fi22
                }
            }
            fmt.Println(fi20)
            _ = fi20
            _ = v21
        }
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = 0
    var v2 int64 = -87
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v1 = 100
    v2 -= (v2 >> ((600) & 63))
    v26 := []int{32767, 852, 2147483647}
    var v27 int32 = -1
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS3{F0: 255}
        _ = fi28
        _ = v29
    }
    fmt.Println(v1, v27, v1)
    var v30 float32 = float32((float64(float32(float64(-2.25))) / 3.14159))
    v31 := GenS2{F0: 42}
    var v32 int8 = v1
    var v33 int8 = (v32 | v1)
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    if false {
        var v34 float64 = -2.25
        var v35 uint64 = 816
        for fi36 := 0; fi36 < 16; fi36++ {
            v2 -= 42
            _ = fi36
        }
        _ = v34
        _ = v35
    } else {
        if (42 >= -1) {
            v2 = 42
        } else {
            v30 = float32((1.5 / float64((1e10 + (0.1 - (((3.14159 - v30) - v30) * (v30 + v30)))))))
            if len(v26) > 0 {
                v26[0] = 4294967295
            }
        }
        switch int(v2) {
        case 42:
            var v37 uint64 = 2
            _ = v37
        }
        v30 = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
