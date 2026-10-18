package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint32) string {
    for fi7 := 0; fi7 < 1; fi7++ {
        if false {
            var v8 float64 = 1e10
            if (false && false) {
                var v9 string = "a"
                var v10 uint32 = 2147483647
                _ = v10
                _ = v9
            } else {
                p1 = (p0 & 1)
                for fi11 := 0; fi11 < 15; fi11++ {
                    if (42 <= 2) {
                        p0 = 2147483647
                        var v12 uint8 = 0
                        _ = v12
                    }
                    _ = fi11
                }
                if true {
                    var v13 int = 4294967295
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v13 += 4294967295
                        _ = fi14
                    }
                    _ = v13
                }
            }
            _ = v8
        } else {
            var v15 uint8 = 128
            _ = v15
        }
        if (!true) {
            v16 := []int{fi7, fi7}
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 float64 = -2.25
                fmt.Println(fi17)
                if false {
                    var v19 string = ("hello" + "")
                    if (!((fi17 >= fi7) || false)) {
                        switch v19 {
                        case "hello":
                            if true {
                                v18 += v18
                                v20 := GenS1{F0: v19, F1: 112, F2: 58}
                                _ = v20
                            } else {
                                var v21 uint = 255
                                _ = v21
                            }
                        case "":
                            if true {
                                v22 := GenS2{F0: -618, F1: 42, F2: 2147483647}
                                if len(v16) > 0 {
                                    v16[0] = -969
                                }
                                _ = v22
                            } else {
                                p1 -= (p1 ^ (^(p1)))
                                v23 := GenS1{F0: (v19 + v19), F1: 2, F2: 255}
                                _ = v23
                            }
                        case "\n\t":
                            v19 = v19
                        }
                        fmt.Println(p0, v18)
                        var v24 uint16 = 1
                        var v25 uint16 = v24
                        _ = v24
                        _ = v25
                    }
                    var v26 uint = 883
                    if false {
                        fmt.Println(v18, p0, fi17)
                    } else {
                        var v27 uint = v26
                        if ((136 != (p1 ^ p1)) || true) {
                            switch 521 {
                            case 32768:
                                if (42 >= 1) {
                                    var v28 int = fi17
                                    _ = len(v16)
                                    _ = v28
                                }
                            case 1:
                                v27 += v26
                            case -411:
                                var v29 bool = (!false)
                                _ = v29
                            }
                            fmt.Println(p1, fi7)
                            v18 = (v18 * float64((-2.25 * 0.1)))
                        }
                        _ = v27
                    }
                    _ = v19
                    _ = v26
                }
                _ = fi17
                _ = v18
            }
            switch fi7 {
            case -911:
                var v30 uint8 = 255
                _ = v30
            }
            _ = v16
        }
        _ = fi7
    }
    if false {
        for fi31 := 0; fi31 < 15; fi31++ {
            p1 = p1
            _ = fi31
        }
    }
    _ = p0
    _ = p1
    return "foo bar"
}
func genFn1(p0 int32, p1 bool) uint16 {
    var v32 int32 = (^(p0))
    fmt.Println(p0, v32)
    v33 := []int{-732, 4294967295, -9223372036854775808, 128}
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    return 65535
}

func main() {
    v1 := []uint32{}
    var v2 float32 = 0.0
    var v3 uint16 = 255
    var v4 int32 = 1
    var v5 int16 = 1000
    v6 := make(map[uint32]bool)
    v3 -= v3
    var v34 bool = ((^(v3)) < genFn1(v4, true))
    v2 += -2.25
    if ((v34 || v34) || v34) {
        v3 = genFn1((v4 - v4), (false || (!(v34 && ((v34 && false) || (!v34))))))
    }
    v5 = v5
    var v35 uint16 = v3
    for fi36 := 0; fi36 < 7; fi36++ {
        if (42 <= 1) {
            v5 += int16(fi36)
        }
        _ = fi36
    }
    v37 := GenS3{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
