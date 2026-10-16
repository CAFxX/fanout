package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int, p1 uint32) int {
    fmt.Println(p1, p0)
    switch 42 {
    case -807:
        switch p0 {
        case 9223372036854775807:
            fmt.Println(p1)
        default:
            var v6 uint8 = 128
            _ = v6
        }
    default:
        v7 := []int{(p0 - (p0 | (p0 % ((p0) | 1)))), (p0 << ((80) & 63)), (p0 | p0), p0}
        _ = v7
    }
    fmt.Println(p1, p0)
    fmt.Println(p0, p0)
    p0 += (p0 >> ((721) & 63))
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 int32 = -1
    var v2 float64 = ((((0.0 - -2.25) + 0.1) + 0.0) + (0.0 / (-2.25 / (1.5 / 0.1))))
    var v3 int8 = 110
    var v4 int16 = -992
    v5 := []string{"", "\n\t"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v8 := GenS1{F0: 42}
    v9 := GenS2{F0: (v4 - (v4 * (v4 - 1)))}
    v1 -= 776
    if (!((255 > 0) || false)) {
        if len(v5) > 0 {
            v5[0] = ""
        }
    }
    if false {
        if ((1.5 + 1e10) >= ((3.14159 / 0.1) * 0.1)) {
            switch 42 {
            case 704:
                var v10 int16 = (v4 + v4)
                _ = v10
            case 128:
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 string = "\n\t"
                    if false {
                        var v13 string = "\n\t"
                        if (240 > 1) {
                            v14 := GenS2{F0: (^(v4))}
                            _ = v14
                        }
                        _ = v13
                    } else {
                        var v15 float32 = 1e10
                        v15 = v15
                        v16 := GenS1{F0: 0}
                        _ = v15
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                }
            }
            v17 := GenS2{F0: v4}
            v18 := GenS2{F0: v4}
            _ = v17
            _ = v18
        }
        v2 = float64(0.0)
        switch "\n\t" {
        case "\n\t":
            switch 2147483647 {
            case 42:
                if len(v5) > 0 {
                    v5[0] = ""
                }
            case 726:
                var v19 string = "foo bar"
                _ = v19
            case 128:
                var v20 float32 = 0.0
                _ = v20
            }
        }
    }
    switch 127 {
    case 42:
        switch ("a" + "foo bar") {
        case "a":
            var v21 int16 = (v4 | v4)
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 15; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    switch "a" {
                    case "":
                        v2 = v2
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            v2 -= (v2 + v2)
                        }
                    case "foo bar":
                        switch "" {
                        case "a":
                            v5 = append(v5, "\n\t")
                        default:
                            v1 += v1
                            if (v4 >= (v4 + (v4 * 0))) {
                                fmt.Println(v4, v4, v3)
                                fmt.Println(v1, v2)
                                var v24 int32 = v1
                                _ = v24
                            } else {
                                _ = len(v5)
                                var v25 uint32 = 0
                                if (true && true) {
                                    v3 += (v3 >> ((2) & 63))
                                }
                                _ = v25
                            }
                        }
                    }
                    v2 = (float64(0.1) * 0.1)
                    v26 := GenS2{F0: int16(v1)}
                    v5 = append(v5, "a")
                    _ = len(v5)
                    _ = fi23
                    _ = v26
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v28 float64 = -2.25
                    _ = fi27
                    _ = v28
                }
                _ = fi22
            }
        case "\n\t":
            v4 += (v4 >> ((1) & 63))
        }
    case 0:
        v1 = v1
    }
    v29 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
