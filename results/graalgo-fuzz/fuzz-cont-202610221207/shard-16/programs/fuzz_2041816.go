package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 int16, p1 uint16, p2 int32) int {
    if true {
        var v4 float32 = 1.5
        if true {
            p1 = p1
            v5 := GenS1{F0: (2 < 104), F1: (p1 * (^(p1)))}
            _ = v5
        }
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "foo bar":
            p2 += p2
        case "hello":
            v6 := GenS1{F0: false, F1: (p1 << ((1) & 63))}
            _ = v6
        }
        _ = v4
    }
    p0 = (p0 & 1000)
    v7 := []int{-2147483648}
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    var v8 uint = 941
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int, p1 bool) uint16 {
    var v9 string = ("a" + "")
    v9 = v9
    var v10 int32 = 575
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 126
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{0, 2147483647}
    var v3 int8 = -113
    fmt.Println(v3)
    if ((65535 > 65535) || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "foo bar" {
        case "foo bar":
            v3 += (v3 % ((v3) | 1))
        case "hello":
            var v11 float64 = 0.0
            _ = v11
        default:
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 float32 = 1e10
                _ = fi12
                _ = v13
            }
        }
    }
    if (v3 > (v3 % ((v3) | 1))) {
        v3 += v3
        v3 -= (v3 << ((626) & 63))
        if (32767 != -32768) {
            v14 := GenS1{F0: true, F1: 2}
            for fi15 := 0; fi15 < 9; fi15++ {
                v3 = -128
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int = 568
                    _ = fi16
                    _ = v17
                }
                switch (("hello" + "a") + "hello") {
                case "a":
                    var v18 uint = 139
                    _ = v18
                case "\n\t":
                    v3 += 42
                default:
                    v1[(fi15 - (^(fi15)))] = "hello"
                }
                _ = fi15
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                if true {
                    var v20 int = genFn0(346, 65535, 261)
                    switch fi19 {
                    case 514:
                        fmt.Println(fi19)
                    case 32768:
                        if ((4294967295 > 0) && true) {
                            fmt.Println(v20)
                            v20 += fi19
                        }
                    case 247:
                        v1[(-(v20))] = ((("" + "a") + ("hello" + "")) + "")
                    }
                    v20 = v20
                    _ = v20
                } else {
                    v21 := GenS1{F0: true, F1: genFn1(0, ((0.1 / 3.14159) == ((1.5 - (-2.25 - (float64(((-2.25 * ((0.0 - 0.0) + 0.0)) * 0.1)) + 1e10))) + 1e10)))}
                    if false {
                        delete(v1, fi19)
                    } else {
                        v3 = 53
                        fmt.Println(fi19, v3)
                    }
                    _ = v21
                }
                _ = fi19
            }
            _ = v14
        } else {
            v3 -= v3
            for fi22 := 0; fi22 < 10; fi22++ {
                if false {
                    _ = len(v1)
                    var v23 uint8 = 42
                    _ = v23
                }
                fmt.Println(fi22)
                v3 = (v3 >> ((4294967295) & 63))
                _ = fi22
            }
        }
    }
    switch ("foo bar" + ((("a" + ("hello" + "hello")) + (((("hello" + "foo bar") + "foo bar") + "hello") + "a")) + ("" + "hello"))) {
    case "":
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 uint64 = 2
                fmt.Println(v3, fi24)
                _ = fi24
                _ = v25
            }
        } else {
            v26 := GenS1{F0: false, F1: 65535}
            _ = v26
        }
    default:
        if true {
            v3 -= (^(v3))
            v3 = (v3 % ((v3) | 1))
        }
    }
    if (!true) {
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(v3)
            var v28 uint8 = 2
            _ = fi27
            _ = v28
        }
    }
    v3 += -76
    fmt.Println(v3)
    switch genFn0(1, 566, 632) {
    case 128:
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int = fi29
            _ = fi29
            _ = v30
        }
    }
    if true {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 float32 = 1e10
            _ = fi31
            _ = v32
        }
        var v33 uint32 = 65535
        if (91 <= 0) {
            var v34 int8 = 100
            _ = v34
        }
        _ = v33
    }
    var v35 int32 = -2147483648
    fmt.Println(v3, v35)
    v36 := GenS1{F0: ((int(v3) != 97) && true), F1: 665}
    var v37 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
