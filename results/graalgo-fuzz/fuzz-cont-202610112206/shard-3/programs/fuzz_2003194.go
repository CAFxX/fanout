package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8, p1 uint, p2 string) string {
    fmt.Println(p2, p0, p1)
    p0 -= p0
    v5 := []int{-178}
    if false {
        var v6 uint32 = 1
        v5 = append(v5, -1)
        var v7 float32 = 1.5
        _ = v6
        _ = v7
    } else {
        fmt.Println(p2, p2)
    }
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("foo bar" + (p2 + (p2 + p2)))
}
func genFn1(p0 uint64, p1 uint, p2 uint16) string {
    p2 -= (^(p2))
    v8 := GenS1{F0: 9223372036854775807, F1: "foo bar", F2: p1}
    if true {
        var v9 uint16 = p2
        if true {
            v10 := GenS1{F0: 128, F1: "foo bar", F2: p1}
            var v11 uint64 = 905
            _ = v10
            _ = v11
        }
        _ = v9
    } else {
        v12 := []int{-28, -752, 32767}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "a"
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    var v4 uint64 = 1
    fmt.Println(v4, v4)
    fmt.Println(v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        switch fi13 {
        case 42:
            v14 := GenS2{F0: -1, F1: 610, F2: "foo bar"}
            _ = v14
        case 32768:
            if true {
                switch (("\n\t" + "hello") + "a") {
                case "\n\t":
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v4 -= (v4 | (v4 + 18446744073709551615))
                        _ = fi15
                    }
                case "":
                    var v16 int8 = 69
                    _ = v16
                default:
                    if false {
                        v17 := []int{}
                        _ = v17
                    }
                }
                if false {
                    v1 = 1
                    v18 := GenS1{F0: -1, F1: "\n\t", F2: (v1 << (((v1 % ((v1) | 1))) & 63))}
                    v19 := GenS2{F0: 0, F1: 0, F2: ""}
                    _ = v18
                    _ = v19
                } else {
                    var v20 int32 = 2147483647
                    var v21 int32 = v20
                    switch (fi13 * -9223372036854775808) {
                    case -209:
                        var v22 uint32 = 42
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                }
                v4 -= v4
            } else {
                if (1 < 42) {
                    var v23 uint8 = 58
                    _ = v23
                } else {
                    if ((((!true) && ((fi13 | fi13) < fi13)) || (!(((3.14159 - 3.14159) - (0.1 * -2.25)) <= (0.1 * 1e10)))) && false) {
                        v24 := []int{fi13, fi13, (fi13 ^ fi13)}
                        for fi25 := 0; fi25 < 14; fi25++ {
                            var v26 uint = (v1 ^ 1)
                            switch ("hello" + "a") {
                            case "a":
                                _ = len(v3)
                            }
                            _ = len(v24)
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    } else {
                        v4 = 264
                    }
                    v4 -= v4
                    fmt.Println(v4, v1)
                }
                if (true || (!(!(((float64(0.1) / float64(1.5)) + 1e10) != 1.5)))) {
                    fmt.Println(v4, v4)
                    var v27 int16 = 980
                    var v28 bool = (!true)
                    _ = v27
                    _ = v28
                }
                v1 -= 489
            }
        default:
            v29 := []int{588, fi13, fi13, fi13}
            _ = v29
        }
        fmt.Println(fi13, fi13)
        v30 := GenS2{F0: 0, F1: 2, F2: ((("hello" + "") + "a") + ("\n\t" + "\n\t"))}
        _ = fi13
        _ = v30
    }
    var v31 bool = (42 != 255)
    v31 = false
    v31 = (!(886 <= 459))
    v32 := GenS1{F0: -2147483648, F1: ("foo bar" + ""), F2: v1}
    var v33 int16 = -32768
    v33 -= v33
    if true {
        if v31 {
            fmt.Println(v31, v31, v1)
        } else {
            var v34 int8 = 47
            v35 := GenS3{F0: 65535, F1: 35, F2: 0}
            _ = v34
            _ = v35
        }
        if (!v31) {
            v36 := []int{4294967295, 4294967295, 4294967295, 1}
            if (!(!v31)) {
                for fi37 := 0; fi37 < 19; fi37++ {
                    v1 = (v1 % ((v1) | 1))
                    _ = fi37
                }
            }
            _ = v36
        } else {
            switch 127 {
            case -267:
                v38 := []int{-2147483648}
                _ = v38
            }
            v31 = (false && (v31 || false))
        }
        var v39 float64 = (0.0 / 1.0)
        _ = v39
    }
    for fi40 := 0; fi40 < 10; fi40++ {
        fmt.Println(v4)
        _ = fi40
    }
    v41 := GenS1{F0: -625, F1: ((("hello" + "") + "") + "a"), F2: 870}
    var v42 float32 = 3.14159
    v43 := GenS4{F0: v4, F1: v4}
    v42 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
