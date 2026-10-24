package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0() uint8 {
    var v5 uint = 0
    var v6 int32 = 2147483647
    v7 := []int{}
    v8 := GenS1{F0: (true || false), F1: 255}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 uint8, p2 int8) int16 {
    switch -524 {
    case -2147483648:
        if (!false) {
            switch -585 {
            case 128:
                p2 = int8(p0)
            case -1:
                v9 := []int{}
                _ = v9
            }
            v10 := []int{127}
            _ = v10
        }
    case 649:
        var v11 string = ("hello" + ("\n\t" + ""))
        _ = v11
    case 4294967295:
        var v12 uint32 = 715
        _ = v12
    default:
        if (1000 < -386) {
            var v13 uint8 = genFn0()
            _ = v13
        } else {
            v14 := []int{32767}
            _ = v14
        }
    }
    if (!false) {
        for fi15 := 0; fi15 < 19; fi15++ {
            p0 += p0
            var v16 float32 = 1.5
            p1 += genFn0()
            _ = fi15
            _ = v16
        }
        var v17 float32 = 1.5
        for fi18 := 0; fi18 < 18; fi18++ {
            p0 = p0
            if false {
                fmt.Println(v17, fi18, fi18)
                fmt.Println(fi18, p0, p1)
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 string = "hello"
                    p0 -= (p0 & (p0 % ((p0) | 1)))
                    v20 = "hello"
                    _ = fi19
                    _ = v20
                }
            } else {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := []int{fi18, fi18, fi21, fi21}
                    switch fi18 {
                    case 16:
                        p0 = (p0 << ((p0) & 63))
                    case 4294967295:
                        if ((p1 > (p1 ^ 0)) && true) {
                            fmt.Println(v17, p2, fi18)
                        } else {
                            p1 -= (p1 - p1)
                            if true {
                                switch fi21 {
                                case -9223372036854775808:
                                    v22 = append(v22, (fi21 - (fi18 & fi21)))
                                case 388:
                                    v23 := GenS1{F0: ((p2 & -1) == p2), F1: 782}
                                    _ = v23
                                }
                                var v24 int = 0
                                _ = len(v22)
                                _ = v24
                            }
                        }
                    case 337:
                        fmt.Println(p0)
                    }
                    _ = fi21
                    _ = v22
                }
                switch ("hello" + "\n\t") {
                case "foo bar":
                    var v25 float64 = 0.1
                    _ = v25
                default:
                    var v26 float32 = (v17 * v17)
                    _ = v26
                }
                var v27 int16 = 1000
                _ = v27
            }
            _ = fi18
        }
        _ = v17
    }
    v28 := []int{42, 372}
    _ = len(v28)
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return -1
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 42
    v3 := make(map[int]int)
    var v4 int16 = 1
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := []int{}
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 279
    var v32 int32 = 382
    if ((v32 * 413) > (v32 >> ((767) & 63))) {
        v3[162] = -2147483648
        var v33 int32 = 1
        v31 = 1
        _ = v33
    }
    switch "" {
    case "":
        if true {
            for fi34 := 0; fi34 < 12; fi34++ {
                fmt.Println(fi34, v2, v4)
                var v35 int16 = (v4 & 42)
                fmt.Println(v32, v1)
                _ = fi34
                _ = v35
            }
            v36 := []int{}
            for fi37 := 0; fi37 < 20; fi37++ {
                for fi38 := 0; fi38 < 4; fi38++ {
                    v39 := GenS1{F0: false, F1: uint16(v31)}
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
            _ = v36
        }
    case "\n\t":
        v32 = (v32 - v32)
    }
    switch "\n\t" {
    case "hello":
        v40 := GenS1{F0: (65535 != 1), F1: 0}
        _ = v40
    default:
        fmt.Println(v31)
    }
    v41 := []int{9223372036854775807, -1, -587, 0}
    var v42 int32 = v32
    v32 = (v32 >> ((18446744073709551615) & 63))
    v32 = v32
    v43 := GenS1{F0: true, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
