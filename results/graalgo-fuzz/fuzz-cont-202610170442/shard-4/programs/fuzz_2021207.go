package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 string; F2 int }
type GenS7 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 int, p1 string, p2 string) int16 {
    p1 = (p2 + (("a" + "") + ("\n\t" + "")))
    v4 := []int{(p0 - (p0 & (^(p0)))), 1, p0, (p0 * (p0 + 2147483647))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1() uint64 {
    var v5 float32 = 0.0
    if (!true) {
        v6 := []int{545, 7, 863}
        fmt.Println(v5)
        for fi7 := 0; fi7 < 1; fi7++ {
            if true {
                var v8 int = -693
                _ = len(v6)
                _ = v8
            } else {
                if (!(false && (3.14159 <= (1e10 / ((1.5 + 0.1) - 3.14159))))) {
                    v5 = v5
                }
                if true {
                    if (v5 != v5) {
                        v9 := GenS1{F0: false, F1: fi7, F2: 90}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v9
                    } else {
                        v5 = (float32(0.1) / v5)
                        var v10 uint8 = 2
                        _ = v10
                    }
                }
                if true {
                    fmt.Println(fi7, v5)
                    v11 := GenS1{F0: (32767 != genFn0(32767, ("\n\t" + ""), "")), F1: fi7, F2: 0}
                    _ = v11
                } else {
                    var v12 uint64 = 18446744073709551615
                    var v13 int8 = -38
                    fmt.Println(fi7, v12)
                    _ = v12
                    _ = v13
                }
            }
            _ = fi7
        }
        _ = v6
    }
    switch -9223372036854775808 {
    case 128:
        v14 := GenS2{F0: -870}
        _ = v14
    default:
        v15 := []int{}
        _ = v15
    }
    _ = v5
    return 203
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 599
    var v3 uint64 = 0
    fmt.Println(v3, v2)
    fmt.Println(v3)
    v16 := GenS3{F0: 0}
    v17 := []int{9223372036854775807, -67, 127, 339}
    var v18 int = 610
    if (false && (!(-1 <= 0))) {
        fmt.Println(v18, v3, v3)
        switch (v18 + v18) {
        case -261:
            for fi19 := 0; fi19 < 3; fi19++ {
                switch "foo bar" {
                case "a":
                    var v20 uint8 = 2
                    _ = v20
                case "hello":
                    v21 := GenS4{F0: (3.14159 >= (0.1 + (0.1 / 3.14159)))}
                    _ = v21
                case "":
                    if true {
                        if (!true) {
                            var v22 uint16 = 32767
                            if (!false) {
                                var v23 int64 = 761
                                v3 = v3
                                _ = v23
                            }
                            _ = v22
                        }
                        fmt.Println(v18, fi19, fi19)
                    } else {
                        v24 := GenS4{F0: (!(!false))}
                        _ = v24
                    }
                default:
                    var v25 uint32 = v2
                    _ = v25
                }
                _ = fi19
            }
        case 2147483647:
            v17 = append(v17, -1)
        case 9223372036854775807:
            v26 := v1[(("foo bar" + "foo bar") + "foo bar")]
            _ = v26
        default:
            var v27 uint16 = 255
            _ = v27
        }
    } else {
        v3 += genFn1()
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                if (42 != 1) {
                    fmt.Println(fi28, v3, fi29)
                    v30 := v1[(("hello" + "a") + ("a" + "hello"))]
                    _ = v30
                }
                _ = fi29
            }
            v31 := GenS5{F0: (fi28 << ((uint(v3)) & 63))}
            if (true || ((true && true) && (-2147483648 > -2147483648))) {
                if (127 != 100) {
                    var v32 float32 = ((0.0 * 0.0) / 1e10)
                    if len(v17) > 0 {
                        v17[0] = fi28
                    }
                    _ = v32
                }
            }
            _ = fi28
            _ = v31
        }
    }
    v18 = -1
    v18 = 1
    v33 := GenS6{F0: 127, F1: "foo bar", F2: (v18 ^ v18)}
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := GenS6{F0: 128, F1: "a", F2: (v18 << ((42) & 63))}
        v36 := GenS7{F0: (^(v3)), F1: 255, F2: 656}
        v3 += v3
        _ = fi34
        _ = v35
        _ = v36
    }
    switch "hello" {
    case "foo bar":
        v18 -= v18
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        v2 = v2
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
}
