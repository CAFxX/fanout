package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 bool, p1 int64) bool {
    p1 += (-(p1))
    p1 = (p1 ^ p1)
    var v4 int16 = 1
    v5 := GenS1{F0: 0, F1: -2147483648, F2: "\n\t"}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 uint64, p1 int) string {
    if true {
        var v6 uint = 4294967295
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                if (false && (63 > -1)) {
                    v9 := []int{-1}
                    _ = v9
                }
                fmt.Println(v6)
                _ = fi8
            }
            var v10 string = "hello"
            _ = fi7
            _ = v10
        }
        v11 := GenS1{F0: 135, F1: 2147483647, F2: "a"}
        _ = v11
        _ = v6
    } else {
        if ((1e10 - 0.0) > 3.14159) {
            v12 := []int{int(p0), p1}
            _ = v12
        }
    }
    v13 := []int{(p1 * p1), (-(p1)), 32768}
    _ = p0
    _ = p1
    _ = v13
    return "a"
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 18446744073709551615
    v3 := []int{2147483647, 2147483647, 127, -2147483648}
    if false {
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                _ = len(v1)
                _ = fi15
            }
            _ = fi14
        }
        v16 := GenS1{F0: 2, F1: -1, F2: "\n\t"}
        _ = v16
    }
    var v17 int16 = 1
    var v18 uint8 = uint8(v2)
    v2 -= (v2 | (v2 - (^(v2))))
    var v19 int = 0
    var v20 int8 = 10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 += 0
    for fi21 := 0; fi21 < 8; fi21++ {
        if (9223372036854775807 >= -494) {
            if true {
                if true {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        fmt.Println(v18)
                        _ = fi22
                    }
                    v2 += v2
                    var v23 int64 = -991
                    _ = v23
                }
                fmt.Println(fi21, fi21)
                if len(v3) > 0 {
                    v3[0] = 915
                }
            } else {
                switch fi21 {
                case 0:
                    if ((-(v20)) >= (^(v20))) {
                        if (!true) {
                            var v24 float32 = ((3.14159 - 1e10) - 0.1)
                            var v25 int64 = -598
                            for fi26 := 0; fi26 < 7; fi26++ {
                                if false {
                                    var v27 float64 = -2.25
                                    var v28 uint32 = 42
                                    var v29 uint64 = v2
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int32 = 0
                                _ = fi26
                                _ = v30
                            }
                            _ = v24
                            _ = v25
                        }
                    } else {
                        var v31 uint64 = (^(v2))
                        _ = v31
                    }
                case 42:
                    fmt.Println(v20, fi21)
                case 396:
                    v20 += -128
                default:
                    var v32 float64 = 3.14159
                    _ = v32
                }
                v2 += (v2 << ((65535) & 63))
                for fi33 := 0; fi33 < 1; fi33++ {
                    v34 := GenS1{F0: (v18 << ((0) & 63)), F1: 184, F2: "foo bar"}
                    _ = fi33
                    _ = v34
                }
            }
        }
        var v35 int64 = 894
        _ = fi21
        _ = v35
    }
    v17 = 42
    for fi36 := 0; fi36 < 2; fi36++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi36
    }
    var v37 uint64 = v2
    v38 := GenS2{F0: (v2 ^ (v2 << ((1) & 63)))}
    _ = len(v3)
    var v39 float32 = (1.5 + -2.25)
    var v40 uint = 284
    v40 += (^(v40))
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
