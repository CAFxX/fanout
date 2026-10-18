package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0() bool {
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 bool = (421 == -625)
        v7 := []int{(fi5 + fi5), fi5, (fi5 % (((fi5 << ((0) & 63))) | 1)), (fi5 / ((-458) | 1))}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v4 = 2147483647
    if false {
        v8 := []int{149}
        v4 += v4
        for fi9 := 0; fi9 < 13; fi9++ {
            v4 = v4
            _ = fi9
        }
        _ = v8
    }
    _ = v4
    return (!(false && false))
}
func genFn1(p0 int64, p1 uint8) int32 {
    v10 := []int{}
    var v11 uint64 = 47
    switch "" {
    case "a":
        _ = len(v10)
    case "foo bar":
        p0 += (p0 ^ 1)
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]bool)
    var v12 bool = (!(42 != 168))
    fmt.Println(v12, v2)
    v12 = v12
    v12 = (2147483647 < -677)
    var v13 float64 = 1.5
    var v14 float32 = 1.5
    if ((uint(v2) < 42) || v12) {
        v14 -= -2.25
        var v15 float64 = v13
        _ = v15
    } else {
        v16 := GenS1{F0: 858, F1: 79}
        var v17 uint32 = 65
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 32767, F1: 121}
    var v19 int64 = 9223372036854775807
    _ = len(v3)
    if true {
        v14 = v14
    } else {
        if (127 > 128) {
            v14 -= (v14 + v14)
            switch ("\n\t" + "a") {
            case "hello":
                switch "foo bar" {
                case "\n\t":
                    v19 = v19
                case "a":
                    if v12 {
                        if ((!((42 > v19) && (!((0.1 + v13) != v13)))) && (100 == -108)) {
                            v20 := v1[1]
                            if true {
                                switch ((v20 + (((v20 + ("hello" + v20)) + (((("" + ((v20 + v20) + v20)) + v20) + "foo bar") + v20)) + v20)) + ("foo bar" + "foo bar")) {
                                case "hello":
                                    for fi21 := 0; fi21 < 18; fi21++ {
                                        v19 = -495
                                        _ = fi21
                                    }
                                default:
                                    for fi22 := 0; fi22 < 12; fi22++ {
                                        if true {
                                            var v23 uint8 = 0
                                            fmt.Println(v20)
                                            var v24 bool = (0 < 0)
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 string = (v20 + (v20 + (v20 + "")))
                                        _ = fi22
                                        _ = v25
                                    }
                                }
                                v26 := []int{42, 142}
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                _ = v26
                            }
                            _ = v20
                        }
                    }
                }
            case "foo bar":
                var v27 uint = 166
                _ = v27
            default:
                fmt.Println(v13, v2)
            }
        }
    }
    fmt.Println(v2, v13, v19)
    fmt.Println(v19, v2, v12)
    v19 = (v19 << ((65535) & 63))
    v14 -= v14
    v13 += ((float64(float32(0.0)) / (0.0 / v13)) - ((3.14159 + v13) - (v13 * (((0.0 / -2.25) / 0.1) + float64(0.0)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
