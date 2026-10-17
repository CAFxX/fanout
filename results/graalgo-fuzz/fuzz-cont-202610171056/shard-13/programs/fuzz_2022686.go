package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 bool = (3.14159 == (1e10 / 0.1))
    v2 := []int{-364, -9223372036854775808, 127}
    v3 := make(map[uint32]int)
    var v4 uint16 = 213
    _ = len(v2)
    var v5 uint = 255
    for fi6 := 0; fi6 < 15; fi6++ {
        if (false && true) {
            v7 := v3[2147483648]
            fmt.Println(v1, v1, v7)
            _ = v7
        }
        var v8 int32 = -695
        _ = fi6
        _ = v8
    }
    v3[521] = -2147483648
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    delete(v3, 756)
    if (-1 != -1) {
        v1 = false
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v5, v4, v4)
    } else {
        fmt.Println(v4, v4, v4)
    }
    v4 = v4
    switch "foo bar" {
    case "a":
        var v10 int64 = 42
        _ = v10
    default:
        if (1 < 802) {
            if ("hello" != "a") {
                if v1 {
                    if (!(v1 && true)) {
                        switch 32768 {
                        case 1:
                            var v11 int64 = 42
                            _ = v11
                        case 9223372036854775807:
                            fmt.Println(v1)
                        case 4294967295:
                            v4 -= (^(v4))
                        default:
                            switch -1 {
                            case -1:
                                switch -9223372036854775808 {
                                case 420:
                                    for fi12 := 0; fi12 < 19; fi12++ {
                                        switch -713 {
                                        case -796:
                                            var v13 float64 = 1e10
                                            _ = v13
                                        case 957:
                                            fmt.Println(fi12, v1, fi12)
                                        }
                                        _ = fi12
                                    }
                                case -1:
                                    var v14 int16 = -32768
                                    _ = v14
                                }
                            case 32768:
                                for fi15 := 0; fi15 < 19; fi15++ {
                                    v4 += v4
                                    fmt.Println(fi15, v4)
                                    fmt.Println(v5, v4, v1)
                                    _ = fi15
                                }
                            }
                        }
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v17 := GenS1{F0: v1}
                            var v18 string = "a"
                            _ = len(v2)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        switch fi19 {
                        case 4294967295:
                            v5 = 526
                        case -9223372036854775808:
                            var v20 uint16 = v4
                            _ = v20
                        }
                        switch ("a" + "foo bar") {
                        case "":
                            v1 = (false && v1)
                        case "\n\t":
                            v5 = 0
                        }
                        _ = fi19
                    }
                } else {
                    var v21 int64 = -828
                    _ = v21
                }
                var v22 uint16 = (v4 ^ v4)
                _ = v22
            }
        }
    }
    if ((2147483647 < 2) && v1) {
        v1 = ((3.14159 - ((float64((1.5 * 3.14159)) / 1.5) + 0.1)) < 1.5)
    }
    v5 += (^(v5))
    v5 += (v5 & (v5 | v5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
