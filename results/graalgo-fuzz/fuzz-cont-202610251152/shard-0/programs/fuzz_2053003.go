package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 int, p1 uint8, p2 int8) uint8 {
    if (-569 <= p0) {
        p2 = (p2 | (p2 + p2))
    }
    v5 := GenS1{F0: 2147483647, F1: false}
    if (!(1 <= 0)) {
        p1 += 1
        var v6 uint8 = p1
        _ = v6
    } else {
        var v7 float64 = 3.14159
        p0 = (p0 | (p0 >> ((164) & 63)))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 36
    v3 := []uint32{254}
    v4 := make(map[int]bool)
    v2 -= (v2 % (((v2 | (v2 | v2))) | 1))
    if (!(42 > 1)) {
        var v8 float64 = ((-2.25 * 0.0) + 1.5)
        if (3.14159 > -2.25) {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int64 = -9223372036854775808
                _ = fi9
                _ = v10
            }
            if true {
                var v11 int = -72
                _ = v11
            } else {
                var v12 int8 = int8(v2)
                for fi13 := 0; fi13 < 13; fi13++ {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v15 := GenS1{F0: 2147483647, F1: (0 < -2147483648)}
                        if true {
                            v8 += ((v8 + float64(1e10)) * 1e10)
                            fmt.Println(v2, fi13, fi14)
                            v8 += v8
                        } else {
                            var v16 bool = (false || false)
                            var v17 uint64 = 0
                            v16 = v16
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                    delete(v4, fi13)
                    fmt.Println(v12, v12)
                    _ = fi13
                }
                _ = v12
            }
            v8 = v8
        } else {
            fmt.Println(v8)
        }
        v18 := GenS1{F0: 1, F1: (true && true)}
        _ = v18
        _ = v8
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint8 = genFn0((fi19 % (((fi19 % ((42) | 1))) | 1)), v2, -128)
            if (uint(v2) < 994) {
                for fi21 := 0; fi21 < 1; fi21++ {
                    if true {
                        var v22 int = 9223372036854775807
                        var v23 uint = 255
                        v20 += (v2 + (v20 ^ v20))
                        _ = v22
                        _ = v23
                    }
                    if true {
                        _ = len(v4)
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = len(v4)
                    }
                    _ = fi21
                }
                if false {
                    v2 = genFn0(fi19, 1, -31)
                    switch (("\n\t" + "") + (("" + (("" + ("hello" + ("\n\t" + ""))) + "a")) + "\n\t")) {
                    case "foo bar":
                        v20 = (v20 ^ genFn0(fi19, (^(v20)), 127))
                    }
                    _ = len(v3)
                } else {
                    v24 := GenS2{F0: 803}
                    v20 = (v20 / (((v20 & 2)) | 1))
                    _ = v24
                }
                if (!(false || false)) {
                    fmt.Println(v20, fi19, v20)
                    v20 -= 79
                } else {
                    v25 := GenS2{F0: 4294967295}
                    v26 := v1[2147483647]
                    _ = v25
                    _ = v26
                }
            }
            _ = fi19
            _ = v20
        }
        var v27 int16 = 209
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS3{F0: (v27 + 1000)}
            if (127 == -1) {
                v30 := GenS4{F0: 844, F1: (-(v27)), F2: true}
                _ = v30
            }
            var v31 float64 = 0.0
            _ = fi28
            _ = v29
            _ = v31
        }
        _ = v27
    }
    fmt.Println(v2)
    v2 -= genFn0(2147483647, (v2 % (((^(v2))) | 1)), 119)
    v32 := GenS3{F0: 42}
    delete(v1, 2)
    switch -1 {
    case 42:
        var v33 uint16 = 2
        _ = v33
    }
    v34 := GenS5{F0: -1}
    switch "" {
    case "foo bar":
        v2 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
}
