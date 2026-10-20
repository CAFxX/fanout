package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 int16, p1 uint16, p2 uint32) int {
    fmt.Println(p1, p1, p0)
    v6 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -50
}
func genFn1() uint {
    var v7 int = 4294967295
    var v8 float32 = -2.25
    v8 -= (-2.25 * ((3.14159 / (-2.25 + (v8 * v8))) - v8))
    v9 := []int{v7, genFn0(32767, 1, 394), 674, v7}
    switch "hello" {
    case "\n\t":
        if (-958 != 284) {
            v8 = v8
            v9 = append(v9, genFn0(153, 744, 4294967295))
            var v10 uint16 = 637
            _ = v10
        }
    case "hello":
        if ((false && (65535 != 42)) && (-2147483648 != 1)) {
            fmt.Println(v7, v7)
        } else {
            if (!false) {
                var v11 int = -1
                var v12 uint = 65535
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                for fi14 := 0; fi14 < 8; fi14++ {
                    if (42 >= -32768) {
                        v8 = v8
                        if len(v9) > 0 {
                            v9[0] = fi14
                        }
                    } else {
                        v15 := GenS1{F0: 32}
                        _ = v15
                    }
                    fmt.Println(v7, v7)
                    v16 := GenS2{F0: 2, F1: ((("\n\t" + "hello") + "") + "a")}
                    _ = fi14
                    _ = v16
                }
                var v17 uint16 = 65535
                if true {
                    v18 := GenS2{F0: 0, F1: "\n\t"}
                    var v19 string = ""
                    v20 := GenS2{F0: 446, F1: v19}
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi13
                _ = v17
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                v7 += v7
                fmt.Println(fi21, v7, v8)
                _ = fi21
            }
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 string = (("hello" + "") + ("\n\t" + "hello"))
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v4 := []string{v1, v1, ("\n\t" + v1)}
    var v5 int32 = 2147483647
    if true {
        if true {
            var v22 uint = 42
            if (!true) {
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 uint32 = 340
                    v25 := GenS1{F0: 255}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if true {
                    if false {
                        v26 := GenS3{F0: ((v1 + "hello") + (("hello" + v1) + (v1 + "")))}
                        _ = v26
                    } else {
                        v22 -= (^(v22))
                        for fi27 := 0; fi27 < 16; fi27++ {
                            switch (fi27 | fi27) {
                            case 617:
                                v22 -= (v22 * v22)
                            case 9223372036854775807:
                                v28 := GenS2{F0: 1, F1: "\n\t"}
                                _ = v28
                            case 32767:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    var v30 int16 = int16(v5)
                                    _ = fi29
                                    _ = v30
                                }
                            }
                            var v31 int32 = int32(v22)
                            _ = fi27
                            _ = v31
                        }
                    }
                }
            } else {
                v32 := v3[65535]
                v33 := GenS2{F0: 1, F1: (v1 + ((v1 + ((v1 + v1) + (v1 + ((v1 + "foo bar") + (v1 + ("a" + v1)))))) + "a"))}
                _ = v32
                _ = v33
            }
            _ = v22
        }
        v5 -= v5
        v34 := v3[2147483647]
        _ = v34
    } else {
        fmt.Println(v5, v1)
        if len(v2) > 0 {
            v2[0] = 1
        }
    }
    v35 := v3[42]
    var v36 int = genFn0(-1, 862, 1)
    fmt.Println(v1)
    v37 := GenS1{F0: 493}
    var v38 int = (v36 * (v36 / (((v36 >> ((244) & 63))) | 1)))
    v39 := GenS4{F0: 9223372036854775807, F1: 1}
    v3[1] = ("" <= (("hello" + "a") + v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
