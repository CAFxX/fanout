package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32, p1 int) uint32 {
    var v7 int = p1
    p1 -= v7
    switch "foo bar" {
    case "a":
        v7 -= (v7 & (v7 + p1))
    }
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint64) int16 {
    var v9 uint16 = 336
    switch 9223372036854775807 {
    case -1:
        for fi10 := 0; fi10 < 6; fi10++ {
            if false {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: 94}
                    var v13 uint = 65535
                    fmt.Println(fi10)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if false {
                    p0 = p0
                    p0 = p0
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS1{F0: 0}
                        if (2147483647 >= 65535) {
                            if (266 >= 42) {
                                v16 := []int{569, (^(fi14)), (fi14 << ((18446744073709551615) & 63)), 32768}
                                p0 += 42
                                _ = v16
                            } else {
                                fmt.Println(p0, fi14, fi14)
                                v17 := GenS1{F0: 42}
                                if true {
                                    var v18 uint16 = (v9 * 2)
                                    v19 := GenS2{F0: true, F1: 2, F2: 65535}
                                    v9 += v9
                                    _ = v18
                                    _ = v19
                                }
                                _ = v17
                            }
                            p0 = (p0 | (p0 * (p0 / (((p0 / ((p0) | 1))) | 1))))
                            v20 := GenS1{F0: 42}
                            switch (((((("foo bar" + ("a" + ("hello" + ""))) + "foo bar") + "a") + "\n\t") + (("" + "a") + (("" + "\n\t") + "foo bar"))) + ("foo bar" + ("" + "\n\t"))) {
                            case "hello":
                                var v21 uint32 = 4294967295
                                _ = v21
                            case "":
                                v22 := GenS1{F0: 100}
                                _ = v22
                            }
                            _ = v20
                        } else {
                            switch "" {
                            case "foo bar":
                                var v23 int16 = 32767
                                _ = v23
                            }
                            v24 := []int{42}
                            var v25 uint8 = 0
                            v24 = append(v24, (fi10 - 127))
                            v26 := GenS2{F0: (!true), F1: 987, F2: uint16(v25)}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := []int{-2147483648, fi14, (fi10 * fi14)}
                        var v28 string = ("a" + "hello")
                        _ = fi14
                        _ = v15
                        _ = v27
                        _ = v28
                    }
                } else {
                    var v29 uint8 = 255
                    _ = v29
                }
            }
            _ = fi10
        }
    case 127:
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := []int{fi30, (fi30 ^ (-(fi30)))}
            p0 = p0
            _ = fi30
            _ = v31
        }
    default:
        switch 4294967295 {
        case 32768:
            fmt.Println(p0, v9)
        case 0:
            if false {
                v32 := []int{9223372036854775807}
                v33 := GenS2{F0: ("foo bar" != ("a" + "\n\t")), F1: v9, F2: 255}
                var v34 uint32 = 312
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v35 uint = 42
                _ = v35
            }
        }
    }
    p0 = 443
    var v36 uint8 = 167
    _ = p0
    _ = v36
    _ = v9
    return -955
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 127
    v3 := []uint32{42, 65535, 42, 2147483648}
    v4 := make(map[uint32]bool)
    var v5 int = 4294967295
    v6 := make(map[int]uint32)
    v37 := GenS3{F0: 2147483647}
    v38 := GenS1{F0: 42}
    var v39 float64 = 3.14159
    for fi40 := 0; fi40 < 16; fi40++ {
        var v41 int64 = 930
        _ = fi40
        _ = v41
    }
    switch ("foo bar" + "\n\t") {
    case "":
        var v42 int32 = -1
        _ = v42
    default:
        if false {
            var v43 uint = 0
            v44 := GenS3{F0: 0}
            _ = v43
            _ = v44
        }
    }
    v2 = 2
    fmt.Println(v2, v5, v5)
    switch (("\n\t" + "") + "\n\t") {
    case "foo bar":
        fmt.Println(v5)
    default:
        v39 += float64((3.14159 * (1.5 - 3.14159)))
    }
    var v45 float64 = (v39 - v39)
    v45 = (float64((float32(0.0) * (1.5 / 1.0))) + (v45 - v39))
    v39 = v39
    v1[4294967295] = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
