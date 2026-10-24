package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 string, p1 int, p2 int32) uint {
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            if (!((false || ((0.1 + (0.0 - -2.25)) >= (0.1 * 1e10))) && false)) {
                fmt.Println(p0, p0, p2)
            } else {
                if (false && true) {
                    if len(v6) > 0 {
                        v6[0] = -9223372036854775808
                    }
                    p2 = (^(p2))
                    p1 += (^(fi7))
                } else {
                    if (42 > 0) {
                        p0 = p0
                        if len(v6) > 0 {
                            v6[0] = (fi7 & (fi7 << ((42) & 63)))
                        }
                    } else {
                        switch ((((p0 + "a") + p0) + "a") + ((((p0 + (p0 + p0)) + "a") + p0) + p0)) {
                        case "a":
                            for fi8 := 0; fi8 < 19; fi8++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi8
                            }
                        case "\n\t":
                            var v9 int64 = -9223372036854775808
                            _ = v9
                        default:
                            var v10 int64 = -9223372036854775808
                            _ = v10
                        }
                    }
                    fmt.Println(p1)
                }
            }
            _ = fi7
        }
        var v11 int16 = 32767
        _ = v11
    } else {
        var v12 int16 = 844
        _ = v12
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 int32, p1 int, p2 int32) int64 {
    var v13 uint8 = 43
    p2 -= -1
    v14 := []int{p1, (-(p1)), 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return -767
}

func main() {
    v1 := []uint32{2147483648, 985, 871, 2147483648}
    var v2 uint32 = 823
    var v3 int32 = 42
    v4 := make(map[int]int)
    v5 := []string{"foo bar", ("\n\t" + ("hello" + "\n\t")), ("" + "foo bar")}
    for fi15 := 0; fi15 < 15; fi15++ {
        v3 = 0
        _ = fi15
    }
    fmt.Println(v2)
    fmt.Println(v3, v3)
    v16 := GenS1{F0: v2}
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v3 -= v3
    delete(v4, 9223372036854775807)
    if false {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 bool = true
            _ = fi17
            _ = v18
        }
    } else {
        var v19 int64 = 0
        switch 9223372036854775807 {
        case 4294967295:
            var v20 bool = (1.5 > (float32(((0.0 * (1e10 + -2.25)) * 1.5)) / (float32(3.14159) + ((0.0 / 0.1) / ((0.1 * (3.14159 * -2.25)) - -2.25)))))
            _ = v20
        case 360:
            fmt.Println(v3)
        default:
            v19 = (v19 * v19)
        }
        for fi21 := 0; fi21 < 16; fi21++ {
            if false {
                if false {
                    v2 = (v2 | (v2 % (((v2 >> ((4294967295) & 63))) | 1)))
                    var v22 string = ("a" + ("a" + "\n\t"))
                    v1 = append(v1, 949)
                    _ = v22
                }
                switch "\n\t" {
                case "a":
                    var v23 uint8 = uint8(v2)
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v4[(fi24 % (((fi21 * int(v19))) | 1))] = 32767
                        v19 = (v19 % (((v19 % (((v19 / ((-9223372036854775808) | 1))) | 1))) | 1))
                        v25 := GenS2{F0: -1}
                        fmt.Println(fi21)
                        v19 = -1
                        v19 += (v19 & v19)
                        if len(v5) > 0 {
                            v5[0] = ("\n\t" + "hello")
                        }
                        _ = fi24
                        _ = v25
                    }
                case "":
                    if false {
                        if (!true) {
                            var v26 float64 = 0.0
                            v2 += (v2 & v2)
                            var v27 uint8 = 125
                            _ = v26
                            _ = v27
                        } else {
                            if false {
                                var v28 int64 = v19
                                var v29 uint = 65535
                                var v30 string = "hello"
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            v31 := GenS1{F0: (v2 + (v2 & (v2 + v2)))}
                            _ = v31
                        }
                    }
                default:
                    v32 := GenS2{F0: 714}
                    _ = v32
                }
            }
            var v33 int = (^(fi21))
            fmt.Println(v33, v19, v19)
            _ = fi21
            _ = v33
        }
        _ = v19
    }
    v2 -= v2
    for fi34 := 0; fi34 < 9; fi34++ {
        v2 = (v2 ^ (v2 << ((genFn0(("hello" + "hello"), fi34, v3)) & 63)))
        _ = fi34
    }
    fmt.Println(v3)
    var v35 uint64 = 0
    var v36 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
