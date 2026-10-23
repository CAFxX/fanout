package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
func genFn0(p0 int16) uint16 {
    if (42 < 127) {
        for fi7 := 0; fi7 < 13; fi7++ {
            if (float32(((((-2.25 + (1e10 - 1.5)) * 3.14159) * (0.0 * 0.0)) / 1.5)) <= 3.14159) {
                for fi8 := 0; fi8 < 5; fi8++ {
                    p0 = p0
                    _ = fi8
                }
                switch "" {
                case "foo bar":
                    var v9 uint64 = 23
                    _ = v9
                case "a":
                    v10 := []int{588}
                    _ = v10
                case "\n\t":
                    if (4294967295 == 586) {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            var v12 uint64 = 305
                            var v13 float64 = (1.5 + 0.0)
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    } else {
                        p0 = p0
                    }
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int8 = 42
                    var v16 uint32 = 1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            if false {
                if false {
                    var v17 uint16 = 1
                    v18 := []int{-1, (fi7 % (((fi7 - -1)) | 1))}
                    fmt.Println(fi7, p0)
                    _ = v17
                    _ = v18
                }
                var v19 int16 = p0
                var v20 float32 = -2.25
                _ = v19
                _ = v20
            }
            _ = fi7
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(p0)
        for fi22 := 0; fi22 < 1; fi22++ {
            p0 = (p0 * p0)
            _ = fi22
        }
        _ = fi21
    }
    _ = p0
    return 0
}
func genFn1(p0 int32, p1 int64) int8 {
    fmt.Println(p1, p0)
    if (!(65535 != 0)) {
        p0 = (p0 | (p0 - (p0 & p0)))
        var v23 bool = ((p1 >> ((2) & 63)) <= p1)
        _ = v23
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    var v4 int64 = 42
    var v5 bool = (!true)
    var v6 uint64 = 14
    var v24 uint8 = v1
    if v5 {
        for fi25 := 0; fi25 < 2; fi25++ {
            if v5 {
                var v26 int64 = (v4 / (((-(v4))) | 1))
                var v27 uint8 = (^(v1))
                _ = v26
                _ = v27
            } else {
                v28 := GenS1{F0: 897, F1: "", F2: (v6 | (^(v6)))}
                _ = v28
            }
            _ = fi25
        }
        v29 := []int{-9223372036854775808}
        _ = v29
    }
    var v30 uint64 = 57
    if (v5 && false) {
        v30 += 0
    }
    v31 := []int{32768}
    if true {
        var v32 uint64 = (v30 & 336)
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
        _ = v32
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v24)
            v4 -= (^(v4))
            _ = fi33
        }
        fmt.Println(v24, v30, v24)
    }
    v6 -= (v30 >> ((24) & 63))
    if v5 {
        var v34 string = ""
        _ = v34
    } else {
        var v35 int64 = (v4 ^ 0)
        v35 = v4
        _ = v35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        if (!true) {
            _ = len(v3)
            _ = len(v3)
        } else {
            var v37 int8 = -24
            v3[1] = 4294967295
            var v38 uint32 = 2147483648
            _ = v37
            _ = v38
        }
        _ = len(v3)
        _ = fi36
    }
    var v39 int32 = -257
    fmt.Println(v5)
    var v40 uint = 0
    var v41 int32 = 159
    if len(v31) > 0 {
        v31[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
