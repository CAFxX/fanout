package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int8) uint {
    if true {
        v4 := []int{1}
        var v5 int64 = -476
        p0 += p0
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if ((4294967295 == 65535) && (float64(3.14159) > 3.14159)) {
            var v7 uint = 42
            switch (fi6 * fi6) {
            case 816:
                var v8 int32 = 397
                _ = v8
            case 9223372036854775807:
                fmt.Println(p0, v7, fi6)
            default:
                v9 := GenS1{F0: 1, F1: 462, F2: 0}
                _ = v9
            }
            _ = v7
        } else {
            v10 := GenS1{F0: 2, F1: -9223372036854775808, F2: p0}
            _ = v10
        }
        v11 := []int{fi6, fi6, 4294967295, fi6}
        _ = fi6
        _ = v11
    }
    _ = p0
    return 2
}
func genFn1(p0 uint) uint8 {
    var v12 uint = (p0 ^ p0)
    fmt.Println(p0)
    for fi13 := 0; fi13 < 7; fi13++ {
        p0 -= 65535
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v12, p0, fi14)
            _ = fi14
        }
        _ = fi13
    }
    var v15 bool = (v12 == 268)
    var v16 uint = 684
    _ = p0
    _ = v12
    _ = v15
    _ = v16
    return 24
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (3.14159 - float64((float32(-2.25) / (((float32((1e10 + 0.0)) + 0.0) / float32(3.14159)) - 0.1))))
    v3 := []string{(("a" + "") + "a")}
    v2 = (v2 * 3.14159)
    fmt.Println(v2)
    var v17 uint = 4294967295
    switch -2147483648 {
    case 32767:
        v18 := GenS2{F0: 800, F1: ""}
        _ = v18
    case -575:
        v17 = (^(v17))
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(fi19)
            if (5 == (^(fi20))) {
                fmt.Println(v2)
                if (255 >= 32767) {
                    v2 = v2
                    v21 := GenS3{F0: (v17 | (v17 - v17)), F1: 76}
                    _ = v21
                }
                switch "\n\t" {
                case "\n\t":
                    v22 := GenS1{F0: 255, F1: -1, F2: 83}
                    _ = v22
                }
            } else {
                if true {
                    if false {
                        fmt.Println(fi20, fi20, fi20)
                        v2 += v2
                        if len(v3) > 0 {
                            v3[0] = (("a" + "foo bar") + "a")
                        }
                    }
                    var v23 float64 = (v2 + (v2 + (float64(float32(v2)) + v2)))
                    fmt.Println(fi19, v17, fi20)
                    _ = v23
                }
                if (fi20 <= -1) {
                    v24 := GenS4{F0: 42, F1: 42}
                    _ = v24
                } else {
                    if false {
                        v25 := GenS5{F0: -614, F1: false, F2: "\n\t"}
                        _ = v25
                    } else {
                        if false {
                            v26 := GenS1{F0: 1, F1: 0, F2: -128}
                            _ = v26
                        }
                    }
                }
                var v27 int = fi20
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v27
            }
            _ = fi20
        }
        v17 += 567
        _ = fi19
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v28 := GenS5{F0: 0, F1: true, F2: ""}
    v17 = (^(v17))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
}
