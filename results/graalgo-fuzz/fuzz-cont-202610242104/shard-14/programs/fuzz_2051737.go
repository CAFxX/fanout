package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint, p1 uint8, p2 uint32) uint32 {
    fmt.Println(p0)
    v4 := GenS1{F0: 1, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1(p0 uint32) uint8 {
    v5 := GenS2{F0: 0}
    var v6 uint32 = (^(p0))
    _ = p0
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -9223372036854775808
    var v3 int16 = 1000
    v3 -= (v3 / ((v3) | 1))
    var v7 uint16 = 602
    v7 = 90
    switch "hello" {
    case "a":
        v8 := []int{0}
        _ = v8
    default:
        v7 += (v7 ^ v7)
    }
    if false {
        var v9 int64 = (v2 << ((0) & 63))
        var v10 float32 = (float32(1.5) / ((3.14159 - 1.5) + -2.25))
        _ = v10
        _ = v9
    } else {
        v3 = (v3 / ((int16(v7)) | 1))
        if true {
            v1[9223372036854775807] = false
            if (true || (v3 == (v3 << ((1) & 63)))) {
                fmt.Println(v3, v2)
                if true {
                    switch ("\n\t" + "foo bar") {
                    case "":
                        fmt.Println(v7)
                    case "hello":
                        if true {
                            v7 = v7
                        } else {
                            var v11 int32 = 840
                            _ = v11
                        }
                    }
                    v3 = -32768
                    var v12 int32 = 431
                    _ = v12
                } else {
                    var v13 uint64 = 131
                    delete(v1, 430)
                    _ = v13
                }
                v14 := GenS1{F0: v7, F1: (^(v7))}
                _ = v14
            } else {
                fmt.Println(v2)
                v7 += v7
                v2 = v2
            }
            fmt.Println(v7, v3, v3)
        } else {
            switch -1 {
            case 42:
                var v15 float64 = 0.1
                _ = v15
            case 128:
                if false {
                    var v16 float64 = 1e10
                    _ = v16
                }
            case 127:
                var v17 int8 = 1
                _ = v17
            }
            v2 = (^(v2))
            var v18 float64 = (-2.25 + 0.0)
            _ = v18
        }
        var v19 string = ("a" + ((("a" + "hello") + "hello") + "\n\t"))
        _ = v19
    }
    if ((((1.5 * ((((-2.25 * float64((0.1 * 0.1))) + 0.1) * 0.0) / 1e10)) * float64((1.5 / (((1.5 + 1.5) / 1.5) - 1.5)))) / -2.25) != 0.0) {
        var v20 uint8 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float32 = (1e10 - 1e10)
        for fi23 := 0; fi23 < 19; fi23++ {
            v3 = v3
            var v24 float64 = (((float64(v22) * -2.25) / ((3.14159 * 1.5) - (-2.25 + 1e10))) * 0.0)
            if (true || (!true)) {
                fmt.Println(v22, fi21, v7)
                v25 := []int{fi21, (fi21 << ((203) & 63)), int(v7)}
                _ = v25
            } else {
                v24 -= ((0.1 - v24) + (-2.25 - 1.5))
                if true {
                    v3 -= v3
                }
            }
            _ = fi23
            _ = v24
        }
        if (fi21 != 128) {
            var v26 int32 = 1
            if (!(v26 == (v26 >> ((42) & 63)))) {
                switch 32767 {
                case 32767:
                    switch ((("a" + "\n\t") + ("" + "hello")) + "\n\t") {
                    case "a":
                        if (fi21 == (fi21 >> ((0) & 63))) {
                            v3 -= v3
                            switch "\n\t" {
                            case "hello":
                                var v27 uint32 = 4294967295
                                _ = v27
                            case "foo bar":
                                if true {
                                    v28 := GenS1{F0: (v7 << ((0) & 63)), F1: v7}
                                    for fi29 := 0; fi29 < 12; fi29++ {
                                        v26 = v26
                                        _ = fi29
                                    }
                                    v30 := []int{(-(fi21)), fi21}
                                    _ = v28
                                    _ = v30
                                } else {
                                    v7 -= (v7 << ((18446744073709551615) & 63))
                                    for fi31 := 0; fi31 < 7; fi31++ {
                                        var v32 uint16 = (^(v7))
                                        _ = fi31
                                        _ = v32
                                    }
                                    var v33 float64 = 3.14159
                                    _ = v33
                                }
                            case "":
                                v34 := []int{fi21, fi21, 9223372036854775807, fi21}
                                _ = v34
                            default:
                                fmt.Println(fi21, v7, fi21)
                            }
                            var v35 bool = (-9223372036854775808 == fi21)
                            _ = v35
                        }
                    case "\n\t":
                        fmt.Println(v7, fi21)
                    case "hello":
                        var v36 uint = 1
                        _ = v36
                    }
                case -1:
                    fmt.Println(v3, v2, v3)
                case 9223372036854775807:
                    v37 := []int{}
                    _ = v37
                }
                for fi38 := 0; fi38 < 20; fi38++ {
                    if true {
                        var v39 uint64 = 199
                        _ = v39
                    } else {
                        var v40 int64 = v2
                        v41 := []int{fi21, 998, (fi38 ^ 864), fi38}
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                }
            }
            if false {
                fmt.Println(v7, v26, v7)
            } else {
                v42 := []int{fi21, fi21}
                _ = v42
            }
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    var v43 float64 = (0.1 / 1.0)
    var v44 uint16 = v7
    var v45 uint16 = v7
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v7
}
