package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string; F1 int64; F2 int }
func genFn0() uint16 {
    var v5 uint32 = 203
    v6 := []int{320, 2147483647, 32767, 32768}
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 uint8) uint {
    v7 := GenS1{F0: 4294967295}
    if false {
        var v8 int32 = p0
        v9 := GenS2{F0: -32768, F1: (42 > p0)}
        _ = v8
        _ = v9
    } else {
        if ((((0.1 / 1.5) / 1.5) <= (float64(float32(1.5)) + float64(float32(1.5)))) || false) {
            var v10 string = ("foo bar" + "")
            _ = v10
        } else {
            if true {
                var v11 string = (((("\n\t" + (((("\n\t" + ((("foo bar" + ("" + "a")) + ("hello" + "")) + "a")) + "a") + "hello") + (("hello" + "a") + "a"))) + "foo bar") + ((("foo bar" + "hello") + ("\n\t" + "foo bar")) + "foo bar")) + "hello")
                v12 := []int{619}
                v11 = v11
                _ = v11
                _ = v12
            } else {
                switch "\n\t" {
                case "\n\t":
                    var v13 uint16 = genFn0()
                    _ = v13
                case "":
                    if false {
                        v14 := []int{679}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        if true {
                            switch 2147483647 {
                            case 9223372036854775807:
                                p0 -= p0
                            default:
                                p1 += p1
                            }
                            switch "hello" {
                            case "foo bar":
                                var v15 uint16 = 65535
                                _ = v15
                            case "a":
                                fmt.Println(p0)
                            case "hello":
                                if (!(!(true && false))) {
                                    if (true && true) {
                                        var v16 uint32 = 4294967295
                                        _ = v16
                                    }
                                    p0 = 0
                                } else {
                                    for fi17 := 0; fi17 < 5; fi17++ {
                                        v18 := GenS3{F0: false, F1: 65535}
                                        var v19 int32 = p0
                                        var v20 uint16 = genFn0()
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    v21 := GenS4{F0: "hello", F1: -1, F2: 32768}
                                    _ = v21
                                }
                            }
                            p1 = 255
                        } else {
                            v22 := GenS1{F0: 42}
                            _ = v22
                        }
                        _ = v14
                    } else {
                        if true {
                            var v23 bool = ((-9223372036854775808 != 0) || (1 < (p0 << ((42) & 63))))
                            _ = v23
                        }
                        p0 -= p0
                        v24 := GenS4{F0: ("a" + ((("\n\t" + "foo bar") + "a") + "")), F1: 0, F2: 1}
                        _ = v24
                    }
                case "a":
                    fmt.Println(p0, p1)
                }
                p1 = (p1 % ((1) | 1))
            }
            var v25 uint16 = 1
            _ = v25
        }
        v26 := GenS2{F0: 0, F1: true}
        _ = v26
    }
    var v27 int = -1
    _ = p0
    _ = p1
    _ = v27
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []int{128, -524, -2, 42}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 float32 = (((((((1.5 + (3.14159 + 0.0)) + 0.1) + -2.25) * (1e10 + (0.0 - 1e10))) + 0.0) / 0.1) / ((1e10 / 1e10) * (((float32(0.1) / 1.5) / -2.25) - -2.25)))
    v4 = (v4 / v4)
    var v28 int = 581
    v29 := GenS1{F0: 2147483647}
    var v30 uint16 = genFn0()
    var v31 float64 = -2.25
    var v32 uint16 = v30
    for fi33 := 0; fi33 < 20; fi33++ {
        fmt.Println(v30)
        var v34 float32 = (v4 / (-2.25 * v4))
        switch -1 {
        case 32767:
            var v35 uint64 = 400
            _ = v35
        case 32768:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi33
        _ = v34
    }
    var v36 uint64 = 845
    fmt.Println(v28, v31)
    v30 = genFn0()
    for fi37 := 0; fi37 < 13; fi37++ {
        v36 -= v36
        _ = fi37
    }
    var v38 int8 = 0
    fmt.Println(v30, v32, v28)
    fmt.Println(v30, v4)
    v39 := GenS3{F0: (996 <= -212), F1: (v32 | (v32 % ((795) | 1)))}
    fmt.Println(v31, v30, v38)
    if (!true) {
        var v40 int32 = 0
        _ = v40
    }
    v32 += (v32 & v32)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v38
    _ = v39
    _ = v4
}
