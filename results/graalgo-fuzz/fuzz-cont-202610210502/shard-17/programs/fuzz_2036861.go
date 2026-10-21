package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 int64, p1 int64) uint64 {
    var v4 float64 = -2.25
    switch "a" {
    case "\n\t":
        fmt.Println(p0)
    case "hello":
        if false {
            p0 += 42
            if (129 > 1) {
                v5 := GenS1{F0: 1}
                if true {
                    switch 4294967295 {
                    case 42:
                        for fi6 := 0; fi6 < 16; fi6++ {
                            var v7 int32 = 1
                            var v8 uint64 = 18446744073709551615
                            v9 := GenS1{F0: 4294967295}
                            _ = fi6
                            _ = v7
                            _ = v8
                            _ = v9
                        }
                    case 127:
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v4 -= 0.1
                            _ = fi10
                        }
                    case 1:
                        v4 = (v4 + v4)
                    default:
                        for fi11 := 0; fi11 < 14; fi11++ {
                            fmt.Println(p0, fi11, p0)
                            var v12 string = ""
                            var v13 uint64 = 0
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    switch ("\n\t" + (((("hello" + "") + (("a" + "a") + ("foo bar" + "a"))) + ("hello" + (("\n\t" + ("a" + "\n\t")) + ("\n\t" + "\n\t")))) + "hello")) {
                    case "":
                        var v14 uint64 = 811
                        _ = v14
                    case "a":
                        var v15 uint64 = 4294967295
                        _ = v15
                    }
                    p1 = (p1 >> ((0) & 63))
                }
                _ = v5
            }
        }
    case "foo bar":
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 string = ""
            _ = fi16
            _ = v17
        }
    default:
        fmt.Println(p1, p1, p0)
    }
    if true {
        fmt.Println(p1)
    } else {
        v18 := []int{}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint, p1 int16, p2 string) int64 {
    var v19 bool = false
    var v20 uint32 = 2
    if (2 >= 2) {
        if v19 {
            v19 = (32768 >= 32768)
            switch -9223372036854775808 {
            case 4294967295:
                p2 = "hello"
            case 42:
                fmt.Println(p2)
            case -135:
                v21 := []int{32767}
                _ = v21
            default:
                v19 = (32767 != 0)
            }
        }
    } else {
        var v22 int64 = -1
        v19 = (v19 && (!((p1 - (p1 * p1)) < p1)))
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{641, -997}
    var v3 float32 = 0.1
    switch "foo bar" {
    case "a":
        v23 := GenS1{F0: 1}
        _ = v23
    case "foo bar":
        v24 := GenS1{F0: 18446744073709551615}
        _ = v24
    case "\n\t":
        var v25 uint32 = 42
        _ = v25
    default:
        if ((0 >= 115) || true) {
            fmt.Println(v3)
            v26 := GenS1{F0: 42}
            _ = v26
        } else {
            if true {
                v3 -= v3
                if false {
                    v3 = v3
                    v3 += v3
                    if true {
                        _ = len(v2)
                        var v27 int8 = 34
                        _ = v27
                    }
                }
                var v28 int = 127
                _ = v28
            } else {
                if (128 == 68) {
                    var v29 string = ("" + "a")
                    v29 = ("hello" + "\n\t")
                    _ = v29
                }
                switch "hello" {
                case "a":
                    _ = len(v2)
                case "":
                    v30 := GenS2{F0: genFn0(-1, 1)}
                    _ = v30
                case "hello":
                    var v31 bool = (((12 > 1) || (0 == 241)) && true)
                    _ = v31
                default:
                    var v32 uint64 = 956
                    _ = v32
                }
            }
            var v33 int = -1
            _ = v33
        }
    }
    var v34 uint32 = 2
    fmt.Println(v34, v3)
    fmt.Println(v34)
    if (18446744073709551615 >= 888) {
        switch "\n\t" {
        case "a":
            fmt.Println(v34, v34)
        case "foo bar":
            v34 -= 71
        }
        v3 = (1.5 / 0.1)
        v3 = (v3 / 0.1)
    } else {
        v35 := GenS3{F0: 128, F1: (!true), F2: 0}
        v34 = (v34 % ((v34) | 1))
        _ = v35
    }
    v36 := GenS1{F0: 1}
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 int8 = 101
        if false {
            v39 := GenS4{F0: fi37}
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    v3 -= 1e10
    var v40 uint = 0
    _ = len(v2)
    v41 := GenS5{F0: 2147483647, F1: 364, F2: 1}
    for fi42 := 0; fi42 < 12; fi42++ {
        fmt.Println(fi42, v3)
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v40
    _ = v41
}
