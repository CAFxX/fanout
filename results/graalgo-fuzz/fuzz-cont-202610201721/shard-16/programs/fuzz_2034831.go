package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint }
type GenS6 struct { F0 int; F1 int32; F2 string }
type GenS7 struct { F0 bool; F1 uint }
func genFn0(p0 int8) uint {
    v6 := []int{127, 2147483647, -9223372036854775808, 0}
    v7 := GenS1{F0: "foo bar", F1: "hello"}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return 611
}
func genFn1(p0 uint32, p1 int64) int64 {
    p1 += p1
    var v8 int8 = -128
    p1 = 0
    v9 := GenS1{F0: "\n\t", F1: ("a" + "\n\t")}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = -1
    var v3 uint32 = 466
    var v4 uint32 = (v3 + v3)
    var v5 int = 687
    v2 = v2
    fmt.Println(v5, v3, v3)
    if (!(42 != 458)) {
        v10 := GenS2{F0: v5}
        for fi11 := 0; fi11 < 9; fi11++ {
            if false {
                v3 = (^(v3))
                v1 = (v1 % (((v1 - v1)) | 1))
                fmt.Println(v2, v3, v1)
            }
            var v12 int = v5
            v5 += (v5 % ((9223372036854775807) | 1))
            _ = fi11
            _ = v12
        }
        var v13 int = (-(v5))
        _ = v10
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 uint8 = 195
                v17 := []int{fi14, fi15, (-(fi14)), fi14}
                v18 := GenS1{F0: "foo bar", F1: "foo bar"}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
    }
    fmt.Println(v1, v3, v4)
    for fi19 := 0; fi19 < 8; fi19++ {
        switch (^(v5)) {
        case 923:
            v20 := []int{v5}
            _ = v20
        case 9223372036854775807:
            if (!false) {
                for fi21 := 0; fi21 < 15; fi21++ {
                    if (v2 >= 1) {
                        v1 = (^(v1))
                    } else {
                        var v22 string = "\n\t"
                        switch (v22 + v22) {
                        case "a":
                            var v23 int16 = -1
                            _ = v23
                        case "":
                            if (!true) {
                                v24 := GenS1{F0: (v22 + v22), F1: v22}
                                var v25 int8 = -1
                                v26 := []int{}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS2{F0: 199}
                                _ = v27
                            }
                        }
                        var v28 uint64 = (^(v1))
                        switch v22 {
                        case "a":
                            v1 += 2
                        case "foo bar":
                            var v29 uint32 = v3
                            v30 := []int{}
                            _ = len(v30)
                            _ = v29
                            _ = v30
                        }
                        _ = v22
                        _ = v28
                    }
                    _ = fi21
                }
            }
        default:
            switch 0 {
            case 127:
                v31 := GenS3{F0: 412}
                _ = v31
            case 9223372036854775807:
                switch "foo bar" {
                case "hello":
                    var v32 bool = false
                    _ = v32
                default:
                    fmt.Println(v5, v1, v1)
                }
            case 426:
                fmt.Println(v2, v2)
            }
        }
        var v33 uint32 = v3
        for fi34 := 0; fi34 < 16; fi34++ {
            v2 += (v2 * 42)
            v35 := []int{(^(v5))}
            var v36 uint8 = 42
            if (1 == 695) {
                v37 := GenS4{F0: 1}
                if true {
                    var v38 uint32 = (v33 ^ v33)
                    v1 = v1
                    _ = v38
                }
                if ((!true) && false) {
                    v39 := GenS5{F0: 2147483647, F1: uint64(v4), F2: 0}
                    _ = v39
                } else {
                    var v40 float64 = float64((-2.25 * (float32(1e10) / -2.25)))
                    _ = v40
                }
                _ = v37
            } else {
                fmt.Println(v4, fi19)
            }
            _ = len(v35)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi19
        _ = v33
    }
    fmt.Println(v1, v3, v2)
    var v41 uint64 = (v1 ^ v1)
    v41 = 42
    v42 := GenS6{F0: v5, F1: v2, F2: "\n\t"}
    v4 += v3
    var v43 uint8 = 1
    v3 += v3
    v44 := GenS7{F0: true, F1: uint(v41)}
    v43 = (v43 + (v43 * 127))
    var v45 uint = 531
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
